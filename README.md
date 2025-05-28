# SnippetShare

A modern web application for easily sharing and managing code snippets with others. Built with Go, this application allows developers to create, store, and share code snippets with syntax highlighting and easy-to-use interface.

## Features

- Create and share code snippets with syntax highlighting
- Public and private snippet support
- Easy-to-use RESTful API
- Secure authentication
- Search and filter snippets
- Copy snippets with one click

## Prerequisites

- Go 1.21 or higher
- MySQL 8.0 or higher
- Git

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/yourusername/snippetshare.git
cd snippetshare
```

2. Install dependencies:

```bash
go mod download
```

3. Set up MySQL database:

```bash
# Login to MySQL as root
mysql -u root -p

# Create database and user (in MySQL shell)
CREATE DATABASE snippetshare;
CREATE USER 'snippetshare'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON snippetshare.* TO 'snippetshare'@'localhost';
FLUSH PRIVILEGES;
EXIT;
```

<!-- 4. Set up environment variables:

```bash
# cp .env.example .env
# Edit .env file with your MySQL configuration:
# DB_HOST=localhost
# DB_PORT=3306
# DB_NAME=snippetshare
# DB_USER=snippetshare
# DB_PASSWORD=your_password
``` -->

<!-- 5. Run database migrations:

```bash
go run cmd/migrate/main.go
``` -->

5. Run the application:

```bash
go run cmd/server/main.go
```

The application will be available at `https://localhost:4000`

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
