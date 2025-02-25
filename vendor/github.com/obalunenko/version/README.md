# version

This package contains build information generated at build time and compiled into the binary.

```go
// Package version provides utilities for retrieving application build information
// such as version number, build date, commit hash, application name, and the Go
// language version used during the build process.
//
// This package leverages the `runtime/debug` package to read build metadata
// embedded in the binary during the build process with Go modules.
//
// Build information can be accessed via the exposed methods:
//
//   - GetGoVersion: Returns the Go version used to build the app.
//   - GetVersion: Returns the application's version.
//   - GetBuildDate: Returns the date the application was built.
//   - GetCommit: Returns the full commit hash the application was built from.
//   - GetShortCommit: Returns a shortened commit hash (7 characters).
//   - GetAppName: Returns the application's module path.
//
// When the application is built from sources with uncommitted changes, the
// commit information will be suffixed with "+CHANGES".
```
