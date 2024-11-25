package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/rtrampox/gqlgen-validation/plugin"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Println("failed to load config", err.Error())
		os.Exit(2)
	}

	if err := api.Generate(cfg, api.AddPlugin(plugin.New())); err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
}
