package main

import (
	"testing"
)

func TestDnsQuery(t *testing.T) {
	host := "www.example.com"
	dnsServer := "8.8.8.8"
	expectedResult := "93.184.216.34"
	actualResult := dnsQuery(host, dnsServer)

	dnsQuery(host, dnsServer)
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
