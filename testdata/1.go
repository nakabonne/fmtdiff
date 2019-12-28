package testdata

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/nakabonne/fmtreporter"
)

func _() {
	_, _ = fmtreporter.Run("", "")
	fmt.Println("fmt")
	pp.Println("pp")
}
