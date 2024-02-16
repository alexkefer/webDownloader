/* Function and helpers to save pages to file at specified location*/

package main

import (
	"fmt"
	"log"
	"os"
)

func makeFileLocation(urlLocation string) {
	// takes in the location of the file and creates the file location
	err := os.MkdirAll(urlLocation, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func saveAsset(content string, url string, saveLocation string, fileType string) {
	// takes in the content of the asset and saves it to the save location
	localPath := buildLocalPath("savedPages/"+saveLocation, url+fileType)
	println("Saving Asset: " + localPath)
	file, err := os.OpenFile(localPath, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err2 := file.WriteString(content)
	if err2 != nil {
		fmt.Println("Error writing to file")
	} else {
		fmt.Println("Successfully saved file")
	}
}

func savePage(context string, url string, saveLocation string, fileType string) {
	if checkIfHomePage(saveLocation) {
		url = "/index"
	}
	localPath := buildLocalPath("savedPages/"+saveLocation, url+fileType)
	println("Saving Asset: " + localPath)
	// takes in the context of the page and saves it to the save location
	file, err := os.OpenFile(localPath, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err2 := file.WriteString(context)
	if err2 != nil {
		fmt.Println("Error writing to file")
	} else {
		fmt.Println("Successfully saved file")
	}
}
