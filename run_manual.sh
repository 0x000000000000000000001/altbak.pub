cd /Users/0x1/Documents/htdocs/altbak.pub/run/bak/rust/output/purust_output
sed -i.bak 's/let _val_eval = act.clone();/println!("About to run act..."); let _val_eval = act.clone();/' src/lib.rs
cargo run -q --release
