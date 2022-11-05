package entryhandler

import (
	"os"
	"path"
)

type DuplicateIndex struct {
	TotalUniqueFiles    int
	TotalDuplicateFiles int
	TotalDuplicateSize  int64
	Duplicates          map[string]string
}

type EntryHandler interface {
	Handle(*DuplicateIndex) error
}

type FileEntry struct {
	fullPath string
	size     int64
}

type DirEntry struct {
	fullPath string
}

type NilEntry struct {
}

func NewEntryHandler(fileEntry os.FileInfo, directory string) EntryHandler {

	path := path.Join(directory, fileEntry.Name())

	if fileEntry.Mode().IsDir() {
		return &DirEntry{path}
	}

	if fileEntry.Mode().IsRegular() {
		return &FileEntry{path, fileEntry.Size()}
	}

	return &NilEntry{}
}

func (entry *DirEntry) Handle(index *DuplicateIndex) error {
	return index.TraverseDirPath(entry.fullPath)
}

func (entry *FileEntry) Handle(index *DuplicateIndex) error {
	//return index.TraverseDirPath(entry.fullPath)
	hash, err := createFileHash(entry.fullPath)
	if err != nil {
		return err
	}

	index.AddHashEntry(entry.fullPath, entry.size, hash)

	return nil
}

func (entry *NilEntry) Handle(index *DuplicateIndex) error {
	return nil
}

func NewDuplicateIndex() DuplicateIndex {

	duplicateIndex := DuplicateIndex{
		Duplicates:          make(map[string]string),
		TotalUniqueFiles:    0,
		TotalDuplicateFiles: 0,
		TotalDuplicateSize:  0,
	}

	return duplicateIndex
}
