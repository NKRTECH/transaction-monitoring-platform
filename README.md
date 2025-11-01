# Transaction Monitoring Platform

[![Java](https://img.shields.io/badge/Java-21-orange.svg)](https://openjdk.java.net/projects/jdk/21/)
[![Spring Boot](https://img.shields.io/badge/Spring%20Boot-3.3.5-brightgreen.svg)](https://spring.io/projects/spring-boot)
[![Go](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![Next.js](https://img.shields.io/badge/Next.js-15-black.svg)](https://nextjs.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://www.docker.com/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Ready-blue.svg)](https://kubernetes.io/)

Enterprise-grade microservices platform for financial transaction monitoring, compliance, and real-time analytics. Built with modern technologies and patterns used by major fintech companies.

## 🏗️ Architecture Overview

This platform demonstrates **true microservices architecture** with polyglot implementation:

- **Transaction Service** (Java Spring Boot) - Core transaction processing
- **Validation Service** (Go) - High-performance validation engine  
- **Analytics Service** (Java Spring Boot) - Real-time analytics with ClickHouse
- **Alert Service** (Go) - Event-driven alerting and notifications
- **Dashboard** (Next.js) - Real-time monitoring interface
- **API Gateway** (Kong) - Centralized routing and security

## 🚀 Quick Start

### Prerequisites
- **Java 21+** (for Spring Boot services)
- **Go 1.21+** (for Go services)
- **Node.js 18+** (for Next.js dashboard)
- **Docker & Docker Compose** (for infrastructure)
- **Git** (for version control)

### 1. Clone Repository
```bash
git clone https://github.com/yourusername/transaction-monitoring-platform.git
cd transaction-monitoring-platform
```

### 2. Start Infrastructure
```bash
# Start PostgreSQL, Redis, Kafka, ClickHouse
docker-compose up -d postgres redis kafka clickhouse
```

### 3. Run Services

#### Transaction Service (Java Spring Boot)
```bash
cd services/transaction-service
./gradlew bootRun
# Service available at: http://localhost:8080
```

#### Validation Service (Go) - Coming Soon
```bash
cd services/validation-service
go run cmd/main.go
# Service available at: http://localhost:8081
```

#### Dashboard (Next.js) - Coming Soon
```bash
cd frontend/dashboard
npm run dev
# Dashboard available at: http://localhost:3000
```

## 📊 Service Status

| Service | Technology | Port | Status | Health Check |
|---------|------------|------|--------|--------------|
| Transaction Service | Java Spring Boot | 8080 | ✅ **Live** | [Health](http://localhost:8080/api/health) |
| Validation Service | Go | 8081 | 🔄 In Progress | Coming Soon |
| Analytics Service | Java Spring Boot | 8082 | ⏳ Planned | Coming Soon |
| Alert Service | Go | 8083 | ⏳ Planned | Coming Soon |
| Dashboard | Next.js | 3000 | ⏳ Planned | Coming Soon |

## 🏢 Enterprise Features

### ✅ **Currently Implemented**
- **Independent Microservices**: Each service is a separate project with own build system
- **Health Monitoring**: Comprehensive health checks and readiness probes
- **Hot Reloading**: Development-friendly with automatic reloading
- **Containerization**: Docker-ready with multi-stage builds
- **API Documentation**: OpenAPI/Swagger integration
- **Testing**: Unit and integration tests with high coverage

### 🔄 **In Development**
- **Go Validation Service**: High-performance validation with gRPC
- **Event-Driven Architecture**: Kafka-based inter-service communication
- **Real-Time Dashboard**: Next.js frontend with live updates

### ⏳ **Planned Features**
- **Real-Time Analytics**: ClickHouse OLAP with sub-second queries
- **Stream Processing**: Apache Flink for fraud detection
- **Kubernetes Deployment**: Production-ready K8s manifests
- **Observability Stack**: Prometheus, Grafana, Jaeger tracing

## 🛠️ Technology Stack

### **Backend Services**
- **Java 21** with **Spring Boot 3.3.5** - Enterprise framework
- **Go 1.21+** with **Gin** - High-performance services
- **PostgreSQL 15** - ACID-compliant OLTP database
- **Redis 7** - Caching and session storage
- **Apache Kafka** - Event streaming and messaging
- **ClickHouse** - High-performance OLAP analytics

### **Frontend & API**
- **Next.js 15** with **TypeScript** - Modern React framework
- **Tailwind CSS** - Utility-first styling
- **Kong API Gateway** - Service routing and security
- **OpenAPI/Swagger** - API documentation

### **Infrastructure & DevOps**
- **Docker & Docker Compose** - Containerization
- **Kubernetes** - Container orchestration
- **Prometheus & Grafana** - Monitoring and alerting
- **Jaeger** - Distributed tracing
- **GitHub Actions** - CI/CD pipeline

## 📁 Project Structure

```
transaction-monitoring-platform/
├── services/                    # Independent microservices
│   ├── transaction-service/     # Java Spring Boot (✅ Implemented)
│   ├── validation-service/      # Go application (🔄 In Progress)
│   ├── analytics-service/       # Java Spring Boot (⏳ Planned)
│   └── alert-service/          # Go application (⏳ Planned)
├── frontend/
│   └── dashboard/              # Next.js application (⏳ Planned)
├── infrastructure/
│   ├── docker-compose.yml      # Local development setup
│   ├── kubernetes/             # K8s deployment manifests
│   └── monitoring/             # Prometheus, Grafana configs
├── shared/
│   └── proto/                  # gRPC protocol definitions
└── .kiro/
    └── specs/                  # Feature specifications and tasks
```

## 🔗 API Endpoints

### Transaction Service (Port 8080) ✅
- **Health**: `GET /api/health`
- **API Docs**: `GET /swagger-ui.html`
- **Actuator**: `GET /actuator/health`

### Validation Service (Port 8081) 🔄
- **Health**: `GET /api/health` (Coming Soon)
- **Validate**: `POST /api/validate` (Coming Soon)
- **gRPC**: Port 9081 (Coming Soon)

*See [Service Communication](/.kiro/specs/transaction-monitoring-dashboard/service-communication.md) for complete API reference.*

## 🧪 Testing

### Run All Tests
```bash
# Transaction Service
cd services/transaction-service && ./gradlew test

# Validation Service (when implemented)
cd services/validation-service && go test ./...

# Integration Tests
docker-compose -f docker-compose.test.yml up --abort-on-container-exit
```

### Performance Testing
```bash
# Load testing with k6
k6 run tests/load/transaction-service.js
```

## 🚀 Deployment

### Local Development
```bash
# Start all infrastructure
docker-compose up -d

# Run services individually (see Quick Start)
```

### Production (Kubernetes)
```bash
# Deploy to Kubernetes
kubectl apply -f infrastructure/kubernetes/

# Monitor deployment
kubectl get pods -n transaction-monitoring
```

## 📈 Performance Targets

- **API Response Time**: < 200ms (p95)
- **Transaction Throughput**: 10,000+ TPS
- **Analytics Query Time**: < 1 second (100M+ records)
- **System Availability**: 99.9%
- **Event Processing Latency**: < 100ms

## 🤝 Contributing

1. **Fork** the repository
2. **Create** feature branch: `git checkout -b feature/amazing-feature`
3. **Commit** changes: `git commit -m 'feat: add amazing feature'`
4. **Push** to branch: `git push origin feature/amazing-feature`
5. **Open** Pull Request

### Commit Convention
We use [Conventional Commits](https://www.conventionalcommits.org/):
- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation updates
- `test:` - Test additions
- `refactor:` - Code refactoring

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🏆 Enterprise Patterns Demonstrated

- **Microservices Architecture** - Independent, scalable services
- **Event-Driven Design** - Kafka-based asynchronous communication
- **CQRS & Event Sourcing** - Separate read/write models
- **API Gateway Pattern** - Centralized routing and security
- **Circuit Breaker** - Resilience and fault tolerance
- **Health Checks** - Kubernetes-ready probes
- **Observability** - Metrics, logging, and tracing
- **Infrastructure as Code** - Docker and Kubernetes manifests

## 🎯 Use Cases

- **Financial Compliance** - Regulatory reporting and monitoring
- **Fraud Detection** - Real-time transaction analysis
- **Risk Management** - Compliance metrics and alerting
- **Operational Efficiency** - Performance monitoring and optimization
- **Audit Trail** - Complete transaction lifecycle tracking

---

**Built with ❤️ for enterprise-grade financial technology**

*This project demonstrates production-ready microservices architecture suitable for fintech companies, banks, and financial institutions.*