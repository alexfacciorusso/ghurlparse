package ghurlparse

import (
	"fmt"
)

func ExampleDestructureRepoURL() {
	valid, owner, repoName := DestructureRepoURL("https://github.com/alexfacciorusso/ghurlparse")

	if valid {
		fmt.Printf("Owner: %s, Repository name: %s", owner, repoName)
	}
	// Output: Owner: alexfacciorusso, Repository name: ghurlparse
}
