pub mod PureScript_Data_Profunctor {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Profunctor_identity() -> &dyn Any {
        static Data_Profunctor_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_identity.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                  &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Profunctor_identity1() -> &dyn Any {
        static Data_Profunctor_identity1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_identity1.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                   &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Profunctor_wrap() -> &dyn Any {
        static Data_Profunctor_wrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_wrap.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_wrap(),
                                                                              &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Profunctor_unwrap() -> &dyn Any {
        static Data_Profunctor_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_unwrap.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Profunctor_Profunctorusd_Dict() -> &dyn Any {
        static Data_Profunctor_Profunctorusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Profunctorusd_Dict.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Data_Profunctor_profunctorFn() -> &dyn Any {
        static Data_Profunctor_profunctorFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_profunctorFn.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_Profunctorusd_Dict(),
                                                                                      &&&add(string("dimap"),
                                                                                             &&Func1::new(move
                                                                                                              |a2b|
                                                                                                              &Func1::new({
                                                                                                                              let a2b
                                                                                                                                  =
                                                                                                                                  a2b.clone();
                                                                                                                              move
                                                                                                                                  |c2d|
                                                                                                                                  &Func1::new({
                                                                                                                                                  let c2d
                                                                                                                                                      =
                                                                                                                                                      c2d.clone();
                                                                                                                                                  move
                                                                                                                                                      |b2c|
                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                          &&&a2b),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                             b2c),
                                                                                                                                                                                                                          &&&c2d))
                                                                                                                                              })
                                                                                                                          })),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Profunctor_dimap() -> &dyn Any {
        static Data_Profunctor_dimap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_dimap.get_or_init(||
                                              &Func1::new(move |dict|
                                                              find(string("dimap"),
                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Profunctor_lcmap() -> &dyn Any {
        static Data_Profunctor_lcmap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_lcmap.get_or_init(||
                                              &Func1::new(move
                                                              |dictProfunctor|
                                                              &Func1::new({
                                                                              let dictProfunctor
                                                                                  =
                                                                                  dictProfunctor.clone();
                                                                              move
                                                                                  |a2b|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_dimap(),
                                                                                                                                                                                         &&&dictProfunctor),
                                                                                                                                                      a2b),
                                                                                                                   &&&PureScript_Data_Profunctor::Data_Profunctor_identity())
                                                                          })))
    }
    pub fn Data_Profunctor_rmap() -> &dyn Any {
        static Data_Profunctor_rmap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_rmap.get_or_init(||
                                             &Func1::new(move |dictProfunctor|
                                                             &Func1::new({
                                                                             let dictProfunctor
                                                                                 =
                                                                                 dictProfunctor.clone();
                                                                             move
                                                                                 |b2c|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_dimap(),
                                                                                                                                                                                        &&&dictProfunctor),
                                                                                                                                                     &&&PureScript_Data_Profunctor::Data_Profunctor_identity1()),
                                                                                                                  b2c)
                                                                         })))
    }
    pub fn Data_Profunctor_unwrapIso() -> &dyn Any {
        static Data_Profunctor_unwrapIso: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_unwrapIso.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictProfunctor|
                                                                  &Func1::new({
                                                                                  let dictProfunctor
                                                                                      =
                                                                                      dictProfunctor.clone();
                                                                                  move
                                                                                      |usd__unused|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_dimap(),
                                                                                                                                                                                             &&&dictProfunctor),
                                                                                                                                                          &&&PureScript_Data_Profunctor::Data_Profunctor_wrap()),
                                                                                                                       &&&PureScript_Data_Profunctor::Data_Profunctor_unwrap())
                                                                              })))
    }
    pub fn Data_Profunctor_wrapIso() -> &dyn Any {
        static Data_Profunctor_wrapIso: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_wrapIso.get_or_init(||
                                                &Func1::new(move
                                                                |dictProfunctor|
                                                                &Func1::new({
                                                                                let dictProfunctor
                                                                                    =
                                                                                    dictProfunctor.clone();
                                                                                move
                                                                                    |usd__unused|
                                                                                    &Func1::new(move
                                                                                                    |v|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_dimap(),
                                                                                                                                                                                                           &&&dictProfunctor),
                                                                                                                                                                        &&&PureScript_Data_Profunctor::Data_Profunctor_unwrap()),
                                                                                                                                     &&&PureScript_Data_Profunctor::Data_Profunctor_wrap()))
                                                                            })))
    }
    pub fn Data_Profunctor_arr() -> &dyn Any {
        static Data_Profunctor_arr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_arr.get_or_init(||
                                            &Func1::new(move |dictCategory|
                                                            {
                                                                let identity2 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                     dictCategory);
                                                                &Func1::new({
                                                                                let identity2
                                                                                    =
                                                                                    identity2.clone();
                                                                                move
                                                                                    |dictProfunctor|
                                                                                    &Func1::new({
                                                                                                    let dictProfunctor
                                                                                                        =
                                                                                                        dictProfunctor.clone();
                                                                                                    move
                                                                                                        |f|
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_rmap(),
                                                                                                                                                                                                               &&&dictProfunctor),
                                                                                                                                                                            f),
                                                                                                                                         &&&identity2)
                                                                                                })
                                                                            })
                                                            }))
    }
}
