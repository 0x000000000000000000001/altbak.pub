pub mod PureScript_Data_Unfoldable {
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
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_b1754f14::PureScript_Data_Unfoldable1;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Data_Unfoldable_FFI {
        use super::*;
        use fable_library_rust::NativeArray_::add as add_1;
        use fable_library_rust::NativeArray_::new_copy;
        use fable_library_rust::NativeArray_::new_empty;
        pub fn unfoldrArrayImpl(isNothing: &dyn Any, fromJust: &dyn Any,
                                fst: &dyn Any, snd: &dyn Any, f: &dyn Any,
                                b: &dyn Any) -> &dyn Any {
            let result = new_empty::<&dyn Any>();
            let value = b.clone();
            let looping: MutCell<bool> = MutCell::new(true);
            while looping.get() {
                let maybe = f(value);
                if isNothing(maybe) {
                    looping.set(false)
                } else {
                    let tuple = fromJust(maybe);
                    add_1(result.clone(), fst(tuple));
                    value.set(snd(tuple))
                }
            }
            &new_copy(result.clone())
        }
    }
    pub fn Data_Unfoldable_unfoldrArrayImpl() -> &dyn Any {
        static Data_Unfoldable_unfoldrArrayImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable_unfoldrArrayImpl.get_or_init(||
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
                                                                                                                                                                        PureScript_Data_Unfoldable::Data_Unfoldable_FFI::unfoldrArrayImpl(&isNothing,
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
    pub fn Data_Unfoldable_Unfoldableusd_Dict() -> &dyn Any {
        static Data_Unfoldable_Unfoldableusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable_Unfoldableusd_Dict.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Data_Unfoldable_unfoldr() -> &dyn Any {
        static Data_Unfoldable_unfoldr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable_unfoldr.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("unfoldr"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Unfoldable_unfoldableMaybe() -> &dyn Any {
        static Data_Unfoldable_unfoldableMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable_unfoldableMaybe.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_Unfoldableusd_Dict(),
                                                                                         &&&add(string("unfoldr"),
                                                                                                &&Func1::new(move
                                                                                                                 |f|
                                                                                                                 &Func1::new({
                                                                                                                                 let f
                                                                                                                                     =
                                                                                                                                     f.clone();
                                                                                                                                 move
                                                                                                                                     |b|
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                            &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                         &&&PureScript_Data_Tuple::Data_Tuple_fst()),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                         b))
                                                                                                                             })),
                                                                                                add(string("Unfoldable10"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldable1Maybe()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
    pub fn Data_Unfoldable_unfoldableArray() -> &dyn Any {
        static Data_Unfoldable_unfoldableArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable_unfoldableArray.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_Unfoldableusd_Dict(),
                                                                                         &&&add(string("unfoldr"),
                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldrArrayImpl(),
                                                                                                                                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_isNothing()),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                             |usd__unused|
                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined())))),
                                                                                                                                                                     &&&PureScript_Data_Tuple::Data_Tuple_fst()),
                                                                                                                                  &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                add(string("Unfoldable10"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused_1|
                                                                                                                     &PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldable1Array()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
    pub fn Data_Unfoldable_replicate() -> &dyn Any {
        static Data_Unfoldable_replicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable_replicate.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictUnfoldable|
                                                                  &Func1::new({
                                                                                  let dictUnfoldable
                                                                                      =
                                                                                      dictUnfoldable.clone();
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
                                                                                                                                              Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                                              &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                           i),
                                                                                                                                                                                                        &&&0_i32));
                                                                                                                                          match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                           &matchValue)
                                                                                                                                              {
                                                                                                                                              0_i32
                                                                                                                                              =>
                                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                              _
                                                                                                                                              =>
                                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&v,
                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                 i),
                                                                                                                                                                                                                                                                                              &&&1_i32))))),
                                                                                                                                          }
                                                                                                                                      }
                                                                                                                              });
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                                                                                     &&&dictUnfoldable),
                                                                                                                                                                                  &&&step),
                                                                                                                                               &&&n)
                                                                                                          }
                                                                                                  })
                                                                              })))
    }
    pub fn Data_Unfoldable_replicateA() -> &dyn Any {
        static Data_Unfoldable_replicateA: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable_replicateA.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictApplicative|
                                                                   &Func1::new({
                                                                                   let dictApplicative
                                                                                       =
                                                                                       dictApplicative.clone();
                                                                                   move
                                                                                       |dictUnfoldable|
                                                                                       &Func1::new({
                                                                                                       let dictUnfoldable
                                                                                                           =
                                                                                                           dictUnfoldable.clone();
                                                                                                       move
                                                                                                           |dictTraversable|
                                                                                                           &Func1::new({
                                                                                                                           let dictTraversable
                                                                                                                               =
                                                                                                                               dictTraversable.clone();
                                                                                                                           move
                                                                                                                               |n|
                                                                                                                               &Func1::new({
                                                                                                                                               let n
                                                                                                                                                   =
                                                                                                                                                   n.clone();
                                                                                                                                               move
                                                                                                                                                   |m|
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                          &&&dictTraversable),
                                                                                                                                                                                                                       &&&dictApplicative),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_replicate(),
                                                                                                                                                                                                                                                                                             &&&dictUnfoldable),
                                                                                                                                                                                                                                                          &&&n),
                                                                                                                                                                                                                       m))
                                                                                                                                           })
                                                                                                                       })
                                                                                                   })
                                                                               })))
    }
    pub fn Data_Unfoldable_none() -> &dyn Any {
        static Data_Unfoldable_none: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable_none.get_or_init(||
                                             &Func1::new(move |dictUnfoldable|
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                                    dictUnfoldable),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                    &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))),
                                                                                              &&&PureScript_Data_Unit::Data_Unit_unit())))
    }
    pub fn Data_Unfoldable_fromMaybe() -> &dyn Any {
        static Data_Unfoldable_fromMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Unfoldable_fromMaybe.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictUnfoldable|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                      dictUnfoldable),
                                                                                                   &&&Func1::new(move
                                                                                                                     |b|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                            &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
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
                                                                                                                                                                                                                                                                                            }))),
                                                                                                                                                                                                                            &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))),
                                                                                                                                                      b)))))
    }
}
