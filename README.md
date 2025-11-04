# irkit-webserver

[![Go Version](https://img.shields.io/badge/Go-1.22-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/docker-kangaechu%2Firkit-blue)](https://hub.docker.com/r/kangaechu/irkit)

A lightweight Go web server for controlling IR (infrared) devices through [IRKit](http://getirkit.com/) hardware. Control your air conditioner, ceiling lights, and other IR-enabled devices via simple HTTP endpoints.

## Features

- 🌡️ **Air Conditioner Control** - Turn your AC on/off remotely
- 💡 **Ceiling Light Control** - Control ceiling lights with multiple modes
- 🚀 **Zero Dependencies** - Built with Go standard library only
- 🐳 **Docker Support** - Ready-to-use Docker images
- 🔧 **Easy Integration** - Simple HTTP API for automation and smart home setups
- 📦 **Multi-platform** - Binaries available for Linux, macOS (darwin), ARM devices

## Architecture

```
HTTP Request → irkit-webserver (port 9090) → Shell Scripts → IRKit Device (IR signals)
```

The server receives HTTP requests and executes corresponding shell scripts that send infrared signal data to your IRKit device via HTTP POST requests.

## Prerequisites

- [IRKit](http://getirkit.com/) hardware device configured on your network
- Go 1.22+ (for building from source)
- Docker (optional, for containerized deployment)
- `curl` installed (for shell scripts)

## Installation

### Option 1: Download Pre-built Binary

Download the latest release for your platform from [GitHub Releases](https://github.com/kangaechu/irkit-webserver/releases).

```bash
# Example for Linux amd64
wget https://github.com/kangaechu/irkit-webserver/releases/latest/download/irkit-webserver_linux_amd64.tar.gz
tar -xzf irkit-webserver_linux_amd64.tar.gz
chmod +x irkit-webserver
```

### Option 2: Build from Source

```bash
git clone https://github.com/kangaechu/irkit-webserver.git
cd irkit-webserver
go build -o irkit-webserver
```

### Option 3: Docker

```bash
docker pull kangaechu/irkit:latest
```

## Configuration

### 1. Update IRKit Device IP Address

Edit the shell scripts in the `command/command/` directory to match your IRKit device's IP address. By default, they're configured for `192.168.1.117`.

```bash
# Example: command/command/aircon_on.sh
curl -i -XPOST 'http://192.168.1.117/messages' ...
```

### 2. Configure Shell Script Path (Optional)

If running locally (not in Docker), update the script paths in [main.go](main.go) line 18-22 to point to your script locations:

```go
const commandPath = "/home/pi/bin/irkit/"  // Update this path
```

### 3. Change Port (Optional)

The server runs on port 9090 by default. To change it, modify [main.go](main.go) line 12:

```go
const port = "9090"  // Change to your preferred port
```

## Usage

### Running the Server

**Local:**
```bash
./irkit-webserver
```

**Docker:**
```bash
docker run -d -p 9090:9090 --name irkit kangaechu/irkit:latest
```

The server will start on `http://localhost:9090`

### Controlling Devices

Use any HTTP client to control your devices:

```bash
# Turn on air conditioner
curl http://localhost:9090/aircon/on

# Turn off air conditioner
curl http://localhost:9090/aircon/off

# Turn on ceiling light
curl http://localhost:9090/ceiling/on

# Turn off ceiling light
curl http://localhost:9090/ceiling/off

# Full brightness ceiling light
curl http://localhost:9090/ceiling/full
```

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Returns usage information and available endpoints |
| `/aircon/on` | GET | Power on air conditioner |
| `/aircon/off` | GET | Power off air conditioner |
| `/ceiling/on` | GET | Turn on ceiling light |
| `/ceiling/off` | GET | Turn off ceiling light |
| `/ceiling/full` | GET | Set ceiling light to full brightness |

All endpoints return:
- `200 OK` on success
- `500 Internal Server Error` if the command execution fails

## Development

### Running Tests

```bash
go test ./...
```

### Linting

This project uses [golangci-lint](https://golangci-lint.run/) for code quality:

```bash
golangci-lint run
```

### Building for Multiple Platforms

Using GoReleaser:

```bash
goreleaser release --snapshot --clean
```

## Docker

### Building Docker Image

```bash
docker build -t irkit-webserver .
```

### Running with Docker Compose (Example)

```yaml
version: '3'
services:
  irkit:
    image: kangaechu/irkit:latest
    ports:
      - "9090:9090"
    restart: unless-stopped
```

## Project Structure

```
irkit-webserver/
├── main.go                    # Main application (HTTP server)
├── command/command/           # Shell scripts for IR control
│   ├── aircon_on.sh          # Air conditioner ON
│   ├── aircon_off.sh         # Air conditioner OFF
│   ├── ceiling_on.sh         # Ceiling light ON
│   ├── ceiling_off.sh        # Ceiling light OFF
│   └── ceiling_full.sh       # Ceiling light FULL
├── .github/workflows/        # CI/CD automation
├── Dockerfile                # Docker container definition
├── .goreleaser.yml          # Release configuration
└── .golangci.yml            # Linting rules
```

## CI/CD

This project uses GitHub Actions for:
- ✅ **Automated Testing** - Runs tests on every push and PR
- 🔍 **Code Linting** - Ensures code quality with golangci-lint
- 📦 **Automated Releases** - Creates multi-platform binaries on version tags
- 🐳 **Docker Publishing** - Automatically builds and pushes Docker images

## Integration Examples

### Home Assistant

```yaml
rest_command:
  aircon_on:
    url: "http://your-server:9090/aircon/on"
  aircon_off:
    url: "http://your-server:9090/aircon/off"
```

### Cron Job

```bash
# Turn on AC every weekday at 5 PM
0 17 * * 1-5 curl http://localhost:9090/aircon/on
```

### Node.js

```javascript
const axios = require('axios');

async function turnOnAC() {
  await axios.get('http://localhost:9090/aircon/on');
}
```

## Troubleshooting

**Server won't start:**
- Check if port 9090 is already in use: `lsof -i :9090`
- Verify Go version: `go version` (should be 1.22+)

**Commands not working:**
- Verify IRKit device IP address in shell scripts
- Ensure IRKit device is on the same network
- Check shell script permissions: `chmod +x command/command/*.sh`
- Test IRKit connectivity: `ping 192.168.1.117`

**Docker container issues:**
- Ensure shell scripts are executable in the container
- Verify network connectivity from container to IRKit device

## License

MIT License - Copyright (c) 2019

See [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Related Projects

- [IRKit](http://getirkit.com/) - The original IRKit hardware and software
- [IRKit API Documentation](http://getirkit.com/#toc-4) - IRKit HTTP API reference

---

Made with ❤️ for home automation enthusiasts
