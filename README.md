# qualitinvest
my quality investing journey

# WORKING WITH THE PROJECT
This section gives advices on how to run this project Go API project with OpenTelemetry logging

# OpenTelemetry Logging
This project uses OpenTelemetry for structured logging. Logs are exported to console by default.

## Configuration
Copy `env.example` to `.env` and configure your environment variables:

```bash
cp env.example .env
```

Available OpenTelemetry environment variables:
- `OTEL_SERVICE_NAME`: Service name (default: qualitinvest-api)
- `OTEL_SERVICE_VERSION`: Service version (default: 1.0.0)
- `OTEL_LOGS_EXPORTER`: Log exporter (default: console)

# Swagger
To generate the swagger files please run the following command

```bash
swag init -g cmd/main.go
```

To access the swagger page, please use the following address localhost:8080/swagger/index.html

## Run The Application
To run the application, please use the following command

```bash
go run cmd/main.go
```

# RESSOURCES

https://www.freecodecamp.org/news/how-to-build-historical-price-charts-with-d3-js-72214aaf6ba3/
https://perso.univ-lemans.fr/~cpiau/BD/SQL_PAGES/SQL0.html