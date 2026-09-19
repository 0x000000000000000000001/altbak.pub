pub mod PureScript_Control_Biapply {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Biapply_Biapplyusd_Dict() -> &dyn Any {
        static Control_Biapply_Biapplyusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapply_Biapplyusd_Dict.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        x.clone()))
    }
    pub fn Control_Biapply_biapplyTuple() -> &dyn Any {
        static Control_Biapply_biapplyTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapply_biapplyTuple.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_Biapplyusd_Dict(),
                                                                                      &&&add(string("biapply"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              &Func1::new({
                                                                                                                              let v
                                                                                                                                  =
                                                                                                                                  v.clone();
                                                                                                                              move
                                                                                                                                  |v1|
                                                                                                                                  {
                                                                                                                                      let matchValue:
                                                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                      let matchValue_1:
                                                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&match matchValue.as_ref()
                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                  },
                                                                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                  }),
                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&match matchValue.as_ref()
                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                  },
                                                                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                  })))
                                                                                                                                  }
                                                                                                                          })),
                                                                                             add(string("Bifunctor0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorTuple()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Control_Biapply_biapply() -> &dyn Any {
        static Control_Biapply_biapply: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapply_biapply.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("biapply"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Biapply_biapplyFirst() -> &dyn Any {
        static Control_Biapply_biapplyFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapply_biapplyFirst.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBiapply|
                                                                     {
                                                                         let Bifunctor0 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictBiapply)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                         &Func1::new({
                                                                                         let Bifunctor0
                                                                                             =
                                                                                             Bifunctor0.clone();
                                                                                         let dictBiapply
                                                                                             =
                                                                                             dictBiapply.clone();
                                                                                         move
                                                                                             |a|
                                                                                             &Func1::new({
                                                                                                             let a
                                                                                                                 =
                                                                                                                 a.clone();
                                                                                                             move
                                                                                                                 |b|
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_biapply(),
                                                                                                                                                                                                                        &&&dictBiapply),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                                    &&&Bifunctor0),
                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Category::Control_Category_categoryFn()))),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Category::Control_Category_categoryFn())))),
                                                                                                                                                                                                                        &&&a)),
                                                                                                                                                  b)
                                                                                                         })
                                                                                     })
                                                                     }))
    }
    pub fn Control_Biapply_biapplySecond() -> &dyn Any {
        static Control_Biapply_biapplySecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapply_biapplySecond.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictBiapply|
                                                                      {
                                                                          let Bifunctor0 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                      Sharpurs_Prelude::unbox(dictBiapply)),
                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                          &Func1::new({
                                                                                          let Bifunctor0
                                                                                              =
                                                                                              Bifunctor0.clone();
                                                                                          let dictBiapply
                                                                                              =
                                                                                              dictBiapply.clone();
                                                                                          move
                                                                                              |a|
                                                                                              &Func1::new({
                                                                                                              let a
                                                                                                                  =
                                                                                                                  a.clone();
                                                                                                              move
                                                                                                                  |b|
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_biapply(),
                                                                                                                                                                                                                         &&&dictBiapply),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                                     &&&Bifunctor0),
                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Function::Data_Function_const()),
                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Function::Data_Function_const())),
                                                                                                                                                                                                                         &&&a)),
                                                                                                                                                   b)
                                                                                                          })
                                                                                      })
                                                                      }))
    }
    pub fn Control_Biapply_bilift2() -> &dyn Any {
        static Control_Biapply_bilift2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapply_bilift2.get_or_init(||
                                                &Func1::new(move |dictBiapply|
                                                                {
                                                                    let Bifunctor0 =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                Sharpurs_Prelude::unbox(dictBiapply)),
                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                    &Func1::new({
                                                                                    let Bifunctor0
                                                                                        =
                                                                                        Bifunctor0.clone();
                                                                                    let dictBiapply
                                                                                        =
                                                                                        dictBiapply.clone();
                                                                                    move
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
                                                                                                                                |a|
                                                                                                                                &Func1::new({
                                                                                                                                                let a
                                                                                                                                                    =
                                                                                                                                                    a.clone();
                                                                                                                                                move
                                                                                                                                                    |b|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_biapply(),
                                                                                                                                                                                                                                                           &&&dictBiapply),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&Bifunctor0),
                                                                                                                                                                                                                                                                                                                                                                    &&&f),
                                                                                                                                                                                                                                                                                                                                 &&&g)),
                                                                                                                                                                                                                                                           &&&a)),
                                                                                                                                                                                     b)
                                                                                                                                            })
                                                                                                                        })
                                                                                                    })
                                                                                })
                                                                }))
    }
    pub fn Control_Biapply_bilift3() -> &dyn Any {
        static Control_Biapply_bilift3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Biapply_bilift3.get_or_init(||
                                                &Func1::new(move |dictBiapply|
                                                                {
                                                                    let Bifunctor0 =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                Sharpurs_Prelude::unbox(dictBiapply)),
                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                    &Func1::new({
                                                                                    let Bifunctor0
                                                                                        =
                                                                                        Bifunctor0.clone();
                                                                                    let dictBiapply
                                                                                        =
                                                                                        dictBiapply.clone();
                                                                                    move
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
                                                                                                                                |a|
                                                                                                                                &Func1::new({
                                                                                                                                                let a
                                                                                                                                                    =
                                                                                                                                                    a.clone();
                                                                                                                                                move
                                                                                                                                                    |b|
                                                                                                                                                    &Func1::new({
                                                                                                                                                                    let b
                                                                                                                                                                        =
                                                                                                                                                                        b.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |c|
                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_biapply(),
                                                                                                                                                                                                                                                                               &&&dictBiapply),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_biapply(),
                                                                                                                                                                                                                                                                                                                                                     &&&dictBiapply),
                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&Bifunctor0),
                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&g)),
                                                                                                                                                                                                                                                                                                                                                     &&&a)),
                                                                                                                                                                                                                                                                               &&&b)),
                                                                                                                                                                                                         c)
                                                                                                                                                                })
                                                                                                                                            })
                                                                                                                        })
                                                                                                    })
                                                                                })
                                                                }))
    }
}
