package envreader

import (
	"bufio"
	"os"
	"strings"
)

// LoadFromFile reads variables from .env file and sets them as environment variables
// This reader is very basic, it only supports the format key = value, with no quotes, no multiline
// variables, no trailing comments.
func LoadFromFile() {
	// The .env is copied to the container during the build process
	// This is not best practise and will need to change and be set differently
	file, err := os.Open("/go/bin/.env")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), "=") {
			continue
		}
		keyValue := strings.Split(scanner.Text(), "=")
		err := os.Setenv(strings.TrimSpace(keyValue[0]), strings.TrimSpace(keyValue[1]))
		if err != nil {
			panic(err)
		}
	}
}

func CheckRequiredEnv() error {
	// todo: check if envs are set
	return nil
}
