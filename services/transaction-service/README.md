# Transaction Service

Independent Spring Boot microservice for transaction processing and management.

## Overview

The Transaction Service is a standalone microservice responsible for:
- Transaction CRUD operations
- Business rule enforcement
- Transaction validation coordination
- Audit logging and compliance

## Technology Stack

- **Java 21** - Latest LTS version
- **Spring Boot 3.3.5** - Enterprise framework
- **PostgreSQL** - Primary database
- **Gradle 8.10.2** - Build system
- **Docker** - Containerization

## Quick Start

### Prerequisites
- Java 21 or higher
- PostgreSQL 15+ (or use H2 for development)
- Docker (optional)

### Running Locally

1. **Development Mode (H2 Database)**:
   ```bash
   ./gradlew bootRunDev
   ```

2. **With PostgreSQL**:
   ```bash
   # Start PostgreSQL (via Docker Compose from root)
   docker-compose up -d postgres
   
   # Run the service
   ./gradlew bootRun
   ```

3. **Using Docker**:
   ```bash
   docker build -t transaction-service .
   docker run -p 8080:8080 transaction-service
   ```

### API Endpoints

- **Health Check**: `GET /api/health`
- **Readiness**: `GET /api/health/ready`
- **Liveness**: `GET /api/health/live`
- **API Documentation**: `http://localhost:8080/swagger-ui.html`
- **Actuator**: `http://localhost:8080/actuator`

### Development Features

- **Hot Reloading**: Enabled with Spring DevTools
- **H2 Console**: Available at `/h2-console` in dev mode
- **OpenAPI Docs**: Swagger UI for API testing
- **CORS**: Configured for frontend integration

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SPRING_PROFILES_ACTIVE` | Active profile | `dev` |
| `DB_URL` | Database URL | `jdbc:postgresql://localhost:5432/gtrs_transactions` |
| `DB_USERNAME` | Database username | `gtrs_user` |
| `DB_PASSWORD` | Database password | `gtrs_password` |

### Profiles

- **dev**: H2 in-memory database, debug logging
- **prod**: PostgreSQL, optimized logging

## Testing

```bash
# Run all tests
./gradlew test

# Run with coverage
./gradlew test jacocoTestReport

# Integration tests
./gradlew integrationTest
```

## Building

```bash
# Build JAR
./gradlew build

# Build Docker image
docker build -t transaction-service .

# Skip tests for faster builds
./gradlew build -x test
```

## Monitoring

- **Health Endpoints**: `/api/health/*`
- **Metrics**: `/actuator/metrics`
- **Prometheus**: `/actuator/prometheus`

## Architecture

This service follows microservices best practices:
- Independent deployment
- Own database schema
- Stateless design
- Health checks
- Observability

## Next Steps

1. Add JPA entities and repositories
2. Implement transaction REST API
3. Add validation service integration
4. Implement comprehensive testing