pub mod PureScript_Data_HeytingAlgebra_Generic {
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
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_HeytingAlgebra_Generic_GenericHeytingAlgebrausd_Dict()
     -> &dyn Any {
        static Data_HeytingAlgebra_Generic_GenericHeytingAlgebrausd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_GenericHeytingAlgebrausd_Dict.get_or_init(||
                                                                                  &Func1::new(move
                                                                                                  |x|
                                                                                                  x.clone()))
    }
    pub fn Data_HeytingAlgebra_Generic_genericTT_prime() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericTT_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericTT_prime.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dict|
                                                                                    find(string("genericTT\'"),
                                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_Generic_genericTT() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericTT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericTT.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictGeneric|
                                                                              &Func1::new({
                                                                                              let dictGeneric
                                                                                                  =
                                                                                                  dictGeneric.clone();
                                                                                              move
                                                                                                  |dictGenericHeytingAlgebra|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                                                                      &&&dictGeneric),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericTT_prime(),
                                                                                                                                                                      dictGenericHeytingAlgebra))
                                                                                          })))
    }
    pub fn Data_HeytingAlgebra_Generic_genericNot_prime() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericNot_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericNot_prime.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dict|
                                                                                     find(string("genericNot\'"),
                                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_Generic_genericNot() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericNot:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericNot.get_or_init(||
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
                                                                                                       |dictGenericHeytingAlgebra|
                                                                                                       &Func1::new({
                                                                                                                       let dictGenericHeytingAlgebra
                                                                                                                           =
                                                                                                                           dictGenericHeytingAlgebra.clone();
                                                                                                                       move
                                                                                                                           |x|
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                               &&&to_var),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericNot_prime(),
                                                                                                                                                                                                                                  &&&dictGenericHeytingAlgebra),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_from(),
                                                                                                                                                                                                                                                                     &&&dictGeneric),
                                                                                                                                                                                                                                  x)))
                                                                                                                   })
                                                                                               })
                                                                               }))
    }
    pub fn Data_HeytingAlgebra_Generic_genericImplies_prime() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericImplies_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericImplies_prime.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |dict|
                                                                                         find(string("genericImplies\'"),
                                                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_Generic_genericImplies() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericImplies:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericImplies.get_or_init(||
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
                                                                                                           |dictGenericHeytingAlgebra|
                                                                                                           &Func1::new({
                                                                                                                           let dictGenericHeytingAlgebra
                                                                                                                               =
                                                                                                                               dictGenericHeytingAlgebra.clone();
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
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericImplies_prime(),
                                                                                                                                                                                                                                                                                             &&&dictGenericHeytingAlgebra),
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
    pub fn Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments()
     -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments.get_or_init(||
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_GenericHeytingAlgebrausd_Dict(),
                                                                                                                      &&&add(string("genericFF\'"),
                                                                                                                             &&LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor),
                                                                                                                             add(string("genericTT\'"),
                                                                                                                                 &&LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor),
                                                                                                                                 add(string("genericImplies\'"),
                                                                                                                                     &&Func1::new(move
                                                                                                                                                      |v|
                                                                                                                                                      &Func1::new(move
                                                                                                                                                                      |v1|
                                                                                                                                                                      &LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor))),
                                                                                                                                     add(string("genericConj\'"),
                                                                                                                                         &&Func1::new(move
                                                                                                                                                          |v_1|
                                                                                                                                                          &Func1::new(move
                                                                                                                                                                          |v1_1|
                                                                                                                                                                          &LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor))),
                                                                                                                                         add(string("genericDisj\'"),
                                                                                                                                             &&Func1::new(move
                                                                                                                                                              |v_2|
                                                                                                                                                              &Func1::new(move
                                                                                                                                                                              |v1_2|
                                                                                                                                                                              &LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor))),
                                                                                                                                             add(string("genericNot\'"),
                                                                                                                                                 &&Func1::new(move
                                                                                                                                                                  |v_3|
                                                                                                                                                                  &LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor)),
                                                                                                                                                 empty::<string,
                                                                                                                                                         &dyn Any>()))))))))
    }
    pub fn Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument()
     -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument.get_or_init(||
                                                                                  &Func1::new(move
                                                                                                  |dictHeytingAlgebra|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_GenericHeytingAlgebrausd_Dict(),
                                                                                                                                   &&&add(string("genericFF\'"),
                                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ff(),
                                                                                                                                                                                                               dictHeytingAlgebra)),
                                                                                                                                          add(string("genericTT\'"),
                                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_tt(),
                                                                                                                                                                                                                   dictHeytingAlgebra)),
                                                                                                                                              add(string("genericImplies\'"),
                                                                                                                                                  &&Func1::new({
                                                                                                                                                                   let dictHeytingAlgebra
                                                                                                                                                                       =
                                                                                                                                                                       dictHeytingAlgebra.clone();
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
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_implies(),
                                                                                                                                                                                                                                                                                                                                         &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                      &&&matchValue),
                                                                                                                                                                                                                                                                   &&&matchValue_1))
                                                                                                                                                                                           }
                                                                                                                                                                                   })
                                                                                                                                                               }),
                                                                                                                                                  add(string("genericConj\'"),
                                                                                                                                                      &&Func1::new({
                                                                                                                                                                       let dictHeytingAlgebra
                                                                                                                                                                           =
                                                                                                                                                                           dictHeytingAlgebra.clone();
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
                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                                                             &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                          &&&matchValue_3),
                                                                                                                                                                                                                                                                       &&&matchValue_4))
                                                                                                                                                                                               }
                                                                                                                                                                                       })
                                                                                                                                                                   }),
                                                                                                                                                      add(string("genericDisj\'"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let dictHeytingAlgebra
                                                                                                                                                                               =
                                                                                                                                                                               dictHeytingAlgebra.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |v_2|
                                                                                                                                                                               &Func1::new({
                                                                                                                                                                                               let v_2
                                                                                                                                                                                                   =
                                                                                                                                                                                                   v_2.clone();
                                                                                                                                                                                               move
                                                                                                                                                                                                   |v1_2|
                                                                                                                                                                                                   {
                                                                                                                                                                                                       let matchValue_6 =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                                                       let matchValue_7 =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(v1_2);
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                                                                                                                                                                                                                                                                 &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                              &&&matchValue_6),
                                                                                                                                                                                                                                                                           &&&matchValue_7))
                                                                                                                                                                                                   }
                                                                                                                                                                                           })
                                                                                                                                                                       }),
                                                                                                                                                          add(string("genericNot\'"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let dictHeytingAlgebra
                                                                                                                                                                                   =
                                                                                                                                                                                   dictHeytingAlgebra.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |v_3|
                                                                                                                                                                                   {
                                                                                                                                                                                       let x_3 =
                                                                                                                                                                                           Sharpurs_Prelude::unbox(v_3);
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                                                                              &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                           &&&x_3))
                                                                                                                                                                                   }
                                                                                                                                                                           }),
                                                                                                                                                              empty::<string,
                                                                                                                                                                      &dyn Any>())))))))))
    }
    pub fn Data_HeytingAlgebra_Generic_genericFF_prime() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericFF_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericFF_prime.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dict|
                                                                                    find(string("genericFF\'"),
                                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_Generic_genericFF() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericFF:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericFF.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictGeneric|
                                                                              &Func1::new({
                                                                                              let dictGeneric
                                                                                                  =
                                                                                                  dictGeneric.clone();
                                                                                              move
                                                                                                  |dictGenericHeytingAlgebra|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_to(),
                                                                                                                                                                      &&&dictGeneric),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericFF_prime(),
                                                                                                                                                                      dictGenericHeytingAlgebra))
                                                                                          })))
    }
    pub fn Data_HeytingAlgebra_Generic_genericDisj_prime() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericDisj_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericDisj_prime.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dict|
                                                                                      find(string("genericDisj\'"),
                                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_Generic_genericDisj() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericDisj:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericDisj.get_or_init(||
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
                                                                                                        |dictGenericHeytingAlgebra|
                                                                                                        &Func1::new({
                                                                                                                        let dictGenericHeytingAlgebra
                                                                                                                            =
                                                                                                                            dictGenericHeytingAlgebra.clone();
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
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericDisj_prime(),
                                                                                                                                                                                                                                                                                          &&&dictGenericHeytingAlgebra),
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
    pub fn Data_HeytingAlgebra_Generic_genericConj_prime() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericConj_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericConj_prime.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dict|
                                                                                      find(string("genericConj\'"),
                                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor()
     -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor.get_or_init(||
                                                                                     &Func1::new(move
                                                                                                     |dictGenericHeytingAlgebra|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_GenericHeytingAlgebrausd_Dict(),
                                                                                                                                      &&&add(string("genericFF\'"),
                                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericFF_prime(),
                                                                                                                                                                                                                  dictGenericHeytingAlgebra)),
                                                                                                                                             add(string("genericTT\'"),
                                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericTT_prime(),
                                                                                                                                                                                                                      dictGenericHeytingAlgebra)),
                                                                                                                                                 add(string("genericImplies\'"),
                                                                                                                                                     &&Func1::new({
                                                                                                                                                                      let dictGenericHeytingAlgebra
                                                                                                                                                                          =
                                                                                                                                                                          dictGenericHeytingAlgebra.clone();
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
                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericImplies_prime(),
                                                                                                                                                                                                                                                                                                                                            &&&dictGenericHeytingAlgebra),
                                                                                                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                                                                                                      &&&matchValue_1))
                                                                                                                                                                                              }
                                                                                                                                                                                      })
                                                                                                                                                                  }),
                                                                                                                                                     add(string("genericConj\'"),
                                                                                                                                                         &&Func1::new({
                                                                                                                                                                          let dictGenericHeytingAlgebra
                                                                                                                                                                              =
                                                                                                                                                                              dictGenericHeytingAlgebra.clone();
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
                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericConj_prime(),
                                                                                                                                                                                                                                                                                                                                                &&&dictGenericHeytingAlgebra),
                                                                                                                                                                                                                                                                                                             &&&matchValue_3),
                                                                                                                                                                                                                                                                          &&&matchValue_4))
                                                                                                                                                                                                  }
                                                                                                                                                                                          })
                                                                                                                                                                      }),
                                                                                                                                                         add(string("genericDisj\'"),
                                                                                                                                                             &&Func1::new({
                                                                                                                                                                              let dictGenericHeytingAlgebra
                                                                                                                                                                                  =
                                                                                                                                                                                  dictGenericHeytingAlgebra.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |v_2|
                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                  let v_2
                                                                                                                                                                                                      =
                                                                                                                                                                                                      v_2.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |v1_2|
                                                                                                                                                                                                      {
                                                                                                                                                                                                          let matchValue_6 =
                                                                                                                                                                                                              Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                                                          let matchValue_7 =
                                                                                                                                                                                                              Sharpurs_Prelude::unbox(v1_2);
                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericDisj_prime(),
                                                                                                                                                                                                                                                                                                                                                    &&&dictGenericHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                 &&&matchValue_6),
                                                                                                                                                                                                                                                                              &&&matchValue_7))
                                                                                                                                                                                                      }
                                                                                                                                                                                              })
                                                                                                                                                                          }),
                                                                                                                                                             add(string("genericNot\'"),
                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                  let dictGenericHeytingAlgebra
                                                                                                                                                                                      =
                                                                                                                                                                                      dictGenericHeytingAlgebra.clone();
                                                                                                                                                                                  move
                                                                                                                                                                                      |v_3|
                                                                                                                                                                                      {
                                                                                                                                                                                          let a =
                                                                                                                                                                                              Sharpurs_Prelude::unbox(v_3);
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericNot_prime(),
                                                                                                                                                                                                                                                                                                 &&&dictGenericHeytingAlgebra),
                                                                                                                                                                                                                                                              &&&a))
                                                                                                                                                                                      }
                                                                                                                                                                              }),
                                                                                                                                                                 empty::<string,
                                                                                                                                                                         &dyn Any>())))))))))
    }
    pub fn Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct()
     -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct.get_or_init(||
                                                                                 &Func1::new(move
                                                                                                 |dictGenericHeytingAlgebra|
                                                                                                 &Func1::new({
                                                                                                                 let dictGenericHeytingAlgebra
                                                                                                                     =
                                                                                                                     dictGenericHeytingAlgebra.clone();
                                                                                                                 move
                                                                                                                     |dictGenericHeytingAlgebra1|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_GenericHeytingAlgebrausd_Dict(),
                                                                                                                                                      &&&add(string("genericFF\'"),
                                                                                                                                                             &&LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericFF_prime(),
                                                                                                                                                                                                                                                                       &&&dictGenericHeytingAlgebra),
                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericFF_prime(),
                                                                                                                                                                                                                                                                       dictGenericHeytingAlgebra1))),
                                                                                                                                                             add(string("genericTT\'"),
                                                                                                                                                                 &&LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericTT_prime(),
                                                                                                                                                                                                                                                                           &&&dictGenericHeytingAlgebra),
                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericTT_prime(),
                                                                                                                                                                                                                                                                           dictGenericHeytingAlgebra1))),
                                                                                                                                                                 add(string("genericImplies\'"),
                                                                                                                                                                     &&Func1::new({
                                                                                                                                                                                      let dictGenericHeytingAlgebra1
                                                                                                                                                                                          =
                                                                                                                                                                                          dictGenericHeytingAlgebra1.clone();
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
                                                                                                                                                                                                                  &LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericImplies_prime(),
                                                                                                                                                                                                                                                                                                                                                                                                 &&&dictGenericHeytingAlgebra),
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
                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericImplies_prime(),
                                                                                                                                                                                                                                                                                                                                                                                                 &&&dictGenericHeytingAlgebra1),
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
                                                                                                                                                                     add(string("genericConj\'"),
                                                                                                                                                                         &&Func1::new({
                                                                                                                                                                                          let dictGenericHeytingAlgebra1
                                                                                                                                                                                              =
                                                                                                                                                                                              dictGenericHeytingAlgebra1.clone();
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
                                                                                                                                                                                                                      &LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericConj_prime(),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&dictGenericHeytingAlgebra),
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
                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericConj_prime(),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&dictGenericHeytingAlgebra1),
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
                                                                                                                                                                         add(string("genericDisj\'"),
                                                                                                                                                                             &&Func1::new({
                                                                                                                                                                                              let dictGenericHeytingAlgebra1
                                                                                                                                                                                                  =
                                                                                                                                                                                                  dictGenericHeytingAlgebra1.clone();
                                                                                                                                                                                              move
                                                                                                                                                                                                  |v_2|
                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                  let v_2
                                                                                                                                                                                                                      =
                                                                                                                                                                                                                      v_2.clone();
                                                                                                                                                                                                                  move
                                                                                                                                                                                                                      |v1_2|
                                                                                                                                                                                                                      {
                                                                                                                                                                                                                          let matchValue_6:
                                                                                                                                                                                                                                  LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                                                                          let matchValue_7:
                                                                                                                                                                                                                                  LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(v1_2);
                                                                                                                                                                                                                          &LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericDisj_prime(),
                                                                                                                                                                                                                                                                                                                                                                                                         &&&dictGenericHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                                                                      &&&match matchValue_6.as_ref()
                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                             Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                                                                                                   &&&match matchValue_7.as_ref()
                                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                                          Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericDisj_prime(),
                                                                                                                                                                                                                                                                                                                                                                                                         &&&dictGenericHeytingAlgebra1),
                                                                                                                                                                                                                                                                                                                                                                      &&&match matchValue_6.as_ref()
                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                             Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                                                                                                   &&&match matchValue_7.as_ref()
                                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                                          Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                                      })))
                                                                                                                                                                                                                      }
                                                                                                                                                                                                              })
                                                                                                                                                                                          }),
                                                                                                                                                                             add(string("genericNot\'"),
                                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                                  let dictGenericHeytingAlgebra1
                                                                                                                                                                                                      =
                                                                                                                                                                                                      dictGenericHeytingAlgebra1.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |v_3|
                                                                                                                                                                                                      {
                                                                                                                                                                                                          let matchValue_9:
                                                                                                                                                                                                                  LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                                                                                                              Sharpurs_Prelude::unbox(v_3);
                                                                                                                                                                                                          &LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericNot_prime(),
                                                                                                                                                                                                                                                                                                                                                      &&&dictGenericHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                   &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                          Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericNot_prime(),
                                                                                                                                                                                                                                                                                                                                                      &&&dictGenericHeytingAlgebra1),
                                                                                                                                                                                                                                                                                                                   &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                          Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                      })))
                                                                                                                                                                                                      }
                                                                                                                                                                                              }),
                                                                                                                                                                                 empty::<string,
                                                                                                                                                                                         &dyn Any>())))))))
                                                                                                             })))
    }
    pub fn Data_HeytingAlgebra_Generic_genericConj() -> &dyn Any {
        static Data_HeytingAlgebra_Generic_genericConj:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_Generic_genericConj.get_or_init(||
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
                                                                                                        |dictGenericHeytingAlgebra|
                                                                                                        &Func1::new({
                                                                                                                        let dictGenericHeytingAlgebra
                                                                                                                            =
                                                                                                                            dictGenericHeytingAlgebra.clone();
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
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra_Generic::Data_HeytingAlgebra_Generic_genericConj_prime(),
                                                                                                                                                                                                                                                                                          &&&dictGenericHeytingAlgebra),
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
