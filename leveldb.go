package main

import (
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

func Byte2int(val []byte) int {
	var result int
	result, _ = strconv.Atoi(string(val))
	return result
}

//for test
func write_leveldb(x, y string) {
	//	db, err := leveldb.OpenFile("/home/brian/.tendermint/data", nil)
	db, err := leveldb.OpenFile("./data/blockstore.db", nil)
	if err != nil {
		panic(err)
	}
	db.Put([]byte(x), []byte(y), nil)
	db.Close()
}

//for test
func read_leveldb(x string) {
	db, err1 := leveldb.OpenFile("./data/blockstore.db", nil)
	if err1 != nil {
		panic(err1)
	}
	ids, err2 := db.Get([]byte(x), nil)
	if err2 != nil {
		panic(err2)
	}
	id := string(ids)
	fmt.Println("read leveldb:", id)
	db.Close()
}

func read_byte(x []byte) {
	db, err1 := leveldb.OpenFile("./data/blockstore.db", nil)
	if err1 != nil {
		panic(err1)
	}
	ids, err2 := db.Get(x, nil)
	if err2 != nil {
		panic(err2)
	}
	id := ids
	db.Close()
	fmt.Println(id)
}

func iterate_leveldb() {
	//db, err := leveldb.OpenFile("./production/ledgersData/chains/index", nil)
	db, err := leveldb.OpenFile("./production/ledgersData/stateLeveldb", nil)
	if err != nil {
		panic(err)
	}
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		//print in string instead of binaries
		//str := fmt.Sprintf("%s", key)
		//str2 := fmt.Sprintf("%s", value)
		fmt.Println("-------------------")
		fmt.Println(key)

		fmt.Println(value)

	}
	iter.Release()
	err = iter.Error()

	db.Close()
}

func main() {
	//write_leveldb("5","12")
	//read_leveldb("5")
	iterate_leveldb()
}
