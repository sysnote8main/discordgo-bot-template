package util

import (
	"os"
)

func WriteFile(path, data string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	b := []byte(data)
	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) (*string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	d := string(f)
	return &d, nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
