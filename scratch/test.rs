fn main() {
    let x: i32 = loop {
        let y: i32 = if true {
            {
                continue;
            }
        } else {
            1
        };
    };
}
