package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func Regex(input, pattern string) string {
	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)

	// Print the results
	for _, match := range matches {
		data := match[1] // match[1] contains the content inside the {}
		fmt.Println(data)
	}

	return ""
}

func ExtractContent(input string) (string, error) {
	var result strings.Builder
	stack := []rune{}
	found := false

	for _, char := range input {
		if char == '{' {
			if len(stack) == 0 {
				found = true // Found the outermost opening brace
			}
			stack = append(stack, char)
		} else if char == '}' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 && found {
				break // Found the outermost closing brace
			}
		}

		if found {
			result.WriteRune(char) // Append characters inside the outermost braces
		}
	}

	if len(stack) > 0 {
		return "", fmt.Errorf("mismatched braces")
	}

	abc := result.String()
	if len(abc) == 0 {
		return input, nil
	}

	if string(abc[len(abc)-1]) != "}" {
		abc += "}"
	}

	return abc, nil
}

func CreateFile(filename string, content []byte) error {
	// Specify the file name

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		// fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close() // Ensure the file is closed after we're done

	// Write some content to the file
	// content := []byte("Hello, World!\n")
	_, err = file.Write(content)
	if err != nil {
		// fmt.Println("Error writing to file:", err)
		return err
	}

	// fmt.Println(filename, "has been created!")
	return nil
}

func ReadFile(filename string) ([]byte, error) {
	_b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return _b, nil
}

func CMDWithStream(name string, args ...string) error {
	// Command to execute (e.g., "ping" command for demonstration)
	cmd := exec.Command(name, args...)

	// Get pipes for standard output and error
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return err
	}

	// Create a scanner for standard output
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println("-", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Println("Error reading stdout:", err)
		}
	}()

	// Create a scanner for standard error
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println("-", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Println("Error reading stderr:", err)
		}
	}()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func CopyFile(src string, dest string) error {
	// Open the source file
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	// Create the destination file
	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}

func PrintText(msg string, value interface{}) string {
	return fmt.Sprintf("  - %-30s %v \n", msg, value)
}

func PrintCommandText(msg string, value interface{}) string {
	return fmt.Sprintf("  - Press (%s) %20s %v \n", msg, "", value)
}

func CMDWithoutStOut(command string, args ...string) error {
	devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer devNull.Close()

	// Create a new command
	cmd := exec.Command(command, args...)

	// Redirect both stdout and stderr to devNull
	cmd.Stdout = devNull
	// cmd.Stderr = devNull

	// Run the command and return any error
	return cmd.Run()
}

func DockerComposeVersion() string {
	script := "docker compose"
	err := CMDWithoutStOut("bash", "-c", script)
	if err != nil {
		if strings.Contains(err.Error(), "is not a docker command") ||
			strings.Contains(err.Error(), "exit status 1") ||
			strings.Contains(err.Error(), "killed") {
			script = "docker-compose"
		}
	}
	return script
}

func IsAppleSilicon(platform string, env string) (bool, string) {
	dFName := "-f docker-compose%s.yml"
	dFName = fmt.Sprintf(dFName, env)
	ff := platform
	if strings.EqualFold(ff, PLATFROM_APPLE_SILLICON) {
		dFName = " -f docker-compose-arm%s.yml"
		dFName = fmt.Sprintf(dFName, env)
		return true, dFName
	}

	return false, dFName
}

func WeiToAmount(wei string) float64 {
	_b, _ := big.NewFloat(0).SetString(wei)
	_f := _b.Quo(_b, big.NewFloat(1e18))
	_f1, _ := _f.Float64()
	return _f1
}

func DockerCommand(serviceName string, dockerComposePath string, platform string, action string, env string) error {
	_, dFName := IsAppleSilicon(platform, env)
	script := fmt.Sprintf(
		`cd %s && \
     %s %s %s %s`,
		dockerComposePath, DockerComposeVersion(), dFName, action, serviceName)
	fmt.Println("====> docker script: ", script)
	err := CMDWithStream("bash", "-c", script)
	if err != nil {
		return err
	}
	return nil
}

func CurrentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return "./"
	}

	return currentDir
}

func IsInArray(array []string, val string) bool {
	for _, i := range array {
		if val == i {
			return true
		}
	}

	return false
}

func RunCommand(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, err
	}

	return output, nil
}

func RandomInRange(min, max int) int {
	// Generate a random integer in the range [min, max)
	return rand.Intn(max-min) + min
}

func ExecCommand(name string, args ...string) (*string, error) {
	cmd := exec.Command(name, args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := string(out.Bytes())
	return &result, nil
}
