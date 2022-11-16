package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	path, _ := config.Get("path")
	fmt.Println(path)
	enabled, _ := config.GetBool("enabled")
	fmt.Println(enabled)
}
