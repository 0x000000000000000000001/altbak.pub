pub mod PureScript_Data_String_Gen {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_851afbc9::PureScript_Control_Monad_Gen_Class;
    use crate::module_e72de349::PureScript_Control_Monad_Gen;
    use crate::module_5ec638cf::PureScript_Data_Char_Gen;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_eea316d6::PureScript_Data_String_CodeUnits;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_String_Gen_genString() -> &dyn Any {
        static Data_String_Gen_genString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Gen_genString.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictMonadRec|
                                                                  &Func1::new({
                                                                                  let dictMonadRec
                                                                                      =
                                                                                      dictMonadRec.clone();
                                                                                  move
                                                                                      |dictMonadGen|
                                                                                      {
                                                                                          let Monad0 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                      Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                                          let Bind1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                      Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                                          let Functor0 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                                          &Func1::new({
                                                                                                          let Bind1
                                                                                                              =
                                                                                                              Bind1.clone();
                                                                                                          let Functor0
                                                                                                              =
                                                                                                              Functor0.clone();
                                                                                                          let dictMonadGen
                                                                                                              =
                                                                                                              dictMonadGen.clone();
                                                                                                          move
                                                                                                              |genChar|
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_sized(),
                                                                                                                                                                                  &&&dictMonadGen),
                                                                                                                                               &&&Func1::new({
                                                                                                                                                                 let genChar
                                                                                                                                                                     =
                                                                                                                                                                     genChar.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |size|
                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                            &&&Bind1),
                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_chooseInt(),
                                                                                                                                                                                                                                                                                                                                                  &&&dictMonadGen),
                                                                                                                                                                                                                                                                                                               &&&1_i32),
                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_max(),
                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                                                                                                                  &&&1_i32),
                                                                                                                                                                                                                                                                                                               size))),
                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                        |newSize|
                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_resize(),
                                                                                                                                                                                                                                                                                                                                                                  &&&dictMonadGen),
                                                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                                                                                                                                                  newSize))),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                  &&&Functor0),
                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_fromCharArray()),
                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen::Control_Monad_Gen_unfoldable(),
                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&dictMonadRec),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&dictMonadGen),
                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldableArray()),
                                                                                                                                                                                                                                                                                                                               &&&genChar)))))
                                                                                                                                                             }))
                                                                                                      })
                                                                                      }
                                                                              })))
    }
    pub fn Data_String_Gen_genUnicodeString() -> &dyn Any {
        static Data_String_Gen_genUnicodeString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Gen_genUnicodeString.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonadRec|
                                                                         &Func1::new({
                                                                                         let dictMonadRec
                                                                                             =
                                                                                             dictMonadRec.clone();
                                                                                         move
                                                                                             |dictMonadGen|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Gen::Data_String_Gen_genString(),
                                                                                                                                                                                                    &&&dictMonadRec),
                                                                                                                                                                 dictMonadGen),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Char_Gen::Data_Char_Gen_genUnicodeChar(),
                                                                                                                                                                 dictMonadGen))
                                                                                     })))
    }
    pub fn Data_String_Gen_genDigitString() -> &dyn Any {
        static Data_String_Gen_genDigitString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Gen_genDigitString.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadRec|
                                                                       &Func1::new({
                                                                                       let dictMonadRec
                                                                                           =
                                                                                           dictMonadRec.clone();
                                                                                       move
                                                                                           |dictMonadGen|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Gen::Data_String_Gen_genString(),
                                                                                                                                                                                                  &&&dictMonadRec),
                                                                                                                                                               dictMonadGen),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Char_Gen::Data_Char_Gen_genDigitChar(),
                                                                                                                                                               dictMonadGen))
                                                                                   })))
    }
    pub fn Data_String_Gen_genAsciiString_prime() -> &dyn Any {
        static Data_String_Gen_genAsciiString_prime: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_Gen_genAsciiString_prime.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictMonadRec|
                                                                             &Func1::new({
                                                                                             let dictMonadRec
                                                                                                 =
                                                                                                 dictMonadRec.clone();
                                                                                             move
                                                                                                 |dictMonadGen|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Gen::Data_String_Gen_genString(),
                                                                                                                                                                                                        &&&dictMonadRec),
                                                                                                                                                                     dictMonadGen),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Char_Gen::Data_Char_Gen_genAsciiChar_prime(),
                                                                                                                                                                     dictMonadGen))
                                                                                         })))
    }
    pub fn Data_String_Gen_genAsciiString() -> &dyn Any {
        static Data_String_Gen_genAsciiString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Gen_genAsciiString.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadRec|
                                                                       &Func1::new({
                                                                                       let dictMonadRec
                                                                                           =
                                                                                           dictMonadRec.clone();
                                                                                       move
                                                                                           |dictMonadGen|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Gen::Data_String_Gen_genString(),
                                                                                                                                                                                                  &&&dictMonadRec),
                                                                                                                                                               dictMonadGen),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Char_Gen::Data_Char_Gen_genAsciiChar(),
                                                                                                                                                               dictMonadGen))
                                                                                   })))
    }
    pub fn Data_String_Gen_genAlphaUppercaseString() -> &dyn Any {
        static Data_String_Gen_genAlphaUppercaseString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Gen_genAlphaUppercaseString.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictMonadRec|
                                                                                &Func1::new({
                                                                                                let dictMonadRec
                                                                                                    =
                                                                                                    dictMonadRec.clone();
                                                                                                move
                                                                                                    |dictMonadGen|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Gen::Data_String_Gen_genString(),
                                                                                                                                                                                                           &&&dictMonadRec),
                                                                                                                                                                        dictMonadGen),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Char_Gen::Data_Char_Gen_genAlphaUppercase(),
                                                                                                                                                                        dictMonadGen))
                                                                                            })))
    }
    pub fn Data_String_Gen_genAlphaString() -> &dyn Any {
        static Data_String_Gen_genAlphaString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Gen_genAlphaString.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadRec|
                                                                       &Func1::new({
                                                                                       let dictMonadRec
                                                                                           =
                                                                                           dictMonadRec.clone();
                                                                                       move
                                                                                           |dictMonadGen|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Gen::Data_String_Gen_genString(),
                                                                                                                                                                                                  &&&dictMonadRec),
                                                                                                                                                               dictMonadGen),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Char_Gen::Data_Char_Gen_genAlpha(),
                                                                                                                                                               dictMonadGen))
                                                                                   })))
    }
    pub fn Data_String_Gen_genAlphaLowercaseString() -> &dyn Any {
        static Data_String_Gen_genAlphaLowercaseString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Gen_genAlphaLowercaseString.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictMonadRec|
                                                                                &Func1::new({
                                                                                                let dictMonadRec
                                                                                                    =
                                                                                                    dictMonadRec.clone();
                                                                                                move
                                                                                                    |dictMonadGen|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Gen::Data_String_Gen_genString(),
                                                                                                                                                                                                           &&&dictMonadRec),
                                                                                                                                                                        dictMonadGen),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Char_Gen::Data_Char_Gen_genAlphaLowercase(),
                                                                                                                                                                        dictMonadGen))
                                                                                            })))
    }
}
