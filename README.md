# Project url-shortner

url-shortner features a robust and efficient URL shortening algorithm that generates unique, random alphanumeric codes. The algorithm ensures that each shortened URL is unique and cannot be predicted, providing an additional layer of security.

The application integrates with a PostgreSQL using GORM ORM, templ and tailwind on th UI.

## Getting Started

1. Clone the repository
2. Add a .env file and copy the configuration inside the .env.example into the .env file and modify accordingly
3. Run `docker compose up -d` to  run postgresql containers in detached mode
3. Run the application using `make watch`
4. Open your browser and enter the proxy url http://localhost:3001 to have access to live reload 


NOTE: Always run the application using port 3000. If you change the port number in the .env file, change it in the .air.toml file also

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```


 Generate templ function files by running `templ generate``