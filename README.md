# REAL FULLSTACK TODO TEMPLATE

### Technology Architecture Stack

1. [Docker](https://www.docker.com/)
2. [PostgreSQL](https://www.postgresql.org/)
3. [Golang](https://go.dev/)

- [Sqlc](https://sqlc.dev/): Compile SQL to type-safe code;
  catch failures before they happen.
- [Goose](https://github.com/pressly/goose): Goose is a database migration tool. Manage your database schema by creating incremental SQL changes or Go functions.

4. [Nextjs](https://nextjs.org/)

- [Tailwind Css](https://tailwindcss.com/)
- [Daisyui](https://daisyui.com/): The most popular component library for Tailwind CSS

5. [Bun](https://bun.sh/)

### Develop the Application

#### DB

Step 1: change the compose.yaml file Expose the database Port `5432`.

```yaml
ports:
  - 5432:5432
```

Step 2 : Install and start your database.

```bash
docker compose up --build -d db
```

#### Frontend

```bash
cd frontend
bun install
bun run dev
```

#### Backend

```bash
cd backend
go run main.go
```

### Run the Application

```bash
docker compose up --build -d
```

#### Backend API

```bash
http GET  "http:localhost:8000/todos"
```

#### Frontend Website

Use Your favorite browser and Open the web page [TODO]("http://localhost:3000").

### Shut down the Application

```bash
docker compose down
```
