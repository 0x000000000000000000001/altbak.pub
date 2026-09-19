pub mod PureScript_Data_Profunctor_Strong {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Profunctor_Strong_Strongusd_Dict() -> &dyn Any {
        static Data_Profunctor_Strong_Strongusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Strong_Strongusd_Dict.get_or_init(||
                                                              &Func1::new(move
                                                                              |x|
                                                                              x.clone()))
    }
    pub fn Data_Profunctor_Strong_strongFn() -> &dyn Any {
        static Data_Profunctor_Strong_strongFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Strong_strongFn.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Strong::Data_Profunctor_Strong_Strongusd_Dict(),
                                                                                         &&&add(string("first"),
                                                                                                &&Func1::new(move
                                                                                                                 |a2b|
                                                                                                                 &Func1::new({
                                                                                                                                 let a2b
                                                                                                                                     =
                                                                                                                                     a2b.clone();
                                                                                                                                 move
                                                                                                                                     |v|
                                                                                                                                     {
                                                                                                                                         let matchValue =
                                                                                                                                             Sharpurs_Prelude::unbox(&&a2b);
                                                                                                                                         let matchValue_1:
                                                                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                 &match matchValue_1.as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                  }))
                                                                                                                                     }
                                                                                                                             })),
                                                                                                add(string("second"),
                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                      &&&PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                                                                    add(string("Profunctor0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_Profunctor::Data_Profunctor_profunctorFn()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))))
    }
    pub fn Data_Profunctor_Strong_second() -> &dyn Any {
        static Data_Profunctor_Strong_second: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Strong_second.get_or_init(||
                                                      &Func1::new(move |dict|
                                                                      find(string("second"),
                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Profunctor_Strong_first() -> &dyn Any {
        static Data_Profunctor_Strong_first: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Strong_first.get_or_init(||
                                                     &Func1::new(move |dict|
                                                                     find(string("first"),
                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Profunctor_Strong_splitStrong() -> &dyn Any {
        static Data_Profunctor_Strong_splitStrong: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Strong_splitStrong.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictSemigroupoid|
                                                                           &Func1::new({
                                                                                           let dictSemigroupoid
                                                                                               =
                                                                                               dictSemigroupoid.clone();
                                                                                           move
                                                                                               |dictStrong|
                                                                                               &Func1::new({
                                                                                                               let dictStrong
                                                                                                                   =
                                                                                                                   dictStrong.clone();
                                                                                                               move
                                                                                                                   |l|
                                                                                                                   &Func1::new({
                                                                                                                                   let l
                                                                                                                                       =
                                                                                                                                       l.clone();
                                                                                                                                   move
                                                                                                                                       |r|
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                              &&&dictSemigroupoid),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Strong::Data_Profunctor_Strong_first(),
                                                                                                                                                                                                                                                                                 &&&dictStrong),
                                                                                                                                                                                                                                              &&&l)),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Strong::Data_Profunctor_Strong_second(),
                                                                                                                                                                                                                                              &&&dictStrong),
                                                                                                                                                                                                           r))
                                                                                                                               })
                                                                                                           })
                                                                                       })))
    }
    pub fn Data_Profunctor_Strong_fanout() -> &dyn Any {
        static Data_Profunctor_Strong_fanout: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Strong_fanout.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictSemigroupoid|
                                                                      &Func1::new({
                                                                                      let dictSemigroupoid
                                                                                          =
                                                                                          dictSemigroupoid.clone();
                                                                                      move
                                                                                          |dictStrong|
                                                                                          {
                                                                                              let Profunctor0 =
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Profunctor0"),
                                                                                                                                          Sharpurs_Prelude::unbox(dictStrong)),
                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                                              &Func1::new({
                                                                                                              let Profunctor0
                                                                                                                  =
                                                                                                                  Profunctor0.clone();
                                                                                                              let dictStrong
                                                                                                                  =
                                                                                                                  dictStrong.clone();
                                                                                                              move
                                                                                                                  |l|
                                                                                                                  &Func1::new({
                                                                                                                                  let l
                                                                                                                                      =
                                                                                                                                      l.clone();
                                                                                                                                  move
                                                                                                                                      |r|
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_lcmap(),
                                                                                                                                                                                                                                             &&&Profunctor0),
                                                                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                                                                            |a|
                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                                                                                                                                                                                                    a.clone())))),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Strong::Data_Profunctor_Strong_splitStrong(),
                                                                                                                                                                                                                                                                                                                   &&&dictSemigroupoid),
                                                                                                                                                                                                                                                                                &&&dictStrong),
                                                                                                                                                                                                                                             &&&l),
                                                                                                                                                                                                          r))
                                                                                                                              })
                                                                                                          })
                                                                                          }
                                                                                  })))
    }
}
