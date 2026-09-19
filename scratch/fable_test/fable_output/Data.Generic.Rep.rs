pub mod PureScript_Data_Generic_Rep {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    #[derive(Clone, Debug,)]
    pub enum Data_Generic_Rep_Sum {
        Data_Generic_Rep_Inlusd_Ctor(&dyn Any),
        Data_Generic_Rep_Inrusd_Ctor(&dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug,)]
    pub enum Data_Generic_Rep_Product {
        Data_Generic_Rep_Productusd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Generic_Rep::Data_Generic_Rep_Product {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
    pub enum Data_Generic_Rep_NoArguments {
        Data_Generic_Rep_NoArgumentsusd_Ctor,
    }
    impl core::fmt::Display for
     PureScript_Data_Generic_Rep::Data_Generic_Rep_NoArguments {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Generic_Rep_Inl() -> &dyn Any {
        static Data_Generic_Rep_Inl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_Inl.get_or_init(||
                                             &Func1::new(move |usd__arg1|
                                                             &LrcPtr::new(PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Generic_Rep_Inr() -> &dyn Any {
        static Data_Generic_Rep_Inr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_Inr.get_or_init(||
                                             &Func1::new(move |usd__arg1|
                                                             &LrcPtr::new(PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Generic_Rep_Product() -> &dyn Any {
        static Data_Generic_Rep_Product: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_Product.get_or_init(||
                                                 &Func1::new(move |usd__arg1|
                                                                 Func1::new({
                                                                                let usd__arg1
                                                                                    =
                                                                                    usd__arg1.clone();
                                                                                move
                                                                                    |usd__arg2|
                                                                                    &LrcPtr::new(PureScript_Data_Generic_Rep::Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(usd__arg1,
                                                                                                                                                                                         usd__arg2.clone()))
                                                                            })))
    }
    pub fn Data_Generic_Rep_NoConstructors() -> &dyn Any {
        static Data_Generic_Rep_NoConstructors: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_NoConstructors.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        x.clone()))
    }
    pub fn Data_Generic_Rep_NoArguments() -> &dyn Any {
        static Data_Generic_Rep_NoArguments: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_NoArguments.get_or_init(||
                                                     &LrcPtr::new(PureScript_Data_Generic_Rep::Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor))
    }
    pub fn Data_Generic_Rep_Genericusd_Dict() -> &dyn Any {
        static Data_Generic_Rep_Genericusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_Genericusd_Dict.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Data_Generic_Rep_Constructor() -> &dyn Any {
        static Data_Generic_Rep_Constructor: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_Constructor.get_or_init(||
                                                     &Func1::new(move |x|
                                                                     x.clone()))
    }
    pub fn Data_Generic_Rep_Argument() -> &dyn Any {
        static Data_Generic_Rep_Argument: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_Argument.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  x.clone()))
    }
    pub fn Data_Generic_Rep_to() -> &dyn Any {
        static Data_Generic_Rep_to: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_to.get_or_init(||
                                            &Func1::new(move |dict|
                                                            find(string("to"),
                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Generic_Rep_showSum() -> &dyn Any {
        static Data_Generic_Rep_showSum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_showSum.get_or_init(||
                                                 &Func1::new(move |dictShow|
                                                                 &Func1::new({
                                                                                 let dictShow
                                                                                     =
                                                                                     dictShow.clone();
                                                                                 move
                                                                                     |dictShow1|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                      &&&add(string("show"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let dictShow1
                                                                                                                                                  =
                                                                                                                                                  dictShow1.clone();
                                                                                                                                              move
                                                                                                                                                  |v|
                                                                                                                                                  {
                                                                                                                                                      let matchValue:
                                                                                                                                                              LrcPtr<PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum> =
                                                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                                                      match matchValue.as_ref()
                                                                                                                                                          {
                                                                                                                                                          PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(matchValue_1_0)
                                                                                                                                                          =>
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                              &&&string("(Inr ")),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                       &&&dictShow1),
                                                                                                                                                                                                                                                                                                    &&matchValue_1_0)),
                                                                                                                                                                                                                              &&&string(")"))),
                                                                                                                                                          PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(matchValue_0_0)
                                                                                                                                                          =>
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                              &&&string("(Inl ")),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                       &&&dictShow),
                                                                                                                                                                                                                                                                                                    &&matchValue_0_0)),
                                                                                                                                                                                                                              &&&string(")"))),
                                                                                                                                                      }
                                                                                                                                                  }
                                                                                                                                          }),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>()))
                                                                             })))
    }
    pub fn Data_Generic_Rep_showProduct() -> &dyn Any {
        static Data_Generic_Rep_showProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_showProduct.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictShow|
                                                                     &Func1::new({
                                                                                     let dictShow
                                                                                         =
                                                                                         dictShow.clone();
                                                                                     move
                                                                                         |dictShow1|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                          &&&add(string("show"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let dictShow1
                                                                                                                                                      =
                                                                                                                                                      dictShow1.clone();
                                                                                                                                                  move
                                                                                                                                                      |v|
                                                                                                                                                      {
                                                                                                                                                          let matchValue:
                                                                                                                                                                  LrcPtr<PureScript_Data_Generic_Rep::Data_Generic_Rep_Product> =
                                                                                                                                                              Sharpurs_Prelude::unbox(v);
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                              &&&string("(Product ")),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                       &&&dictShow),
                                                                                                                                                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                           PureScript_Data_Generic_Rep::Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                       })),
                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                    &&&string(" ")),
                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                             &&&dictShow1),
                                                                                                                                                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                                                                                 PureScript_Data_Generic_Rep::Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                                                                                             })),
                                                                                                                                                                                                                                                                                                    &&&string(")")))))
                                                                                                                                                      }
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>()))
                                                                                 })))
    }
    pub fn Data_Generic_Rep_showNoArguments() -> &dyn Any {
        static Data_Generic_Rep_showNoArguments: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_showNoArguments.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                          &&&add(string("show"),
                                                                                                 &&Func1::new(move
                                                                                                                  |v|
                                                                                                                  &string("NoArguments")),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>())))
    }
    pub fn Data_Generic_Rep_showConstructor() -> &dyn Any {
        static Data_Generic_Rep_showConstructor: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_showConstructor.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictIsSymbol|
                                                                         &Func1::new({
                                                                                         let dictIsSymbol
                                                                                             =
                                                                                             dictIsSymbol.clone();
                                                                                         move
                                                                                             |dictShow|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                              &&&add(string("show"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let dictShow
                                                                                                                                                          =
                                                                                                                                                          dictShow.clone();
                                                                                                                                                      move
                                                                                                                                                          |v|
                                                                                                                                                          {
                                                                                                                                                              let a =
                                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                     &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                  &&&string("(Constructor @")),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Show::Data_Show_showString()),
                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                                                                              &&&dictIsSymbol),
                                                                                                                                                                                                                                                                                                                                           &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)))),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                        &&&string(" ")),
                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                                 &&&dictShow),
                                                                                                                                                                                                                                                                                                                                                                              &&&a)),
                                                                                                                                                                                                                                                                                                        &&&string(")")))))
                                                                                                                                                          }
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>()))
                                                                                     })))
    }
    pub fn Data_Generic_Rep_showArgument() -> &dyn Any {
        static Data_Generic_Rep_showArgument: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_showArgument.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictShow|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                       &&&add(string("show"),
                                                                                                              &&Func1::new({
                                                                                                                               let dictShow
                                                                                                                                   =
                                                                                                                                   dictShow.clone();
                                                                                                                               move
                                                                                                                                   |v|
                                                                                                                                   {
                                                                                                                                       let a =
                                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                           &&&string("(Argument ")),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                    &&&dictShow),
                                                                                                                                                                                                                                                                                 &&&a)),
                                                                                                                                                                                                           &&&string(")")))
                                                                                                                                   }
                                                                                                                           }),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Generic_Rep_repOf() -> &dyn Any {
        static Data_Generic_Rep_repOf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_repOf.get_or_init(||
                                               &Func1::new(move |dictGeneric|
                                                               &Func1::new(move
                                                                               |v|
                                                                               &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))))
    }
    pub fn Data_Generic_Rep_from() -> &dyn Any {
        static Data_Generic_Rep_from: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Generic_Rep_from.get_or_init(||
                                              &Func1::new(move |dict|
                                                              find(string("from"),
                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}
