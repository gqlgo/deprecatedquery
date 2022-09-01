# deprecatedquery

[![pkg.go.dev][gopkg-badge]][gopkg]

`deprecatedquery` finds a deprecated query in your GraphQL query files.

```graphql
# In your schema
type User {
    id: ID!
    addr: String! @deprecated(reason: "use address instead")
    address: String!
}

# Query
query user {
    user {
        addr # want "addr is deprecated reason: use address instead"
    }
}
```

## How to use

A runnable linter can be created with multichecker package.
You can create own linter with your favorite Analyzers.

```go
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
```

`deprecatedquery` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/gqlgo/deprecatedquery/cmd/deprecatedquery@latest
```

The `deprecatedquery` command has two flags, `schema` and `query` which will be parsed and analyzed by deprecatedquery's Analyzer.

```sh
$ deprecatedquery -schema="server/graphql/schema/**/*.graphql" -query="client/**/*.graphql" -excludes="field1,field2"
```

The default value of `schema` is "schema/*/**.graphql" and `query` is `query/*/**.graphql`.

## Author

[![Appify Technologies, Inc.](appify-logo.png)](http://github.com/appify-technologies)

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/deprecatedquery
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/deprecatedquery?status.svg
