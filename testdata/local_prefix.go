package testdata

import (
	"fmt"

	"github.com/nakabonne/fmtdiff"
	"github.com/nakabonne/unusedparam/pkg/unusedparam"
)

func _() {
	_, _ = fmtdiff.Run("", nil)
	fmt.Println("fmt")
	_, _ = unusedparam.Check("")
}
