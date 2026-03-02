# Commit Message Convention

This project follows a structured commit message format to maintain a clean and readable git history.

## Format

```
<type>: <subject>

[optional body]

[optional footer]
```

## Types

- `feat`: A new feature or implementation
- `fix`: A bug fix
- `docs`: Documentation only changes
- `refactor`: Code change that neither fixes a bug nor adds a feature
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

## Subject Line Rules

1. Use imperative mood ("add" not "added" or "adds")
2. Don't capitalize the first letter
3. No period (.) at the end
4. Keep it under 50 characters
5. Be concise but descriptive

## Body (Optional)

- Use when the subject line isn't enough to explain the change
- Wrap at 72 characters
- Explain the "what" and "why", not the "how"
- Separate from subject with a blank line

## Footer (Optional)

- Reference issues: `Closes #123` or `Fixes #456`
- Note breaking changes: `BREAKING CHANGE: description`

## Examples

### Simple feature addition
```
feat: add virtual nodes to consistent hashing
```

### Bug fix with body
```
fix: correct slot calculation in resolver

The slot calculation was off by one when virtualizations
factor exceeded the number of targets. Updated the math
to properly distribute virtual nodes across the slot space.
```

### Refactoring with issue reference
```
refactor: simplify rate limiter middleware chain

Extracted common middleware logic into a reusable function
to reduce code duplication and improve maintainability.

Closes #45
```

### Documentation update
```
docs: add usage examples for rate limiter configuration
```

### Breaking change
```
feat: change consistent hashing API to use context

Updated all Resolve methods to accept context.Context as
the first parameter for better cancellation support.

BREAKING CHANGE: Resolve() now requires a context parameter
```

## Tips

- Make atomic commits (one logical change per commit)
- Commit early and often during development
- Write clear, descriptive messages
- Think about future developers reading your commit history
- Use the body to provide context that code alone cannot

## Running Pre-commit Checks

Before committing, ensure:
```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Run tests
go test ./...
```
