# Trade Agent Gate

[**Trade Agent Gate-demo**](https://DhanapalKulanthaivel23.github.io/trade-agent-gate/)


A small Go API example with a synthetic request playground.

```sh
go run .
```

Open http://127.0.0.1:8090. The API serves `/health` and `/api/demo-trades`. The UI includes illustrative API and MCP messages. It does not perform live screening or authorize shipments.

```sh
go test -race ./...
```

## Domain API workbench

Open `/workbench.html` for preloaded REST requests, Postman export, applicable file/MCP tabs, a functional walkthrough and explicit technology/status explanations. Sample mode performs no network/model calls. Live mode needs a separately configured private gateway; the public shell does not implement private endpoints. No credentials are exported. The existing IAM request/ticket studio is preserved.

## Private workspace boundary (supersedes public export instructions)

The workbench now loads authorized collections and MT/MX context locally from its private core. Public download/export controls and embedded request fixtures have been retired. Keep workspace.private.json out of public web roots. Reference replay is distinct from live execution; the interface does not claim a live banking connection or conformance. Earlier export instructions describe the previous UI only.

## Current collection and MCP interface

The latest generic interface source is in [docs/project-pages](docs/project-pages/README.md). It provides separate collection and MCP pages plus applicable file, graph and cited-question views. The private project host supplies request metadata and executes the existing services; no private collections or implementation are bundled here.

## One interface for local and hosted use

Run `go run .` and open http://127.0.0.1:8080/. API Collections, MCP Tools, Files, Knowledge Graph and Questions & Answers use the same client assets as the GitHub Pages link above. GitHub Pages navigation uses URL fragments so repository subpaths work without server rewrites.

The public client calls the owner's running core API using `client-config.json`. The configured endpoint must be reachable HTTPS for remote visitors. `CORE_API_URL` can override it for local integration tests. No GitHub credential or access to private source is required by API users. Collections and business implementation are supplied by the private core, not copied into this repository.

A repository push alone does not start a core service. The endpoint is currently awaiting the owner's hosting configuration; an unavailable core is shown as a connection failure rather than a successful response.

### Updating the core connection

The owner sets the GitHub Actions repository variable `CORE_API_URL` to the HTTPS core presentation API and runs **Publish project interface**. The Pages deployment updates its public connection configuration. Local clones automatically read that same hosted configuration, so they do not need new code or private-repository access. Until the endpoint is supplied, the interface reports the unavailable connection.

For local integration testing on the owner's machine: `CORE_API_URL=http://127.0.0.1:1880X go run .`, using the corresponding Trade/Compliance/IAM/Payments core port 18801/18802/18803/18804. The endpoint value is public routing information, never a provider key or GitHub token.
