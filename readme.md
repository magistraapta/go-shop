# Go-Shop Backend



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

# Todo
- [x] complete CRUD operation on product
- [x] add checkout feature
- [x] add transaction model
- [ ] setup RBAC for auth
- [ ] Write test
- [ ] Setup Docker
- [ ] Setup Github Actions
- [ ] Create API Docs using Swagger