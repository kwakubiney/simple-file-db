package sqlitego

import (
	"encoding/gob"
	"log"
	"os"
)

func WriteToIndexMap(db *DB, r Row) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	fileInfo, err := db.File.Stat()
	if err != nil {
		log.Println(err)
	}
	fileOffset := fileInfo.Size()
	db.Bucket[r.ID] = fileOffset
}

func WriteToIndexMapWithoutLock(db *DB, r Row) {
	fileInfo, err := db.File.Stat()
	if err != nil {
		log.Println(err)
	}
	fileOffset := fileInfo.Size()
	db.Bucket[r.ID] = fileOffset
}

func WriteToIndexFile(db *DB) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	encoder := gob.NewEncoder(db.IndexFile)
	encoder.Encode(db.Bucket)
}

func ReadMapFromIndexFile(db *DB) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	decoder := gob.NewDecoder(db.IndexFile)
	decoder.Decode(&db.Bucket)
}

func RemoveIndexFile(db *DB) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	if err := os.Truncate(db.IndexFilePath, 0); err != nil {
		log.Printf("failed to truncate: %v", err)
	}
}
