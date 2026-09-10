#[global_allocator]
static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;

fn main() {
    let mut _effect = Purs_App::main();
    (_effect.unwrap_func1())(purust_core::Value::Unit);
}
