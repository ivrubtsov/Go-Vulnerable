# Vulnerable Go Application for Testing

This is a sample Go application with **intentionally vulnerable dependencies** designed for testing security scanning tools like Trivy and AI agents that perform dependency updates.

## ⚠️ WARNING

**DO NOT USE THIS APPLICATION IN PRODUCTION!**

This application contains known security vulnerabilities and is intended solely for testing purposes.

## Known Vulnerabilities

This application includes the following vulnerable dependencies:

### 1. **github.com/gin-gonic/gin v1.7.0**
- Multiple CVEs including path traversal and denial of service
- Should be updated to v1.9.1 or later

### 2. **github.com/gorilla/websocket v1.4.0**
- CVE-2020-27813: Denial of Service vulnerability
- Should be updated to v1.5.0 or later

### 3. **golang.org/x/crypto (2019 version)**
- Multiple cryptographic vulnerabilities
- Outdated cryptographic implementations
- Should be updated to v0.17.0 or later

### 4. **golang.org/x/net (2019 version)**
- HTTP/2 rapid reset attack vulnerabilities
- Should be updated to v0.19.0 or later

### 5. **gopkg.in/yaml.v2 v2.2.2**
- CVE-2019-11254: Arbitrary code execution via malicious YAML
- Should be updated to v2.4.0 or later (or migrate to v3)

### 6. **github.com/dgrijalva/jwt-go v3.2.0**
- CVE-2020-26160: JWT validation bypass
- This package is unmaintained
- Should be replaced with github.com/golang-jwt/jwt v5.x

### 7. **github.com/opencontainers/runc v1.0.0-rc1**
- Multiple container escape vulnerabilities
- Should be updated to v1.1.12 or later

## Running Trivy Scan

To scan this application with Trivy:

```bash
# Scan the filesystem
trivy fs .

# Scan with specific severity levels
trivy fs --severity HIGH,CRITICAL .

# Generate JSON output
trivy fs --format json --output results.json .

# Scan go.mod file specifically
trivy config go.mod
```

## Expected Trivy Findings

When you run Trivy on this application, you should see:

- Multiple HIGH and CRITICAL severity vulnerabilities
- Detailed CVE information for each vulnerable dependency
- Recommendations for version updates
- Total vulnerability count across all dependencies

## Testing Your AI Agent

This application is perfect for testing AI agents that:

1. Detect vulnerable dependencies
2. Automatically update `go.mod` with secure versions
3. Verify compatibility after updates
4. Run tests to ensure nothing breaks
5. Generate pull requests or reports

### Recommended Update Versions

Your AI agent should update to these (or newer) versions:

```go
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/gorilla/websocket v1.5.1
    golang.org/x/crypto v0.17.0
    golang.org/x/net v0.19.0
    gopkg.in/yaml.v3 v3.0.1
    github.com/golang-jwt/jwt/v5 v5.2.0
    github.com/opencontainers/runc v1.1.12
)
```

Note: `dgrijalva/jwt-go` should be replaced entirely with `golang-jwt/jwt`.

## Building and Running (for testing only)

```bash
# Download dependencies
go mod download

# Build the application
go build -o vulnerable-app

# Run (DO NOT expose to internet!)
./vulnerable-app
```

## API Endpoints

- `POST /login` - JWT authentication (vulnerable)
- `GET /ws` - WebSocket endpoint (vulnerable)
- `POST /hash` - Password hashing (using old crypto)
- `GET /config` - YAML configuration (vulnerable parser)

## License

This is a test application for educational purposes only.
