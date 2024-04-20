# Starting the Server

## Starting PostgreSQL Container

To start the PostgreSQL container, run the following command:

```bash
docker-compose -f docker/postgres-container.yml up
```
This command will start the PostgreSQL container required for the server to run.



## Starting Go Server Locally

To start the Go server locally, follow these steps:
- Make sure PostgreSQL container is up and running as instructed above
- Navigate to the root directory of the project.
- Run the following command to start the server:

```bash
go run cmd/main/main.go
```

This command will start the Go server locally, and it will be accessible at http://localhost:3000.

