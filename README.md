# nopapp

A minimal Go application used as the victim consumer in the Go supply chain security research.

It sets a secret in its environment and calls into `noplib`:

```go
os.Setenv("MY_SECRET_CREDENTIAL", "secret_cred")
fmt.Println(noplib.CallToNopLib())
```

This simulates a CI runner that has credentials in its environment (API tokens, cloud keys, etc.) while building Go code — the same environment targeted in the Trivy-action attack. If a malicious `init()` is present in `noplib`, it runs during `go build` and can read `MY_SECRET_CREDENTIAL` before `main()` is ever reached.

The tests vary the `GOPROXY`, `GONOSUMDB`, and `go.sum` state of this module to observe how Go's three-layer verification chain responds to a compromised `noplib`.
