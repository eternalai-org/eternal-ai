package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"agent-remote-svc/cmd/server/discover"
	pb "agent-remote-svc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

var (
	port        = ":50051"
	validAPIKey string
)

func init() {
	if len(os.Args) < 2 {
		log.Fatal("API key must be provided as command line argument")
	}
	validAPIKey = os.Args[1]
	log.Printf("API key: %s", validAPIKey)
}

type server struct {
	pb.UnimplementedScriptServiceServer
}

// ScriptParams represents the parameters for script execution
type ScriptParams struct {
	Script string `json:"script"`
}

// UploadFileParams represents the parameters for file upload
type UploadFileParams struct {
	FileName       string `json:"file_name"`
	FileDataBase64 string `json:"file_data_base64"`
}

// DownloadFileParams represents the parameters for file download
type DownloadFileParams struct {
	FileUrl  string `json:"file_url"`
	FileName string `json:"file_name"`
}

// DockerParams represents the parameters for docker command
type DockerParams struct {
	Command string `json:"command"`
}

type ProgressReader struct {
	OnProgress func(pos int64, size int64, percent int64)
	Reader     io.Reader
	Size       int64
	Pos        int64
	Percent    int64
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	if err == nil {
		pr.Pos += int64(n)
		percent := pr.Pos * 100 / pr.Size
		if percent != pr.Percent {
			pr.Percent = percent
			if pr.OnProgress != nil {
				pr.OnProgress(pr.Pos, pr.Size, percent)
			}
		}
	}
	return n, err
}

func (s *server) ExecuteRPC(req *pb.RPCRequest, stream pb.ScriptService_ExecuteRPCServer) error {
	// get request ip
	p, ok := peer.FromContext(stream.Context())
	if !ok || p == nil {
		return s.responseError(stream, req, status.Error(codes.Internal, "failed to get peer info"))
	}
	ip, _, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to get request ip: %v", err)))
	}
	fmt.Printf("ExecuteRPC: %s %v\n", ip, req)

	// Get API key from metadata
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return s.responseError(stream, req, status.Error(codes.Unauthenticated, "missing metadata"))
	}

	apiKeys := md.Get("api-key")
	if len(apiKeys) == 0 || apiKeys[0] != validAPIKey {
		stream.Send(&pb.RPCResponse{
			Id:      req.Id,
			Method:  req.Method,
			Output:  "invalid API key",
			IsError: true,
		})
		return s.responseError(stream, req, status.Error(codes.Unauthenticated, "invalid API key"))
	}

	// Handle different methods
	switch req.Method {
	case "ping":
		{
			stream.Send(&pb.RPCResponse{
				Id:      req.Id,
				Method:  req.Method,
				Output:  "pong",
				IsError: false,
			})
		}
	case "execute_script":
		{
			var params ScriptParams
			if err := json.Unmarshal([]byte(req.Params), &params); err != nil {
				return s.responseError(stream, req, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid params: %v", err)))
			}

			// Create a shell command
			cmd := exec.Command("sh", "-c", params.Script)

			// Get stdout pipe
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to create stdout pipe: %v", err)))
			}
			// Get stderr pipe
			stderr, err := cmd.StderrPipe()
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to create stderr pipe: %v", err)))
			}

			// Start the command
			if err := cmd.Start(); err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to start command: %v", err)))
			}

			// Create a WaitGroup to track goroutines
			var wg sync.WaitGroup
			wg.Add(2)

			// Create a goroutine to read stdout
			go func() {
				defer wg.Done()
				reader := bufio.NewReader(stdout)
				for {
					line, err := reader.ReadString('\n')
					if err != nil {
						if err != io.EOF {
							log.Printf("Error reading stdout: %v", err)
						}
						return
					}
					stream.Send(&pb.RPCResponse{
						Id:      req.Id,
						Method:  req.Method,
						Output:  strings.TrimSpace(line),
						IsError: false,
					})
				}
			}()

			// Create a goroutine to read stderr
			go func() {
				defer wg.Done()
				reader := bufio.NewReader(stderr)
				for {
					line, err := reader.ReadString('\n')
					if err != nil {
						if err != io.EOF {
							log.Printf("Error reading stderr: %v", err)
						}
						return
					}
					stream.Send(&pb.RPCResponse{
						Id:      req.Id,
						Method:  req.Method,
						Output:  strings.TrimSpace(line),
						IsError: true,
					})
				}
			}()

			// Wait for the command to finish
			if err := cmd.Wait(); err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("command failed: %v", err)))
			}

			// Wait for all goroutines to finish
			wg.Wait()
		}

	case "get_upload_dir":
		{
			currentDir, err := os.Getwd()
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to get current directory: %v", err)))
			}

			stream.Send(&pb.RPCResponse{
				Id:      req.Id,
				Method:  req.Method,
				Output:  currentDir + "/uploads",
				IsError: false,
			})

			return s.responseError(stream, req, status.Error(codes.OK, currentDir+"/uploads"))
		}

	case "upload_file":
		{
			var params UploadFileParams
			if err := json.Unmarshal([]byte(req.Params), &params); err != nil {
				return s.responseError(stream, req, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid params: %v", err)))
			}

			// Decode base64 file data
			fileData, err := base64.StdEncoding.DecodeString(params.FileDataBase64)
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode file data: %v", err)))
			}

			// Create directory if it doesn't exist
			currentDir, err := os.Getwd()
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to get current directory: %v", err)))
			}
			err = os.MkdirAll(currentDir+"/uploads/", 0755)
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to create directory: %v", err)))
			}

			// Create file path
			filePath := currentDir + "/uploads/" + strings.TrimPrefix(params.FileName, "/")
			// Create directory if it doesn't exist
			fileDir := filepath.Dir(filePath)
			// check if directory exists
			if _, err := os.Stat(fileDir); os.IsNotExist(err) {
				err = os.MkdirAll(fileDir, 0755)
				if err != nil {
					return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to create directory: %v", err)))
				}
			}
			err = os.WriteFile(filePath, fileData, 0644)
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to write file: %v", err)))
			}

			stream.Send(&pb.RPCResponse{
				Id:      req.Id,
				Method:  req.Method,
				Output:  filePath,
				IsError: false,
			})
		}

	case "download_file":
		{
			var params DownloadFileParams
			if err := json.Unmarshal([]byte(req.Params), &params); err != nil {
				return s.responseError(stream, req, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid params: %v", err)))
			}

			// Download the file from the URL
			resp, err := http.Get(params.FileUrl)
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to download file: %v", err)))
			}
			defer resp.Body.Close()

			// Create directory if it doesn't exist
			currentDir, err := os.Getwd()
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to get current directory: %v", err)))
			}
			err = os.MkdirAll(currentDir+"/uploads/", 0755)
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to create directory: %v", err)))
			}

			// Create file path
			filePath := currentDir + "/uploads/" + strings.TrimPrefix(params.FileName, "/")
			// Create directory if it doesn't exist
			fileDir := filepath.Dir(filePath)
			// check if directory exists
			if _, err := os.Stat(fileDir); os.IsNotExist(err) {
				err = os.MkdirAll(fileDir, 0755)
				if err != nil {
					return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to create directory: %v", err)))
				}
			}

			// Create a new file
			file, err := os.Create(filePath)
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to create file: %v", err)))
			}

			// Create a progress reader
			progressReader := &ProgressReader{
				Reader: resp.Body,
				Size:   resp.ContentLength,
				OnProgress: func(pos int64, size int64, percent int64) {
					stream.Send(&pb.RPCResponse{
						Id:      req.Id,
						Method:  req.Method,
						Output:  fmt.Sprintf("Downloading file %d bytes of %d (%d%%)", pos, size, percent),
						IsError: false,
					})
				},
			}
			if _, err := io.Copy(file, progressReader); err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to copy file: %v", err)))
			}

			// Close the file
			err = file.Close()
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to close file: %v", err)))
			}

			stream.Send(&pb.RPCResponse{
				Id:      req.Id,
				Method:  req.Method,
				Output:  filePath,
				IsError: false,
			})
		}

	case "docker":
		{
			var params DockerParams
			if err := json.Unmarshal([]byte(req.Params), &params); err != nil {
				return s.responseError(stream, req, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid params: %v", err)))
			}

			commands := strings.Split(params.Command, " ")
			if len(commands) == 0 {
				return s.responseError(stream, req, status.Error(codes.InvalidArgument, "command is required"))
			}

			// Create a shell command
			cmd := exec.Command("docker", commands...)

			// Get stdout and stderr pipes
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to create stdout pipe: %v", err)))
			}

			stderr, err := cmd.StderrPipe()
			if err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to create stderr pipe: %v", err)))
			}

			// Start the command
			if err := cmd.Start(); err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("failed to start command: %v", err)))
			}

			// Create a WaitGroup to track goroutines
			var wg sync.WaitGroup
			wg.Add(2)

			// Create a goroutine to read stdout
			go func() {
				defer wg.Done()
				reader := bufio.NewReader(stdout)
				for {
					line, err := reader.ReadString('\n')
					if err != nil {
						if err != io.EOF {
							log.Printf("Error reading stdout: %v", err)
						}
						return
					}
					stream.Send(&pb.RPCResponse{
						Id:      req.Id,
						Method:  req.Method,
						Output:  strings.TrimSpace(line),
						IsError: false,
					})
				}
			}()

			// Create a goroutine to read stderr
			go func() {
				defer wg.Done()
				reader := bufio.NewReader(stderr)
				for {
					line, err := reader.ReadString('\n')
					if err != nil {
						if err != io.EOF {
							log.Printf("Error reading stderr: %v", err)
						}
						return
					}
					stream.Send(&pb.RPCResponse{
						Id:      req.Id,
						Method:  req.Method,
						Output:  strings.TrimSpace(line),
						IsError: true,
					})
				}
			}()

			// Wait for the command to finish
			if err := cmd.Wait(); err != nil {
				return s.responseError(stream, req, status.Error(codes.Internal, fmt.Sprintf("command failed: %v", err)))
			}

			// Wait for all goroutines to finish
			wg.Wait()
		}

	default:
		return s.responseError(stream, req, status.Error(codes.Unimplemented, fmt.Sprintf("method %s not implemented", req.Method)))
	}

	return nil
}

func (s *server) responseError(stream pb.ScriptService_ExecuteRPCServer, req *pb.RPCRequest, err error) error {
	stream.Send(&pb.RPCResponse{
		Id:      req.Id,
		Method:  req.Method,
		Output:  err.Error(),
		IsError: true,
	})
	return err
}

func main() {
	// Start the discovery server
	go discover.DiscoverStart()

	// Start the RPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterScriptServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
