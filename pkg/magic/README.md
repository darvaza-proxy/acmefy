# magic

`magic` is a TLS store that falls back to self-signed if needed,
uses ACME to acquire new certs, and renews those about to expire.

## Sequence

```mermaid
sequenceDiagram
    actor client
    participant SRV as http.Server
    participant AC as magic
    participant R as ACME Server

    client ->>+ SRV: HTTP GET

    SRV -->>+ AC: GetCertificate()

    opt not cached?
    AC -->>+ R: ACME Request
        alt success?
        R -->>- AC: *tls.Certificate{}
        else
        AC -->> AC: IssueCertificate()
        end
    end

    AC -->>- SRV: *tls.Certificate{}

    SRV -->>- client: HTTP Response
```

## Related Projects

* [Darvaza Autocert](https://darvaza.org/darvaza/shared/storage/autocert)
