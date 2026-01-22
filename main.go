package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Initialize Gin router with vulnerable version
	r := gin.Default()

	// Vulnerable TLS configuration
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS10, // Vulnerable: allows old TLS versions
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_RC4_128_SHA, // Vulnerable cipher
		},
	}
	_ = tlsConfig

	// Routes
	r.POST("/login", loginHandler)
	r.GET("/ws", handleWebSocket)
	r.POST("/hash", hashPassword)
	r.GET("/config", loadConfig)

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func loginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Using vulnerable JWT library
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})

	// Weak secret key
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func handleWebSocket(c *gin.Context) {
	// Using vulnerable websocket library
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func hashPassword(c *gin.Context) {
	var req struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Using vulnerable bcrypt from golang.org/x/crypto
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hash": string(hash),
	})
}

func loadConfig(c *gin.Context) {
	// Using vulnerable YAML parser
	yamlData := `
server:
  host: localhost
  port: 8080
database:
  host: localhost
  port: 5432
  username: admin
  password: secret123
`

	var config Config
	err := yaml.Unmarshal([]byte(yamlData), &config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse config"})
		return
	}

	// Using vulnerable context package
	ctx := context.Background()
	_ = ctx

	c.JSON(http.StatusOK, config)
}

func demonstrateVulnerabilities() {
	fmt.Println("This application contains multiple known vulnerabilities:")
	fmt.Println("1. gin-gonic/gin v1.7.0 - Multiple CVEs")
	fmt.Println("2. gorilla/websocket v1.4.0 - Denial of Service vulnerabilities")
	fmt.Println("3. golang.org/x/crypto (old version) - Cryptographic vulnerabilities")
	fmt.Println("4. golang.org/x/net (old version) - HTTP/2 vulnerabilities")
	fmt.Println("5. gopkg.in/yaml.v2 v2.2.2 - Arbitrary code execution")
	fmt.Println("6. dgrijalva/jwt-go v3.2.0 - JWT validation bypass")
	fmt.Println("7. opencontainers/runc v1.0.0-rc1 - Container escape vulnerabilities")
}
