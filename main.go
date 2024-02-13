// Alex Kefer - 2024 - Built for P2P Web Cache Project
// Downloads webpages and stores them in a cache for later retrieval
// Usage: go run . <URL>
// Example: go run . https://www.google.com
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <URL>")
		fmt.Println("Example: go run . https://www.google.com")
		return
	}
	url := os.Args[1]
	fmt.Println("Downloading: " + url)
	//Download Function Here
	html := downloadHTML(url)
	fmt.Println("Page Location: " + parsePageLocation(url))
	//Save Location Function Here
	html = findAndReplaceLinks(html, url)
	makeFileLocation("savedPages/" + parsePageLocation(url))
	fmt.Println("Page Name: " + parsePageName(url))
	savePage(html, parsePageName(url), parsePageLocation(url), ".html")
	//Save Page Function Here
}

func downloadHTML(url string) string {
	// takes in the URL and returns the MHTML of the page
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading webpage:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading webpage content:", err)
		return ""
	}
	html := string(body)
	return html
}
