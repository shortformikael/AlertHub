# Environment Variables Best Practices for AlertHub

## Overview

This document outlines the best practices for handling environment variables in the AlertHub project, covering both development and production environments.

## How Environment Variables Work

1. **OS Environment Variables** - Always available via `os.Getenv()`
2. **`.env` files** - Only loaded by packages like `godotenv` when explicitly called
3. **Priority** - OS env vars typically override `.env` file values

## Best Practices for Production

### 1. Use OS Environment Variables in Production

```go
// Production-ready approach
func Init() error {
    // Only load .env in development
    if os.Getenv("APP_ENV") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Println("No .env file found, using system environment variables")
        }
    }
    
    // Always read from OS environment variables
    Port = getEnvWithDefault("PORT", "8080")
    Name = getEnvWithDefault("NAME", "AlertHub")
    // etc...
}

func getEnvWithDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

### 2. Recommended Production Setup

**Development:**
- Use `.env` files for local development
- Load with `godotenv`

**Production:**
- Set environment variables directly on the system
- Use container orchestration (Docker, Kubernetes)
- Use secrets management (HashiCorp Vault, AWS Secrets Manager)

### 3. Docker Example

```dockerfile
# Dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o alerthub main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/alerthub .
CMD ["./alerthub"]
```

```yaml
# docker-compose.yml
version: '3.8'
services:
  alerthub:
    build: .
    environment:
      - PORT=8080
      - NAME=AlertHub
      - VERSION=1.0.0
    ports:
      - "8080:8080"
```

### 4. Kubernetes Example

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: alerthub
spec:
  template:
    spec:
      containers:
      - name: alerthub
        image: alerthub:latest
        env:
        - name: PORT
          value: "8080"
        - name: NAME
          value: "AlertHub"
        # Or use ConfigMaps/Secrets
        envFrom:
        - configMapRef:
            name: alerthub-config
```

## Recommended Implementation for AlertHub

```go
func Init() error {
    // Load .env only in development
    if os.Getenv("APP_ENV") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Println("No .env file found")
        }
    }
    
    // Read from environment with defaults
    Port = getEnvWithDefault("ALERTHUB_PORT", "8080")
    Name = getEnvWithDefault("ALERTHUB_NAME", "AlertHub")
    Version = getEnvWithDefault("ALERTHUB_VERSION", "0.1")
    
    // Use PWD for current directory, but allow override
    Dir.Current = getEnvWithDefault("ALERTHUB_CURRENT_DIR", os.Getenv("PWD"))
    Dir.Root = getEnvWithDefault("ALERTHUB_ROOT_DIR", "/"+Name)
    Dir.Tmp = getEnvWithDefault("ALERTHUB_TMP_DIR", "/tmp")
    Dir.Data = getEnvWithDefault("ALERTHUB_DATA_DIR", Dir.Tmp+"/"+Name+"/data")
    Dir.Logs = getEnvWithDefault("ALERTHUB_LOGS_DIR", Dir.Tmp+"/"+Name+"/logs")
    
    return makeAllDirectories()
}

func getEnvWithDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

## Environment Variable Packages

### 1. godotenv (joho/godotenv)
```bash
go get github.com/joho/godotenv
```

**Pros:**
- Most popular (6k+ stars)
- Simple API
- Supports multiple .env files
- Well-maintained

### 2. viper (spf13/viper)
```bash
go get github.com/spf13/viper
```

**Pros:**
- Very powerful configuration management
- Supports JSON, YAML, TOML, HCL, env files
- Automatic type conversion
- Great for complex applications

### 3. env (caarlos0/env)
```bash
go get github.com/caarlos0/env/v6
```

**Pros:**
- Struct-based configuration
- Automatic type conversion
- Default values
- Very clean API

## Why This Approach

1. **Development** - Easy to use `.env` files
2. **Production** - Relies on OS environment variables
3. **Flexibility** - Can override any setting via environment
4. **Security** - No sensitive data in code or files
5. **Container-friendly** - Works well with Docker/Kubernetes

## Key Points

- Your binary will always read OS environment variables
- `.env` files are just a convenience for development
- Production should rely on OS environment variables
- Use environment-specific prefixes (e.g., `ALERTHUB_PORT`) to avoid conflicts
- Always provide sensible defaults