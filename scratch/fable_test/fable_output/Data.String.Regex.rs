pub mod PureScript_Data_String_Regex {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_eea316d6::PureScript_Data_String_CodeUnits;
    use crate::module_91241446::PureScript_Data_String_Pattern;
    use crate::module_df7b54b2::PureScript_Data_String_Regex_Flags;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Data_String_Regex_FFI {
        use super::*;
        use fable_library_rust::Exception_::finally;
        use fable_library_rust::Exception_::try_catch;
        use fable_library_rust::Interfaces_::System::Collections::Generic::IEnumerator_1;
        use fable_library_rust::Interfaces_::System::IDisposable;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::NativeArray_::Array;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_init;
        use fable_library_rust::RegExp_::Match;
        use fable_library_rust::RegExp_::MatchCollection;
        use fable_library_rust::RegExp_::Regex;
        use fable_library_rust::String_::append;
        use fable_library_rust::System::Exception;
        #[derive(Clone, Debug, Default,)]
        pub struct RegexWrapper {
            source: string,
            flags: string,
            options: MutCell<i32>,
            globalFlag: MutCell<bool>,
            ignoreCase: MutCell<bool>,
            multiline: MutCell<bool>,
            dotAll: MutCell<bool>,
            sticky: MutCell<bool>,
            unicode: MutCell<bool>,
        }
        impl PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper
         {
            pub fn _ctor__Z384F8060(source_1: string, flags: string)
             ->
                 LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> {
                let source: string;
                let flags_1: string;
                let options: i32;
                let globalFlag: bool;
                let ignoreCase: bool;
                let multiline: bool;
                let dotAll: bool;
                let sticky: bool;
                let unicode: bool;
                ();
                source = source_1;
                flags_1 = flags;
                options = 0_i32;
                globalFlag = false;
                ignoreCase = false;
                multiline = false;
                dotAll = false;
                sticky = false;
                unicode = false;
                {
                    let inputSequence: string = flags_1.clone();
                    let enumerator: LrcPtr<dyn IEnumerator_1<_>> =
                        inputSequence.GetEnumerator();
                    {
                        finally(||
                                    if ((&enumerator) as
                                            &dyn Any).is::<LrcPtr<dyn IDisposable>>()
                                       {
                                        (&enumerator).Dispose();
                                    });
                        while enumerator.MoveNext() {
                            let c: char = enumerator.get_Current();
                            match &c {
                                'g' => globalFlag = true,
                                'i' => {
                                    options = options | 1_i32;
                                    ignoreCase = true
                                }
                                'm' => {
                                    options = options | 2_i32;
                                    multiline = true
                                }
                                's' => {
                                    options = options | 16_i32;
                                    dotAll = true
                                }
                                'u' => unicode = true,
                                'y' => sticky = true,
                                _ => (),
                            }
                        }
                    }
                }
                ();
                LrcPtr::new(PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper{source:
                                                                                                  source,
                                                                                              flags:
                                                                                                  flags_1,
                                                                                              options:
                                                                                                  MutCell::new(options),
                                                                                              globalFlag:
                                                                                                  MutCell::new(globalFlag),
                                                                                              ignoreCase:
                                                                                                  MutCell::new(ignoreCase),
                                                                                              multiline:
                                                                                                  MutCell::new(multiline),
                                                                                              dotAll:
                                                                                                  MutCell::new(dotAll),
                                                                                              sticky:
                                                                                                  MutCell::new(sticky),
                                                                                              unicode:
                                                                                                  MutCell::new(unicode),})
            }
            pub fn get_Regex(&self) -> LrcPtr<Regex> {
                Regex::new__sn(self.source.clone(),
                               self.options.get().clone())
            }
            pub fn get_Source(&self) -> string { self.source.clone() }
            pub fn get_Global(&self) -> bool { self.globalFlag.get().clone() }
            pub fn get_IgnoreCase(&self) -> bool {
                self.ignoreCase.get().clone()
            }
            pub fn get_Multiline(&self) -> bool {
                self.multiline.get().clone()
            }
            pub fn get_DotAll(&self) -> bool { self.dotAll.get().clone() }
            pub fn get_Sticky(&self) -> bool { self.sticky.get().clone() }
            pub fn get_Unicode(&self) -> bool { self.unicode.get().clone() }
            pub fn get_Flags(&self) -> string { self.flags.clone() }
        }
        impl core::fmt::Display for
         PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn regexImpl(left: &dyn Any, right: &dyn Any, s1: &dyn Any,
                         s2: &dyn Any) -> &dyn Any {
            let source_1: string = Sharpurs_Prelude::unbox(s1);
            let flags: string = Sharpurs_Prelude::unbox(s2);
            try_catch(||
                          {
                              let wrapper:
                                      LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> =
                                  PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper::_ctor__Z384F8060(source_1,
                                                                                                                      flags);
                              wrapper.get_Regex();
                              Sharpurs_Prelude::sharpurs_apply(right,
                                                               &&wrapper)
                          },
                      |e: LrcPtr<Exception>|
                          Sharpurs_Prelude::sharpurs_apply(left,
                                                           &&e.get_Message()))
        }
        pub fn showRegexImpl(r: &dyn Any) -> &dyn Any {
            let wrapper:
                    LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> =
                Sharpurs_Prelude::unbox(r);
            &append(append(append(string("/"), wrapper.get_Source()),
                           string("/")), wrapper.get_Flags())
        }
        pub fn source(r: &dyn Any) -> &dyn Any {
            &(Sharpurs_Prelude::unbox(r)).get_Source()
        }
        pub fn flagsImpl(r: &dyn Any) -> &dyn Any {
            let wrapper:
                    LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> =
                Sharpurs_Prelude::unbox(r);
            &{
                 let table_5 =
                     {
                         let table_4 =
                             {
                                 let table_3 =
                                     {
                                         let table_2 =
                                             {
                                                 let table_1 =
                                                     {
                                                         let table =
                                                             empty::<string,
                                                                     &dyn Any>();
                                                         add(string("global"),
                                                             &wrapper.get_Global(),
                                                             table)
                                                     };
                                                 add(string("ignoreCase"),
                                                     &wrapper.get_IgnoreCase(),
                                                     table_1)
                                             };
                                         add(string("multiline"),
                                             &wrapper.get_Multiline(),
                                             table_2)
                                     };
                                 add(string("dotAll"), &wrapper.get_DotAll(),
                                     table_3)
                             };
                         add(string("sticky"), &wrapper.get_Sticky(), table_4)
                     };
                 add(string("unicode"), &wrapper.get_Unicode(), table_5)
             }
        }
        pub fn test(r: &dyn Any, s: &dyn Any) -> &dyn Any {
            let wrapper:
                    LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> =
                Sharpurs_Prelude::unbox(r);
            let str: string = Sharpurs_Prelude::unbox(s);
            &wrapper.get_Regex().isMatch_s(str)
        }
        pub fn _match(just: &dyn Any, nothing: &dyn Any, r: &dyn Any,
                      s: &dyn Any) -> &dyn Any {
            let wrapper:
                    LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> =
                Sharpurs_Prelude::unbox(r);
            let str: string = Sharpurs_Prelude::unbox(s);
            if wrapper.get_Global() {
                let matches: LrcPtr<MatchCollection> =
                    wrapper.get_Regex().matches_s(str.clone());
                if matches.count() == 0_i32 {
                    nothing.clone()
                } else {
                    let arr = new_init(&defaultOf(), matches.count());
                    for i in 0_i32..=matches.count() - 1_i32 {
                        arr.get_mut()[i as usize] =
                            Sharpurs_Prelude::sharpurs_apply(just,
                                                             &&matches.item_n(i).value());
                    }
                    Sharpurs_Prelude::sharpurs_apply(just, &&arr)
                }
            } else {
                let m: Match = wrapper.get_Regex().match_s(str);
                if !m.success() {
                    nothing.clone()
                } else {
                    let arr_1 = new_init(&defaultOf(), m.groups().count());
                    for i_1 in 0_i32..=m.groups().count() - 1_i32 {
                        if m.groups().item_n(i_1).success() {
                            arr_1.get_mut()[i_1 as usize] =
                                Sharpurs_Prelude::sharpurs_apply(just,
                                                                 &&m.groups().item_n(i_1).value())
                        } else {
                            arr_1.get_mut()[i_1 as usize] = nothing.clone()
                        };
                    }
                    Sharpurs_Prelude::sharpurs_apply(just, &&arr_1)
                }
            }
        }
        pub fn replace(r: &dyn Any, s1: &dyn Any, s2: &dyn Any) -> &dyn Any {
            let wrapper:
                    LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> =
                Sharpurs_Prelude::unbox(r);
            let replacement: string = Sharpurs_Prelude::unbox(s1);
            let str: string = Sharpurs_Prelude::unbox(s2);
            if wrapper.get_Global() {
                &wrapper.get_Regex().replace_ss(str.clone(),
                                                replacement.clone())
            } else {
                &wrapper.get_Regex().replace_ssn(str, replacement, 1_i32)
            }
        }
        pub fn _replaceBy(just: &dyn Any, nothing: &dyn Any, r: &dyn Any,
                          f: &dyn Any, s: &dyn Any) -> &dyn Any {
            let wrapper:
                    LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> =
                Sharpurs_Prelude::unbox(r);
            let str: string = Sharpurs_Prelude::unbox(s);
            let evaluator =
                Func1::new({
                               let f = f.clone();
                               let just = just.clone();
                               let nothing = nothing.clone();
                               move |m: Match|
                                   {
                                       let res = f;
                                       res.set(Sharpurs_Prelude::sharpurs_apply(&res,
                                                                                &&m.groups().item_n(0_i32).value()));
                                       {
                                           let groupsArr =
                                               new_init(&defaultOf(),
                                                        m.groups().count() -
                                                            1_i32);
                                           for i in
                                               1_i32..=m.groups().count() -
                                                           1_i32 {
                                               if m.groups().item_n(i).success()
                                                  {
                                                   groupsArr.get_mut()[(i -
                                                                            1_i32)
                                                                           as
                                                                           usize]
                                                       =
                                                       Sharpurs_Prelude::sharpurs_apply(&just,
                                                                                        &&m.groups().item_n(i).value())
                                               } else {
                                                   groupsArr.get_mut()[(i -
                                                                            1_i32)
                                                                           as
                                                                           usize]
                                                       = nothing
                                               };
                                           }
                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&res,
                                                                                                     &&groupsArr))
                                       }
                                   }
                           });
            if wrapper.get_Global() {
                &wrapper.get_Regex().replace_sf(str.clone(),
                                                Func1::new({
                                                               let evaluator =
                                                                   evaluator.clone();
                                                               move
                                                                   |delegateArg:
                                                                        Match|
                                                                   evaluator(delegateArg)
                                                           }))
            } else {
                &wrapper.get_Regex().replace_sfn(str,
                                                 Func1::new({
                                                                let evaluator
                                                                    =
                                                                    evaluator.clone();
                                                                move
                                                                    |delegateArg_1:
                                                                         Match|
                                                                    evaluator(delegateArg_1)
                                                            }), 1_i32)
            }
        }
        pub fn _search(just: &dyn Any, nothing: &dyn Any, r: &dyn Any,
                       s: &dyn Any) -> &dyn Any {
            let wrapper:
                    LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> =
                Sharpurs_Prelude::unbox(r);
            let str: string = Sharpurs_Prelude::unbox(s);
            let m: Match = wrapper.get_Regex().match_s(str);
            if !m.success() {
                nothing.clone()
            } else { Sharpurs_Prelude::sharpurs_apply(just, &&m.index()) }
        }
        pub fn split(r: &dyn Any, s: &dyn Any) -> &dyn Any {
            let wrapper:
                    LrcPtr<PureScript_Data_String_Regex::Data_String_Regex_FFI::RegexWrapper> =
                Sharpurs_Prelude::unbox(r);
            let str: string = Sharpurs_Prelude::unbox(s);
            if if str.clone() == string("") {
                   wrapper.get_Source() == string("")
               } else { false } {
                &new_init(&defaultOf(), 0_i32)
            } else {
                let parts: Array<string> = wrapper.get_Regex().split_s(str);
                if if wrapper.get_Source() == string("") {
                       count(parts.clone()) >= 2_i32
                   } else { false } {
                    let len: i32 = count(parts.clone()) - 2_i32;
                    let arr = new_init(&defaultOf(), len);
                    for i in 0_i32..=len - 1_i32 {
                        arr.get_mut()[i as usize] = &parts[i + 1_i32].clone();
                    }
                    &arr
                } else {
                    let arr_1 = new_init(&defaultOf(), count(parts.clone()));
                    for i_1 in 0_i32..=count(parts.clone()) - 1_i32 {
                        arr_1.get_mut()[i_1 as usize] = &parts[i_1].clone();
                    }
                    &arr_1
                }
            }
        }
    }
    pub fn Data_String_Regex__match() -> &dyn Any {
        static Data_String_Regex__match: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex__match.get_or_init(||
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
                                                                                                       |r|
                                                                                                       Func1::new({
                                                                                                                      let r
                                                                                                                          =
                                                                                                                          r.clone();
                                                                                                                      move
                                                                                                                          |s|
                                                                                                                          PureScript_Data_String_Regex::Data_String_Regex_FFI::_match(&just,
                                                                                                                                                                                      &nothing,
                                                                                                                                                                                      &r,
                                                                                                                                                                                      s)
                                                                                                                  })
                                                                                               })
                                                                            })))
    }
    pub fn Data_String_Regex__replaceBy() -> &dyn Any {
        static Data_String_Regex__replaceBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex__replaceBy.get_or_init(||
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
                                                                                                           |r|
                                                                                                           Func1::new({
                                                                                                                          let r
                                                                                                                              =
                                                                                                                              r.clone();
                                                                                                                          move
                                                                                                                              |f|
                                                                                                                              Func1::new({
                                                                                                                                             let f
                                                                                                                                                 =
                                                                                                                                                 f.clone();
                                                                                                                                             move
                                                                                                                                                 |s|
                                                                                                                                                 PureScript_Data_String_Regex::Data_String_Regex_FFI::_replaceBy(&just,
                                                                                                                                                                                                                 &nothing,
                                                                                                                                                                                                                 &r,
                                                                                                                                                                                                                 &f,
                                                                                                                                                                                                                 s)
                                                                                                                                         })
                                                                                                                      })
                                                                                                   })
                                                                                })))
    }
    pub fn Data_String_Regex__search() -> &dyn Any {
        static Data_String_Regex__search: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex__search.get_or_init(||
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
                                                                                                        |r|
                                                                                                        Func1::new({
                                                                                                                       let r
                                                                                                                           =
                                                                                                                           r.clone();
                                                                                                                       move
                                                                                                                           |s|
                                                                                                                           PureScript_Data_String_Regex::Data_String_Regex_FFI::_search(&just,
                                                                                                                                                                                        &nothing,
                                                                                                                                                                                        &r,
                                                                                                                                                                                        s)
                                                                                                                   })
                                                                                                })
                                                                             })))
    }
    pub fn Data_String_Regex_flagsImpl() -> &dyn Any {
        static Data_String_Regex_flagsImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_flagsImpl.get_or_init(||
                                                    &Func1::new(move |r|
                                                                    PureScript_Data_String_Regex::Data_String_Regex_FFI::flagsImpl(r)))
    }
    pub fn Data_String_Regex_regexImpl() -> &dyn Any {
        static Data_String_Regex_regexImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_regexImpl.get_or_init(||
                                                    &Func1::new(move |left|
                                                                    Func1::new({
                                                                                   let left
                                                                                       =
                                                                                       left.clone();
                                                                                   move
                                                                                       |right|
                                                                                       Func1::new({
                                                                                                      let right
                                                                                                          =
                                                                                                          right.clone();
                                                                                                      move
                                                                                                          |s|
                                                                                                          Func1::new({
                                                                                                                         let s
                                                                                                                             =
                                                                                                                             s.clone();
                                                                                                                         move
                                                                                                                             |s_1|
                                                                                                                             PureScript_Data_String_Regex::Data_String_Regex_FFI::regexImpl(&left,
                                                                                                                                                                                            &right,
                                                                                                                                                                                            &s,
                                                                                                                                                                                            s_1)
                                                                                                                     })
                                                                                                  })
                                                                               })))
    }
    pub fn Data_String_Regex_replace() -> &dyn Any {
        static Data_String_Regex_replace: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_replace.get_or_init(||
                                                  &Func1::new(move |r|
                                                                  Func1::new({
                                                                                 let r
                                                                                     =
                                                                                     r.clone();
                                                                                 move
                                                                                     |s|
                                                                                     Func1::new({
                                                                                                    let s
                                                                                                        =
                                                                                                        s.clone();
                                                                                                    move
                                                                                                        |s_1|
                                                                                                        PureScript_Data_String_Regex::Data_String_Regex_FFI::replace(&r,
                                                                                                                                                                     &s,
                                                                                                                                                                     s_1)
                                                                                                })
                                                                             })))
    }
    pub fn Data_String_Regex_showRegexImpl() -> &dyn Any {
        static Data_String_Regex_showRegexImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_showRegexImpl.get_or_init(||
                                                        &Func1::new(move |r|
                                                                        PureScript_Data_String_Regex::Data_String_Regex_FFI::showRegexImpl(r)))
    }
    pub fn Data_String_Regex_source() -> &dyn Any {
        static Data_String_Regex_source: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_source.get_or_init(||
                                                 &Func1::new(move |r|
                                                                 PureScript_Data_String_Regex::Data_String_Regex_FFI::source(r)))
    }
    pub fn Data_String_Regex_split() -> &dyn Any {
        static Data_String_Regex_split: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_split.get_or_init(||
                                                &Func1::new(move |r|
                                                                Func1::new({
                                                                               let r
                                                                                   =
                                                                                   r.clone();
                                                                               move
                                                                                   |s|
                                                                                   PureScript_Data_String_Regex::Data_String_Regex_FFI::split(&r,
                                                                                                                                              s)
                                                                           })))
    }
    pub fn Data_String_Regex_test() -> &dyn Any {
        static Data_String_Regex_test: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_test.get_or_init(||
                                               &Func1::new(move |r|
                                                               Func1::new({
                                                                              let r
                                                                                  =
                                                                                  r.clone();
                                                                              move
                                                                                  |s|
                                                                                  PureScript_Data_String_Regex::Data_String_Regex_FFI::test(&r,
                                                                                                                                            s)
                                                                          })))
    }
    pub fn Data_String_Regex_showRegex() -> &dyn Any {
        static Data_String_Regex_showRegex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_showRegex.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                     &&&add(string("show"),
                                                                                            &&PureScript_Data_String_Regex::Data_String_Regex_showRegexImpl(),
                                                                                            empty::<string,
                                                                                                    &dyn Any>())))
    }
    pub fn Data_String_Regex_search() -> &dyn Any {
        static Data_String_Regex_search: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_search.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Regex::Data_String_Regex__search(),
                                                                                                                     &&&Func1::new(move
                                                                                                                                       |usd__arg1|
                                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                  &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_String_Regex_replace_prime() -> &dyn Any {
        static Data_String_Regex_replace_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_replace_prime.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Regex::Data_String_Regex__replaceBy(),
                                                                                                                            &&&Func1::new(move
                                                                                                                                              |usd__arg1|
                                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                         &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_String_Regex_renderFlags() -> &dyn Any {
        static Data_String_Regex_renderFlags: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_renderFlags.get_or_init(||
                                                      &Func1::new(move |v|
                                                                      {
                                                                          let f =
                                                                              Sharpurs_Prelude::unbox(v);
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                              &&{
                                                                                                                                                    let matchValue =
                                                                                                                                                        Sharpurs_Prelude::unbox(&find(string("global"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&f)));
                                                                                                                                                    match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                     &matchValue)
                                                                                                                                                        {
                                                                                                                                                        0_i32
                                                                                                                                                        =>
                                                                                                                                                        &string("g"),
                                                                                                                                                        _
                                                                                                                                                        =>
                                                                                                                                                        &string(""),
                                                                                                                                                    }
                                                                                                                                                }),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                    &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                 &&{
                                                                                                                                                                                       let matchValue_1 =
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&find(string("ignoreCase"),
                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&f)));
                                                                                                                                                                                       match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                        &matchValue_1)
                                                                                                                                                                                           {
                                                                                                                                                                                           0_i32
                                                                                                                                                                                           =>
                                                                                                                                                                                           &string("i"),
                                                                                                                                                                                           _
                                                                                                                                                                                           =>
                                                                                                                                                                                           &string(""),
                                                                                                                                                                                       }
                                                                                                                                                                                   }),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                       &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                    &&{
                                                                                                                                                                                                                          let matchValue_2 =
                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(&find(string("multiline"),
                                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&f)));
                                                                                                                                                                                                                          match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                           &matchValue_2)
                                                                                                                                                                                                                              {
                                                                                                                                                                                                                              0_i32
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              &string("m"),
                                                                                                                                                                                                                              _
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              &string(""),
                                                                                                                                                                                                                          }
                                                                                                                                                                                                                      }),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                       &&{
                                                                                                                                                                                                                                                             let matchValue_3 =
                                                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&find(string("dotAll"),
                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&f)));
                                                                                                                                                                                                                                                             match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                                                              &matchValue_3)
                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                 0_i32
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 &string("s"),
                                                                                                                                                                                                                                                                 _
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 &string(""),
                                                                                                                                                                                                                                                             }
                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                          &&{
                                                                                                                                                                                                                                                                                                let matchValue_4 =
                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&find(string("sticky"),
                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&f)));
                                                                                                                                                                                                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                                                                                                 &matchValue_4)
                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                    0_i32
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    &string("y"),
                                                                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    &string(""),
                                                                                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                       &&{
                                                                                                                                                                                                                                                             let matchValue_5 =
                                                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&find(string("unicode"),
                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&f)));
                                                                                                                                                                                                                                                             match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                                                              &matchValue_5)
                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                 0_i32
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 &string("u"),
                                                                                                                                                                                                                                                                 _
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 &string(""),
                                                                                                                                                                                                                                                             }
                                                                                                                                                                                                                                                         })))))
                                                                      }))
    }
    pub fn Data_String_Regex_regex() -> &dyn Any {
        static Data_String_Regex_regex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_regex.get_or_init(||
                                                &Func1::new(move |s|
                                                                &Func1::new({
                                                                                let s
                                                                                    =
                                                                                    s.clone();
                                                                                move
                                                                                    |f|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Regex::Data_String_Regex_regexImpl(),
                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                   |usd__arg1|
                                                                                                                                                                                                                                                                                   &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                |usd__arg1_1|
                                                                                                                                                                                                                                                &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                           &&&s)),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Regex::Data_String_Regex_renderFlags(),
                                                                                                                                                        f))
                                                                            })))
    }
    pub fn Data_String_Regex_parseFlags() -> &dyn Any {
        static Data_String_Regex_parseFlags: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_parseFlags.get_or_init(||
                                                     &Func1::new(move |s|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Regex_Flags::Data_String_Regex_Flags_RegexFlags(),
                                                                                                      &&&add(string("global"),
                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_contains(),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Pattern::Data_String_Pattern_Pattern(),
                                                                                                                                                                                                                     &&&string("g"))),
                                                                                                                                               s),
                                                                                                             add(string("ignoreCase"),
                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_contains(),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Pattern::Data_String_Pattern_Pattern(),
                                                                                                                                                                                                                         &&&string("i"))),
                                                                                                                                                   s),
                                                                                                                 add(string("multiline"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_contains(),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Pattern::Data_String_Pattern_Pattern(),
                                                                                                                                                                                                                             &&&string("m"))),
                                                                                                                                                       s),
                                                                                                                     add(string("dotAll"),
                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_contains(),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Pattern::Data_String_Pattern_Pattern(),
                                                                                                                                                                                                                                 &&&string("s"))),
                                                                                                                                                           s),
                                                                                                                         add(string("sticky"),
                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_contains(),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Pattern::Data_String_Pattern_Pattern(),
                                                                                                                                                                                                                                     &&&string("y"))),
                                                                                                                                                               s),
                                                                                                                             add(string("unicode"),
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_contains(),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Pattern::Data_String_Pattern_Pattern(),
                                                                                                                                                                                                                                         &&&string("u"))),
                                                                                                                                                                   s),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>())))))))))
    }
    pub fn Data_String_Regex_match() -> &dyn Any {
        static Data_String_Regex_match: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_match.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Regex::Data_String_Regex__match(),
                                                                                                                    &&&Func1::new(move
                                                                                                                                      |usd__arg1|
                                                                                                                                      &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                 &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_String_Regex_flags() -> &dyn Any {
        static Data_String_Regex_flags: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_flags.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                    &&&PureScript_Data_String_Regex_Flags::Data_String_Regex_Flags_RegexFlags()),
                                                                                 &&&PureScript_Data_String_Regex::Data_String_Regex_flagsImpl()))
    }
}
