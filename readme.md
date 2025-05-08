# Go-Shop Backend

Go-Shop is a robust e-commerce backend service built with Go, implementing a layered architecture pattern. This project provides a complete set of features for an online shopping platform, including user authentication, product management, shopping cart functionality, and order processing.

## Features
- User authentication and authorization (JWT-based)
- Product management (CRUD operations)
- Shopping cart functionality
- Order processing and checkout
- Transaction management
- Role-based access control (Admin/User)

## Project Structure
While the codebase is organized into separate modules (auth, product, cart, order) for better maintainability and separation of concerns, Go-Shop is implemented as a monolithic application. This means:

- All components are deployed as a single unit
- Modules share the same database and resources
- Direct communication between modules through function calls
- Single codebase and deployment pipeline

This monolithic approach was chosen for:
- Simpler deployment and maintenance
- Easier development and debugging
- Reduced operational complexity
- Better performance for smaller to medium-scale applications

## Architecture
This layered architecture provides several benefits:
- Separation of concerns
- Maintainability and testability
- Clear dependencies between components
- Scalability and flexibility

The project follows clean architecture principles with clear boundaries between layers and dependency injection (e.g., repositories are injected into services, services into handlers).

## Entity Relationship Diagram
![ER Diagram](https://github.com/magistraapta/go-shop/blob/ea26075301cca3536d30420fc05564d7c01186c0/images/ERD.png)

## Directory Structure
```
.
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── auth/             # Authentication related code
│   ├── cart/             # Shopping cart functionality
│   ├── order/            # Order management
│   └── product/          # Product management
├── middleware/           # HTTP middleware components
├── initializers/         # Application initialization code
├── test/                # Test files
├── images/              # Project images and diagrams
├── .github/             # GitHub specific files
├── dockerfile           # Docker configuration
├── docker-compose.yml   # Docker compose configuration
├── prometheus.yml       # Prometheus monitoring config
├── go.mod              # Go module definition
└── go.sum              # Go module checksums
```

## Directory Structure
```
.
├── cmd/                    # Application entry points
├── internal/              # Private application code
│   ├── auth/             # Authentication related code
│   ├── cart/             # Shopping cart functionality
│   ├── order/            # Order management
│   └── product/          # Product management
├── middleware/           # HTTP middleware components
├── initializers/         # Application initialization code
├── test/                # Test files
├── images/              # Project images and diagrams
├── .github/             # GitHub specific files
├── dockerfile           # Docker configuration
├── docker-compose.yml   # Docker compose configuration
├── prometheus.yml       # Prometheus monitoring config
├── go.mod              # Go module definition
└── go.sum              # Go module checksums
```

# Todo
- [x] complete CRUD operation on product
- [x] add checkout feature
- [x] add transaction model
- [ ] setup RBAC for auth
- [x] add checkout feature
- [x] add transaction model
- [ ] setup RBAC for auth
- [ ] Write test
- [ ] Setup Docker
- [ ] Setup Github Actions
- [ ] Create API Docs using Swagger