package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	IP   string `json:"ip"`
	Cidr string `json:"cidr"`
	ASN  uint32 `json:"asn"`
	Desc string `json:"desc"`
}

type Result struct {
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	Addresses []Address `json:"addresses"`
	Tag       string    `json:"cert"`
	Sources   []string  `json:"srouces"`
}

func main() {
	f, err := os.Open("utils/.amass4owasp.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		var result Result
		json.Unmarshal([]byte(m), &result)

		fmt.Println(result)
    break
	}
}
