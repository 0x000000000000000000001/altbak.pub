pub mod PureScript_Data_Show_Generic {
    use super::*;
    use fable_library_rust::Array_::equals;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::NativeArray_::count;
    use fable_library_rust::NativeArray_::new_array;
    use fable_library_rust::NativeArray_::new_empty;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Product;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    use fable_library_rust::System::Lazy_1;
    pub mod Data_Show_Generic_FFI {
        use super::*;
        pub fn intercalate() -> &dyn Any {
            static intercalate: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            intercalate.get_or_init(|| Sharpurs_Prelude::undefined())
        }
    }
    pub fn Data_Show_Generic_intercalate() -> &dyn Any {
        static Data_Show_Generic_intercalate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_intercalate.get_or_init(||
                                                      &PureScript_Data_Show_Generic::Data_Show_Generic_FFI::intercalate())
    }
    pub fn Data_Show_Generic_GenericShowArgsusd_Dict() -> &dyn Any {
        static Data_Show_Generic_GenericShowArgsusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_GenericShowArgsusd_Dict.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |x|
                                                                                  x.clone()))
    }
    pub fn Data_Show_Generic_GenericShowusd_Dict() -> &dyn Any {
        static Data_Show_Generic_GenericShowusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_GenericShowusd_Dict.get_or_init(||
                                                              &Func1::new(move
                                                                              |x|
                                                                              x.clone()))
    }
    pub fn Data_Show_Generic_genericShowArgsNoArguments() -> &dyn Any {
        static Data_Show_Generic_genericShowArgsNoArguments:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_genericShowArgsNoArguments.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_GenericShowArgsusd_Dict(),
                                                                                                      &&&add(string("genericShowArgs"),
                                                                                                             &&Func1::new(move
                                                                                                                              |v|
                                                                                                                              &new_empty::<&dyn Any>()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
    }
    pub fn Data_Show_Generic_genericShowArgsArgument() -> &dyn Any {
        static Data_Show_Generic_genericShowArgsArgument:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_genericShowArgsArgument.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictShow|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_GenericShowArgsusd_Dict(),
                                                                                                                   &&&add(string("genericShowArgs"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictShow
                                                                                                                                               =
                                                                                                                                               dictShow.clone();
                                                                                                                                           move
                                                                                                                                               |v|
                                                                                                                                               {
                                                                                                                                                   let a =
                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                   &new_array(&[Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                    &&&dictShow),
                                                                                                                                                                                                 &&&a)])
                                                                                                                                               }
                                                                                                                                       }),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Show_Generic_genericShowArgs() -> &dyn Any {
        static Data_Show_Generic_genericShowArgs: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_genericShowArgs.get_or_init(||
                                                          &Func1::new(move
                                                                          |dict|
                                                                          find(string("genericShowArgs"),
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Show_Generic_genericShowArgsProduct() -> &dyn Any {
        static Data_Show_Generic_genericShowArgsProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_genericShowArgsProduct.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictGenericShowArgs|
                                                                                 &Func1::new({
                                                                                                 let dictGenericShowArgs
                                                                                                     =
                                                                                                     dictGenericShowArgs.clone();
                                                                                                 move
                                                                                                     |dictGenericShowArgs1|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_GenericShowArgsusd_Dict(),
                                                                                                                                      &&&add(string("genericShowArgs"),
                                                                                                                                             &&Func1::new({
                                                                                                                                                              let dictGenericShowArgs1
                                                                                                                                                                  =
                                                                                                                                                                  dictGenericShowArgs1.clone();
                                                                                                                                                              move
                                                                                                                                                                  |v|
                                                                                                                                                                  {
                                                                                                                                                                      let matchValue:
                                                                                                                                                                              LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupArray()),
                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_genericShowArgs(),
                                                                                                                                                                                                                                                                                                                &&&dictGenericShowArgs),
                                                                                                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                    Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                })),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_genericShowArgs(),
                                                                                                                                                                                                                                                                             &&&dictGenericShowArgs1),
                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                 Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                             }))
                                                                                                                                                                  }
                                                                                                                                                          }),
                                                                                                                                             empty::<string,
                                                                                                                                                     &dyn Any>()))
                                                                                             })))
    }
    pub fn Data_Show_Generic_genericShowConstructor() -> &dyn Any {
        static Data_Show_Generic_genericShowConstructor:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_genericShowConstructor.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictGenericShowArgs|
                                                                                 &Func1::new({
                                                                                                 let dictGenericShowArgs
                                                                                                     =
                                                                                                     dictGenericShowArgs.clone();
                                                                                                 move
                                                                                                     |dictIsSymbol|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_GenericShowusd_Dict(),
                                                                                                                                      &&&add(string("genericShow\'"),
                                                                                                                                             &&Func1::new({
                                                                                                                                                              let dictIsSymbol
                                                                                                                                                                  =
                                                                                                                                                                  dictIsSymbol.clone();
                                                                                                                                                              move
                                                                                                                                                                  |v|
                                                                                                                                                                  {
                                                                                                                                                                      let a =
                                                                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                                                                      let ctor =
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                              &&&dictIsSymbol),
                                                                                                                                                                                                           &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                      let matchValue =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_genericShowArgs(),
                                                                                                                                                                                                                                                                        &&&dictGenericShowArgs),
                                                                                                                                                                                                                                     &&&a));
                                                                                                                                                                      if if !equals(matchValue.clone(),
                                                                                                                                                                                    new_empty::<&dyn Any>())
                                                                                                                                                                            {
                                                                                                                                                                             count(matchValue.clone())
                                                                                                                                                                                 ==
                                                                                                                                                                                 0_i32
                                                                                                                                                                         } else {
                                                                                                                                                                             false
                                                                                                                                                                         }
                                                                                                                                                                         {
                                                                                                                                                                          &ctor
                                                                                                                                                                      } else {
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                              &&&string("(")),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_intercalate(),
                                                                                                                                                                                                                                                                                                                                                       &&&string(" ")),
                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupArray()),
                                                                                                                                                                                                                                                                                                                                                                                          &&&new_array(&[&ctor])),
                                                                                                                                                                                                                                                                                                                                                       &&&matchValue))),
                                                                                                                                                                                                                                              &&&string(")")))
                                                                                                                                                                      }
                                                                                                                                                                  }
                                                                                                                                                          }),
                                                                                                                                             empty::<string,
                                                                                                                                                     &dyn Any>()))
                                                                                             })))
    }
    pub fn Data_Show_Generic_genericShow_prime() -> &dyn Any {
        static Data_Show_Generic_genericShow_prime: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Show_Generic_genericShow_prime.get_or_init(||
                                                            &Func1::new(move
                                                                            |dict|
                                                                            find(string("genericShow\'"),
                                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Show_Generic_genericShowNoConstructors_004030() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_GenericShowusd_Dict(),
                                         &&&add(string("genericShow\'"),
                                                &&Func1::new({
                                                                 let Data_Show_Generic_genericShowNoConstructors_004030_002d1
                                                                     =
                                                                     Data_Show_Generic_genericShowNoConstructors_004030_002d1.clone();
                                                                 move |a|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_genericShow_prime(),
                                                                                                                                         &&&Data_Show_Generic_genericShowNoConstructors_004030_002d1.Value),
                                                                                                      a)
                                                             }),
                                                empty::<string, &dyn Any>()))
    }
    pub fn Data_Show_Generic_genericShowNoConstructors_004030_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Show_Generic_genericShowNoConstructors_004030_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Show_Generic_genericShowNoConstructors_004030_002d1.get_or_init(||
                                                                                 Lazy(Data_Show_Generic_genericShowNoConstructors_004030.clone()))
    }
    pub fn Data_Show_Generic_genericShowNoConstructors() -> &dyn Any {
        static Data_Show_Generic_genericShowNoConstructors:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_genericShowNoConstructors.get_or_init(||
                                                                    Data_Show_Generic_genericShowNoConstructors_004030_002d1.Value)
    }
    pub fn Data_Show_Generic_genericShowSum() -> &dyn Any {
        static Data_Show_Generic_genericShowSum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_genericShowSum.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictGenericShow|
                                                                         &Func1::new({
                                                                                         let dictGenericShow
                                                                                             =
                                                                                             dictGenericShow.clone();
                                                                                         move
                                                                                             |dictGenericShow1|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_GenericShowusd_Dict(),
                                                                                                                              &&&add(string("genericShow\'"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let dictGenericShow1
                                                                                                                                                          =
                                                                                                                                                          dictGenericShow1.clone();
                                                                                                                                                      move
                                                                                                                                                          |v|
                                                                                                                                                          {
                                                                                                                                                              let matchValue:
                                                                                                                                                                      LrcPtr<Data_Generic_Rep_Sum> =
                                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                                              match matchValue.as_ref()
                                                                                                                                                                  {
                                                                                                                                                                  Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(matchValue_1_0)
                                                                                                                                                                  =>
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_genericShow_prime(),
                                                                                                                                                                                                                                      &&&dictGenericShow1),
                                                                                                                                                                                                   &&matchValue_1_0),
                                                                                                                                                                  Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(matchValue_0_0)
                                                                                                                                                                  =>
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_genericShow_prime(),
                                                                                                                                                                                                                                      &&&dictGenericShow),
                                                                                                                                                                                                   &&matchValue_0_0),
                                                                                                                                                              }
                                                                                                                                                          }
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>()))
                                                                                     })))
    }
    pub fn Data_Show_Generic_genericShow() -> &dyn Any {
        static Data_Show_Generic_genericShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Generic_genericShow.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictGeneric|
                                                                      &Func1::new({
                                                                                      let dictGeneric
                                                                                          =
                                                                                          dictGeneric.clone();
                                                                                      move
                                                                                          |dictGenericShow|
                                                                                          &Func1::new({
                                                                                                          let dictGenericShow
                                                                                                              =
                                                                                                              dictGenericShow.clone();
                                                                                                          move
                                                                                                              |x|
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show_Generic::Data_Show_Generic_genericShow_prime(),
                                                                                                                                                                                  &&&dictGenericShow),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                     &&&dictGeneric),
                                                                                                                                                                                  x))
                                                                                                      })
                                                                                  })))
    }
}
