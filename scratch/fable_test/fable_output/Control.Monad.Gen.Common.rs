pub mod PureScript_Control_Monad_Gen_Common {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_851afbc9::PureScript_Control_Monad_Gen_Class;
    use crate::module_e72de349::PureScript_Control_Monad_Gen;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty::Data_NonEmpty_NonEmpty;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Gen_Common_genTuple() -> &dyn Any {
        static Control_Monad_Gen_Common_genTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Common_genTuple.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictApply|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                              dictApply),
                                                                                                           &&&Func1::new(move
                                                                                                                             |usd__arg1|
                                                                                                                             Func1::new({
                                                                                                                                            let usd__arg1
                                                                                                                                                =
                                                                                                                                                usd__arg1.clone();
                                                                                                                                            move
                                                                                                                                                |usd__arg2|
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                        usd__arg2.clone()))
                                                                                                                                        })))))
    }
    pub fn Control_Monad_Gen_Common_genNonEmpty() -> &dyn Any {
        static Control_Monad_Gen_Common_genNonEmpty: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Gen_Common_genNonEmpty.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictMonadRec|
                                                                             &Func1::new({
                                                                                             let dictMonadRec
                                                                                                 =
                                                                                                 dictMonadRec.clone();
                                                                                             move
                                                                                                 |dictMonadGen|
                                                                                                 {
                                                                                                     let Bind1 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                     let Apply0 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                 Sharpurs_Prelude::unbox(&&Bind1)),
                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                     let Functor0 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Bind1)),
                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                     &Func1::new({
                                                                                                                     let Apply0
                                                                                                                         =
                                                                                                                         Apply0.clone();
                                                                                                                     let Functor0
                                                                                                                         =
                                                                                                                         Functor0.clone();
                                                                                                                     let dictMonadGen
                                                                                                                         =
                                                                                                                         dictMonadGen.clone();
                                                                                                                     move
                                                                                                                         |dictUnfoldable|
                                                                                                                         &Func1::new({
                                                                                                                                         let dictUnfoldable
                                                                                                                                             =
                                                                                                                                             dictUnfoldable.clone();
                                                                                                                                         move
                                                                                                                                             |gen|
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                    &&&Apply0),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                                         |usd__arg1|
                                                                                                                                                                                                                                                                                                         Func1::new({
                                                                                                                                                                                                                                                                                                                        let usd__arg1
                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                            usd__arg1.clone();
                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                            |usd__arg2|
                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                                usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                                    }))),
                                                                                                                                                                                                                                                    gen)),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_resize(),
                                                                                                                                                                                                                                                                                       &&&dictMonadGen),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_max(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                                                                                                                             &&&0_i32)),
                                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                                         |v|
                                                                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                                                             v),
                                                                                                                                                                                                                                                                                                                                          &&&1_i32)))),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen::Control_Monad_Gen_unfoldable(),
                                                                                                                                                                                                                                                                                                                                                             &&&dictMonadRec),
                                                                                                                                                                                                                                                                                                                          &&&dictMonadGen),
                                                                                                                                                                                                                                                                                       &&&dictUnfoldable),
                                                                                                                                                                                                                                                    gen)))
                                                                                                                                     })
                                                                                                                 })
                                                                                                 }
                                                                                         })))
    }
    pub fn Control_Monad_Gen_Common_genMaybe_prime() -> &dyn Any {
        static Control_Monad_Gen_Common_genMaybe_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Common_genMaybe_prime.get_or_init(||
                                                                &Func1::new(move
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
                                                                                    let Applicative0 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    &Func1::new({
                                                                                                    let Applicative0
                                                                                                        =
                                                                                                        Applicative0.clone();
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
                                                                                                        |bias|
                                                                                                        &Func1::new({
                                                                                                                        let bias
                                                                                                                            =
                                                                                                                            bias.clone();
                                                                                                                        move
                                                                                                                            |gen|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                   &&&Bind1),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_chooseFloat(),
                                                                                                                                                                                                                                                                                                         &&&dictMonadGen),
                                                                                                                                                                                                                                                                      &&&0.0_f64),
                                                                                                                                                                                                                                   &&&1.0_f64)),
                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                               let gen
                                                                                                                                                                                   =
                                                                                                                                                                                   gen.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |n|
                                                                                                                                                                                   {
                                                                                                                                                                                       let matchValue =
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                                                                                                                        n),
                                                                                                                                                                                                                                                     &&&bias));
                                                                                                                                                                                       match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                        &matchValue)
                                                                                                                                                                                           {
                                                                                                                                                                                           0_i32
                                                                                                                                                                                           =>
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                  &&&Functor0),
                                                                                                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                                                                                                 |usd__arg1|
                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                            &&&gen),
                                                                                                                                                                                           _
                                                                                                                                                                                           =>
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                               &&&Applicative0),
                                                                                                                                                                                                                            &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                                                       }
                                                                                                                                                                                   }
                                                                                                                                                                           }))
                                                                                                                    })
                                                                                                })
                                                                                }))
    }
    pub fn Control_Monad_Gen_Common_genMaybe() -> &dyn Any {
        static Control_Monad_Gen_Common_genMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Common_genMaybe.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonadGen|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Common::Control_Monad_Gen_Common_genMaybe_prime(),
                                                                                                                                              dictMonadGen),
                                                                                                           &&&0.75_f64)))
    }
    pub fn Control_Monad_Gen_Common_genIdentity() -> &dyn Any {
        static Control_Monad_Gen_Common_genIdentity: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Gen_Common_genIdentity.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictFunctor|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                 dictFunctor),
                                                                                                              &&&PureScript_Data_Identity::Data_Identity_Identity())))
    }
    pub fn Control_Monad_Gen_Common_genEither_prime() -> &dyn Any {
        static Control_Monad_Gen_Common_genEither_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Common_genEither_prime.get_or_init(||
                                                                 &Func1::new(move
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
                                                                                                         |bias|
                                                                                                         &Func1::new({
                                                                                                                         let bias
                                                                                                                             =
                                                                                                                             bias.clone();
                                                                                                                         move
                                                                                                                             |genA|
                                                                                                                             &Func1::new({
                                                                                                                                             let genA
                                                                                                                                                 =
                                                                                                                                                 genA.clone();
                                                                                                                                             move
                                                                                                                                                 |genB|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                        &&&Bind1),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_chooseFloat(),
                                                                                                                                                                                                                                                                                                                              &&&dictMonadGen),
                                                                                                                                                                                                                                                                                           &&&0.0_f64),
                                                                                                                                                                                                                                                        &&&1.0_f64)),
                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                    let genB
                                                                                                                                                                                                        =
                                                                                                                                                                                                        genB.clone();
                                                                                                                                                                                                    move
                                                                                                                                                                                                        |n|
                                                                                                                                                                                                        {
                                                                                                                                                                                                            let matchValue =
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                                                                                                                                             n),
                                                                                                                                                                                                                                                                          &&&bias));
                                                                                                                                                                                                            match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                             &matchValue)
                                                                                                                                                                                                                {
                                                                                                                                                                                                                0_i32
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                       &&&Functor0),
                                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                                      |usd__arg1|
                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                                 &&&genA),
                                                                                                                                                                                                                _
                                                                                                                                                                                                                =>
                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                       &&&Functor0),
                                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                                      |usd__arg1_1|
                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                                                                                 &&&genB),
                                                                                                                                                                                                            }
                                                                                                                                                                                                        }
                                                                                                                                                                                                }))
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                                 }))
    }
    pub fn Control_Monad_Gen_Common_genEither() -> &dyn Any {
        static Control_Monad_Gen_Common_genEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Common_genEither.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictMonadGen|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Common::Control_Monad_Gen_Common_genEither_prime(),
                                                                                                                                               dictMonadGen),
                                                                                                            &&&0.5_f64)))
    }
}
