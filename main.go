package main

import (
	entryhandler "clean-code-workshop/entryHandler"
	"log"
)

func main() {

	dir := "./test"

	duplicateIndex := entryhandler.DuplicateIndex{
		Duplicates:          make(map[string]string),
		TotalUniqueFiles:    0,
		TotalDuplicateFiles: 0,
		TotalDuplicateSize:  0,
	}

	err := duplicateIndex.TraverseDirPath(dir)
	if err != nil {
		log.Fatal(err)
	}

	duplicateIndex.PrintDuplicateResult()
}
