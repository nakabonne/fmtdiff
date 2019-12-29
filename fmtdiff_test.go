package fmtdiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
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
				LocalPrefix: localPrefix,
				Fragment:    true,
				TabWidth:    8,
				TabIndent:   true,
				FormatOnly:  false,
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
	_, _ = fmtdiff.Run("", nil)
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
	_, _ = fmtdiff.Run("", nil)
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
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f, err := Run(tc.filename, tc.options)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.expected, f)
		})
	}
}
