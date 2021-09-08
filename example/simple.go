package main

import (
    gorocksdb "github.com/DGHeroin/rocksdb"
    "log"
    "time"
)

func main()  {
    opt := gorocksdb.NewDefaultOptions()
    opt.SetCreateIfMissing(true)
    db, err :=  gorocksdb.OpenDb(opt, "db/test")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    db.Put(gorocksdb.NewDefaultWriteOptions(), []byte("key"), []byte(time.Now().String()))
    rs, _ := db.Get(gorocksdb.NewDefaultReadOptions(), []byte("key"))
    log.Println(string(rs.Data()))
}

