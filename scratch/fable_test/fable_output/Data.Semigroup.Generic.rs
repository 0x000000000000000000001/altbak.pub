pub mod PureScript_Data_Semigroup_Generic {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Product;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Semigroup_Generic_GenericSemigroupusd_Dict() -> &dyn Any {
        static Data_Semigroup_Generic_GenericSemigroupusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Generic_GenericSemigroupusd_Dict.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |x|
                                                                                        x.clone()))
    }
    pub fn Data_Semigroup_Generic_genericSemigroupNoConstructors()
     -> &dyn Any {
        static Data_Semigroup_Generic_genericSemigroupNoConstructors:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Generic_genericSemigroupNoConstructors.get_or_init(||
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Generic::Data_Semigroup_Generic_GenericSemigroupusd_Dict(),
                                                                                                               &&&add(string("genericAppend\'"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |a|
                                                                                                                                       &Func1::new({
                                                                                                                                                       let a
                                                                                                                                                           =
                                                                                                                                                           a.clone();
                                                                                                                                                       move
                                                                                                                                                           |v|
                                                                                                                                                           &a
                                                                                                                                                   })),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>())))
    }
    pub fn Data_Semigroup_Generic_genericSemigroupNoArguments() -> &dyn Any {
        static Data_Semigroup_Generic_genericSemigroupNoArguments:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Generic_genericSemigroupNoArguments.get_or_init(||
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Generic::Data_Semigroup_Generic_GenericSemigroupusd_Dict(),
                                                                                                            &&&add(string("genericAppend\'"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |a|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let a
                                                                                                                                                        =
                                                                                                                                                        a.clone();
                                                                                                                                                    move
                                                                                                                                                        |v|
                                                                                                                                                        &a
                                                                                                                                                })),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>())))
    }
    pub fn Data_Semigroup_Generic_genericSemigroupArgument() -> &dyn Any {
        static Data_Semigroup_Generic_genericSemigroupArgument:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Generic_genericSemigroupArgument.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictSemigroup|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Generic::Data_Semigroup_Generic_GenericSemigroupusd_Dict(),
                                                                                                                         &&&add(string("genericAppend\'"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let dictSemigroup
                                                                                                                                                     =
                                                                                                                                                     dictSemigroup.clone();
                                                                                                                                                 move
                                                                                                                                                     |v|
                                                                                                                                                     &Func1::new({
                                                                                                                                                                     let v
                                                                                                                                                                         =
                                                                                                                                                                         v.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |v1|
                                                                                                                                                                         {
                                                                                                                                                                             let matchValue =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                             let matchValue_1 =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                       &&&dictSemigroup),
                                                                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                                                                 &&&matchValue_1))
                                                                                                                                                                         }
                                                                                                                                                                 })
                                                                                                                                             }),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Semigroup_Generic_genericAppend_prime() -> &dyn Any {
        static Data_Semigroup_Generic_genericAppend_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Generic_genericAppend_prime.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dict|
                                                                                   find(string("genericAppend\'"),
                                                                                        Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semigroup_Generic_genericSemigroupConstructor() -> &dyn Any {
        static Data_Semigroup_Generic_genericSemigroupConstructor:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Generic_genericSemigroupConstructor.get_or_init(||
                                                                           &Func1::new(move
                                                                                           |dictGenericSemigroup|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Generic::Data_Semigroup_Generic_GenericSemigroupusd_Dict(),
                                                                                                                            &&&add(string("genericAppend\'"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let dictGenericSemigroup
                                                                                                                                                        =
                                                                                                                                                        dictGenericSemigroup.clone();
                                                                                                                                                    move
                                                                                                                                                        |v|
                                                                                                                                                        &Func1::new({
                                                                                                                                                                        let v
                                                                                                                                                                            =
                                                                                                                                                                            v.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |v1|
                                                                                                                                                                            {
                                                                                                                                                                                let matchValue =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                let matchValue_1 =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Generic::Data_Semigroup_Generic_genericAppend_prime(),
                                                                                                                                                                                                                                                                                                                          &&&dictGenericSemigroup),
                                                                                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                                                                                    &&&matchValue_1))
                                                                                                                                                                            }
                                                                                                                                                                    })
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Semigroup_Generic_genericSemigroupProduct() -> &dyn Any {
        static Data_Semigroup_Generic_genericSemigroupProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Generic_genericSemigroupProduct.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictGenericSemigroup|
                                                                                       &Func1::new({
                                                                                                       let dictGenericSemigroup
                                                                                                           =
                                                                                                           dictGenericSemigroup.clone();
                                                                                                       move
                                                                                                           |dictGenericSemigroup1|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Generic::Data_Semigroup_Generic_GenericSemigroupusd_Dict(),
                                                                                                                                            &&&add(string("genericAppend\'"),
                                                                                                                                                   &&Func1::new({
                                                                                                                                                                    let dictGenericSemigroup1
                                                                                                                                                                        =
                                                                                                                                                                        dictGenericSemigroup1.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |v|
                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                        let v
                                                                                                                                                                                            =
                                                                                                                                                                                            v.clone();
                                                                                                                                                                                        move
                                                                                                                                                                                            |v1|
                                                                                                                                                                                            {
                                                                                                                                                                                                let matchValue:
                                                                                                                                                                                                        LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                let matchValue_1:
                                                                                                                                                                                                        LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                                                    Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                &LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Generic::Data_Semigroup_Generic_genericAppend_prime(),
                                                                                                                                                                                                                                                                                                                                                                               &&&dictGenericSemigroup),
                                                                                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                                   Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Generic::Data_Semigroup_Generic_genericAppend_prime(),
                                                                                                                                                                                                                                                                                                                                                                               &&&dictGenericSemigroup1),
                                                                                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                                   Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                            })))
                                                                                                                                                                                            }
                                                                                                                                                                                    })
                                                                                                                                                                }),
                                                                                                                                                   empty::<string,
                                                                                                                                                           &dyn Any>()))
                                                                                                   })))
    }
    pub fn Data_Semigroup_Generic_genericAppend() -> &dyn Any {
        static Data_Semigroup_Generic_genericAppend: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_Generic_genericAppend.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictGeneric|
                                                                             &Func1::new({
                                                                                             let dictGeneric
                                                                                                 =
                                                                                                 dictGeneric.clone();
                                                                                             move
                                                                                                 |dictGenericSemigroup|
                                                                                                 &Func1::new({
                                                                                                                 let dictGenericSemigroup
                                                                                                                     =
                                                                                                                     dictGenericSemigroup.clone();
                                                                                                                 move
                                                                                                                     |x|
                                                                                                                     &Func1::new({
                                                                                                                                     let x
                                                                                                                                         =
                                                                                                                                         x.clone();
                                                                                                                                     move
                                                                                                                                         |y|
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                                                                                                             &&&dictGeneric),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Generic::Data_Semigroup_Generic_genericAppend_prime(),
                                                                                                                                                                                                                                                                                   &&&dictGenericSemigroup),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                                                                      &&&dictGeneric),
                                                                                                                                                                                                                                                                                   &&&x)),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                                   &&&dictGeneric),
                                                                                                                                                                                                                                                y)))
                                                                                                                                 })
                                                                                                             })
                                                                                         })))
    }
}
