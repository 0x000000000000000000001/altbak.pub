pub mod PureScript_Control_Monad_List_Trans {
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
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step;
    use crate::module_8aa70462::PureScript_Control_Monad_ST_Class;
    use crate::module_f5fe307f::PureScript_Control_Monad_Trans_Class;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_53a2e11::PureScript_Control_MonadPlus;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_720e12db::PureScript_Data_Lazy;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_b1754f14::PureScript_Data_Unfoldable1;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_21d6b3bc::PureScript_Effect_Class;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    use fable_library_rust::System::Lazy_1;
    #[derive(Clone, Debug,)]
    pub enum Control_Monad_List_Trans_Step {
        Control_Monad_List_Trans_Yieldusd_Ctor(&dyn Any, &dyn Any),
        Control_Monad_List_Trans_Skipusd_Ctor(&dyn Any),
        Control_Monad_List_Trans_Doneusd_Ctor,
    }
    impl core::fmt::Display for
     PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Control_Monad_List_Trans_identity() -> &dyn Any {
        static Control_Monad_List_Trans_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_identity.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                           &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Control_Monad_List_Trans_identity1() -> &dyn Any {
        static Control_Monad_List_Trans_identity1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_identity1.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                            &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Control_Monad_List_Trans_Yield() -> &dyn Any {
        static Control_Monad_List_Trans_Yield: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_Yield.get_or_init(||
                                                       &Func1::new(move
                                                                       |usd__arg1|
                                                                       Func1::new({
                                                                                      let usd__arg1
                                                                                          =
                                                                                          usd__arg1.clone();
                                                                                      move
                                                                                          |usd__arg2|
                                                                                          &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(usd__arg1,
                                                                                                                                                                                                                  usd__arg2.clone()))
                                                                                  })))
    }
    pub fn Control_Monad_List_Trans_Skip() -> &dyn Any {
        static Control_Monad_List_Trans_Skip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_Skip.get_or_init(||
                                                      &Func1::new(move
                                                                      |usd__arg1|
                                                                      &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Control_Monad_List_Trans_Done() -> &dyn Any {
        static Control_Monad_List_Trans_Done: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_Done.get_or_init(||
                                                      &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor))
    }
    pub fn Control_Monad_List_Trans_ListT() -> &dyn Any {
        static Control_Monad_List_Trans_ListT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_ListT.get_or_init(||
                                                       &Func1::new(move |x|
                                                                       x.clone()))
    }
    pub fn Control_Monad_List_Trans_wrapLazy() -> &dyn Any {
        static Control_Monad_List_Trans_wrapLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_wrapLazy.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictApplicative|
                                                                          &Func1::new({
                                                                                          let dictApplicative
                                                                                              =
                                                                                              dictApplicative.clone();
                                                                                          move
                                                                                              |v|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                  &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_ListT()),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                     &&&dictApplicative),
                                                                                                                                                                  &&&LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(v.clone()))))
                                                                                      })))
    }
    pub fn Control_Monad_List_Trans_wrapEffect() -> &dyn Any {
        static Control_Monad_List_Trans_wrapEffect: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_List_Trans_wrapEffect.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFunctor|
                                                                            &Func1::new({
                                                                                            let dictFunctor
                                                                                                =
                                                                                                dictFunctor.clone();
                                                                                            move
                                                                                                |v|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                    &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_ListT()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                          &&&dictFunctor),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                             &&&Func1::new(move
                                                                                                                                                                                                                                                                                               |usd__arg1|
                                                                                                                                                                                                                                                                                               &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Lazy::Data_Lazy_defer()),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_Function::Data_Function_const()))),
                                                                                                                                                                    v))
                                                                                        })))
    }
    pub fn Control_Monad_List_Trans_unfold_004029() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_unfold_tco(dictMonad))
    }
    pub fn Control_Monad_List_Trans_unfold_004029_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_unfold_004029_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_unfold_004029_002d1.get_or_init(||
                                                                     Lazy(Control_Monad_List_Trans_unfold_004029.clone()))
    }
    pub fn Control_Monad_List_Trans_unfold_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Control_Monad_List_Trans_unfold_004029_002d1 =
                            Control_Monad_List_Trans_unfold_004029_002d1.clone();
                        let Functor0 = Functor0.clone();
                        let dictMonad = dictMonad.clone();
                        move |f|
                            &Func1::new({
                                            let Control_Monad_List_Trans_unfold_004029_002d1
                                                =
                                                Control_Monad_List_Trans_unfold_004029_002d1.clone();
                                            let f = f.clone();
                                            move |z|
                                                {
                                                    let g =
                                                        &Func1::new({
                                                                        let Control_Monad_List_Trans_unfold_004029_002d1
                                                                            =
                                                                            Control_Monad_List_Trans_unfold_004029_002d1.clone();
                                                                        move
                                                                            |v|
                                                                            {
                                                                                let matchValue:
                                                                                        LrcPtr<Data_Maybe_Maybe> =
                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                match matchValue.as_ref()
                                                                                    {
                                                                                    Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                    =>
                                                                                    &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                                                                    Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                    =>
                                                                                    {
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                                   _
                                                                                                                                   =>
                                                                                                                                   unreachable!(),
                                                                                                                               });
                                                                                        &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                 },
                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                                                                                                   let Control_Monad_List_Trans_unfold_004029_002d1
                                                                                                                                                                                                                                                                       =
                                                                                                                                                                                                                                                                       Control_Monad_List_Trans_unfold_004029_002d1.clone();
                                                                                                                                                                                                                                                                   let activePatternResult
                                                                                                                                                                                                                                                                       =
                                                                                                                                                                                                                                                                       activePatternResult.clone();
                                                                                                                                                                                                                                                                   move
                                                                                                                                                                                                                                                                       |v1|
                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_unfold_004029_002d1.Value,
                                                                                                                                                                                                                                                                                                                                                                              &&&dictMonad),
                                                                                                                                                                                                                                                                                                                                           &&&f),
                                                                                                                                                                                                                                                                                                        &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                           })
                                                                                                                                                                                                                                                               }))))
                                                                                    }
                                                                                }
                                                                            }
                                                                    });
                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                        &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_ListT()),
                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                              &&&Functor0),
                                                                                                                                                           &&&g),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                           z)))
                                                }
                                        })
                    })
    }
    pub fn Control_Monad_List_Trans_unfold() -> &dyn Any {
        static Control_Monad_List_Trans_unfold: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_unfold.get_or_init(||
                                                        Control_Monad_List_Trans_unfold_004029_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_uncons_004033() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_uncons_tco(dictMonad))
    }
    pub fn Control_Monad_List_Trans_uncons_004033_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_uncons_004033_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_uncons_004033_002d1.get_or_init(||
                                                                     Lazy(Control_Monad_List_Trans_uncons_004033.clone()))
    }
    pub fn Control_Monad_List_Trans_uncons_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let pure_var =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                       Sharpurs_Prelude::unbox(dictMonad)),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Applicative0 = Applicative0.clone();
                        let Bind1 = Bind1.clone();
                        let Control_Monad_List_Trans_uncons_004033_002d1 =
                            Control_Monad_List_Trans_uncons_004033_002d1.clone();
                        let dictMonad = dictMonad.clone();
                        let pure_var = pure_var.clone();
                        move |v|
                            {
                                let l = Sharpurs_Prelude::unbox(v);
                                let g =
                                    &Func1::new({
                                                    let Control_Monad_List_Trans_uncons_004033_002d1
                                                        =
                                                        Control_Monad_List_Trans_uncons_004033_002d1.clone();
                                                    move |v1|
                                                        {
                                                            let matchValue:
                                                                    LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                                                Sharpurs_Prelude::unbox(v1);
                                                            match matchValue.as_ref()
                                                                {
                                                                PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_1_0)
                                                                =>
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_uncons_004033_002d1.Value,
                                                                                                                                    &&&dictMonad),
                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                    &&matchValue_1_0)),
                                                                PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                                                =>
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                    &&&Applicative0),
                                                                                                 &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_0_0,
                                                                                                                                                                           matchValue_0_1)
                                                                =>
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                    &&&pure_var),
                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                         |usd__arg1|
                                                                                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                    &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(matchValue_0_0,
                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                               &&matchValue_0_1))))),
                                                            }
                                                        }
                                                });
                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                       &&&Bind1),
                                                                                                    &&&l),
                                                                 &&&g)
                            }
                    })
    }
    pub fn Control_Monad_List_Trans_uncons() -> &dyn Any {
        static Control_Monad_List_Trans_uncons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_uncons.get_or_init(||
                                                        Control_Monad_List_Trans_uncons_004033_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_tail() -> &dyn Any {
        static Control_Monad_List_Trans_tail: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_tail.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonad|
                                                                      {
                                                                          let Functor0 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                          &Func1::new({
                                                                                          let Functor0
                                                                                              =
                                                                                              Functor0.clone();
                                                                                          let dictMonad
                                                                                              =
                                                                                              dictMonad.clone();
                                                                                          move
                                                                                              |l|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                     &&&Functor0),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                        &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                     &&&PureScript_Data_Tuple::Data_Tuple_snd())),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_uncons(),
                                                                                                                                                                                                     &&&dictMonad),
                                                                                                                                                                  l))
                                                                                      })
                                                                      }))
    }
    pub fn Control_Monad_List_Trans_stepMap() -> &dyn Any {
        static Control_Monad_List_Trans_stepMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_stepMap.get_or_init(||
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
                                                                                                                 {
                                                                                                                     let matchValue =
                                                                                                                         Sharpurs_Prelude::unbox(&&f);
                                                                                                                     let matchValue_1 =
                                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                         &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_ListT()),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                               &&&dictFunctor),
                                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                                         &&&matchValue_1))
                                                                                                                 }
                                                                                                         })
                                                                                     })))
    }
    pub fn Control_Monad_List_Trans_takeWhile_004041() -> &dyn Any {
        &Func1::new(move |dictApplicative|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_takeWhile_tco(dictApplicative))
    }
    pub fn Control_Monad_List_Trans_takeWhile_004041_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_takeWhile_004041_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_takeWhile_004041_002d1.get_or_init(||
                                                                        Lazy(Control_Monad_List_Trans_takeWhile_004041.clone()))
    }
    pub fn Control_Monad_List_Trans_takeWhile_tco(dictApplicative: &dyn Any)
     -> &dyn Any {
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Control_Monad_List_Trans_takeWhile_004041_002d1 =
                            Control_Monad_List_Trans_takeWhile_004041_002d1.clone();
                        let Functor0 = Functor0.clone();
                        let dictApplicative = dictApplicative.clone();
                        move |f|
                            {
                                let g =
                                    &Func1::new({
                                                    let Control_Monad_List_Trans_takeWhile_004041_002d1
                                                        =
                                                        Control_Monad_List_Trans_takeWhile_004041_002d1.clone();
                                                    let f = f.clone();
                                                    move |v|
                                                        {
                                                            let matchValue:
                                                                    LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                                                Sharpurs_Prelude::unbox(v);
                                                            match matchValue.as_ref()
                                                                {
                                                                PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_1_0)
                                                                =>
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                    &&&Func1::new(move
                                                                                                                                                      |usd__arg1|
                                                                                                                                                      &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(usd__arg1.clone())))),
                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                          &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_takeWhile_004041_002d1.Value,
                                                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                                                          &&&f)),
                                                                                                                                    &&matchValue_1_0)),
                                                                PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                                                =>
                                                                &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                                                PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_0_0,
                                                                                                                                                                           matchValue_0_1)
                                                                => {
                                                                    let a =
                                                                        matchValue_0_0.clone();
                                                                    let matchValue_1 =
                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                  &&&a));
                                                                    match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                     &matchValue_1)
                                                                        {
                                                                        0_i32
                                                                        =>
                                                                        &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(&a,
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_takeWhile_004041_002d1.Value,
                                                                                                                                                                                                                                                                                                                                          &&&dictApplicative),
                                                                                                                                                                                                                                                                                                       &&&f)),
                                                                                                                                                                                                                                 &&matchValue_0_1))),
                                                                        _ =>
                                                                        &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                                                    }
                                                                }
                                                            }
                                                        }
                                                });
                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_stepMap(),
                                                                                                    &&&Functor0),
                                                                 &&&g)
                            }
                    })
    }
    pub fn Control_Monad_List_Trans_takeWhile() -> &dyn Any {
        static Control_Monad_List_Trans_takeWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_takeWhile.get_or_init(||
                                                           Control_Monad_List_Trans_takeWhile_004041_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_scanl() -> &dyn Any {
        static Control_Monad_List_Trans_scanl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_scanl.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonad|
                                                                       {
                                                                           let Functor0 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                           &Func1::new({
                                                                                           let Functor0
                                                                                               =
                                                                                               Functor0.clone();
                                                                                           let dictMonad
                                                                                               =
                                                                                               dictMonad.clone();
                                                                                           move
                                                                                               |f|
                                                                                               &Func1::new({
                                                                                                               let f
                                                                                                                   =
                                                                                                                   f.clone();
                                                                                                               move
                                                                                                                   |b|
                                                                                                                   &Func1::new({
                                                                                                                                   let b
                                                                                                                                       =
                                                                                                                                       b.clone();
                                                                                                                                   move
                                                                                                                                       |l|
                                                                                                                                       {
                                                                                                                                           let g =
                                                                                                                                               &Func1::new(move
                                                                                                                                                               |v|
                                                                                                                                                               {
                                                                                                                                                                   let matchValue:
                                                                                                                                                                           LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                                   let b_prime =
                                                                                                                                                                       match matchValue.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                      _)
                                                                                                                                                                           =>
                                                                                                                                                                           x.clone(),
                                                                                                                                                                       };
                                                                                                                                                                   let h =
                                                                                                                                                                       &Func1::new({
                                                                                                                                                                                       let b_prime
                                                                                                                                                                                           =
                                                                                                                                                                                           b_prime.clone();
                                                                                                                                                                                       move
                                                                                                                                                                                           |v1|
                                                                                                                                                                                           {
                                                                                                                                                                                               let matchValue_1:
                                                                                                                                                                                                       LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                               match matchValue_1.as_ref()
                                                                                                                                                                                                   {
                                                                                                                                                                                                   PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                         |usd__arg1_1|
                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                                                                    &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b_prime,
                                                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                                                                                                       &&matchValue_1_1_0))),
                                                                                                                                                                                                                                                                                              &b_prime))),
                                                                                                                                                                                                   PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                                                   PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_1_0_0,
                                                                                                                                                                                                                                                                                                              matchValue_1_0_1)
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   {
                                                                                                                                                                                                       let b_prime_prime =
                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                               &&&b_prime),
                                                                                                                                                                                                                                            &&matchValue_1_0_0);
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                                             |usd__arg1|
                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                        &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b_prime_prime,
                                                                                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                                                                                                           &&matchValue_1_0_1))),
                                                                                                                                                                                                                                                                                                  &b_prime)))
                                                                                                                                                                                                   }
                                                                                                                                                                                               }
                                                                                                                                                                                           }
                                                                                                                                                                                   });
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                                                                                                       &&&h),
                                                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                                                           {
                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                           =>
                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                       })
                                                                                                                                                               });
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_unfold(),
                                                                                                                                                                                                                                                  &&&dictMonad),
                                                                                                                                                                                                               &&&g),
                                                                                                                                                                            &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                      l.clone())))
                                                                                                                                       }
                                                                                                                               })
                                                                                                           })
                                                                                       })
                                                                       }))
    }
    pub fn Control_Monad_List_Trans_prepend_prime() -> &dyn Any {
        static Control_Monad_List_Trans_prepend_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_prepend_prime.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictApplicative|
                                                                               &Func1::new({
                                                                                               let dictApplicative
                                                                                                   =
                                                                                                   dictApplicative.clone();
                                                                                               move
                                                                                                   |h|
                                                                                                   &Func1::new({
                                                                                                                   let h
                                                                                                                       =
                                                                                                                       h.clone();
                                                                                                                   move
                                                                                                                       |t|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                           &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_ListT()),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                              &&&dictApplicative),
                                                                                                                                                                                           &&&LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(&h,
                                                                                                                                                                                                                                                                                                                     t.clone()))))
                                                                                                               })
                                                                                           })))
    }
    pub fn Control_Monad_List_Trans_prepend() -> &dyn Any {
        static Control_Monad_List_Trans_prepend: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_prepend.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictApplicative|
                                                                         &Func1::new({
                                                                                         let dictApplicative
                                                                                             =
                                                                                             dictApplicative.clone();
                                                                                         move
                                                                                             |h|
                                                                                             &Func1::new({
                                                                                                             let h
                                                                                                                 =
                                                                                                                 h.clone();
                                                                                                             move
                                                                                                                 |t|
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_prepend_prime(),
                                                                                                                                                                                                                        &&&dictApplicative),
                                                                                                                                                                                     &&&h),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                        &&&PureScript_Data_Lazy::Data_Lazy_defer()),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                        t)))
                                                                                                         })
                                                                                     })))
    }
    pub fn Control_Monad_List_Trans_nil() -> &dyn Any {
        static Control_Monad_List_Trans_nil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_nil.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictApplicative|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                         &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_ListT()),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                            dictApplicative),
                                                                                                                                         &&&LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor)))))
    }
    pub fn Control_Monad_List_Trans_singleton() -> &dyn Any {
        static Control_Monad_List_Trans_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_singleton.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictApplicative|
                                                                           {
                                                                               let nil1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_nil(),
                                                                                                                    dictApplicative);
                                                                               &Func1::new({
                                                                                               let dictApplicative
                                                                                                   =
                                                                                                   dictApplicative.clone();
                                                                                               let nil1
                                                                                                   =
                                                                                                   nil1.clone();
                                                                                               move
                                                                                                   |a|
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_prepend(),
                                                                                                                                                                                                          &&&dictApplicative),
                                                                                                                                                                       a),
                                                                                                                                    &&&nil1)
                                                                                           })
                                                                           }))
    }
    pub fn Control_Monad_List_Trans_take_004055() -> &dyn Any {
        &Func1::new(move |dictApplicative|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_take_tco(dictApplicative))
    }
    pub fn Control_Monad_List_Trans_take_004055_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_take_004055_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_take_004055_002d1.get_or_init(||
                                                                   Lazy(Control_Monad_List_Trans_take_004055.clone()))
    }
    pub fn Control_Monad_List_Trans_take_tco(dictApplicative: &dyn Any)
     -> &dyn Any {
        let nil1 =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_nil(),
                                             dictApplicative);
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Control_Monad_List_Trans_take_004055_002d1 =
                            Control_Monad_List_Trans_take_004055_002d1.clone();
                        let Functor0 = Functor0.clone();
                        let dictApplicative = dictApplicative.clone();
                        let nil1 = nil1.clone();
                        move |v|
                            &Func1::new({
                                            let Control_Monad_List_Trans_take_004055_002d1
                                                =
                                                Control_Monad_List_Trans_take_004055_002d1.clone();
                                            let v = v.clone();
                                            move |v1|
                                                {
                                                    let matchValue =
                                                        Sharpurs_Prelude::unbox(&&v);
                                                    let matchValue_1 =
                                                        Sharpurs_Prelude::unbox(v1);
                                                    match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                    &matchValue)
                                                        {
                                                        0_i32 => &nil1,
                                                        _ => {
                                                            let n =
                                                                matchValue;
                                                            let f =
                                                                &Func1::new({
                                                                                let Control_Monad_List_Trans_take_004055_002d1
                                                                                    =
                                                                                    Control_Monad_List_Trans_take_004055_002d1.clone();
                                                                                let n
                                                                                    =
                                                                                    n.clone();
                                                                                move
                                                                                    |v2|
                                                                                    {
                                                                                        let matchValue_3:
                                                                                                LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                                                                            Sharpurs_Prelude::unbox(v2);
                                                                                        match matchValue_3.as_ref()
                                                                                            {
                                                                                            PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_3_1_0)
                                                                                            =>
                                                                                            &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_take_004055_002d1.Value,
                                                                                                                                                                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                          &&&n)),
                                                                                                                                                                                                                                                    &&matchValue_3_1_0))),
                                                                                            PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                                                                            =>
                                                                                            &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                                                                            PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_3_0_0,
                                                                                                                                                                                                       matchValue_3_0_1)
                                                                                            =>
                                                                                            &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_3_0_0,
                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_take_004055_002d1.Value,
                                                                                                                                                                                                                                                                                                                                                              &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                 &&&n),
                                                                                                                                                                                                                                                                                                                                                              &&&1_i32))),
                                                                                                                                                                                                                                                     &&matchValue_3_0_1))),
                                                                                        }
                                                                                    }
                                                                            });
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_stepMap(),
                                                                                                                                                                   &&&Functor0),
                                                                                                                                &&&f),
                                                                                             &&&matchValue_1)
                                                        }
                                                    }
                                                }
                                        })
                    })
    }
    pub fn Control_Monad_List_Trans_take() -> &dyn Any {
        static Control_Monad_List_Trans_take: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_take.get_or_init(||
                                                      Control_Monad_List_Trans_take_004055_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_zipWith_prime_004059() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_zipWith_prime_tco(dictMonad))
    }
    pub fn Control_Monad_List_Trans_zipWith_prime_004059_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_zipWith_prime_004059_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_zipWith_prime_004059_002d1.get_or_init(||
                                                                            Lazy(Control_Monad_List_Trans_zipWith_prime_004059.clone()))
    }
    pub fn Control_Monad_List_Trans_zipWith_prime_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let prepend_prime1 =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_prepend_prime(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                       Sharpurs_Prelude::unbox(dictMonad)),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Applicative0 = Applicative0.clone();
                        let Bind1 = Bind1.clone();
                        let Control_Monad_List_Trans_zipWith_prime_004059_002d1
                            =
                            Control_Monad_List_Trans_zipWith_prime_004059_002d1.clone();
                        let Functor0 = Functor0.clone();
                        let dictMonad = dictMonad.clone();
                        let prepend_prime1 = prepend_prime1.clone();
                        move |f|
                            {
                                let g =
                                    &Func1::new({
                                                    let Control_Monad_List_Trans_zipWith_prime_004059_002d1
                                                        =
                                                        Control_Monad_List_Trans_zipWith_prime_004059_002d1.clone();
                                                    let f = f.clone();
                                                    move |v|
                                                        &Func1::new({
                                                                        let Control_Monad_List_Trans_zipWith_prime_004059_002d1
                                                                            =
                                                                            Control_Monad_List_Trans_zipWith_prime_004059_002d1.clone();
                                                                        let v
                                                                            =
                                                                            v.clone();
                                                                        move
                                                                            |v1|
                                                                            {
                                                                                let matchValue:
                                                                                        LrcPtr<Data_Maybe_Maybe> =
                                                                                    Sharpurs_Prelude::unbox(&&v);
                                                                                let matchValue_1:
                                                                                        LrcPtr<Data_Maybe_Maybe> =
                                                                                    Sharpurs_Prelude::unbox(v1);
                                                                                if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                       =
                                                                                       matchValue_1.as_ref()
                                                                                   {
                                                                                    if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                           =
                                                                                           matchValue.as_ref()
                                                                                       {
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                                   _
                                                                                                                                   =>
                                                                                                                                   unreachable!(),
                                                                                                                               });
                                                                                        let activePatternResult_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                                   _
                                                                                                                                   =>
                                                                                                                                   unreachable!(),
                                                                                                                               });
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                               &&&Functor0),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                  &&&prepend_prime1),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                                                    let Control_Monad_List_Trans_zipWith_prime_004059_002d1
                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                        Control_Monad_List_Trans_zipWith_prime_004059_002d1.clone();
                                                                                                                                                                                                                                                    let activePatternResult
                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                        activePatternResult.clone();
                                                                                                                                                                                                                                                    let activePatternResult_1
                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                        activePatternResult_1.clone();
                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                        |v2|
                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_zipWith_prime_004059_002d1.Value,
                                                                                                                                                                                                                                                                                                                                                                                                  &&&dictMonad),
                                                                                                                                                                                                                                                                                                                                                               &&&f),
                                                                                                                                                                                                                                                                                                                            &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                                         &&&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                })))),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                               &&&match activePatternResult.as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                  }),
                                                                                                                                                            &&&match activePatternResult_1.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                              _)
                                                                                                                                                                   =>
                                                                                                                                                                   x.clone(),
                                                                                                                                                               }))
                                                                                    } else {
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                            &&&Applicative0),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_nil(),
                                                                                                                                                            &&&Applicative0))
                                                                                    }
                                                                                } else {
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                        &&&Applicative0),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_nil(),
                                                                                                                                                        &&&Applicative0))
                                                                                }
                                                                            }
                                                                    })
                                                });
                                &&Func1::new({
                                                 let g = g.clone();
                                                 move |fa|
                                                     &Func1::new({
                                                                     let fa =
                                                                         fa.clone();
                                                                     move |fb|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_wrapEffect(),
                                                                                                                                             &&&Functor0),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                   &&&Bind1),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_uncons(),
                                                                                                                                                                                                                                                      &&&dictMonad),
                                                                                                                                                                                                                   &&&fa)),
                                                                                                                                             &&&Func1::new({
                                                                                                                                                               let fb
                                                                                                                                                                   =
                                                                                                                                                                   fb.clone();
                                                                                                                                                               move
                                                                                                                                                                   |ua|
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                          &&&Bind1),
                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_uncons(),
                                                                                                                                                                                                                                                                                                             &&&dictMonad),
                                                                                                                                                                                                                                                                          &&&fb)),
                                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                                      let ua
                                                                                                                                                                                                                          =
                                                                                                                                                                                                                          ua.clone();
                                                                                                                                                                                                                      move
                                                                                                                                                                                                                          |ub|
                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                                                                              &&&ua),
                                                                                                                                                                                                                                                           ub)
                                                                                                                                                                                                                  }))
                                                                                                                                                           })))
                                                                 })
                                             })
                            }
                    })
    }
    pub fn Control_Monad_List_Trans_zipWith_prime() -> &dyn Any {
        static Control_Monad_List_Trans_zipWith_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_zipWith_prime.get_or_init(||
                                                               Control_Monad_List_Trans_zipWith_prime_004059_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_zipWith() -> &dyn Any {
        static Control_Monad_List_Trans_zipWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_zipWith.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonad|
                                                                         {
                                                                             let pure_var =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             &Func1::new({
                                                                                             let dictMonad
                                                                                                 =
                                                                                                 dictMonad.clone();
                                                                                             let pure_var
                                                                                                 =
                                                                                                 pure_var.clone();
                                                                                             move
                                                                                                 |f|
                                                                                                 {
                                                                                                     let g =
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
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                     &&&pure_var),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                        &&&a),
                                                                                                                                                                                                                     b))
                                                                                                                                         })
                                                                                                                     });
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_zipWith_prime(),
                                                                                                                                                                         &&&dictMonad),
                                                                                                                                      &&&g)
                                                                                                 }
                                                                                         })
                                                                         }))
    }
    pub fn Control_Monad_List_Trans_newtypeListT() -> &dyn Any {
        static Control_Monad_List_Trans_newtypeListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_newtypeListT.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                               &&&add(string("Coercible0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &Sharpurs_Prelude::Prim_undefined()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))
    }
    pub fn Control_Monad_List_Trans_mapMaybe_004067() -> &dyn Any {
        &Func1::new(move |dictFunctor|
                        Func1::new({
                                       let dictFunctor = dictFunctor.clone();
                                       move |f|
                                           PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_mapMaybe_tco(&dictFunctor,
                                                                                                                      f)
                                   }))
    }
    pub fn Control_Monad_List_Trans_mapMaybe_004067_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_mapMaybe_004067_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_mapMaybe_004067_002d1.get_or_init(||
                                                                       Lazy(Control_Monad_List_Trans_mapMaybe_004067.clone()))
    }
    pub fn Control_Monad_List_Trans_mapMaybe_tco(dictFunctor: &dyn Any,
                                                 f: &dyn Any) -> &dyn Any {
        let g =
            &Func1::new({
                            let Control_Monad_List_Trans_mapMaybe_004067_002d1
                                =
                                Control_Monad_List_Trans_mapMaybe_004067_002d1.clone();
                            let dictFunctor = dictFunctor.clone();
                            let f = f.clone();
                            move |v|
                                {
                                    let matchValue:
                                            LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                        Sharpurs_Prelude::unbox(v);
                                    match matchValue.as_ref() {
                                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_1_0)
                                        =>
                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                            &&&Func1::new(move
                                                                                                                              |usd__arg1_2|
                                                                                                                              &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(usd__arg1_2.clone())))),
                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_mapMaybe_004067_002d1.Value,
                                                                                                                                                                                                                     &&&dictFunctor),
                                                                                                                                                                                  &&&f)),
                                                                                                            &&matchValue_1_0)),
                                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                        =>
                                        &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_0_0,
                                                                                                                                                   matchValue_0_1)
                                        =>
                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromMaybe(),
                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                 |usd__arg1|
                                                                                                                                                                 &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(usd__arg1.clone())))),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                     &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                    |usd__arg1_1|
                                                                                                                                                                                                    Func1::new({
                                                                                                                                                                                                                   let usd__arg1_1
                                                                                                                                                                                                                       =
                                                                                                                                                                                                                       usd__arg1_1.clone();
                                                                                                                                                                                                                   move
                                                                                                                                                                                                                       |usd__arg2|
                                                                                                                                                                                                                       &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                                                                                                               usd__arg2.clone()))
                                                                                                                                                                                                               }))),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                  &&matchValue_0_0))),
                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_mapMaybe_004067_002d1.Value,
                                                                                                                                                                                                                     &&&dictFunctor),
                                                                                                                                                                                  &&&f)),
                                                                                                            &&matchValue_0_1)),
                                    }
                                }
                        });
        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_stepMap(),
                                                                            dictFunctor),
                                         &&&g)
    }
    pub fn Control_Monad_List_Trans_mapMaybe() -> &dyn Any {
        static Control_Monad_List_Trans_mapMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_mapMaybe.get_or_init(||
                                                          Control_Monad_List_Trans_mapMaybe_004067_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_iterate() -> &dyn Any {
        static Control_Monad_List_Trans_iterate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_iterate.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonad|
                                                                         {
                                                                             let pure_var =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             &Func1::new({
                                                                                             let dictMonad
                                                                                                 =
                                                                                                 dictMonad.clone();
                                                                                             let pure_var
                                                                                                 =
                                                                                                 pure_var.clone();
                                                                                             move
                                                                                                 |f|
                                                                                                 &Func1::new({
                                                                                                                 let f
                                                                                                                     =
                                                                                                                     f.clone();
                                                                                                                 move
                                                                                                                     |a|
                                                                                                                     {
                                                                                                                         let g =
                                                                                                                             &Func1::new(move
                                                                                                                                             |x|
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                 &&&pure_var),
                                                                                                                                                                              &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                x),
                                                                                                                                                                                                                                                                                               x.clone()))))));
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_unfold(),
                                                                                                                                                                                                                                &&&dictMonad),
                                                                                                                                                                                             &&&g),
                                                                                                                                                          a)
                                                                                                                     }
                                                                                                             })
                                                                                         })
                                                                         }))
    }
    pub fn Control_Monad_List_Trans_repeat() -> &dyn Any {
        static Control_Monad_List_Trans_repeat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_repeat.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictMonad|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_iterate(),
                                                                                                                                            dictMonad),
                                                                                                         &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_identity())))
    }
    pub fn Control_Monad_List_Trans_head() -> &dyn Any {
        static Control_Monad_List_Trans_head: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_head.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonad|
                                                                      {
                                                                          let Functor0 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                          &Func1::new({
                                                                                          let Functor0
                                                                                              =
                                                                                              Functor0.clone();
                                                                                          let dictMonad
                                                                                              =
                                                                                              dictMonad.clone();
                                                                                          move
                                                                                              |l|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                     &&&Functor0),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                        &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                     &&&PureScript_Data_Tuple::Data_Tuple_fst())),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_uncons(),
                                                                                                                                                                                                     &&&dictMonad),
                                                                                                                                                                  l))
                                                                                      })
                                                                      }))
    }
    pub fn Control_Monad_List_Trans_functorListT_004077() -> &dyn Any {
        &Func1::new(move |dictFunctor|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_functorListT_tco(dictFunctor))
    }
    pub fn Control_Monad_List_Trans_functorListT_004077_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_functorListT_004077_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_functorListT_004077_002d1.get_or_init(||
                                                                           Lazy(Control_Monad_List_Trans_functorListT_004077.clone()))
    }
    pub fn Control_Monad_List_Trans_functorListT_tco(dictFunctor: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                         &&&add(string("map"),
                                                &&Func1::new({
                                                                 let dictFunctor
                                                                     =
                                                                     dictFunctor.clone();
                                                                 move |f|
                                                                     {
                                                                         let g =
                                                                             &Func1::new({
                                                                                             let f
                                                                                                 =
                                                                                                 f.clone();
                                                                                             move
                                                                                                 |v|
                                                                                                 {
                                                                                                     let matchValue:
                                                                                                             LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                     match matchValue.as_ref()
                                                                                                         {
                                                                                                         PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_1_0)
                                                                                                         =>
                                                                                                         &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                          &&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_functorListT_tco(&&dictFunctor)),
                                                                                                                                                                                                                                                                                                                                       &&&f)),
                                                                                                                                                                                                                                                                 &&matchValue_1_0))),
                                                                                                         PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                                                                                         =>
                                                                                                         &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                                                                                         PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_0_0,
                                                                                                                                                                                                                    matchValue_0_1)
                                                                                                         =>
                                                                                                         &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                  &&matchValue_0_0),
                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                           &&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_functorListT_tco(&&dictFunctor)),
                                                                                                                                                                                                                                                                                                                                        &&&f)),
                                                                                                                                                                                                                                                                  &&matchValue_0_1))),
                                                                                                     }
                                                                                                 }
                                                                                         });
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_stepMap(),
                                                                                                                                             &&&dictFunctor),
                                                                                                          &&&g)
                                                                     }
                                                             }),
                                                empty::<string, &dyn Any>()))
    }
    pub fn Control_Monad_List_Trans_functorListT() -> &dyn Any {
        static Control_Monad_List_Trans_functorListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_functorListT.get_or_init(||
                                                              Control_Monad_List_Trans_functorListT_004077_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_fromEffect() -> &dyn Any {
        static Control_Monad_List_Trans_fromEffect: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_List_Trans_fromEffect.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictApplicative|
                                                                            {
                                                                                let Functor0 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                             Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                                &Func1::new({
                                                                                                let Functor0
                                                                                                    =
                                                                                                    Functor0.clone();
                                                                                                let dictApplicative
                                                                                                    =
                                                                                                    dictApplicative.clone();
                                                                                                move
                                                                                                    |fa|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                        &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_ListT()),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                              &&&Functor0),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                                   |usd__arg1|
                                                                                                                                                                                                                                                                                                   Func1::new({
                                                                                                                                                                                                                                                                                                                  let usd__arg1
                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                      usd__arg1.clone();
                                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                                      |usd__arg2|
                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                                                                              usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                              }))),
                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Lazy::Data_Lazy_defer()),
                                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                                   |v|
                                                                                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_nil(),
                                                                                                                                                                                                                                                                                                                                    &&&dictApplicative))))),
                                                                                                                                                                        fa))
                                                                                            })
                                                                            }))
    }
    pub fn Control_Monad_List_Trans_monadTransListT() -> &dyn Any {
        static Control_Monad_List_Trans_monadTransListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_monadTransListT.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_MonadTransusd_Dict(),
                                                                                                  &&&add(string("lift"),
                                                                                                         &&Func1::new(move
                                                                                                                          |dictMonad|
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_fromEffect(),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Control_Monad_List_Trans_lift() -> &dyn Any {
        static Control_Monad_List_Trans_lift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_lift.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_lift(),
                                                                                       &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_monadTransListT()))
    }
    pub fn Control_Monad_List_Trans_foldlRec_prime() -> &dyn Any {
        static Control_Monad_List_Trans_foldlRec_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_foldlRec_prime.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictMonadRec|
                                                                                {
                                                                                    let Monad0 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    let Applicative0 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    let Bind1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    let Monad01 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    &Func1::new({
                                                                                                    let Applicative0
                                                                                                        =
                                                                                                        Applicative0.clone();
                                                                                                    let Bind1
                                                                                                        =
                                                                                                        Bind1.clone();
                                                                                                    let Monad01
                                                                                                        =
                                                                                                        Monad01.clone();
                                                                                                    let dictMonadRec
                                                                                                        =
                                                                                                        dictMonadRec.clone();
                                                                                                    move
                                                                                                        |f|
                                                                                                        {
                                                                                                            let r#loop =
                                                                                                                &Func1::new({
                                                                                                                                let f
                                                                                                                                    =
                                                                                                                                    f.clone();
                                                                                                                                move
                                                                                                                                    |b|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let b
                                                                                                                                                        =
                                                                                                                                                        b.clone();
                                                                                                                                                    move
                                                                                                                                                        |l|
                                                                                                                                                        {
                                                                                                                                                            let g =
                                                                                                                                                                &Func1::new(move
                                                                                                                                                                                |v|
                                                                                                                                                                                {
                                                                                                                                                                                    let matchValue:
                                                                                                                                                                                            LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                    match matchValue.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                        =>
                                                                                                                                                                                        {
                                                                                                                                                                                            let activePatternResult:
                                                                                                                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                Sharpurs_Prelude::_007cUnbox_007c(matchValue_1_0);
                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                   &&&Bind1),
                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                      &&&b),
                                                                                                                                                                                                                                                                                                   &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                      })),
                                                                                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                                                                                               let activePatternResult
                                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                                   activePatternResult.clone();
                                                                                                                                                                                                                                               move
                                                                                                                                                                                                                                                   |b_prime|
                                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                       &&&Applicative0),
                                                                                                                                                                                                                                                                                    &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&add(string("a"),
                                                                                                                                                                                                                                                                                                                                                                           b_prime.clone(),
                                                                                                                                                                                                                                                                                                                                                                           add(string("b"),
                                                                                                                                                                                                                                                                                                                                                                               &&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                                                                                                                               empty::<string,
                                                                                                                                                                                                                                                                                                                                                                                       &dyn Any>())))))
                                                                                                                                                                                                                                           }))
                                                                                                                                                                                        }
                                                                                                                                                                                        _
                                                                                                                                                                                        =>
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                            &&&Applicative0),
                                                                                                                                                                                                                         &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&b))),
                                                                                                                                                                                    }
                                                                                                                                                                                });
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                   &&&Bind1),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_uncons(),
                                                                                                                                                                                                                                                                                                      &&&Monad01),
                                                                                                                                                                                                                                                                   l)),
                                                                                                                                                                                             &&&g)
                                                                                                                                                        }
                                                                                                                                                })
                                                                                                                            });
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM2(),
                                                                                                                                                                                &&&dictMonadRec),
                                                                                                                                             &&&r#loop)
                                                                                                        }
                                                                                                })
                                                                                }))
    }
    pub fn Control_Monad_List_Trans_runListTRec() -> &dyn Any {
        static Control_Monad_List_Trans_runListTRec: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_List_Trans_runListTRec.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictMonadRec|
                                                                             {
                                                                                 let Applicative0 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                             Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_foldlRec_prime(),
                                                                                                                                                                                        dictMonadRec),
                                                                                                                                                     &&&Func1::new({
                                                                                                                                                                       let Applicative0
                                                                                                                                                                           =
                                                                                                                                                                           Applicative0.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |v|
                                                                                                                                                                           &Func1::new(move
                                                                                                                                                                                           |v1|
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                               &&&Applicative0),
                                                                                                                                                                                                                            &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                                   })),
                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit())
                                                                             }))
    }
    pub fn Control_Monad_List_Trans_foldlRec() -> &dyn Any {
        static Control_Monad_List_Trans_foldlRec: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_foldlRec.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonadRec|
                                                                          {
                                                                              let Monad0 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                          Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              let Applicative0 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                          Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              let Bind1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                          Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              let Monad01 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                          Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              &Func1::new({
                                                                                              let Applicative0
                                                                                                  =
                                                                                                  Applicative0.clone();
                                                                                              let Bind1
                                                                                                  =
                                                                                                  Bind1.clone();
                                                                                              let Monad01
                                                                                                  =
                                                                                                  Monad01.clone();
                                                                                              let dictMonadRec
                                                                                                  =
                                                                                                  dictMonadRec.clone();
                                                                                              move
                                                                                                  |f|
                                                                                                  {
                                                                                                      let r#loop =
                                                                                                          &Func1::new({
                                                                                                                          let f
                                                                                                                              =
                                                                                                                              f.clone();
                                                                                                                          move
                                                                                                                              |b|
                                                                                                                              &Func1::new({
                                                                                                                                              let b
                                                                                                                                                  =
                                                                                                                                                  b.clone();
                                                                                                                                              move
                                                                                                                                                  |l|
                                                                                                                                                  {
                                                                                                                                                      let g =
                                                                                                                                                          &Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          {
                                                                                                                                                                              let matchValue:
                                                                                                                                                                                      LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                                                              match matchValue.as_ref()
                                                                                                                                                                                  {
                                                                                                                                                                                  Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                  =>
                                                                                                                                                                                  {
                                                                                                                                                                                      let activePatternResult:
                                                                                                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                          Sharpurs_Prelude::_007cUnbox_007c(matchValue_1_0);
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                          &&&Applicative0),
                                                                                                                                                                                                                       &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&add(string("a"),
                                                                                                                                                                                                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                   &&&b),
                                                                                                                                                                                                                                                                                                                                                &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                                                   }),
                                                                                                                                                                                                                                                                                                              add(string("b"),
                                                                                                                                                                                                                                                                                                                  &&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                    },
                                                                                                                                                                                                                                                                                                                  empty::<string,
                                                                                                                                                                                                                                                                                                                          &dyn Any>())))))
                                                                                                                                                                                  }
                                                                                                                                                                                  _
                                                                                                                                                                                  =>
                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                      &&&Applicative0),
                                                                                                                                                                                                                   &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&b))),
                                                                                                                                                                              }
                                                                                                                                                                          });
                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                             &&&Bind1),
                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_uncons(),
                                                                                                                                                                                                                                                                                                &&&Monad01),
                                                                                                                                                                                                                                                             l)),
                                                                                                                                                                                       &&&g)
                                                                                                                                                  }
                                                                                                                                          })
                                                                                                                      });
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM2(),
                                                                                                                                                                          &&&dictMonadRec),
                                                                                                                                       &&&r#loop)
                                                                                                  }
                                                                                          })
                                                                          }))
    }
    pub fn Control_Monad_List_Trans_foldl_prime() -> &dyn Any {
        static Control_Monad_List_Trans_foldl_prime: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_List_Trans_foldl_prime.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictMonad|
                                                                             {
                                                                                 let Applicative0 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                             Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                 let Bind1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                             Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                 &Func1::new({
                                                                                                 let Applicative0
                                                                                                     =
                                                                                                     Applicative0.clone();
                                                                                                 let Bind1
                                                                                                     =
                                                                                                     Bind1.clone();
                                                                                                 let dictMonad
                                                                                                     =
                                                                                                     dictMonad.clone();
                                                                                                 move
                                                                                                     |f|
                                                                                                     {
                                                                                                         let loop_2 =
                                                                                                             Func0::new({
                                                                                                                            let loop_tco
                                                                                                                                =
                                                                                                                                loop_tco.clone();
                                                                                                                            move
                                                                                                                                ||
                                                                                                                                &Func1::new({
                                                                                                                                                let loop_tco
                                                                                                                                                    =
                                                                                                                                                    loop_tco.clone();
                                                                                                                                                move
                                                                                                                                                    |b|
                                                                                                                                                    Func1::new({
                                                                                                                                                                   let b
                                                                                                                                                                       =
                                                                                                                                                                       b.clone();
                                                                                                                                                                   let loop_tco
                                                                                                                                                                       =
                                                                                                                                                                       loop_tco.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |l|
                                                                                                                                                                       loop_tco(b)(l.clone())
                                                                                                                                                               })
                                                                                                                                            })
                                                                                                                        });
                                                                                                         let loop_1 =
                                                                                                             Lazy(loop_2);
                                                                                                         let loop_tco =
                                                                                                             Func1::new({
                                                                                                                            let f
                                                                                                                                =
                                                                                                                                f.clone();
                                                                                                                            let loop_1
                                                                                                                                =
                                                                                                                                loop_1.clone();
                                                                                                                            move
                                                                                                                                |b_1|
                                                                                                                                Func1::new({
                                                                                                                                               let b_1
                                                                                                                                                   =
                                                                                                                                                   b_1.clone();
                                                                                                                                               let loop_1
                                                                                                                                                   =
                                                                                                                                                   loop_1.clone();
                                                                                                                                               move
                                                                                                                                                   |l_1|
                                                                                                                                                   {
                                                                                                                                                       let g =
                                                                                                                                                           &Func1::new({
                                                                                                                                                                           let loop_1
                                                                                                                                                                               =
                                                                                                                                                                               loop_1.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |v|
                                                                                                                                                                               {
                                                                                                                                                                                   let matchValue:
                                                                                                                                                                                           LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                   match matchValue.as_ref()
                                                                                                                                                                                       {
                                                                                                                                                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                       =>
                                                                                                                                                                                       {
                                                                                                                                                                                           let activePatternResult:
                                                                                                                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                               Sharpurs_Prelude::_007cUnbox_007c(matchValue_1_0);
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                  &&&Bind1),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                     &&&b_1),
                                                                                                                                                                                                                                                                                                  &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                     })),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                  &&&loop_1.Value),
                                                                                                                                                                                                                                                               &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                  }))
                                                                                                                                                                                       }
                                                                                                                                                                                       _
                                                                                                                                                                                       =>
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                           &&&Applicative0),
                                                                                                                                                                                                                        &&&b_1),
                                                                                                                                                                                   }
                                                                                                                                                                               }
                                                                                                                                                                       });
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                              &&&Bind1),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_uncons(),
                                                                                                                                                                                                                                                                                                 &&&dictMonad),
                                                                                                                                                                                                                                                              l_1)),
                                                                                                                                                                                        &&&g)
                                                                                                                                                   }
                                                                                                                                           })
                                                                                                                        });
                                                                                                         let r#loop =
                                                                                                             loop_1.Value;
                                                                                                         &r#loop
                                                                                                     }
                                                                                             })
                                                                             }))
    }
    pub fn Control_Monad_List_Trans_runListT() -> &dyn Any {
        static Control_Monad_List_Trans_runListT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_runListT.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonad|
                                                                          {
                                                                              let Applicative0 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                          Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_foldl_prime(),
                                                                                                                                                                                     dictMonad),
                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                    let Applicative0
                                                                                                                                                                        =
                                                                                                                                                                        Applicative0.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |v|
                                                                                                                                                                        &Func1::new(move
                                                                                                                                                                                        |v1|
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                            &&&Applicative0),
                                                                                                                                                                                                                         &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                                })),
                                                                                                               &&&PureScript_Data_Unit::Data_Unit_unit())
                                                                          }))
    }
    pub fn Control_Monad_List_Trans_foldl() -> &dyn Any {
        static Control_Monad_List_Trans_foldl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_foldl.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonad|
                                                                       {
                                                                           let Applicative0 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                       Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                           let Bind1 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                       Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                           &Func1::new({
                                                                                           let Applicative0
                                                                                               =
                                                                                               Applicative0.clone();
                                                                                           let Bind1
                                                                                               =
                                                                                               Bind1.clone();
                                                                                           let dictMonad
                                                                                               =
                                                                                               dictMonad.clone();
                                                                                           move
                                                                                               |f|
                                                                                               {
                                                                                                   let loop_2 =
                                                                                                       Func0::new({
                                                                                                                      let loop_tco
                                                                                                                          =
                                                                                                                          loop_tco.clone();
                                                                                                                      move
                                                                                                                          ||
                                                                                                                          &Func1::new({
                                                                                                                                          let loop_tco
                                                                                                                                              =
                                                                                                                                              loop_tco.clone();
                                                                                                                                          move
                                                                                                                                              |b|
                                                                                                                                              Func1::new({
                                                                                                                                                             let b
                                                                                                                                                                 =
                                                                                                                                                                 b.clone();
                                                                                                                                                             let loop_tco
                                                                                                                                                                 =
                                                                                                                                                                 loop_tco.clone();
                                                                                                                                                             move
                                                                                                                                                                 |l|
                                                                                                                                                                 loop_tco(b)(l.clone())
                                                                                                                                                         })
                                                                                                                                      })
                                                                                                                  });
                                                                                                   let loop_1 =
                                                                                                       Lazy(loop_2);
                                                                                                   let loop_tco =
                                                                                                       Func1::new({
                                                                                                                      let f
                                                                                                                          =
                                                                                                                          f.clone();
                                                                                                                      move
                                                                                                                          |b_1|
                                                                                                                          fix1(&(move
                                                                                                                                     |loop_tco,
                                                                                                                                      b_1|
                                                                                                                                     Func1::new({
                                                                                                                                                    let b_1
                                                                                                                                                        =
                                                                                                                                                        b_1.clone();
                                                                                                                                                    let loop_tco
                                                                                                                                                        =
                                                                                                                                                        loop_tco.clone();
                                                                                                                                                    move
                                                                                                                                                        |l_1|
                                                                                                                                                        {
                                                                                                                                                            let g =
                                                                                                                                                                &Func1::new({
                                                                                                                                                                                let loop_tco
                                                                                                                                                                                    =
                                                                                                                                                                                    loop_tco.clone();
                                                                                                                                                                                move
                                                                                                                                                                                    |v|
                                                                                                                                                                                    {
                                                                                                                                                                                        let matchValue:
                                                                                                                                                                                                LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                        match matchValue.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                            =>
                                                                                                                                                                                            {
                                                                                                                                                                                                let activePatternResult:
                                                                                                                                                                                                        LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                    Sharpurs_Prelude::_007cUnbox_007c(matchValue_1_0);
                                                                                                                                                                                                loop_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                             &&&b_1),
                                                                                                                                                                                                                                          &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                             }))(&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                  })
                                                                                                                                                                                            }
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                &&&Applicative0),
                                                                                                                                                                                                                             &&&b_1),
                                                                                                                                                                                        }
                                                                                                                                                                                    }
                                                                                                                                                                            });
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                   &&&Bind1),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_uncons(),
                                                                                                                                                                                                                                                                                                      &&&dictMonad),
                                                                                                                                                                                                                                                                   l_1)),
                                                                                                                                                                                             &&&g)
                                                                                                                                                        }
                                                                                                                                                })),
                                                                                                                               b_1.clone())
                                                                                                                  });
                                                                                                   let r#loop =
                                                                                                       loop_1.Value;
                                                                                                   &r#loop
                                                                                               }
                                                                                       })
                                                                       }))
    }
    pub fn Control_Monad_List_Trans_filter_0040111() -> &dyn Any {
        &Func1::new(move |dictFunctor|
                        Func1::new({
                                       let dictFunctor = dictFunctor.clone();
                                       move |f|
                                           PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_filter_tco(&dictFunctor,
                                                                                                                    f)
                                   }))
    }
    pub fn Control_Monad_List_Trans_filter_0040111_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_filter_0040111_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_filter_0040111_002d1.get_or_init(||
                                                                      Lazy(Control_Monad_List_Trans_filter_0040111.clone()))
    }
    pub fn Control_Monad_List_Trans_filter_tco(dictFunctor: &dyn Any,
                                               f: &dyn Any) -> &dyn Any {
        let g =
            &Func1::new({
                            let Control_Monad_List_Trans_filter_0040111_002d1
                                =
                                Control_Monad_List_Trans_filter_0040111_002d1.clone();
                            let dictFunctor = dictFunctor.clone();
                            let f = f.clone();
                            move |v|
                                {
                                    let matchValue:
                                            LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                        Sharpurs_Prelude::unbox(v);
                                    match matchValue.as_ref() {
                                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_1_0)
                                        =>
                                        &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                       &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_filter_0040111_002d1.Value,
                                                                                                                                                                                                                                                                                                          &&&dictFunctor),
                                                                                                                                                                                                                                                                       &&&f)),
                                                                                                                                                                                                 &&matchValue_1_0))),
                                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                        =>
                                        &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_0_0,
                                                                                                                                                   matchValue_0_1)
                                        => {
                                            let a = matchValue_0_0.clone();
                                            let s_prime =
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                       &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_filter_0040111_002d1.Value,
                                                                                                                                                                                          &&&dictFunctor),
                                                                                                                                                       &&&f)),
                                                                                 &&matchValue_0_1);
                                            let matchValue_1 =
                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                          &&&a));
                                            match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                             &matchValue_1)
                                                {
                                                0_i32 =>
                                                &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(&a,
                                                                                                                                                                        &s_prime)),
                                                _ =>
                                                &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(&s_prime)),
                                            }
                                        }
                                    }
                                }
                        });
        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_stepMap(),
                                                                            dictFunctor),
                                         &&&g)
    }
    pub fn Control_Monad_List_Trans_filter() -> &dyn Any {
        static Control_Monad_List_Trans_filter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_filter.get_or_init(||
                                                        Control_Monad_List_Trans_filter_0040111_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_dropWhile_0040115() -> &dyn Any {
        &Func1::new(move |dictApplicative|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_dropWhile_tco(dictApplicative))
    }
    pub fn Control_Monad_List_Trans_dropWhile_0040115_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_dropWhile_0040115_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_dropWhile_0040115_002d1.get_or_init(||
                                                                         Lazy(Control_Monad_List_Trans_dropWhile_0040115.clone()))
    }
    pub fn Control_Monad_List_Trans_dropWhile_tco(dictApplicative: &dyn Any)
     -> &dyn Any {
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Control_Monad_List_Trans_dropWhile_0040115_002d1 =
                            Control_Monad_List_Trans_dropWhile_0040115_002d1.clone();
                        let Functor0 = Functor0.clone();
                        let dictApplicative = dictApplicative.clone();
                        move |f|
                            {
                                let g =
                                    &Func1::new({
                                                    let Control_Monad_List_Trans_dropWhile_0040115_002d1
                                                        =
                                                        Control_Monad_List_Trans_dropWhile_0040115_002d1.clone();
                                                    let f = f.clone();
                                                    move |v|
                                                        {
                                                            let matchValue:
                                                                    LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                                                Sharpurs_Prelude::unbox(v);
                                                            match matchValue.as_ref()
                                                                {
                                                                PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_1_0)
                                                                =>
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                    &&&Func1::new(move
                                                                                                                                                      |usd__arg1|
                                                                                                                                                      &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(usd__arg1.clone())))),
                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                          &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_dropWhile_0040115_002d1.Value,
                                                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                                                          &&&f)),
                                                                                                                                    &&matchValue_1_0)),
                                                                PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                                                =>
                                                                &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                                                PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_0_0,
                                                                                                                                                                           matchValue_0_1)
                                                                => {
                                                                    let s =
                                                                        matchValue_0_1.clone();
                                                                    let a =
                                                                        matchValue_0_0.clone();
                                                                    let matchValue_1 =
                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                  &&&a));
                                                                    match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                     &matchValue_1)
                                                                        {
                                                                        0_i32
                                                                        =>
                                                                        &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_dropWhile_0040115_002d1.Value,
                                                                                                                                                                                                                                                                                                                                         &&&dictApplicative),
                                                                                                                                                                                                                                                                                                      &&&f)),
                                                                                                                                                                                                                                &&&s))),
                                                                        _ =>
                                                                        &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(&a,
                                                                                                                                                                                                &s)),
                                                                    }
                                                                }
                                                            }
                                                        }
                                                });
                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_stepMap(),
                                                                                                    &&&Functor0),
                                                                 &&&g)
                            }
                    })
    }
    pub fn Control_Monad_List_Trans_dropWhile() -> &dyn Any {
        static Control_Monad_List_Trans_dropWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_dropWhile.get_or_init(||
                                                           Control_Monad_List_Trans_dropWhile_0040115_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_drop_0040119() -> &dyn Any {
        &Func1::new(move |dictApplicative|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_drop_tco(dictApplicative))
    }
    pub fn Control_Monad_List_Trans_drop_0040119_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_drop_0040119_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_drop_0040119_002d1.get_or_init(||
                                                                    Lazy(Control_Monad_List_Trans_drop_0040119.clone()))
    }
    pub fn Control_Monad_List_Trans_drop_tco(dictApplicative: &dyn Any)
     -> &dyn Any {
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Control_Monad_List_Trans_drop_0040119_002d1 =
                            Control_Monad_List_Trans_drop_0040119_002d1.clone();
                        let Functor0 = Functor0.clone();
                        let dictApplicative = dictApplicative.clone();
                        move |v|
                            &Func1::new({
                                            let Control_Monad_List_Trans_drop_0040119_002d1
                                                =
                                                Control_Monad_List_Trans_drop_0040119_002d1.clone();
                                            let v = v.clone();
                                            move |v1|
                                                {
                                                    let matchValue =
                                                        Sharpurs_Prelude::unbox(&&v);
                                                    let matchValue_1 =
                                                        Sharpurs_Prelude::unbox(v1);
                                                    match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                    &matchValue)
                                                        {
                                                        0_i32 => fa.clone(),
                                                        _ => {
                                                            let n =
                                                                matchValue;
                                                            let f =
                                                                &Func1::new({
                                                                                let Control_Monad_List_Trans_drop_0040119_002d1
                                                                                    =
                                                                                    Control_Monad_List_Trans_drop_0040119_002d1.clone();
                                                                                let n
                                                                                    =
                                                                                    n.clone();
                                                                                move
                                                                                    |v2|
                                                                                    {
                                                                                        let matchValue_3:
                                                                                                LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                                                                            Sharpurs_Prelude::unbox(v2);
                                                                                        match matchValue_3.as_ref()
                                                                                            {
                                                                                            PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_3_1_0)
                                                                                            =>
                                                                                            &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_drop_0040119_002d1.Value,
                                                                                                                                                                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                          &&&n)),
                                                                                                                                                                                                                                                    &&matchValue_3_1_0))),
                                                                                            PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                                                                            =>
                                                                                            &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                                                                            PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_3_0_0,
                                                                                                                                                                                                       matchValue_3_0_1)
                                                                                            =>
                                                                                            &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Control_Monad_List_Trans_drop_0040119_002d1.Value,
                                                                                                                                                                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                &&&n),
                                                                                                                                                                                                                                                                                                                                                             &&&1_i32))),
                                                                                                                                                                                                                                                    &&matchValue_3_0_1))),
                                                                                        }
                                                                                    }
                                                                            });
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_stepMap(),
                                                                                                                                                                   &&&Functor0),
                                                                                                                                &&&f),
                                                                                             &&&matchValue_1)
                                                        }
                                                    }
                                                }
                                        })
                    })
    }
    pub fn Control_Monad_List_Trans_drop() -> &dyn Any {
        static Control_Monad_List_Trans_drop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_drop.get_or_init(||
                                                      Control_Monad_List_Trans_drop_0040119_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_cons() -> &dyn Any {
        static Control_Monad_List_Trans_cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_cons.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictApplicative|
                                                                      {
                                                                          let pure_var =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                               dictApplicative);
                                                                          &Func1::new({
                                                                                          let pure_var
                                                                                              =
                                                                                              pure_var.clone();
                                                                                          move
                                                                                              |lh|
                                                                                              &Func1::new({
                                                                                                              let lh
                                                                                                                  =
                                                                                                                  lh.clone();
                                                                                                              move
                                                                                                                  |t|
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                      &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_ListT()),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                         &&&pure_var),
                                                                                                                                                                                      &&&LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                                                                 &&&lh),
                                                                                                                                                                                                                                                                                                                t.clone()))))
                                                                                                          })
                                                                                      })
                                                                      }))
    }
    pub fn Control_Monad_List_Trans_unfoldable1ListT() -> &dyn Any {
        static Control_Monad_List_Trans_unfoldable1ListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_unfoldable1ListT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonad|
                                                                                  {
                                                                                      let Applicative0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                  Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_Unfoldable1usd_Dict(),
                                                                                                                       &&&add(string("unfoldr1"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let Applicative0
                                                                                                                                                   =
                                                                                                                                                   Applicative0.clone();
                                                                                                                                               move
                                                                                                                                                   |f|
                                                                                                                                                   &Func1::new({
                                                                                                                                                                   let f
                                                                                                                                                                       =
                                                                                                                                                                       f.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |b|
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
                                                                                                                                                                           fn go_tco(v_1:
                                                                                                                                                                                         _)
                                                                                                                                                                            ->
                                                                                                                                                                                &dyn Any {
                                                                                                                                                                               let matchValue:
                                                                                                                                                                                       LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                               if let Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                                                      =
                                                                                                                                                                                      Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                             {
                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                         }).as_ref()
                                                                                                                                                                                  {
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_singleton(),
                                                                                                                                                                                                                                                       &&&Applicative0),
                                                                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                                                                           {
                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                       })
                                                                                                                                                                               } else {
                                                                                                                                                                                   let activePatternResult_1:
                                                                                                                                                                                           LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                       Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                              {
                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                          });
                                                                                                                                                                                   if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(activePatternResult_1_1_0)
                                                                                                                                                                                          =
                                                                                                                                                                                          activePatternResult_1.as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_cons(),
                                                                                                                                                                                                                                                                                              &&&Applicative0),
                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Lazy::Data_Lazy_applicativeLazy()),
                                                                                                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                                                                             let activePatternResult_1
                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                 activePatternResult_1.clone();
                                                                                                                                                                                                                                                                             let go_tco
                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                 go_tco.clone();
                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                 |v1|
                                                                                                                                                                                                                                                                                 go_tco(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                         &&&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                                                                                                                            }))
                                                                                                                                                                                                                                                                         })))
                                                                                                                                                                                   } else {
                                                                                                                                                                                       panic!("{}",
                                                                                                                                                                                              LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Control.Monad.List.Trans.fs"),
                                  Data1: 126_i32,
                                  Data2: 240_i32,}).get_Message(),)
                                                                                                                                                                                   }
                                                                                                                                                                               }
                                                                                                                                                                           }
                                                                                                                                                                           let go =
                                                                                                                                                                               go_1.Value;
                                                                                                                                                                           go_tco(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                   b))
                                                                                                                                                                       }
                                                                                                                                                               })
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>()))
                                                                                  }))
    }
    pub fn Control_Monad_List_Trans_unfoldableListT() -> &dyn Any {
        static Control_Monad_List_Trans_unfoldableListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_unfoldableListT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictMonad|
                                                                                 {
                                                                                     let Applicative0 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                 Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                                     let unfoldable1ListT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_unfoldable1ListT(),
                                                                                                                          dictMonad);
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_Unfoldableusd_Dict(),
                                                                                                                      &&&add(string("unfoldr"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let Applicative0
                                                                                                                                                  =
                                                                                                                                                  Applicative0.clone();
                                                                                                                                              move
                                                                                                                                                  |f|
                                                                                                                                                  &Func1::new({
                                                                                                                                                                  let f
                                                                                                                                                                      =
                                                                                                                                                                      f.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |b|
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
                                                                                                                                                                          fn go_tco(v_1:
                                                                                                                                                                                        _)
                                                                                                                                                                           ->
                                                                                                                                                                               &dyn Any {
                                                                                                                                                                              let matchValue:
                                                                                                                                                                                      LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                              match matchValue.as_ref()
                                                                                                                                                                                  {
                                                                                                                                                                                  Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                  =>
                                                                                                                                                                                  {
                                                                                                                                                                                      let activePatternResult:
                                                                                                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                          Sharpurs_Prelude::_007cUnbox_007c(matchValue_1_0);
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_cons(),
                                                                                                                                                                                                                                                                                             &&&Applicative0),
                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Lazy::Data_Lazy_applicativeLazy()),
                                                                                                                                                                                                                                                                                             &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                })),
                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                                                                                                                            let activePatternResult
                                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                                activePatternResult.clone();
                                                                                                                                                                                                                                                                            let go_tco
                                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                                go_tco.clone();
                                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                                |v1|
                                                                                                                                                                                                                                                                                go_tco(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                        &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                           }))
                                                                                                                                                                                                                                                                        })))
                                                                                                                                                                                  }
                                                                                                                                                                                  _
                                                                                                                                                                                  =>
                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_nil(),
                                                                                                                                                                                                                   &&&Applicative0),
                                                                                                                                                                              }
                                                                                                                                                                          }
                                                                                                                                                                          let go =
                                                                                                                                                                              go_1.Value;
                                                                                                                                                                          go_tco(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                  b))
                                                                                                                                                                      }
                                                                                                                                                              })
                                                                                                                                          }),
                                                                                                                             add(string("Unfoldable10"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let unfoldable1ListT1
                                                                                                                                                      =
                                                                                                                                                      unfoldable1ListT1.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused|
                                                                                                                                                      &unfoldable1ListT1
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>())))
                                                                                 }))
    }
    pub fn Control_Monad_List_Trans_semigroupListT_0040141() -> &dyn Any {
        &Func1::new(move |dictApplicative|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_semigroupListT_tco(dictApplicative))
    }
    pub fn Control_Monad_List_Trans_semigroupListT_0040141_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_semigroupListT_0040141_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_semigroupListT_0040141_002d1.get_or_init(||
                                                                              Lazy(Control_Monad_List_Trans_semigroupListT_0040141.clone()))
    }
    pub fn Control_Monad_List_Trans_concat_0040143() -> &dyn Any {
        &Func1::new(move |dictApplicative|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_concat_tco(dictApplicative))
    }
    pub fn Control_Monad_List_Trans_concat_0040143_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_concat_0040143_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_concat_0040143_002d1.get_or_init(||
                                                                      Lazy(Control_Monad_List_Trans_concat_0040143.clone()))
    }
    pub fn Control_Monad_List_Trans_semigroupListT_tco(dictApplicative:
                                                           &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                         &&&add(string("append"),
                                                &PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_concat_tco(dictApplicative),
                                                empty::<string, &dyn Any>()))
    }
    pub fn Control_Monad_List_Trans_semigroupListT() -> &dyn Any {
        static Control_Monad_List_Trans_semigroupListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_semigroupListT.get_or_init(||
                                                                Control_Monad_List_Trans_semigroupListT_0040141_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_concat_tco(dictApplicative: &dyn Any)
     -> &dyn Any {
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Functor0 = Functor0.clone();
                        let dictApplicative = dictApplicative.clone();
                        move |x|
                            &Func1::new({
                                            let x = x.clone();
                                            move |y|
                                                {
                                                    let f =
                                                        &Func1::new({
                                                                        let y
                                                                            =
                                                                            y.clone();
                                                                        move
                                                                            |v|
                                                                            {
                                                                                let matchValue:
                                                                                        LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                match matchValue.as_ref()
                                                                                    {
                                                                                    PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_1_0)
                                                                                    =>
                                                                                    &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                                                                                                                 |v1_1|
                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                        &&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_semigroupListT_tco(&&dictApplicative)),
                                                                                                                                                                                                                                                                                                                                                                     v1_1),
                                                                                                                                                                                                                                                                                                                                  &&&y))),
                                                                                                                                                                                                                                            &&matchValue_1_0))),
                                                                                    PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                                                                    =>
                                                                                    &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                               &&&PureScript_Data_Lazy::Data_Lazy_defer()),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                                                               &&&y)))),
                                                                                    PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_0_0,
                                                                                                                                                                                               matchValue_0_1)
                                                                                    =>
                                                                                    &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_0_0,
                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                                                                  |v1|
                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                         &&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_semigroupListT_tco(&&dictApplicative)),
                                                                                                                                                                                                                                                                                                                                                                      v1),
                                                                                                                                                                                                                                                                                                                                   &&&y))),
                                                                                                                                                                                                                                             &&matchValue_0_1))),
                                                                                }
                                                                            }
                                                                    });
                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_stepMap(),
                                                                                                                                                           &&&Functor0),
                                                                                                                        &&&f),
                                                                                     &&&x)
                                                }
                                        })
                    })
    }
    pub fn Control_Monad_List_Trans_concat() -> &dyn Any {
        static Control_Monad_List_Trans_concat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_concat.get_or_init(||
                                                        Control_Monad_List_Trans_concat_0040143_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_monoidListT() -> &dyn Any {
        static Control_Monad_List_Trans_monoidListT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_List_Trans_monoidListT.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictApplicative|
                                                                             {
                                                                                 let semigroupListT1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_semigroupListT(),
                                                                                                                      dictApplicative);
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                  &&&add(string("mempty"),
                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_nil(),
                                                                                                                                                           dictApplicative),
                                                                                                                         add(string("Semigroup0"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let semigroupListT1
                                                                                                                                                  =
                                                                                                                                                  semigroupListT1.clone();
                                                                                                                                              move
                                                                                                                                                  |usd__unused|
                                                                                                                                                  &semigroupListT1
                                                                                                                                          }),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>())))
                                                                             }))
    }
    pub fn Control_Monad_List_Trans_catMaybes() -> &dyn Any {
        static Control_Monad_List_Trans_catMaybes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_catMaybes.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictFunctor|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_mapMaybe(),
                                                                                                                                               dictFunctor),
                                                                                                            &&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_identity1())))
    }
    pub fn Control_Monad_List_Trans_monadListT_0040151() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_monadListT_tco(dictMonad))
    }
    pub fn Control_Monad_List_Trans_monadListT_0040151_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_monadListT_0040151_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_monadListT_0040151_002d1.get_or_init(||
                                                                          Lazy(Control_Monad_List_Trans_monadListT_0040151.clone()))
    }
    pub fn Control_Monad_List_Trans_bindListT_0040153() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_bindListT_tco(dictMonad))
    }
    pub fn Control_Monad_List_Trans_bindListT_0040153_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_bindListT_0040153_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_bindListT_0040153_002d1.get_or_init(||
                                                                         Lazy(Control_Monad_List_Trans_bindListT_0040153.clone()))
    }
    pub fn Control_Monad_List_Trans_applyListT_0040155() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_applyListT_tco(dictMonad))
    }
    pub fn Control_Monad_List_Trans_applyListT_0040155_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_applyListT_0040155_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_applyListT_0040155_002d1.get_or_init(||
                                                                          Lazy(Control_Monad_List_Trans_applyListT_0040155.clone()))
    }
    pub fn Control_Monad_List_Trans_applicativeListT_0040157() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_applicativeListT_tco(dictMonad))
    }
    pub fn Control_Monad_List_Trans_applicativeListT_0040157_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_List_Trans_applicativeListT_0040157_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_List_Trans_applicativeListT_0040157_002d1.get_or_init(||
                                                                                Lazy(Control_Monad_List_Trans_applicativeListT_0040157.clone()))
    }
    pub fn Control_Monad_List_Trans_monadListT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                         &&&add(string("Applicative0"),
                                                &&Func1::new({
                                                                 let dictMonad
                                                                     =
                                                                     dictMonad.clone();
                                                                 move
                                                                     |usd__unused|
                                                                     PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_applicativeListT_tco(&&dictMonad)
                                                             }),
                                                add(string("Bind1"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused_1|
                                                                         PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_bindListT_tco(&&dictMonad)
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_List_Trans_monadListT() -> &dyn Any {
        static Control_Monad_List_Trans_monadListT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_List_Trans_monadListT.get_or_init(||
                                                            Control_Monad_List_Trans_monadListT_0040151_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_bindListT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let semigroupListT1 =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_semigroupListT(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                       Sharpurs_Prelude::unbox(dictMonad)),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                         &&&add(string("bind"),
                                                &&Func1::new({
                                                                 let Functor0
                                                                     =
                                                                     Functor0.clone();
                                                                 let dictMonad
                                                                     =
                                                                     dictMonad.clone();
                                                                 let semigroupListT1
                                                                     =
                                                                     semigroupListT1.clone();
                                                                 move |fa|
                                                                     &Func1::new({
                                                                                     let fa
                                                                                         =
                                                                                         fa.clone();
                                                                                     move
                                                                                         |f|
                                                                                         {
                                                                                             let g =
                                                                                                 &Func1::new({
                                                                                                                 let f
                                                                                                                     =
                                                                                                                     f.clone();
                                                                                                                 move
                                                                                                                     |v|
                                                                                                                     {
                                                                                                                         let matchValue:
                                                                                                                                 LrcPtr<PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step> =
                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                         match matchValue.as_ref()
                                                                                                                             {
                                                                                                                             PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(matchValue_1_0)
                                                                                                                             =>
                                                                                                                             &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                          |v1|
                                                                                                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_bindListT_tco(&&dictMonad)),
                                                                                                                                                                                                                                                                                                                                                                                                              v1),
                                                                                                                                                                                                                                                                                                                                                                           &&&f))),
                                                                                                                                                                                                                                                                                     &&matchValue_1_0))),
                                                                                                                             PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor
                                                                                                                             =>
                                                                                                                             &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Doneusd_Ctor),
                                                                                                                             PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Yieldusd_Ctor(matchValue_0_0,
                                                                                                                                                                                                                                        matchValue_0_1)
                                                                                                                             =>
                                                                                                                             {
                                                                                                                                 let h =
                                                                                                                                     &Func1::new(move
                                                                                                                                                     |s_prime|
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                            &&&semigroupListT1),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                            &&matchValue_0_0)),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                               &&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_bindListT_tco(&&dictMonad)),
                                                                                                                                                                                                                                                            s_prime),
                                                                                                                                                                                                                         &&&f)));
                                                                                                                                 &LrcPtr::new(PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_Step::Control_Monad_List_Trans_Skipusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                                                            &&&h),
                                                                                                                                                                                                                                                                                         &&matchValue_0_1)))
                                                                                                                             }
                                                                                                                         }
                                                                                                                     }
                                                                                                             });
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_stepMap(),
                                                                                                                                                                                                    &&&Functor0),
                                                                                                                                                                 &&&g),
                                                                                                                              &&&fa)
                                                                                         }
                                                                                 })
                                                             }),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_applyListT_tco(&&dictMonad)
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_List_Trans_bindListT() -> &dyn Any {
        static Control_Monad_List_Trans_bindListT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_bindListT.get_or_init(||
                                                           Control_Monad_List_Trans_bindListT_0040153_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_applyListT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        let functorListT1 =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_functorListT(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                         &&&add(string("apply"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_ap(),
                                                                                  &&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_monadListT_tco(dictMonad)),
                                                add(string("Functor0"),
                                                    &&Func1::new({
                                                                     let functorListT1
                                                                         =
                                                                         functorListT1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &functorListT1
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_List_Trans_applyListT() -> &dyn Any {
        static Control_Monad_List_Trans_applyListT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_List_Trans_applyListT.get_or_init(||
                                                            Control_Monad_List_Trans_applyListT_0040155_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_applicativeListT_tco(dictMonad: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                         &&&add(string("pure"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_singleton(),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                            Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined())),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let dictMonad
                                                                         =
                                                                         dictMonad.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_applyListT_tco(&&dictMonad)
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_List_Trans_applicativeListT() -> &dyn Any {
        static Control_Monad_List_Trans_applicativeListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_applicativeListT.get_or_init(||
                                                                  Control_Monad_List_Trans_applicativeListT_0040157_002d1.Value)
    }
    pub fn Control_Monad_List_Trans_monadEffectListT() -> &dyn Any {
        static Control_Monad_List_Trans_monadEffectListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_monadEffectListT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonadEffect|
                                                                                  {
                                                                                      let Monad0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                  Sharpurs_Prelude::unbox(dictMonadEffect)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let monadListT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_monadListT(),
                                                                                                                           &&&Monad0);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_MonadEffectusd_Dict(),
                                                                                                                       &&&add(string("liftEffect"),
                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_lift(),
                                                                                                                                                                                                                                      &&&Monad0)),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                                   dictMonadEffect)),
                                                                                                                              add(string("Monad0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let monadListT1
                                                                                                                                                       =
                                                                                                                                                       monadListT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &monadListT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
    pub fn Control_Monad_List_Trans_monadSTListT() -> &dyn Any {
        static Control_Monad_List_Trans_monadSTListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_monadSTListT.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictMonadST|
                                                                              {
                                                                                  let Monad0 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                              Sharpurs_Prelude::unbox(dictMonadST)),
                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                                  let monadListT1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_monadListT(),
                                                                                                                       &&&Monad0);
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                                                                   &&&add(string("liftST"),
                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_lift(),
                                                                                                                                                                                                                                  &&&Monad0)),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_liftST(),
                                                                                                                                                                                               dictMonadST)),
                                                                                                                          add(string("Monad0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let monadListT1
                                                                                                                                                   =
                                                                                                                                                   monadListT1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &monadListT1
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))
                                                                              }))
    }
    pub fn Control_Monad_List_Trans_altListT() -> &dyn Any {
        static Control_Monad_List_Trans_altListT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_altListT.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictApplicative|
                                                                          {
                                                                              let functorListT1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_functorListT(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                                               &&&add(string("alt"),
                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_concat(),
                                                                                                                                                        dictApplicative),
                                                                                                                      add(string("Functor0"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let functorListT1
                                                                                                                                               =
                                                                                                                                               functorListT1.clone();
                                                                                                                                           move
                                                                                                                                               |usd__unused|
                                                                                                                                               &functorListT1
                                                                                                                                       }),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>())))
                                                                          }))
    }
    pub fn Control_Monad_List_Trans_plusListT() -> &dyn Any {
        static Control_Monad_List_Trans_plusListT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_plusListT.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictMonad|
                                                                           {
                                                                               let Applicative0 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                           Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                               let altListT1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_altListT(),
                                                                                                                    &&&Applicative0);
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                                                &&&add(string("empty"),
                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_nil(),
                                                                                                                                                         &&&Applicative0),
                                                                                                                       add(string("Alt0"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let altListT1
                                                                                                                                                =
                                                                                                                                                altListT1.clone();
                                                                                                                                            move
                                                                                                                                                |usd__unused|
                                                                                                                                                &altListT1
                                                                                                                                        }),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>())))
                                                                           }))
    }
    pub fn Control_Monad_List_Trans_alternativeListT() -> &dyn Any {
        static Control_Monad_List_Trans_alternativeListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_alternativeListT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictMonad|
                                                                                  {
                                                                                      let applicativeListT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_applicativeListT(),
                                                                                                                           dictMonad);
                                                                                      let plusListT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_plusListT(),
                                                                                                                           dictMonad);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                                                       &&&add(string("Applicative0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let applicativeListT1
                                                                                                                                                   =
                                                                                                                                                   applicativeListT1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &applicativeListT1
                                                                                                                                           }),
                                                                                                                              add(string("Plus1"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let plusListT1
                                                                                                                                                       =
                                                                                                                                                       plusListT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused_1|
                                                                                                                                                       &plusListT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
    pub fn Control_Monad_List_Trans_monadPlusListT() -> &dyn Any {
        static Control_Monad_List_Trans_monadPlusListT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_List_Trans_monadPlusListT.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictMonad|
                                                                                {
                                                                                    let monadListT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_monadListT(),
                                                                                                                         dictMonad);
                                                                                    let alternativeListT1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_List_Trans::Control_Monad_List_Trans_alternativeListT(),
                                                                                                                         dictMonad);
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_MonadPlus::Control_MonadPlus_MonadPlususd_Dict(),
                                                                                                                     &&&add(string("Monad0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let monadListT1
                                                                                                                                                 =
                                                                                                                                                 monadListT1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &monadListT1
                                                                                                                                         }),
                                                                                                                            add(string("Alternative1"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let alternativeListT1
                                                                                                                                                     =
                                                                                                                                                     alternativeListT1.clone();
                                                                                                                                                 move
                                                                                                                                                     |usd__unused_1|
                                                                                                                                                     &alternativeListT1
                                                                                                                                             }),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>())))
                                                                                }))
    }
}
