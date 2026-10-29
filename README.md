# Trade Agent Gate

A small Go API example with a synthetic request playground.

```sh
go run .
```

Open http://127.0.0.1:8090. The API serves `/health` and `/api/demo-trades`. The UI includes illustrative API and MCP messages. It does not perform live screening or authorize shipments.

```sh
go test -race ./...
```
