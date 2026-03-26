# Employee API

A sample employee REST API service built with Go.

## Getting Started

```bash
go mod download
go run .
```

## API Endpoints

- `GET /employees` - List all employees
- `GET /employees/{id}` - Get employee by ID
- `POST /employees` - Create a new employee

## Example

```bash
curl -X POST http://localhost:8080/employees \
  -H "Content-Type: application/json" \
  -d '{"first_name": "Jane", "last_name": "Doe", "email": "jane@example.com", "role": "Engineer"}'
```
