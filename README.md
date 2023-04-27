# Simple CRM Go Fiber

![GitHub last commit](https://img.shields.io/github/last-commit/andrevks/simple-crm-go-fiber?style=flat-square)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/andrevks/simple-crm-go-fiber?style=flat-square)
![GitHub](https://img.shields.io/github/license/andrevks/simple-crm-go-fiber?style=flat-square)

This is a simple CRM (Customer Relationship Management) application written in Go using the Fiber web framework and GORM for database management.

## Requirements

* Go 1.19 or higher

## Getting Started

1. Clone this repository:

```sh
  git clone https://github.com/andrevks/simple-crm-go-fiber.git
```

2. Install dependencies:
```sh
go mod tidy
```
3. Run the application:

```sh
go run main.go
```

4. Access the application in your browser at `http://localhost:3000/api/v1/users`

## Endpoints

* `GET /api/v1/users` - Get all users
* `GET /api/v1/users/:id` - Get a user by ID
* `POST /api/v1/users` - Create a new user
* `DELETE /api/v1/users/:id` - Delete a user by ID

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for more information.

## License

This project is licensed under the [MIT License](LICENSE).

## Credits 

[Akhil Sharma who inspired this project](https://www.youtube.com/@AkhilSharmaTech)