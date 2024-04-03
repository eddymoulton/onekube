package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/eddymoulton/onekube/internal/onepassword"
	"github.com/iancoleman/strcase"
)

func GetKubeConfigFilePath(name string) string {
	configDirectory := getConfigDirectory()
	return filepath.Join(configDirectory, strcase.ToKebab(name))
}

func Clean() {
	os.RemoveAll(getConfigDirectory())
}

func getConfigDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(home, ".config", "onekube")
}

func getConfigFilePath() string {
	configDirectory := getConfigDirectory()
	return filepath.Join(configDirectory, "configs")
}

func Read() ([]onepassword.Item, error) {
	itemsJson, err := os.ReadFile(getConfigFilePath())

	if err != nil {
		return nil, err
	}

	var items []onepassword.Item

	err = json.Unmarshal(itemsJson, &items)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return items, nil
}

func Write(items []onepassword.Item) error {
	itemsJson, _ := json.Marshal(items)

	err := os.WriteFile(getConfigFilePath(), []byte(itemsJson), 0644)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func EnsureDirectoryExists() error {
	configDirectory := getConfigDirectory()

	err := os.MkdirAll(configDirectory, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	return err
}