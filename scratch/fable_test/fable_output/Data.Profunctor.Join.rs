pub mod PureScript_Data_Profunctor_Join {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Profunctor_Join_Join() -> &dyn Any {
        static Data_Profunctor_Join_Join: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Join_Join.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  x.clone()))
    }
    pub fn Data_Profunctor_Join_showJoin() -> &dyn Any {
        static Data_Profunctor_Join_showJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Join_showJoin.get_or_init(||
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
    pub fn Data_Profunctor_Join_semigroupJoin() -> &dyn Any {
        static Data_Profunctor_Join_semigroupJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Join_semigroupJoin.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictSemigroupoid|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                            &&&add(string("append"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let dictSemigroupoid
                                                                                                                                        =
                                                                                                                                        dictSemigroupoid.clone();
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
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Join::Data_Profunctor_Join_Join(),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                          &&&dictSemigroupoid),
                                                                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                                                                    &&&matchValue_1))
                                                                                                                                                            }
                                                                                                                                                    })
                                                                                                                                }),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Profunctor_Join_ordJoin() -> &dyn Any {
        static Data_Profunctor_Join_ordJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Join_ordJoin.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictOrd|
                                                                     dictOrd.clone()))
    }
    pub fn Data_Profunctor_Join_newtypeJoin() -> &dyn Any {
        static Data_Profunctor_Join_newtypeJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Join_newtypeJoin.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                          &&&add(string("Coercible0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &Sharpurs_Prelude::Prim_undefined()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>())))
    }
    pub fn Data_Profunctor_Join_monoidJoin() -> &dyn Any {
        static Data_Profunctor_Join_monoidJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Join_monoidJoin.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictCategory|
                                                                        {
                                                                            let semigroupJoin1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Join::Data_Profunctor_Join_semigroupJoin(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroupoid0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictCategory)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                             &&&add(string("mempty"),
                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Join::Data_Profunctor_Join_Join(),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                         dictCategory)),
                                                                                                                    add(string("Semigroup0"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupJoin1
                                                                                                                                             =
                                                                                                                                             semigroupJoin1.clone();
                                                                                                                                         move
                                                                                                                                             |usd__unused|
                                                                                                                                             &semigroupJoin1
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>())))
                                                                        }))
    }
    pub fn Data_Profunctor_Join_invariantJoin() -> &dyn Any {
        static Data_Profunctor_Join_invariantJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Join_invariantJoin.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictProfunctor|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                                            &&&add(string("imap"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let dictProfunctor
                                                                                                                                        =
                                                                                                                                        dictProfunctor.clone();
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
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Join::Data_Profunctor_Join_Join(),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_dimap(),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictProfunctor),
                                                                                                                                                                                                                                                                                                                              &&&matchValue_1),
                                                                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                                                                        &&&matchValue_2))
                                                                                                                                                                                }
                                                                                                                                                                        })
                                                                                                                                                    })
                                                                                                                                }),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Profunctor_Join_eqJoin() -> &dyn Any {
        static Data_Profunctor_Join_eqJoin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Join_eqJoin.get_or_init(||
                                                    &Func1::new(move |dictEq|
                                                                    dictEq.clone()))
    }
}
