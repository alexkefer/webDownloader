/* Utility functions to parse various file information out of the URL */

package main

import (
	"strings"
)

func checkIfHomePage(url string) bool {
	// takes in url and returns true if it is the homepage
	if len(url) <= 1 {
		return true
	}
	return url[len(url)-1] == '/' || !strings.Contains(url, "/")
}

func urlCleaner(url string) string {
	// takes in url and returns the cleaned url (removes http(s):// and www.)
	if len(url) >= 8 && url[:8] == "https://" {
		url = url[8:]
	}
	if len(url) >= 7 && url[:7] == "http://" {
		url = url[9:]
	}
	if len(url) >= 4 && url[:4] == "www." {
		url = url[4:]
	}
	return url
}

func parsePageLocation(url string) string {
	// takes in url and returns the location of the page
	url = urlCleaner(url)
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			url = url[:i]
			break
		}
	}
	return url
}

func parsePageName(url string) string {
	// takes in url and returns the name of the page
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			url = url[i+1:]
			break
		}
	}
	return url
}

// Function to build the local path for the asset, removes http, https, and www and places the asset in the correct location
func buildLocalPath(baseURL, assetURL string) string {
	// takes in a base url and an asset url and returns the full url
	if strings.HasPrefix(assetURL, "http://") || strings.HasPrefix(assetURL, "https://") {
		assetURL = assetURL[8:]
	}
	if strings.HasPrefix(assetURL, "www.") {
		assetURL = assetURL[4:]
	}
	if strings.HasPrefix(assetURL, "//") {
		assetURL = assetURL[2:]
	}
	if strings.HasPrefix(assetURL, "/") {
		assetURL = assetURL[1:]
	}
	return baseURL + "/" + assetURL
}

func determineAssetType(url string) string {
	if strings.HasSuffix(url, ".css") {
		return "css"
	} else if strings.HasSuffix(url, ".js") {
		return "js"
	} else if strings.Contains(url, "/images/") {
		return "img"
	}

	return "unknown"
}
