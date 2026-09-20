pub fn check_structure() {
    struct Name;
    struct Flag;
    assert_eq!(keys(&RowNil), 0);
    let record = field::<Name, _, _>(String::from("value"), field::<Flag, _, _>(true, RowNil));
    assert_eq!(keys(&record), 2);
    assert_eq!(record.value, "value");
    assert!(record.tail.value);
    let other_values = field::<Name, _, _>(String::new(), field::<Flag, _, _>(false, RowNil));
    assert_eq!(keys(&other_values), 2, "the dictionary counts types, not field values");
}
