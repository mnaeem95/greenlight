# Greenlight - A JSON API for Movies

Greenlight is a JSON API application built using Golang. It allows users to retrieve and manage information about movies through various endpoints. The project layout follows the book "Let's Go Further" and is structured as follows:

## Project Layout

- **cmd/api**: This directory contains the application-specific code for our Greenlight API. It includes the code for running the server, handling HTTP requests, and managing authentication.

- **internal**: The `internal` directory houses various ancillary packages used by our API. It contains code for interacting with the database, data validation, email sending, and more. Any code that isn't application-specific and can potentially be reused resides in this directory. The Go code under `cmd/api` will import the packages from the `internal` directory, but not the other way around.

- **migrations**: The `migrations` directory holds the SQL migration files for our database. It helps manage the evolution of the database schema over time.

- **go.mod**: The `go.mod` file declares our project dependencies, versions, and module path.

- **Makefile**: The `Makefile` contains recipes for automating common administrative tasks, such as auditing Go code, building binaries, and executing database migrations.

## Features Implemented

Greenlight incorporates various features to enhance its functionality. Some of the key features that have been implemented include:

- Custom JSON Encoding: Implementation of the `MarshalJSON` interface for custom JSON encoding.

- Custom JSON Decoding: Implementation of the `UnmarshalJSON` interface for custom JSON decoding.

- Validation Wrapper: A tiny validation wrapper to validate incoming JSON requests for data integrity.

- CRUD Operations: Support for Create, Read, Update, and Delete operations, allowing users to interact with movie data effectively.

- Partial Updates Handling: Support for partial updates, enabling users to update specific fields of a movie.

- Optimistic Locking: Implementation of optimistic locking to prevent conflicts when updating movie data concurrently.

- Filtering, Sorting & Pagination: Features to filter, sort, and paginate movie data to efficiently retrieve specific subsets.

- Structured JSON Log Entries: A lightweight wrapper for generating structured JSON log entries.

- Panic Recovery Middleware: Middleware to handle panic recovery for all endpoints, enhancing the application's stability.

- Rate Limiting: Utilization of the `x/time/rate` package for rate limiting to control the number of requests per second.

- Graceful Shutdown: Implementation of a graceful shutdown mechanism for the server, ensuring clean termination.

- User Registration & Activation: Support for user registration and activation using tokens for secure authentication.

- Permission-Based Authorization: Generic middleware for permission-based authorization to control user access.
