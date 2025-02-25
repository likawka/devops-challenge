# Software Engineering Intern DevOps home challenge

## Overview

This project is a **Golang-based API** for managing deals, with a structured architecture following best practices. It supports **RESTful API endpoints**, **Swagger documentation**, and **Docker deployment**. The project also includes **CI/CD pipelines** for automated builds and deployments.

## Features

* **Versioned API Endpoints** (e.g., `<span>/v1/deals</span>`)
* **Swagger Documentation**
* **Docker Support** for easy deployment
* **Metrics Middleware** for logging API performance
* **CI/CD Pipelines** using GitHub Actions
* **Metrics** available at `/metrics`

## Project Structure

```
.github/workflows/   # CI/CD configurations
├── cd.yml          # Continuous Deployment pipeline
├── ci.yml          # Continuous Integration pipeline
api/                 # API logic
├── deals.go        # Deals API handler
handlers/            # HTTP handlers
├── deals.go        # Deal request handlers
├── swagger.go      # Swagger setup
└── docs/           # API documentation
    ├── swagger.json
    ├── swagger.yaml
.env                 # Environment variables (not committed)
.gitignore           # Git ignored files
Dockerfile           # Docker build instructions
docker-compose.yml   # Docker Compose configuration
go.mod, go.sum       # Go dependencies
main.go              # Application entry point
README.md            # Project documentation
```

## Getting Started

### 1. Prerequisites

Ensure you have the following installed:

* **Golang** (1.21)
* **Docker** & **Docker Compose**

### 2. Running Locally

1. **Clone the repository:**

   ```
   git clone https://github.com/likawka/devops-challenge.git
   cd devops-challenge
   ```
2. **Set up environment variables:**

   * Create a `<span>.env</span>` file in the project root and add the necessary variables, e.g.:
     ```
     PIPEDRIVE_API_KEY=your_api_key_here
     ```
3. **Run the application:**

   ```
   go mod tidy
   go run main.go
   ```

   or with Bash:

   ```
   bash devRun.sh
   ```
4. **Swagger UI:** Open `<span>http://localhost:8080/swagger/index.html</span>`.
5. **Metrics Endpoint:** Available at `<span>http://localhost:8080/metrics</span>`.

### 3. Running with Docker

For **Mac (M1/M2)**, ensure you build with the correct platform:

```
docker build --platform=linux/arm64 -t devops-challenge .
docker run -p 8080:8080 --env-file .env devops-challenge
```

For general use:

```
docker-compose up --build
```

### 4. API Documentation

Swagger documentation is automatically generated. After running the server, visit:

```
http://localhost:8080/swagger/index.html
```

### 5. CI/CD Pipelines

The project includes **CI/CD automation** with GitHub Actions:

* **Continuous Integration (CI):** Runs tests and lints the code.
* **Continuous Deployment (CD):** Builds and deploys the application automatically.

### 6. Author

**Iryna Zinko**
[LinkedIn](https://www.linkedin.com/in/iryna-zinko-143a4125b/)
