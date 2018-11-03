package main

import (
	"fmt"
	"github.com/miekg/dns"
	"strings"
)

var answer *dns.A

func main() {
	dnsServer()
}

func dnsServer() {
	dns.HandleFunc(".", handleRequest)
	server := &dns.Server{Addr: ":53", Net: "udp"}
	fmt.Println("Starting DNS Server.....")
	fmt.Println("Listening on:", server.Addr)
	server.ListenAndServe()
	defer server.Shutdown()
}

func getIpFromDnsServer(host string, dnsServer string) string {
	c := dns.Client{Net: "tcp-tls"}
	m := dns.Msg{}
	m.SetQuestion(host+".", dns.TypeA)
	r, _, err := c.Exchange(&m, dnsServer+":853")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	for _, ans := range r.Answer {
		Arecord := ans.(*dns.A)
		// fmt.Printf("%s\n", Arecord)
		answer = Arecord
	}
	// return ip address
	s := strings.Fields(answer.String())
	return s[4]
}

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			fmt.Printf("Query for %s\n", q.Name)
			// Can't have trailing "." removing it
			host := strings.TrimSuffix(q.Name, ".")
			// Query Cloudflare DNS
			ip := getIpFromDnsServer(host, "1.1.1.1")
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
	m.Compress = false
	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}
	w.WriteMsg(m)
}
