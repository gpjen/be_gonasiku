# GONASI Backend

GONASI Backend is the server-side application for the GONASI donation platform. It is built with Go Fiber, GORM, JWT, and MySQL to provide a robust and secure backend infrastructure.

## Features

- **Go Fiber**: GONASI utilizes [Go Fiber](https://github.com/gofiber/fiber) as the web framework to build a fast and efficient backend server.

- **GORM**: Database interactions are handled seamlessly using [GORM](https://gorm.io/), an excellent Object Relational Mapper (ORM) library for Go.

- **JWT Authentication**: Secure user authentication is implemented using JSON Web Tokens (JWT) to ensure the privacy and integrity of user data.

- **MySQL Database**: GONASI stores data in a MySQL database. Make sure to set up your MySQL database and update the configuration accordingly.

## Prerequisites

Before running the GONASI Backend, make sure you have the following installed:

- [Go](https://golang.org/)
- [MySQL](https://www.mysql.com/)

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/GONASI-Backend.git
    ```

2. Install dependencies:

    ```bash
    cd GONASI-Backend
    go get -v ./...
    ```

3. Configure the database:

    - Create a MySQL database for GONASI.
    - Update the database configuration in `config/config.go`.

4. Run the application:

    ```bash
    go run main.go
    ```

The backend server will be running at `http://localhost:3000`.

## API Endpoints

- Document your API endpoints here, providing details on the available routes and their functionalities.

## Configuration

Update the configuration in `config/config.go` to match your environment and preferences.

## Contribution Guidelines

Contributions to GONASI Backend are welcome! Please check out our [Contribution Guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the [MIT License](LICENSE).
