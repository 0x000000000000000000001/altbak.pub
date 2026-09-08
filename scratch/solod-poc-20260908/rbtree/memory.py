#!/usr/bin/env python3
"""One fresh parent, one measured child: child's peak RSS without sysctl.

Darwin reports ru_maxrss in bytes. This intentionally runs separately from
timed samples and excludes the Python parent's memory from the measurement.
"""
import resource
import subprocess
import sys

if sys.platform != 'darwin':
    raise SystemExit('RSS units in this helper are defined for macOS only')
result = subprocess.run(sys.argv[1:], check=False)
usage = resource.getrusage(resource.RUSAGE_CHILDREN)
print(f'POC_RSS_BYTES: {usage.ru_maxrss}', file=sys.stderr)
raise SystemExit(result.returncode)
