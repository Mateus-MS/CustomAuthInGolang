# Use an official Go image for the backend
FROM golang:latest AS backend-builder

# Set the working directory in the container
WORKDIR /app

# Copy the entire backend folder to the working directory
COPY /backend .

# Build the Go application
RUN go build -o main .

# Use the official PostgreSQL image for the database
FROM postgres:latest

# Copy the initialization scripts and .env file into the container
COPY init.sql /docker-entrypoint-initdb.d/
COPY .env /docker-entrypoint-initdb.d/env

# Copy the built Go application from the builder stage
COPY --from=backend-builder /app/main /app/main

# Set environment variables for PostgreSQL from .env file
ENV POSTGRES_USER=$DBuser \
    POSTGRES_PASSWORD=$DBpass \
    POSTGRES_DB=users

# Expose PostgreSQL and application ports
EXPOSE 5432 8080

# Command to start both PostgreSQL and the Go application
CMD ["sh", "-c", "docker-entrypoint.sh postgres & /app/main"]