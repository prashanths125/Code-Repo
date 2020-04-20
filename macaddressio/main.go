package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {

	APIKEY := os.Getenv("MACADDRESSIO_API_KEY")
	if APIKEY == "" {
		fmt.Println("APIKEY is missing - Please set MACADDRESSIO_API_KEY environment variable.")
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Println("Must provide exactly one argument - the MAC address")
		os.Exit(1)
	}
	MACADDRESS := os.Args[1]
	if !isValidMACAddress(MACADDRESS) {
		fmt.Println("MAC address provided was invalid. A sample valid MAC address looks like 44:38:39:ff:ef:57")
		os.Exit(1)
	}

	URL := fmt.Sprintf("https://api.macaddress.io/v1?apiKey=%s&output=json&search=%s", APIKEY, MACADDRESS)

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("There was an error retrieving data from an upstream system")
		os.Exit(2)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	name, err := extractCompanyName(body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(name)
}

func extractCompanyName(body []byte) (companyName string, err error) {
	var macResponse map[string]interface{}
	err = json.Unmarshal([]byte(body), &macResponse)
	if err != nil || macResponse["error"] != nil {
		return "", errors.New("API response was not correct. Please check APIKey or MACaddress")
	}

	return macResponse["vendorDetails"].(map[string]interface{})["companyName"].(string), nil
}

func isValidMACAddress(mac string) bool {
	validMAC := regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)
	return validMAC.Match([]byte(mac))
}
