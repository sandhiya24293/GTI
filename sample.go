package main

import (
	Db "Dgraph/Common/DB/Mysql"
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
	Domain       string
	Reputation   int   `json:"reputation"`
	Categories   []int `json:"categories"`
	Categoryname []string
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

//type Dmcatype []map[string]string

func getDMcategory(Cat []int) (resstring []string) {

	var data []Datastruct

	file, err := ioutil.ReadFile("url.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	var res string
	for i := 0; i < len(Cat); i++ {
		for _, v := range data {

			if v.Key == Cat[i] {

				res = v.Value
				fmt.Println(res)

			}

		}
		resstring = append(resstring, res)
		fmt.Println(resstring)

	}

	return
}

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
	getDmcat := getDMcategory(res.Categories)
	fmt.Println(getDmcat)

	Response.Domain = domain
	Response.Reputation = res.Reputation
	Response.Categories = res.Categories
	Response.Categoryname = getDmcat

	Db.InsertDomain(Response.Domain, Response.Reputation, Response.Categories, Response.Categoryname)

	//	jsonresponse = append(jsonresponse, Response)

	//	rankingsJson, _ := json.Marshal(jsonresponse)
	//	err = ioutil.WriteFile("sample1.json", rankingsJson, 0644)

}
