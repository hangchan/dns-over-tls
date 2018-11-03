package main

import (
	"fmt"
	"github.com/miekg/dns"
)

func main() {
	dnsQuery()
}

func dnsQuery() {

	host := "www.google.com"
	dnsServer := "1.1.1.1"

	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(host+".", dns.TypeA)
	r, _, err := c.Exchange(&m, dnsServer+":53")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	for _, ans := range r.Answer {
		Arecord := ans.(*dns.A)
		fmt.Printf("%s\n", Arecord)
	}
}
