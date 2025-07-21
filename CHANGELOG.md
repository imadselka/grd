# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial implementation of `Try`, `Then`, `Catch`, and `Finally` functions
- Comprehensive test suite
- Example code
- Documentation

## [1.0.0] - 2025-07-20

### Added
- First stable release
- Core functionality:
  - `Try[T](fn func() (T, error)) *TryResult[T]`
  - `Then(fn func(T) (T, error)) *TryResult[T]`
  - `Catch(fn func(error) T) T`
  - `Finally(fn func()) *TryResult[T]`
- Type-safe error handling with Go generics
- Chainable operations
- Full test coverage
- Documentation and examples

[Unreleased]: https://github.com/imadselka/grd/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/imadselka/grd/releases/tag/v1.0.0
