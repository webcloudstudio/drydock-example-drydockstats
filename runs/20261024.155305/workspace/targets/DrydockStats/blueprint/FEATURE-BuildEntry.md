# FEATURE: Build Entry

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | The module builds the ddkstat executable with go build using only the standard library. |
| Depends On  | ARCHITECTURE.md |
| Provides    | ddkstat executable |
| Consumes    | - |

## Programmatic Acceptance

=== AC build-entry-builds ===
Intent: go build produces the ddkstat executable from the application root.
Suite: scoped
Requires: executable=go; scope=test

import subprocess
import tempfile
import os

out = os.path.join(tempfile.mkdtemp(), "ddkstat")
result = subprocess.run(["go", "build", "-o", out, "./cmd/ddkstat"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
assert os.path.isfile(out)
=== END AC build-entry-builds ===

=== AC build-entry-stdlib-only ===
Intent: The module declares no requirements, so every import resolves to the standard library.
Suite: scoped

from pathlib import Path

assert "require" not in Path("go.mod").read_text()
=== END AC build-entry-stdlib-only ===

## User Acceptance

- None.

## Guardrails

- The build output path is supplied by the caller and never points into a Blueprint, evidence, or sources directory.
