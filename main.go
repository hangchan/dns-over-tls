package main

import (
	"fmt"
	"github.com/miekg/dns"
)

func main() {
	dnsServer()
	dnsQuery()
}

func dnsServer() {
	dns.HandleFunc(".", handleRequest)
	server := &dns.Server{Addr: ":53", Net: "udp"}
	server.ListenAndServe()
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

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	w.WriteMsg(m)
}
