# Evidence: Block 7 · Service (block-7)

- block type: block
- date: 2026-10-26
- resulting state: closed/verified
- story points (combined assembled cost): 4750
- execution id: 20261026.154625.159Z-e7211486

## Stories built
- Wire the ddkstat main package and keep the module standard-library only. (build-exe) [story]

## Acceptance tooling authorization
- FEATURE-BuildExecutable.md#build-exe-builds: executable=go; scope=test; authorization=existing Target environment
- FEATURE-BuildExecutable.md#build-exe-usage-exit: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 1037)
- implements: FEATURE-BuildExecutable.md (SP 976)
- context: ARCHITECTURE_compact.md (SP 493)

## Pre-build acceptance observation
- GREEN (prepassed): build-exe-builds (FEATURE-BuildExecutable.md)
  intent: go build of ./cmd/ddkstat from the module root succeeds.
  return code: 0
- GREEN (prepassed): build-exe-stdlib-only (FEATURE-BuildExecutable.md)
  intent: The module declares no third-party requirements.
  return code: 0
  stdout:
    ['module ddkstat', '', 'go 1.22']
- GREEN (prepassed): build-exe-usage-exit (FEATURE-BuildExecutable.md)
  intent: The built executable exits 2 and writes nothing to standard output when run with no arguments or an unknown report.
  return code: 0
  stdout:
    usage: ddkstat {summary|commands|targets|models|daily|failures} [--target NAME] [--since DATE] [--until DATE] LOG
    
    
    usage: ddkstat {summary|commands|targets|models|daily|failures} [--target NAME] [--since DATE] [--until DATE] LOG

## Post-build programmatic acceptance
- PASS: build-exe-builds (FEATURE-BuildExecutable.md)
  intent: go build of ./cmd/ddkstat from the module root succeeds.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: build-exe-stdlib-only (FEATURE-BuildExecutable.md)
  intent: The module declares no third-party requirements.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ['module ddkstat', '', 'go 1.22']
- PASS: build-exe-usage-exit (FEATURE-BuildExecutable.md)
  intent: The built executable exits 2 and writes nothing to standard output when run with no arguments or an unknown report.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    usage: ddkstat {summary|commands|targets|models|daily|failures} [--target NAME] [--since DATE] [--until DATE] LOG
    
    
    usage: ddkstat {summary|commands|targets|models|daily|failures} [--target NAME] [--since DATE] [--until DATE] LOG

## Build summary
The module already has most of the packages from earlier blocks. Next I'll read the existing main package, its layout test, the CLI, and the test runner.`main.go` and `go.mod` already match the spec, so this step may be verification-only. I'll run the declared gate, `sh bin/test.sh`, and check the build and the usage exit directly.The build block is already in place, so this step required no code changes. I checked it by running the declared gate and the acceptance checks directly.

RESULT: SUCCESS

FILES CHANGED:
- (none)

SUMMARY:
Verification-only step; the existing implementation already satisfies FEATURE-BuildExecutable.md.
- `cmd/ddkstat/main.go` is `package main` and contains only `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))`, with no application logic.
- `go.mod` declares `module ddkstat` and `go 1.22`, with no `require` lines. There is no `go.sum` and no `vendor/`.
- `go build ./cmd/ddkstat` from the module root exits 0 (toolchain go1.24.6).
- Running with no arguments exits 2 with 0 bytes on stdout. Running with an unknown report and a log path also exits 2 with 0 bytes on stdout.
- `sh bin/test.sh` (`go vet ./...` then `go test ./...`) exits 0 across all four packages. The test results were served from cache, so they were not re-executed in this run.
- I did not run `sources/run.sh`. It is the terminal story's check and is scoped to run once, last.
- No files were written, and no Git history was touched.

BLOCKERS:
- None.
