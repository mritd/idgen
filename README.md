# idgen

> Chinese identity information generator written in Go. Supports generating name, ID number, mobile phone number, bank card number, email address and address.

Part of the code is translated from [java-testdata-generator](https://github.com/binarywang/java-testdata-generator). Thanks to the original author [binarywang](https://github.com/binarywang).

## Installation

Download pre-compiled binaries from the [Release](https://github.com/mritd/idgen/releases) page.

Docker users can pull the image directly:

```bash
docker pull mritd/idgen
```

### Build from source

Requires Go 1.25+:

```bash
# Install task runner
go install github.com/go-task/task/v3/cmd/task@latest

# Build
task build
```

## CLI Mode

Run the binary directly to generate identity information. Results are automatically copied to clipboard by default.

```
$ idgen --help

Identity information generator for Chinese name, ID number,
bank card number, mobile phone number, address and Email.

Generate ID number by default without sub-command.

Usage:
  idgen [flags]
  idgen [command]

Available Commands:
  addr        Generate address information
  all         Generate all information
  bank        Generate bank card number
  email       Generate email address
  idno        Generate ID number
  mobile      Generate mobile phone number
  name        Generate name
  server      Run as http server
  version     Print version

Flags:
  -C, --copy            Copy to clipboard (default: true for single, false for batch)
  -c, --count int       Number of records to generate (default 1)
  -f, --format string   Output format: table|json|csv (default "table")
  -h, --help            help for idgen
```

### Examples

```bash
# Generate single ID number (copied to clipboard)
idgen

# Generate 10 records in table format
idgen -c 10

# Generate all information
idgen all

# Batch generate in JSON format
idgen all -c 5 -f json

# Export to CSV
idgen all -c 100 -f csv > data.csv
```

## Server Mode

Start an HTTP server with web UI and API endpoints.

```
$ idgen server --help

Run a simple http server to provide page access and API.

Access endpoints:
  http://BINDADDR:PORT/                  HTML page with theme
  http://BINDADDR:PORT/api/v1/generate   Generate single record (JSON)
  http://BINDADDR:PORT/api/v1/batch      Batch generate (JSON)
  http://BINDADDR:PORT/api/v1/export     Export as CSV

Flags:
  -l, --listen string   HTTP listen address (default "0.0.0.0")
  -p, --port int        HTTP listen port (default 8080)
  -t, --theme string    Default theme: cyber|terminal (default "cyber")
```

### Run with Docker

```bash
docker run -d -p 8080:8080 mritd/idgen server
```

### Themes

The web UI supports two themes:

- **Cyberpunk** (`cyber`): Neon colors with particle animation and mouse interaction
- **Terminal** (`terminal`): Matrix-style falling characters with green monochrome

Switch themes via the button in the top-right corner, or set default theme with `--theme` flag.

### API Endpoints

```bash
# Generate single record
curl http://localhost:8080/api/v1/generate

# Batch generate (default 10)
curl http://localhost:8080/api/v1/batch

# Batch generate with count
curl http://localhost:8080/api/v1/batch?count=50

# Export as CSV
curl http://localhost:8080/api/v1/export?count=100
```

## License

MIT
