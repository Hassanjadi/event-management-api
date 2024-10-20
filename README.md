# Event Management System API

This system allows users to create, manage, and participate in events. Users can register for specific events, provide reviews, and view event details. The system also manages event categories.

## Features

- CRUD for users
- CRUD for categories
- CRUD for events
- CRUD for reviews
- CRUD for participants

## Technologies Used

- Golang
- Gin Framework
- PostgreSQL
- dotenv for configuration settings
- SQL Migrate for database migration management
- Basic Auth

## Prerequisites

Before you begin, make sure you have the following installed:

- [Golang](https://golang.org/dl/) the latest version
- [PostgreSQL](https://www.postgresql.org/download/) according to your operating system

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/Hassanjadi/event-management-api.git
   ```

2. **Navigate to the project directory**

   ```bash
   cd event-management-api
   ```

3. **Install dependencies**

   ```bash
   go mod tidy
   ```

4. **Copy the configuration file**

   Copy the env example file to .env and update the environment settings as needed.

   ```bash
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database_name

   ```

5. **Run database migrations**

   ```bash
   go run main.go
   ```

## Entity Relationship Diagram

![ERD](https://i.ibb.co.com/bKxdhc6/emsy.png)

## API Endpoints

### Users

- GET `/api/users` Retrieve all users
- POST `/api/users` Add a new user
- GET `/api/users/:id` Retrieve user details by ID
- PUT `/api/users/:id` Update user by ID
- DELETE `/api/users/:id` Delete user by ID

### Categories

- GET `/api/categories` Retrieve all categories
- POST `/api/categories` Add a new category
- GET `/api/categories/:id` Retrieve category details by ID
- PUT `/api/categories/:id` Update category by ID
- DELETE `/api/categories/:id` Delete category by ID

### Events

- GET `/api/events` Retrieve all events
- POST `/api/events` Add a new event
- GET `/api/events/:id` Retrieve event details by ID
- PUT `/api/events/:id` Update event by ID
- DELETE `/api/events/:id` Delete event by ID

### Reviews

- GET `/api/reviews` Retrieve all reviews
- POST `/api/reviews` Add a new review
- GET `/api/reviews/:id` Retrieve review details by ID
- PUT `/api/reviews/:id` Update review by ID
- DELETE `/api/reviews/:id` Delete review by ID

### Participants

- GET `/api/participants` Retrieve all participants
- POST `/api/participants` Add a new participant
- GET `/api/participants/:id` Retrieve participant details by ID
- PUT `/api/participants/:id` Update participant by ID
- DELETE `/api/participants/:id` Delete participant by ID

## Authentication

This API uses Basic Auth for authentication. Make sure to use a valid username and password (admin / admin as an example).

The application can be accessed at `http://localhost:8080/`.
