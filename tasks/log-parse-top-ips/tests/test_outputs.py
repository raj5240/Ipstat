import subprocess

def test_go_verifier_exits_zero():
    subprocess.run(["/usr/bin/env", "bash", "-lc", "go run /tests/verify.go"], check=True)