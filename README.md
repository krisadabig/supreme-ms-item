# Supreme MS Item Service

A RESTful API service for managing items with user association, built with Go and Supabase.

## Features

- Create, read, update, and delete items
- User-specific item management
- Input validation
- Clean architecture with separation of concerns
- Environment-based configuration
- Supabase integration for data persistence

## Prerequisites

- Go 1.21 or higher
- Supabase account and project
- Git

## Getting Started

1. **Clone the repository**
   ```bash
   git clone git@github.com:krisadabig/supreme-ms-item.git
   cd supreme-ms-item
   ```

2. **Set up environment variables**
   Copy the example environment file and update it with your Supabase credentials:
   ```bash
   cp .env.example .env
   ```
   Then edit the `.env` file with your Supabase URL and anon key.

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run the application**
   ```bash
   go run cmd/main.go
   ```
   The server will start on `http://localhost:1323` by default.

## API Endpoints

- `GET /` - Health check endpoint
- `GET /items` - List all items for the authenticated user
- `GET /items/:id` - Get a specific item by ID
- `POST /items` - Create a new item
- `PUT /items/:id` - Update an existing item
- `DELETE /items/:id` - Delete an item

## Project Structure

```
.
├── adapters/           # Adapters for external services
│   ├── handlers/       # HTTP handlers
│   └── repositories/   # Database repositories
├── application/        # Application services
├── cmd/                # Application entry points
│   ├── http_server/    # HTTP server implementation
│   └── main.go         # Main application entry point
├── domain/             # Domain models and interfaces
├── internal/           # Internal application code
│   ├── bootstrap/      # Application bootstrap and initialization
│   ├── config/         # Configuration management
│   └── router/         # HTTP router setup
├── .env                # Environment variables (gitignored)
├── config.yaml         # Application configuration
├── go.mod              # Go module definition
└── go.sum              # Go module checksums
```

## Configuration

The application can be configured using `config.yaml` and environment variables:

```yaml
server:
  port: ":1323"  # HTTP server port

supabase:
  url: "${SUPABASE_URL}"      # Supabase project URL
  anon_key: "${SUPABASE_KEY}" # Supabase anon/public key
```

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```
SUPABASE_URL=your-supabase-url
SUPABASE_KEY=your-supabase-anon-key
```

## Development

### Running Tests

```bash
go test ./...
```

### Building the Application

```bash
go build -o bin/item-service cmd/main.go
```

## Deployment

The application can be containerized using the provided Dockerfile:

```bash
docker build -t item-service .
docker run -p 1323:1323 --env-file .env item-service
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
