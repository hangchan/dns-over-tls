package main

import (
	"fmt"
	"github.com/miekg/dns"
)

var records = map[string]string{
	"test.service.": "192.168.0.2",
}
var answer *dns.A

func main() {
	dnsQuery("www.google.com", "8.8.8.8")
	fmt.Println(answer)
	dnsServer()
}

func dnsServer() {
	dns.HandleFunc(".", handleRequest)
	server := &dns.Server{Addr: ":53", Net: "udp"}
	server.ListenAndServe()
}

func dnsQuery(host string, dnsServer string) string {
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
		// fmt.Printf("%s\n", Arecord)
		answer = Arecord
	}
	return answer.String()
}

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			fmt.Printf("Query for %s\n", q.Name)
			ip := records[q.Name]
			if ip != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		}
	}
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}
	w.WriteMsg(m)
}
