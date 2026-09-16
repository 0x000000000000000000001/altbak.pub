#[derive(Clone)]
struct DictD { e: i64, f: i64 }
#[derive(Clone)]
struct DictB { c: i64, d: Box<DictD> }
#[derive(Clone)]
struct DictR { a: i64, b: Box<DictB> }

pub fn Test_RecordsFFI_runRecordsFFI(mut limit: i64) -> i64 {
    let mut rec = Box::new(DictR {
        a: 0,
        b: Box::new(DictB {
            c: 0,
            d: Box::new(DictD { e: 0, f: 0 }),
        }),
    });
    while limit > 0 {
        rec = Box::new(DictR {
            a: rec.a + 1,
            b: Box::new(DictB {
                c: rec.b.c + 2,
                d: Box::new(DictD {
                    e: rec.b.d.e + 3,
                    f: rec.b.d.f + (limit % 5),
                }),
            }),
        });
        limit -= 1;
    }
    rec.b.d.f
}
