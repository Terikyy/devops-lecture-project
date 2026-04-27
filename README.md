# DevOps Lecture Project - Slipper Shop

An e-commerce webshop specializing in novelty slippers, built with Go to demonstrate DevOps principles and best practices. This project serves as the practical foundation for the course "Introduction to DevOps, Continuous Delivery Tools and Mindset" (T3INF4902) at DHBW Stuttgart.

## Table of Contents

- [About the Shop](#about-the-shop)
- [Features](#features)
- [Getting Started](#getting-started)
- [Running with Docker](#running-with-docker)
- [Deploying with ArgoCD](#deploying-with-argocd)
- [Infrastructure as Code](#infrastructure-as-code)
- [API Endpoints](#api-endpoints)
- [DevOps Journey](#devops-journey)
- [Team](#team)
- [Branching Strategy](#branching-strategy)
- [Known Issues](#known-issues)

## About the Shop

Our webshop offers a selection of fun novelty slippers:
- Gary Slippers
- Fish Slippers
- Avocado Slippers
- Croissant Slippers
- Cat Slippers

Customers can browse products, authenticate, and place orders through a REST API.

## Features

- Microservices architecture (Auth, Checkout, Product)
- Product catalog with detailed product information
- User authentication with JWT
- Order checkout functionality
- Automated releases and changelog generation via Release Please
- Infrastructure as Code (OpenTofu/Terraform)
- GitOps deployment with ArgoCD
- Comprehensive observability (LGTM stack)

## Getting Started

### Prerequisites

- Go 1.25.7 or higher
- Make (optional, for convenience commands)

### Running the Application

The project is structured as a monorepo containing multiple microservices. You can run or build each service individually.

```bash
# Clone and navigate to the project
git clone https://github.com/Terikyy/devops-lecture-project.git
cd devops-lecture-project

# Example: Run the Auth Service
cd services/auth-service
go mod tidy
go run cmd/main.go
```

Alternatively, you can use the provided `Makefile` from the root directory:

```bash
# Build a specific service
make build service=auth-service

# Run tests for a service
make test service=checkout-service
```

## Running with Docker

### Pull from Docker Hub

```bash
docker pull oleschmid/slipper-shop-auth-service:latest
docker pull oleschmid/slipper-shop-product-service:latest
docker pull oleschmid/slipper-shop-checkout-service:latest
```

### Build Locally

```bash
# Build one service image
docker build --build-arg SERVICE=auth-service -t slipper-shop-auth-service:local .

# Run one service image
docker run -p 8080:8080 slipper-shop-auth-service:local
```

The application will be available at `http://localhost:8080`.

**Docker Hub**:
- https://hub.docker.com/u/oleschmid

## Deploying with ArgoCD

Bootstrap ArgoCD once by applying the root app:

```bash
kubectl apply -f argocd/root.yml
```

ArgoCD will then automatically manage everything from there:

1. The `root` Application syncs `argocd/applicationsets/`
2. Two ApplicationSets are deployed — one for the slipper shop services, one for monitoring
3. Each ApplicationSet uses a Git directory generator to auto-discover and deploy services from the corresponding `kubernetes/` subdirectory

**Repository structure:**

```
argocd/
  root.yml                          # Apply this once manually
  applicationsets/
    slipper-shop-services.yml       # Auto-discovers kubernetes/slipper-shop-services/*
    monitoring.yml                  # Auto-discovers kubernetes/monitoring/*

kubernetes/
  slipper-shop-services/            # One subdirectory per service
    auth-service/
    checkout-service/
    product-service/
  monitoring/                       # Monitoring stack manifests go here
```

Adding a new service only requires creating a new subdirectory under `kubernetes/slipper-shop-services/` — no ArgoCD manifest needed.

## Infrastructure as Code

The cloud infrastructure (Azure Kubernetes Service, Resource Groups, and initially bootstrapping ArgoCD) is managed via OpenTofu (Terraform). Detailed instructions for initializing and deploying the Azure resources can be found in the [Infrastructure README](infrastructure/README.md).

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
| 6 | Observability & Resilience | Integrate LGTM Stack (Loki, Grafana, Tempo, Mimir) |
| 7 | DevSecOps & Platform Engineering | Security scanning, SAST, and platform engineering practices |
| 8 | Infrastructure as Code | Provision infrastructure with Terraform/OpenTofu |

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

### Automated Releases & Versioning

This repository leverages Google's **Release Please** to automate semantic versioning and changelog generation. 
Based on the Conventional Commits, Release Please automatically:
- Bumps semantic versions for our microservices (under `services/`)
- Generates and updates `CHANGELOG.md` files
- Creates Github Releases

## Known Issues

The application runs seamlessly in local environments. However, when deploying the infrastructure to Azure using OpenTofu on an Azure for Students subscription, individual pods might not start successfully. This is a known issue caused by insufficient resource limits available on the student plan.
