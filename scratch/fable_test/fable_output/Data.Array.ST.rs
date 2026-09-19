pub mod PureScript_Data_Array_ST {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::NativeArray_::new_array;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_fcf3066b::PureScript_Control_Monad_ST_Internal;
    use crate::module_1d594369::PureScript_Control_Monad_ST_Uncurried;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Data_Array_ST_FFI {
        use super::*;
        use fable_library_rust::Array_::addRangeInPlace;
        use fable_library_rust::Array_::getSubArray;
        use fable_library_rust::Interfaces_::System::Collections::Generic::IComparer_1;
        use fable_library_rust::Interfaces_::System::Collections::Generic::IEnumerable_1;
        use fable_library_rust::Map_::add;
        use fable_library_rust::Map_::empty;
        use fable_library_rust::Native_::Func0;
        use fable_library_rust::Native_::Lrc;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::Native_::interface_cast;
        use fable_library_rust::NativeArray_::add as add_1;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_copy;
        use fable_library_rust::NativeArray_::new_empty;
        use fable_library_rust::NativeArray_::new_init;
        use fable_library_rust::Seq_::toArray;
        use fable_library_rust::String_::string;
        pub fn new() -> &dyn Any {
            static new: MutCell<Option<&dyn Any>> = MutCell::new(None);
            new.get_or_init(|| &Func0::new(move || &new_empty::<&dyn Any>()))
        }
        pub fn peekImpl(just: &dyn Any, nothing: &dyn Any, iVal: &dyn Any,
                        xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let iVal = iVal.clone();
                            let just = just.clone();
                            let nothing = nothing.clone();
                            let xs = xs.clone();
                            move ||
                                {
                                    let arr = xs;
                                    let i: i32 = iVal;
                                    if if i >= 0_i32 {
                                           i < count(arr.clone())
                                       } else { false } {
                                        Sharpurs_Prelude::sharpurs_apply(&just,
                                                                         &arr[i].clone())
                                    } else { nothing }
                                }
                        })
        }
        pub fn pokeImpl(iVal: &dyn Any, a: &dyn Any, xs: &dyn Any)
         -> &dyn Any {
            &Func0::new({
                            let a = a.clone();
                            let iVal = iVal.clone();
                            let xs = xs.clone();
                            move ||
                                {
                                    let arr = xs;
                                    let i: i32 = iVal;
                                    let ret: bool =
                                        if i >= 0_i32 {
                                            i < count(arr.clone())
                                        } else { false };
                                    if ret { arr.get_mut()[i as usize] = a; }
                                    &ret
                                }
                        })
        }
        pub fn lengthImpl(xs: &dyn Any) -> &dyn Any {
            &Func0::new({ let xs = xs.clone(); move || &count(xs) })
        }
        pub fn popImpl(just: &dyn Any, nothing: &dyn Any, xs: &dyn Any)
         -> &dyn Any {
            &Func0::new({
                            let just = just.clone();
                            let nothing = nothing.clone();
                            let xs = xs.clone();
                            move ||
                                {
                                    let arr = xs;
                                    if count(arr.clone()) > 0_i32 {
                                        let el =
                                            arr[count(arr.clone()) -
                                                    1_i32].clone();
                                        arr.get_mut().remove((count(arr.clone())
                                                                  - 1_i32) as
                                                                 usize);
                                        Sharpurs_Prelude::sharpurs_apply(&just,
                                                                         &el)
                                    } else { nothing }
                                }
                        })
        }
        pub fn pushAllImpl(asVal: &dyn Any, xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let xs = xs.clone();
                            move || { defaultOf::<&dyn Any>(); &count(xs) }
                        })
        }
        pub fn shiftImpl(just: &dyn Any, nothing: &dyn Any, xs: &dyn Any)
         -> &dyn Any {
            &Func0::new({
                            let just = just.clone();
                            let nothing = nothing.clone();
                            let xs = xs.clone();
                            move ||
                                {
                                    let arr = xs;
                                    if count(arr.clone()) > 0_i32 {
                                        let el = arr[0_i32].clone();
                                        arr.get_mut().remove(0_i32 as usize);
                                        Sharpurs_Prelude::sharpurs_apply(&just,
                                                                         &el)
                                    } else { nothing }
                                }
                        })
        }
        pub fn unshiftAllImpl(asVal: &dyn Any, xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let xs = xs.clone();
                            move || { defaultOf::<&dyn Any>(); &count(xs) }
                        })
        }
        pub fn spliceImpl(iVal: &dyn Any, howManyVal: &dyn Any,
                          bsVal: &dyn Any, xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let howManyVal = howManyVal.clone();
                            let iVal = iVal.clone();
                            let xs = xs.clone();
                            move ||
                                {
                                    let arr = xs;
                                    let i: i32 = iVal;
                                    let howMany: i32 = howManyVal;
                                    let removed =
                                        getSubArray(arr.clone(), i, howMany);
                                    arr.splice(i, howMany);
                                    defaultOf::<&dyn Any>();
                                    &new_copy(removed)
                                }
                        })
        }
        pub fn unsafeFreezeImpl(xs: &dyn Any) -> &dyn Any {
            &Func0::new({ let xs = xs.clone(); move || &new_copy(xs) })
        }
        pub fn unsafeThawImpl(xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let xs = xs.clone();
                            move ||
                                &toArray(interface_cast!(xs,
                                                         Lrc<dyn IEnumerable_1<&dyn Any>>,).clone())
                        })
        }
        pub fn freezeImpl(xs: &dyn Any) -> &dyn Any {
            &Func0::new({ let xs = xs.clone(); move || &new_copy(xs) })
        }
        pub fn thawImpl(xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let xs = xs.clone();
                            move ||
                                &toArray(interface_cast!(xs,
                                                         Lrc<dyn IEnumerable_1<&dyn Any>>,).clone())
                        })
        }
        pub fn cloneImpl(xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let xs = xs.clone();
                            move ||
                                &toArray(interface_cast!(xs,
                                                         Lrc<dyn IEnumerable_1<&dyn Any>>,).clone())
                        })
        }
        pub fn sortByImpl(compare: &dyn Any, fromOrdering: &dyn Any,
                          xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let compare = compare.clone();
                            let fromOrdering = fromOrdering.clone();
                            let xs = xs.clone();
                            move ||
                                {
                                    let arr = xs;
                                    if count(arr.clone()) < 2_i32 {
                                        &arr
                                    } else {
                                        let outArr = new_copy(arr.clone());
                                        let comparer =
                                            {
                                                struct ObjectExpr {
                                                }
                                                impl <T: Clone + 'static>
                                                 core::fmt::Display for
                                                 dyn IComparer_1<_> {
                                                    fn fmt(&self,
                                                           f:
                                                               &mut core::fmt::Formatter)
                                                     -> core::fmt::Result {
                                                        write!(f, "{}", core::any::type_name::<Self>())
                                                    }
                                                }
                                                impl <T: Clone + 'static>
                                                 IComparer_1<T> for
                                                 ObjectExpr<T> {
                                                    fn Compare(&self, a: _,
                                                               b: _) -> i32 {
                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&fromOrdering,
                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&compare,
                                                                                                                                                                                      a),
                                                                                                                                                    b)))
                                                    }
                                                }
                                                interface_cast!(LrcPtr::new(ObjectExpr::<&dyn Any>{}),
                                                                Lrc<dyn IComparer_1<&dyn Any>>,)
                                            };
                                        let sorted =
                                            panic!("{}",
                                                   1_i32.get_Message(),);
                                        arr.get_mut().clear();
                                        addRangeInPlace(sorted, arr.clone());
                                        &arr
                                    }
                                }
                        })
        }
        pub fn toAssocArrayImpl(xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let xs = xs.clone();
                            move ||
                                {
                                    let arr = xs;
                                    let result =
                                        new_init(&defaultOf(),
                                                 count(arr.clone()));
                                    for i in
                                        0_i32..=count(arr.clone()) - 1_i32 {
                                        let map =
                                            add(string("index"), &i,
                                                {
                                                    let table =
                                                        empty::<string,
                                                                &dyn Any>();
                                                    add(string("value"),
                                                        arr[i].clone(), table)
                                                });
                                        result.get_mut()[i as usize] = &map
                                    }
                                    &result
                                }
                        })
        }
        pub fn pushImpl(a: &dyn Any, xs: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let a = a.clone();
                            let xs = xs.clone();
                            move ||
                                {
                                    let arr = xs;
                                    add_1(arr.clone(), a);
                                    &count(arr)
                                }
                        })
        }
    }
    pub fn Data_Array_ST_cloneImpl() -> &dyn Any {
        static Data_Array_ST_cloneImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_cloneImpl.get_or_init(||
                                                &Func1::new(move |xs|
                                                                PureScript_Data_Array_ST::Data_Array_ST_FFI::cloneImpl(xs)))
    }
    pub fn Data_Array_ST_freezeImpl() -> &dyn Any {
        static Data_Array_ST_freezeImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_freezeImpl.get_or_init(||
                                                 &Func1::new(move |xs|
                                                                 PureScript_Data_Array_ST::Data_Array_ST_FFI::freezeImpl(xs)))
    }
    pub fn Data_Array_ST_lengthImpl() -> &dyn Any {
        static Data_Array_ST_lengthImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_lengthImpl.get_or_init(||
                                                 &Func1::new(move |xs|
                                                                 PureScript_Data_Array_ST::Data_Array_ST_FFI::lengthImpl(xs)))
    }
    pub fn Data_Array_ST_new() -> &dyn Any {
        static Data_Array_ST_new: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_new.get_or_init(||
                                          &PureScript_Data_Array_ST::Data_Array_ST_FFI::new())
    }
    pub fn Data_Array_ST_peekImpl() -> &dyn Any {
        static Data_Array_ST_peekImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_peekImpl.get_or_init(||
                                               &Func1::new(move |just|
                                                               Func1::new({
                                                                              let just
                                                                                  =
                                                                                  just.clone();
                                                                              move
                                                                                  |nothing|
                                                                                  Func1::new({
                                                                                                 let nothing
                                                                                                     =
                                                                                                     nothing.clone();
                                                                                                 move
                                                                                                     |iVal|
                                                                                                     Func1::new({
                                                                                                                    let iVal
                                                                                                                        =
                                                                                                                        iVal.clone();
                                                                                                                    move
                                                                                                                        |xs|
                                                                                                                        PureScript_Data_Array_ST::Data_Array_ST_FFI::peekImpl(&just,
                                                                                                                                                                              &nothing,
                                                                                                                                                                              &iVal,
                                                                                                                                                                              xs)
                                                                                                                })
                                                                                             })
                                                                          })))
    }
    pub fn Data_Array_ST_pokeImpl() -> &dyn Any {
        static Data_Array_ST_pokeImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_pokeImpl.get_or_init(||
                                               &Func1::new(move |iVal|
                                                               Func1::new({
                                                                              let iVal
                                                                                  =
                                                                                  iVal.clone();
                                                                              move
                                                                                  |a|
                                                                                  Func1::new({
                                                                                                 let a
                                                                                                     =
                                                                                                     a.clone();
                                                                                                 move
                                                                                                     |xs|
                                                                                                     PureScript_Data_Array_ST::Data_Array_ST_FFI::pokeImpl(&iVal,
                                                                                                                                                           &a,
                                                                                                                                                           xs)
                                                                                             })
                                                                          })))
    }
    pub fn Data_Array_ST_popImpl() -> &dyn Any {
        static Data_Array_ST_popImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_popImpl.get_or_init(||
                                              &Func1::new(move |just|
                                                              Func1::new({
                                                                             let just
                                                                                 =
                                                                                 just.clone();
                                                                             move
                                                                                 |nothing|
                                                                                 Func1::new({
                                                                                                let nothing
                                                                                                    =
                                                                                                    nothing.clone();
                                                                                                move
                                                                                                    |xs|
                                                                                                    PureScript_Data_Array_ST::Data_Array_ST_FFI::popImpl(&just,
                                                                                                                                                         &nothing,
                                                                                                                                                         xs)
                                                                                            })
                                                                         })))
    }
    pub fn Data_Array_ST_pushAllImpl() -> &dyn Any {
        static Data_Array_ST_pushAllImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_pushAllImpl.get_or_init(||
                                                  &Func1::new(move |asVal|
                                                                  Func1::new({
                                                                                 let asVal
                                                                                     =
                                                                                     asVal.clone();
                                                                                 move
                                                                                     |xs|
                                                                                     PureScript_Data_Array_ST::Data_Array_ST_FFI::pushAllImpl(&asVal,
                                                                                                                                              xs)
                                                                             })))
    }
    pub fn Data_Array_ST_pushImpl() -> &dyn Any {
        static Data_Array_ST_pushImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_pushImpl.get_or_init(||
                                               &Func1::new(move |a|
                                                               Func1::new({
                                                                              let a
                                                                                  =
                                                                                  a.clone();
                                                                              move
                                                                                  |xs|
                                                                                  PureScript_Data_Array_ST::Data_Array_ST_FFI::pushImpl(&a,
                                                                                                                                        xs)
                                                                          })))
    }
    pub fn Data_Array_ST_shiftImpl() -> &dyn Any {
        static Data_Array_ST_shiftImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_shiftImpl.get_or_init(||
                                                &Func1::new(move |just|
                                                                Func1::new({
                                                                               let just
                                                                                   =
                                                                                   just.clone();
                                                                               move
                                                                                   |nothing|
                                                                                   Func1::new({
                                                                                                  let nothing
                                                                                                      =
                                                                                                      nothing.clone();
                                                                                                  move
                                                                                                      |xs|
                                                                                                      PureScript_Data_Array_ST::Data_Array_ST_FFI::shiftImpl(&just,
                                                                                                                                                             &nothing,
                                                                                                                                                             xs)
                                                                                              })
                                                                           })))
    }
    pub fn Data_Array_ST_sortByImpl() -> &dyn Any {
        static Data_Array_ST_sortByImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_sortByImpl.get_or_init(||
                                                 &Func1::new(move |compare|
                                                                 Func1::new({
                                                                                let compare
                                                                                    =
                                                                                    compare.clone();
                                                                                move
                                                                                    |fromOrdering|
                                                                                    Func1::new({
                                                                                                   let fromOrdering
                                                                                                       =
                                                                                                       fromOrdering.clone();
                                                                                                   move
                                                                                                       |xs|
                                                                                                       PureScript_Data_Array_ST::Data_Array_ST_FFI::sortByImpl(&compare,
                                                                                                                                                               &fromOrdering,
                                                                                                                                                               xs)
                                                                                               })
                                                                            })))
    }
    pub fn Data_Array_ST_spliceImpl() -> &dyn Any {
        static Data_Array_ST_spliceImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_spliceImpl.get_or_init(||
                                                 &Func1::new(move |iVal|
                                                                 Func1::new({
                                                                                let iVal
                                                                                    =
                                                                                    iVal.clone();
                                                                                move
                                                                                    |howManyVal|
                                                                                    Func1::new({
                                                                                                   let howManyVal
                                                                                                       =
                                                                                                       howManyVal.clone();
                                                                                                   move
                                                                                                       |bsVal|
                                                                                                       Func1::new({
                                                                                                                      let bsVal
                                                                                                                          =
                                                                                                                          bsVal.clone();
                                                                                                                      move
                                                                                                                          |xs|
                                                                                                                          PureScript_Data_Array_ST::Data_Array_ST_FFI::spliceImpl(&iVal,
                                                                                                                                                                                  &howManyVal,
                                                                                                                                                                                  &bsVal,
                                                                                                                                                                                  xs)
                                                                                                                  })
                                                                                               })
                                                                            })))
    }
    pub fn Data_Array_ST_thawImpl() -> &dyn Any {
        static Data_Array_ST_thawImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_thawImpl.get_or_init(||
                                               &Func1::new(move |xs|
                                                               PureScript_Data_Array_ST::Data_Array_ST_FFI::thawImpl(xs)))
    }
    pub fn Data_Array_ST_toAssocArrayImpl() -> &dyn Any {
        static Data_Array_ST_toAssocArrayImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_toAssocArrayImpl.get_or_init(||
                                                       &Func1::new(move |xs|
                                                                       PureScript_Data_Array_ST::Data_Array_ST_FFI::toAssocArrayImpl(xs)))
    }
    pub fn Data_Array_ST_unsafeFreezeImpl() -> &dyn Any {
        static Data_Array_ST_unsafeFreezeImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_unsafeFreezeImpl.get_or_init(||
                                                       &Func1::new(move |xs|
                                                                       PureScript_Data_Array_ST::Data_Array_ST_FFI::unsafeFreezeImpl(xs)))
    }
    pub fn Data_Array_ST_unsafeThawImpl() -> &dyn Any {
        static Data_Array_ST_unsafeThawImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_unsafeThawImpl.get_or_init(||
                                                     &Func1::new(move |xs|
                                                                     PureScript_Data_Array_ST::Data_Array_ST_FFI::unsafeThawImpl(xs)))
    }
    pub fn Data_Array_ST_unshiftAllImpl() -> &dyn Any {
        static Data_Array_ST_unshiftAllImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_unshiftAllImpl.get_or_init(||
                                                     &Func1::new(move |asVal|
                                                                     Func1::new({
                                                                                    let asVal
                                                                                        =
                                                                                        asVal.clone();
                                                                                    move
                                                                                        |xs|
                                                                                        PureScript_Data_Array_ST::Data_Array_ST_FFI::unshiftAllImpl(&asVal,
                                                                                                                                                    xs)
                                                                                })))
    }
    pub fn Data_Array_ST_unshiftAll() -> &dyn Any {
        static Data_Array_ST_unshiftAll: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_unshiftAll.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn2(),
                                                                                  &&&PureScript_Data_Array_ST::Data_Array_ST_unshiftAllImpl()))
    }
    pub fn Data_Array_ST_unshift() -> &dyn Any {
        static Data_Array_ST_unshift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_unshift.get_or_init(||
                                              &Func1::new(move |a|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn2(),
                                                                                                                                  &&&PureScript_Data_Array_ST::Data_Array_ST_unshiftAllImpl()),
                                                                                               &&&new_array(&[a.clone()]))))
    }
    pub fn Data_Array_ST_unsafeThaw() -> &dyn Any {
        static Data_Array_ST_unsafeThaw: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_unsafeThaw.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn1(),
                                                                                  &&&PureScript_Data_Array_ST::Data_Array_ST_unsafeThawImpl()))
    }
    pub fn Data_Array_ST_unsafeFreeze() -> &dyn Any {
        static Data_Array_ST_unsafeFreeze: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_unsafeFreeze.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn1(),
                                                                                    &&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreezeImpl()))
    }
    pub fn Data_Array_ST_toAssocArray() -> &dyn Any {
        static Data_Array_ST_toAssocArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_toAssocArray.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn1(),
                                                                                    &&&PureScript_Data_Array_ST::Data_Array_ST_toAssocArrayImpl()))
    }
    pub fn Data_Array_ST_thaw() -> &dyn Any {
        static Data_Array_ST_thaw: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_thaw.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn1(),
                                                                            &&&PureScript_Data_Array_ST::Data_Array_ST_thawImpl()))
    }
    pub fn Data_Array_ST_withArray() -> &dyn Any {
        static Data_Array_ST_withArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_withArray.get_or_init(||
                                                &Func1::new(move |f|
                                                                &Func1::new({
                                                                                let f
                                                                                    =
                                                                                    f.clone();
                                                                                move
                                                                                    |xs|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                           &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_thaw(),
                                                                                                                                                                                           xs)),
                                                                                                                     &&&Func1::new(move
                                                                                                                                       |result|
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                              &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                              result)),
                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                          let result
                                                                                                                                                                                              =
                                                                                                                                                                                              result.clone();
                                                                                                                                                                                          move
                                                                                                                                                                                              |usd__unused|
                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze(),
                                                                                                                                                                                                                               &&&result)
                                                                                                                                                                                      }))))
                                                                            })))
    }
    pub fn Data_Array_ST_splice() -> &dyn Any {
        static Data_Array_ST_splice: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_splice.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn4(),
                                                                              &&&PureScript_Data_Array_ST::Data_Array_ST_spliceImpl()))
    }
    pub fn Data_Array_ST_sortBy() -> &dyn Any {
        static Data_Array_ST_sortBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_sortBy.get_or_init(||
                                             &Func1::new(move |comp|
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn3(),
                                                                                                                                                                    &&&PureScript_Data_Array_ST::Data_Array_ST_sortByImpl()),
                                                                                                                                 comp),
                                                                                              &&&Func1::new(move
                                                                                                                |v|
                                                                                                                {
                                                                                                                    let matchValue:
                                                                                                                            LrcPtr<Data_Ordering_Ordering> =
                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                    match matchValue.as_ref()
                                                                                                                        {
                                                                                                                        Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                        =>
                                                                                                                        &0_i32,
                                                                                                                        Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                        =>
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                                                                                                                            &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                         &&&1_i32),
                                                                                                                        _
                                                                                                                        =>
                                                                                                                        &1_i32,
                                                                                                                    }
                                                                                                                }))))
    }
    pub fn Data_Array_ST_sortWith() -> &dyn Any {
        static Data_Array_ST_sortWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_sortWith.get_or_init(||
                                               &Func1::new(move |dictOrd|
                                                               &Func1::new({
                                                                               let dictOrd
                                                                                   =
                                                                                   dictOrd.clone();
                                                                               move
                                                                                   |f|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_sortBy(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_comparing(),
                                                                                                                                                                                          &&&dictOrd),
                                                                                                                                                       f))
                                                                           })))
    }
    pub fn Data_Array_ST_sort() -> &dyn Any {
        static Data_Array_ST_sort: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_sort.get_or_init(||
                                           &Func1::new(move |dictOrd|
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_sortBy(),
                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                               dictOrd))))
    }
    pub fn Data_Array_ST_shift() -> &dyn Any {
        static Data_Array_ST_shift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_shift.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn3(),
                                                                                                                                                   &&&PureScript_Data_Array_ST::Data_Array_ST_shiftImpl()),
                                                                                                                &&&Func1::new(move
                                                                                                                                  |usd__arg1|
                                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                             &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Array_ST_run() -> &dyn Any {
        static Data_Array_ST_run: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_run.get_or_init(||
                                          &Func1::new(move |st|
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_run(),
                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                    &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                 st),
                                                                                                                              &&&PureScript_Data_Array_ST::Data_Array_ST_unsafeFreeze()))))
    }
    pub fn Data_Array_ST_pushAll() -> &dyn Any {
        static Data_Array_ST_pushAll: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_pushAll.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn2(),
                                                                               &&&PureScript_Data_Array_ST::Data_Array_ST_pushAllImpl()))
    }
    pub fn Data_Array_ST_push() -> &dyn Any {
        static Data_Array_ST_push: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_push.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn2(),
                                                                            &&&PureScript_Data_Array_ST::Data_Array_ST_pushImpl()))
    }
    pub fn Data_Array_ST_pop() -> &dyn Any {
        static Data_Array_ST_pop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_pop.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn3(),
                                                                                                                                                 &&&PureScript_Data_Array_ST::Data_Array_ST_popImpl()),
                                                                                                              &&&Func1::new(move
                                                                                                                                |usd__arg1|
                                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                           &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Array_ST_poke() -> &dyn Any {
        static Data_Array_ST_poke: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_poke.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn3(),
                                                                            &&&PureScript_Data_Array_ST::Data_Array_ST_pokeImpl()))
    }
    pub fn Data_Array_ST_peek() -> &dyn Any {
        static Data_Array_ST_peek: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_peek.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn4(),
                                                                                                                                                  &&&PureScript_Data_Array_ST::Data_Array_ST_peekImpl()),
                                                                                                               &&&Func1::new(move
                                                                                                                                 |usd__arg1|
                                                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                            &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Array_ST_modify() -> &dyn Any {
        static Data_Array_ST_modify: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_modify.get_or_init(||
                                             &Func1::new(move |i|
                                                             &Func1::new({
                                                                             let i
                                                                                 =
                                                                                 i.clone();
                                                                             move
                                                                                 |f|
                                                                                 &Func1::new({
                                                                                                 let f
                                                                                                     =
                                                                                                     f.clone();
                                                                                                 move
                                                                                                     |xs|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                            &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_peek(),
                                                                                                                                                                                                                                               &&&i),
                                                                                                                                                                                                            xs)),
                                                                                                                                      &&&Func1::new({
                                                                                                                                                        let xs
                                                                                                                                                            =
                                                                                                                                                            xs.clone();
                                                                                                                                                        move
                                                                                                                                                            |entry|
                                                                                                                                                            {
                                                                                                                                                                let matchValue:
                                                                                                                                                                        LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                    Sharpurs_Prelude::unbox(entry);
                                                                                                                                                                match matchValue.as_ref()
                                                                                                                                                                    {
                                                                                                                                                                    Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                                    =>
                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                        &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                                                                     &&&false),
                                                                                                                                                                    Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                    =>
                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_poke(),
                                                                                                                                                                                                                                                                           &&&i),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                  Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                  _
                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                                                                                              })),
                                                                                                                                                                                                     &&&xs),
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                    }))
                                                                                             })
                                                                         })))
    }
    pub fn Data_Array_ST_length() -> &dyn Any {
        static Data_Array_ST_length: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_length.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn1(),
                                                                              &&&PureScript_Data_Array_ST::Data_Array_ST_lengthImpl()))
    }
    pub fn Data_Array_ST_freeze() -> &dyn Any {
        static Data_Array_ST_freeze: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_freeze.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn1(),
                                                                              &&&PureScript_Data_Array_ST::Data_Array_ST_freezeImpl()))
    }
    pub fn Data_Array_ST_clone() -> &dyn Any {
        static Data_Array_ST_clone: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_clone.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn1(),
                                                                             &&&PureScript_Data_Array_ST::Data_Array_ST_cloneImpl()))
    }
}
