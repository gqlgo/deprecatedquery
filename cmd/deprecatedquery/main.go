package main

import (
	"flag"
	"github.com/gqlgo/deprecatedquery"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	var excludes string
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL query field name. it can specify multiple values separated by `,`")
	flag.Parse()

	multichecker.Main(
		deprecatedquery.Analyzer(excludes),
	)
}
