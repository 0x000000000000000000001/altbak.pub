#[global_allocator]
static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;

fn main() {
    purust_core::microtasks::run_main(|| { let _effect = Purs_AppX::main();
    (_effect.unwrap_func1())(purust_core::Value::Unit) });
}
