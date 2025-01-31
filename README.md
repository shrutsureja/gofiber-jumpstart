# GoFiber-Jumpstart

## Technologies Used

- [Go](https://golang.org/) - Language
- [Fiber](https://gofiber.io/) - Framwork
- [GORM](https://gorm.io/) - ORM
- [PostgreSQL](https://www.postgresql.org/) - Current Primary Database
- [Goose](https://github.com/pressly/goose) - For maintaing migrations
- [Validator](https://github.com/go-playground/validator) - For validation of data

## Setup the Codebase Locally

Follow these steps to set up the XYZ project on your local machine:

1. **Install Go** - Ensure you have Go installed. You can download it from [go.dev](https://go.dev/dl/) version `1.23.5`

2. **Install the necessary dependencies**:

    ```bash
    go mod tidy
    ```

3. **Create a .env file**: Copy the `example.env` file to create your own `.env` file and configure it with necessary values:

    ```bash
    cp .env.example .env
    ```

5. **Run the application locally**

    ```bash
    go run src/main.go
    ```

The server will start on `localhost:3000` by default. Migrations will run at server startup by default unless explicitly disabled in the `.env` file.

### Setup with Docker

1. **Build the Docker image**:

    ```bash
    docker build -t gofiber-jumpstart .
    ```

2. **Database Setup**:

    - Create a new PostgreSQL database and update the `.env` file with the database credentials.

3. **Run the Docker container**:

    ```bash
    docker run --env-file .env -p 3000:3000 gofiber-jumpstart
    ```

The server will start on `localhost:3000` by default. Migrations will run at server startup by default unless explicitly disabled in the `.env` file.

## Setup with Docker Compose

1. **Database Setup**:

    - Update the `.env` file with the database credentials.

2. **Run the Docker container**:

    ```bash
    docker-compose up
    ```

The server will start on `localhost:3000` by default. Migrations will run at server startup by default unless explicitly disabled in the `.env` file.

## Migration Commands 

This project uses Goose for managing database migrations. Use the following commands to apply and manage migrations:

> Before running migrations command make sure to install goose locally ([link](https://pressly.github.io/goose/installation/))

- **Apply all up migrations**:

    ```bash
    make migrate-up
    ```

- **Roll back the last migration**:

    ```bash
    make migrate-down
    ```

- **Create a new migration**:

    ```bash
    make create name=<migration_name>
    ```

- **Check the current migration status**:

    ```bash
    make status
    ```

- **Roll back all migrations**:

    ```bash
    make migrate-reset
    ```

- **Migrate to a specific version**:

    ```bash
    make migrate-to v=<version_number>
    ```

To get more control of the migrations use the **`goose`** commands ([link](https://github.com/pressly/goose))