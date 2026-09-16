"""Verify independent input families, no timings."""
from pathlib import Path
import subprocess
HERE = Path(__file__).resolve().parent
reference = None
out = []
for name in ['baseline', 'borrow', 'borrow-ins', 'borrow-ins-depth']:
    main = HERE / f'orders-{name}.rs'
    main.write_text('#![allow(warnings)]\n'
        + f'include!("kernel-{name}.rs");\ninclude!("check_orders.rs");\n')
    binary = HERE / f'orders-{name}'
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(main), '-o', str(binary)], check=True)
    result = subprocess.check_output([str(binary)], text=True, timeout=45)
    if reference is None:
        reference = result
    assert result == reference, (name, result, reference)
    out.append(name + '\n' + result)
(HERE / 'orders.log').write_text('\n'.join(out))
print('\n'.join(out))
