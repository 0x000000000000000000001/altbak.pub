pub mod PureScript_Control_Monad {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Monadusd_Dict() -> &dyn Any {
        static Control_Monad_Monadusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Monadusd_Dict.get_or_init(||
                                                    &Func1::new(move |x|
                                                                    x.clone()))
    }
    pub fn Control_Monad_whenM() -> &dyn Any {
        static Control_Monad_whenM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_whenM.get_or_init(||
                                            &Func1::new(move |dictMonad|
                                                            {
                                                                let Bind1 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                            Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                let Applicative0 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                            Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                &Func1::new({
                                                                                let Applicative0
                                                                                    =
                                                                                    Applicative0.clone();
                                                                                let Bind1
                                                                                    =
                                                                                    Bind1.clone();
                                                                                move
                                                                                    |mb|
                                                                                    &Func1::new({
                                                                                                    let mb
                                                                                                        =
                                                                                                        mb.clone();
                                                                                                    move
                                                                                                        |m|
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                               &&&Bind1),
                                                                                                                                                                            &&&mb),
                                                                                                                                         &&&Func1::new({
                                                                                                                                                           let m
                                                                                                                                                               =
                                                                                                                                                               m.clone();
                                                                                                                                                           move
                                                                                                                                                               |b|
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_when(),
                                                                                                                                                                                                                                                                      &&&Applicative0),
                                                                                                                                                                                                                                   b),
                                                                                                                                                                                                &&&m)
                                                                                                                                                       }))
                                                                                                })
                                                                            })
                                                            }))
    }
    pub fn Control_Monad_unlessM() -> &dyn Any {
        static Control_Monad_unlessM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_unlessM.get_or_init(||
                                              &Func1::new(move |dictMonad|
                                                              {
                                                                  let Bind1 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                              Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                  let Applicative0 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                              Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                  &Func1::new({
                                                                                  let Applicative0
                                                                                      =
                                                                                      Applicative0.clone();
                                                                                  let Bind1
                                                                                      =
                                                                                      Bind1.clone();
                                                                                  move
                                                                                      |mb|
                                                                                      &Func1::new({
                                                                                                      let mb
                                                                                                          =
                                                                                                          mb.clone();
                                                                                                      move
                                                                                                          |m|
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                 &&&Bind1),
                                                                                                                                                                              &&&mb),
                                                                                                                                           &&&Func1::new({
                                                                                                                                                             let m
                                                                                                                                                                 =
                                                                                                                                                                 m.clone();
                                                                                                                                                             move
                                                                                                                                                                 |b|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_unless(),
                                                                                                                                                                                                                                                                        &&&Applicative0),
                                                                                                                                                                                                                                     b),
                                                                                                                                                                                                  &&&m)
                                                                                                                                                         }))
                                                                                                  })
                                                                              })
                                                              }))
    }
    pub fn Control_Monad_monadProxy() -> &dyn Any {
        static Control_Monad_monadProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_monadProxy.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                  &&&add(string("Applicative0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Control_Applicative::Control_Applicative_applicativeProxy()),
                                                                                         add(string("Bind1"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused_1|
                                                                                                              &PureScript_Control_Bind::Control_Bind_bindProxy()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>()))))
    }
    pub fn Control_Monad_monadFn() -> &dyn Any {
        static Control_Monad_monadFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_monadFn.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                               &&&add(string("Applicative0"),
                                                                                      &&Func1::new(move
                                                                                                       |usd__unused|
                                                                                                       &PureScript_Control_Applicative::Control_Applicative_applicativeFn()),
                                                                                      add(string("Bind1"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused_1|
                                                                                                           &PureScript_Control_Bind::Control_Bind_bindFn()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
    pub fn Control_Monad_monadArray() -> &dyn Any {
        static Control_Monad_monadArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_monadArray.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                  &&&add(string("Applicative0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Control_Applicative::Control_Applicative_applicativeArray()),
                                                                                         add(string("Bind1"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused_1|
                                                                                                              &PureScript_Control_Bind::Control_Bind_bindArray()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>()))))
    }
    pub fn Control_Monad_liftM1() -> &dyn Any {
        static Control_Monad_liftM1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_liftM1.get_or_init(||
                                             &Func1::new(move |dictMonad|
                                                             {
                                                                 let Bind1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                             Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                 let Applicative0 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                             Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                 &Func1::new({
                                                                                 let Applicative0
                                                                                     =
                                                                                     Applicative0.clone();
                                                                                 let Bind1
                                                                                     =
                                                                                     Bind1.clone();
                                                                                 move
                                                                                     |f|
                                                                                     &Func1::new({
                                                                                                     let f
                                                                                                         =
                                                                                                         f.clone();
                                                                                                     move
                                                                                                         |a|
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                &&&Bind1),
                                                                                                                                                                             a),
                                                                                                                                          &&&Func1::new(move
                                                                                                                                                            |a_prime|
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                &&&Applicative0),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                a_prime))))
                                                                                                 })
                                                                             })
                                                             }))
    }
    pub fn Control_Monad_ap() -> &dyn Any {
        static Control_Monad_ap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ap.get_or_init(||
                                         &Func1::new(move |dictMonad|
                                                         {
                                                             let Bind1 =
                                                                 Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                         Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                  &&&Sharpurs_Prelude::Prim_undefined());
                                                             let Applicative0 =
                                                                 Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                         Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                  &&&Sharpurs_Prelude::Prim_undefined());
                                                             &Func1::new({
                                                                             let Applicative0
                                                                                 =
                                                                                 Applicative0.clone();
                                                                             let Bind1
                                                                                 =
                                                                                 Bind1.clone();
                                                                             move
                                                                                 |f|
                                                                                 &Func1::new({
                                                                                                 let f
                                                                                                     =
                                                                                                     f.clone();
                                                                                                 move
                                                                                                     |a|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                            &&&Bind1),
                                                                                                                                                                         &&&f),
                                                                                                                                      &&&Func1::new({
                                                                                                                                                        let a
                                                                                                                                                            =
                                                                                                                                                            a.clone();
                                                                                                                                                        move
                                                                                                                                                            |f_prime|
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                   &&&Bind1),
                                                                                                                                                                                                                                &&&a),
                                                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                                                               let f_prime
                                                                                                                                                                                                                   =
                                                                                                                                                                                                                   f_prime.clone();
                                                                                                                                                                                                               move
                                                                                                                                                                                                                   |a_prime|
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                       &&&Applicative0),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&f_prime,
                                                                                                                                                                                                                                                                                       a_prime))
                                                                                                                                                                                                           }))
                                                                                                                                                    }))
                                                                                             })
                                                                         })
                                                         }))
    }
}
