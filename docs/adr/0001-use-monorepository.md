# ADR-001: Use a monorepository

- Status: Accepted
- Date: 2026-09-02

## Context

DocMind will consist of several technology-specific components: a React frontend, a modular Go backend, a separate Python AI component, shared API contracts, database migrations, architecture documentation, and deployment configuration.

The project is developed and maintained by one person with approximately 10–20 hours available per week. Maintaining several repositories would introduce additional coordination, versioning, CI, access-management, and release overhead.

Some changes will affect more than one component. For example, an OpenAPI contract change may require coordinated updates to the Go backend and the generated frontend client. Such changes should be reviewable and mergeable as one consistent unit when necessary.

The repository boundary must not be confused with a runtime boundary. Components stored in one repository may still run as separate processes or containers.

## Decision

DocMind will use a single Git monorepository for application components, shared contracts, documentation, and deployment configuration.

The target top-level areas are:

- `frontend/` for the React and TypeScript application;
- `backend/` for the modular Go application;
- `ai/` for the Python AI component;
- `api/` for shared API contracts;
- `docs/` for ADRs and other architecture documentation;
- `deploy/` for deployment configuration.

These directories will be created only when their corresponding implementation capability begins. Empty placeholder directories will not be added in advance.

Each component will retain its own dependency manifest, build process, tests, and internal boundaries. A shared repository does not imply that all components form one runtime process or may depend on each other without explicit contracts.

## Alternatives considered

### Separate repository for every component

The frontend, Go backend, Python AI component, and deployment configuration could each have a dedicated repository.

This would provide stronger repository-level isolation and independent access control. It was rejected for the current project because it would add coordination and maintenance overhead without providing a practical benefit to a single-developer team.

### Separate frontend repository

The frontend could be separated while the backend, AI component, and infrastructure remain together.

This would allow an independent frontend lifecycle, but it would make coordinated API contract changes more difficult. The project currently has no separate frontend team or release process that would justify this split.

## Consequences

### Positive

- Cross-component changes can be reviewed atomically when they must remain consistent.
- Architecture documentation and shared contracts remain close to the implementation.
- Local development and repository discovery are simpler.
- Git, CI, issue tracking, and access management require less administration.
- One developer can maintain the full system without synchronizing multiple repositories.

### Negative

- Repository size and CI duration will grow as components are added.
- CI will eventually require path-based filtering to avoid unnecessary work.
- Component boundaries must be enforced through structure and contracts rather than repository separation.
- Independent access control and release histories are harder to implement.

### Reconsider when

This decision should be revisited if components gain independent teams, release cycles, access requirements, or operational ownership, or if repository size and CI performance become measurable problems.
