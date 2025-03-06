## RAGO
> I hate Python and JS but there are no great tools for that in Golang. So I made one. That's it. Oh-zzo-rago.

## Overview

### Knowledge Base
* Datasource that contains knowledge, and performs similarity search.
* Normally, it supports vector type, and also called as Vector Search Engine.

### Ingester

#### Concept

The `Ingester` interface is designed to handle the process of loading, modifying, generating embeddings, and storing documents into a Vector Search Engine. This process is essential for applications that require efficient and accurate similarity searches, such as recommendation systems or search engines.

#### Interface

The `Ingester` interface defines a single method, `Ingest`, which is responsible for orchestrating the entire ingestion process. This method should be implemented by any struct that aims to perform the ingestion of documents.

#### Implementation

The `DefaultIngester` struct is an implementation of the `Ingester` interface. It combines several components to perform the ingestion process:

- `DocumentLoader`: Responsible for loading documents from a data source.
- `DocumentModifiers`: Applies modifications to the loaded documents.
- `EmbeddingGenerator`: Generates embeddings for the modified documents.
- `KnowledgeAddable`: Adds the generated embeddings and modified documents to the Vector Search Engine.

The `Ingest` method of `DefaultIngester` orchestrates these components to complete the ingestion process.

#### Components

##### DocumentLoader

The `DocumentLoader` interface defines a method for loading documents from a data source.

```go
type DocumentLoader interface {
	Load() ([]string, error)
}
```

##### DocumentModifiers

The `DocumentModifiers` interface defines a method for applying modifications to a list of documents.

```go
type DocumentModifiers interface {
	ApplyAll([]string) ([]string, error)
}
```

##### EmbeddingGenerator

The `EmbeddingGenerator` interface defines a method for generating embeddings for a list of documents.

```go
type EmbeddingGenerator interface {
	Generate([]string) ([]Embedding, error)
}
```

##### KnowledgeAddable

The `KnowledgeAddable` interface defines a method for adding embeddings and documents to the Vector Search Engine.

```go
type KnowledgeAddable interface {
	Add([]Embedding, []string) error
}
```

#### Summary

The `Ingester` interface and its implementation in `DefaultIngester` provide a structured way to handle the ingestion of documents into a Vector Search Engine. By breaking down the process into distinct components, it ensures that each step is modular and can be easily maintained or replaced. This design promotes clean code and separation of concerns, making the ingestion process more manageable and scalable.

### Retriever

#### Concept

The `Retriever` interface is designed to handle the process of retrieving documents from a Vector Search Engine based on their similarity to a given input. This process is essential for applications that require efficient and accurate similarity searches, such as recommendation systems or search engines.

#### Interface

The `Retriever` interface defines a single method, `Retrieve`, which is responsible for orchestrating the entire retrieval process. This method should be implemented by any struct that aims to perform the retrieval of documents.

#### Implementation

The `DefaultRetriever` struct is an implementation of the `Retriever` interface. It combines several components to perform the retrieval process:

- `EmbeddingGenerator`: Generates embeddings for the input documents.
- `KnowledgeSearchable`: Searches the Vector Search Engine for documents similar to the input embeddings.

The `Retrieve` method of `DefaultRetriever` orchestrates these components to complete the retrieval process.

#### Components

##### EmbeddingGenerator

The `EmbeddingGenerator` interface defines a method for generating embeddings for a list of documents.

```go
type EmbeddingGenerator interface {
    Generate([]string) ([]Embedding, error)
}
```

##### KnowledgeSearchable

The `KnowledgeSearchable` interface defines a method for searching the Vector Search Engine for documents similar to the input embeddings.

```go
type KnowledgeSearchable interface {
    Search(embeddings []Embedding, topK int) (Retrieved, error)
}
```

#### Summary

The `Retriever` interface and its implementation in `DefaultRetriever` provide a structured way to handle the retrieval of documents from a Vector Search Engine. By breaking down the process into distinct components, it ensures that each step is modular and can be easily maintained or replaced. This design promotes clean code and separation of concerns, making the retrieval process more manageable and scalable.


## API Documentation

### Base URL
```
http://localhost:1323
```

### Endpoints
#### Create kNN Index
##### `POST /index/knn/:indexName`
- **Description**: Create a kNN index.
- **Parameters**:
  - `indexName` (path): The name of the index to create.
- **Response**:
  - `200 OK`: `{ "message": "success" }`
  - `500 Internal Server Error`: `{ "message": "error message" }`

#### Ingest Documents
##### `POST /ingest/:indexName`
- **Description**: Ingest documents into the specified index.
- **Parameters**:
  - `indexName` (path): The name of the index to ingest documents into.
- **Request Body**:
  ```json
  {
    "documents": ["document1", "document2", ...]
  }
  ```
- **Response**:
    - `200 OK`: `{ "message": "success" }`
    - `400 Bad Request`: `{ "message": "error message" }`

#### Retrieve Documents
##### `GET /retrieve/:indexName`
- **Description**: Retrieve documents from the specified index.
- **Parameters**:
    - `indexName` (path): The name of the index to retrieve documents from.
- **Request Body**:
  ```json
  {
    "query": "query string",
    "message": "message string"
  }
  ```
- **Response**:
    - `200 OK`: `{ "message": "retrieved documents" }`
    - `400 Bad Request`: `{ "message": "error message" }`
    - `500 Internal Server Error`: `{ "message": "error message" }`

#### Ingest Similar Items
##### `POST /ingest/similar/item/:indexName`
- **Description**: Ingest similar items into the specified index.
- **Parameters**:
    - `indexName` (path): The name of the index to ingest similar items into.
- **Request Body**:
  ```json
  {
    "documents": ["item1", "item2", ...]
  }
  ```
- **Response**:
    - `200 OK`: `{ "message": "success" }`
    - `400 Bad Request`: `{ "message": "error message" }`

#### Retrieve Similar Items
##### `GET /retrieve/similar/item/:indexName`
- **Description**: Retrieve similar items from the specified index.
- **Parameters**:
    - `indexName` (path): The name of the index to retrieve similar items from.
- **Request Body**:
  ```json
  {
    "query": "query string"
  }
  ```
- **Response**:
    - `200 OK`: `{ "message": "retrieved items" }`
    - `400 Bad Request`: `{ "message": "error message" }`
    - `500 Internal Server Error`: `{ "message": "error message" }`

### Common Response Format
- **Success**:
  ```json
  {
    "message": "success"
  }
  ```
- **Error**:
  ```json
  {
    "message": "error message"
  }
  ```
```

