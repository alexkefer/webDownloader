/* This is a helper utility built to regex through the html and modify the locations to where they are downloaded rather than their links */

package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func findAndReplaceLinks(html string, url string) string {
	// takes in the html and the url and returns the html with the links modified
	links := findAssets(html)
	for i := 0; i < len(links); i++ {
		if links[i][1][0] != '#' {
			link := buildPageUrl(url, links[i][1])
			retrieveAsset(link)
			print("Retrieving Asset: " + link + "\n")
		}
	}
	return html
}

func findAssets(html string) [][]string {
	// takes in the html and returns the links to the assets in the html
	regex := regexp.MustCompile(`src="([^"]+)"`)
	return regex.FindAllStringSubmatch(html, -1)
}

func retrieveAsset(url string) string {
	// takes in the url and returns the asset
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading asset:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading asset content:", err)
		return ""
	}
	return string(body)
}

func buildPageUrl(url string, assetUrl string) string {
	// takes in the source url and the asset url and returns the full url
	if strings.HasPrefix(assetUrl, "http://") || strings.HasPrefix(assetUrl, "https://") {
		return assetUrl
	} else if strings.HasPrefix(assetUrl, "//") {
		assetUrl = "https:" + assetUrl
	} else if assetUrl[0] == '/' {
		assetUrl = parsePageSource(parsePageLocation(url)) + assetUrl
	} else {
		for i := 0; i < len(assetUrl); i++ {
			if assetUrl[i] == '.' {
				assetUrl = "https://" + assetUrl
				break
			}
			if assetUrl[i] == '/' {
				assetUrl = parsePageSource(parsePageLocation(url)) + assetUrl
				break
			}
		}
	}
	return assetUrl
}

func parsePageSource(url string) string {
	// takes in the url and returns only the https://www.{url}
	for i := 0; i < len(url); i++ {
		if url[i] == '/' {
			return url[:i]
		}
	}
	return "https://" + url
}
