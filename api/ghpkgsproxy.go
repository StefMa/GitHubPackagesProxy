package api

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GitHubPackagesProxy(w http.ResponseWriter, r *http.Request) {
	packagePath := strings.TrimPrefix(r.URL.Path, "/")

	githubProxyUserName := os.Getenv("GITHUB_USERNAME")
	githubProxyPAT := os.Getenv("GITHUB_PAT")
	githubPackagesURL := "https://" + githubProxyUserName + ":" + githubProxyPAT + "@maven.pkg.github.com/" + packagePath

	targetURL, err := url.Parse(githubPackagesURL)
	fmt.Println(targetURL)
	if err != nil {
		http.Error(w, "Failed to parse GitHub Packages URL", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, targetURL.String(), http.StatusTemporaryRedirect)
}
