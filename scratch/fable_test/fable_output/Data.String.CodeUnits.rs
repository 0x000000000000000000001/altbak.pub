pub mod PureScript_Data_String_CodeUnits {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Option_::getValue;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_b0fd79e4::PureScript_Data_String_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub mod Data_String_CodeUnits_FFI {
        use super::*;
        use fable_library_rust::List_::ofArray;
        use fable_library_rust::Map_::ofList;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::NativeArray_::Array;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_array;
        use fable_library_rust::NativeArray_::new_init;
        use fable_library_rust::String_::fromChars;
        use fable_library_rust::String_::getCharAt;
        use fable_library_rust::String_::indexOf2;
        use fable_library_rust::String_::indexOf3;
        use fable_library_rust::String_::lastIndexOf2;
        use fable_library_rust::String_::length as length_1;
        use fable_library_rust::String_::ofChar;
        use fable_library_rust::String_::substring;
        use fable_library_rust::String_::substring2;
        pub fn fromCharArray(a: &dyn Any) -> &dyn Any {
            let arr = Sharpurs_Prelude::unbox(a);
            let chars: Array<char> =
                new_init(&'\u{0000}', count(arr.clone()));
            for i in 0_i32..=count(arr.clone()) - 1_i32 {
                chars.get_mut()[i as usize] =
                    Sharpurs_Prelude::unbox(&arr[i].clone());
            }
            &fromChars(chars.clone())
        }
        pub fn toCharArray(str: &dyn Any) -> &dyn Any {
            let s: string = Sharpurs_Prelude::unbox(str);
            let arr = new_init(&defaultOf(), length_1(s.clone()));
            for i in 0_i32..=length_1(s.clone()) - 1_i32 {
                arr.get_mut()[i as usize] = &getCharAt(s.clone(), i);
            }
            &arr
        }
        pub fn singleton(c: &dyn Any) -> &dyn Any {
            &ofChar(Sharpurs_Prelude::unbox(c))
        }
        pub fn _charAt(just: &dyn Any, nothing: &dyn Any, idx: &dyn Any,
                       str: &dyn Any) -> &dyn Any {
            let i: i32 = Sharpurs_Prelude::unbox(idx);
            let s: string = Sharpurs_Prelude::unbox(str);
            if if i >= 0_i32 { i < length_1(s.clone()) } else { false } {
                Sharpurs_Prelude::sharpurs_apply(just, &&getCharAt(s, i))
            } else { nothing.clone() }
        }
        pub fn _toChar(just: &dyn Any, nothing: &dyn Any, str: &dyn Any)
         -> &dyn Any {
            let s: string = Sharpurs_Prelude::unbox(str);
            if length_1(s.clone()) == 1_i32 {
                Sharpurs_Prelude::sharpurs_apply(just, &&getCharAt(s, 0_i32))
            } else { nothing.clone() }
        }
        pub fn length(str: &dyn Any) -> &dyn Any {
            &length_1(Sharpurs_Prelude::unbox(str))
        }
        pub fn countPrefix(p: &dyn Any, str: &dyn Any) -> &dyn Any {
            let s: string = Sharpurs_Prelude::unbox(str);
            let i: MutCell<i32> = MutCell::new(0_i32);
            while if i.get() < length_1(s.clone()) {
                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(p,
                                                                                &&getCharAt(s.clone(),
                                                                                            i.get())))
                  } else { false } {
                i.set(i.get() + 1_i32);
            }
            &i.get()
        }
        pub fn _indexOf(just: &dyn Any, nothing: &dyn Any, x: &dyn Any,
                        s: &dyn Any) -> &dyn Any {
            let idx: i32 =
                indexOf2(Sharpurs_Prelude::unbox(s),
                         Sharpurs_Prelude::unbox(x), 4_i32);
            if idx == -1_i32 {
                nothing.clone()
            } else { Sharpurs_Prelude::sharpurs_apply(just, &&idx) }
        }
        pub fn _indexOfStartingAt(just: &dyn Any, nothing: &dyn Any,
                                  x: &dyn Any, startIdx: &dyn Any,
                                  str: &dyn Any) -> &dyn Any {
            let s: string = Sharpurs_Prelude::unbox(str);
            let sub: string = Sharpurs_Prelude::unbox(x);
            let start: i32 = Sharpurs_Prelude::unbox(startIdx);
            if if start < 0_i32 { true } else { start > length_1(s.clone()) }
               {
                nothing.clone()
            } else {
                let idx: i32 = indexOf3(s, sub, start, 4_i32);
                if idx == -1_i32 {
                    nothing.clone()
                } else { Sharpurs_Prelude::sharpurs_apply(just, &&idx) }
            }
        }
        pub fn _lastIndexOf(just: &dyn Any, nothing: &dyn Any, x: &dyn Any,
                            s: &dyn Any) -> &dyn Any {
            let idx: i32 =
                lastIndexOf2(Sharpurs_Prelude::unbox(s),
                             Sharpurs_Prelude::unbox(x), 4_i32);
            if idx == -1_i32 {
                nothing.clone()
            } else { Sharpurs_Prelude::sharpurs_apply(just, &&idx) }
        }
        pub fn _lastIndexOfStartingAt(just: &dyn Any, nothing: &dyn Any,
                                      x: &dyn Any, startIdx: &dyn Any,
                                      str: &dyn Any) -> &dyn Any {
            let s: string = Sharpurs_Prelude::unbox(str);
            let sub: string = Sharpurs_Prelude::unbox(x);
            let start: i32 = Sharpurs_Prelude::unbox(startIdx);
            let safeStart: i32 = 0_i32.max(length_1(s.clone()).min(start));
            let idx: i32 =
                lastIndexOf2(substring2(s.clone(), 0_i32,
                                        length_1(s).min(safeStart +
                                                            length_1(sub.clone()))),
                             sub, 4_i32);
            if idx == -1_i32 {
                nothing.clone()
            } else { Sharpurs_Prelude::sharpurs_apply(just, &&idx) }
        }
        pub fn take(idx: &dyn Any, str: &dyn Any) -> &dyn Any {
            let i: i32 = Sharpurs_Prelude::unbox(idx);
            let s: string = Sharpurs_Prelude::unbox(str);
            &substring2(s.clone(), 0_i32, 0_i32.max(length_1(s).min(i)))
        }
        pub fn drop_(idx: &dyn Any, str: &dyn Any) -> &dyn Any {
            let i: i32 = Sharpurs_Prelude::unbox(idx);
            let s: string = Sharpurs_Prelude::unbox(str);
            &substring(s.clone(), 0_i32.max(length_1(s).min(i)))
        }
        pub fn slice(start: &dyn Any, end_: &dyn Any, str: &dyn Any)
         -> &dyn Any {
            let s: string = Sharpurs_Prelude::unbox(str);
            let st: MutCell<i32> =
                MutCell::new(Sharpurs_Prelude::unbox(start));
            let en: MutCell<i32> =
                MutCell::new(Sharpurs_Prelude::unbox(end_));
            if st.get() < 0_i32 { st.set(length_1(s.clone()) + st.get()); }
            if en.get() < 0_i32 { en.set(length_1(s.clone()) + en.get()); }
            st.set(0_i32.max(length_1(s.clone()).min(st.get())));
            en.set(0_i32.max(length_1(s.clone()).min(en.get())));
            if st.get() > en.get() {
                &string("")
            } else { &substring2(s, st.get(), en.get() - st.get()) }
        }
        pub fn splitAt(idx: &dyn Any, str: &dyn Any) -> &dyn Any {
            let s: string = Sharpurs_Prelude::unbox(str);
            let i: i32 = Sharpurs_Prelude::unbox(idx);
            let safeI: i32 = 0_i32.max(length_1(s.clone()).min(i));
            &ofList(ofArray(new_array(&[LrcPtr::new((string("before"),
                                                     &substring2(s.clone(),
                                                                 0_i32,
                                                                 safeI))),
                                        LrcPtr::new((string("after"),
                                                     &substring(s,
                                                                safeI)))])))
        }
    }
    pub fn Data_String_CodeUnits__charAt() -> &dyn Any {
        static Data_String_CodeUnits__charAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits__charAt.get_or_init(||
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
                                                                                                            |idx|
                                                                                                            Func1::new({
                                                                                                                           let idx
                                                                                                                               =
                                                                                                                               idx.clone();
                                                                                                                           move
                                                                                                                               |str|
                                                                                                                               PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::_charAt(&just,
                                                                                                                                                                                                    &nothing,
                                                                                                                                                                                                    &idx,
                                                                                                                                                                                                    str)
                                                                                                                       })
                                                                                                    })
                                                                                 })))
    }
    pub fn Data_String_CodeUnits__indexOf() -> &dyn Any {
        static Data_String_CodeUnits__indexOf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits__indexOf.get_or_init(||
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
                                                                                                             |x|
                                                                                                             Func1::new({
                                                                                                                            let x
                                                                                                                                =
                                                                                                                                x.clone();
                                                                                                                            move
                                                                                                                                |s|
                                                                                                                                PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::_indexOf(&just,
                                                                                                                                                                                                      &nothing,
                                                                                                                                                                                                      &x,
                                                                                                                                                                                                      s)
                                                                                                                        })
                                                                                                     })
                                                                                  })))
    }
    pub fn Data_String_CodeUnits__indexOfStartingAt() -> &dyn Any {
        static Data_String_CodeUnits__indexOfStartingAt:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits__indexOfStartingAt.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |just|
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
                                                                                                                       |x|
                                                                                                                       Func1::new({
                                                                                                                                      let x
                                                                                                                                          =
                                                                                                                                          x.clone();
                                                                                                                                      move
                                                                                                                                          |startIdx|
                                                                                                                                          Func1::new({
                                                                                                                                                         let startIdx
                                                                                                                                                             =
                                                                                                                                                             startIdx.clone();
                                                                                                                                                         move
                                                                                                                                                             |str|
                                                                                                                                                             PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::_indexOfStartingAt(&just,
                                                                                                                                                                                                                                             &nothing,
                                                                                                                                                                                                                                             &x,
                                                                                                                                                                                                                                             &startIdx,
                                                                                                                                                                                                                                             str)
                                                                                                                                                     })
                                                                                                                                  })
                                                                                                               })
                                                                                            })))
    }
    pub fn Data_String_CodeUnits__lastIndexOf() -> &dyn Any {
        static Data_String_CodeUnits__lastIndexOf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits__lastIndexOf.get_or_init(||
                                                           &Func1::new(move
                                                                           |just|
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
                                                                                                                 |x|
                                                                                                                 Func1::new({
                                                                                                                                let x
                                                                                                                                    =
                                                                                                                                    x.clone();
                                                                                                                                move
                                                                                                                                    |s|
                                                                                                                                    PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::_lastIndexOf(&just,
                                                                                                                                                                                                              &nothing,
                                                                                                                                                                                                              &x,
                                                                                                                                                                                                              s)
                                                                                                                            })
                                                                                                         })
                                                                                      })))
    }
    pub fn Data_String_CodeUnits__lastIndexOfStartingAt() -> &dyn Any {
        static Data_String_CodeUnits__lastIndexOfStartingAt:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits__lastIndexOfStartingAt.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |just|
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
                                                                                                                           |x|
                                                                                                                           Func1::new({
                                                                                                                                          let x
                                                                                                                                              =
                                                                                                                                              x.clone();
                                                                                                                                          move
                                                                                                                                              |startIdx|
                                                                                                                                              Func1::new({
                                                                                                                                                             let startIdx
                                                                                                                                                                 =
                                                                                                                                                                 startIdx.clone();
                                                                                                                                                             move
                                                                                                                                                                 |str|
                                                                                                                                                                 PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::_lastIndexOfStartingAt(&just,
                                                                                                                                                                                                                                                     &nothing,
                                                                                                                                                                                                                                                     &x,
                                                                                                                                                                                                                                                     &startIdx,
                                                                                                                                                                                                                                                     str)
                                                                                                                                                         })
                                                                                                                                      })
                                                                                                                   })
                                                                                                })))
    }
    pub fn Data_String_CodeUnits__toChar() -> &dyn Any {
        static Data_String_CodeUnits__toChar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits__toChar.get_or_init(||
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
                                                                                                            |str|
                                                                                                            PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::_toChar(&just,
                                                                                                                                                                                 &nothing,
                                                                                                                                                                                 str)
                                                                                                    })
                                                                                 })))
    }
    pub fn Data_String_CodeUnits_countPrefix() -> &dyn Any {
        static Data_String_CodeUnits_countPrefix: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_countPrefix.get_or_init(||
                                                          &Func1::new(move |p|
                                                                          Func1::new({
                                                                                         let p
                                                                                             =
                                                                                             p.clone();
                                                                                         move
                                                                                             |str|
                                                                                             PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::countPrefix(&p,
                                                                                                                                                                      str)
                                                                                     })))
    }
    pub fn Data_String_CodeUnits_drop() -> &dyn Any {
        static Data_String_CodeUnits_drop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_drop.get_or_init(||
                                                   &Func1::new(move |idx|
                                                                   Func1::new({
                                                                                  let idx
                                                                                      =
                                                                                      idx.clone();
                                                                                  move
                                                                                      |str|
                                                                                      PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::drop_(&idx,
                                                                                                                                                         str)
                                                                              })))
    }
    pub fn Data_String_CodeUnits_fromCharArray() -> &dyn Any {
        static Data_String_CodeUnits_fromCharArray: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_CodeUnits_fromCharArray.get_or_init(||
                                                            &Func1::new(move
                                                                            |a|
                                                                            PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::fromCharArray(a)))
    }
    pub fn Data_String_CodeUnits_length() -> &dyn Any {
        static Data_String_CodeUnits_length: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_length.get_or_init(||
                                                     &Func1::new(move |str|
                                                                     PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::length(str)))
    }
    pub fn Data_String_CodeUnits_singleton() -> &dyn Any {
        static Data_String_CodeUnits_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_singleton.get_or_init(||
                                                        &Func1::new(move |c|
                                                                        PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::singleton(c)))
    }
    pub fn Data_String_CodeUnits_slice() -> &dyn Any {
        static Data_String_CodeUnits_slice: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_slice.get_or_init(||
                                                    &Func1::new(move |start|
                                                                    Func1::new({
                                                                                   let start
                                                                                       =
                                                                                       start.clone();
                                                                                   move
                                                                                       |end_|
                                                                                       Func1::new({
                                                                                                      let end_
                                                                                                          =
                                                                                                          end_.clone();
                                                                                                      move
                                                                                                          |str|
                                                                                                          PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::slice(&start,
                                                                                                                                                                             &end_,
                                                                                                                                                                             str)
                                                                                                  })
                                                                               })))
    }
    pub fn Data_String_CodeUnits_splitAt() -> &dyn Any {
        static Data_String_CodeUnits_splitAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_splitAt.get_or_init(||
                                                      &Func1::new(move |idx|
                                                                      Func1::new({
                                                                                     let idx
                                                                                         =
                                                                                         idx.clone();
                                                                                     move
                                                                                         |str|
                                                                                         PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::splitAt(&idx,
                                                                                                                                                              str)
                                                                                 })))
    }
    pub fn Data_String_CodeUnits_take() -> &dyn Any {
        static Data_String_CodeUnits_take: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_take.get_or_init(||
                                                   &Func1::new(move |idx|
                                                                   Func1::new({
                                                                                  let idx
                                                                                      =
                                                                                      idx.clone();
                                                                                  move
                                                                                      |str|
                                                                                      PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::take(&idx,
                                                                                                                                                        str)
                                                                              })))
    }
    pub fn Data_String_CodeUnits_toCharArray() -> &dyn Any {
        static Data_String_CodeUnits_toCharArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_toCharArray.get_or_init(||
                                                          &Func1::new(move
                                                                          |str|
                                                                          PureScript_Data_String_CodeUnits::Data_String_CodeUnits_FFI::toCharArray(str)))
    }
    pub fn Data_String_CodeUnits_zero() -> &dyn Any {
        static Data_String_CodeUnits_zero: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_zero.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                    &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()))
    }
    pub fn Data_String_CodeUnits_one() -> &dyn Any {
        static Data_String_CodeUnits_one: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_one.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()))
    }
    pub fn Data_String_CodeUnits_uncons() -> &dyn Any {
        static Data_String_CodeUnits_uncons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_uncons.get_or_init(||
                                                     &Func1::new(move |v|
                                                                     {
                                                                         let matchValue =
                                                                             Sharpurs_Prelude::unbox(v);
                                                                         match &Sharpurs_Prelude::_007cLitString_007c__007c(string(""),
                                                                                                                            &matchValue)
                                                                             {
                                                                             0_i32
                                                                             =>
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                             _
                                                                             =>
                                                                             {
                                                                                 let s =
                                                                                     matchValue;
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&add(string("head"),
                                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Unsafe::Data_String_Unsafe_charAt(),
                                                                                                                                                                                                                  &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_zero()),
                                                                                                                                                                               &&&s),
                                                                                                                                             add(string("tail"),
                                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_drop(),
                                                                                                                                                                                                                      &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_one()),
                                                                                                                                                                                   &&&s),
                                                                                                                                                 empty::<string,
                                                                                                                                                         &dyn Any>()))))
                                                                             }
                                                                         }
                                                                     }))
    }
    pub fn Data_String_CodeUnits_toChar() -> &dyn Any {
        static Data_String_CodeUnits_toChar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_toChar.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits__toChar(),
                                                                                                                         &&&Func1::new(move
                                                                                                                                           |usd__arg1|
                                                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                      &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_String_CodeUnits_takeWhile() -> &dyn Any {
        static Data_String_CodeUnits_takeWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_takeWhile.get_or_init(||
                                                        &Func1::new(move |p|
                                                                        &Func1::new({
                                                                                        let p
                                                                                            =
                                                                                            p.clone();
                                                                                        move
                                                                                            |s|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_take(),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_countPrefix(),
                                                                                                                                                                                                                                      &&&p),
                                                                                                                                                                                                   s)),
                                                                                                                             s)
                                                                                    })))
    }
    pub fn Data_String_CodeUnits_takeRight() -> &dyn Any {
        static Data_String_CodeUnits_takeRight: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_takeRight.get_or_init(||
                                                        &Func1::new(move |i|
                                                                        &Func1::new({
                                                                                        let i
                                                                                            =
                                                                                            i.clone();
                                                                                        move
                                                                                            |s|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_drop(),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_length(),
                                                                                                                                                                                                                                                                         s)),
                                                                                                                                                                                                   &&&i)),
                                                                                                                             s)
                                                                                    })))
    }
    pub fn Data_String_CodeUnits_stripSuffix() -> &dyn Any {
        static Data_String_CodeUnits_stripSuffix: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_stripSuffix.get_or_init(||
                                                          &Func1::new(move |v|
                                                                          &Func1::new({
                                                                                          let v
                                                                                              =
                                                                                              v.clone();
                                                                                          move
                                                                                              |str|
                                                                                              {
                                                                                                  let suffix =
                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                  let str1 =
                                                                                                      Sharpurs_Prelude::unbox(str);
                                                                                                  let matchValue_3 =
                                                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_splitAt(),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_length(),
                                                                                                                                                                                                                                                                                                             &&&str1)),
                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_length(),
                                                                                                                                                                                                                                                                          &&&suffix))),
                                                                                                                                                                 &&&str1));
                                                                                                  {
                                                                                                      let activePatternResult =
                                                                                                          Sharpurs_Prelude::_007cHasProp_007c__007c(string("before"),
                                                                                                                                                    &matchValue_3);
                                                                                                      if activePatternResult.is_some()
                                                                                                         {
                                                                                                          let activePatternResult_1 =
                                                                                                              Sharpurs_Prelude::_007cHasProp_007c__007c(string("after"),
                                                                                                                                                        &matchValue_3);
                                                                                                          if activePatternResult_1.is_some()
                                                                                                             {
                                                                                                              let after =
                                                                                                                  getValue(activePatternResult_1);
                                                                                                              let before =
                                                                                                                  getValue(activePatternResult);
                                                                                                              let matchValue_4 =
                                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Eq::Data_Eq_eqString()),
                                                                                                                                                                                                               &&&after),
                                                                                                                                                                            &&&suffix));
                                                                                                              match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                               &matchValue_4)
                                                                                                                  {
                                                                                                                  0_i32
                                                                                                                  =>
                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&before)),
                                                                                                                  _
                                                                                                                  =>
                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                              }
                                                                                                          } else {
                                                                                                              panic!("{}",
                                                                                                                     LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.String.CodeUnits.fs"),
                                  Data1: 143_i32,
                                  Data2: 599_i32,}).get_Message(),)
                                                                                                          }
                                                                                                      } else {
                                                                                                          panic!("{}",
                                                                                                                 LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.String.CodeUnits.fs"),
                                  Data1: 143_i32,
                                  Data2: 599_i32,}).get_Message(),)
                                                                                                      }
                                                                                                  }
                                                                                              }
                                                                                      })))
    }
    pub fn Data_String_CodeUnits_stripPrefix() -> &dyn Any {
        static Data_String_CodeUnits_stripPrefix: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_stripPrefix.get_or_init(||
                                                          &Func1::new(move |v|
                                                                          &Func1::new({
                                                                                          let v
                                                                                              =
                                                                                              v.clone();
                                                                                          move
                                                                                              |str|
                                                                                              {
                                                                                                  let matchValue =
                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                  let matchValue_1 =
                                                                                                      Sharpurs_Prelude::unbox(str);
                                                                                                  let prefix =
                                                                                                      matchValue;
                                                                                                  let matchValue_3 =
                                                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_splitAt(),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_length(),
                                                                                                                                                                                                                                       &&&prefix)),
                                                                                                                                                                 &&&matchValue_1));
                                                                                                  {
                                                                                                      let activePatternResult =
                                                                                                          Sharpurs_Prelude::_007cHasProp_007c__007c(string("before"),
                                                                                                                                                    &matchValue_3);
                                                                                                      if activePatternResult.is_some()
                                                                                                         {
                                                                                                          let activePatternResult_1 =
                                                                                                              Sharpurs_Prelude::_007cHasProp_007c__007c(string("after"),
                                                                                                                                                        &matchValue_3);
                                                                                                          if activePatternResult_1.is_some()
                                                                                                             {
                                                                                                              let after =
                                                                                                                  getValue(activePatternResult_1);
                                                                                                              let before =
                                                                                                                  getValue(activePatternResult);
                                                                                                              let matchValue_4 =
                                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Eq::Data_Eq_eqString()),
                                                                                                                                                                                                               &&&before),
                                                                                                                                                                            &&&prefix));
                                                                                                              match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                               &matchValue_4)
                                                                                                                  {
                                                                                                                  0_i32
                                                                                                                  =>
                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&after)),
                                                                                                                  _
                                                                                                                  =>
                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                              }
                                                                                                          } else {
                                                                                                              panic!("{}",
                                                                                                                     LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.String.CodeUnits.fs"),
                                  Data1: 145_i32,
                                  Data2: 377_i32,}).get_Message(),)
                                                                                                          }
                                                                                                      } else {
                                                                                                          panic!("{}",
                                                                                                                 LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.String.CodeUnits.fs"),
                                  Data1: 145_i32,
                                  Data2: 377_i32,}).get_Message(),)
                                                                                                      }
                                                                                                  }
                                                                                              }
                                                                                      })))
    }
    pub fn Data_String_CodeUnits_startsWith() -> &dyn Any {
        static Data_String_CodeUnits_startsWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_startsWith.get_or_init(||
                                                         &Func1::new(move
                                                                         |pat|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                             &&&PureScript_Data_Maybe::Data_Maybe_isJust()),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_stripPrefix(),
                                                                                                                                             pat))))
    }
    pub fn Data_String_CodeUnits_lastIndexOf_prime() -> &dyn Any {
        static Data_String_CodeUnits_lastIndexOf_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_lastIndexOf_prime.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits__lastIndexOfStartingAt(),
                                                                                                                                    &&&Func1::new(move
                                                                                                                                                      |usd__arg1|
                                                                                                                                                      &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                 &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_String_CodeUnits_lastIndexOf() -> &dyn Any {
        static Data_String_CodeUnits_lastIndexOf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_lastIndexOf.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits__lastIndexOf(),
                                                                                                                              &&&Func1::new(move
                                                                                                                                                |usd__arg1|
                                                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                           &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_String_CodeUnits_indexOf_prime() -> &dyn Any {
        static Data_String_CodeUnits_indexOf_prime: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_CodeUnits_indexOf_prime.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits__indexOfStartingAt(),
                                                                                                                                &&&Func1::new(move
                                                                                                                                                  |usd__arg1|
                                                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                             &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_String_CodeUnits_indexOf() -> &dyn Any {
        static Data_String_CodeUnits_indexOf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_indexOf.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits__indexOf(),
                                                                                                                          &&&Func1::new(move
                                                                                                                                            |usd__arg1|
                                                                                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                       &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_String_CodeUnits_endsWith() -> &dyn Any {
        static Data_String_CodeUnits_endsWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_endsWith.get_or_init(||
                                                       &Func1::new(move |pat|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_isJust()),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_stripSuffix(),
                                                                                                                                           pat))))
    }
    pub fn Data_String_CodeUnits_dropWhile() -> &dyn Any {
        static Data_String_CodeUnits_dropWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_dropWhile.get_or_init(||
                                                        &Func1::new(move |p|
                                                                        &Func1::new({
                                                                                        let p
                                                                                            =
                                                                                            p.clone();
                                                                                        move
                                                                                            |s|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_drop(),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_countPrefix(),
                                                                                                                                                                                                                                      &&&p),
                                                                                                                                                                                                   s)),
                                                                                                                             s)
                                                                                    })))
    }
    pub fn Data_String_CodeUnits_dropRight() -> &dyn Any {
        static Data_String_CodeUnits_dropRight: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_dropRight.get_or_init(||
                                                        &Func1::new(move |i|
                                                                        &Func1::new({
                                                                                        let i
                                                                                            =
                                                                                            i.clone();
                                                                                        move
                                                                                            |s|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_take(),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_length(),
                                                                                                                                                                                                                                                                         s)),
                                                                                                                                                                                                   &&&i)),
                                                                                                                             s)
                                                                                    })))
    }
    pub fn Data_String_CodeUnits_contains() -> &dyn Any {
        static Data_String_CodeUnits_contains: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_contains.get_or_init(||
                                                       &Func1::new(move |pat|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_isJust()),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_indexOf(),
                                                                                                                                           pat))))
    }
    pub fn Data_String_CodeUnits_charAt() -> &dyn Any {
        static Data_String_CodeUnits_charAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CodeUnits_charAt.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits__charAt(),
                                                                                                                         &&&Func1::new(move
                                                                                                                                           |usd__arg1|
                                                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                      &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
}
