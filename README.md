# simpleproxy

A simple proxy tool that can be run locally to send requests to a remote URL. It's intended primarily as a collection of hacks for fixing snags in development processes.

Clone with:

```
go get github.com/kujenga/simpleproxy
```

Run `simpleproxy -h` for usage information.

### Use cases:

- Communicating with a HTTP endpoint that doesn't support HTTPS on a domain that has [HSTS](https://en.wikipedia.org/wiki/HTTP_Strict_Transport_Security). Obviously the better option here is to add HTTPS to the endpoint.
- Adding basic authentication to all proxied requests. The better option here is to add real authentication support to your app, perhaps with something like [OpenID Connect](http://openid.net).
