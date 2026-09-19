pub mod PureScript_Control_Monad_Gen {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::fix1;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_851afbc9::PureScript_Control_Monad_Gen_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_3b72fe33::PureScript_Data_Monoid_Additive;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_2563115d::PureScript_Data_Semigroup_Foldable;
    use crate::module_8e57e1bc::PureScript_Data_Semigroup_Last;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    #[derive(Clone, Debug,)]
    pub enum Control_Monad_Gen_LL {
        Control_Monad_Gen_Consusd_Ctor(&dyn Any, &dyn Any),
        Control_Monad_Gen_Nilusd_Ctor,
    }
    impl core::fmt::Display for
     PureScript_Control_Monad_Gen::Control_Monad_Gen_LL {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Control_Monad_Gen_monoidAdditive() -> &dyn Any {
        static Control_Monad_Gen_monoidAdditive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_monoidAdditive.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_monoidAdditive(),
                                                                                          &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()))
    }
    pub fn Control_Monad_Gen_Cons() -> &dyn Any {
        static Control_Monad_Gen_Cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Cons.get_or_init(||
                                               &Func1::new(move |usd__arg1|
                                                               Func1::new({
                                                                              let usd__arg1
                                                                                  =
                                                                                  usd__arg1.clone();
                                                                              move
                                                                                  |usd__arg2|
                                                                                  &LrcPtr::new(PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(usd__arg1,
                                                                                                                                                                                  usd__arg2.clone()))
                                                                          })))
    }
    pub fn Control_Monad_Gen_Nil() -> &dyn Any {
        static Control_Monad_Gen_Nil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Nil.get_or_init(||
                                              &LrcPtr::new(PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Nilusd_Ctor))
    }
    pub fn Control_Monad_Gen_FreqSemigroup() -> &dyn Any {
        static Control_Monad_Gen_FreqSemigroup: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_FreqSemigroup.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        x.clone()))
    }
    pub fn Control_Monad_Gen_unfoldable() -> &dyn Any {
        static Control_Monad_Gen_unfoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_unfoldable.get_or_init(||
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
                                                                                             let pure_var =
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                            Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
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
                                                                                                             let pure_var
                                                                                                                 =
                                                                                                                 pure_var.clone();
                                                                                                             move
                                                                                                                 |dictUnfoldable|
                                                                                                                 &Func1::new({
                                                                                                                                 let dictUnfoldable
                                                                                                                                     =
                                                                                                                                     dictUnfoldable.clone();
                                                                                                                                 move
                                                                                                                                     |gen|
                                                                                                                                     {
                                                                                                                                         let unfold =
                                                                                                                                             &Func1::new(move
                                                                                                                                                             |v|
                                                                                                                                                             {
                                                                                                                                                                 let matchValue:
                                                                                                                                                                         LrcPtr<PureScript_Control_Monad_Gen::Control_Monad_Gen_LL> =
                                                                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                                                                 match matchValue.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                     PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(matchValue_0_0,
                                                                                                                                                                                                                                                        matchValue_0_1)
                                                                                                                                                                     =>
                                                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                         PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                                                     },
                                                                                                                                                                                                                                                                                    &match matchValue.as_ref()
                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                         PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                                                     })))),
                                                                                                                                                                     _
                                                                                                                                                                     =>
                                                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                 }
                                                                                                                                                             });
                                                                                                                                         let loopGen =
                                                                                                                                             &Func1::new({
                                                                                                                                                             let gen
                                                                                                                                                                 =
                                                                                                                                                                 gen.clone();
                                                                                                                                                             move
                                                                                                                                                                 |v_1|
                                                                                                                                                                 {
                                                                                                                                                                     let matchValue_1:
                                                                                                                                                                             LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                         Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                     if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                                     &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                  &&&0_i32))
                                                                                                                                                                        {
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                             &&&pure_var),
                                                                                                                                                                                                          &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                             })))
                                                                                                                                                                     } else {
                                                                                                                                                                         if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                                                            {
                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                    &&&Bind1),
                                                                                                                                                                                                                                                 &&&gen),
                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                let matchValue_1
                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                    matchValue_1.clone();
                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                    |x_1|
                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                        &&&pure_var),
                                                                                                                                                                                                                                                                     &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(x_1.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                })),
                                                                                                                                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&1_i32))))))
                                                                                                                                                                                                                            }))
                                                                                                                                                                         } else {
                                                                                                                                                                             panic!("{}",
                                                                                                                                                                                    LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Control.Monad.Gen.fs"),
                                  Data1: 19_i32,
                                  Data2: 1286_i32,}).get_Message(),)
                                                                                                                                                                         }
                                                                                                                                                                     }
                                                                                                                                                                 }
                                                                                                                                                         });
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                &&&Functor0),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                                                                                                                                                   &&&dictUnfoldable),
                                                                                                                                                                                                                                                &&&unfold)),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_sized(),
                                                                                                                                                                                                                                                &&&dictMonadGen),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                                                                                                                                                         &&&dictMonadRec),
                                                                                                                                                                                                                                                                                                                      &&&loopGen)),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                                                                     |usd__arg1|
                                                                                                                                                                                                                                                                                                     Func1::new({
                                                                                                                                                                                                                                                                                                                    let usd__arg1
                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                        usd__arg1.clone();
                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                        |usd__arg2|
                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                                })),
                                                                                                                                                                                                                                                                                   &&&LrcPtr::new(PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Nilusd_Ctor)))))
                                                                                                                                     }
                                                                                                                             })
                                                                                                         })
                                                                                         }
                                                                                 })))
    }
    pub fn Control_Monad_Gen_semigroupFreqSemigroup() -> &dyn Any {
        static Control_Monad_Gen_semigroupFreqSemigroup:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_semigroupFreqSemigroup.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                  &&&add(string("append"),
                                                                                                         &&Func1::new(move
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
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen::Control_Monad_Gen_FreqSemigroup(),
                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                     let matchValue_1
                                                                                                                                                                                                         =
                                                                                                                                                                                                         matchValue_1.clone();
                                                                                                                                                                                                     move
                                                                                                                                                                                                         |pos|
                                                                                                                                                                                                         {
                                                                                                                                                                                                             let matchValue_3:
                                                                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                            pos));
                                                                                                                                                                                                             {
                                                                                                                                                                                                                 let activePatternResult:
                                                                                                                                                                                                                         LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                                     Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                        });
                                                                                                                                                                                                                 if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(activePatternResult_1_0)
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        activePatternResult.as_ref()
                                                                                                                                                                                                                    {
                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                      &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                             Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                                         })
                                                                                                                                                                                                                 } else {
                                                                                                                                                                                                                     &matchValue_3
                                                                                                                                                                                                                 }
                                                                                                                                                                                                             }
                                                                                                                                                                                                         }
                                                                                                                                                                                                 }))
                                                                                                                                              }
                                                                                                                                      })),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Control_Monad_Gen_getFreqVal() -> &dyn Any {
        static Control_Monad_Gen_getFreqVal: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_getFreqVal.get_or_init(||
                                                     &Func1::new(move |v|
                                                                     {
                                                                         let f =
                                                                             Sharpurs_Prelude::unbox(v);
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                             &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                          &&&f)
                                                                     }))
    }
    pub fn Control_Monad_Gen_fromIndex() -> &dyn Any {
        static Control_Monad_Gen_fromIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_fromIndex.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictFoldable1|
                                                                    {
                                                                        let Foldable0 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                    Sharpurs_Prelude::unbox(dictFoldable1)),
                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                        &Func1::new({
                                                                                        let Foldable0
                                                                                            =
                                                                                            Foldable0.clone();
                                                                                        let dictFoldable1
                                                                                            =
                                                                                            dictFoldable1.clone();
                                                                                        move
                                                                                            |i|
                                                                                            &Func1::new({
                                                                                                            let i
                                                                                                                =
                                                                                                                i.clone();
                                                                                                            move
                                                                                                                |xs|
                                                                                                                {
                                                                                                                    let go_2 =
                                                                                                                        Func0::new({
                                                                                                                                       let go_tco
                                                                                                                                           =
                                                                                                                                           go_tco.clone();
                                                                                                                                       move
                                                                                                                                           ||
                                                                                                                                           &Func1::new({
                                                                                                                                                           let go_tco
                                                                                                                                                               =
                                                                                                                                                               go_tco.clone();
                                                                                                                                                           move
                                                                                                                                                               |v|
                                                                                                                                                               Func1::new({
                                                                                                                                                                              let go_tco
                                                                                                                                                                                  =
                                                                                                                                                                                  go_tco.clone();
                                                                                                                                                                              let v
                                                                                                                                                                                  =
                                                                                                                                                                                  v.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |v1|
                                                                                                                                                                                  go_tco(v)(v1.clone())
                                                                                                                                                                          })
                                                                                                                                                       })
                                                                                                                                   });
                                                                                                                    let go_1 =
                                                                                                                        Lazy(go_2);
                                                                                                                    let go_tco =
                                                                                                                        Func1::new({
                                                                                                                                       let xs
                                                                                                                                           =
                                                                                                                                           xs.clone();
                                                                                                                                       move
                                                                                                                                           |v_1|
                                                                                                                                           fix1(&(move
                                                                                                                                                      |go_tco,
                                                                                                                                                       v_1|
                                                                                                                                                      Func1::new({
                                                                                                                                                                     let go_tco
                                                                                                                                                                         =
                                                                                                                                                                         go_tco.clone();
                                                                                                                                                                     let v_1
                                                                                                                                                                         =
                                                                                                                                                                         v_1.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |v1_1|
                                                                                                                                                                         {
                                                                                                                                                                             let matchValue =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                             let matchValue_1:
                                                                                                                                                                                     LrcPtr<PureScript_Control_Monad_Gen::Control_Monad_Gen_LL> =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                             if let PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Nilusd_Ctor
                                                                                                                                                                                    =
                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_un(),
                                                                                                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                     &&&PureScript_Data_Semigroup_Last::Data_Semigroup_Last_Last()),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                                                                                                                                                                              &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup_Last::Data_Semigroup_Last_semigroupLast()),
                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Semigroup_Last::Data_Semigroup_Last_Last()),
                                                                                                                                                                                                                                                     &&&xs))
                                                                                                                                                                             } else {
                                                                                                                                                                                 if let PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Nilusd_Ctor
                                                                                                                                                                                        =
                                                                                                                                                                                        Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                               PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                               _
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                           }).as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                     &match matchValue_1.as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                          PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                          =>
                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                          _
                                                                                                                                                                                          =>
                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                      }
                                                                                                                                                                                 } else {
                                                                                                                                                                                     if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                                                                  &&&0_i32))
                                                                                                                                                                                        {
                                                                                                                                                                                         &match matchValue_1.as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          }
                                                                                                                                                                                     } else {
                                                                                                                                                                                         go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                                                 &&&1_i32))(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                 PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                 _
                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                                                                                             })
                                                                                                                                                                                     }
                                                                                                                                                                                 }
                                                                                                                                                                             }
                                                                                                                                                                         }
                                                                                                                                                                 })),
                                                                                                                                                v_1.clone())
                                                                                                                                   });
                                                                                                                    let go =
                                                                                                                        go_1.Value;
                                                                                                                    go_tco(&i)(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                         &&&Foldable0),
                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                        |usd__arg1|
                                                                                                                                                                                                                                                        Func1::new({
                                                                                                                                                                                                                                                                       let usd__arg1
                                                                                                                                                                                                                                                                           =
                                                                                                                                                                                                                                                                           usd__arg1.clone();
                                                                                                                                                                                                                                                                       move
                                                                                                                                                                                                                                                                           |usd__arg2|
                                                                                                                                                                                                                                                                           &LrcPtr::new(PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Consusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                           usd__arg2.clone()))
                                                                                                                                                                                                                                                                   }))),
                                                                                                                                                                                                   &&&LrcPtr::new(PureScript_Control_Monad_Gen::Control_Monad_Gen_LL::Control_Monad_Gen_Nilusd_Ctor)),
                                                                                                                                                                xs))
                                                                                                                }
                                                                                                        })
                                                                                    })
                                                                    }))
    }
    pub fn Control_Monad_Gen_oneOf() -> &dyn Any {
        static Control_Monad_Gen_oneOf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_oneOf.get_or_init(||
                                                &Func1::new(move
                                                                |dictMonadGen|
                                                                {
                                                                    let Bind1 =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                    &Func1::new({
                                                                                    let Bind1
                                                                                        =
                                                                                        Bind1.clone();
                                                                                    let dictMonadGen
                                                                                        =
                                                                                        dictMonadGen.clone();
                                                                                    move
                                                                                        |dictFoldable1|
                                                                                        {
                                                                                            let Foldable0 =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                        Sharpurs_Prelude::unbox(dictFoldable1)),
                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                                            &Func1::new({
                                                                                                            let Foldable0
                                                                                                                =
                                                                                                                Foldable0.clone();
                                                                                                            let dictFoldable1
                                                                                                                =
                                                                                                                dictFoldable1.clone();
                                                                                                            move
                                                                                                                |xs|
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                       &&&Bind1),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_chooseInt(),
                                                                                                                                                                                                                                                                                             &&&dictMonadGen),
                                                                                                                                                                                                                                                          &&&0_i32),
                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_length(),
                                                                                                                                                                                                                                                                                                                                                                                                      &&&Foldable0),
                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                xs)),
                                                                                                                                                                                                                                                          &&&1_i32))),
                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                   let xs
                                                                                                                                                                       =
                                                                                                                                                                       xs.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |n|
                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen::Control_Monad_Gen_fromIndex(),
                                                                                                                                                                                                                                                                              &&&dictFoldable1),
                                                                                                                                                                                                                                           n),
                                                                                                                                                                                                        &&&xs)
                                                                                                                                                               }))
                                                                                                        })
                                                                                        }
                                                                                })
                                                                }))
    }
    pub fn Control_Monad_Gen_freqSemigroup() -> &dyn Any {
        static Control_Monad_Gen_freqSemigroup: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_freqSemigroup.get_or_init(||
                                                        &Func1::new(move |v|
                                                                        {
                                                                            let matchValue:
                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                Sharpurs_Prelude::unbox(v);
                                                                            let x =
                                                                                match matchValue.as_ref()
                                                                                    {
                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                               x)
                                                                                    =>
                                                                                    x.clone(),
                                                                                };
                                                                            let weight =
                                                                                match matchValue.as_ref()
                                                                                    {
                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                               _)
                                                                                    =>
                                                                                    x.clone(),
                                                                                };
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen::Control_Monad_Gen_FreqSemigroup(),
                                                                                                             &&&Func1::new({
                                                                                                                               let weight
                                                                                                                                   =
                                                                                                                                   weight.clone();
                                                                                                                               let x
                                                                                                                                   =
                                                                                                                                   x.clone();
                                                                                                                               move
                                                                                                                                   |pos|
                                                                                                                                   {
                                                                                                                                       let matchValue_1 =
                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                                                           &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                                                                        pos),
                                                                                                                                                                                                     &&&weight));
                                                                                                                                       match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                        &matchValue_1)
                                                                                                                                           {
                                                                                                                                           0_i32
                                                                                                                                           =>
                                                                                                                                           &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Ring::Data_Ring_ringNumber()),
                                                                                                                                                                                                                                                                                                                              pos),
                                                                                                                                                                                                                                                                                           &&&weight))),
                                                                                                                                                                                                   &x)),
                                                                                                                                           _
                                                                                                                                           =>
                                                                                                                                           &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                                                   &x)),
                                                                                                                                       }
                                                                                                                                   }
                                                                                                                           }))
                                                                        }))
    }
    pub fn Control_Monad_Gen_frequency() -> &dyn Any {
        static Control_Monad_Gen_frequency: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_frequency.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictMonadGen|
                                                                    {
                                                                        let Bind1 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                        &Func1::new({
                                                                                        let Bind1
                                                                                            =
                                                                                            Bind1.clone();
                                                                                        let dictMonadGen
                                                                                            =
                                                                                            dictMonadGen.clone();
                                                                                        move
                                                                                            |dictFoldable1|
                                                                                            {
                                                                                                let foldMap =
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictFoldable1)),
                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                     &&&PureScript_Control_Monad_Gen::Control_Monad_Gen_monoidAdditive());
                                                                                                &Func1::new({
                                                                                                                let dictFoldable1
                                                                                                                    =
                                                                                                                    dictFoldable1.clone();
                                                                                                                let foldMap
                                                                                                                    =
                                                                                                                    foldMap.clone();
                                                                                                                move
                                                                                                                    |xs|
                                                                                                                    {
                                                                                                                        let total =
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_alaF(),
                                                                                                                                                                                                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                      &&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive()),
                                                                                                                                                                                                                                   &&&foldMap),
                                                                                                                                                                                                &&&PureScript_Data_Tuple::Data_Tuple_fst()),
                                                                                                                                                             xs);
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                               &&&Bind1),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_chooseFloat(),
                                                                                                                                                                                                                                                                                                     &&&dictMonadGen),
                                                                                                                                                                                                                                                                  &&&0.0_f64),
                                                                                                                                                                                                                               &&&total)),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen::Control_Monad_Gen_getFreqVal(),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                                                                                                                                                        &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_Gen::Control_Monad_Gen_semigroupFreqSemigroup()),
                                                                                                                                                                                                                                                                  &&&PureScript_Control_Monad_Gen::Control_Monad_Gen_freqSemigroup()),
                                                                                                                                                                                                                               xs)))
                                                                                                                    }
                                                                                                            })
                                                                                            }
                                                                                    })
                                                                    }))
    }
    pub fn Control_Monad_Gen_filtered() -> &dyn Any {
        static Control_Monad_Gen_filtered: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_filtered.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonadRec|
                                                                   &Func1::new({
                                                                                   let dictMonadRec
                                                                                       =
                                                                                       dictMonadRec.clone();
                                                                                   move
                                                                                       |dictMonadGen|
                                                                                       {
                                                                                           let Functor0 =
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                                                                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                                           &Func1::new({
                                                                                                           let Functor0
                                                                                                               =
                                                                                                               Functor0.clone();
                                                                                                           move
                                                                                                               |gen|
                                                                                                               {
                                                                                                                   let go =
                                                                                                                       &Func1::new({
                                                                                                                                       let gen
                                                                                                                                           =
                                                                                                                                           gen.clone();
                                                                                                                                       move
                                                                                                                                           |v|
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_mapFlipped(),
                                                                                                                                                                                                                                                  &&&Functor0),
                                                                                                                                                                                                               &&&gen),
                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                              |a|
                                                                                                                                                                                              {
                                                                                                                                                                                                  let matchValue:
                                                                                                                                                                                                          LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                      Sharpurs_Prelude::unbox(a);
                                                                                                                                                                                                  match matchValue.as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                      Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_1_0)),
                                                                                                                                                                                                      _
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                                                                                                                  }
                                                                                                                                                                                              }))
                                                                                                                                   });
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                          &&&dictMonadRec),
                                                                                                                                                                                       &&&go),
                                                                                                                                                    &&&PureScript_Data_Unit::Data_Unit_unit())
                                                                                                               }
                                                                                                       })
                                                                                       }
                                                                               })))
    }
    pub fn Control_Monad_Gen_suchThat() -> &dyn Any {
        static Control_Monad_Gen_suchThat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_suchThat.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonadRec|
                                                                   {
                                                                       let filtered1 =
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen::Control_Monad_Gen_filtered(),
                                                                                                            dictMonadRec);
                                                                       &Func1::new({
                                                                                       let filtered1
                                                                                           =
                                                                                           filtered1.clone();
                                                                                       move
                                                                                           |dictMonadGen|
                                                                                           {
                                                                                               let filtered2 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&filtered1,
                                                                                                                                    dictMonadGen);
                                                                                               let Functor0 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                               &Func1::new({
                                                                                                               let Functor0
                                                                                                                   =
                                                                                                                   Functor0.clone();
                                                                                                               let filtered2
                                                                                                                   =
                                                                                                                   filtered2.clone();
                                                                                                               move
                                                                                                                   |gen|
                                                                                                                   &Func1::new({
                                                                                                                                   let gen
                                                                                                                                       =
                                                                                                                                       gen.clone();
                                                                                                                                   move
                                                                                                                                       |pred|
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                           &&&filtered2),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_mapFlipped(),
                                                                                                                                                                                                                                                                                 &&&Functor0),
                                                                                                                                                                                                                                              &&&gen),
                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                             let pred
                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                 pred.clone();
                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                 |a|
                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                     let matchValue =
                                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&pred,
                                                                                                                                                                                                                                                                                                   a));
                                                                                                                                                                                                                                     match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                                      &matchValue)
                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                         0_i32
                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(a.clone())),
                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                                                                                     }
                                                                                                                                                                                                                                 }
                                                                                                                                                                                                                         })))
                                                                                                                               })
                                                                                                           })
                                                                                           }
                                                                                   })
                                                                   }))
    }
    pub fn Control_Monad_Gen_elements() -> &dyn Any {
        static Control_Monad_Gen_elements: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_elements.get_or_init(||
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
                                                                       let pure_var =
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                      Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                       &Func1::new({
                                                                                       let Bind1
                                                                                           =
                                                                                           Bind1.clone();
                                                                                       let dictMonadGen
                                                                                           =
                                                                                           dictMonadGen.clone();
                                                                                       let pure_var
                                                                                           =
                                                                                           pure_var.clone();
                                                                                       move
                                                                                           |dictFoldable1|
                                                                                           {
                                                                                               let Foldable0 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                           Sharpurs_Prelude::unbox(dictFoldable1)),
                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                               &Func1::new({
                                                                                                               let Foldable0
                                                                                                                   =
                                                                                                                   Foldable0.clone();
                                                                                                               let dictFoldable1
                                                                                                                   =
                                                                                                                   dictFoldable1.clone();
                                                                                                               move
                                                                                                                   |xs|
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                          &&&Bind1),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_chooseInt(),
                                                                                                                                                                                                                                                                                                &&&dictMonadGen),
                                                                                                                                                                                                                                                             &&&0_i32),
                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_length(),
                                                                                                                                                                                                                                                                                                                                                                                                         &&&Foldable0),
                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                   xs)),
                                                                                                                                                                                                                                                             &&&1_i32))),
                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                      let xs
                                                                                                                                                                          =
                                                                                                                                                                          xs.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |n|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                              &&&pure_var),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen::Control_Monad_Gen_fromIndex(),
                                                                                                                                                                                                                                                                                                                    &&&dictFoldable1),
                                                                                                                                                                                                                                                                                 n),
                                                                                                                                                                                                                                              &&&xs))
                                                                                                                                                                  }))
                                                                                                           })
                                                                                           }
                                                                                   })
                                                                   }))
    }
    pub fn Control_Monad_Gen_choose() -> &dyn Any {
        static Control_Monad_Gen_choose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_choose.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictMonadGen|
                                                                 {
                                                                     let Bind1 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     let chooseBool =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_chooseBool(),
                                                                                                          dictMonadGen);
                                                                     &Func1::new({
                                                                                     let Bind1
                                                                                         =
                                                                                         Bind1.clone();
                                                                                     let chooseBool
                                                                                         =
                                                                                         chooseBool.clone();
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
                                                                                                                                                                                 &&&chooseBool),
                                                                                                                                              &&&Func1::new({
                                                                                                                                                                let genB
                                                                                                                                                                    =
                                                                                                                                                                    genB.clone();
                                                                                                                                                                move
                                                                                                                                                                    |v|
                                                                                                                                                                    {
                                                                                                                                                                        let matchValue =
                                                                                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                                                                                        match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                         &matchValue)
                                                                                                                                                                            {
                                                                                                                                                                            0_i32
                                                                                                                                                                            =>
                                                                                                                                                                            &genA,
                                                                                                                                                                            _
                                                                                                                                                                            =>
                                                                                                                                                                            &genB,
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                            }))
                                                                                                     })
                                                                                 })
                                                                 }))
    }
}
