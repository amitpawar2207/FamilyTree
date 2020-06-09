package fileoperations

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

//ReadInputFile performs file operations
func ReadInputFile(filePath string) ([]string, error) {
	filePath = filepath.Join(filePath)

	if _, err := os.Stat(filePath); err != nil {
		return nil, fmt.Errorf("no such file or directory exists")
	}

	file, err := os.Open(filePath) // For read access.
	if err != nil {
		return nil, fmt.Errorf("error while opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		ok, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z_]+\s[a-zA-Z]+\s([a-zA-Z]+)|([a-zA-Z_]+\s[a-zA-Z]+)$`, text)
		if !ok {
			return nil, fmt.Errorf("Incorrect string. Contains either numbers or special character %s", text)
		} else {
			inputs = append(inputs, text)
		}
	}
	if len(inputs) == 0 {
		return inputs, fmt.Errorf("file is empty")
	}

	if err2 := scanner.Err(); err2 != nil {
		log.Fatal(err2)
		return nil, err2
	}

	return inputs, nil

}
