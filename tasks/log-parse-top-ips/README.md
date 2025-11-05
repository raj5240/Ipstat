log-parse-top-ips

A small Terminal-Bench task that parses a web access log and computes the top 3 IPs by request count. The reference solution and verifier are implemented in Go.

Files
- Dockerfile: Base Ubuntu image with Go installed.
- task.yaml: Neutral, procedural task instructions.
- solution.sh: Writes `/app/web.log`, compiles and runs a Go program to create `/app/top_ips.txt`.
- tests/test_outputs.py: Minimal pytest wrapper that runs the Go verifier.
- tests/verify.go: Go verifier that checks exact file contents and ordering.
- run-tests.sh: Standard test runner for Terminal-Bench.

What the task expects
1. Create `/app/web.log` with exactly 12 lines (provided in instructions).
2. Produce `/app/top_ips.txt` with the top 3 IPs in the format `IP count`.
   - Sort by count (descending). Break ties by IP (ascending).

How the Go solution works (short)
- Read each line from `web.log`.
- Take the first token (IP) and count occurrences in a `map[string]int`.
- Sort by count desc and IP asc for ties.
- Write the top 3 as "IP count" to `top_ips.txt`.

How the Go verifier works (short)
- Ensure `web.log` and `top_ips.txt` exist.
- Compare `web.log` to the expected 12 lines (exact match).
- Check `top_ips.txt` has exactly 3 lines and matches the expected order and format.

Run locally with uv (recommended)
Prereqs: Docker Desktop running, Python 3.12, uv, Git.

```powershell
cd "C:\\Users\\raj52\\OneDrive\\Desktop"
git clone https://github.com/laude-institute/terminal-bench.git
cd .\\terminal-bench
uv sync

# Copy the task into the repo
Copy-Item -Recurse -Force "C:\\Users\\raj52\\OneDrive\\Desktop\\AfterQuery Assignment\\tasks\\log-parse-top-ips" \
  ".\\tasks\\log-parse-top-ips"

# From repo root
uv run tb run --agent oracle --task-id log-parse-top-ips   # expect pass
uv run tb run --agent nop --task-id log-parse-top-ips      # expect 0%
```

Run locally without uv (fallback)
```powershell
cd "C:\\Users\\raj52\\OneDrive\\Desktop\\terminal-bench"
py -3.12 -m venv .venv
.\\.venv\\Scripts\\Activate.ps1
pip install terminal-bench pytest

# Copy the task if not already present
Copy-Item -Recurse -Force "C:\\Users\\raj52\\OneDrive\\Desktop\\AfterQuery Assignment\\tasks\\log-parse-top-ips" \
  ".\\tasks\\log-parse-top-ips"

# From repo root
python -m tb run --agent oracle --task-id log-parse-top-ips
python -m tb run --agent nop --task-id log-parse-top-ips
```

Create the submission zip
From the task folder:
```powershell
cd "C:\\Users\\raj52\\OneDrive\\Desktop\\AfterQuery Assignment\\tasks\\log-parse-top-ips"
Compress-Archive -Path * -DestinationPath log-parse-top-ips.zip -Force
```
Rename the zip to `firstname-lastname-date.zip` before uploading.

Notes
- Keep Python at 3.12 to avoid dependency issues.
- Always run commands from the `terminal-bench` repository root for `tb run`.

