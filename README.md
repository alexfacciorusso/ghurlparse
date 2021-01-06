# ghurlparse
Go library that parses a GitHub URL returning useful information, like the owner and the user of that repository.

All the parsing is done locally from the URL (the GitHub API is not used).

## Usage

```go
valid, owner, repoName := DestructureRepoURL("https://github.com/alexfacciorusso/ghurlparse")

if valid {
    fmt.Printf("Owner: %s, Repository name: %s", owner, repoName)
}
// Output: Owner: alexfacciorusso, Repository name: ghurlparse
```

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.
