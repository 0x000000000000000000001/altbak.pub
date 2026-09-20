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
PROBES = {
    'ListOpsFFI': '''
        List<String> words = new Cons<>("a", new Cons<>("b", new Nil<>()));
        if (!foldl((String a, String b) -> a + b, "!", words).equals("!ab")) throw new AssertionError("generic list/fold");
        if (!foldl((String a, Integer b) -> a + b, "", filterEvens(range(1, 6))).equals("642")) throw new AssertionError("filter order");
    ''',
    'PrimesFFI': '''
        List<String> words = new List<>("a", new List<>("skip", new List<>("b", null)));
        List<String> selected = filter(value -> !value.equals("skip"), words);
        if (!selected.head.equals("a") || !selected.tail.head.equals("b") || selected.tail.tail != null) throw new AssertionError("generic filter/reverse");
    ''',
    'PolymorphismFFI': '''
        Monoidish<String> dictionary = new Monoidish<>("x", left -> right -> left + right);
        if (!polyLoop(dictionary, 3, "s").equals("sxxx")) throw new AssertionError("generic dictionary");
    ''',
    'ChurchFFI': '''
        Church<String> two = successor(successor(zero()));
        if (!multiply(two, two).apply(value -> value + "!").apply("s").equals("s!!!!")) throw new AssertionError("generic Church");
    ''',
    'StateMonadFFI': '''
        State<String, String> get = getState();
        State<String, Integer> action = bindState(get, initial ->
            bindState(putState(initial + "!"), ignored -> pureState(initial.length())));
        StateResult<String, Integer> result = action.run("abc");
        if (result.value() != 3 || !result.state().equals("abc!")) throw new AssertionError("generic State");
    ''',
    'LazyEvaluationFFI': '''
        int[] count = {0};
        Lazy<String> thunk = defer(() -> { count[0]++; return "word"; });
        if (!force(thunk).equals("word") || !force(thunk).equals("word") || count[0] != 2) throw new AssertionError("non-memoizing generic Lazy");
    ''',
    'RowToListFFI': '''
        var empty = new RowNil();
        var single = new RowCons<>("one", empty);
        var pair = new RowCons<>(true, single);
        if (keys(keysNil(), empty) != 0 || keys(keysCons(keysNil()), single) != 1
                || keys(keysCons(keysCons(keysNil())), pair) != 2) throw new AssertionError("typed heterogeneous row");
    ''',
}
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
            probe = '\npublic static void functionalChecks() {\n' + PROBES[stem] + '\n}\n' if stem in PROBES else ''
            (folder/(stem+'.java')).write_text('public final class '+stem+' {\n'+(ROOT/f'src/Test/{stem}.java').read_text()+probe+'\n}\n')
    lines = ['public final class NativeCheck {', 'public static void main(String[] args) {']
    lines += [stem + '.functionalChecks();' for stem in PROBES]
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
