# Evidence: Block 1 · Service (block-1)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 8800
- execution id: 20261024.142407.402Z-8b26b5c7

## Stories built
- Parse execution records, skip invalid and duplicate lines, and apply filters. (log-intake) [story]

## Stacked context
- stack: go_compact.md (SP 1374)
- compass: COMPASS.md (SP 677)
- implements: FEATURE-LogIntake.md (SP 1029)
- context: llm_sample.jsonl (SP 3862)
- stack: common.md (SP 1812)

## Build directory changes
- bin/test.sh
- cmd/ddkstat/main.go
- ddkstat
- go.mod
- internal/cli/run.go
- internal/logintake/intake.go
- internal/logintake/intake_test.go

## Pre-build acceptance observation
- GREEN (prepassed): intake-duplicate-skip (FEATURE-LogIntake.md)
  intent: The parser skips a repeated execution ID, retains the first record, and reports one skipped line.
  return code: 0
- RED: intake-invalid-skip (FEATURE-LogIntake.md)
  intent: The command continues successfully with skipped malformed and incomplete lines.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      errno_num = 2
      err_msg = 'No such file or directory'
      err_filename = './ddkstat'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "intake-invalid-skip.py", line 17, in <module>
        result = subprocess.run(["./ddkstat", "summary", path], capture_output=True, text=True)
                 ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/subprocess.py", line 548, in run
        with Popen(*popenargs, **kwargs) as process:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/subprocess.py", line 1026, in __init__
        self._execute_child(args, executable, preexec_fn, close_fds,
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/subprocess.py", line 1955, in _execute_child
        raise child_exception_type(errno_num, err_msg, err_filename)
    FileNotFoundError: [Errno 2] No such file or directory: './ddkstat'
- RED: intake-filtering (FEATURE-LogIntake.md)
  intent: Target and inclusive date filters select matching records while skipped counts remain whole-file.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      errno_num = 2
      err_msg = 'No such file or directory'
      err_filename = './ddkstat'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "intake-filtering.py", line 25, in <module>
        result = subprocess.run(
                 ^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/subprocess.py", line 548, in run
        with Popen(*popenargs, **kwargs) as process:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/subprocess.py", line 1026, in __init__
        self._execute_child(args, executable, preexec_fn, close_fds,
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/subprocess.py", line 1955, in _execute_child
        raise child_exception_type(errno_num, err_msg, err_filename)
    FileNotFoundError: [Errno 2] No such file or directory: './ddkstat'

## Post-build programmatic acceptance
- PASS: intake-duplicate-skip (FEATURE-LogIntake.md)
  intent: The parser skips a repeated execution ID, retains the first record, and reports one skipped line.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: intake-invalid-skip (FEATURE-LogIntake.md)
  intent: The command continues successfully with skipped malformed and incomplete lines.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: intake-filtering (FEATURE-LogIntake.md)
  intent: Target and inclusive date filters select matching records while skipped counts remain whole-file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Repair attempts
- attempt 0 (initial build): failed; 1/3 checks model=gpt-6-luna; execution 20261024.142222.151Z-f1318b3a; reason: programmatic acceptance failed: intake-invalid-skip, intake-filtering
- attempt 1 (repair 1): failed; 1/3 checks model=gpt-6-luna; execution 20261024.142334.761Z-8219e1aa; reason: programmatic acceptance failed: intake-invalid-skip, intake-filtering
- attempt 2 (repair 2): built; 3/3 checks model=gpt-6-luna; execution 20261024.142407.402Z-8b26b5c7

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- ddkstat

SUMMARY:
Built the `./ddkstat` executable in the build directory, which the two failing acceptance checks expected to exist. Both malformed-line skipping and filtered summary checks now pass. `sh bin/test.sh` also passes.

BLOCKERS:
- None
