package main

import (
	"fmt"
	"time"
	"SSG/simpleDB"
	//"strconv"
	//"math"
)

func main(){
	defer timeTrack(time.Now(), "main")
	defer func() {
		if r:=recover(); r != nil {
			fmt.Println(r)
		}
	}()
	db := simpleDB.NewDB("test")
	
	if data, err := db.Get("1234"); err == nil {
		fmt.Println("got data")
		fmt.Println(data)
	} else {
		fmt.Println(err)
	}

	if data, err := db.GetAll(); err == nil {
		fmt.Println(data)
	} else {
		fmt.Println(err)
	}

	if id, err := db.Post("{'hello':'world'}"); err == nil {
		fmt.Println(id)
		data, err := db.Get(id)
		fmt.Println(data)
		fmt.Println(err)
	} else {
		fmt.Println(err)
	}

	if id, err := db.Put("555555", "{'test':'tester'}"); err == nil {
		fmt.Println(id)
		data, err := db.Get("555555")
		fmt.Println(data)
		fmt.Println(err)
	} else {
		fmt.Println(err)
	}
	
	//fmt.Println(Uniqid())
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s\n", name, elapsed)
}