package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := os.Open("utils/enum_result.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	ipv4_regex := `(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for count := 0; count < 10 && scanner.Scan(); count++ {
		m := scanner.Text()

		domain := strings.Split(m, " ")[0]
		ipv4 := regexp.MustCompile(ipv4_regex).FindAllString(m, -1)

		if len(ipv4) > 0 {
			fmt.Printf("domain: %v\nipv4: %v\n", domain, ipv4)
		}
	}

	log.Printf("finished...")

}
