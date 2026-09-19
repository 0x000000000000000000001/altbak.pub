pub mod PureScript_Control_Extend {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Control_Extend_FFI {
        use super::*;
        use fable_library_rust::Array_::getSlice;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_init;
        pub fn arrayExtend(f: &dyn Any) -> Func1<&dyn Any, &dyn Any> {
            Func1::new({
                           let f = f.clone();
                           move |xs|
                               {
                                   let arr = xs.clone();
                                   let res =
                                       new_init(&defaultOf(),
                                                count(arr.clone()));
                                   for i in 0_i32..=count(arr.clone()) - 1_i32
                                       {
                                       res.get_mut()[i as usize] =
                                           f(&getSlice(arr.clone(), Some(i),
                                                       None::<i32>));
                                   }
                                   &res
                               }
                       })
        }
    }
    pub fn Control_Extend_arrayExtend() -> &dyn Any {
        static Control_Extend_arrayExtend: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_arrayExtend.get_or_init(||
                                                   &Func1::new(move |arg0|
                                                                   &PureScript_Control_Extend::Control_Extend_FFI::arrayExtend(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Control_Extend_identity() -> &dyn Any {
        static Control_Extend_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_identity.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                 &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Control_Extend_Extendusd_Dict() -> &dyn Any {
        static Control_Extend_Extendusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_Extendusd_Dict.get_or_init(||
                                                      &Func1::new(move |x|
                                                                      x.clone()))
    }
    pub fn Control_Extend_extendFn() -> &dyn Any {
        static Control_Extend_extendFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_extendFn.get_or_init(||
                                                &Func1::new(move
                                                                |dictSemigroup|
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                                 &&&add(string("extend"),
                                                                                                        &&Func1::new({
                                                                                                                         let dictSemigroup
                                                                                                                             =
                                                                                                                             dictSemigroup.clone();
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
                                                                                                                                                                     |w|
                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                        let w
                                                                                                                                                                                                                            =
                                                                                                                                                                                                                            w.clone();
                                                                                                                                                                                                                        move
                                                                                                                                                                                                                            |w_prime|
                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                      &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                                                   &&&w),
                                                                                                                                                                                                                                                                                                w_prime))
                                                                                                                                                                                                                    }))
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     }),
                                                                                                        add(string("Functor0"),
                                                                                                            &&Func1::new(move
                                                                                                                             |usd__unused|
                                                                                                                             &PureScript_Data_Functor::Data_Functor_functorFn()),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>())))))
    }
    pub fn Control_Extend_extendArray() -> &dyn Any {
        static Control_Extend_extendArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_extendArray.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                    &&&add(string("extend"),
                                                                                           &&PureScript_Control_Extend::Control_Extend_arrayExtend(),
                                                                                           add(string("Functor0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Functor::Data_Functor_functorArray()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
    pub fn Control_Extend_extend() -> &dyn Any {
        static Control_Extend_extend: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_extend.get_or_init(||
                                              &Func1::new(move |dict|
                                                              find(string("extend"),
                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Extend_extendFlipped() -> &dyn Any {
        static Control_Extend_extendFlipped: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_extendFlipped.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictExtend|
                                                                     &Func1::new({
                                                                                     let dictExtend
                                                                                         =
                                                                                         dictExtend.clone();
                                                                                     move
                                                                                         |w|
                                                                                         &Func1::new({
                                                                                                         let w
                                                                                                             =
                                                                                                             w.clone();
                                                                                                         move
                                                                                                             |f|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                    &&&dictExtend),
                                                                                                                                                                                 f),
                                                                                                                                              &&&w)
                                                                                                     })
                                                                                 })))
    }
    pub fn Control_Extend_duplicate() -> &dyn Any {
        static Control_Extend_duplicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_duplicate.get_or_init(||
                                                 &Func1::new(move |dictExtend|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                     dictExtend),
                                                                                                  &&&PureScript_Control_Extend::Control_Extend_identity())))
    }
    pub fn Control_Extend_composeCoKleisliFlipped() -> &dyn Any {
        static Control_Extend_composeCoKleisliFlipped:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_composeCoKleisliFlipped.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictExtend|
                                                                               &Func1::new({
                                                                                               let dictExtend
                                                                                                   =
                                                                                                   dictExtend.clone();
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
                                                                                                                                           |w|
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                                                                                     &&&dictExtend),
                                                                                                                                                                                                                                                  &&&g),
                                                                                                                                                                                                               w))
                                                                                                                                   })
                                                                                                               })
                                                                                           })))
    }
    pub fn Control_Extend_composeCoKleisli() -> &dyn Any {
        static Control_Extend_composeCoKleisli: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Extend_composeCoKleisli.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictExtend|
                                                                        &Func1::new({
                                                                                        let dictExtend
                                                                                            =
                                                                                            dictExtend.clone();
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
                                                                                                                                    |w|
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                                                                              &&&dictExtend),
                                                                                                                                                                                                                                           &&&f),
                                                                                                                                                                                                        w))
                                                                                                                            })
                                                                                                        })
                                                                                    })))
    }
}
