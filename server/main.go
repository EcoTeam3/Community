package main

import (
	"log"
	"community/storage"
)

func main(){
	db, err := storage.Connect()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
}