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
package version

import (
	"runtime/debug"
)

const unset = "unset"

var ( // build info
	version     = unset
	builddate   = unset
	commit      = unset
	shortcommit = unset
	appname     = unset
	goversion   = unset
)

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	goversion = info.GoVersion

	var modified bool

	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			commit = setting.Value
		case "vcs.time":
			builddate = setting.Value
		case "vcs.modified":
			modified = true
		}
	}

	shortcommit = commit[:7]

	if modified {
		commit += "+CHANGES"
		shortcommit += "+CHANGES"
	}

	appname = info.Path
	version = info.Main.Version
}

// GetGoVersion returns the go version
func GetGoVersion() string {
	return goversion
}

// GetVersion returns the app version
func GetVersion() string {
	return version
}

// GetBuildDate returns the build date
func GetBuildDate() string {
	return builddate
}

// GetCommit returns the git commit
func GetCommit() string {
	return commit
}

// GetAppName returns the app name
func GetAppName() string {
	return appname
}

// GetShortCommit returns the short git commit
func GetShortCommit() string {
	return shortcommit
}
