package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/rtrampox/gqlgen-validation/hooks"
	"github.com/rtrampox/gqlgen-validation/plugin"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Println("failed to load config", err.Error())
		os.Exit(2)
	}

	p := modelgen.Plugin{
		MutateHook: hooks.CamelCaseMutateHook,
	}

	err = api.Generate(cfg, api.AddPlugin(plugin.New()), api.AddPlugin(&p))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
}
