pub mod PureScript_Data_Functor_Invariant {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_3b72fe33::PureScript_Data_Monoid_Additive;
    use crate::module_ee682245::PureScript_Data_Monoid_Alternate;
    use crate::module_89cc47bd::PureScript_Data_Monoid_Conj;
    use crate::module_6a9d0e21::PureScript_Data_Monoid_Disj;
    use crate::module_5023b7e9::PureScript_Data_Monoid_Dual;
    use crate::module_f1079b5::PureScript_Data_Monoid_Endo;
    use crate::module_7d83a5e5::PureScript_Data_Monoid_Multiplicative;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Functor_Invariant_Invariantusd_Dict() -> &dyn Any {
        static Data_Functor_Invariant_Invariantusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Invariant_Invariantusd_Dict.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |x|
                                                                                 x.clone()))
    }
    pub fn Data_Functor_Invariant_invariantMultiplicative() -> &dyn Any {
        static Data_Functor_Invariant_invariantMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Invariant_invariantMultiplicative.get_or_init(||
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                                        &&&add(string("imap"),
                                                                                                               &&Func1::new(move
                                                                                                                                |f|
                                                                                                                                &Func1::new({
                                                                                                                                                let f
                                                                                                                                                    =
                                                                                                                                                    f.clone();
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
                                                                                                                                                                                Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                            let matchValue_1 =
                                                                                                                                                                                Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                            let matchValue_2 =
                                                                                                                                                                                Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_Multiplicative(),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                &&&matchValue_2))
                                                                                                                                                                        }
                                                                                                                                                                })
                                                                                                                                            })),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>())))
    }
    pub fn Data_Functor_Invariant_invariantEndo() -> &dyn Any {
        static Data_Functor_Invariant_invariantEndo: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Invariant_invariantEndo.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                              &&&add(string("imap"),
                                                                                                     &&Func1::new(move
                                                                                                                      |ab|
                                                                                                                      &Func1::new({
                                                                                                                                      let ab
                                                                                                                                          =
                                                                                                                                          ab.clone();
                                                                                                                                      move
                                                                                                                                          |ba|
                                                                                                                                          &Func1::new({
                                                                                                                                                          let ba
                                                                                                                                                              =
                                                                                                                                                              ba.clone();
                                                                                                                                                          move
                                                                                                                                                              |v|
                                                                                                                                                              {
                                                                                                                                                                  let matchValue =
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&ab);
                                                                                                                                                                  let matchValue_1 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&ba);
                                                                                                                                                                  let matchValue_2 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                            &&&matchValue_2),
                                                                                                                                                                                                                                                                         &&&matchValue_1)))
                                                                                                                                                              }
                                                                                                                                                      })
                                                                                                                                  })),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))
    }
    pub fn Data_Functor_Invariant_invariantDual() -> &dyn Any {
        static Data_Functor_Invariant_invariantDual: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Invariant_invariantDual.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                              &&&add(string("imap"),
                                                                                                     &&Func1::new(move
                                                                                                                      |f|
                                                                                                                      &Func1::new({
                                                                                                                                      let f
                                                                                                                                          =
                                                                                                                                          f.clone();
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
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                  let matchValue_1 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                  let matchValue_2 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                      &&&matchValue_2))
                                                                                                                                                              }
                                                                                                                                                      })
                                                                                                                                  })),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))
    }
    pub fn Data_Functor_Invariant_invariantDisj() -> &dyn Any {
        static Data_Functor_Invariant_invariantDisj: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Invariant_invariantDisj.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                              &&&add(string("imap"),
                                                                                                     &&Func1::new(move
                                                                                                                      |f|
                                                                                                                      &Func1::new({
                                                                                                                                      let f
                                                                                                                                          =
                                                                                                                                          f.clone();
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
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                  let matchValue_1 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                  let matchValue_2 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_Disj(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                      &&&matchValue_2))
                                                                                                                                                              }
                                                                                                                                                      })
                                                                                                                                  })),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))
    }
    pub fn Data_Functor_Invariant_invariantConj() -> &dyn Any {
        static Data_Functor_Invariant_invariantConj: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Invariant_invariantConj.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                              &&&add(string("imap"),
                                                                                                     &&Func1::new(move
                                                                                                                      |f|
                                                                                                                      &Func1::new({
                                                                                                                                      let f
                                                                                                                                          =
                                                                                                                                          f.clone();
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
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                  let matchValue_1 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                  let matchValue_2 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                      &&&matchValue_2))
                                                                                                                                                              }
                                                                                                                                                      })
                                                                                                                                  })),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))
    }
    pub fn Data_Functor_Invariant_invariantAdditive() -> &dyn Any {
        static Data_Functor_Invariant_invariantAdditive:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Invariant_invariantAdditive.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                                  &&&add(string("imap"),
                                                                                                         &&Func1::new(move
                                                                                                                          |f|
                                                                                                                          &Func1::new({
                                                                                                                                          let f
                                                                                                                                              =
                                                                                                                                              f.clone();
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
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                      let matchValue_1 =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                      let matchValue_2 =
                                                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive(),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                          &&&matchValue_2))
                                                                                                                                                                  }
                                                                                                                                                          })
                                                                                                                                      })),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Data_Functor_Invariant_imapF() -> &dyn Any {
        static Data_Functor_Invariant_imapF: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Invariant_imapF.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictFunctor|
                                                                     &Func1::new({
                                                                                     let dictFunctor
                                                                                         =
                                                                                         dictFunctor.clone();
                                                                                     move
                                                                                         |f|
                                                                                         &Func1::new({
                                                                                                         let f
                                                                                                             =
                                                                                                             f.clone();
                                                                                                         move
                                                                                                             |v|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                 &&&dictFunctor),
                                                                                                                                              &&&f)
                                                                                                     })
                                                                                 })))
    }
    pub fn Data_Functor_Invariant_invariantArray() -> &dyn Any {
        static Data_Functor_Invariant_invariantArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Invariant_invariantArray.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                               &&&add(string("imap"),
                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imapF(),
                                                                                                                                        &&&PureScript_Data_Functor::Data_Functor_functorArray()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))
    }
    pub fn Data_Functor_Invariant_invariantFn() -> &dyn Any {
        static Data_Functor_Invariant_invariantFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Invariant_invariantFn.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                            &&&add(string("imap"),
                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imapF(),
                                                                                                                                     &&&PureScript_Data_Functor::Data_Functor_functorFn()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>())))
    }
    pub fn Data_Functor_Invariant_imap() -> &dyn Any {
        static Data_Functor_Invariant_imap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Invariant_imap.get_or_init(||
                                                    &Func1::new(move |dict|
                                                                    find(string("imap"),
                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Functor_Invariant_invariantAlternate() -> &dyn Any {
        static Data_Functor_Invariant_invariantAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Invariant_invariantAlternate.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictInvariant|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                                                   &&&add(string("imap"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictInvariant
                                                                                                                                               =
                                                                                                                                               dictInvariant.clone();
                                                                                                                                           move
                                                                                                                                               |f|
                                                                                                                                               &Func1::new({
                                                                                                                                                               let f
                                                                                                                                                                   =
                                                                                                                                                                   f.clone();
                                                                                                                                                               move
                                                                                                                                                                   |g|
                                                                                                                                                                   &Func1::new({
                                                                                                                                                                                   let g
                                                                                                                                                                                       =
                                                                                                                                                                                       g.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |v|
                                                                                                                                                                                       {
                                                                                                                                                                                           let matchValue =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                           let matchValue_1 =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&g);
                                                                                                                                                                                           let matchValue_2 =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Alternate::Data_Monoid_Alternate_Alternate(),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imap(),
                                                                                                                                                                                                                                                                                                                                                                        &&&dictInvariant),
                                                                                                                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                                                                                                                  &&&matchValue_1),
                                                                                                                                                                                                                                                               &&&matchValue_2))
                                                                                                                                                                                       }
                                                                                                                                                                               })
                                                                                                                                                           })
                                                                                                                                       }),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))
    }
}
