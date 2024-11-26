package hooks

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/99designs/gqlgen/plugin/modelgen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toCamelCase(str string) string {
	// First, convert to snake_case
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = strings.ToLower(snake)

	// Then convert to camelCase
	parts := strings.Split(snake, "_")
	for i := 1; i < len(parts); i++ {
		caser := cases.Title(language.English)
		parts[i] = caser.String(parts[i])
	}

	return strings.Join(parts, "")
}

var jsonTagRegexp = regexp.MustCompile(`json:".*?"`)
var jsonTagGroupRegexp = regexp.MustCompile(`json:"(.*?)"`)

func CamelCaseMutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			jsonTagGrouped := jsonTagGroupRegexp.FindStringSubmatch(field.Tag)
			camelCase := toCamelCase(jsonTagGrouped[1])
			field.Tag = jsonTagRegexp.ReplaceAllString(field.Tag, fmt.Sprintf(`json:"%s"`, camelCase))
		}
	}
	return b
}
