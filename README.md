# Go ERP NLM Mongo

A simple ERP backend built with Go, Fiber, and MongoDB, featuring integration with OpenAI and Ollama for natural language queries.

## Features

- User, Inventory, and Sales management via REST API
- MongoDB for persistent storage
- Natural language query support using OpenAI and Ollama
- Written in Go using the Fiber web framework

## Project Structure

```
.
├── main.go
├── config/
│   └── db.go
├── handlers/
│   ├── inventory.go
│   ├── ollamai.go
│   ├── openai.go
│   ├── sale.go
│   └── user.go
├── models/
│   ├── inventory.go
│   ├── ollama.go
│   ├── sale.go
│   └── user.go
├── routes/
│   └── routes.go
├── .env
├── go.mod
├── go.sum
└── LICENSE
```

## Getting Started

### Prerequisites

- Go 1.24+
- MongoDB instance running
- [OpenAI API key](https://platform.openai.com/)
- [Ollama](https://ollama.com/) running locally or accessible

### Setup

1. **Clone the repository**

   ```sh
   git clone <your-repo-url>
   cd go-erp-nlm-mongo
   ```

2. **Configure environment variables**

   Copy `.env` and set the following:

   ```
   OPENAI_API_KEY=your-openai-key
   MONGO_URI=mongodb://admin:secret123@localhost:27017
   OLLAMA_URL=localhost:11434
   ```

3. **Install dependencies**

   ```sh
   go mod tidy
   ```

4. **Run the application**

   ```sh
   go run main.go
   ```

   The server will start on `http://localhost:3000`.

## API Endpoints

### Users

- `POST /users` - Create a user
- `GET /users` - List all users

### Inventory

- `POST /inventory` - Add inventory item
- `GET /inventory` - List inventory

### Sales

- `POST /sales` - Record a sale
- `GET /sales` - List sales

### Natural Language Queries

- `POST /openai` - Query using OpenAI (expects `{ "message": "..." }`)
- `POST /ollamai` - Query using Ollama (expects `{ "message": "..." }`)

## License

This project is licensed under the [GNU GPL v3](LICENSE).

---

**Note:** For production, ensure your `.env` is secured and not committed to version control.