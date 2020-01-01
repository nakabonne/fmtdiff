package fmtdiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	localPrefix := "github.com/nakabonne/fmtdiff"
	cases := []struct {
		name     string
		filename string
		options  *Options
		expected *FileDiff
		wantErr  bool
	}{
		{
			name:     "local package is grouped into 3rd-party packages",
			filename: "testdata/local_prefix.go",
			options: &Options{
				LocalPrefixes: []string{localPrefix},
			},
			expected: &FileDiff{
				Name: "testdata/local_prefix.go",
				Before: []byte(`package testdata

import (
	"fmt"

	"github.com/nakabonne/fmtdiff"
	"github.com/nakabonne/unusedparam/pkg/unusedparam"
)

func _() {
	_, _ = fmtdiff.Process("", nil)
	fmt.Println("fmt")
	_, _ = unusedparam.Check("")
}
`),
				After: []byte(`package testdata

import (
	"fmt"

	"github.com/nakabonne/unusedparam/pkg/unusedparam"

	"github.com/nakabonne/fmtdiff"
)

func _() {
	_, _ = fmtdiff.Process("", nil)
	fmt.Println("fmt")
	_, _ = unusedparam.Check("")
}
`),
				Hunks: []*Hunk{
					&Hunk{
						OrigStartLine: 3,
						OrigLines:     8,
						NewStartLine:  3,
						NewLines:      9,
						Body: []byte(` import (
 	"fmt"
 
-	"github.com/nakabonne/fmtdiff"
 	"github.com/nakabonne/unusedparam/pkg/unusedparam"
+
+	"github.com/nakabonne/fmtdiff"
 )
 
 func _() {
`),
					},
				},
			},
			wantErr: false,
		},
		{
			name:     "set multiple prefixes as local packages",
			filename: "testdata/local_prefix.go",
			options: &Options{
				LocalPrefixes: []string{localPrefix, "github.com/nakabonne/unusedparam"},
			},
			expected: &FileDiff{
				Name: "testdata/local_prefix.go",
				Before: []byte(`package testdata

import (
	"fmt"

	"github.com/nakabonne/fmtdiff"
	"github.com/nakabonne/unusedparam/pkg/unusedparam"
)

func _() {
	_, _ = fmtdiff.Process("", nil)
	fmt.Println("fmt")
	_, _ = unusedparam.Check("")
}
`),
				After: []byte(`package testdata

import (
	"fmt"

	"github.com/nakabonne/fmtdiff"
	"github.com/nakabonne/unusedparam/pkg/unusedparam"
)

func _() {
	_, _ = fmtdiff.Process("", nil)
	fmt.Println("fmt")
	_, _ = unusedparam.Check("")
}
`),
				Hunks: nil,
			},
			wantErr: false,
		},
		{
			name:     "no diff",
			filename: "testdata/fmted.go",
			options:  &Options{},
			expected: &FileDiff{
				Name: "testdata/fmted.go",
				Before: []byte(`package testdata

func _(m, n int) {
	return
}

/*


 */

func _() {
	_ = 1
}
`),
				After: []byte(`package testdata

func _(m, n int) {
	return
}

/*


 */

func _() {
	_ = 1
}
`),
				Hunks: nil,
			},
			wantErr: false,
		},
		{
			name:     "no options",
			filename: "testdata/fmted.go",
			options:  nil,
			expected: &FileDiff{
				Name: "testdata/fmted.go",
				Before: []byte(`package testdata

func _(m, n int) {
	return
}

/*


 */

func _() {
	_ = 1
}
`),
				After: []byte(`package testdata

func _(m, n int) {
	return
}

/*


 */

func _() {
	_ = 1
}
`),
				Hunks: nil,
			},
			wantErr: false,
		},
		{
			name:     "wrong file path specified",
			filename: "xxx/yyy.go",
			options:  nil,
			wantErr:  true,
		},
		{
			name:     "not go source file given",
			filename: "testdata/empty.txt",
			options:  nil,
			wantErr:  true,
		},
		{
			name:     "emtpy go source file given",
			filename: "testdata/empty.go",
			options:  nil,
			wantErr:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f, err := Process(tc.filename, tc.options)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.expected, f)
		})
	}
}
