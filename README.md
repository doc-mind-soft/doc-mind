# DocMind

DocMind is a long-term full-stack product for uploading, processing, indexing, and semantically searching user documents with retrieval-augmented generation (RAG).

The project also serves as a practical environment for studying modern backend, frontend, data, and AI architecture.

## Status

Architecture version 0.2 is approved. Application implementation has not started.

## MVP scope

The MVP will be a single-user web application that supports:

- uploading one TXT, Markdown, PDF, or DOCX file per request;
- storing original files in a local filesystem;
- indexing documents asynchronously;
- viewing documents and their indexing status;
- downloading and deleting documents;
- performing semantic search over indexed content;
- generating answers grounded in retrieved chunks;
- returning citations with the source document and its format-specific location.

## Out of scope for MVP

- authentication and authorization;
- multitenancy;
- document versioning;
- OCR;
- chat history;
- WebSocket progress and answer streaming;
- hybrid search and reranking;
- external integrations;
- MinIO or another object storage service.

## Target stack

- Frontend: React and TypeScript
- Backend: Go
- AI component: Python
- Data: PostgreSQL with pgvector
- Async messaging: Redis Streams

## Development approach

DocMind is built through small, independently verifiable iterations. Each iteration has a narrow scope, an objective validation method, and its own commit.
