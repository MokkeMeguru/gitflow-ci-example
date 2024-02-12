package cmd_test

import (
	"create-release-pr-body/cmd"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateReleasePrBodyCmd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		prListFile string
		expected   string
	}{
		{
			name:       "create release pr body",
			prListFile: "../test/testdata/pr-list.txt",
			expected: `# Release PRs
## meguru-mokke
- [ ] #112
- [ ] #912
- [ ] #823
## takuya-ebata
- [ ] #938
- [ ] #935
- [ ] #931
- [ ] #923
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := os.ReadFile(tt.prListFile)
			require.NoError(t, err)
			prInfos := []cmd.PRInfo{}
			for _, line := range strings.Split(string(bytes), "\n") {
				raw := strings.Split(line, "|")
				if len(raw) < 2 {
					continue
				}
				author := raw[0]
				mergeMsg := raw[1]
				prHashNumber := cmd.PRHashNumberRegexp.FindString(mergeMsg)
				prInfos = append(prInfos, cmd.PRInfo{Author: author, PRHashNumber: prHashNumber})
			}
			acturl := cmd.CreateReleasePRBody(prInfos)
			require.Equal(t, tt.expected, acturl)
		})
	}
}
