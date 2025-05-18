# Oolio ECommerce API

## Database

Import coupon data into Postgres DB.

1. Extract coupon data compressed `.gz` files into the db/migrations directory
2. `cd` into `db/migrations` 
3. Run `import_coupon_codes.sh` script
4. Run `index_coupon_codes.sh` script
5. Run `create_tables.sh` scripts
Run `chmod +x <script_file_name>` if you get permission errors. 

This migration script assumes you are running your Postgres DB in a Docker container.

## Web Service

### Prerequisites

- Go
- Docker

### Development

- Run `docker-compose up --build`
- Application running at `http://localhost:8080`
- Alternatively run directly with `go run cmd/main.go`

### Testing

- Run `go test ./src/...`

### Code Structure

```
src/
├── cmd/                 
│   └── main.go          # Main application file
├── internal/            
│   ├── model/           # Primary domain models
│   ├── repository/      # Data access
│   ├── service/         # Services for logic and business operations
│   └── webapp/          # Web application/API
│       ├── handler/     # HTTP request handlers/controllers
│       ├── container.go # Dependency injection container
│       └── router.go    # HTTP routes
```

## NOTES

Out of scope items
- No AuthR/N
- No access control
- No user context
- No paging, filtering, sorting of data
- No caching
- No versioning
- No hosting, build/deployment pipeline, PAC, or IAC
- Hardcoded string in some places
- Secrets are hardcoded in env files and deliberately pushed to source control
- Git commits are not grouped logically, and treated like a playground