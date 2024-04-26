package utility

import (
	"encoding/json"
	"os"
	"path/filepath"
	"fmt"
)

var Users = map[string]Properties{}

type Properties struct {
	Color string `json:"color"`
}

func LoadUserConfig() (map[string]Properties, error) {
	data, err := loadFile("userconfig.json")
	if err != nil {
		return nil, err
	}
	var users map[string]Properties
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserConfig() (map[string]Properties, error) {
	if len(Users) == 0 {
		users, err := LoadUserConfig()
		if err != nil {fmt.Printf("Error loading user config: %v", err)}
		Users = users
		return Users, nil
	}
	return Users, nil
}

func SaveUserConfig(users map[string]Properties) error {
	Users = users
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	if err := saveFile("userconfig.json", data); err != nil {return err}
	return nil
}

func saveFile(name string, data []byte) error {
	filepath := filepath.Join("data", name)
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func loadFile(name string) ([]byte, error) {
	filePath := filepath.Join("data", name)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil ,err
	}
	return file, nil
}
