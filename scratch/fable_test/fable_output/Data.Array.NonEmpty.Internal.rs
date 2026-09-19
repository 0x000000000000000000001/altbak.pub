pub mod PureScript_Data_Array_NonEmpty_Internal {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_9201da02::PureScript_Data_FoldableWithIndex;
    use crate::module_d2b7e5fc::PureScript_Data_Function_Uncurried;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f8b1f47e::PureScript_Data_FunctorWithIndex;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_2563115d::PureScript_Data_Semigroup_Foldable;
    use crate::module_abab3d09::PureScript_Data_Semigroup_Traversable;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_829cacf6::PureScript_Data_TraversableWithIndex;
    use crate::module_b1754f14::PureScript_Data_Unfoldable1;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub mod Data_Array_NonEmpty_Internal_FFI {
        use super::*;
        use fable_library_rust::Array_::append;
        use fable_library_rust::Native_::Func2;
        use fable_library_rust::Native_::fix2;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_array;
        pub fn foldr1Impl(f: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let acc = arr[count(arr.clone()) - 1_i32].clone();
            for i in (0_i32..=count(arr.clone()) - 2_i32).rev() {
                let step1 =
                    Sharpurs_Prelude::sharpurs_apply(f, &arr[i].clone());
                acc.set(Sharpurs_Prelude::sharpurs_apply(&step1, &acc))
            }
            acc
        }
        pub fn foldl1Impl(f: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let acc = arr[0_i32].clone();
            for i in 1_i32..=count(arr.clone()) - 1_i32 {
                let step1 = Sharpurs_Prelude::sharpurs_apply(f, &acc);
                acc.set(Sharpurs_Prelude::sharpurs_apply(&step1,
                                                         &arr[i].clone()))
            }
            acc
        }
        pub fn traverse1Impl(apply: &dyn Any, mapFn: &dyn Any, f: &dyn Any,
                             arrayVal: &dyn Any) -> &dyn Any {
            let arr = arrayVal.clone();
            let go =
                Func2::new({
                               let apply = apply.clone();
                               let arr = arr.clone();
                               let f = f.clone();
                               let mapFn = mapFn.clone();
                               move |bot: i32, top: i32|
                                   fix2(&(move |go, bot: i32, top: i32|
                                              {
                                                  let diff: i32 = top - bot;
                                                  if diff == 1_i32 {
                                                      Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&mapFn,
                                                                                                                         &&Func1::new(move
                                                                                                                                          |a|
                                                                                                                                          &new_array(&[a.clone()]))),
                                                                                       &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                         &arr[bot].clone()))
                                                  } else {
                                                      if diff == 2_i32 {
                                                          Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&apply,
                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&mapFn,
                                                                                                                                                                                                 &&Func1::new(move
                                                                                                                                                                                                                  |a_1|
                                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                                  let a_1
                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                      a_1.clone();
                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                      |b|
                                                                                                                                                                                                                                      &new_array(&[a_1,
                                                                                                                                                                                                                                                   b.clone()])
                                                                                                                                                                                                                              }))),
                                                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                                                                                 &arr[bot].clone()))),
                                                                                           &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                             &arr[bot
                                                                                                                                      +
                                                                                                                                      1_i32].clone()))
                                                      } else {
                                                          if diff == 3_i32 {
                                                              Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&apply,
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&apply,
                                                                                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&mapFn,
                                                                                                                                                                                                                                                                         &&Func1::new(move
                                                                                                                                                                                                                                                                                          |a_2|
                                                                                                                                                                                                                                                                                          &Func1::new({
                                                                                                                                                                                                                                                                                                          let a_2
                                                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                                                              a_2.clone();
                                                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                                                              |b_1|
                                                                                                                                                                                                                                                                                                              &Func1::new({
                                                                                                                                                                                                                                                                                                                              let b_1
                                                                                                                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                                                                                                                  b_1.clone();
                                                                                                                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                                                                                                                  |c|
                                                                                                                                                                                                                                                                                                                                  &new_array(&[a_2,
                                                                                                                                                                                                                                                                                                                                               b_1,
                                                                                                                                                                                                                                                                                                                                               c.clone()])
                                                                                                                                                                                                                                                                                                                          })
                                                                                                                                                                                                                                                                                                      }))),
                                                                                                                                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                                                                                                                                                         &arr[bot].clone()))),
                                                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                                                                                     &arr[bot
                                                                                                                                                                                                              +
                                                                                                                                                                                                              1_i32].clone()))),
                                                                                               &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                 &arr[bot
                                                                                                                                          +
                                                                                                                                          2_i32].clone()))
                                                          } else {
                                                              let pivot: i32 =
                                                                  bot +
                                                                      diff /
                                                                          4_i32
                                                                          *
                                                                          2_i32;
                                                              Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&apply,
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&mapFn,
                                                                                                                                                                                                     &&Func1::new(move
                                                                                                                                                                                                                      |xsVal|
                                                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                                                      let xsVal
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          xsVal.clone();
                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                          |ysVal|
                                                                                                                                                                                                                                          &append(xsVal,
                                                                                                                                                                                                                                                  ysVal.clone())
                                                                                                                                                                                                                                  }))),
                                                                                                                                                                   &go(bot,
                                                                                                                                                                       pivot))),
                                                                                               &go(pivot,
                                                                                                   top))
                                                          }
                                                      }
                                                  }
                                              }), bot, top)
                           });
            go(0_i32, count(arr.clone()))
        }
    }
    pub fn Data_Array_NonEmpty_Internal_foldl1Impl() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_foldl1Impl:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_foldl1Impl.get_or_init(||
                                                                &Func1::new(move
                                                                                |f|
                                                                                Func1::new({
                                                                                               let f
                                                                                                   =
                                                                                                   f.clone();
                                                                                               move
                                                                                                   |xs|
                                                                                                   PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_FFI::foldl1Impl(&f,
                                                                                                                                                                                         xs)
                                                                                           })))
    }
    pub fn Data_Array_NonEmpty_Internal_foldr1Impl() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_foldr1Impl:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_foldr1Impl.get_or_init(||
                                                                &Func1::new(move
                                                                                |f|
                                                                                Func1::new({
                                                                                               let f
                                                                                                   =
                                                                                                   f.clone();
                                                                                               move
                                                                                                   |xs|
                                                                                                   PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_FFI::foldr1Impl(&f,
                                                                                                                                                                                         xs)
                                                                                           })))
    }
    pub fn Data_Array_NonEmpty_Internal_traverse1Impl() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_traverse1Impl:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_traverse1Impl.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |apply|
                                                                                   Func1::new({
                                                                                                  let apply
                                                                                                      =
                                                                                                      apply.clone();
                                                                                                  move
                                                                                                      |mapFn|
                                                                                                      Func1::new({
                                                                                                                     let mapFn
                                                                                                                         =
                                                                                                                         mapFn.clone();
                                                                                                                     move
                                                                                                                         |f|
                                                                                                                         Func1::new({
                                                                                                                                        let f
                                                                                                                                            =
                                                                                                                                            f.clone();
                                                                                                                                        move
                                                                                                                                            |arrayVal|
                                                                                                                                            PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_FFI::traverse1Impl(&apply,
                                                                                                                                                                                                                                     &mapFn,
                                                                                                                                                                                                                                     &f,
                                                                                                                                                                                                                                     arrayVal)
                                                                                                                                    })
                                                                                                                 })
                                                                                              })))
    }
    pub fn Data_Array_NonEmpty_Internal_NonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_NonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_NonEmptyArray.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |x|
                                                                                   x.clone()))
    }
    pub fn Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray()
     -> &dyn Any {
        static Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray.get_or_init(||
                                                                              &PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldable1Array())
    }
    pub fn Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray()
     -> &dyn Any {
        static Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray.get_or_init(||
                                                                                       &PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traversableWithIndexArray())
    }
    pub fn Data_Array_NonEmpty_Internal_traversableNonEmptyArray()
     -> &dyn Any {
        static Data_Array_NonEmpty_Internal_traversableNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_traversableNonEmptyArray.get_or_init(||
                                                                              &PureScript_Data_Traversable::Data_Traversable_traversableArray())
    }
    pub fn Data_Array_NonEmpty_Internal_showNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_showNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_showNonEmptyArray.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictShow|
                                                                                       {
                                                                                           let showArray =
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_showArray(),
                                                                                                                                dictShow);
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                            &&&add(string("show"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let showArray
                                                                                                                                                        =
                                                                                                                                                        showArray.clone();
                                                                                                                                                    move
                                                                                                                                                        |v|
                                                                                                                                                        {
                                                                                                                                                            let xs =
                                                                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                &&&string("(NonEmptyArray ")),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                         &&&showArray),
                                                                                                                                                                                                                                                                                                      &&&xs)),
                                                                                                                                                                                                                                &&&string(")")))
                                                                                                                                                        }
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>()))
                                                                                       }))
    }
    pub fn Data_Array_NonEmpty_Internal_semigroupNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_semigroupNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_semigroupNonEmptyArray.get_or_init(||
                                                                            &PureScript_Data_Semigroup::Data_Semigroup_semigroupArray())
    }
    pub fn Data_Array_NonEmpty_Internal_ordNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_ordNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_ordNonEmptyArray.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictOrd|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordArray(),
                                                                                                                       dictOrd)))
    }
    pub fn Data_Array_NonEmpty_Internal_ord1NonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_ord1NonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_ord1NonEmptyArray.get_or_init(||
                                                                       &PureScript_Data_Ord::Data_Ord_ord1Array())
    }
    pub fn Data_Array_NonEmpty_Internal_monadNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_monadNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_monadNonEmptyArray.get_or_init(||
                                                                        &PureScript_Control_Monad::Control_Monad_monadArray())
    }
    pub fn Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray()
     -> &dyn Any {
        static Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray.get_or_init(||
                                                                                   &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexArray())
    }
    pub fn Data_Array_NonEmpty_Internal_functorNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_functorNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_functorNonEmptyArray.get_or_init(||
                                                                          &PureScript_Data_Functor::Data_Functor_functorArray())
    }
    pub fn Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray()
     -> &dyn Any {
        static Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray.get_or_init(||
                                                                                    &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexArray())
    }
    pub fn Data_Array_NonEmpty_Internal_foldableNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_foldableNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_foldableNonEmptyArray.get_or_init(||
                                                                           &PureScript_Data_Foldable::Data_Foldable_foldableArray())
    }
    pub fn Data_Array_NonEmpty_Internal_foldable1NonEmptyArray_004097()
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_Foldable1usd_Dict(),
                                         &&&add(string("foldMap1"),
                                                &&Func1::new({
                                                                 let Data_Array_NonEmpty_Internal_foldable1NonEmptyArray_004097_002d1
                                                                     =
                                                                     Data_Array_NonEmpty_Internal_foldable1NonEmptyArray_004097_002d1.clone();
                                                                 move
                                                                     |dictSemigroup|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1DefaultL(),
                                                                                                                                                                            &&&Data_Array_NonEmpty_Internal_foldable1NonEmptyArray_004097_002d1.Value),
                                                                                                                                         &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_functorNonEmptyArray()),
                                                                                                      dictSemigroup)
                                                             }),
                                                add(string("foldr1"),
                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                                      &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_foldr1Impl()),
                                                    add(string("foldl1"),
                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                                          &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_foldl1Impl()),
                                                        add(string("Foldable0"),
                                                            &&Func1::new(move
                                                                             |usd__unused|
                                                                             &PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_foldableNonEmptyArray()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_Array_NonEmpty_Internal_foldable1NonEmptyArray_004097_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static
         Data_Array_NonEmpty_Internal_foldable1NonEmptyArray_004097_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_foldable1NonEmptyArray_004097_002d1.get_or_init(||
                                                                                         Lazy(Data_Array_NonEmpty_Internal_foldable1NonEmptyArray_004097.clone()))
    }
    pub fn Data_Array_NonEmpty_Internal_foldable1NonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_foldable1NonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_foldable1NonEmptyArray.get_or_init(||
                                                                            Data_Array_NonEmpty_Internal_foldable1NonEmptyArray_004097_002d1.Value)
    }
    pub fn Data_Array_NonEmpty_Internal_traversable1NonEmptyArray_0040100()
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_Traversable1usd_Dict(),
                                         &&&add(string("traverse1"),
                                                &&Func1::new(move |dictApply|
                                                                 {
                                                                     let apply =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                          dictApply);
                                                                     let map =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                    Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                     &Func1::new({
                                                                                     let apply
                                                                                         =
                                                                                         apply.clone();
                                                                                     let map
                                                                                         =
                                                                                         map.clone();
                                                                                     move
                                                                                         |f|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                                                                                                                                                                                   &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_traverse1Impl()),
                                                                                                                                                                                                &&&apply),
                                                                                                                                                             &&&map),
                                                                                                                          f)
                                                                                 })
                                                                 }),
                                                add(string("sequence1"),
                                                    &&Func1::new({
                                                                     let Data_Array_NonEmpty_Internal_traversable1NonEmptyArray_0040100_002d1
                                                                         =
                                                                         Data_Array_NonEmpty_Internal_traversable1NonEmptyArray_0040100_002d1.clone();
                                                                     move
                                                                         |dictApply_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_sequence1Default(),
                                                                                                                                             &&&Data_Array_NonEmpty_Internal_traversable1NonEmptyArray_0040100_002d1.Value),
                                                                                                          dictApply_1)
                                                                 }),
                                                    add(string("Foldable10"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()),
                                                        add(string("Traversable1"),
                                                            &&Func1::new(move
                                                                             |usd__unused_1|
                                                                             &PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_traversableNonEmptyArray()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_Array_NonEmpty_Internal_traversable1NonEmptyArray_0040100_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static
         Data_Array_NonEmpty_Internal_traversable1NonEmptyArray_0040100_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_traversable1NonEmptyArray_0040100_002d1.get_or_init(||
                                                                                             Lazy(Data_Array_NonEmpty_Internal_traversable1NonEmptyArray_0040100.clone()))
    }
    pub fn Data_Array_NonEmpty_Internal_traversable1NonEmptyArray()
     -> &dyn Any {
        static Data_Array_NonEmpty_Internal_traversable1NonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_traversable1NonEmptyArray.get_or_init(||
                                                                               Data_Array_NonEmpty_Internal_traversable1NonEmptyArray_0040100_002d1.Value)
    }
    pub fn Data_Array_NonEmpty_Internal_eqNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_eqNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_eqNonEmptyArray.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictEq|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eqArray(),
                                                                                                                      dictEq)))
    }
    pub fn Data_Array_NonEmpty_Internal_eq1NonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_eq1NonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_eq1NonEmptyArray.get_or_init(||
                                                                      &PureScript_Data_Eq::Data_Eq_eq1Array())
    }
    pub fn Data_Array_NonEmpty_Internal_bindNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_bindNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_bindNonEmptyArray.get_or_init(||
                                                                       &PureScript_Control_Bind::Control_Bind_bindArray())
    }
    pub fn Data_Array_NonEmpty_Internal_applyNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_applyNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_applyNonEmptyArray.get_or_init(||
                                                                        &PureScript_Control_Apply::Control_Apply_applyArray())
    }
    pub fn Data_Array_NonEmpty_Internal_applicativeNonEmptyArray()
     -> &dyn Any {
        static Data_Array_NonEmpty_Internal_applicativeNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_applicativeNonEmptyArray.get_or_init(||
                                                                              &PureScript_Control_Applicative::Control_Applicative_applicativeArray())
    }
    pub fn Data_Array_NonEmpty_Internal_altNonEmptyArray() -> &dyn Any {
        static Data_Array_NonEmpty_Internal_altNonEmptyArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_Internal_altNonEmptyArray.get_or_init(||
                                                                      &PureScript_Control_Alt::Control_Alt_altArray())
    }
}
