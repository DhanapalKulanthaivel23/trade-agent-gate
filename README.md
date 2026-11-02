# Trade Agent Gate

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
