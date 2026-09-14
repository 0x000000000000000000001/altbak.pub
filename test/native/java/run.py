#!/usr/bin/env python3
"""Compile exact native Java snippets in a temporary package; no benchmark runs."""
from pathlib import Path
import os
import runpy
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[3]
ORACLE = runpy.run_path(str(ROOT/'test/native/oracle.py'))
CASES = ORACLE['cases']()
def java_tool(name):
    if os.environ.get(name.upper()):
        return os.environ[name.upper()]
    for home in [os.environ.get('JAVA_HOME'), '/opt/homebrew/opt/openjdk']:
        if home and (Path(home)/'bin'/name).is_file():
            return str(Path(home)/'bin'/name)
    return name
with tempfile.TemporaryDirectory(prefix='altbak-native-java-') as directory:
    folder = Path(directory)
    names = ['Bench']
    (folder/'Bench.java').write_text('public final class Bench {\n'+(ROOT/'src/Bench.java').read_text()+'\n}\n')
    for name in ORACLE['NOMINAL']:
        for suffix in ['FFI', 'FFICheatcode']:
            stem = name+suffix; names.append(stem)
            (folder/(stem+'.java')).write_text('public final class '+stem+' {\n'+(ROOT/f'src/Test/{stem}.java').read_text()+'\n}\n')
    lines = ['public final class NativeCheck {', 'public static void main(String[] args) {']
    for name, argument, expected in CASES:
        for suffix in ['FFI','FFICheatcode']:
            stem = name+suffix
            lines += ['{',f'int value = ((Number){stem}.run{stem}.apply({argument})).intValue();',f'if(value != {expected}) throw new AssertionError("{stem}({argument}): expected {expected}, got "+value);','}']
    lines += ['double before = System.nanoTime()/1000.0;', 'double clock = ((Number)Bench.benchNow.get()).doubleValue();', 'double after = System.nanoTime()/1000.0;',
              'if(clock < before-1.0 || clock > after+1.0) throw new AssertionError("clock must use monotonic microseconds");',
              'for(Object value : new Object[]{123, 1.25, "ok"}) if(!((java.util.function.Supplier<?>)Bench.opaque.apply(value)).get().equals(value)) throw new AssertionError("opaque changed its input");',
              f'System.out.println("PASS native Java: {2*len(CASES)} values and clock/opaque contracts; no benchmark timing");', '}', '}']
    (folder/'NativeCheck.java').write_text('\n'.join(lines)+'\n'); names.append('NativeCheck')
    subprocess.run([java_tool('javac'),'-d',str(folder),*map(lambda n:str(folder/(n+'.java')),names)],check=True)
    subprocess.run([java_tool('java'),'-Xss100M','-cp',str(folder),'NativeCheck'],check=True)
