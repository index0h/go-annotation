package annotation

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const maxGenerateImportAliasLevel = 1024

var upperCaseRegexp = regexp.MustCompile(`(^|\p{Ll})(\p{Lu})`)
var symbolRegexp = regexp.MustCompile(`(\p{S}|\p{P}|\p{C})+`)
var duplicateUnderscoreRegexp = regexp.MustCompile(`_+`)

type SafeImportGenerator struct {
	validator Validator
}

func NewSafeImportGenerator(validator Validator) *SafeImportGenerator {
	if validator == nil {
		panic(errors.New("Variable 'validator' must be not nil"))
	}

	return &SafeImportGenerator{
		validator: validator,
	}
}

func (g *SafeImportGenerator) Generate(dstFile *File, namespace string) *Import {
	if dstFile == nil {
		panic(errors.New("Variable 'dstFile' must be not nil"))
	}

	result := &Import{
		Namespace: namespace,
	}

	if err := g.validator.Validate(result); err != nil {
		panic(err)
	}

	existImports := map[string]string{}

	for _, importGroup := range dstFile.ImportGroups {
		for _, importEntity := range importGroup.Imports {
			if importEntity.Namespace == namespace {
				result.Alias = importEntity.Alias

				return result
			}

			existImports[importEntity.RealAlias()] = importEntity.Namespace
		}
	}

	if _, ok := existImports[result.RealAlias()]; !ok {
		return result
	}

	parts := strings.Split(result.Namespace, "/")

	for i := len(parts) - 1; i >= 0; i-- {
		alias := g.cleanAlias(strings.Join(parts[i:], "_"))

		if _, ok := existImports[alias]; !ok {
			result.Alias = alias

			return result
		}
	}

	for i := 0; i < maxGenerateImportAliasLevel; i++ {
		alias := g.cleanAlias(filepath.Base(namespace) + "_" + strconv.Itoa(i))

		if _, ok := existImports[alias]; !ok {
			result.Alias = alias

			return result
		}
	}

	panic(errors.Errorf("Can't generate alias for import %s", result.Namespace))
}

func (g *SafeImportGenerator) cleanAlias(alias string) string {
	alias = upperCaseRegexp.ReplaceAllString(alias, "${1}_${2}")
	alias = symbolRegexp.ReplaceAllString(alias, "_")
	alias = duplicateUnderscoreRegexp.ReplaceAllString(alias, "_")

	return strings.ToLower(strings.Trim(alias, "_"))
}
