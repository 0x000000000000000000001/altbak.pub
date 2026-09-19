pub mod PureScript_Control_Apply {
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
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub mod Control_Apply_FFI {
        use super::*;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_init;
        pub fn arrayApply(fs: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let fsArr = Sharpurs_Prelude::unbox(fs);
            let xsArr = Sharpurs_Prelude::unbox(xs);
            let lenF: i32 = count(fsArr.clone());
            let lenX: i32 = count(xsArr.clone());
            let result = new_init(&defaultOf(), lenF * lenX);
            let n: MutCell<i32> = MutCell::new(0_i32);
            for i in 0_i32..=lenF - 1_i32 {
                let f = fsArr[i].clone();
                for j in 0_i32..=lenX - 1_i32 {
                    result.get_mut()[n.get() as usize] =
                        Sharpurs_Prelude::sharpurs_apply(&f,
                                                         &xsArr[j].clone());
                    n.set(n.get() + 1_i32)
                }
            }
            &result
        }
    }
    pub fn Control_Apply_arrayApply() -> &dyn Any {
        static Control_Apply_arrayApply: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_arrayApply.get_or_init(||
                                                 &Func1::new(move |arg0|
                                                                 &Func1::new({
                                                                                 let arg0
                                                                                     =
                                                                                     arg0.clone();
                                                                                 move
                                                                                     |arg1|
                                                                                     &PureScript_Control_Apply::Control_Apply_FFI::arrayApply(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                              &Sharpurs_Prelude::unbox(arg1))
                                                                             })))
    }
    pub fn Control_Apply_Applyusd_Dict() -> &dyn Any {
        static Control_Apply_Applyusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_Applyusd_Dict.get_or_init(||
                                                    &Func1::new(move |x|
                                                                    x.clone()))
    }
    pub fn Control_Apply_applyProxy() -> &dyn Any {
        static Control_Apply_applyProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_applyProxy.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                  &&&add(string("apply"),
                                                                                         &&Func1::new(move
                                                                                                          |v|
                                                                                                          &Func1::new(move
                                                                                                                          |v1|
                                                                                                                          &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                         add(string("Functor0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &PureScript_Data_Functor::Data_Functor_functorProxy()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>()))))
    }
    pub fn Control_Apply_applyFn() -> &dyn Any {
        static Control_Apply_applyFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_applyFn.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                               &&&add(string("apply"),
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
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                   x),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                   x))
                                                                                                                                       })
                                                                                                                   })),
                                                                                      add(string("Functor0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &PureScript_Data_Functor::Data_Functor_functorFn()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
    pub fn Control_Apply_applyArray() -> &dyn Any {
        static Control_Apply_applyArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_applyArray.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                  &&&add(string("apply"),
                                                                                         &&PureScript_Control_Apply::Control_Apply_arrayApply(),
                                                                                         add(string("Functor0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &PureScript_Data_Functor::Data_Functor_functorArray()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>()))))
    }
    pub fn Control_Apply_apply() -> &dyn Any {
        static Control_Apply_apply: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_apply.get_or_init(||
                                            &Func1::new(move |dict|
                                                            find(string("apply"),
                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Apply_applyFirst() -> &dyn Any {
        static Control_Apply_applyFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_applyFirst.get_or_init(||
                                                 &Func1::new(move |dictApply|
                                                                 {
                                                                     let Functor0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                 Sharpurs_Prelude::unbox(dictApply)),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     &Func1::new({
                                                                                     let Functor0
                                                                                         =
                                                                                         Functor0.clone();
                                                                                     let dictApply
                                                                                         =
                                                                                         dictApply.clone();
                                                                                     move
                                                                                         |a|
                                                                                         &Func1::new({
                                                                                                         let a
                                                                                                             =
                                                                                                             a.clone();
                                                                                                         move
                                                                                                             |b|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                    &&&dictApply),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                                                                                                                       &&&PureScript_Data_Function::Data_Function_const()),
                                                                                                                                                                                                                    &&&a)),
                                                                                                                                              b)
                                                                                                     })
                                                                                 })
                                                                 }))
    }
    pub fn Control_Apply_applySecond() -> &dyn Any {
        static Control_Apply_applySecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_applySecond.get_or_init(||
                                                  &Func1::new(move |dictApply|
                                                                  {
                                                                      let Functor0 =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                  Sharpurs_Prelude::unbox(dictApply)),
                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                      &Func1::new({
                                                                                      let Functor0
                                                                                          =
                                                                                          Functor0.clone();
                                                                                      let dictApply
                                                                                          =
                                                                                          dictApply.clone();
                                                                                      move
                                                                                          |a|
                                                                                          &Func1::new({
                                                                                                          let a
                                                                                                              =
                                                                                                              a.clone();
                                                                                                          move
                                                                                                              |b|
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                     &&&dictApply),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                           &&&Functor0),
                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Category::Control_Category_categoryFn()))),
                                                                                                                                                                                                                     &&&a)),
                                                                                                                                               b)
                                                                                                      })
                                                                                  })
                                                                  }))
    }
    pub fn Control_Apply_lift2() -> &dyn Any {
        static Control_Apply_lift2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_lift2.get_or_init(||
                                            &Func1::new(move |dictApply|
                                                            {
                                                                let Functor0 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                            Sharpurs_Prelude::unbox(dictApply)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                &Func1::new({
                                                                                let Functor0
                                                                                    =
                                                                                    Functor0.clone();
                                                                                let dictApply
                                                                                    =
                                                                                    dictApply.clone();
                                                                                move
                                                                                    |f|
                                                                                    &Func1::new({
                                                                                                    let f
                                                                                                        =
                                                                                                        f.clone();
                                                                                                    move
                                                                                                        |a|
                                                                                                        &Func1::new({
                                                                                                                        let a
                                                                                                                            =
                                                                                                                            a.clone();
                                                                                                                        move
                                                                                                                            |b|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                   &&&dictApply),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                         &&&Functor0),
                                                                                                                                                                                                                                                                      &&&f),
                                                                                                                                                                                                                                   &&&a)),
                                                                                                                                                             b)
                                                                                                                    })
                                                                                                })
                                                                            })
                                                            }))
    }
    pub fn Control_Apply_lift3() -> &dyn Any {
        static Control_Apply_lift3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_lift3.get_or_init(||
                                            &Func1::new(move |dictApply|
                                                            {
                                                                let Functor0 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                            Sharpurs_Prelude::unbox(dictApply)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                &Func1::new({
                                                                                let Functor0
                                                                                    =
                                                                                    Functor0.clone();
                                                                                let dictApply
                                                                                    =
                                                                                    dictApply.clone();
                                                                                move
                                                                                    |f|
                                                                                    &Func1::new({
                                                                                                    let f
                                                                                                        =
                                                                                                        f.clone();
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
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                       &&&dictApply),
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                             &&&dictApply),
                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                   &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                &&&f),
                                                                                                                                                                                                                                                                                                                             &&&a)),
                                                                                                                                                                                                                                                       &&&b)),
                                                                                                                                                                                 c)
                                                                                                                                        })
                                                                                                                    })
                                                                                                })
                                                                            })
                                                            }))
    }
    pub fn Control_Apply_lift4() -> &dyn Any {
        static Control_Apply_lift4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_lift4.get_or_init(||
                                            &Func1::new(move |dictApply|
                                                            {
                                                                let Functor0 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                            Sharpurs_Prelude::unbox(dictApply)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                &Func1::new({
                                                                                let Functor0
                                                                                    =
                                                                                    Functor0.clone();
                                                                                let dictApply
                                                                                    =
                                                                                    dictApply.clone();
                                                                                move
                                                                                    |f|
                                                                                    &Func1::new({
                                                                                                    let f
                                                                                                        =
                                                                                                        f.clone();
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
                                                                                                                                                &Func1::new({
                                                                                                                                                                let c
                                                                                                                                                                    =
                                                                                                                                                                    c.clone();
                                                                                                                                                                move
                                                                                                                                                                    |d|
                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                           &&&dictApply),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                 &&&dictApply),
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&dictApply),
                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&a)),
                                                                                                                                                                                                                                                                                                                                                 &&&b)),
                                                                                                                                                                                                                                                                           &&&c)),
                                                                                                                                                                                                     d)
                                                                                                                                                            })
                                                                                                                                        })
                                                                                                                    })
                                                                                                })
                                                                            })
                                                            }))
    }
    pub fn Control_Apply_lift5() -> &dyn Any {
        static Control_Apply_lift5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Apply_lift5.get_or_init(||
                                            &Func1::new(move |dictApply|
                                                            {
                                                                let Functor0 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                            Sharpurs_Prelude::unbox(dictApply)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                &Func1::new({
                                                                                let Functor0
                                                                                    =
                                                                                    Functor0.clone();
                                                                                let dictApply
                                                                                    =
                                                                                    dictApply.clone();
                                                                                move
                                                                                    |f|
                                                                                    &Func1::new({
                                                                                                    let f
                                                                                                        =
                                                                                                        f.clone();
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
                                                                                                                                                &Func1::new({
                                                                                                                                                                let c
                                                                                                                                                                    =
                                                                                                                                                                    c.clone();
                                                                                                                                                                move
                                                                                                                                                                    |d|
                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                    let d
                                                                                                                                                                                        =
                                                                                                                                                                                        d.clone();
                                                                                                                                                                                    move
                                                                                                                                                                                        |e|
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                               &&&dictApply),
                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                     &&&dictApply),
                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&dictApply),
                                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&dictApply),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&a)),
                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&b)),
                                                                                                                                                                                                                                                                                                                                                                     &&&c)),
                                                                                                                                                                                                                                                                                               &&&d)),
                                                                                                                                                                                                                         e)
                                                                                                                                                                                })
                                                                                                                                                            })
                                                                                                                                        })
                                                                                                                    })
                                                                                                })
                                                                            })
                                                            }))
    }
}
