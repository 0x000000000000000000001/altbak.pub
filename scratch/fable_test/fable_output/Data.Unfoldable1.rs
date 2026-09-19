pub mod PureScript_Data_Unfoldable1 {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_abab3d09::PureScript_Data_Semigroup_Traversable;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub mod Data_Unfoldable1_FFI {
        use super::*;
        use fable_library_rust::NativeArray_::add as add_1;
        use fable_library_rust::NativeArray_::new_copy;
        use fable_library_rust::NativeArray_::new_empty;
        pub fn unfoldr1ArrayImpl(isNothing: &dyn Any, fromJust: &dyn Any,
                                 fst: &dyn Any, snd: &dyn Any, f: &dyn Any,
                                 b: &dyn Any) -> &dyn Any {
            let result = new_empty::<&dyn Any>();
            let value = b.clone();
            let looping: MutCell<bool> = MutCell::new(true);
            while looping.get() {
                let tuple = f(value);
                add_1(result.clone(), fst(tuple));
                {
                    let maybe = snd(tuple);
                    if isNothing(maybe) {
                        looping.set(false)
                    } else { value.set(fromJust(maybe)) }
                }
            }
            &new_copy(result.clone())
        }
    }
    pub fn Data_Unfoldable1_unfoldr1ArrayImpl() -> &dyn Any {
        static Data_Unfoldable1_unfoldr1ArrayImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable1_unfoldr1ArrayImpl.get_or_init(||
                                                           &Func1::new(move
                                                                           |isNothing|
                                                                           Func1::new({
                                                                                          let isNothing
                                                                                              =
                                                                                              isNothing.clone();
                                                                                          move
                                                                                              |fromJust|
                                                                                              Func1::new({
                                                                                                             let fromJust
                                                                                                                 =
                                                                                                                 fromJust.clone();
                                                                                                             move
                                                                                                                 |fst|
                                                                                                                 Func1::new({
                                                                                                                                let fst
                                                                                                                                    =
                                                                                                                                    fst.clone();
                                                                                                                                move
                                                                                                                                    |snd|
                                                                                                                                    Func1::new({
                                                                                                                                                   let snd
                                                                                                                                                       =
                                                                                                                                                       snd.clone();
                                                                                                                                                   move
                                                                                                                                                       |f|
                                                                                                                                                       Func1::new({
                                                                                                                                                                      let f
                                                                                                                                                                          =
                                                                                                                                                                          f.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |b|
                                                                                                                                                                          PureScript_Data_Unfoldable1::Data_Unfoldable1_FFI::unfoldr1ArrayImpl(&isNothing,
                                                                                                                                                                                                                                               &fromJust,
                                                                                                                                                                                                                                               &fst,
                                                                                                                                                                                                                                               &snd,
                                                                                                                                                                                                                                               &f,
                                                                                                                                                                                                                                               b)
                                                                                                                                                                  })
                                                                                                                                               })
                                                                                                                            })
                                                                                                         })
                                                                                      })))
    }
    pub fn Data_Unfoldable1_Unfoldable1usd_Dict() -> &dyn Any {
        static Data_Unfoldable1_Unfoldable1usd_Dict: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Unfoldable1_Unfoldable1usd_Dict.get_or_init(||
                                                             &Func1::new(move
                                                                             |x|
                                                                             x.clone()))
    }
    pub fn Data_Unfoldable1_unfoldr1() -> &dyn Any {
        static Data_Unfoldable1_unfoldr1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable1_unfoldr1.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("unfoldr1"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Unfoldable1_unfoldable1Maybe() -> &dyn Any {
        static Data_Unfoldable1_unfoldable1Maybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable1_unfoldable1Maybe.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_Unfoldable1usd_Dict(),
                                                                                           &&&add(string("unfoldr1"),
                                                                                                  &&Func1::new(move
                                                                                                                   |f|
                                                                                                                   &Func1::new({
                                                                                                                                   let f
                                                                                                                                       =
                                                                                                                                       f.clone();
                                                                                                                                   move
                                                                                                                                       |b|
                                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_fst(),
                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                  b))))
                                                                                                                               })),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Unfoldable1_unfoldable1Array() -> &dyn Any {
        static Data_Unfoldable1_unfoldable1Array: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable1_unfoldable1Array.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_Unfoldable1usd_Dict(),
                                                                                           &&&add(string("unfoldr1"),
                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1ArrayImpl(),
                                                                                                                                                                                                                                             &&&PureScript_Data_Maybe::Data_Maybe_isNothing()),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                                                                             &&&Func1::new(move
                                                                                                                                                                                                                                                               |usd__unused|
                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined())))),
                                                                                                                                                                       &&&PureScript_Data_Tuple::Data_Tuple_fst()),
                                                                                                                                    &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Unfoldable1_replicate1() -> &dyn Any {
        static Data_Unfoldable1_replicate1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable1_replicate1.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictUnfoldable1|
                                                                    &Func1::new({
                                                                                    let dictUnfoldable1
                                                                                        =
                                                                                        dictUnfoldable1.clone();
                                                                                    move
                                                                                        |n|
                                                                                        &Func1::new({
                                                                                                        let n
                                                                                                            =
                                                                                                            n.clone();
                                                                                                        move
                                                                                                            |v|
                                                                                                            {
                                                                                                                let step =
                                                                                                                    &Func1::new({
                                                                                                                                    let v
                                                                                                                                        =
                                                                                                                                        v.clone();
                                                                                                                                    move
                                                                                                                                        |i|
                                                                                                                                        {
                                                                                                                                            let matchValue =
                                                                                                                                                Sharpurs_Prelude::unbox(i);
                                                                                                                                            if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                                               &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                                                         &&&0_i32))
                                                                                                                                               {
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&v,
                                                                                                                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
                                                                                                                                            } else {
                                                                                                                                                if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                                   {
                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&v,
                                                                                                                                                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                                                                                                                                    &&&1_i32)))))
                                                                                                                                                } else {
                                                                                                                                                    panic!("{}",
                                                                                                                                                           LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Unfoldable1.fs"),
                                  Data1: 43_i32,
                                  Data2: 156_i32,}).get_Message(),)
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                });
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                                                                                       &&&dictUnfoldable1),
                                                                                                                                                                                    &&&step),
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                          &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                       &&&n),
                                                                                                                                                                                    &&&1_i32))
                                                                                                            }
                                                                                                    })
                                                                                })))
    }
    pub fn Data_Unfoldable1_replicate1A() -> &dyn Any {
        static Data_Unfoldable1_replicate1A: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable1_replicate1A.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictApply|
                                                                     &Func1::new({
                                                                                     let dictApply
                                                                                         =
                                                                                         dictApply.clone();
                                                                                     move
                                                                                         |dictUnfoldable1|
                                                                                         &Func1::new({
                                                                                                         let dictUnfoldable1
                                                                                                             =
                                                                                                             dictUnfoldable1.clone();
                                                                                                         move
                                                                                                             |dictTraversable1|
                                                                                                             &Func1::new({
                                                                                                                             let dictTraversable1
                                                                                                                                 =
                                                                                                                                 dictTraversable1.clone();
                                                                                                                             move
                                                                                                                                 |n|
                                                                                                                                 &Func1::new({
                                                                                                                                                 let n
                                                                                                                                                     =
                                                                                                                                                     n.clone();
                                                                                                                                                 move
                                                                                                                                                     |m|
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_sequence1(),
                                                                                                                                                                                                                                                            &&&dictTraversable1),
                                                                                                                                                                                                                         &&&dictApply),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_replicate1(),
                                                                                                                                                                                                                                                                                               &&&dictUnfoldable1),
                                                                                                                                                                                                                                                            &&&n),
                                                                                                                                                                                                                         m))
                                                                                                                                             })
                                                                                                                         })
                                                                                                     })
                                                                                 })))
    }
    pub fn Data_Unfoldable1_singleton() -> &dyn Any {
        static Data_Unfoldable1_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable1_singleton.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictUnfoldable1|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_replicate1(),
                                                                                                                                       dictUnfoldable1),
                                                                                                    &&&1_i32)))
    }
    pub fn Data_Unfoldable1_range() -> &dyn Any {
        static Data_Unfoldable1_range: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable1_range.get_or_init(||
                                               &Func1::new(move
                                                               |dictUnfoldable1|
                                                               &Func1::new({
                                                                               let dictUnfoldable1
                                                                                   =
                                                                                   dictUnfoldable1.clone();
                                                                               move
                                                                                   |start|
                                                                                   &Func1::new({
                                                                                                   let start
                                                                                                       =
                                                                                                       start.clone();
                                                                                                   move
                                                                                                       |end_var|
                                                                                                       {
                                                                                                           let go =
                                                                                                               &Func1::new({
                                                                                                                               let end_var
                                                                                                                                   =
                                                                                                                                   end_var.clone();
                                                                                                                               move
                                                                                                                                   |delta|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let delta
                                                                                                                                                       =
                                                                                                                                                       delta.clone();
                                                                                                                                                   move
                                                                                                                                                       |i|
                                                                                                                                                       {
                                                                                                                                                           let i_prime =
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                   i),
                                                                                                                                                                                                &&&delta);
                                                                                                                                                           &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(i.clone(),
                                                                                                                                                                                                                   {
                                                                                                                                                                                                                       let matchValue =
                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                                                                                                        i),
                                                                                                                                                                                                                                                                                     &&&end_var));
                                                                                                                                                                                                                       match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                        &matchValue)
                                                                                                                                                                                                                           {
                                                                                                                                                                                                                           0_i32
                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                                                                           _
                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&i_prime)),
                                                                                                                                                                                                                       }
                                                                                                                                                                                                                   }))
                                                                                                                                                       }
                                                                                                                                               })
                                                                                                                           });
                                                                                                           let delta_1 =
                                                                                                               {
                                                                                                                   let matchValue_1 =
                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                                       &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                    end_var),
                                                                                                                                                                                 &&&start));
                                                                                                                   match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                    &matchValue_1)
                                                                                                                       {
                                                                                                                       0_i32
                                                                                                                       =>
                                                                                                                       &1_i32,
                                                                                                                       _
                                                                                                                       =>
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                                                                                                                           &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                        &&&1_i32),
                                                                                                                   }
                                                                                                               };
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                                                                                  &&&dictUnfoldable1),
                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                                                                                  &&&delta_1)),
                                                                                                                                            &&&start)
                                                                                                       }
                                                                                               })
                                                                           })))
    }
    pub fn Data_Unfoldable1_iterateN() -> &dyn Any {
        static Data_Unfoldable1_iterateN: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable1_iterateN.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictUnfoldable1|
                                                                  &Func1::new({
                                                                                  let dictUnfoldable1
                                                                                      =
                                                                                      dictUnfoldable1.clone();
                                                                                  move
                                                                                      |n|
                                                                                      &Func1::new({
                                                                                                      let n
                                                                                                          =
                                                                                                          n.clone();
                                                                                                      move
                                                                                                          |f|
                                                                                                          &Func1::new({
                                                                                                                          let f
                                                                                                                              =
                                                                                                                              f.clone();
                                                                                                                          move
                                                                                                                              |s|
                                                                                                                              {
                                                                                                                                  let go =
                                                                                                                                      &Func1::new(move
                                                                                                                                                      |v|
                                                                                                                                                      {
                                                                                                                                                          let matchValue:
                                                                                                                                                                  LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                              Sharpurs_Prelude::unbox(v);
                                                                                                                                                          let x =
                                                                                                                                                              match matchValue.as_ref()
                                                                                                                                                                  {
                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                             _)
                                                                                                                                                                  =>
                                                                                                                                                                  x.clone(),
                                                                                                                                                              };
                                                                                                                                                          let n_prime =
                                                                                                                                                              match matchValue.as_ref()
                                                                                                                                                                  {
                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                             x)
                                                                                                                                                                  =>
                                                                                                                                                                  x.clone(),
                                                                                                                                                              };
                                                                                                                                                          &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&x,
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                      let matchValue_1 =
                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                                                                                       &&&n_prime),
                                                                                                                                                                                                                                                                                    &&&0_i32));
                                                                                                                                                                                                                      match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                       &matchValue_1)
                                                                                                                                                                                                                          {
                                                                                                                                                                                                                          0_i32
                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                |usd__arg1|
                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                      |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                                      Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                     let usd__arg1_1
                                                                                                                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                                                                                                                         usd__arg1_1.clone();
                                                                                                                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                                                                                                                         |usd__arg2|
                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                 usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                       &&&x))),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                                                    &&&n_prime),
                                                                                                                                                                                                                                                                                                                                 &&&1_i32))),
                                                                                                                                                                                                                          _
                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                          &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                                                                      }
                                                                                                                                                                                                                  }))
                                                                                                                                                      });
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                                                                                                                                            &&&dictUnfoldable1),
                                                                                                                                                                                                                                         &&&go)),
                                                                                                                                                                   &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(s.clone(),
                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                 &&&n),
                                                                                                                                                                                                                                                              &&&1_i32))))
                                                                                                                              }
                                                                                                                      })
                                                                                                  })
                                                                              })))
    }
}
