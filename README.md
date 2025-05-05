# 🪽 Wingobank Backend

Welcome to the **Wingobank Backend** — a modern, modular, and scalable banking platform built using microservices architecture in Go.

## 🧩 Microservices Included

- `user-service`: User management, registration, profiles
- `account-service`: Bank accounts and balances
- `transaction-service`: Transfers between accounts
- `payment-service`: External/internal payments
- `notification-service`: Emails, SMS, push notifications
- `auth-service`: Login, JWT, OAuth2
- `ledger-service`: Internal double-entry bookkeeping
- `fraud-service`: Fraud detection and risk scoring
- `kyc-service`: Identity verification (KYC/AML)
- `reporting-service`: Financial reports and analytics

## 🛠 Tech Stack

- **Go** (Golang) for services
- **PostgreSQL** as the main database
- **Apache Kafka** for asynchronous events
- **Docker** and **Docker Compose** for local development
- **Kubernetes** for orchestration
- **GitHub Actions** for CI/CD
- **Terraform** for infrastructure provisioning

## 🚀 Getting Started

To spin up the environment locally:

```bash
docker-compose up --build
```

Each service is containerized and will log its own startup message to the terminal.

## 📁 Repository Structure

```bash
wingobank-backend/
├── services/              # Microservices
├── api-gateway/           # Gateway/API routing
├── deployments/           # k8s, Helm, Docker configs
├── kafka/                 # Kafka topic configs
├── proto/                 # gRPC/protobuf (optional)
├── infra/                 # Terraform, scripts
├── .github/workflows/     # CI/CD pipelines
├── docker-compose.yml     # Dev environment
└── README.md              # This file
```

---

### 🧠 Philosophy

This project was built with the idea of mimicking the internal structure of a real bank, but in a microbank scale. Each service is decoupled, scalable, and testable, following modern architectural principles.

> **"Move fast. Scale smart. Build trust."**

---

### 📬 Questions?

Open an issue, or send a PR. Contributions are always welcome! 😄

---
