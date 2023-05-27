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
		fmt.Printf("Failed to open .env file: %s\n", err.Error())
		return
	}
	defer file.Close()

	envContent := make(map[string]string)
	comments := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			comments = append(comments, line)
		} else {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				envContent[parts[0]] = parts[1]
			}
		}
	}

	envContent["CONTRACT_ADDRESS"] = contractAddress

	file.Truncate(0)
	file.Seek(0, 0)
	writer := bufio.NewWriter(file)
	for _, comment := range comments {
		fmt.Fprintln(writer, comment)
	}
	for key, value := range envContent {
		fmt.Fprintf(writer, "%s=%s\n", key, value)
	}
	writer.Flush()

	fmt.Println(".env updated.")
}
