package deprecatedquery_test

import (
	"testing"

	"github.com/gqlgo/deprecatedquery"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, deprecatedquery.Analyzer("excludeDeprecatedField"), "a")
}
