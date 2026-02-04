# Contributing

## Setup

```bash
git clone <repository-url>
cd math-skills
make build
make test
```

## Commands

```bash
make build    # Compile
make run      # Build and run
make test     # Run tests
make clean    # Remove binaries
make fmt      # Format code
make vet      # Check quality
```

## Guidelines

- Follow Go conventions
- Run `make fmt` and `make vet` before committing
- Add tests for new features
- Keep functions small and focused
