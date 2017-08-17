package main

import (
	"fmt"
	"time"
	"SSG/simpleDB"
	"net/http"
	"regexp"
	//"os"
	//"strconv"
	//"math"
)

func dbHandler(w http.ResponseWriter, r *http.Request){
	// get the GET and POST variables from the request, load them into a map
	// extract out the $_CONTROL, $_DATA, and $_FILE variables from the request

	
}

var validPath = regexp.MustCompile("^/([a-zA-Z0-9_-]+)/(edit|save|view)/([a-zA-Z0-9]+)$")
  
  func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
  	return func(w http.ResponseWriter, r *http.Request) {
  		m := validPath.FindStringSubmatch(r.URL.Path)
  		if m == nil {
  			http.NotFound(w, r)
  			return
  		}
  		fn(w, r, m[2])
  	}
  }

func main(){
	defer timeTrack(time.Now(), "main")
	defer func() {
		if r:=recover(); r != nil {
			fmt.Println(r)
		}
	}()

	db := simpleDB.NewDB("test")
	
	// if data, err := db.Get("598a6a5e09c27"); err == nil {
	// 	fmt.Println("got data")
	// 	fmt.Println(string(data))
	// } else {
	// 	fmt.Println(err)
	// }

	if data, err := db.GetAll(); err == nil {
		fmt.Println(data)
	} else {
		fmt.Println(err)
	}
	// m := make(map[string]interface{})
	// m["color"] = "blue"
	// m["shape"] = "round"
	// if id, err := db.Post(m); err == nil {
	// 	fmt.Println(id)
	// 	// if data, err := db.Get(id); err != nil {
	// 	// 	fmt.Println(data)
	// 	// } else {
	// 	// 	fmt.Println(err)
	// 	// }
	// } else {
	// 	fmt.Println(err)
	// }

	// if id, err := db.Put("555555", "{'test':'tester'}"); err == nil {
	// 	fmt.Println(id)
	// 	data, err := db.Get("555555")
	// 	fmt.Println(data)
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(err)
	// }
	
	//fmt.Println(Uniqid())
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s\n", name, elapsed)
}