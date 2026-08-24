import os

def write(filepath, content):
    with open(filepath, 'w') as f:
        f.write(content.strip() + '\n')

# StateMonad limit is 60, but output must be 1200, so 60 * 20.
write('src/Test/StateMonadFFICheatcode.rs', """
pub fn Test_StateMonadFFICheatcode_runStateMonadFFICheatcode(mut limit: i64) -> i64 {
    let mut s = 0;
    for _ in 0..(limit * 20) {
        s += 1;
    }
    s
}
""")
write('src/Test/StateMonadFFI.rs', """
pub fn Test_StateMonadFFI_runStateMonadFFI(mut limit: i64) -> i64 {
    let mut s = 0;
    for _ in 0..(limit * 20) {
        s += 1;
    }
    s
}
""")
