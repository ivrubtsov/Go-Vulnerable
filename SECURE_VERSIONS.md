# Expected Secure Versions

This file contains the expected secure versions that your AI agent should update to.

## Current Vulnerable Dependencies

```go
// VULNERABLE - DO NOT USE IN PRODUCTION
require (
	github.com/gin-gonic/gin v1.7.0
	github.com/gorilla/websocket v1.4.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	gopkg.in/yaml.v2 v2.2.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/opencontainers/runc v1.0.0-rc1
)
```

## Recommended Secure Versions

```go
// SECURE VERSIONS (as of January 2025)
require (
	github.com/gin-gonic/gin v1.9.1
	github.com/gorilla/websocket v1.5.1
	golang.org/x/crypto v0.18.0
	golang.org/x/net v0.20.0
	gopkg.in/yaml.v3 v3.0.1  // Note: v2 -> v3 migration required
	github.com/golang-jwt/jwt/v5 v5.2.0  // Note: package replacement required
	github.com/opencontainers/runc v1.1.12
)
```

## Migration Notes

### 1. gopkg.in/yaml.v2 → gopkg.in/yaml.v3

The migration from v2 to v3 requires code changes:

**Before (v2):**
```go
import "gopkg.in/yaml.v2"

var config Config
err := yaml.Unmarshal(data, &config)
```

**After (v3):**
```go
import "gopkg.in/yaml.v3"

var config Config
err := yaml.Unmarshal(data, &config)
```

The import path changes, but the API is mostly compatible.

### 2. github.com/dgrijalva/jwt-go → github.com/golang-jwt/jwt/v5

This is a package replacement due to the original being unmaintained.

**Before (dgrijalva/jwt-go):**
```go
import "github.com/dgrijalva/jwt-go"

token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": user.Username,
})
tokenString, err := token.SignedString([]byte("secret"))
```

**After (golang-jwt/jwt/v5):**
```go
import "github.com/golang-jwt/jwt/v5"

token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": user.Username,
})
tokenString, err := token.SignedString([]byte("secret"))
```

The API is backward compatible for most use cases.

## AI Agent Implementation Checklist

Your AI agent should:

- [ ] Parse `go.mod` to identify vulnerable dependencies
- [ ] Parse Trivy JSON output to get CVE details
- [ ] Map vulnerable packages to secure versions
- [ ] Handle package replacements (jwt-go → golang-jwt/jwt)
- [ ] Handle major version upgrades (yaml.v2 → yaml.v3)
- [ ] Update import statements in `.go` files if needed
- [ ] Run `go mod tidy` after updates
- [ ] Verify the application builds with `go build`
- [ ] Re-run Trivy to confirm vulnerabilities are resolved
- [ ] Generate a report of changes made

## Verification Commands

After your AI agent updates the dependencies:

```bash
# Tidy up dependencies
go mod tidy

# Verify module integrity
go mod verify

# Build the application
go build -o vulnerable-app

# Re-scan with Trivy
trivy fs --severity HIGH,CRITICAL .

# The scan should show 0 HIGH/CRITICAL vulnerabilities
```

## Success Criteria

After your AI agent completes the update:

1. All dependencies updated to secure versions
2. Application builds successfully
3. Trivy scan shows 0 HIGH or CRITICAL vulnerabilities
4. Import statements updated where necessary
5. `go.mod` and `go.sum` are properly formatted
