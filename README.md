# Hospital Management System API

A Go-based backend API for a hospital management system that allows doctors and receptionists to manage patient records.

## Features

- **Role-Based Authentication**: Separate login for doctors and receptionists with JWT token-based authentication
- **Patient Management**: Create, read, update, and delete patient records
- **Role-Based Access Control**: Different permissions for doctors and receptionists
  - Receptionists: Can create, read, update, and delete patient records
  - Doctors: Can read and update patient records

## Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Containerization**: Docker

## Project Structure

```
├── config/             # Configuration and database setup
├── handlers/           # HTTP request handlers
├── middleware/         # Authentication and CORS middleware
├── models/             # Data models
├── router/             # API routes definition
├── server/             # Application entry point
├── services/           # Business logic
├── DockerFile          # Docker configuration
└── docker-compose.yaml # Docker Compose configuration
```

## API Endpoints

### Authentication

- **POST /login** - Authenticate user (doctor or receptionist)

### Patient Management

#### Receptionist Access

- **GET /patients** - List all patients
- **GET /patients/:id** - Get patient by ID
- **POST /patients** - Create a new patient
- **PUT /patients/:id** - Update patient information
- **DELETE /patients/:id** - Delete a patient

#### Doctor Access

- **GET /patients** - List all patients
- **GET /patients/:id** - Get patient by ID
- **PUT /patients/:id** - Update patient information

## Setup and Installation

### Prerequisites

- Docker and Docker Compose
- Go 1.24+ (for local development without Docker)
- PostgreSQL (for local development without Docker)

### Using Docker

1. Clone the repository
2. Navigate to the project directory
3. Run the application using Docker Compose:

```bash
docker-compose up -d
```

The API will be available at http://localhost:8080

### Local Development

1. Clone the repository
2. Set up environment variables:

```
DATABASE_URL=postgres://postgres:password123@localhost:5432/patientdb?sslmode=disable
JWT_SECRET=supersecretkey
```

3. Run the application:

```bash
go run server/main.go
```

## Default Users

The system is pre-seeded with the following users:

- **Receptionist**:
  - Username: receptionist1
  - Password: pass123

- **Doctor**:
  - Username: doctor1
  - Password: pass123

## API Usage Examples

### Login as Receptionist

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "receptionist1", "password": "pass123"}'
```

### Create a Patient (Receptionist Only)

```bash
curl -X POST http://localhost:8080/patients \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "John Doe",
    "age": 30,
    "gender": "male",
    "notes": "No known allergies"
  }'
```

### Get Patient by ID (Both Roles)

```bash
curl -X GET http://localhost:8080/patients/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Deployment

This application can be deployed to any cloud provider that supports Docker containers. The project includes a Dockerfile and docker-compose.yaml for easy deployment.

The application is currently deployed at: https://go-backend-crud.onrender.com/