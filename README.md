# qualitinvest

**Quality Investing API** - A comprehensive REST API for fundamental stock analysis and screening

[![Go](https://img.shields.io/badge/Go-1.23.1-blue.svg)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-blue.svg)](https://postgresql.org)
[![Alpha Vantage](https://img.shields.io/badge/Alpha%20Vantage-API-green.svg)](https://www.alphavantage.co)

## 📋 Overview

qualitinvest is a Go-based REST API that provides comprehensive financial analysis and stock screening capabilities. It integrates with Alpha Vantage for real-time financial data collection and offers intelligent caching, structured logging with OpenTelemetry, and a clean architecture approach.

## 🏗️ Architecture

### Clean Architecture Pattern

The application follows **Clean Architecture** principles with clear separation of concerns:

```
├── cmd/                    # Application entrypoints
├── internal/
│   ├── alpha_vantage/      # External API client
│   ├── config/            # Configuration management
│   ├── controllers/       # HTTP request handlers
│   ├── models/            # Domain entities & DTOs
│   ├── repository/        # Data access layer
│   └── services/          # Business logic layer
├── pkg/                   # Shared packages (logger, responses)
├── sqitch/               # Database migrations
└── docs/                 # Generated documentation
```

### Key Components

- **Controllers**: HTTP request handling with Gin framework
- **Services**: Business logic with dependency injection
- **Repository**: Database abstraction with PostgreSQL
- **Models**: Domain entities and data transfer objects
- **Alpha Vantage Client**: External API integration for financial data

## 🚀 Features

### ✅ Implemented

- **RESTful API** with JSON responses
- **Intelligent Data Collection** from Alpha Vantage API
- **Automatic Data Fetching** (collects data on-demand if not available)
- **Comprehensive Stock Analysis** (overview, financials, ratios)
- **Stock Screening** with customizable criteria
- **Structured Logging** with OpenTelemetry
- **Database Migrations** with Sqitch
- **Graceful Shutdown** handling
- **Swagger Documentation** generation

### 🔄 Data Collection Strategy

**Smart On-Demand Collection:**
- Checks local PostgreSQL database first
- Automatically fetches from Alpha Vantage if data is missing
- Stores data locally for future requests
- Comprehensive error handling and logging

**Supported Data Types:**
- Company Overview (profile, ratios, fundamentals)
- Income Statements (quarterly & annual)
- Balance Sheets
- Cash Flow Statements
- Earnings data
- Stock prices and volume

## 🗄️ Database Schema

### PostgreSQL with Sqitch Migrations

The database schema includes tables for:

- **companies**: Basic company information
- **company_overviews**: Financial overview and ratios
- **income_statements**: Revenue, expenses, profits
- **balance_sheets**: Assets, liabilities, equity
- **cash_flow_statements**: Operating, investing, financing cash flows
- **earnings**: Historical earnings data
- **dividends**: Dividend history
- **splits**: Stock split information

### Migrations
```bash
cd sqitch
./sqitch deploy db:pg://user:pass@localhost/qualitinvest
```

## 🔧 Installation & Setup

### Prerequisites

- **Go 1.23.1+**
- **PostgreSQL 13+**
- **Sqitch** (for database migrations)
- **Alpha Vantage API Key** (free at https://www.alphavantage.co)

### Quick Start

1. **Clone and setup:**
```bash
git clone <repository-url>
cd qualitinvest
```

2. **Database setup:**
```bash
# Create database
createdb qualitinvest

# Run migrations
cd sqitch
./sqitch deploy db:pg://localhost/qualitinvest
```

3. **Environment configuration:**
```bash
cp env.example .env
# Edit .env with your settings
```

4. **Install dependencies:**
```bash
go mod tidy
```

5. **Run the application:**
```bash
source .env
go run cmd/main.go
```

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `DATABASE_URL` | PostgreSQL connection string | - | ✅ |
| `ALPHA_VANTAGE_API_KEY` | Alpha Vantage API key | - | ✅ |
| `ENV` | Environment (development/production) | development | ❌ |
| `PORT` | HTTP server port | 8080 | ❌ |
| `OTEL_SERVICE_NAME` | Service name for logging | qualitinvest-api | ❌ |
| `OTEL_SERVICE_VERSION` | Service version | 1.0.0 | ❌ |
| `OTEL_LOGS_EXPORTER` | Log exporter | console | ❌ |

### Example .env file:
```bash
# Database
DATABASE_URL=postgres://user:password@localhost:5432/qualitinvest?sslmode=disable

# Alpha Vantage API
ALPHA_VANTAGE_API_KEY=your_api_key_here

# Environment
ENV=development
PORT=8080

# OpenTelemetry (optional)
OTEL_SERVICE_NAME=qualitinvest-api
OTEL_SERVICE_VERSION=1.0.0
OTEL_LOGS_EXPORTER=console
```

## 📡 API Endpoints

### Stock Analysis
```
GET /api/v1/analysis/{symbol}      # Complete company analysis (auto-collects if needed)
GET /api/v1/analysis/{symbol}/valuation # Valuation analysis (planned)
```

### Stock Screening
```
POST /api/v1/screening             # Screen companies by criteria
GET /api/v1/screening/templates    # Get screening templates
```

### Data Management
```
POST /api/v1/collect/{symbol}      # Manual data collection (admin/testing)
```

### Documentation
```
GET /swagger/*                     # Swagger UI
```

## 🔍 Usage Examples

### Get Company Analysis (with automatic data collection)
```bash
curl http://localhost:8080/api/v1/analysis/AAPL
```

### Screen Companies
```bash
curl -X POST http://localhost:8080/api/v1/screening \
  -H "Content-Type: application/json" \
  -d '{
    "min_pe_ratio": 0,
    "max_pe_ratio": 20,
    "min_roe": 0.10,
    "limit": 10
  }'
```

### Manual Data Collection
```bash
curl -X POST http://localhost:8080/api/v1/collect/MSFT
```

## 📊 Logging & Observability

### OpenTelemetry Structured Logging

The application uses **OpenTelemetry** for comprehensive observability:

- **Structured JSON logs** with rich attributes
- **Trace context** propagation
- **Multiple log levels** (DEBUG, INFO, WARN, ERROR)
- **Console export** (easily configurable for Loki, Elasticsearch, etc.)

### Log Examples

**Data Collection:**
```json
{
  "Timestamp": "2025-12-21T20:18:52.197358902+01:00",
  "Severity": 9,
  "Body": {"Type":"String","Value":"🔄 Données manquantes - démarrage collecte automatique"},
  "Attributes": [
    {"Key": "symbol", "Value": {"Type": "String", "Value": "MSFT"}}
  ]
}
```

**API Request:**
```json
{
  "Body": {"Type":"String","Value":"📥 Requête d'analyse d'entreprise reçue"},
  "Attributes": [
    {"Key": "symbol", "Value": {"Type": "String", "Value": "AAPL"}},
    {"Key": "endpoint", "Value": {"Type": "String", "Value": "/api/v1/analysis/AAPL"}},
    {"Key": "method", "Value": {"Type": "String", "Value": "GET"}}
  ]
}
```

## 🧪 Testing

### API Testing
```bash
# Health check
curl http://localhost:8080/swagger/index.html

# Test data collection
curl -X POST http://localhost:8080/api/v1/collect/AAPL

# Test analysis
curl http://localhost:8080/api/v1/analysis/AAPL
```

### Database Testing
```bash
# Check migration status
cd sqitch
./sqitch status db:pg://localhost/qualitinvest

# Verify data
psql qualitinvest -c "SELECT symbol, name FROM company_overviews LIMIT 5;"
```

## 🚀 Deployment

### Production Considerations

1. **Environment Variables**: Use production values
2. **Database**: Configure connection pooling
3. **Logging**: Configure OTLP exporter for centralized logging
4. **API Keys**: Secure storage for Alpha Vantage key
5. **Rate Limiting**: Consider implementing request limits

### Docker Example
```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/env.example .env
EXPOSE 8080
CMD ["./main"]
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📈 Roadmap

### Short Term
- [ ] Unit tests for critical services
- [ ] Complete Swagger documentation
- [ ] Web interface (React/Vue)
- [ ] Symbol search API endpoint

### Medium Term
- [ ] Grafana dashboards for metrics
- [ ] Intelligent caching strategies
- [ ] Multi-market support
- [ ] Alert system for financial thresholds

### Long Term
- [ ] Machine learning predictions
- [ ] Portfolio optimization
- [ ] Real-time data streaming
- [ ] Mobile application

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- **Alpha Vantage** for providing comprehensive financial data APIs
- **OpenTelemetry** community for observability standards
- **Gin Framework** for the excellent HTTP router
- **PostgreSQL** for reliable data storage

---

**Built with ❤️ for quality investing insights**