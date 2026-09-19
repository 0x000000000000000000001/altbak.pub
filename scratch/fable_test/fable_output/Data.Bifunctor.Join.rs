pub mod PureScript_Data_Bifunctor_Join {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_2f2d2d01::PureScript_Control_Biapplicative;
    use crate::module_c52ab4fd::PureScript_Control_Biapply;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Bifunctor_Join_Join() -> &dyn Any {
        static Data_Bifunctor_Join_Join: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_Join_Join.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Bifunctor_Join_showJoin() -> &dyn Any {
        static Data_Bifunctor_Join_showJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_Join_showJoin.get_or_init(||
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
                                                                                                                                                                                                          &&&string("(Join ")),
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
    pub fn Data_Bifunctor_Join_ordJoin() -> &dyn Any {
        static Data_Bifunctor_Join_ordJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_Join_ordJoin.get_or_init(||
                                                    &Func1::new(move |dictOrd|
                                                                    dictOrd.clone()))
    }
    pub fn Data_Bifunctor_Join_newtypeJoin() -> &dyn Any {
        static Data_Bifunctor_Join_newtypeJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_Join_newtypeJoin.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                         &&&add(string("Coercible0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &Sharpurs_Prelude::Prim_undefined()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Bifunctor_Join_eqJoin() -> &dyn Any {
        static Data_Bifunctor_Join_eqJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_Join_eqJoin.get_or_init(||
                                                   &Func1::new(move |dictEq|
                                                                   dictEq.clone()))
    }
    pub fn Data_Bifunctor_Join_bifunctorJoin() -> &dyn Any {
        static Data_Bifunctor_Join_bifunctorJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_Join_bifunctorJoin.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictBifunctor|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                           &&&add(string("map"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let dictBifunctor
                                                                                                                                       =
                                                                                                                                       dictBifunctor.clone();
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
                                                                                                                                                               let f1 =
                                                                                                                                                                   matchValue;
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor_Join::Data_Bifunctor_Join_Join(),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                            &&&dictBifunctor),
                                                                                                                                                                                                                                                                                                         &&&f1),
                                                                                                                                                                                                                                                                      &&&f1),
                                                                                                                                                                                                                                   &&&matchValue_1))
                                                                                                                                                           }
                                                                                                                                                   })
                                                                                                                               }),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Bifunctor_Join_biapplyJoin() -> &dyn Any {
        static Data_Bifunctor_Join_biapplyJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_Join_biapplyJoin.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictBiapply|
                                                                        {
                                                                            let bifunctorJoin1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor_Join::Data_Bifunctor_Join_bifunctorJoin(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictBiapply)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                             &&&add(string("apply"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictBiapply
                                                                                                                                         =
                                                                                                                                         dictBiapply.clone();
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
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor_Join::Data_Bifunctor_Join_Join(),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_biapply(),
                                                                                                                                                                                                                                                                                                           &&&dictBiapply),
                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                     &&&matchValue_1))
                                                                                                                                                             }
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    add(string("Functor0"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let bifunctorJoin1
                                                                                                                                             =
                                                                                                                                             bifunctorJoin1.clone();
                                                                                                                                         move
                                                                                                                                             |usd__unused|
                                                                                                                                             &bifunctorJoin1
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>())))
                                                                        }))
    }
    pub fn Data_Bifunctor_Join_biapplicativeJoin() -> &dyn Any {
        static Data_Bifunctor_Join_biapplicativeJoin:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_Join_biapplicativeJoin.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictBiapplicative|
                                                                              {
                                                                                  let biapplyJoin1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor_Join::Data_Bifunctor_Join_biapplyJoin(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Biapply0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictBiapplicative)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                                   &&&add(string("pure"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictBiapplicative
                                                                                                                                               =
                                                                                                                                               dictBiapplicative.clone();
                                                                                                                                           move
                                                                                                                                               |a|
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor_Join::Data_Bifunctor_Join_Join(),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapplicative::Control_Biapplicative_bipure(),
                                                                                                                                                                                                                                                                                         &&&dictBiapplicative),
                                                                                                                                                                                                                                                      a),
                                                                                                                                                                                                                   a))
                                                                                                                                       }),
                                                                                                                          add(string("Apply0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let biapplyJoin1
                                                                                                                                                   =
                                                                                                                                                   biapplyJoin1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &biapplyJoin1
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))
                                                                              }))
    }
}
