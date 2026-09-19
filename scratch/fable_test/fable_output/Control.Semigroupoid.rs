pub mod PureScript_Control_Semigroupoid {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Semigroupoid_Semigroupoidusd_Dict() -> &dyn Any {
        static Control_Semigroupoid_Semigroupoidusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Semigroupoid_Semigroupoidusd_Dict.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |x|
                                                                                  x.clone()))
    }
    pub fn Control_Semigroupoid_semigroupoidFn() -> &dyn Any {
        static Control_Semigroupoid_semigroupoidFn: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Semigroupoid_semigroupoidFn.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_Semigroupoidusd_Dict(),
                                                                                             &&&add(string("compose"),
                                                                                                    &&Func1::new(move
                                                                                                                     |f|
                                                                                                                     &Func1::new({
                                                                                                                                     let f
                                                                                                                                         =
                                                                                                                                         f.clone();
                                                                                                                                     move
                                                                                                                                         |g|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let g
                                                                                                                                                             =
                                                                                                                                                             g.clone();
                                                                                                                                                         move
                                                                                                                                                             |x|
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                 x))
                                                                                                                                                     })
                                                                                                                                 })),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Control_Semigroupoid_compose() -> &dyn Any {
        static Control_Semigroupoid_compose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Semigroupoid_compose.get_or_init(||
                                                     &Func1::new(move |dict|
                                                                     find(string("compose"),
                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Semigroupoid_composeFlipped() -> &dyn Any {
        static Control_Semigroupoid_composeFlipped: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Semigroupoid_composeFlipped.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroupoid|
                                                                            &Func1::new({
                                                                                            let dictSemigroupoid
                                                                                                =
                                                                                                dictSemigroupoid.clone();
                                                                                            move
                                                                                                |f|
                                                                                                &Func1::new({
                                                                                                                let f
                                                                                                                    =
                                                                                                                    f.clone();
                                                                                                                move
                                                                                                                    |g|
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                           &&&dictSemigroupoid),
                                                                                                                                                                                        g),
                                                                                                                                                     &&&f)
                                                                                                            })
                                                                                        })))
    }
}
