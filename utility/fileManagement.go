package utility

import (
	"os"
	"path/filepath"
)

func SaveFile(name string, data []byte) error {
	filepath := filepath.Join("data", name)
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadFile(name string) ([]byte, error) {
	filePath := filepath.Join("data", name)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil ,err
	}
	return file, nil
}
