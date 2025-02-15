# Restaurant Booking System

This is a restaurant booking system built with Go and PostgreSQL. The system allows users to book and cancel reservations for tables at a restaurant.

## Prerequisites

- Docker
- Docker Compose

## Setup

1. **Clone the repository:**

   ```sh
   git clone https://github.com/AmiraliFarazmand/1DLJ5-73943.git
   cd 1DLJ5-73943
   ```

2. **Create a `.env` file (optional):**

   Create a `.env` file in the project root with the following content *(In the case if you want to run it without docker and do it with $go run)*:

   ```properties
   DSN="host=db user=myuser password=mypassword dbname=restaurant_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
   SECRET_KEY="some_secret_key_kalfjdlkfmkmfklsdmfklsdmfkslmdfklsmfdklsmdfh"
   ```


3. **Build and run the Docker containers:**

   ```sh
   docker-compose up --build -d
   ```

   This will build and start both the PostgreSQL database and the Go application. The `init.sql` file will be executed, creating the `tables` table and inserting 10 rows into it.
   \
   **Note:** If you encounter an issue where the application fails to connect to the database, it may be due to the database not being fully ready when the application starts. You can resolve this by running the `docker-compose up` command again:

   ```sh
   docker-compose up -d
   ```

4. **Access the application:**

   The application will be available at `http://localhost:8080`.

## API Endpoints

### Authentication

- **Signup:** `POST /signup`
- **Login:** `POST /login`
- **Validate:** `GET /validate`

### Reservations

- **Book a Table:** `POST /book`
- **Cancel a Reservation:** `POST /cancel`

