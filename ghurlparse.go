package ghurlparse

import (
	"regexp"
	"strings"
)

var (
	repoRegex = regexp.MustCompile(`(?:https|git)(?:://|@)github\.com[/:]([^/:#]+)/([^/#]*).*`)
)

// DestructureRepoURL returns the user and the repo name from a repository url.
func DestructureRepoURL(repoURL string) (valid bool, owner string, repoName string) {
	sanitisedRepoURL := strings.TrimSuffix(repoURL, ".git")
	submatches := repoRegex.FindStringSubmatch(sanitisedRepoURL)
	if len(submatches) < 3 {
		return
	}
	return true, submatches[1], submatches[2]
}
