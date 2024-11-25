package plugin

import (
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/v2/ast"
)

type generator struct{}

func New() plugin.Plugin {
	return &generator{}
}

func (g *generator) Name() string {
	return "gql_generator"
}

func (g *generator) MutateConfig(cfg *config.Config) error {
	cfg.Directives["binding"] = config.DirectiveConfig{SkipRuntime: false}
	return nil
}

func (g *generator) InjectSourceEarly() *ast.Source {
	return &ast.Source{
		Name:    "validationDirective.graphqls",
		BuiltIn: false,
		Input:   `@binding(constraint: String!) on FIELD_DEFINITION | ARGUMENT_DEFINITION`,
	}
}
