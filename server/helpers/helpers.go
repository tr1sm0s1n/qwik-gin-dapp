package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func UpdateEnv(contractAddress string) {
	file, err := os.OpenFile(".env", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Failed to open .env file: %s\n\033[?25h", err.Error())
		return
	}
	defer file.Close()

	envContent := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "CONTRACT_ADDRESS") {
			modified := fmt.Sprintf("CONTRACT_ADDRESS=%s", contractAddress)
			envContent = append(envContent, modified)
		} else {
			envContent = append(envContent, line)
		}
	}

	file.Truncate(0)
	file.Seek(0, 0)
	writer := bufio.NewWriter(file)
	for _, line := range envContent {
		fmt.Fprintln(writer, line)
	}

	writer.Flush()
	fmt.Println("\033[35mDeployment details are saved to the '.env' file.\033[0m\033[?25h")
}

func ParseInt(input string) (int64, error) {
	output, err := strconv.ParseInt(input, 10, 64)
	return output, err
}

func Loading(done chan bool) {
	spinner := []string{"◳", "◲", "◱", "◰"}
	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\r\033[34mDeploying Contract... %s\033[0m\033[?25l", spinner[i])
			i = (i + 1) % len(spinner)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
