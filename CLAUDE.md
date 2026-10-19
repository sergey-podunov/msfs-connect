# CLAUDE.md

## Project Overview

msfs-connect is a Go wrapper over the Microsoft Flight Simulator SimConnect SDK (`SimConnect.dll`). It is used by apps in the sibling folders (e.g. `../atc_freq`) via the Go workspace `../go.work`.

Package `simconnect` (`github.com/sergey-podunov/msfs-connect/simconnect`) contains:
- `Connection` interface, `DllConnection` (real DLL calls via windows syscalls), `MockConnection` (testify/mock)
- Generated `SIMCONNECT_*` types and consts (`defs_generated.go`)
- `ExceptionName`, `CString` / `TrimCString`

App-specific logic (frequencies, METAR, clouds) does not belong here.

## Build Commands

```bash
# Regenerate defs_generated.go from SimConnect.h (needs g++)
make generate

# Build and test
make test
```

## Misc
Don't edit defs_generated.go directly. When some new consts or types from SimConnect.h are needed — add them to `simconnect/defs.go` and run `make generate`. `defs_generated.go` is committed so consumers can build without g++ and SimConnect.h.

Keep comments sparse: no doc comment that restates the name or signature. Comment only what isn't obvious from the code: non-obvious behavior, constraints, the reason behind a decision.
