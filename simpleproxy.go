package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/skratchdot/open-golang/open"
)

var (
	port = flag.String("port", "6600", "address to listed on")
	user = flag.String("username", "", "username for basic authentication")
	pass = flag.String("password", "", "password for basic authentication")
	o    = flag.Bool("o", false, "open the proxy in the browser")
)

func main() {
	flag.Parse()

	target := flag.Arg(0)
	if target == "" {
		log.Fatal("Specify a proxy target as a command line argument")
	}

	tgt, err := url.Parse(target)
	if err != nil {
		log.Fatalln("error parsing target:", err)
	}

	log.Printf("Proxying to %s on port %s", tgt.String(), *port)

	proxy := httputil.NewSingleHostReverseProxy(tgt)
	if *user != "" || *pass != "" {
		proxy.Transport = &basicAuthTransport{
			Username: *user,
			Password: *pass,
		}
	}

	if *o {
		go func() {
			open.Run("http://localhost:" + *port)
		}()
	}

	log.Fatal(http.ListenAndServe(":"+*port, proxy))
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
