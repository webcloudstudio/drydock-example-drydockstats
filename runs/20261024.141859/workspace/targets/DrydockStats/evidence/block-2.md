# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 10398
- execution id: 20261024.142503.963Z-cb2e3a1b

## Stories built
- Aggregate and format the six CSV reports. (report-generation) [story]
- Deliver and verify the command-line interface, test runner, and full conformance gate. (cli-delivery) [story]

## Missing context
- context: INSTRUCTIONS.md

## Stacked context
- stack: go_compact.md (SP 1374)
- compass: COMPASS.md (SP 677)
- implements: FEATURE-Reports.md (SP 777)
- context: llm_sample.jsonl (SP 3862)
- stack: common_compact.md (SP 1079)
- implements: FEATURE-CLI.md (SP 888)
- context: run.sh (SP 1624)

## Build directory changes
- bin/test.sh
- internal/cli/run.go
- internal/cli/run_test.go
- internal/logintake/intake.go
- internal/reports/reports.go
- internal/reports/reports_test.go

## Pre-build acceptance observation
- GREEN (prepassed): reports-conformance (FEATURE-Reports.md)
  intent: The six reports, filters, skipped-line handling, formatting, ordering, and empty-result output pass the supplied known-answer suite.
  return code: 0
  stdout:
    ?   	drydockstats/cmd/ddkstat	[no test files]
    ?   	drydockstats/internal/cli	[no test files]
    ok  	drydockstats/internal/logintake	(cached)
- GREEN (prepassed): model-selection (FEATURE-Reports.md)
  intent: The report aggregation assigns a record to its executed model when present and its requested model otherwise.
  return code: 0
  stdout:
    ?   	drydockstats/cmd/ddkstat	[no test files]
    ?   	drydockstats/internal/cli	[no test files]
    ok  	drydockstats/internal/logintake	(cached)
- RED: cli-error-contract (FEATURE-CLI.md)
  intent: Missing arguments and an unreadable log produce their specified exit statuses with no standard output.
  return code: 1
  stdout:
    PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    FAIL A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,9d0
      < command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
      < build,17,0,129705,8268539,209318,71.1
      < score release,5,0,31244,2819917,46278,22.7
      < plan,4,0,1862,2244487,59759,25.1
      < analyze,3,0,32267,1662483,40833,20.2
      < diagnose,2,2,33682,1010463,17499,10.3
      < import_compass,1,0,843,263315,7751,0.9
      < plan repair,1,0,32923,784113,17142,9.1
      < TOTAL,33,2,262526,17053317,398580,159.4
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A3: ddkstat targets /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,5d0
      < target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
      < CommonMark,12,1,5547,5786766,119677,3.7919,51.5
      < ReadingList,11,0,4104,5161259,145664,1.2485,47.6
      < Toml,10,1,252875,6105292,133239,0.4841,60.3
      < TOTAL,33,2,262526,17053317,398580,5.5245,159.4
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A4: ddkstat models /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,7d0
      < provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
      < claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
      < claude,claude-sonnet-5-5,11,1,5547,5786766,119677,3.7919
      < claude,sonnet,1,0,0,0,0,0.0000
      < codex,gpt-5.6-luna,4,0,110058,2466012,40962,0.0000
      < gemini,gemini-3-flash,6,1,142817,3639280,92277,0.4841
      < TOTAL,,33,2,262526,17053317,398580,5.5245
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A5: ddkstat daily /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,5d0
      < date,calls,failed,tokens,minutes
      < 2026-10-14,14,1,6944443,86.3
      < 2026-10-15,12,1,7016674,57.8
      < 2026-10-16,7,0,3753306,15.3
      < TOTAL,33,2,17714423,159.4
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A6: ddkstat failures /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,3d0
      < started_at,target,command,model,returncode,timed_out
      < 2026-10-14T02:38:41.904Z,Toml,diagnose,gemini-3-flash,1,false
      < 2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    PASS B1: ddkstat summary --target Toml --since 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    FAIL B2: ddkstat commands --until 2026-10-14 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,9d0
      < command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
      < build,6,0,67524,2290028,87813,37.5
      < analyze,2,0,31740,1346394,25842,12.0
      < plan,2,0,467,1010857,35477,15.6
      < diagnose,1,1,33526,770301,17499,9.8
      < import_compass,1,0,843,263315,7751,0.9
      < plan repair,1,0,32923,784113,17142,9.1
      < score release,1,0,379,110560,9949,1.3
      < TOTAL,14,1,167402,6575568,201473,86.3
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL B3: ddkstat daily --since 2026-10-15 --until 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,3d0
      < date,calls,failed,tokens,minutes
      < 2026-10-15,12,1,7016674,57.8
      < TOTAL,12,1,7016674,57.8
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL B4: ddkstat failures --target CommonMark /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,2d0
      < started_at,target,command,model,returncode,timed_out
      < 2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL B5: ddkstat models --target ReadingList /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,3d0
      < provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
      < claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
      < TOTAL,,11,0,4104,5161259,145664,1.2485
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL C1: ddkstat commands --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,2d0
      < command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
      < TOTAL,0,0,0,0,0,0.0
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    PASS C2: ddkstat summary --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS E1: ddkstat summary /tmp/tmp.zw9ZpDVTYC/no-such-log.jsonl
    PASS U1: ddkstat 
    PASS U2: ddkstat frobnicate /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U3: ddkstat summary
    PASS U4: ddkstat summary --since 2026-13-01 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U5: ddkstat summary --colour red /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U6: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    DrydockStats: 10 passed, 10 failed (of 20)
  stderr:
    --- drydock: values at failure ---
      missing = CompletedProcess(args=['sh', 'sources/run.sh'], returncode=1, stdout='PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl\nFAIL A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl\n  exit: expected 0, got 2\n  1,9d0\n  < c… (+6247 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-error-contract.py", line 13, in <module>
        assert missing.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: test-runner (FEATURE-CLI.md)
  intent: The project test runner exits successfully after vet and all Go tests pass.
  return code: 1
  stdout:
    PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    FAIL A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,9d0
      < command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
      < build,17,0,129705,8268539,209318,71.1
      < score release,5,0,31244,2819917,46278,22.7
      < plan,4,0,1862,2244487,59759,25.1
      < analyze,3,0,32267,1662483,40833,20.2
      < diagnose,2,2,33682,1010463,17499,10.3
      < import_compass,1,0,843,263315,7751,0.9
      < plan repair,1,0,32923,784113,17142,9.1
      < TOTAL,33,2,262526,17053317,398580,159.4
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A3: ddkstat targets /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,5d0
      < target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
      < CommonMark,12,1,5547,5786766,119677,3.7919,51.5
      < ReadingList,11,0,4104,5161259,145664,1.2485,47.6
      < Toml,10,1,252875,6105292,133239,0.4841,60.3
      < TOTAL,33,2,262526,17053317,398580,5.5245,159.4
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A4: ddkstat models /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,7d0
      < provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
      < claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
      < claude,claude-sonnet-5-5,11,1,5547,5786766,119677,3.7919
      < claude,sonnet,1,0,0,0,0,0.0000
      < codex,gpt-5.6-luna,4,0,110058,2466012,40962,0.0000
      < gemini,gemini-3-flash,6,1,142817,3639280,92277,0.4841
      < TOTAL,,33,2,262526,17053317,398580,5.5245
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A5: ddkstat daily /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,5d0
      < date,calls,failed,tokens,minutes
      < 2026-10-14,14,1,6944443,86.3
      < 2026-10-15,12,1,7016674,57.8
      < 2026-10-16,7,0,3753306,15.3
      < TOTAL,33,2,17714423,159.4
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A6: ddkstat failures /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,3d0
      < started_at,target,command,model,returncode,timed_out
      < 2026-10-14T02:38:41.904Z,Toml,diagnose,gemini-3-flash,1,false
      < 2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    PASS B1: ddkstat summary --target Toml --since 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    FAIL B2: ddkstat commands --until 2026-10-14 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,9d0
      < command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
      < build,6,0,67524,2290028,87813,37.5
      < analyze,2,0,31740,1346394,25842,12.0
      < plan,2,0,467,1010857,35477,15.6
      < diagnose,1,1,33526,770301,17499,9.8
      < import_compass,1,0,843,263315,7751,0.9
      < plan repair,1,0,32923,784113,17142,9.1
      < score release,1,0,379,110560,9949,1.3
      < TOTAL,14,1,167402,6575568,201473,86.3
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL B3: ddkstat daily --since 2026-10-15 --until 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,3d0
      < date,calls,failed,tokens,minutes
      < 2026-10-15,12,1,7016674,57.8
      < TOTAL,12,1,7016674,57.8
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL B4: ddkstat failures --target CommonMark /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,2d0
      < started_at,target,command,model,returncode,timed_out
      < 2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL B5: ddkstat models --target ReadingList /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,3d0
      < provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
      < claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
      < TOTAL,,11,0,4104,5161259,145664,1.2485
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL C1: ddkstat commands --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,2d0
      < command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
      < TOTAL,0,0,0,0,0,0.0
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    PASS C2: ddkstat summary --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS E1: ddkstat summary /tmp/tmp.MlKo4oBhpD/no-such-log.jsonl
    PASS U1: ddkstat 
    PASS U2: ddkstat frobnicate /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U3: ddkstat summary
    PASS U4: ddkstat summary --since 2026-13-01 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U5: ddkstat summary --colour red /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U6: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    DrydockStats: 10 passed, 10 failed (of 20)
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/run.sh'], returncode=1, stdout='PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl\nFAIL A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl\n  exit: expected 0, got 2\n  1,9d0\n  < c… (+6247 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "test-runner.py", line 13, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: full-conformance (FEATURE-CLI.md)
  intent: The terminal story passes all 20 supplied grader checks.
  return code: 1
  stdout:
    PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    FAIL A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,9d0
      < command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
      < build,17,0,129705,8268539,209318,71.1
      < score release,5,0,31244,2819917,46278,22.7
      < plan,4,0,1862,2244487,59759,25.1
      < analyze,3,0,32267,1662483,40833,20.2
      < diagnose,2,2,33682,1010463,17499,10.3
      < import_compass,1,0,843,263315,7751,0.9
      < plan repair,1,0,32923,784113,17142,9.1
      < TOTAL,33,2,262526,17053317,398580,159.4
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A3: ddkstat targets /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,5d0
      < target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
      < CommonMark,12,1,5547,5786766,119677,3.7919,51.5
      < ReadingList,11,0,4104,5161259,145664,1.2485,47.6
      < Toml,10,1,252875,6105292,133239,0.4841,60.3
      < TOTAL,33,2,262526,17053317,398580,5.5245,159.4
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A4: ddkstat models /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,7d0
      < provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
      < claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
      < claude,claude-sonnet-5-5,11,1,5547,5786766,119677,3.7919
      < claude,sonnet,1,0,0,0,0,0.0000
      < codex,gpt-5.6-luna,4,0,110058,2466012,40962,0.0000
      < gemini,gemini-3-flash,6,1,142817,3639280,92277,0.4841
      < TOTAL,,33,2,262526,17053317,398580,5.5245
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A5: ddkstat daily /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,5d0
      < date,calls,failed,tokens,minutes
      < 2026-10-14,14,1,6944443,86.3
      < 2026-10-15,12,1,7016674,57.8
      < 2026-10-16,7,0,3753306,15.3
      < TOTAL,33,2,17714423,159.4
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL A6: ddkstat failures /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,3d0
      < started_at,target,command,model,returncode,timed_out
      < 2026-10-14T02:38:41.904Z,Toml,diagnose,gemini-3-flash,1,false
      < 2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    PASS B1: ddkstat summary --target Toml --since 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    FAIL B2: ddkstat commands --until 2026-10-14 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,9d0
      < command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
      < build,6,0,67524,2290028,87813,37.5
      < analyze,2,0,31740,1346394,25842,12.0
      < plan,2,0,467,1010857,35477,15.6
      < diagnose,1,1,33526,770301,17499,9.8
      < import_compass,1,0,843,263315,7751,0.9
      < plan repair,1,0,32923,784113,17142,9.1
      < score release,1,0,379,110560,9949,1.3
      < TOTAL,14,1,167402,6575568,201473,86.3
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL B3: ddkstat daily --since 2026-10-15 --until 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,3d0
      < date,calls,failed,tokens,minutes
      < 2026-10-15,12,1,7016674,57.8
      < TOTAL,12,1,7016674,57.8
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL B4: ddkstat failures --target CommonMark /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,2d0
      < started_at,target,command,model,returncode,timed_out
      < 2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL B5: ddkstat models --target ReadingList /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,3d0
      < provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
      < claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
      < TOTAL,,11,0,4104,5161259,145664,1.2485
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    FAIL C1: ddkstat commands --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
      exit: expected 0, got 2
      1,2d0
      < command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
      < TOTAL,0,0,0,0,0,0.0
      stderr: usage: ddkstat summary [--target TARGET] [--since DATE] [--until DATE] LOG
    PASS C2: ddkstat summary --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS E1: ddkstat summary /tmp/tmp.uf1J7bLYZn/no-such-log.jsonl
    PASS U1: ddkstat 
    PASS U2: ddkstat frobnicate /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U3: ddkstat summary
    PASS U4: ddkstat summary --since 2026-13-01 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U5: ddkstat summary --colour red /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U6: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    DrydockStats: 10 passed, 10 failed (of 20)
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/run.sh'], returncode=1, stdout='PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl\nFAIL A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl\n  exit: expected 0, got 2\n  1,9d0\n  < c… (+6247 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "full-conformance.py", line 13, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: reports-conformance (FEATURE-Reports.md)
  intent: The six reports, filters, skipped-line handling, formatting, ordering, and empty-result output pass the supplied known-answer suite.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ?   	drydockstats/cmd/ddkstat	[no test files]
    ok  	drydockstats/internal/cli	(cached)
    ok  	drydockstats/internal/logintake	(cached)
    ok  	drydockstats/internal/reports	(cached)
- PASS: model-selection (FEATURE-Reports.md)
  intent: The report aggregation assigns a record to its executed model when present and its requested model otherwise.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ?   	drydockstats/cmd/ddkstat	[no test files]
    ok  	drydockstats/internal/cli	(cached)
    ok  	drydockstats/internal/logintake	(cached)
    ok  	drydockstats/internal/reports	(cached)
- PASS: cli-error-contract (FEATURE-CLI.md)
  intent: Missing arguments and an unreadable log produce their specified exit statuses with no standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A3: ddkstat targets /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A4: ddkstat models /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A5: ddkstat daily /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A6: ddkstat failures /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B1: ddkstat summary --target Toml --since 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B2: ddkstat commands --until 2026-10-14 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B3: ddkstat daily --since 2026-10-15 --until 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B4: ddkstat failures --target CommonMark /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B5: ddkstat models --target ReadingList /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS C1: ddkstat commands --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS C2: ddkstat summary --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS E1: ddkstat summary /tmp/tmp.2rQYXpuLud/no-such-log.jsonl
    PASS U1: ddkstat 
    PASS U2: ddkstat frobnicate /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U3: ddkstat summary
    PASS U4: ddkstat summary --since 2026-13-01 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U5: ddkstat summary --colour red /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U6: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    DrydockStats: 20 passed, 0 failed (of 20)
- PASS: test-runner (FEATURE-CLI.md)
  intent: The project test runner exits successfully after vet and all Go tests pass.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A3: ddkstat targets /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A4: ddkstat models /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A5: ddkstat daily /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A6: ddkstat failures /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B1: ddkstat summary --target Toml --since 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B2: ddkstat commands --until 2026-10-14 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B3: ddkstat daily --since 2026-10-15 --until 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B4: ddkstat failures --target CommonMark /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B5: ddkstat models --target ReadingList /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS C1: ddkstat commands --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS C2: ddkstat summary --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS E1: ddkstat summary /tmp/tmp.v5ZGLdlRdu/no-such-log.jsonl
    PASS U1: ddkstat 
    PASS U2: ddkstat frobnicate /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U3: ddkstat summary
    PASS U4: ddkstat summary --since 2026-13-01 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U5: ddkstat summary --colour red /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U6: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    DrydockStats: 20 passed, 0 failed (of 20)
- PASS: full-conformance (FEATURE-CLI.md)
  intent: The terminal story passes all 20 supplied grader checks.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    PASS A1: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A2: ddkstat commands /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A3: ddkstat targets /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A4: ddkstat models /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A5: ddkstat daily /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS A6: ddkstat failures /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B1: ddkstat summary --target Toml --since 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B2: ddkstat commands --until 2026-10-14 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B3: ddkstat daily --since 2026-10-15 --until 2026-10-15 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B4: ddkstat failures --target CommonMark /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS B5: ddkstat models --target ReadingList /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS C1: ddkstat commands --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS C2: ddkstat summary --target NoSuchTarget /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS E1: ddkstat summary /tmp/tmp.FUJygyUWV4/no-such-log.jsonl
    PASS U1: ddkstat 
    PASS U2: ddkstat frobnicate /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U3: ddkstat summary
    PASS U4: ddkstat summary --since 2026-13-01 /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U5: ddkstat summary --colour red /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    PASS U6: ddkstat summary /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl /mnt/c/Users/barlo/projects/drydock/uat/DrydockStats/runs/20261024.141859/build/DrydockStats/sources/llm_sample.jsonl
    DrydockStats: 20 passed, 0 failed (of 20)

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- `bin/test.sh`
- `internal/cli/run.go`
- `internal/cli/run_test.go`
- `internal/logintake/intake.go`
- `internal/reports/reports.go`
- `internal/reports/reports_test.go`

SUMMARY:
Implemented all six CSV reports, model selection, filtering, CLI error handling, and the vet-and-test runner. `sh bin/test.sh` passed; `sh sources/run.sh` passed all 20 checks.

BLOCKERS:
- None
