package main

import (
	"testing"
)

func TestGetIpFromDnsServer(t *testing.T) {
	host := "www.example.com"
	dnsServer := "1.1.1.1"
	expectedResult := "93.184.216.34"
	actualResult := getIpFromDnsServer(host, dnsServer)
	getIpFromDnsServer(host, dnsServer)
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
