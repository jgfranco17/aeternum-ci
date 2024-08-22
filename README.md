<h1 align="center">AeternumCI</h1>

<div align="center">

![STATUS](https://img.shields.io/badge/status-active-brightgreen?style=for-the-badge)
![LICENSE](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

</div>

---

## 🔎 About <a name = "about"></a>

The objective of this project is to showcase the concepts of SWE and DevOps that I have learned so far. This project features CI/CD workflows and automations, unittesting, and containerization. Git & Github were used as the VCS, and a Docker image is published for container deployment.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before running the AeternumCI API, make sure you have the following prerequisites installed:

- Go 1.21 or higher

### Installing

To get started with this project, clone the repository to your local machine and install the required dependencies.

```bash
git clone https://github.com/jgfranco17/aeternum-ci.git
cd aeternum-ci
go mod tidy
```

## 🚀 Usage <a name = "usage"></a>

### Dev mode

To run the API server in dev mode, simply execute either of the following commands:

```bash
# Default Go execution
go run service/cmd/main.go --port=8080 --dev=true

# Or, after editing the Makefile to set a port number
just run-local 8080
```

### Build with Docker

To run the microservice in a container, the package comes with both a Dockerfile and a Compose YAML configuration. Run either of the following to get the API launched in a container; by default, the API will be set to listen on port 5050 for the Compose.

```bash
# Plain Docker build
docker build -t aeternum-ci .
docker run -p 8080:8080 aeternum-ci

# Docker-Compose build
docker compose up
```

## 🔧 Testing <a name = "testing"></a>

### Running unittest suite

To run the full test suite, run the Justfile command as follows:

```bash
just test
```

## ✒️ Authors <a name = "authors"></a>

- [Chino Franco](https://github.com/jgfranco17)
