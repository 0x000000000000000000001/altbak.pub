pub mod PureScript_Data_Semiring_Generic {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_NoArguments;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Product;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Semiring_Generic_GenericSemiringusd_Dict() -> &dyn Any {
        static Data_Semiring_Generic_GenericSemiringusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_GenericSemiringusd_Dict.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |x|
                                                                                      x.clone()))
    }
    pub fn Data_Semiring_Generic_genericZero_prime() -> &dyn Any {
        static Data_Semiring_Generic_genericZero_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericZero_prime.get_or_init(||
                                                                &Func1::new(move
                                                                                |dict|
                                                                                find(string("genericZero\'"),
                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_Generic_genericZero() -> &dyn Any {
        static Data_Semiring_Generic_genericZero: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericZero.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictGeneric|
                                                                          &Func1::new({
                                                                                          let dictGeneric
                                                                                              =
                                                                                              dictGeneric.clone();
                                                                                          move
                                                                                              |dictGenericSemiring|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                                                                  &&&dictGeneric),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericZero_prime(),
                                                                                                                                                                  dictGenericSemiring))
                                                                                      })))
    }
    pub fn Data_Semiring_Generic_genericSemiringNoArguments() -> &dyn Any {
        static Data_Semiring_Generic_genericSemiringNoArguments:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericSemiringNoArguments.get_or_init(||
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_GenericSemiringusd_Dict(),
                                                                                                          &&&add(string("genericAdd\'"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |v|
                                                                                                                                  &Func1::new(move
                                                                                                                                                  |v1|
                                                                                                                                                  &LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor))),
                                                                                                                 add(string("genericZero\'"),
                                                                                                                     &&LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor),
                                                                                                                     add(string("genericMul\'"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |v_1|
                                                                                                                                          &Func1::new(move
                                                                                                                                                          |v1_1|
                                                                                                                                                          &LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor))),
                                                                                                                         add(string("genericOne\'"),
                                                                                                                             &&LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>()))))))
    }
    pub fn Data_Semiring_Generic_genericSemiringArgument() -> &dyn Any {
        static Data_Semiring_Generic_genericSemiringArgument:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericSemiringArgument.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictSemiring|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_GenericSemiringusd_Dict(),
                                                                                                                       &&&add(string("genericAdd\'"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let dictSemiring
                                                                                                                                                   =
                                                                                                                                                   dictSemiring.clone();
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
                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                     &&&dictSemiring),
                                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                                               &&&matchValue_1))
                                                                                                                                                                       }
                                                                                                                                                               })
                                                                                                                                           }),
                                                                                                                              add(string("genericZero\'"),
                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                                                                       dictSemiring)),
                                                                                                                                  add(string("genericMul\'"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let dictSemiring
                                                                                                                                                           =
                                                                                                                                                           dictSemiring.clone();
                                                                                                                                                       move
                                                                                                                                                           |v_1|
                                                                                                                                                           &Func1::new({
                                                                                                                                                                           let v_1
                                                                                                                                                                               =
                                                                                                                                                                               v_1.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |v1_1|
                                                                                                                                                                               {
                                                                                                                                                                                   let matchValue_3 =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                   let matchValue_4 =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                                                                                             &&&dictSemiring),
                                                                                                                                                                                                                                                                                          &&&matchValue_3),
                                                                                                                                                                                                                                                       &&&matchValue_4))
                                                                                                                                                                               }
                                                                                                                                                                       })
                                                                                                                                                   }),
                                                                                                                                      add(string("genericOne\'"),
                                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                                                               dictSemiring)),
                                                                                                                                          empty::<string,
                                                                                                                                                  &dyn Any>())))))))
    }
    pub fn Data_Semiring_Generic_genericOne_prime() -> &dyn Any {
        static Data_Semiring_Generic_genericOne_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericOne_prime.get_or_init(||
                                                               &Func1::new(move
                                                                               |dict|
                                                                               find(string("genericOne\'"),
                                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_Generic_genericOne() -> &dyn Any {
        static Data_Semiring_Generic_genericOne: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericOne.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictGeneric|
                                                                         &Func1::new({
                                                                                         let dictGeneric
                                                                                             =
                                                                                             dictGeneric.clone();
                                                                                         move
                                                                                             |dictGenericSemiring|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                                                                 &&&dictGeneric),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericOne_prime(),
                                                                                                                                                                 dictGenericSemiring))
                                                                                     })))
    }
    pub fn Data_Semiring_Generic_genericMul_prime() -> &dyn Any {
        static Data_Semiring_Generic_genericMul_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericMul_prime.get_or_init(||
                                                               &Func1::new(move
                                                                               |dict|
                                                                               find(string("genericMul\'"),
                                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_Generic_genericMul() -> &dyn Any {
        static Data_Semiring_Generic_genericMul: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericMul.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictGeneric|
                                                                         {
                                                                             let to_var =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                  dictGeneric);
                                                                             &Func1::new({
                                                                                             let dictGeneric
                                                                                                 =
                                                                                                 dictGeneric.clone();
                                                                                             let to_var
                                                                                                 =
                                                                                                 to_var.clone();
                                                                                             move
                                                                                                 |dictGenericSemiring|
                                                                                                 &Func1::new({
                                                                                                                 let dictGenericSemiring
                                                                                                                     =
                                                                                                                     dictGenericSemiring.clone();
                                                                                                                 move
                                                                                                                     |x|
                                                                                                                     &Func1::new({
                                                                                                                                     let x
                                                                                                                                         =
                                                                                                                                         x.clone();
                                                                                                                                     move
                                                                                                                                         |y|
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                             &&&to_var),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericMul_prime(),
                                                                                                                                                                                                                                                                                   &&&dictGenericSemiring),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                                                                      &&&dictGeneric),
                                                                                                                                                                                                                                                                                   &&&x)),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                                   &&&dictGeneric),
                                                                                                                                                                                                                                                y)))
                                                                                                                                 })
                                                                                                             })
                                                                                         })
                                                                         }))
    }
    pub fn Data_Semiring_Generic_genericAdd_prime() -> &dyn Any {
        static Data_Semiring_Generic_genericAdd_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericAdd_prime.get_or_init(||
                                                               &Func1::new(move
                                                                               |dict|
                                                                               find(string("genericAdd\'"),
                                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_Generic_genericSemiringConstructor() -> &dyn Any {
        static Data_Semiring_Generic_genericSemiringConstructor:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericSemiringConstructor.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |dictGenericSemiring|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_GenericSemiringusd_Dict(),
                                                                                                                          &&&add(string("genericAdd\'"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let dictGenericSemiring
                                                                                                                                                      =
                                                                                                                                                      dictGenericSemiring.clone();
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
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericAdd_prime(),
                                                                                                                                                                                                                                                                                                                        &&&dictGenericSemiring),
                                                                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                                                                  &&&matchValue_1))
                                                                                                                                                                          }
                                                                                                                                                                  })
                                                                                                                                              }),
                                                                                                                                 add(string("genericZero\'"),
                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericZero_prime(),
                                                                                                                                                                                                          dictGenericSemiring)),
                                                                                                                                     add(string("genericMul\'"),
                                                                                                                                         &&Func1::new({
                                                                                                                                                          let dictGenericSemiring
                                                                                                                                                              =
                                                                                                                                                              dictGenericSemiring.clone();
                                                                                                                                                          move
                                                                                                                                                              |v_1|
                                                                                                                                                              &Func1::new({
                                                                                                                                                                              let v_1
                                                                                                                                                                                  =
                                                                                                                                                                                  v_1.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |v1_1|
                                                                                                                                                                                  {
                                                                                                                                                                                      let matchValue_3 =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                      let matchValue_4 =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericMul_prime(),
                                                                                                                                                                                                                                                                                                                                &&&dictGenericSemiring),
                                                                                                                                                                                                                                                                                             &&&matchValue_3),
                                                                                                                                                                                                                                                          &&&matchValue_4))
                                                                                                                                                                                  }
                                                                                                                                                                          })
                                                                                                                                                      }),
                                                                                                                                         add(string("genericOne\'"),
                                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericOne_prime(),
                                                                                                                                                                                                                  dictGenericSemiring)),
                                                                                                                                             empty::<string,
                                                                                                                                                     &dyn Any>())))))))
    }
    pub fn Data_Semiring_Generic_genericSemiringProduct() -> &dyn Any {
        static Data_Semiring_Generic_genericSemiringProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericSemiringProduct.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictGenericSemiring|
                                                                                     &Func1::new({
                                                                                                     let dictGenericSemiring
                                                                                                         =
                                                                                                         dictGenericSemiring.clone();
                                                                                                     move
                                                                                                         |dictGenericSemiring1|
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_GenericSemiringusd_Dict(),
                                                                                                                                          &&&add(string("genericAdd\'"),
                                                                                                                                                 &&Func1::new({
                                                                                                                                                                  let dictGenericSemiring1
                                                                                                                                                                      =
                                                                                                                                                                      dictGenericSemiring1.clone();
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
                                                                                                                                                                                              &LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericAdd_prime(),
                                                                                                                                                                                                                                                                                                                                                                             &&&dictGenericSemiring),
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
                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericAdd_prime(),
                                                                                                                                                                                                                                                                                                                                                                             &&&dictGenericSemiring1),
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
                                                                                                                                                 add(string("genericZero\'"),
                                                                                                                                                     &&LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericZero_prime(),
                                                                                                                                                                                                                                                               &&&dictGenericSemiring),
                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericZero_prime(),
                                                                                                                                                                                                                                                               dictGenericSemiring1))),
                                                                                                                                                     add(string("genericMul\'"),
                                                                                                                                                         &&Func1::new({
                                                                                                                                                                          let dictGenericSemiring1
                                                                                                                                                                              =
                                                                                                                                                                              dictGenericSemiring1.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |v_1|
                                                                                                                                                                              &Func1::new({
                                                                                                                                                                                              let v_1
                                                                                                                                                                                                  =
                                                                                                                                                                                                  v_1.clone();
                                                                                                                                                                                              move
                                                                                                                                                                                                  |v1_1|
                                                                                                                                                                                                  {
                                                                                                                                                                                                      let matchValue_3:
                                                                                                                                                                                                              LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                                      let matchValue_4:
                                                                                                                                                                                                              LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                                                          Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                                      &LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericMul_prime(),
                                                                                                                                                                                                                                                                                                                                                                                     &&&dictGenericSemiring),
                                                                                                                                                                                                                                                                                                                                                  &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                         Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                                                                                                                               &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                      Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericMul_prime(),
                                                                                                                                                                                                                                                                                                                                                                                     &&&dictGenericSemiring1),
                                                                                                                                                                                                                                                                                                                                                  &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                         Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                                                                                                                               &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                      Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                  })))
                                                                                                                                                                                                  }
                                                                                                                                                                                          })
                                                                                                                                                                      }),
                                                                                                                                                         add(string("genericOne\'"),
                                                                                                                                                             &&LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericOne_prime(),
                                                                                                                                                                                                                                                                       &&&dictGenericSemiring),
                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericOne_prime(),
                                                                                                                                                                                                                                                                       dictGenericSemiring1))),
                                                                                                                                                             empty::<string,
                                                                                                                                                                     &dyn Any>())))))
                                                                                                 })))
    }
    pub fn Data_Semiring_Generic_genericAdd() -> &dyn Any {
        static Data_Semiring_Generic_genericAdd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Generic_genericAdd.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictGeneric|
                                                                         {
                                                                             let to_var =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                  dictGeneric);
                                                                             &Func1::new({
                                                                                             let dictGeneric
                                                                                                 =
                                                                                                 dictGeneric.clone();
                                                                                             let to_var
                                                                                                 =
                                                                                                 to_var.clone();
                                                                                             move
                                                                                                 |dictGenericSemiring|
                                                                                                 &Func1::new({
                                                                                                                 let dictGenericSemiring
                                                                                                                     =
                                                                                                                     dictGenericSemiring.clone();
                                                                                                                 move
                                                                                                                     |x|
                                                                                                                     &Func1::new({
                                                                                                                                     let x
                                                                                                                                         =
                                                                                                                                         x.clone();
                                                                                                                                     move
                                                                                                                                         |y|
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                             &&&to_var),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring_Generic::Data_Semiring_Generic_genericAdd_prime(),
                                                                                                                                                                                                                                                                                   &&&dictGenericSemiring),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                                                                      &&&dictGeneric),
                                                                                                                                                                                                                                                                                   &&&x)),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                                   &&&dictGeneric),
                                                                                                                                                                                                                                                y)))
                                                                                                                                 })
                                                                                                             })
                                                                                         })
                                                                         }))
    }
}
