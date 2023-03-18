package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Ipv4     string `json:"ip"`
	Region   string `json:"region"`
	Country  string `json:"country_name"`
	City     string `json:"city"`
	Isp      string `json:"isp"`
	As       string `json:"asn"`
	Org      string `json:"asn_org"`
	TimeZone string `json:"timezone_name"`
}

func main() {
	if len(os.Args) <= 1 { /* Length Check  */
		fmt.Printf(" Syntax: iplookup.go <ip>\r\n")
	}

	/* Api */
	var url string = "https://json.geoiplookup.io/" + os.Args[1]

	/* Makes the request to the api */
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Could not unmarsh data")
	}

	/* Prints out the data */
	fmt.Println("IP:", result.Ipv4)
	fmt.Println()
	fmt.Println("Region:", result.Region)
	fmt.Println("Country:", result.Country)
	fmt.Println("City:", result.City)
	fmt.Println("ISP:", result.Isp)
	fmt.Println("AS:", result.As)
	fmt.Println("Org:", result.Org)
	fmt.Println("TimeZone:", result.TimeZone)
}
