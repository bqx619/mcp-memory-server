# MCP Memory Server

## Overview

MCP Memory Server is an open-source project currently in development that provides a memory storage, retrieval, and management system using Golang and Upstash's vector database. It offers a simple HTTP API for AI assistants to store and retrieve contextual information.

The project implements the MCP (Model Control Protocol) specification, allowing AI models to interact with external tools in a standardized way.

## Features

Currently, the project is in development with three core API methods implemented:

1. **Store Memory** - Save information with automatic vector embeddings
2. **Retrieve Memory** - Semantic search for relevant information using natural language queries
3. **Delete Memory** - Remove specific memories by their ID

## Technology Stack

- **Backend**: Go (Golang)
- **Vector Database**: currently supports Upstash Vector
- **Transport**: HTTP API

## Getting Started

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/bqx619/mcp-memory-server.git
   cd mcp-memory-server
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

3. Create a `.env` file with the following variables:
   ```
   VECTOR_PROVIDER=upstash
   VECTOR_URL=your-upstash-vector-url
   VECTOR_TOKEN=your-upstash-vector-token
   HTTP_PORT=8080
   ```

### Running the Server

```bash
go run .
```

The server will start on the configured port (default: 8080).

## API Reference

### Store Memory

Stores new information in the vector database.

**Endpoint**: POST /tools/store_memory
**Request Body**:
```json
{
  "content": "Memory content to store"
}
```

**Response**:
```json
{
  "content": "Successfully stored memory with ID: abcd1234"
}
```

### Retrieve Memory

Performs semantic search to find relevant information based on a query.

**Endpoint**: POST /tools/retrieve_memory
**Request Body**:
```json
{
  "query": "search term",
  "n_results": 5
}
```

**Response**:
```json
{
  "content": "Found the following memories:
Memory 1:
Content: The first relevant memory
Id: abcd1234
Relevance Score: 0.92
Memory 2:
Content: Another related memory
Id: efgh5678
Relevance Score: 0.85
..."
}
```

### Delete Memory

Removes a specific memory by its ID.

**Endpoint**: POST /tools/delete_memory
**Request Body**:
```json
{
  "memory_id": "abcd1234"
}
```

**Response**:
```json
{
  "content": "Successfully deleted memory with ID: abcd1234"
}
```

## Development Roadmap

- Adding authentication
- Implementing memory tagging and categorization
- Supporting additional vector database providers
- Enhancing search capabilities
- Adding memory expiration and TTL features
- ...

## Contributing

As this project is currently in development, contributions are welcome. Please feel free to submit issues and pull requests.

