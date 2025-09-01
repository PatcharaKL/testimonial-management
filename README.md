# Testimonial Management API

A Go Fiber REST API for managing testimonials with PostgreSQL.

## Features
- Create and fetch testimonials
- Clean architecture: handler, usecase, repository, DTO, entity
- Unit tests for repository and usecase layers
- PostgreSQL migrations

## Project Structure
```
app/main.go                # App entrypoint, DI, server setup
internal/entities/         # Domain models
internal/dto/              # Request/response DTOs
internal/repository/       # DB access logic
internal/usecases/         # Business logic
internal/handler/          # HTTP route handlers
pkg/                       # Shared utilities (DB connection)
migrations/                # DB schema migrations
```

## Getting Started
1. **Clone the repo**
2. **Configure `.env`**
   ```
   DB_USER=youruser
   DB_PASS=yourpass
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=yourdb
   ```
3. **Run migrations**
   ```bash
   # Use your preferred migration tool or psql
   psql -U $DB_USER -d $DB_NAME -f migrations/000001_init-db.up.sql
   ```
4. **Start the server**
   ```bash
   go run app/main.go
   ```
5. **API Endpoints**
   - `POST /testimonial` — Create a testimonial
   - `GET /testimonial` — List all testimonials

## Testing
Run unit tests:
```bash
go test ./internal/repository/...
go test ./internal/usecases/...
```

## Example Request
**POST /testimonial**
```json
{
  "full_name": "John Doe",
  "email": "john@example.com",
  "role": "Developer",
  "company": "Acme Corp",
  "testimonial": "Great service!",
  "photo_url": "http://example.com/photo.jpg",
  "is_public": true
}
```

## License
MIT
