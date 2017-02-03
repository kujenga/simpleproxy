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
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
