from pathlib import Path
import subprocess
here = Path(__file__).resolve().parent
source = (here / 'validate.rs').read_text()
source = source.replace('mod region_alloc;', '#[path="region_local.rs"] mod region_alloc;')
(here / 'validate_local.rs').write_text(source)
subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=3', str(here/'validate_local.rs'), '-o', str(here/'validate_local')], check=True)
result = subprocess.check_output([str(here/'validate_local')], text=True, timeout=60)
(here/'validation-local.txt').write_text(result)
print(result, end='')
