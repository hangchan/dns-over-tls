package main

import (
	"strings"
	"testing"
)

func TestDnsQuery(t *testing.T) {
	host := "www.example.com"
	dnsServer := "8.8.8.8"
	expectedResult := "93.184.216.34"
	actualResult := strings.Fields(dnsQuery(host, dnsServer))

	dnsQuery(host, dnsServer)
	if actualResult[4] != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
