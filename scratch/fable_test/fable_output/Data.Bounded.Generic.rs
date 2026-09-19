pub mod PureScript_Data_Bounded_Generic {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_NoArguments;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Product;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Bounded_Generic_GenericTopusd_Dict() -> &dyn Any {
        static Data_Bounded_Generic_GenericTopusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_GenericTopusd_Dict.get_or_init(||
                                                                &Func1::new(move
                                                                                |x|
                                                                                x.clone()))
    }
    pub fn Data_Bounded_Generic_GenericBottomusd_Dict() -> &dyn Any {
        static Data_Bounded_Generic_GenericBottomusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_GenericBottomusd_Dict.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |x|
                                                                                   x.clone()))
    }
    pub fn Data_Bounded_Generic_genericTopNoArguments() -> &dyn Any {
        static Data_Bounded_Generic_genericTopNoArguments:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericTopNoArguments.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericTopusd_Dict(),
                                                                                                    &&&add(string("genericTop\'"),
                                                                                                           &&LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))
    }
    pub fn Data_Bounded_Generic_genericTopArgument() -> &dyn Any {
        static Data_Bounded_Generic_genericTopArgument:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericTopArgument.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictBounded|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericTopusd_Dict(),
                                                                                                                 &&&add(string("genericTop\'"),
                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                             dictBounded)),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))))
    }
    pub fn Data_Bounded_Generic_genericTop_prime() -> &dyn Any {
        static Data_Bounded_Generic_genericTop_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericTop_prime.get_or_init(||
                                                              &Func1::new(move
                                                                              |dict|
                                                                              find(string("genericTop\'"),
                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bounded_Generic_genericTopConstructor() -> &dyn Any {
        static Data_Bounded_Generic_genericTopConstructor:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericTopConstructor.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictGenericTop|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericTopusd_Dict(),
                                                                                                                    &&&add(string("genericTop\'"),
                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericTop_prime(),
                                                                                                                                                                                                dictGenericTop)),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))))
    }
    pub fn Data_Bounded_Generic_genericTopProduct() -> &dyn Any {
        static Data_Bounded_Generic_genericTopProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericTopProduct.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictGenericTop|
                                                                               &Func1::new({
                                                                                               let dictGenericTop
                                                                                                   =
                                                                                                   dictGenericTop.clone();
                                                                                               move
                                                                                                   |dictGenericTop1|
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericTopusd_Dict(),
                                                                                                                                    &&&add(string("genericTop\'"),
                                                                                                                                           &&LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericTop_prime(),
                                                                                                                                                                                                                                                     &&&dictGenericTop),
                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericTop_prime(),
                                                                                                                                                                                                                                                     dictGenericTop1))),
                                                                                                                                           empty::<string,
                                                                                                                                                   &dyn Any>()))
                                                                                           })))
    }
    pub fn Data_Bounded_Generic_genericTopSum() -> &dyn Any {
        static Data_Bounded_Generic_genericTopSum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericTopSum.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictGenericTop|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericTopusd_Dict(),
                                                                                                            &&&add(string("genericTop\'"),
                                                                                                                   &&LrcPtr::new(Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericTop_prime(),
                                                                                                                                                                                                                     dictGenericTop))),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Bounded_Generic_genericTop() -> &dyn Any {
        static Data_Bounded_Generic_genericTop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericTop.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictGeneric|
                                                                        &Func1::new({
                                                                                        let dictGeneric
                                                                                            =
                                                                                            dictGeneric.clone();
                                                                                        move
                                                                                            |dictGenericTop|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                                                                &&&dictGeneric),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericTop_prime(),
                                                                                                                                                                dictGenericTop))
                                                                                    })))
    }
    pub fn Data_Bounded_Generic_genericBottomNoArguments() -> &dyn Any {
        static Data_Bounded_Generic_genericBottomNoArguments:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericBottomNoArguments.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericBottomusd_Dict(),
                                                                                                       &&&add(string("genericBottom\'"),
                                                                                                              &&LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))
    }
    pub fn Data_Bounded_Generic_genericBottomArgument() -> &dyn Any {
        static Data_Bounded_Generic_genericBottomArgument:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericBottomArgument.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictBounded|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericBottomusd_Dict(),
                                                                                                                    &&&add(string("genericBottom\'"),
                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                dictBounded)),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))))
    }
    pub fn Data_Bounded_Generic_genericBottom_prime() -> &dyn Any {
        static Data_Bounded_Generic_genericBottom_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericBottom_prime.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dict|
                                                                                 find(string("genericBottom\'"),
                                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bounded_Generic_genericBottomConstructor() -> &dyn Any {
        static Data_Bounded_Generic_genericBottomConstructor:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericBottomConstructor.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictGenericBottom|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericBottomusd_Dict(),
                                                                                                                       &&&add(string("genericBottom\'"),
                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericBottom_prime(),
                                                                                                                                                                                                   dictGenericBottom)),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Bounded_Generic_genericBottomProduct() -> &dyn Any {
        static Data_Bounded_Generic_genericBottomProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericBottomProduct.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictGenericBottom|
                                                                                  &Func1::new({
                                                                                                  let dictGenericBottom
                                                                                                      =
                                                                                                      dictGenericBottom.clone();
                                                                                                  move
                                                                                                      |dictGenericBottom1|
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericBottomusd_Dict(),
                                                                                                                                       &&&add(string("genericBottom\'"),
                                                                                                                                              &&LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericBottom_prime(),
                                                                                                                                                                                                                                                        &&&dictGenericBottom),
                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericBottom_prime(),
                                                                                                                                                                                                                                                        dictGenericBottom1))),
                                                                                                                                              empty::<string,
                                                                                                                                                      &dyn Any>()))
                                                                                              })))
    }
    pub fn Data_Bounded_Generic_genericBottomSum() -> &dyn Any {
        static Data_Bounded_Generic_genericBottomSum:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericBottomSum.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictGenericBottom|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_GenericBottomusd_Dict(),
                                                                                                               &&&add(string("genericBottom\'"),
                                                                                                                      &&LrcPtr::new(Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericBottom_prime(),
                                                                                                                                                                                                                        dictGenericBottom))),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))))
    }
    pub fn Data_Bounded_Generic_genericBottom() -> &dyn Any {
        static Data_Bounded_Generic_genericBottom: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Generic_genericBottom.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictGeneric|
                                                                           &Func1::new({
                                                                                           let dictGeneric
                                                                                               =
                                                                                               dictGeneric.clone();
                                                                                           move
                                                                                               |dictGenericBottom|
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                                                                   &&&dictGeneric),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded_Generic::Data_Bounded_Generic_genericBottom_prime(),
                                                                                                                                                                   dictGenericBottom))
                                                                                       })))
    }
}
