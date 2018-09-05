package db

import (
	"sync"
	"os"
	"fmt"
	"encoding/json"
)

type DB struct {
	mtx       sync.Mutex
	cacheFile string
	filePath  string
	data      map[string]interface{}
}

func (self *DB) Load(filePath string) error {
	self.data = map[string]interface{}{}
	self.filePath = filePath

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil
	}
	r, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("%s error opening file: %v", filePath, err)
	}
	defer r.Close()

	dec := json.NewDecoder(r)
	err = dec.Decode(&self.data)
	if err != nil {
		return fmt.Errorf("error reading %s: %v", filePath, err)
	}

	return nil
}

func (self *DB) Save() error {
	w, err := os.Create(self.filePath)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(w)
	defer w.Close()
	if err := enc.Encode(&self.data); err != nil {
		return err
	}

	return nil
}

func (self *DB) Set(key string, value interface{}) error {
	self.data[key] = value
	return nil
}

func (self *DB) Get(key string) interface{} {
	value, ok := self.data[key]
	if ok {
		return value
	} else {
		return nil
	}
}

func (self *DB) GetAll() map[string]interface{} {
	return self.data
}
