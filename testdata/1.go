package testdata

import (
	"fmt"

	"github.com/nakabonne/fmtreporter"
	"github.com/nakabonne/unusedparam/pkg/unusedparam"
)

func _() {
	_, _ = fmtreporter.Run("", "")
	fmt.Println("fmt")
	_, _ := unusedparam.Check(path)
}
