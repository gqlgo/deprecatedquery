package deprecatedquery

import (
	"fmt"
	"github.com/gqlgo/gqlanalysis"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

var excludes []string

func isExcludeField(fieldName string) bool {
	for _, exclude := range excludes {
		if exclude == fieldName {
			return true
		}
	}
	return false
}

func Analyzer(excludes string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "deprecatedquery",
		Doc:  "deprecatedquery finds a deprecatedquery in your GraphQL query files",
		Run:  run(excludes),
	}
}

func run(excludesStr string) func(pass *gqlanalysis.Pass) (interface{}, error) {
	excludes = strings.Split(excludesStr, ",")

	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		for _, q := range pass.Queries {
			checkOperations(pass, q.Operations)
			checkFragments(pass, q.Fragments)
		}

		return nil, nil
	}
}

func checkOperations(pass *gqlanalysis.Pass, ops ast.OperationList) {
	for _, op := range ops {
		for _, s := range op.SelectionSet {
			checkField(pass, s)
			checkInlineFragment(pass, s)
		}
	}
}

func checkFragments(pass *gqlanalysis.Pass, fs ast.FragmentDefinitionList) {
	for _, f := range fs {
		checkFragment(pass, f)
	}
}

func checkFragment(pass *gqlanalysis.Pass, f *ast.FragmentDefinition) {
	for _, s := range f.SelectionSet {
		checkField(pass, s)
	}
}

func checkInlineFragment(pass *gqlanalysis.Pass, sel ast.Selection) {
	f, _ := sel.(*ast.InlineFragment)
	if f == nil {
		return
	}

	for _, s := range f.SelectionSet {
		checkField(pass, s)
		checkInlineFragment(pass, s)
	}
}

func checkField(pass *gqlanalysis.Pass, sel ast.Selection) {
	field, _ := sel.(*ast.Field)
	if field == nil {
		return
	}

	for _, s := range field.SelectionSet {
		checkField(pass, s)
		checkInlineFragment(pass, s)
	}

	if isExcludeField(field.Name) {
		return
	}

	if field.Definition == nil {
		return
	}

	if deprecatedDirective := field.Definition.Directives.ForName("deprecated"); deprecatedDirective != nil {
		msg := fmt.Sprintf("%v is deprecated", field.Name)
		if reason, ok := deprecatedDirective.ArgumentMap(nil)["reason"]; ok {
			msg += fmt.Sprintf(" reason: %v", reason)
		}
		pass.Reportf(field.Position, msg)
	}
}
