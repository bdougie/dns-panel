package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var open bool

func Open() error {
	var err error
	_, filename, _, _ := runtime.Caller(0) // get full path of this file
	dbfile := path.Join(path.Dir(filename), "data.db")
	config := &bolt.Options{Timeout: 1 * time.Second}
	db, err = bolt.Open(dbfile, 0600, config)
	if err != nil {
		log.Fatal(err)
	}
	open = true
	return nil
}

func Close() {
	open = false
	db.Close()
}

func (r *Record) save() error {
	if !open {
		return fmt.Errorf("db must be opened before saving!")
	}
	err := db.Update(func(tx *bolt.Tx) error {
		records, err := tx.CreateBucketIfNotExists([]byte("records"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		enc, err := r.encode()
		if err != nil {
			return fmt.Errorf("could not encode Record %s: %s", r.ID, err)
		}
		err = records.Put([]byte(r.ID), enc)
		return err
	})
	return err
}

func (r *Record) encode() ([]byte, error) {
	enc, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func decode(data []byte) (*Record, error) {
	var r *Record
	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func seedDB() {
	Data := Records{
		Record{ID: "1", Domain: "web.com", Name: "A", Address: "192.62.1.1"},
		Record{ID: "2", Domain: "da-web.com", Name: "MX", Address: "172.42.1.1"},
		Record{ID: "3", Domain: "internetweb.com", Name: "CNAME", Address: "196.52.1.1"},
		Record{ID: "4", Domain: "google.com", Name: "CNAME", Address: "192.63.1.1"},
	}

	// Persist Records in the Database
	for _, r := range Data {
		r.save()
	}
}

func GetRecord(id string) (*Record, error) {
	if !open {
		return nil, fmt.Errorf("db must be opened before fetching!")
	}
	var r *Record
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("records"))
		k := []byte(id)
		r, _ = decode(b.Get(k))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Could not get Record ID %s", id)
		return nil, err
	}
	return r, nil
}

func AllRecords() ([]*Record, error) {
	if !open {
		return nil, fmt.Errorf("db must be opened before fetching!")
	}

	records := []*Record{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("records"))
		b.ForEach(func(k, v []byte) error {
			r, err := decode(b.Get(k))

			records = append(records, r)
			if err != nil {
				return err
			}

			return nil
		})

		return nil
	})

	if err != nil {
		fmt.Println("Could not get Records")
		return nil, err
	}
	return records, nil
}

func init() {
	seedDB()
}
