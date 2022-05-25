// Package id provides unique uint64 identifiers for datasets designated by a string.
// The uint64 identifiers start at 1 for a new dataset and are monotonically increasing,
// persisting through restarts of the server.  The package also allows retrieving
// data associated with the id like a user and branch string.
package id

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const idFileName = "idfile"

var (
	filePath string // where to store the id file
	curID    uint64
	idMu     sync.Mutex
)

// LoadID initializes the unique id and must be called before
// any other id functions.
func LoadID(path string) error {
	if err := ensureDir(path); err != nil {
		return err
	}
	filePath = filepath.Join(path, idFileName)
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("No id file @ %q found.  Starting with new id set.\n", filePath)
			return nil
		}
		return fmt.Errorf("couldn't open id file %q: %v", filePath, err)
	}
	defer file.Close()
	if _, err = fmt.Fscanf(file, "%d\n", &curID); err != nil {
		return fmt.Errorf("couldn't read from id file %q: %v", filePath, err)
	}
	log.Printf("Loaded id: %d\n", curID)
	return nil
}

// Generate returns a unique id.
func GenerateID() (newid uint64, err error) {
	idMu.Lock()
	defer idMu.Unlock()
	if err = persistID(curID + 1); err != nil {
		return
	}
	curID++
	return curID, nil
}

// GenerateIDs returns a range of unique ids: [first, last].
func GenerateIDs(count uint64) (idrange [2]uint64, err error) {
	idMu.Lock()
	defer idMu.Unlock()
	if err = persistID(curID + count); err != nil {
		return
	}
	start := curID + 1
	curID += count
	return [2]uint64{start, curID}, nil
}

func persistID(id uint64) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("couldn't open id file %q: %v", filePath, err)
	}
	defer file.Close()
	if _, err = file.WriteString(fmt.Sprintf("%d\n", id)); err != nil {
		return fmt.Errorf("couldn't write to id file %q: %v", filePath, err)
	}
	if err = file.Sync(); err != nil {
		return fmt.Errorf("couldn't sync id file %q: %v", filePath, err)
	}
	return nil
}

// ensures a directory exists or will create it, or returns error
// if path is not directory
func ensureDir(path string) error {
	if fileinfo, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Creating directory: %s\n", path)
		err := os.MkdirAll(path, 0744)
		if err != nil {
			return fmt.Errorf("can't make directory: %v", err)
		}
	} else if !fileinfo.IsDir() {
		return fmt.Errorf("path (%s) is not a directory", path)
	}
	return nil
}
