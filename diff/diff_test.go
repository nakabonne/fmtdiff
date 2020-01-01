package diff

import (
	"testing"

	godiff "github.com/sourcegraph/go-diff/diff"
	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	cases := []struct {
		name          string
		filename      string
		b1            []byte
		b2            []byte
		expectedHunks []*godiff.Hunk
		wantErr       bool
	}{
		{
			name:     "simple text",
			filename: "foo.txt",
			b1:       []byte("foo"),
			b2:       []byte("bar"),
			expectedHunks: []*godiff.Hunk{
				&godiff.Hunk{
					OrigStartLine:   1,
					OrigLines:       1,
					OrigNoNewlineAt: 5,
					NewStartLine:    1,
					NewLines:        1,
					Section:         "",
					StartPosition:   1,
					Body: []byte(`-foo
+bar`),
				},
			},
			wantErr: false,
		},
		{
			name:     "unified output format",
			filename: "foo.go",
			b1: []byte(`package foo

// ...
func foo() int {
  sum := 1 + 2
  return sum
}
`),
			b2: []byte(`package foo

// ...
func bar() int {
  sum := 1 + 2
  return sum
}
`),
			expectedHunks: []*godiff.Hunk{
				&godiff.Hunk{
					OrigStartLine:   1,
					OrigLines:       7,
					OrigNoNewlineAt: 0,
					NewStartLine:    1,
					NewLines:        7,
					Section:         "",
					StartPosition:   1,
					Body:            []byte(" package foo\n \n // ...\n-func foo() int {\n+func bar() int {\n   sum := 1 + 2\n   return sum\n }\n"),
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d, err := Diff(tc.b1, tc.b2, tc.filename)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.filename, d.NewName)
			assert.ElementsMatch(t, tc.expectedHunks, d.Hunks)
		})
	}
}

func TestReplaceTempFilename(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		arg      []byte
		expected []byte
		wantErr  bool
	}{
		{
			name:     "line of arg less than three",
			filename: "foo.txt",
			arg:      []byte("foo\nbar"),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := replaceTempFilename(tc.arg, tc.filename)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.expected, b)
		})
	}
}
