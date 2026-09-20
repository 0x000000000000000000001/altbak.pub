pub fn check_structure() {
    let dictionary = Monoidish {
        mempty_: String::from("x"),
        mappend_: Box::new(|left: String| Box::new(move |right: String| left.clone() + &right)),
    };
    assert_eq!(poly_loop(&dictionary, 3, String::from("!")), "!xxx");
    assert_eq!(poly_loop(&dictionary, 0, String::from("unchanged")), "unchanged");
}
