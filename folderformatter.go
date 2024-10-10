package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {

	// get the current directory name, throw if there is an error doing so
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v\n", err)
	}

	directoryObjects, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatalf("Error getting current directory: %v\n", err)
	}

	// for each object in the current directory, if the object is a folder we want to rename it so there are not any whitespaces
	for _, object := range directoryObjects {

		if object.IsDir() {

			// if the directory contains a blank rune(character), get the individual words in the file name
			if strings.ContainsRune(object.Name(), ' ') {

				indivFolderWords := strings.Fields(object.Name())

				processedWordsList := []string{}

				// for each word in the indiv works capitalize the first letter
				for _, word := range indivFolderWords {
					capLetter := string(unicode.ToUpper(rune(word[0]))) + word[1:]
					processedWordsList = append(processedWordsList, capLetter)
				}

				// create the new file name by capitalizing the words and combining them with a _
				newFileName := strings.Join(processedWordsList, "_")

				// rename the file
				err := os.Rename(object.Name(), newFileName)

				if err != nil {
					log.Fatalf("Error renaming directory %v\n", err)
				}

				// if the folder already contains underscores, then split the words using split vs fields since there is no whitespaces
			} else if strings.ContainsRune(object.Name(), '_') {

				indivFolderWords := strings.Split(object.Name(), "_")

				processedWordsList := []string{}

				// for each word in the indiv works capitalize the first letter
				for _, word := range indivFolderWords {
					capLetter := string(unicode.ToUpper(rune(word[0]))) + word[1:]
					processedWordsList = append(processedWordsList, capLetter)
				}

				// create the new file name by capitalizing the words and combining them with a _
				newFileName := strings.Join(processedWordsList, "_")

				// rename the file
				err := os.Rename(object.Name(), newFileName)

				if err != nil {
					log.Fatalf("Error renaming directory %v\n", err)
				}

			}
		}
	}
}
