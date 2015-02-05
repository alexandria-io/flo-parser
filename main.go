package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("/home/ubuntu/.litecoin-insight/txs", nil)

	if err != nil {
		fmt.Println("\n===\n")
		fmt.Println(err)
		fmt.Println("\n===\n")
		log.Fatal("error reading database")
	}

	iter := db.NewIterator(nil, nil)

	alltxc := []string{}

	txprotocol := []byte("txc-")

	for iter.Next() {
		key := iter.Key()

		// only care about keys that correspond with tx-comments
		if bytes.HasPrefix(key, txprotocol) {
			value := iter.Value()
			alltxc = append(alltxc, string(value))
			// fmt.Printf("txid: %v\n%v\n\n", key, value)
		}
	}

}
