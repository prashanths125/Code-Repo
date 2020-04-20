package main

import "testing"

func TestExtractCompanyNameFromGoodJSON(t *testing.T) {
	inputJSON := []byte(`{"vendorDetails":{"companyName":"ABC Systems"}}`)
	wanted := "ABC Systems"

	got, err := extractCompanyName(inputJSON)

	if got != wanted {
		t.Errorf("Got %v, wanted %v", got, wanted)
	}
	if err != nil {
		t.Errorf("Expected nil error")
	}
}

func TestExtractCompanyNameFromErrorJSON(t *testing.T) {
	malformedJSON := []byte(`{"error":{"data":"Access denied"}}`)
	wanted := ""
	got, err := extractCompanyName(malformedJSON)

	if err == nil {
		t.Errorf("Expected error")
	}
	if got != wanted {
		t.Errorf("Expected empty result but got %s", got)
	}
}

func TestValidMACAddress(t *testing.T) {
	goodMACAddress := "44:38:39:ff:ef:57"
	wanted := true

	got := isValidMACAddress(goodMACAddress)

	if got != wanted {
		t.Errorf("Wanted %v, but got %v", wanted, got)
	}

}

func TestInvalidMACAddress(t *testing.T) {
	badMACAddress := "xa:38:39:ff:ef:57"
	wanted := false

	got := isValidMACAddress(badMACAddress)

	if got != wanted {
		t.Errorf("Wanted %v, but got %v", wanted, got)
	}

}
