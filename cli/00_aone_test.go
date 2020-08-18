package cli

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func TestLoadIncrementer(t *testing.T) {
	cc:= cliTest(false, false,
		"plugin_providers", "upload", "incrementer", "from", path.Join("../bin", runtime.GOOS, runtime.GOARCH, "incrementer")).run(t)
	fmt.Println("passsed 1")
	fmt.Println (cc)
	tt:= cliTest(false, false, "plugin_providers", "list").run(t)
	fmt.Println("passsed 2")
	fmt.Println (tt)
	rr:= cliTest(false, false, "plugin_providers", "show", "incrementer").run(t)
	fmt.Println("passsed 3")
	fmt.Println (rr)
}
