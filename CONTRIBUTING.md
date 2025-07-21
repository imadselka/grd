# Contributing to grd

Thank you for considering contributing to grd! This document outlines the process for contributing to the project.

## Code of Conduct

By participating in this project, you agree to abide by the [Code of Conduct](CODE_OF_CONDUCT.md).

## How Can I Contribute?

### Reporting Bugs

- **Ensure the bug was not already reported** by searching on GitHub under [Issues](https://github.com/imadselka/grd/issues).
- If you're unable to find an open issue addressing the problem, [open a new one](https://github.com/imadselka/grd/issues/new). Be sure to include a **title and clear description**, as much relevant information as possible, and a **code sample** demonstrating the expected behavior that is not occurring.

### Suggesting Enhancements

- **Check if the enhancement has already been suggested** by searching on GitHub under [Issues](https://github.com/imadselka/grd/issues).
- If it hasn't, [create a new issue](https://github.com/imadselka/grd/issues/new) with a clear title and description of the suggested enhancement.

### Pull Requests

1. **Fork the Repository** and create your branch from `main`.
2. **Make your changes** in your fork.
3. **Write or update tests** for the changes you made.
4. **Ensure your code passes all tests** by running `go test ./...`.
5. **Submit a pull request** to the original repository.

## Development Setup

1. **Clone the repository**:
   ```
   git clone https://github.com/imadselka/grd.git
   cd grd
   ```

2. **Install dependencies** (there are none currently, but if any are added in the future):
   ```
   go mod download
   ```

3. **Run tests**:
   ```
   go test ./...
   ```

## Style Guidelines

### Go Code

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
- Ensure your code is formatted with `gofmt`.
- Use meaningful variable and function names.
- Write comments for public functions and types.

### Git Commit Messages

- Use the present tense ("Add feature" not "Added feature").
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...").
- Limit the first line to 72 characters or less.
- Reference issues and pull requests liberally after the first line.

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## Questions?

Feel free to [open an issue](https://github.com/imadselka/grd/issues/new) with your question or contact the maintainer directly.

Thank you for contributing to grd!
