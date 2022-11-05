package entryhandler

import (
	"clean-code-workshop/utils"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"sync/atomic"
)

func (index *DuplicateIndex) TraverseDirPath(directory string) error {

	entries, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryHandler := NewEntryHandler(entry, directory)
		entryHandler.Handle(index)
	}
	return nil

}

func createFileHash(fullPath string) (string, error) {

	file, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return "", err
	}

	hash := sha1.New()
	if _, err := hash.Write(file); err != nil {
		return "", err
	}

	hashSum := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashSum)

	return hashString, nil

}

func (index *DuplicateIndex) AddHashEntry(fullPath string, fileSize int64, hash string) {

	if hashEntry, ok := index.Duplicates[hash]; ok {
		index.Duplicates[hashEntry] = fullPath
		index.TotalDuplicateFiles++
		atomic.AddInt64(&index.TotalDuplicateSize, fileSize)
	} else {
		index.Duplicates[hash] = fullPath
		index.TotalUniqueFiles++
	}

}

func (index *DuplicateIndex) PrintDuplicateResult() {
	fmt.Println("DUPLICATES")

	fmt.Println("TOTAL UNIQUE FILES: ", index.TotalUniqueFiles)
	fmt.Println("DUPLICATES FILES: ", index.TotalDuplicateFiles)
	fmt.Println("TOTAL DUPLICATE SIZE: ", utils.ToReadableSize(index.TotalDuplicateSize))
}
