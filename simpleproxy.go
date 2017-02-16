package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	target = flag.String("target", "", "proxy target")
	addr   = flag.String("addr", ":6600", "address to listed on")
	user   = flag.String("username", "", "username for basic authentication")
	pass   = flag.String("password", "", "password for basic authentication")
)

func main() {
	flag.Parse()

	if *target == "" {
		log.Fatal("Specify a proxy with the -target flag")
	}

	tgt, err := url.Parse(*target)
	if err != nil {
		log.Fatalln("error parsing target:", err)
	}

	log.Printf("Proxying to %s on %s", tgt.String(), *addr)

	proxy := httputil.NewSingleHostReverseProxy(tgt)
	if *user != "" || *pass != "" {
		proxy.Transport = &basicAuthTransport{
			Username: *user,
			Password: *pass,
		}
	}

	log.Fatal(http.ListenAndServe(*addr, proxy))
}

// based on net/http/httputil internals

type basicAuthTransport struct {
	Username string
	Password string
}

func (t basicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.Username, t.Password)
	return http.DefaultTransport.RoundTrip(req)
}

func (t *basicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}
