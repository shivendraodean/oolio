## Database

Import coupon data into Postgres DB.

1. Extract coupon data compressed `.gz` files into the db/migrations directory
2. Run `import_coupon_codes.sh` script
3. Run `index_coupon_codes.sh` script
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