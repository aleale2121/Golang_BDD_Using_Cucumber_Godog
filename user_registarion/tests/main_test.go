package tests

import (
	"flag"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"testing"
)

var opt = godog.Options{
	Format: "progress",
	Output: colors.Colored(os.Stdout),
}

func init() {
	godog.BindCommandLineFlags("godog", &opt)

}

func TestMain(m *testing.M) {
	flag.Parse()

	opt.Paths = flag.Args()

}
