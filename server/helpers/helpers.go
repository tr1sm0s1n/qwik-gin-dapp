package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
