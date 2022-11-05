package main

import (
	entryhandler "clean-code-workshop/entryHandler"
	"log"
)

func main() {

	dir := "./test"

	duplicateIndex := entryhandler.NewDuplicateIndex()

	err := duplicateIndex.TraverseDirPath(dir)
	if err != nil {
		log.Fatal(err)
	}

	duplicateIndex.PrintDuplicateResult()
}
