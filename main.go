package main

import (
	"flag"
	"log"

	"github.com/jamesroutley/guestbook/datastore"
	"github.com/jamesroutley/guestbook/handler"
)

func main() {
	port := flag.Int("port", 80, "port number")
	flag.Parse()

	ds, err := datastore.NewSQLiteDatastore("./db/guestbook.db")
	if err != nil {
		log.Fatal(err)
	}
	handler.Serve(*port, ds)
}
