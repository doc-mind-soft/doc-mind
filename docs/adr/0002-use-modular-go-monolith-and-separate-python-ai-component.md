# ADR-002: Use a modular Go monolith and a separate Python AI component

- Status: Accepted
- Date: 2026-09-02

## Context

DocMind combines two substantially different groups of responsibilities.

The first group includes the public HTTP API, document lifecycle, metadata management, validation, transactional workflows, indexing job creation, transactional outbox, and operational endpoints.

The second group includes document parsing, text normalization, chunking, token counting, embedding generation, vector retrieval, context construction, and interaction with large language models.

Splitting every responsibility into an independent service would introduce deployment, networking, observability, contract-versioning, and failure-handling overhead. That complexity is not justified for a project maintained by one developer with approximately 10–20 hours available per week.

At the same time, implementing all AI-related functionality in Go would make it harder to use the established Python ecosystem for document processing, tokenization, evaluation, embeddings, and language-model integrations.

The architecture therefore needs to minimize distributed-system boundaries while retaining a clear boundary around AI-specific processing.

## Decision

DocMind will use two primary backend components:

1. a modular monolith implemented in Go;
2. a separate AI component implemented in Python.

### Go backend

The Go backend will be one deployable application organized into explicit internal modules.

It will own:

- the public HTTP API;
- request validation and the public error model;
- document metadata and document lifecycle use cases;
- file upload and download orchestration;
- creation of documents, indexing jobs, and outbox events in PostgreSQL transactions;
- the outbox publisher;
- public indexing-status and search endpoints;
- health, readiness, configuration, logging, and graceful shutdown.

For the MVP, the outbox publisher will run as a managed background goroutine inside the Go API process.

Modules inside the Go application will communicate through in-process interfaces and application contracts. They will not be separated by network calls merely to imitate microservices.

### Python AI component

The Python AI component will run separately from the Go process and will retain its own dependencies, tests, configuration, and entry points.

It will own:

- consumption of indexing commands from Redis Streams;
- document parsing and normalization;
- chunking and source locator construction;
- model-specific token counting;
- embedding generation;
- persistence of chunks and embeddings;
- indexing warnings, chunk counts, and execution outcomes;
- query embedding and vector retrieval;
- context construction and interaction with the configured language model.

The Python worker must be idempotent. Repeated delivery of the same indexing job must not create duplicate chunks or produce an invalid job state.

### Component interaction

PostgreSQL is the source of truth for documents, indexing jobs, outbox events, chunks, embeddings, and execution results.

Redis Streams is a delivery mechanism for asynchronous indexing commands. It is not the source of truth for job state.

The Go API and Python worker will access original files through the same storage abstraction. In the initial Docker Compose deployment, both containers will mount the same named volume, with the worker receiving read-only access to original files.

The database schema, Redis message format, storage key format, and future synchronous interaction required by search are cross-component contracts. Changes to these contracts must be explicit, versioned when necessary, and tested across the boundary.

The exact protocol for synchronous interaction between the Go backend and the Python AI component will be decided before implementation of semantic search.

The components must not import each other's source code. Sharing one Git repository does not remove their runtime or dependency boundaries.

## Alternatives considered

### Implement all backend and AI functionality in Go

A single Go process would reduce the number of runtimes and deployment units.

This alternative was rejected because document parsing, tokenization, embedding, evaluation, and language-model tooling are better supported by the Python ecosystem. Reimplementing or wrapping these capabilities in Go would add work without improving the MVP.

### Implement the entire backend in Python

A single Python application would provide direct access to AI libraries and remove the cross-language boundary.

This alternative was rejected because Go is the selected language for the transactional application core, public API, concurrency control, and operational backend development. Moving the whole system to Python would discard that deliberate project goal.

### Split backend capabilities into independent microservices

Documents, indexing, search, storage, and other capabilities could each be implemented as separately deployed services.

This alternative was rejected because it would introduce multiple network boundaries, distributed tracing, independent deployments, service discovery, and additional failure modes before the product has demonstrated a need for them.

## Consequences

### Positive

- Transactional application workflows remain inside one Go application.
- Internal Go modules can evolve without network or distributed-transaction overhead.
- AI processing can use mature Python libraries and model tooling.
- Go and Python dependencies remain isolated.
- Asynchronous indexing can be scaled separately from the public API when required.
- Component boundaries are explicit without creating a service for every capability.

### Negative

- Local development and deployment require both Go and Python runtimes.
- Cross-component contracts require integration testing.
- Both components require coordinated database schema changes.
- Asynchronous processing introduces retries, duplicate delivery, and eventual consistency.
- Operational diagnostics must correlate activity across two processes.
- Some end-to-end changes will require coordinated modifications in both languages.

### Reconsider when

This decision should be revisited if the AI component becomes small enough that a separate runtime no longer provides meaningful value, or if individual capabilities require independent teams, scaling, security boundaries, release cycles, or operational ownership.
