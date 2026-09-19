pub mod PureScript_Control_Monad_Rec_Class {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_9d129c3::PureScript_Effect_Ref;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    #[derive(Clone, Debug,)]
    pub enum Control_Monad_Rec_Class_Step {
        Control_Monad_Rec_Class_Loopusd_Ctor(&dyn Any),
        Control_Monad_Rec_Class_Doneusd_Ctor(&dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Control_Monad_Rec_Class_Loop() -> &dyn Any {
        static Control_Monad_Rec_Class_Loop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_Loop.get_or_init(||
                                                     &Func1::new(move
                                                                     |usd__arg1|
                                                                     &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Control_Monad_Rec_Class_Done() -> &dyn Any {
        static Control_Monad_Rec_Class_Done: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_Done.get_or_init(||
                                                     &Func1::new(move
                                                                     |usd__arg1|
                                                                     &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Control_Monad_Rec_Class_MonadRecusd_Dict() -> &dyn Any {
        static Control_Monad_Rec_Class_MonadRecusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_MonadRecusd_Dict.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |x|
                                                                                 x.clone()))
    }
    pub fn Control_Monad_Rec_Class_tailRecM() -> &dyn Any {
        static Control_Monad_Rec_Class_tailRecM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_tailRecM.get_or_init(||
                                                         &Func1::new(move
                                                                         |dict|
                                                                         find(string("tailRecM"),
                                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Rec_Class_tailRecM2() -> &dyn Any {
        static Control_Monad_Rec_Class_tailRecM2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_tailRecM2.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonadRec|
                                                                          &Func1::new({
                                                                                          let dictMonadRec
                                                                                              =
                                                                                              dictMonadRec.clone();
                                                                                          move
                                                                                              |f|
                                                                                              &Func1::new({
                                                                                                              let f
                                                                                                                  =
                                                                                                                  f.clone();
                                                                                                              move
                                                                                                                  |a|
                                                                                                                  &Func1::new({
                                                                                                                                  let a
                                                                                                                                      =
                                                                                                                                      a.clone();
                                                                                                                                  move
                                                                                                                                      |b|
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                                             &&&dictMonadRec),
                                                                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                                                                            |o|
                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                &&find(string("a"),
                                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(o))),
                                                                                                                                                                                                                                                             &&find(string("b"),
                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(o))))),
                                                                                                                                                                       &&&add(string("a"),
                                                                                                                                                                              &&a,
                                                                                                                                                                              add(string("b"),
                                                                                                                                                                                  b.clone(),
                                                                                                                                                                                  empty::<string,
                                                                                                                                                                                          &dyn Any>())))
                                                                                                                              })
                                                                                                          })
                                                                                      })))
    }
    pub fn Control_Monad_Rec_Class_tailRecM3() -> &dyn Any {
        static Control_Monad_Rec_Class_tailRecM3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_tailRecM3.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonadRec|
                                                                          &Func1::new({
                                                                                          let dictMonadRec
                                                                                              =
                                                                                              dictMonadRec.clone();
                                                                                          move
                                                                                              |f|
                                                                                              &Func1::new({
                                                                                                              let f
                                                                                                                  =
                                                                                                                  f.clone();
                                                                                                              move
                                                                                                                  |a|
                                                                                                                  &Func1::new({
                                                                                                                                  let a
                                                                                                                                      =
                                                                                                                                      a.clone();
                                                                                                                                  move
                                                                                                                                      |b|
                                                                                                                                      &Func1::new({
                                                                                                                                                      let b
                                                                                                                                                          =
                                                                                                                                                          b.clone();
                                                                                                                                                      move
                                                                                                                                                          |c|
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                                                                 &&&dictMonadRec),
                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                |o|
                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                       &&find(string("a"),
                                                                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(o))),
                                                                                                                                                                                                                                                                                                                    &&find(string("b"),
                                                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(o))),
                                                                                                                                                                                                                                                                                 &&find(string("c"),
                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(o))))),
                                                                                                                                                                                           &&&add(string("a"),
                                                                                                                                                                                                  &&a,
                                                                                                                                                                                                  add(string("b"),
                                                                                                                                                                                                      &&b,
                                                                                                                                                                                                      add(string("c"),
                                                                                                                                                                                                          c.clone(),
                                                                                                                                                                                                          empty::<string,
                                                                                                                                                                                                                  &dyn Any>()))))
                                                                                                                                                  })
                                                                                                                              })
                                                                                                          })
                                                                                      })))
    }
    pub fn Control_Monad_Rec_Class_untilJust() -> &dyn Any {
        static Control_Monad_Rec_Class_untilJust: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_untilJust.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonadRec|
                                                                          {
                                                                              let Functor0 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                          Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                                                                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              &Func1::new({
                                                                                              let Functor0
                                                                                                  =
                                                                                                  Functor0.clone();
                                                                                              let dictMonadRec
                                                                                                  =
                                                                                                  dictMonadRec.clone();
                                                                                              move
                                                                                                  |m|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_applyFlipped(),
                                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                         &&&dictMonadRec),
                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                        let m
                                                                                                                                                                                            =
                                                                                                                                                                                            m.clone();
                                                                                                                                                                                        move
                                                                                                                                                                                            |v|
                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_mapFlipped(),
                                                                                                                                                                                                                                                                                                   &&&Functor0),
                                                                                                                                                                                                                                                                &&&m),
                                                                                                                                                                                                                             &&&Func1::new(move
                                                                                                                                                                                                                                               |v1|
                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                   let matchValue:
                                                                                                                                                                                                                                                           LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                                                   match matchValue.as_ref()
                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                       &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_1_0)),
                                                                                                                                                                                                                                                       _
                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                       &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                                                                                                                                                                   }
                                                                                                                                                                                                                                               }))
                                                                                                                                                                                    })))
                                                                                          })
                                                                          }))
    }
    pub fn Control_Monad_Rec_Class_whileJust() -> &dyn Any {
        static Control_Monad_Rec_Class_whileJust: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_whileJust.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonoid|
                                                                          {
                                                                              let Semigroup0 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                          Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              &Func1::new({
                                                                                              let Semigroup0
                                                                                                  =
                                                                                                  Semigroup0.clone();
                                                                                              let dictMonoid
                                                                                                  =
                                                                                                  dictMonoid.clone();
                                                                                              move
                                                                                                  |dictMonadRec|
                                                                                                  {
                                                                                                      let Functor0 =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                      &Func1::new({
                                                                                                                      let Functor0
                                                                                                                          =
                                                                                                                          Functor0.clone();
                                                                                                                      let dictMonadRec
                                                                                                                          =
                                                                                                                          dictMonadRec.clone();
                                                                                                                      move
                                                                                                                          |m|
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_applyFlipped(),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                 &&&dictMonoid)),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                                 &&&dictMonadRec),
                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                let m
                                                                                                                                                                                                                    =
                                                                                                                                                                                                                    m.clone();
                                                                                                                                                                                                                move
                                                                                                                                                                                                                    |v|
                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_mapFlipped(),
                                                                                                                                                                                                                                                                                                                           &&&Functor0),
                                                                                                                                                                                                                                                                                        &&&m),
                                                                                                                                                                                                                                                     &&&Func1::new({
                                                                                                                                                                                                                                                                       let v
                                                                                                                                                                                                                                                                           =
                                                                                                                                                                                                                                                                           v.clone();
                                                                                                                                                                                                                                                                       move
                                                                                                                                                                                                                                                                           |v1|
                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                               let matchValue:
                                                                                                                                                                                                                                                                                       LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                                                                                                   Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                                                                               match matchValue.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                         |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                             &&&Semigroup0),
                                                                                                                                                                                                                                                                                                                                                                                          &&&v),
                                                                                                                                                                                                                                                                                                                                                       &&matchValue_1_0)),
                                                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&v)),
                                                                                                                                                                                                                                                                               }
                                                                                                                                                                                                                                                                           }
                                                                                                                                                                                                                                                                   }))
                                                                                                                                                                                                            })))
                                                                                                                  })
                                                                                                  }
                                                                                          })
                                                                          }))
    }
    pub fn Control_Monad_Rec_Class_tailRec() -> &dyn Any {
        static Control_Monad_Rec_Class_tailRec: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_tailRec.get_or_init(||
                                                        &Func1::new(move |f|
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
                                                                                                                       go_tco(v.clone())
                                                                                                               })
                                                                                           });
                                                                            let go_1 =
                                                                                Lazy(go_2);
                                                                            let go_tco =
                                                                                Func1::new({
                                                                                               let f
                                                                                                   =
                                                                                                   f.clone();
                                                                                               move
                                                                                                   |v_1|
                                                                                                   {
                                                                                                       let v_1 =
                                                                                                           v_1.clone();
                                                                                                       '_go_tco:
                                                                                                           loop 
                                                                                                                {
                                                                                                               break
                                                                                                                   '_go_tco
                                                                                                                   ({
                                                                                                                        let matchValue:
                                                                                                                                LrcPtr<PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step> =
                                                                                                                            Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                        match matchValue.as_ref()
                                                                                                                            {
                                                                                                                            PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_1_0)
                                                                                                                            =>
                                                                                                                            matchValue_1_0,
                                                                                                                            PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(matchValue_0_0)
                                                                                                                            =>
                                                                                                                            {
                                                                                                                                let v_1_temp =
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                     &&matchValue_0_0);
                                                                                                                                v_1.set(v_1_temp);
                                                                                                                                continue
                                                                                                                                    '_go_tco

                                                                                                                            }
                                                                                                                        }
                                                                                                                    })
                                                                                                                   ;
                                                                                                           }
                                                                                                   }
                                                                                           });
                                                                            let go =
                                                                                go_1.Value;
                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                &&&go),
                                                                                                             f)
                                                                        }))
    }
    pub fn Control_Monad_Rec_Class_tailRec2() -> &dyn Any {
        static Control_Monad_Rec_Class_tailRec2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_tailRec2.get_or_init(||
                                                         &Func1::new(move |f|
                                                                         &Func1::new({
                                                                                         let f
                                                                                             =
                                                                                             f.clone();
                                                                                         move
                                                                                             |a|
                                                                                             &Func1::new({
                                                                                                             let a
                                                                                                                 =
                                                                                                                 a.clone();
                                                                                                             move
                                                                                                                 |b|
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRec(),
                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                       |o|
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                           &&find(string("a"),
                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(o))),
                                                                                                                                                                                                                                        &&find(string("b"),
                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(o))))),
                                                                                                                                                  &&&add(string("a"),
                                                                                                                                                         &&a,
                                                                                                                                                         add(string("b"),
                                                                                                                                                             b.clone(),
                                                                                                                                                             empty::<string,
                                                                                                                                                                     &dyn Any>())))
                                                                                                         })
                                                                                     })))
    }
    pub fn Control_Monad_Rec_Class_tailRec3() -> &dyn Any {
        static Control_Monad_Rec_Class_tailRec3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_tailRec3.get_or_init(||
                                                         &Func1::new(move |f|
                                                                         &Func1::new({
                                                                                         let f
                                                                                             =
                                                                                             f.clone();
                                                                                         move
                                                                                             |a|
                                                                                             &Func1::new({
                                                                                                             let a
                                                                                                                 =
                                                                                                                 a.clone();
                                                                                                             move
                                                                                                                 |b|
                                                                                                                 &Func1::new({
                                                                                                                                 let b
                                                                                                                                     =
                                                                                                                                     b.clone();
                                                                                                                                 move
                                                                                                                                     |c|
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRec(),
                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                           |o|
                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                  &&find(string("a"),
                                                                                                                                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(o))),
                                                                                                                                                                                                                                                                                               &&find(string("b"),
                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(o))),
                                                                                                                                                                                                                                                            &&find(string("c"),
                                                                                                                                                                                                                                                                   Sharpurs_Prelude::unbox(o))))),
                                                                                                                                                                      &&&add(string("a"),
                                                                                                                                                                             &&a,
                                                                                                                                                                             add(string("b"),
                                                                                                                                                                                 &&b,
                                                                                                                                                                                 add(string("c"),
                                                                                                                                                                                     c.clone(),
                                                                                                                                                                                     empty::<string,
                                                                                                                                                                                             &dyn Any>()))))
                                                                                                                             })
                                                                                                         })
                                                                                     })))
    }
    pub fn Control_Monad_Rec_Class_monadRecMaybe() -> &dyn Any {
        static Control_Monad_Rec_Class_monadRecMaybe:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_monadRecMaybe.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                               &&&add(string("tailRecM"),
                                                                                                      &&Func1::new(move
                                                                                                                       |f|
                                                                                                                       &Func1::new({
                                                                                                                                       let f
                                                                                                                                           =
                                                                                                                                           f.clone();
                                                                                                                                       move
                                                                                                                                           |a0|
                                                                                                                                           {
                                                                                                                                               let g =
                                                                                                                                                   &Func1::new(move
                                                                                                                                                                   |v|
                                                                                                                                                                   {
                                                                                                                                                                       let matchValue:
                                                                                                                                                                               LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                                                                       if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                              =
                                                                                                                                                                              matchValue.as_ref()
                                                                                                                                                                          {
                                                                                                                                                                           let activePatternResult:
                                                                                                                                                                                   LrcPtr<PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step> =
                                                                                                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                      {
                                                                                                                                                                                                                      Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                      _
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                  });
                                                                                                                                                                           if let PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(activePatternResult_0_0)
                                                                                                                                                                                  =
                                                                                                                                                                                  activePatternResult.as_ref()
                                                                                                                                                                              {
                                                                                                                                                                               &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                    &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                           PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                           _
                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                                                                                                                                                       })))
                                                                                                                                                                           } else {
                                                                                                                                                                               let activePatternResult_1:
                                                                                                                                                                                       LrcPtr<PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                   Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                          {
                                                                                                                                                                                                                          Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                          _
                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                                      });
                                                                                                                                                                               if let PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(activePatternResult_1_1_0)
                                                                                                                                                                                      =
                                                                                                                                                                                      activePatternResult_1.as_ref()
                                                                                                                                                                                  {
                                                                                                                                                                                   &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                                                   PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                                                                                                                                               }))))
                                                                                                                                                                               } else {
                                                                                                                                                                                   panic!("{}",
                                                                                                                                                                                          string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                                                                                                                                                                               }
                                                                                                                                                                           }
                                                                                                                                                                       } else {
                                                                                                                                                                           &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
                                                                                                                                                                       }
                                                                                                                                                                   });
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRec(),
                                                                                                                                                                                                                   &&&g),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                   a0))
                                                                                                                                           }
                                                                                                                                   })),
                                                                                                      add(string("Monad0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Maybe::Data_Maybe_monadMaybe()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Control_Monad_Rec_Class_monadRecIdentity() -> &dyn Any {
        static Control_Monad_Rec_Class_monadRecIdentity:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_monadRecIdentity.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                                  &&&add(string("tailRecM"),
                                                                                                         &&Func1::new(move
                                                                                                                          |f|
                                                                                                                          {
                                                                                                                              let runIdentity =
                                                                                                                                  &Func1::new(move
                                                                                                                                                  |v|
                                                                                                                                                  &Sharpurs_Prelude::unbox(v));
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                  &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRec(),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                        &&&runIdentity),
                                                                                                                                                                                                                                     f)))
                                                                                                                          }),
                                                                                                         add(string("Monad0"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused|
                                                                                                                              &PureScript_Data_Identity::Data_Identity_monadIdentity()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))
    }
    pub fn Control_Monad_Rec_Class_monadRecFunction() -> &dyn Any {
        static Control_Monad_Rec_Class_monadRecFunction:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_monadRecFunction.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                                  &&&add(string("tailRecM"),
                                                                                                         &&Func1::new(move
                                                                                                                          |f|
                                                                                                                          &Func1::new({
                                                                                                                                          let f
                                                                                                                                              =
                                                                                                                                              f.clone();
                                                                                                                                          move
                                                                                                                                              |a0|
                                                                                                                                              &Func1::new({
                                                                                                                                                              let a0
                                                                                                                                                                  =
                                                                                                                                                                  a0.clone();
                                                                                                                                                              move
                                                                                                                                                                  |e|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRec(),
                                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                                        let e
                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                            e.clone();
                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                            |a|
                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                a),
                                                                                                                                                                                                                                                                                             &&&e)
                                                                                                                                                                                                                                                    })),
                                                                                                                                                                                                   &&&a0)
                                                                                                                                                          })
                                                                                                                                      })),
                                                                                                         add(string("Monad0"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused|
                                                                                                                              &PureScript_Control_Monad::Control_Monad_monadFn()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))
    }
    pub fn Control_Monad_Rec_Class_monadRecEither() -> &dyn Any {
        static Control_Monad_Rec_Class_monadRecEither:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_monadRecEither.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                                &&&add(string("tailRecM"),
                                                                                                       &&Func1::new(move
                                                                                                                        |f|
                                                                                                                        &Func1::new({
                                                                                                                                        let f
                                                                                                                                            =
                                                                                                                                            f.clone();
                                                                                                                                        move
                                                                                                                                            |a0|
                                                                                                                                            {
                                                                                                                                                let g =
                                                                                                                                                    &Func1::new(move
                                                                                                                                                                    |v|
                                                                                                                                                                    {
                                                                                                                                                                        let matchValue:
                                                                                                                                                                                LrcPtr<Data_Either_Either> =
                                                                                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                                                                                        if let Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                                                               =
                                                                                                                                                                               matchValue.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                            let activePatternResult:
                                                                                                                                                                                    LrcPtr<PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                       {
                                                                                                                                                                                                                       Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                       _
                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                       unreachable!(),
                                                                                                                                                                                                                   });
                                                                                                                                                                            if let PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(activePatternResult_0_0)
                                                                                                                                                                                   =
                                                                                                                                                                                   activePatternResult.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                                &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                     &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                            PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                            _
                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                                                                                                                                                        })))
                                                                                                                                                                            } else {
                                                                                                                                                                                let activePatternResult_1:
                                                                                                                                                                                        LrcPtr<PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                    Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                           {
                                                                                                                                                                                                                           Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                           _
                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                                       });
                                                                                                                                                                                if let PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(activePatternResult_1_1_0)
                                                                                                                                                                                       =
                                                                                                                                                                                       activePatternResult_1.as_ref()
                                                                                                                                                                                   {
                                                                                                                                                                                    &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                        PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                                                                                                                                                    }))))
                                                                                                                                                                                } else {
                                                                                                                                                                                    panic!("{}",
                                                                                                                                                                                           string("Match failure: PureScript_Data_Either.Data_Either_Either"),)
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                        } else {
                                                                                                                                                                            &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                               Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                               _
                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                                                                                                                                                           }))))
                                                                                                                                                                        }
                                                                                                                                                                    });
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRec(),
                                                                                                                                                                                                                    &&&g),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                    a0))
                                                                                                                                            }
                                                                                                                                    })),
                                                                                                       add(string("Monad0"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused|
                                                                                                                            &PureScript_Data_Either::Data_Either_monadEither()),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>()))))
    }
    pub fn Control_Monad_Rec_Class_monadRecEffect() -> &dyn Any {
        static Control_Monad_Rec_Class_monadRecEffect:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_monadRecEffect.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                                &&&add(string("tailRecM"),
                                                                                                       &&Func1::new(move
                                                                                                                        |f|
                                                                                                                        &Func1::new({
                                                                                                                                        let f
                                                                                                                                            =
                                                                                                                                            f.clone();
                                                                                                                                        move
                                                                                                                                            |a|
                                                                                                                                            {
                                                                                                                                                let fromDone =
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                       |usd__unused|
                                                                                                                                                                                                       &Func1::new(move
                                                                                                                                                                                                                       |v|
                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&Func1::new({
                                                                                                                                                                                                                                                                          let v
                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                              v.clone();
                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                              |usd__unused_1|
                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                  let matchValue:
                                                                                                                                                                                                                                                                                          LrcPtr<PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                                                                                                  if let PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                         matchValue.as_ref()
                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                      &match matchValue.as_ref()
                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                           PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(x)
                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                           _
                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                                                                                                       }
                                                                                                                                                                                                                                                                                  } else {
                                                                                                                                                                                                                                                                                      panic!("{}",
                                                                                                                                                                                                                                                                                             string("Match failure"),)
                                                                                                                                                                                                                                                                                  }
                                                                                                                                                                                                                                                                              }
                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()))));
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                       &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bindFlipped(),
                                                                                                                                                                                                                                                                                                                             &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                                                                                                                          &&&PureScript_Effect_Ref::Effect_Ref_new()),
                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                          a))),
                                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                                   let fromDone
                                                                                                                                                                                                       =
                                                                                                                                                                                                       fromDone.clone();
                                                                                                                                                                                                   move
                                                                                                                                                                                                       |r|
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                              &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_untilE(),
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Ref::Effect_Ref_read(),
                                                                                                                                                                                                                                                                                                                                                                                                                       r)),
                                                                                                                                                                                                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                   let r
                                                                                                                                                                                                                                                                                                                                                                       =
                                                                                                                                                                                                                                                                                                                                                                       r.clone();
                                                                                                                                                                                                                                                                                                                                                                   move
                                                                                                                                                                                                                                                                                                                                                                       |v_1|
                                                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                                                           let matchValue_1:
                                                                                                                                                                                                                                                                                                                                                                                   LrcPtr<PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                                                                                                                                                                                           match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                               PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                                                                                                                                                                                                                                                                                                &&&true),
                                                                                                                                                                                                                                                                                                                                                                               PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(matchValue_1_0_0)
                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&matchValue_1_0_0)),
                                                                                                                                                                                                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                  |e|
                                                                                                                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Ref::Effect_Ref_write(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&r)),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |usd__unused_2|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&false))))),
                                                                                                                                                                                                                                                                                                                                                                           }
                                                                                                                                                                                                                                                                                                                                                                       }
                                                                                                                                                                                                                                                                                                                                                               })))),
                                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                                          let r
                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                              r.clone();
                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                              |usd__unused_3|
                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Effect::Effect_functorEffect()),
                                                                                                                                                                                                                                                                                                                                  &&&fromDone),
                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Ref::Effect_Ref_read(),
                                                                                                                                                                                                                                                                                                                                  &&&r))
                                                                                                                                                                                                                                                      }))
                                                                                                                                                                                               }))
                                                                                                                                            }
                                                                                                                                    })),
                                                                                                       add(string("Monad0"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused_4|
                                                                                                                            &PureScript_Effect::Effect_monadEffect()),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>()))))
    }
    pub fn Control_Monad_Rec_Class_loop3() -> &dyn Any {
        static Control_Monad_Rec_Class_loop3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_loop3.get_or_init(||
                                                      &Func1::new(move |a|
                                                                      &Func1::new({
                                                                                      let a
                                                                                          =
                                                                                          a.clone();
                                                                                      move
                                                                                          |b|
                                                                                          &Func1::new({
                                                                                                          let b
                                                                                                              =
                                                                                                              b.clone();
                                                                                                          move
                                                                                                              |c|
                                                                                                              &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&add(string("a"),
                                                                                                                                                                                                                                       &&a,
                                                                                                                                                                                                                                       add(string("b"),
                                                                                                                                                                                                                                           &&b,
                                                                                                                                                                                                                                           add(string("c"),
                                                                                                                                                                                                                                               c.clone(),
                                                                                                                                                                                                                                               empty::<string,
                                                                                                                                                                                                                                                       &dyn Any>())))))
                                                                                                      })
                                                                                  })))
    }
    pub fn Control_Monad_Rec_Class_loop2() -> &dyn Any {
        static Control_Monad_Rec_Class_loop2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_loop2.get_or_init(||
                                                      &Func1::new(move |a|
                                                                      &Func1::new({
                                                                                      let a
                                                                                          =
                                                                                          a.clone();
                                                                                      move
                                                                                          |b|
                                                                                          &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&add(string("a"),
                                                                                                                                                                                                                   &&a,
                                                                                                                                                                                                                   add(string("b"),
                                                                                                                                                                                                                       b.clone(),
                                                                                                                                                                                                                       empty::<string,
                                                                                                                                                                                                                               &dyn Any>()))))
                                                                                  })))
    }
    pub fn Control_Monad_Rec_Class_functorStep() -> &dyn Any {
        static Control_Monad_Rec_Class_functorStep: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Rec_Class_functorStep.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                             &&&add(string("map"),
                                                                                                    &&Func1::new(move
                                                                                                                     |f|
                                                                                                                     &Func1::new({
                                                                                                                                     let f
                                                                                                                                         =
                                                                                                                                         f.clone();
                                                                                                                                     move
                                                                                                                                         |m|
                                                                                                                                         {
                                                                                                                                             let matchValue:
                                                                                                                                                     LrcPtr<PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step> =
                                                                                                                                                 Sharpurs_Prelude::unbox(m);
                                                                                                                                             match matchValue.as_ref()
                                                                                                                                                 {
                                                                                                                                                 PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_1_0)
                                                                                                                                                 =>
                                                                                                                                                 &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                      &&matchValue_1_0))),
                                                                                                                                                 PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(matchValue_0_0)
                                                                                                                                                 =>
                                                                                                                                                 &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(matchValue_0_0)),
                                                                                                                                             }
                                                                                                                                         }
                                                                                                                                 })),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Control_Monad_Rec_Class_forever() -> &dyn Any {
        static Control_Monad_Rec_Class_forever: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_forever.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictMonadRec|
                                                                        {
                                                                            let Functor0 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                            &Func1::new({
                                                                                            let Functor0
                                                                                                =
                                                                                                Functor0.clone();
                                                                                            let dictMonadRec
                                                                                                =
                                                                                                dictMonadRec.clone();
                                                                                            move
                                                                                                |ma|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                       &&&dictMonadRec),
                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                      let ma
                                                                                                                                                                                          =
                                                                                                                                                                                          ma.clone();
                                                                                                                                                                                      move
                                                                                                                                                                                          |u|
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_voidRight(),
                                                                                                                                                                                                                                                                                                 &&&Functor0),
                                                                                                                                                                                                                                                              &&&LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(u.clone()))),
                                                                                                                                                                                                                           &&&ma)
                                                                                                                                                                                  })),
                                                                                                                                 &&&PureScript_Data_Unit::Data_Unit_unit())
                                                                                        })
                                                                        }))
    }
    pub fn Control_Monad_Rec_Class_bifunctorStep() -> &dyn Any {
        static Control_Monad_Rec_Class_bifunctorStep:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Rec_Class_bifunctorStep.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                               &&&add(string("bimap"),
                                                                                                      &&Func1::new(move
                                                                                                                       |v|
                                                                                                                       &Func1::new({
                                                                                                                                       let v
                                                                                                                                           =
                                                                                                                                           v.clone();
                                                                                                                                       move
                                                                                                                                           |v1|
                                                                                                                                           &Func1::new({
                                                                                                                                                           let v1
                                                                                                                                                               =
                                                                                                                                                               v1.clone();
                                                                                                                                                           move
                                                                                                                                                               |v2|
                                                                                                                                                               {
                                                                                                                                                                   let matchValue =
                                                                                                                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                   let matchValue_1 =
                                                                                                                                                                       Sharpurs_Prelude::unbox(&&v1);
                                                                                                                                                                   let matchValue_2:
                                                                                                                                                                           LrcPtr<PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step> =
                                                                                                                                                                       Sharpurs_Prelude::unbox(v2);
                                                                                                                                                                   match matchValue_2.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                       PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_2_1_0)
                                                                                                                                                                       =>
                                                                                                                                                                       &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                            &&matchValue_2_1_0))),
                                                                                                                                                                       PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(matchValue_2_0_0)
                                                                                                                                                                       =>
                                                                                                                                                                       &LrcPtr::new(PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                            &&matchValue_2_0_0))),
                                                                                                                                                                   }
                                                                                                                                                               }
                                                                                                                                                       })
                                                                                                                                   })),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))
    }
}
