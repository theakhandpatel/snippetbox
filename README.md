# SnippetBox

SnippetBox is a web application that allows users to create and share text snippets with an expiration time. The service uses sessions for authentication and provides various endpoints for different functionalities.

## Features

- Create, view, and share text snippets.
- Set an expiration time for each snippet.
- User registration and authentication.
- User account settings and password update.
- Secure session management and anti-CSRF protection.
- Static file serving for CSS and JavaScript.

## Endpoints

### Non-Authenticated Routes

- `GET /` - Home page displaying the list of available snippets.
- `GET /about` - About page providing information about the application.
- `GET /ping` - A simple endpoint for checking if the server is running.
- `GET /user/signup` - User registration page.
- `POST /user/signup` - User registration form submission.
- `GET /user/login` - User login page.
- `POST /user/login` - User login form submission.
- `GET /snippet/view/:id` - View a specific snippet.

### Authenticated Routes

To access these routes, users must be logged in.

- `GET /snippet/create` - Create a new snippet.
- `POST /snippet/create` - Create a new snippet form submission.
- `POST /user/logout` - Log out the user.
- `GET /user/me` - User account information.
- `GET /user/password/update` - Change user's password page.
- `POST /user/password/update` - Change user's password form submission.

## Installation and Usage

1. Clone the repository: `git clone https://github.com/theakhandpatel/snippetbox.git`
2. Navigate to the project directory: `cd snippetbox`
3. Install dependencies: `go mod tidy`https://github.com/theakhandpatel/snippetbox.git
4. Set up your MySQL database and update the `dsn` value in `main.go`.
5. Generate TLS certificates or update the paths in `main.go` for development use.
6. Run the application: `go run main.go`
   - You can customize the behavior using the following command-line flags:
     - `-addr`: Set the HTTP network address (default: ":4000").
     - `-dsn`: Set the database connection string (default: "web:password@/snippetbox?parseTime=true").
     - `-debug`: Enable debug mode (default: false).

The application will be available at `https://localhost:4000`.

## Running Tests

1. Ensure that the application is running.
2. Run tests: `go test ./...`

## License

This project is licensed under the [MIT License](LICENSE).