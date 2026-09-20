use std::marker::PhantomData;

struct RowNil;
struct RowCons<Label, A, Tail> {
    value: A,
    tail: Tail,
    label: PhantomData<Label>,
}

fn field<Label, A, Tail>(value: A, tail: Tail) -> RowCons<Label, A, Tail> {
    RowCons { value, tail, label: PhantomData }
}

// The trait instance is the typed RecordKeys dictionary. Its recursive method
// counts the record's row type without reading its values. Monomorphization and
// constant folding are compiler choices, as for the PureScript type class.
trait RecordKeys {
    fn keys_impl() -> i64;
}

impl RecordKeys for RowNil {
    fn keys_impl() -> i64 { 0 }
}

impl<Label, A, Tail: RecordKeys> RecordKeys for RowCons<Label, A, Tail> {
    fn keys_impl() -> i64 { 1 + Tail::keys_impl() }
}

fn keys<Row: RecordKeys>(_: &Row) -> i64 {
    Row::keys_impl()
}

struct ALabel;
struct BLabel;
struct CLabel;
struct DLabel;
struct ELabel;

pub fn Test_RowToListFFI_runRowToListFFI(_: i64) -> i64 {
    let record = field::<ALabel, _, _>(1_i64,
        field::<BLabel, _, _>(String::from("two"),
            field::<CLabel, _, _>(true,
                field::<DLabel, _, _>(4.0_f64,
                    field::<ELabel, _, _>(String::from("five"), RowNil)))));
    keys(&record)
}
