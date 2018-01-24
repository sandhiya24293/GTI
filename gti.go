package main

import (
	Db "GTI/Common/DB/Mysql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type resStr struct {
	Reputation int   `json:"reputation"`
	Categories []int `json:"categories"`
}

type ReputationCat struct {
	Domain     string
	Reputation int   `json:"reputation"`
	Categories []int `json:"categories"`
}

type datatype []string

func main() {
	var data datatype
	file, err := ioutil.ReadFile("sample.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data {
		getdomain(v)
	}

}

type Datastruct struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

var jsonresponse []ReputationCat
var Response ReputationCat

func getdomain(domain string) {

	urlStr := "http://localhost:8087/local/gti/" + domain + "/rate"

	resp, err := http.Get(urlStr)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	res := new(resStr)
	if err = json.Unmarshal(content, &res); err != nil {
		fmt.Println("Error :", err)
		return
	}

	Response.Domain = domain
	Response.Reputation = res.Reputation
	Response.Categories = res.Categories
	source = "https://adaway.org/hosts.txt"

	Db.InsertDomain(Response.Domain, Response.Reputation, Response.Categories, source)

}
