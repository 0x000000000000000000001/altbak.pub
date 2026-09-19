pub mod PureScript_Data_Monoid_Generic {
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
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_NoArguments;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Product;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Monoid_Generic_GenericMonoidusd_Dict() -> &dyn Any {
        static Data_Monoid_Generic_GenericMonoidusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Generic_GenericMonoidusd_Dict.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |x|
                                                                                  x.clone()))
    }
    pub fn Data_Monoid_Generic_genericMonoidNoArguments() -> &dyn Any {
        static Data_Monoid_Generic_genericMonoidNoArguments:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Generic_genericMonoidNoArguments.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Generic::Data_Monoid_Generic_GenericMonoidusd_Dict(),
                                                                                                      &&&add(string("genericMempty\'"),
                                                                                                             &&LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
    }
    pub fn Data_Monoid_Generic_genericMonoidArgument() -> &dyn Any {
        static Data_Monoid_Generic_genericMonoidArgument:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Generic_genericMonoidArgument.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonoid|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Generic::Data_Monoid_Generic_GenericMonoidusd_Dict(),
                                                                                                                   &&&add(string("genericMempty\'"),
                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                               dictMonoid)),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Monoid_Generic_genericMempty_prime() -> &dyn Any {
        static Data_Monoid_Generic_genericMempty_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Generic_genericMempty_prime.get_or_init(||
                                                                &Func1::new(move
                                                                                |dict|
                                                                                find(string("genericMempty\'"),
                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Monoid_Generic_genericMonoidConstructor() -> &dyn Any {
        static Data_Monoid_Generic_genericMonoidConstructor:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Generic_genericMonoidConstructor.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictGenericMonoid|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Generic::Data_Monoid_Generic_GenericMonoidusd_Dict(),
                                                                                                                      &&&add(string("genericMempty\'"),
                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Generic::Data_Monoid_Generic_genericMempty_prime(),
                                                                                                                                                                                                  dictGenericMonoid)),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>()))))
    }
    pub fn Data_Monoid_Generic_genericMonoidProduct() -> &dyn Any {
        static Data_Monoid_Generic_genericMonoidProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Generic_genericMonoidProduct.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictGenericMonoid|
                                                                                 &Func1::new({
                                                                                                 let dictGenericMonoid
                                                                                                     =
                                                                                                     dictGenericMonoid.clone();
                                                                                                 move
                                                                                                     |dictGenericMonoid1|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Generic::Data_Monoid_Generic_GenericMonoidusd_Dict(),
                                                                                                                                      &&&add(string("genericMempty\'"),
                                                                                                                                             &&LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Generic::Data_Monoid_Generic_genericMempty_prime(),
                                                                                                                                                                                                                                                       &&&dictGenericMonoid),
                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Generic::Data_Monoid_Generic_genericMempty_prime(),
                                                                                                                                                                                                                                                       dictGenericMonoid1))),
                                                                                                                                             empty::<string,
                                                                                                                                                     &dyn Any>()))
                                                                                             })))
    }
    pub fn Data_Monoid_Generic_genericMempty() -> &dyn Any {
        static Data_Monoid_Generic_genericMempty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Generic_genericMempty.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictGeneric|
                                                                          &Func1::new({
                                                                                          let dictGeneric
                                                                                              =
                                                                                              dictGeneric.clone();
                                                                                          move
                                                                                              |dictGenericMonoid|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                                                                  &&&dictGeneric),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Generic::Data_Monoid_Generic_genericMempty_prime(),
                                                                                                                                                                  dictGenericMonoid))
                                                                                      })))
    }
}
