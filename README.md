# DevOps Lecture Project - Go Webshop

A modern e-commerce webshop built with Go, demonstrating DevOps principles and best practices throughout its development lifecycle. This project serves as the practical foundation for the course "Introduction to DevOps, Continuous Delivery Tools and Mindset" (T3INF4902) at DHBW Stuttgart.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Features](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [DevOps Journey](#devops-journey)
- [Team](#team)
- [Technology Stack](#technology-stack)

## Overview

This webshop application specializes in selling novelty slippers and serves as a hands-on project to implement a complete DevOps pipeline from code to production. The project evolves throughout the semester, incorporating industry-standard tools and practices including containerization, CI/CD, Kubernetes orchestration, GitOps, Infrastructure as Code, observability, and security scanning.

The application is built with Go's standard library following clean architecture principles, ensuring modularity, testability, and maintainability.

## Project Structure

The project follows Go's standard project layout conventions:

```
devops-lecture-project/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/                   # Private application code
│   ├── auth/
│   │   └── handler.go         # Authentication handlers
│   ├── checkout/
│   │   └── handler.go         # Order processing handlers
│   └── products/
│       ├── handler.go         # Product HTTP handlers
│       ├── product.go         # Product domain models
│       └── repository.go      # Product data access
├── pkg/                       # Public reusable packages
│   ├── httputil/
│   │   └── response.go       # HTTP response helpers
│   └── jwt/
│       └── jwt.go            # JWT authentication utilities
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
└── README.md
```

### Directory Explanation

- **cmd/**: Contains the main application entry point and executable definitions
- **internal/**: Houses private application code that cannot be imported by other projects
- **pkg/**: Contains public packages that can be reused across projects

## Features

- **Product Management**: Browse and view detailed information about products
- **Authentication System**: User login and logout functionality with JWT-based authentication
- **Checkout Process**: Order placement and processing
- **RESTful API**: Clean HTTP API following REST principles
- **Modular Architecture**: Separation of concerns with clear package boundaries

## Getting Started

### Prerequisites

- Go 1.25.7 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Terikyy/devops-lecture-project.git
cd devops-lecture-project
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080`.

### Building

To create a production binary:

```bash
go build -o webshop cmd/main.go
./webshop
```

## API Endpoints

### Authentication

- **POST** `/auth/login` - User authentication
- **POST** `/auth/logout` - User logout

### Products

- **GET** `/products` - List all available products
- **GET** `/products/{id}` - Get detailed information about a specific product

### Checkout

- **POST** `/checkout/placeorder` - Place a new order

### Example Usage

List all products:
```bash
curl http://localhost:8080/products
```

Get specific product:
```bash
curl http://localhost:8080/products/1
```

## DevOps Journey

This project serves as the foundation for implementing a comprehensive DevOps pipeline across eight lectures:

| Week | Topic | Implementation |
|------|-------|----------------|
| 1 | Version Control & Git | GitHub repository setup, branching strategies, pull requests |
| 2 | From VMs to Containers | Dockerize the application |
| 3 | CI/CD & Testing | Implement GitHub Actions pipelines |
| 4 | Container Orchestration | Deploy to local Kubernetes cluster |
| 5 | GitOps & Progressive Delivery | Implement GitOps with Argo CD |
| 6 | Infrastructure as Code | Provision infrastructure with Terraform/OpenTofu |
| 7 | Observability & Resilience | Integrate LGTM Stack (Loki, Grafana, Tempo, Mimir) |
| 8 | DevSecOps & Platform Engineering | Security scanning, SAST, and platform engineering practices |

## Team

This project is developed as part of the DevOps course at DHBW Stuttgart:

- **Kevin Kienle** - 5409606
- **Ole Schmid** - 5070696
- **Erik von Heyden** - 8720832

**Course**: T3INF4902 - Introduction to DevOps, Continuous Delivery Tools and Mindset
**Instructor**: Robin Lieb
**Institution**: DHBW Stuttgart
**Semester**: Winter 2025/2026

## Technology Stack

- **Language**: Go 1.25.7
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **HTTP Router**: Go standard library `net/http`
- **Version Control**: Git & GitHub
- **Future Technologies**: Docker, Kubernetes, GitHub Actions, Argo CD, Terraform, Grafana Stack

## Development Principles

This project adheres to key DevOps principles:

- **Flow Principle**: Optimize the entire development pipeline, not just individual stages
- **Feedback Principle**: Fast feedback loops through automated testing and monitoring
- **Continuous Learning**: Iterative improvements based on metrics and retrospectives
- **Small Batch Sizes**: Frequent, small commits over large, infrequent changes
- **Trunk-Based Development**: Short-lived feature branches with regular integration

## Contributing

This is an educational project. Contributions are managed through pull requests with mandatory code review. All changes must follow conventional commit messages and semantic versioning principles.

### Commit Message Format

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

**Types**: `feat`, `fix`, `docs`, `chore`, `refactor`, `test`, `ci`

### Branching Strategy

We follow GitHub Flow:
- Main branch is always deployable
- Feature branches are created from main
- Pull requests require approval before merging
- Automated tests must pass before merge

## License

This project is created for educational purposes as part of the DevOps curriculum at DHBW Stuttgart.

---

**Repository**: https://github.com/Terikyy/devops-lecture-project
**Last Updated**: February 2026
