package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	pb "agent-remote-svc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var requestID int64 = 0

func getNextRequestID() int64 {
	requestID++
	return requestID
}

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewScriptServiceClient(conn)

	// Create context with longer timeout and API key
	ctx := context.Background()

	// Add API key to context metadata
	if len(os.Args) < 2 {
		log.Fatal("API key must be provided as command line argument")
	}
	log.Printf("API key: %s", os.Args[1])
	ctx = metadata.AppendToOutgoingContext(ctx, "api-key", os.Args[1])

	// Create scanner for reading input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Agent Remote Service Client")
	fmt.Println("Enter commands (type 'exit' to quit):")
	fmt.Println("----------------------------------------")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		cmd := strings.TrimSpace(scanner.Text())

		if cmd == "exit" {
			break
		}

		if cmd == "" {
			continue
		}

		strs := strings.Split(cmd, " ")
		if len(strs) < 1 {
			fmt.Println("Invalid command")
			continue
		}
		method := strs[0]
		var paramsJSON string
		if len(strs) > 1 {
			paramsJSON = strings.Join(strs[1:], " ")
		}

		// Create RPC request with ID
		request := &pb.RPCRequest{
			Method: method,
			Params: paramsJSON,
			Id:     getNextRequestID(),
		}

		// Call ExecuteRPC and get stream with retry logic
		var stream pb.ScriptService_ExecuteRPCClient
		stream, err = client.ExecuteRPC(ctx, request)
		if err != nil {
			log.Printf("error calling ExecuteRPC: %v", err)
			continue
		}

		// Receive streaming responses
		fmt.Println("----------------------------------------")
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("----------------------------------------")
				break
			}
			if err != nil {
				if ctx.Err() == context.DeadlineExceeded {
					log.Printf("Stream timeout: %v", err)
					break
				}
				log.Printf("Error receiving response: %v", err)
				break
			}
			if resp.IsError {
				fmt.Printf("[ID:%d] Error: %s\n", resp.Id, resp.Output)
			} else {
				fmt.Printf("[ID:%d] %s\n", resp.Id, resp.Output)
			}
		}
	}
}
