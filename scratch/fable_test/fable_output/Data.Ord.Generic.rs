pub mod PureScript_Data_Ord_Generic {
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
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Ord_Generic_GenericOrdusd_Dict() -> &dyn Any {
        static Data_Ord_Generic_GenericOrdusd_Dict: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Ord_Generic_GenericOrdusd_Dict.get_or_init(||
                                                            &Func1::new(move
                                                                            |x|
                                                                            x.clone()))
    }
    pub fn Data_Ord_Generic_genericOrdNoConstructors() -> &dyn Any {
        static Data_Ord_Generic_genericOrdNoConstructors:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Generic_genericOrdNoConstructors.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_GenericOrdusd_Dict(),
                                                                                                   &&&add(string("genericCompare\'"),
                                                                                                          &&Func1::new(move
                                                                                                                           |v|
                                                                                                                           &Func1::new(move
                                                                                                                                           |v1|
                                                                                                                                           &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
    }
    pub fn Data_Ord_Generic_genericOrdNoArguments() -> &dyn Any {
        static Data_Ord_Generic_genericOrdNoArguments:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Generic_genericOrdNoArguments.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_GenericOrdusd_Dict(),
                                                                                                &&&add(string("genericCompare\'"),
                                                                                                       &&Func1::new(move
                                                                                                                        |v|
                                                                                                                        &Func1::new(move
                                                                                                                                        |v1|
                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>())))
    }
    pub fn Data_Ord_Generic_genericOrdArgument() -> &dyn Any {
        static Data_Ord_Generic_genericOrdArgument: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Ord_Generic_genericOrdArgument.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictOrd|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_GenericOrdusd_Dict(),
                                                                                                             &&&add(string("genericCompare\'"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictOrd
                                                                                                                                         =
                                                                                                                                         dictOrd.clone();
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
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                        &&&dictOrd),
                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                  &&&matchValue_1)
                                                                                                                                                             }
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>()))))
    }
    pub fn Data_Ord_Generic_genericCompare_prime() -> &dyn Any {
        static Data_Ord_Generic_genericCompare_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Generic_genericCompare_prime.get_or_init(||
                                                              &Func1::new(move
                                                                              |dict|
                                                                              find(string("genericCompare\'"),
                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Ord_Generic_genericOrdConstructor() -> &dyn Any {
        static Data_Ord_Generic_genericOrdConstructor:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Generic_genericOrdConstructor.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictGenericOrd|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_GenericOrdusd_Dict(),
                                                                                                                &&&add(string("genericCompare\'"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let dictGenericOrd
                                                                                                                                            =
                                                                                                                                            dictGenericOrd.clone();
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
                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_genericCompare_prime(),
                                                                                                                                                                                                                                                                           &&&dictGenericOrd),
                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                     &&&matchValue_1)
                                                                                                                                                                }
                                                                                                                                                        })
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>()))))
    }
    pub fn Data_Ord_Generic_genericOrdProduct() -> &dyn Any {
        static Data_Ord_Generic_genericOrdProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Generic_genericOrdProduct.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictGenericOrd|
                                                                           &Func1::new({
                                                                                           let dictGenericOrd
                                                                                               =
                                                                                               dictGenericOrd.clone();
                                                                                           move
                                                                                               |dictGenericOrd1|
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_GenericOrdusd_Dict(),
                                                                                                                                &&&add(string("genericCompare\'"),
                                                                                                                                       &&Func1::new({
                                                                                                                                                        let dictGenericOrd1
                                                                                                                                                            =
                                                                                                                                                            dictGenericOrd1.clone();
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
                                                                                                                                                                                    let matchValue_3:
                                                                                                                                                                                            LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_genericCompare_prime(),
                                                                                                                                                                                                                                                                                                                         &&&dictGenericOrd),
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
                                                                                                                                                                                                                                                      }));
                                                                                                                                                                                    if let Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                                                                           =
                                                                                                                                                                                           matchValue_3.as_ref()
                                                                                                                                                                                       {
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_genericCompare_prime(),
                                                                                                                                                                                                                                                                                               &&&dictGenericOrd1),
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
                                                                                                                                                                                                                            })
                                                                                                                                                                                    } else {
                                                                                                                                                                                        &matchValue_3
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                        })
                                                                                                                                                    }),
                                                                                                                                       empty::<string,
                                                                                                                                               &dyn Any>()))
                                                                                       })))
    }
    pub fn Data_Ord_Generic_genericOrdSum() -> &dyn Any {
        static Data_Ord_Generic_genericOrdSum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Generic_genericOrdSum.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictGenericOrd|
                                                                       &Func1::new({
                                                                                       let dictGenericOrd
                                                                                           =
                                                                                           dictGenericOrd.clone();
                                                                                       move
                                                                                           |dictGenericOrd1|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_GenericOrdusd_Dict(),
                                                                                                                            &&&add(string("genericCompare\'"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let dictGenericOrd1
                                                                                                                                                        =
                                                                                                                                                        dictGenericOrd1.clone();
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
                                                                                                                                                                                        LrcPtr<Data_Generic_Rep_Sum> =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                let matchValue_1:
                                                                                                                                                                                        LrcPtr<Data_Generic_Rep_Sum> =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                if let Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(matchValue_1_0)
                                                                                                                                                                                       =
                                                                                                                                                                                       matchValue.as_ref()
                                                                                                                                                                                   {
                                                                                                                                                                                    if let Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(matchValue_1_0_0)
                                                                                                                                                                                           =
                                                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                                                       {
                                                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                    } else {
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_genericCompare_prime(),
                                                                                                                                                                                                                                                                                               &&&dictGenericOrd1),
                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                   Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(x)
                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(x)
                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                            })
                                                                                                                                                                                    }
                                                                                                                                                                                } else {
                                                                                                                                                                                    if let Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                           =
                                                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                                                       {
                                                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                    } else {
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_genericCompare_prime(),
                                                                                                                                                                                                                                                                                               &&&dictGenericOrd),
                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                   Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(x)
                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(x)
                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                            })
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                    })
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>()))
                                                                                   })))
    }
    pub fn Data_Ord_Generic_genericCompare() -> &dyn Any {
        static Data_Ord_Generic_genericCompare: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Generic_genericCompare.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictGeneric|
                                                                        &Func1::new({
                                                                                        let dictGeneric
                                                                                            =
                                                                                            dictGeneric.clone();
                                                                                        move
                                                                                            |dictGenericOrd|
                                                                                            &Func1::new({
                                                                                                            let dictGenericOrd
                                                                                                                =
                                                                                                                dictGenericOrd.clone();
                                                                                                            move
                                                                                                                |x|
                                                                                                                &Func1::new({
                                                                                                                                let x
                                                                                                                                    =
                                                                                                                                    x.clone();
                                                                                                                                move
                                                                                                                                    |y|
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Generic::Data_Ord_Generic_genericCompare_prime(),
                                                                                                                                                                                                                                           &&&dictGenericOrd),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                              &&&dictGeneric),
                                                                                                                                                                                                                                           &&&x)),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                           &&&dictGeneric),
                                                                                                                                                                                                        y))
                                                                                                                            })
                                                                                                        })
                                                                                    })))
    }
}
