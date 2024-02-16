/* This is a helper utility built to regex through the html and modify the locations to where they are downloaded rather than their links */

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

/*
Legacy code that was replaced by the tokenizer
func findAndReplaceLinks(html string, url string) string {
	// takes in the html and the url and returns the html with the links modified
	links := findAssets(html)
	for i := 0; i < len(links); i++ {
		if links[i][1][0] != '#' {
			link := buildPageUrl(url, links[i][1])
			assetInfo := retrieveAsset(link)
			if assetInfo != "" {
				makeFileLocation("savedPages/" + parsePageLocation(link))
				saveAsset(assetInfo, parsePageName(links[i][1]), parsePageLocation(links[i][1]), "")
				print("Retrieving Asset: " + link + "\n")
			}
		}
	}
	return html
}

func findAssets(html string) [][]string {
	// takes in the html and returns the links to the assets in the html
	regex := regexp.MustCompile(`src="([^"]+)"`)
	return regex.FindAllStringSubmatch(html, -1)
}*/

func retrieveAsset(url string) string {
	// takes in the url and returns the asset
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "P2PWebCache")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error downloading asset:", url, err)
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

func DownloadAllAssets(url, htmlContent string) string {
	tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return htmlContent
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "link": // Download CSS
				for _, attr := range token.Attr {
					if attr.Key == "rel" && strings.Contains(attr.Val, "stylesheet") {
						if href, ok := getAttributeValue(token, "href"); ok {
							link := buildPageUrl(url, href)
							fmt.Println("Retrieving Asset: " + link)
							assetInfo := retrieveAsset(link)
							if assetInfo != "" {
								makeFileLocation("savedPages/" + parsePageLocation(link))
								href = trimLongURL(href)
								var fileType string
								if strings.HasSuffix(href, ".css") {
									fileType = ""
								} else {
									fileType = ".css"
								}
								saveAsset(assetInfo, parsePageName(href), parsePageLocation(link), fileType)
							}
						}
					}
				}
			case "script": // Download JS
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						link := buildPageUrl(url, attr.Val)
						fmt.Println("Retrieving Asset: " + link)
						assetInfo := retrieveAsset(link)
						if assetInfo != "" {
							makeFileLocation("savedPages/" + parsePageLocation(link))
							var fileType string
							if strings.HasSuffix(link, ".js") {
								fileType = ""
							} else {
								fileType = ".js"
							}
							saveAsset(assetInfo, parsePageName(attr.Val), parsePageLocation(link), fileType)
						}
					}
				}
			}
		}
	}
}

/* Helper Functions */

func getAttributeValue(token html.Token, key string) (string, bool) {
	for _, attr := range token.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func trimLongURL(url string) string {
	// takes in url and returns the trimmed url
	if len(url) > 50 {
		return url[:50]
	}
	return url
}

func buildPageUrl(url string, assetUrl string) string {
	// takes in the source url and the asset url and returns the full url
	if strings.HasPrefix(assetUrl, "http://") || strings.HasPrefix(assetUrl, "https://") {
		return assetUrl
	}
	if strings.HasPrefix(assetUrl, "//") {
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
			return "https://" + url[:i]
		}
	}
	return "https://" + url
}
