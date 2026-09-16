"""Bound a diagnostic command and terminate only the process group it starts."""
import argparse
import os
from pathlib import Path
import signal
import subprocess
import time

parser = argparse.ArgumentParser()
parser.add_argument('--seconds', type=int, required=True)
parser.add_argument('--log', type=Path, required=True)
parser.add_argument('command', nargs=argparse.REMAINDER)
args = parser.parse_args()
command = args.command[1:] if args.command[0] == '--' else args.command
start = time.monotonic()
with args.log.open('w') as log:
    process = subprocess.Popen(command, stdout=log, stderr=subprocess.STDOUT, start_new_session=True)
    try:
        result = process.wait(timeout=args.seconds)
    except subprocess.TimeoutExpired:
        os.killpg(process.pid, signal.SIGTERM)
        process.wait(timeout=5)
        result = 124
print('exit:', result, 'seconds:', round(time.monotonic() - start, 2), 'log:', args.log)
print('\n'.join(args.log.read_text().splitlines()[-25:]))
raise SystemExit(result)
