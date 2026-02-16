# DevOps Lecture Project - Slipper Shop

An e-commerce webshop specializing in novelty slippers, built with Go to demonstrate DevOps principles and best practices. This project serves as the practical foundation for the course "Introduction to DevOps, Continuous Delivery Tools and Mindset" (T3INF4902) at DHBW Stuttgart.

## Table of Contents

- [About the Shop](#about-the-shop)
- [Features](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [DevOps Journey](#devops-journey)
- [Team](#team)
- [Branching Strategy](#branching-strategy)
- [Technology Stack](#technology-stack)

## About the Shop

Our webshop offers a selection of fun novelty slippers:
- Gary Slippers
- Fish Slippers
- Avocado Slippers
- Croissant Slippers
- Cat Slippers

Customers can browse products, authenticate, and place orders through a REST API.

## Features

- Product catalog with detailed product information
- User authentication with JWT
- Order checkout functionality
- RESTful API design

## Getting Started

### Prerequisites

- Go 1.25.7 or higher

### Running the Application

```bash
# Clone and navigate to the project
git clone https://github.com/Terikyy/devops-lecture-project.git
cd devops-lecture-project

# Install dependencies
go mod tidy

# Run the application
go run cmd/main.go
```

The server will start on `http://localhost:8080`.

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

- Kevin Kienle
- Ole Schmid
- Erik von Heyden

**Course**: T3INF4902 - Introduction to DevOps, Continuous Delivery Tools and Mindset
**Instructor**: Robin Lieb
**Institution**: DHBW Stuttgart
**Semester**: Summer 2026

## Branching Strategy

We follow GitHub Flow to support continuous delivery:
- `main` branch is always deployable
- Feature branches are created from `main` for new work
- Pull requests are used for code review and discussion
- After approval, changes are merged back to `main`
- Small, frequent commits over large changes

This approach enables fast feedback cycles and reduces merge conflicts, aligning with DevOps principles.

### Commit Message Format

We use Conventional Commits for clear and consistent commit messages:

```
<type>(<scope>): <description>
```

**Common types:**
- `feat` - New features
- `fix` - Bug fixes
- `docs` - Documentation changes
- `chore` - Maintenance tasks
- `refactor` - Code refactoring
- `test` - Adding or updating tests
- `ci` - CI/CD pipeline changes

Example: `feat(products): add new slipper category`

## Technology Stack

- Go 1.25.7
- JWT Authentication (golang-jwt/jwt/v5)
- Git & GitHub
- Future: Docker, Kubernetes, GitHub Actions, Argo CD, Terraform, Grafana Stack
