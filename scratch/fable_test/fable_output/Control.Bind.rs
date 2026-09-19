pub mod PureScript_Control_Bind {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub mod Control_Bind_FFI {
        use super::*;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_copy;
        use fable_library_rust::NativeArray_::new_empty;
        pub fn arrayBind(xs: &dyn Any, f: &dyn Any) -> &dyn Any {
            let arr = Sharpurs_Prelude::unbox(xs);
            let result = new_empty::<&dyn Any>();
            for idx in 0_i32..=count(arr.clone()) - 1_i32 {
                let res =
                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(f,
                                                                              &arr[idx].clone()));
                defaultOf::<&dyn Any>()
            }
            &new_copy(result)
        }
    }
    pub fn Control_Bind_arrayBind() -> &dyn Any {
        static Control_Bind_arrayBind: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_arrayBind.get_or_init(||
                                               &Func1::new(move |xs|
                                                               Func1::new({
                                                                              let xs
                                                                                  =
                                                                                  xs.clone();
                                                                              move
                                                                                  |f|
                                                                                  PureScript_Control_Bind::Control_Bind_FFI::arrayBind(&xs,
                                                                                                                                       f)
                                                                          })))
    }
    pub fn Control_Bind_identity() -> &dyn Any {
        static Control_Bind_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_identity.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                               &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Control_Bind_Bindusd_Dict() -> &dyn Any {
        static Control_Bind_Bindusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_Bindusd_Dict.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  x.clone()))
    }
    pub fn Control_Bind_Discardusd_Dict() -> &dyn Any {
        static Control_Bind_Discardusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_Discardusd_Dict.get_or_init(||
                                                     &Func1::new(move |x|
                                                                     x.clone()))
    }
    pub fn Control_Bind_discard() -> &dyn Any {
        static Control_Bind_discard: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_discard.get_or_init(||
                                             &Func1::new(move |dict|
                                                             find(string("discard"),
                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Bind_bindProxy() -> &dyn Any {
        static Control_Bind_bindProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_bindProxy.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                &&&add(string("bind"),
                                                                                       &&Func1::new(move
                                                                                                        |v|
                                                                                                        &Func1::new(move
                                                                                                                        |v1|
                                                                                                                        &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                       add(string("Apply0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Control_Apply::Control_Apply_applyProxy()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))))
    }
    pub fn Control_Bind_bindFn() -> &dyn Any {
        static Control_Bind_bindFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_bindFn.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                             &&&add(string("bind"),
                                                                                    &&Func1::new(move
                                                                                                     |m|
                                                                                                     &Func1::new({
                                                                                                                     let m
                                                                                                                         =
                                                                                                                         m.clone();
                                                                                                                     move
                                                                                                                         |f|
                                                                                                                         &Func1::new({
                                                                                                                                         let f
                                                                                                                                             =
                                                                                                                                             f.clone();
                                                                                                                                         move
                                                                                                                                             |x|
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&m,
                                                                                                                                                                                                                                                    x)),
                                                                                                                                                                              x)
                                                                                                                                     })
                                                                                                                 })),
                                                                                    add(string("Apply0"),
                                                                                        &&Func1::new(move
                                                                                                         |usd__unused|
                                                                                                         &PureScript_Control_Apply::Control_Apply_applyFn()),
                                                                                        empty::<string,
                                                                                                &dyn Any>()))))
    }
    pub fn Control_Bind_bindArray() -> &dyn Any {
        static Control_Bind_bindArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_bindArray.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                &&&add(string("bind"),
                                                                                       &&PureScript_Control_Bind::Control_Bind_arrayBind(),
                                                                                       add(string("Apply0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Control_Apply::Control_Apply_applyArray()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))))
    }
    pub fn Control_Bind_bind() -> &dyn Any {
        static Control_Bind_bind: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_bind.get_or_init(||
                                          &Func1::new(move |dict|
                                                          find(string("bind"),
                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Bind_bindFlipped() -> &dyn Any {
        static Control_Bind_bindFlipped: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_bindFlipped.get_or_init(||
                                                 &Func1::new(move |dictBind|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                     dictBind))))
    }
    pub fn Control_Bind_composeKleisliFlipped() -> &dyn Any {
        static Control_Bind_composeKleisliFlipped: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_composeKleisliFlipped.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictBind|
                                                                           &Func1::new({
                                                                                           let dictBind
                                                                                               =
                                                                                               dictBind.clone();
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
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bindFlipped(),
                                                                                                                                                                                                                                              &&&dictBind),
                                                                                                                                                                                                           &&&f),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                           a))
                                                                                                                               })
                                                                                                           })
                                                                                       })))
    }
    pub fn Control_Bind_composeKleisli() -> &dyn Any {
        static Control_Bind_composeKleisli: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_composeKleisli.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictBind|
                                                                    &Func1::new({
                                                                                    let dictBind
                                                                                        =
                                                                                        dictBind.clone();
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
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                       &&&dictBind),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                       a)),
                                                                                                                                                                 &&&g)
                                                                                                                        })
                                                                                                    })
                                                                                })))
    }
    pub fn Control_Bind_discardProxy() -> &dyn Any {
        static Control_Bind_discardProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_discardProxy.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Discardusd_Dict(),
                                                                                   &&&add(string("discard"),
                                                                                          &&Func1::new(move
                                                                                                           |dictBind|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                            dictBind)),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Control_Bind_discardUnit() -> &dyn Any {
        static Control_Bind_discardUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_discardUnit.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Discardusd_Dict(),
                                                                                  &&&add(string("discard"),
                                                                                         &&Func1::new(move
                                                                                                          |dictBind|
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                           dictBind)),
                                                                                         empty::<string,
                                                                                                 &dyn Any>())))
    }
    pub fn Control_Bind_ifM() -> &dyn Any {
        static Control_Bind_ifM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_ifM.get_or_init(||
                                         &Func1::new(move |dictBind|
                                                         &Func1::new({
                                                                         let dictBind
                                                                             =
                                                                             dictBind.clone();
                                                                         move
                                                                             |cond|
                                                                             &Func1::new({
                                                                                             let cond
                                                                                                 =
                                                                                                 cond.clone();
                                                                                             move
                                                                                                 |t|
                                                                                                 &Func1::new({
                                                                                                                 let t
                                                                                                                     =
                                                                                                                     t.clone();
                                                                                                                 move
                                                                                                                     |f|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                            &&&dictBind),
                                                                                                                                                                                         &&&cond),
                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                        let f
                                                                                                                                                                            =
                                                                                                                                                                            f.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |cond_prime|
                                                                                                                                                                            {
                                                                                                                                                                                let matchValue =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(cond_prime);
                                                                                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                 &matchValue)
                                                                                                                                                                                    {
                                                                                                                                                                                    0_i32
                                                                                                                                                                                    =>
                                                                                                                                                                                    &t,
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    &f,
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                    }))
                                                                                                             })
                                                                                         })
                                                                     })))
    }
    pub fn Control_Bind_join() -> &dyn Any {
        static Control_Bind_join: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Bind_join.get_or_init(||
                                          &Func1::new(move |dictBind|
                                                          &Func1::new({
                                                                          let dictBind
                                                                              =
                                                                              dictBind.clone();
                                                                          move
                                                                              |m|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                     &&&dictBind),
                                                                                                                                                  m),
                                                                                                               &&&PureScript_Control_Bind::Control_Bind_identity())
                                                                      })))
    }
}
