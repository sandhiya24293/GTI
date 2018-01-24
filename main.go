package main

import (
	Db "GTI/Common/DB/Mysql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	//"net/http"
)

type Datastruct struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

func main() {
	var data []Datastruct

	file, err := ioutil.ReadFile("url.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data {
		fmt.Println(v.Key, v.Value)
		Db.InsertGti(v.Value, v.Key)
	}

}
