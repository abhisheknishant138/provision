package cli

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func TestLoadIncrementer(t *testing.T) {
	 cliTest(false, false,
		"plugin_providers", "upload", "incrementer", "from", path.Join("../bin", runtime.GOOS, runtime.GOARCH, "incrementer")).run(t)
	fmt.Println("passsed 1")
	cliTest(false, false, "plugin_providers", "list").run(t)
	fmt.Println("passsed 2")
	cliTest(false, false, "plugin_providers", "show", "incrementer").run(t)
	fmt.Println("passsed 3")
}
