pub mod PureScript_Data_Functor_Joker {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_2f2d2d01::PureScript_Control_Biapplicative;
    use crate::module_c52ab4fd::PureScript_Control_Biapply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_8c6f434a::PureScript_Data_Profunctor_Choice;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Functor_Joker_Joker() -> &dyn Any {
        static Data_Functor_Joker_Joker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_Joker.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Functor_Joker_showJoker() -> &dyn Any {
        static Data_Functor_Joker_showJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_showJoker.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictShow|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                      &&&add(string("show"),
                                                                                                             &&Func1::new({
                                                                                                                              let dictShow
                                                                                                                                  =
                                                                                                                                  dictShow.clone();
                                                                                                                              move
                                                                                                                                  |v|
                                                                                                                                  {
                                                                                                                                      let x =
                                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                          &&&string("(Joker ")),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                   &&&dictShow),
                                                                                                                                                                                                                                                                                &&&x)),
                                                                                                                                                                                                          &&&string(")")))
                                                                                                                                  }
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))
    }
    pub fn Data_Functor_Joker_profunctorJoker() -> &dyn Any {
        static Data_Functor_Joker_profunctorJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_profunctorJoker.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictFunctor|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_Profunctorusd_Dict(),
                                                                                                            &&&add(string("dimap"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let dictFunctor
                                                                                                                                        =
                                                                                                                                        dictFunctor.clone();
                                                                                                                                    move
                                                                                                                                        |v|
                                                                                                                                        &Func1::new({
                                                                                                                                                        let v
                                                                                                                                                            =
                                                                                                                                                            v.clone();
                                                                                                                                                        move
                                                                                                                                                            |g|
                                                                                                                                                            &Func1::new({
                                                                                                                                                                            let g
                                                                                                                                                                                =
                                                                                                                                                                                g.clone();
                                                                                                                                                                            move
                                                                                                                                                                                |v1|
                                                                                                                                                                                {
                                                                                                                                                                                    let matchValue =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                    let matchValue_1 =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&g);
                                                                                                                                                                                    let matchValue_2 =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker(),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                              &&&dictFunctor),
                                                                                                                                                                                                                                                                                           &&&matchValue_1),
                                                                                                                                                                                                                                                        &&&matchValue_2))
                                                                                                                                                                                }
                                                                                                                                                                        })
                                                                                                                                                    })
                                                                                                                                }),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Functor_Joker_ordJoker() -> &dyn Any {
        static Data_Functor_Joker_ordJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_ordJoker.get_or_init(||
                                                    &Func1::new(move |dictOrd|
                                                                    dictOrd.clone()))
    }
    pub fn Data_Functor_Joker_newtypeJoker() -> &dyn Any {
        static Data_Functor_Joker_newtypeJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_newtypeJoker.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                         &&&add(string("Coercible0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &Sharpurs_Prelude::Prim_undefined()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Functor_Joker_hoistJoker() -> &dyn Any {
        static Data_Functor_Joker_hoistJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_hoistJoker.get_or_init(||
                                                      &Func1::new(move |f|
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
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                  &&&matchValue_1))
                                                                                          }
                                                                                  })))
    }
    pub fn Data_Functor_Joker_functorJoker() -> &dyn Any {
        static Data_Functor_Joker_functorJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_functorJoker.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictFunctor|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                         &&&add(string("map"),
                                                                                                                &&Func1::new({
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
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker(),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                       &&&dictFunctor),
                                                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                                                 &&&matchValue_1))
                                                                                                                                                         }
                                                                                                                                                 })
                                                                                                                             }),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Functor_Joker_eqJoker() -> &dyn Any {
        static Data_Functor_Joker_eqJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_eqJoker.get_or_init(||
                                                   &Func1::new(move |dictEq|
                                                                   dictEq.clone()))
    }
    pub fn Data_Functor_Joker_choiceJoker() -> &dyn Any {
        static Data_Functor_Joker_choiceJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_choiceJoker.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictFunctor|
                                                                       {
                                                                           let profunctorJoker1 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_profunctorJoker(),
                                                                                                                dictFunctor);
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Choice::Data_Profunctor_Choice_Choiceusd_Dict(),
                                                                                                            &&&add(string("left"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let dictFunctor
                                                                                                                                        =
                                                                                                                                        dictFunctor.clone();
                                                                                                                                    move
                                                                                                                                        |v|
                                                                                                                                        {
                                                                                                                                            let f =
                                                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                &&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker()),
                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                      &&&dictFunctor),
                                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                                     |usd__arg1|
                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                &&&f))
                                                                                                                                        }
                                                                                                                                }),
                                                                                                                   add(string("right"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let dictFunctor
                                                                                                                                            =
                                                                                                                                            dictFunctor.clone();
                                                                                                                                        move
                                                                                                                                            |v_1|
                                                                                                                                            {
                                                                                                                                                let f_1 =
                                                                                                                                                    Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                    &&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker()),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                          &&&dictFunctor),
                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                         |usd__arg1_1|
                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                                                    &&&f_1))
                                                                                                                                            }
                                                                                                                                    }),
                                                                                                                       add(string("Profunctor0"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let profunctorJoker1
                                                                                                                                                =
                                                                                                                                                profunctorJoker1.clone();
                                                                                                                                            move
                                                                                                                                                |usd__unused|
                                                                                                                                                &profunctorJoker1
                                                                                                                                        }),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))))
                                                                       }))
    }
    pub fn Data_Functor_Joker_bifunctorJoker() -> &dyn Any {
        static Data_Functor_Joker_bifunctorJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_bifunctorJoker.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictFunctor|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                                           &&&add(string("bimap"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let dictFunctor
                                                                                                                                       =
                                                                                                                                       dictFunctor.clone();
                                                                                                                                   move
                                                                                                                                       |v|
                                                                                                                                       &Func1::new({
                                                                                                                                                       let v
                                                                                                                                                           =
                                                                                                                                                           v.clone();
                                                                                                                                                       move
                                                                                                                                                           |g|
                                                                                                                                                           &Func1::new({
                                                                                                                                                                           let g
                                                                                                                                                                               =
                                                                                                                                                                               g.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |v1|
                                                                                                                                                                               {
                                                                                                                                                                                   let matchValue =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                   let matchValue_1 =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&g);
                                                                                                                                                                                   let matchValue_2 =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker(),
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                             &&&dictFunctor),
                                                                                                                                                                                                                                                                                          &&&matchValue_1),
                                                                                                                                                                                                                                                       &&&matchValue_2))
                                                                                                                                                                               }
                                                                                                                                                                       })
                                                                                                                                                   })
                                                                                                                               }),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Functor_Joker_biapplyJoker() -> &dyn Any {
        static Data_Functor_Joker_biapplyJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_biapplyJoker.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictApply|
                                                                        {
                                                                            let bifunctorJoker1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_bifunctorJoker(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_Biapplyusd_Dict(),
                                                                                                             &&&add(string("biapply"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictApply
                                                                                                                                         =
                                                                                                                                         dictApply.clone();
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
                                                                                                                                                                     Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                 let matchValue_1 =
                                                                                                                                                                     Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker(),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                           &&&dictApply),
                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                     &&&matchValue_1))
                                                                                                                                                             }
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    add(string("Bifunctor0"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let bifunctorJoker1
                                                                                                                                             =
                                                                                                                                             bifunctorJoker1.clone();
                                                                                                                                         move
                                                                                                                                             |usd__unused|
                                                                                                                                             &bifunctorJoker1
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>())))
                                                                        }))
    }
    pub fn Data_Functor_Joker_biapplicativeJoker() -> &dyn Any {
        static Data_Functor_Joker_biapplicativeJoker:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_biapplicativeJoker.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictApplicative|
                                                                              {
                                                                                  let biapplyJoker1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_biapplyJoker(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapplicative::Control_Biapplicative_Biapplicativeusd_Dict(),
                                                                                                                   &&&add(string("bipure"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictApplicative
                                                                                                                                               =
                                                                                                                                               dictApplicative.clone();
                                                                                                                                           move
                                                                                                                                               |v|
                                                                                                                                               &Func1::new(move
                                                                                                                                                               |b|
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker(),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                      &&&dictApplicative),
                                                                                                                                                                                                                                   b)))
                                                                                                                                       }),
                                                                                                                          add(string("Biapply0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let biapplyJoker1
                                                                                                                                                   =
                                                                                                                                                   biapplyJoker1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &biapplyJoker1
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))
                                                                              }))
    }
    pub fn Data_Functor_Joker_applyJoker() -> &dyn Any {
        static Data_Functor_Joker_applyJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_applyJoker.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictApply|
                                                                      {
                                                                          let functorJoker1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_functorJoker(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                         Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                           &&&add(string("apply"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let dictApply
                                                                                                                                       =
                                                                                                                                       dictApply.clone();
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
                                                                                                                                                                   Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                               let matchValue_1 =
                                                                                                                                                                   Sharpurs_Prelude::unbox(v1);
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                   &&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker()),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                         &&&dictApply),
                                                                                                                                                                                                                                                                      &&&matchValue),
                                                                                                                                                                                                                                   &&&matchValue_1))
                                                                                                                                                           }
                                                                                                                                                   })
                                                                                                                               }),
                                                                                                                  add(string("Functor0"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let functorJoker1
                                                                                                                                           =
                                                                                                                                           functorJoker1.clone();
                                                                                                                                       move
                                                                                                                                           |usd__unused|
                                                                                                                                           &functorJoker1
                                                                                                                                   }),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>())))
                                                                      }))
    }
    pub fn Data_Functor_Joker_bindJoker() -> &dyn Any {
        static Data_Functor_Joker_bindJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_bindJoker.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBind|
                                                                     {
                                                                         let applyJoker1 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_applyJoker(),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                        Sharpurs_Prelude::unbox(dictBind)),
                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                                          &&&add(string("bind"),
                                                                                                                 &&Func1::new({
                                                                                                                                  let dictBind
                                                                                                                                      =
                                                                                                                                      dictBind.clone();
                                                                                                                                  move
                                                                                                                                      |v|
                                                                                                                                      &Func1::new({
                                                                                                                                                      let v
                                                                                                                                                          =
                                                                                                                                                          v.clone();
                                                                                                                                                      move
                                                                                                                                                          |amb|
                                                                                                                                                          {
                                                                                                                                                              let matchValue =
                                                                                                                                                                  Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                              let matchValue_1 =
                                                                                                                                                                  Sharpurs_Prelude::unbox(amb);
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                  &&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker()),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                        &&&dictBind),
                                                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                        &&&matchValue_1),
                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_un(),
                                                                                                                                                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker()))))
                                                                                                                                                          }
                                                                                                                                                  })
                                                                                                                              }),
                                                                                                                 add(string("Apply0"),
                                                                                                                     &&Func1::new({
                                                                                                                                      let applyJoker1
                                                                                                                                          =
                                                                                                                                          applyJoker1.clone();
                                                                                                                                      move
                                                                                                                                          |usd__unused|
                                                                                                                                          &applyJoker1
                                                                                                                                  }),
                                                                                                                     empty::<string,
                                                                                                                             &dyn Any>())))
                                                                     }))
    }
    pub fn Data_Functor_Joker_applicativeJoker() -> &dyn Any {
        static Data_Functor_Joker_applicativeJoker: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Joker_applicativeJoker.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictApplicative|
                                                                            {
                                                                                let applyJoker1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_applyJoker(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                               Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                                 &&&add(string("pure"),
                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                             &&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker()),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                             dictApplicative)),
                                                                                                                        add(string("Apply0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let applyJoker1
                                                                                                                                                 =
                                                                                                                                                 applyJoker1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &applyJoker1
                                                                                                                                         }),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>())))
                                                                            }))
    }
    pub fn Data_Functor_Joker_monadJoker() -> &dyn Any {
        static Data_Functor_Joker_monadJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Joker_monadJoker.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonad|
                                                                      {
                                                                          let applicativeJoker1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_applicativeJoker(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                         Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                          let bindJoker1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_bindJoker(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                         Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                                           &&&add(string("Applicative0"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let applicativeJoker1
                                                                                                                                       =
                                                                                                                                       applicativeJoker1.clone();
                                                                                                                                   move
                                                                                                                                       |usd__unused|
                                                                                                                                       &applicativeJoker1
                                                                                                                               }),
                                                                                                                  add(string("Bind1"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let bindJoker1
                                                                                                                                           =
                                                                                                                                           bindJoker1.clone();
                                                                                                                                       move
                                                                                                                                           |usd__unused_1|
                                                                                                                                           &bindJoker1
                                                                                                                                   }),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>())))
                                                                      }))
    }
}
