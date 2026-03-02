# OpenAPI MCP Server

HTTP server that converts OpenAPI specifications into MCP tools and proxies calls to upstream APIs.

## Quick Start

1. Clone the repository:

```sh
git clone https://github.com/merzzzl/openapi-mcp-server.git
cd openapi-mcp-server
```

2. Copy and edit the configuration:

```sh
cp config.example.yaml config.yaml
```

3. Run the server:

```sh
go run ./cmd/app -config config.yaml
```

4. MCP endpoints are available at:

```
http://localhost:9090/{server_name}/mcp   — Streamable HTTP
http://localhost:9090/{server_name}/sse   — SSE
```

5. Health check:

```
http://localhost:8080/healthz
```

## Configuration

```yaml
# Main HTTP server port (defaults to ":8080").
port: ":9090"

# Enable TOON format for tool responses.
enable_toon: false

# OpenTelemetry: traces, logs, metrics via OTLP gRPC.
# If endpoint is empty — traces and metrics are disabled, logs go to stdout.
telemetry:
  endpoint: "localhost:4317"
  insecure: true

# List of OpenAPI servers. Each one is registered as a separate MCP server.
servers:
  - # Server name — used in URL path (/{name}/mcp, /{name}/sse).
    name: petstore

    # OpenAPI specification URL (v2 or v3).
    schema_url: "https://petstore3.swagger.io/api/v3/openapi.json"

    # Base URL for proxying tool calls.
    base_url: "https://petstore3.swagger.io/api/v3"

    # Operation filter rules. allow — whitelist, block — blacklist.
    # Each rule: methods (HTTP methods) + regex (path pattern).
    allow:
      - methods: ["GET", "POST", "PUT", "DELETE"]
        regex: ".*"
    block: []
```

## Prerequisites

- Go >= 1.24
- OpenTelemetry Collector (optional, for telemetry collection)
