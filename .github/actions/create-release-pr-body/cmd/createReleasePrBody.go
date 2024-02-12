package cmd

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var (
	PRHashNumberRegexp = regexp.MustCompile(`#[0-9]*`)
)

var createReleasePrBodyCmd = &cobra.Command{
	Use:   "createReleasePrBody",
	Short: "Create Release PR Body",
	Long: `Create Release PR Body like this:

# Release PRs
## Author1
- [ ] #12
- [ ] #13
## Author2
- [ ] #14
- [ ] #15
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := os.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		prInfos := []PRInfo{}
		for _, line := range strings.Split(string(bytes), "\n") {
			raw := strings.Split(line, "|")
			if len(raw) < 2 {
				continue
			}
			author := raw[0]
			mergeMsg := raw[1]
			prHashNumber := PRHashNumberRegexp.FindString(mergeMsg)
			prInfos = append(prInfos, PRInfo{Author: author, PRHashNumber: prHashNumber})
		}
		body := CreateReleasePRBody(prInfos)
		fmt.Println(body)
	},
}

type PRInfo struct {
	Author       string `json:"author"`
	PRHashNumber string `json:"prHashNumber"`
}

func CreateReleasePRBody(prInfos []PRInfo) string {
	authorMap := map[string][]string{}
	for _, prInfo := range prInfos {
		authorMap[prInfo.Author] = append(authorMap[prInfo.Author], prInfo.PRHashNumber)
	}
	return `# Release PRs
` + CreateReleasePRBodySections(authorMap)
}

func CreateReleasePRBodySections(authorMap map[string][]string) string {
	body := ""
	sortAuthors := []string{}
	for author := range authorMap {
		sortAuthors = append(sortAuthors, author)
	}
	sort.SliceStable(sortAuthors, func(i, j int) bool {
		return sortAuthors[i] < sortAuthors[j]
	})
	for _, author := range sortAuthors {
		body = body + fmt.Sprintf("## %s\n", author)
		for _, prHashNumber := range authorMap[author] {
			body = body + fmt.Sprintf("- [ ] %s\n", prHashNumber)
		}
	}
	return body
}

func init() {
	rootCmd.AddCommand(createReleasePrBodyCmd)
}
