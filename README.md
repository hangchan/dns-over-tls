# dns-over-tls

This starts a dns server listening on udp port 53.  When a dns lookup is made, the server will initiate a tls connection with Cloudflare DNS, retrieve the ip address and return it to the client.

## Getting Started

Clone repo from github.com or dns-over-tls.zip

### Prerequisites

```
GO
package github.com/miekg/dns
```

### Installing

Building
```
#go build
```

Running
```
#sudo ./dns-over-tls 
Password:
Starting DNS Server.....
Listening on: :53
```

Lookup DNS Record
```
#dig @127.0.0.1 www.example.com

; <<>> DiG 9.8.3-P1 <<>> @127.0.0.1 www.example.com
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 13303
;; flags: qr rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;www.example.com.		IN	A

;; ANSWER SECTION:
www.example.com.	3600	IN	A	93.184.216.34

;; Query time: 150 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Mon Nov  5 13:13:59 2018
;; MSG SIZE  rcvd: 64
```

## Running the tests

Fetches an ip address from Cloudflare DNS Server over tls

```
#go test
PASS
ok  	github.com/hangchan/dns-over-tls	0.222s
```

## Deployment

Build docker image
```
docker build
docker tag
docker push
```

Run the server
```
sudo docker run -i -p 53:53/udp -t hangchan/dns-over-tls:0.0.6
```
