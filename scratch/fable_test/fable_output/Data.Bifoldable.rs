pub mod PureScript_Data_Bifoldable {
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
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_3079b2b5::PureScript_Data_Functor_Product2::Data_Functor_Product2_Product2;
    use crate::module_89cc47bd::PureScript_Data_Monoid_Conj;
    use crate::module_6a9d0e21::PureScript_Data_Monoid_Disj;
    use crate::module_5023b7e9::PureScript_Data_Monoid_Dual;
    use crate::module_f1079b5::PureScript_Data_Monoid_Endo;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_Bifoldable_identity() -> &dyn Any {
        static Data_Bifoldable_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_identity.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                  &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Bifoldable_identity1() -> &dyn Any {
        static Data_Bifoldable_identity1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_identity1.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                   &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Bifoldable_monoidDual() -> &dyn Any {
        static Data_Bifoldable_monoidDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_monoidDual.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_monoidDual(),
                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_monoidEndo(),
                                                                                                                       &&&PureScript_Control_Category::Control_Category_categoryFn())))
    }
    pub fn Data_Bifoldable_monoidEndo() -> &dyn Any {
        static Data_Bifoldable_monoidEndo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_monoidEndo.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_monoidEndo(),
                                                                                    &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Bifoldable_identity2() -> &dyn Any {
        static Data_Bifoldable_identity2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_identity2.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                   &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Bifoldable_unwrap() -> &dyn Any {
        static Data_Bifoldable_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_unwrap.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Bifoldable_Bifoldableusd_Dict() -> &dyn Any {
        static Data_Bifoldable_Bifoldableusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_Bifoldableusd_Dict.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Data_Bifoldable_bifoldr() -> &dyn Any {
        static Data_Bifoldable_bifoldr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldr.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("bifoldr"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bifoldable_bitraverse_() -> &dyn Any {
        static Data_Bifoldable_bitraverse_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bitraverse_.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictBifoldable|
                                                                    &Func1::new({
                                                                                    let dictBifoldable
                                                                                        =
                                                                                        dictBifoldable.clone();
                                                                                    move
                                                                                        |dictApplicative|
                                                                                        {
                                                                                            let Apply0 =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                        Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                                            let applySecond =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_applySecond(),
                                                                                                                                 &&&Apply0);
                                                                                            let applySecond1 =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_applySecond(),
                                                                                                                                 &&&Apply0);
                                                                                            &Func1::new({
                                                                                                            let applySecond
                                                                                                                =
                                                                                                                applySecond.clone();
                                                                                                            let applySecond1
                                                                                                                =
                                                                                                                applySecond1.clone();
                                                                                                            let dictApplicative
                                                                                                                =
                                                                                                                dictApplicative.clone();
                                                                                                            move
                                                                                                                |f|
                                                                                                                &Func1::new({
                                                                                                                                let f
                                                                                                                                    =
                                                                                                                                    f.clone();
                                                                                                                                move
                                                                                                                                    |g|
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldr(),
                                                                                                                                                                                                                                                                              &&&dictBifoldable),
                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                 &&&applySecond),
                                                                                                                                                                                                                                                                              &&&f)),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                              &&&applySecond1),
                                                                                                                                                                                                                                           g)),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                           &&&dictApplicative),
                                                                                                                                                                                                        &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                            })
                                                                                                        })
                                                                                        }
                                                                                })))
    }
    pub fn Data_Bifoldable_bifor_() -> &dyn Any {
        static Data_Bifoldable_bifor_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifor_.get_or_init(||
                                               &Func1::new(move
                                                               |dictBifoldable|
                                                               &Func1::new({
                                                                               let dictBifoldable
                                                                                   =
                                                                                   dictBifoldable.clone();
                                                                               move
                                                                                   |dictApplicative|
                                                                                   &Func1::new({
                                                                                                   let dictApplicative
                                                                                                       =
                                                                                                       dictApplicative.clone();
                                                                                                   move
                                                                                                       |t|
                                                                                                       &Func1::new({
                                                                                                                       let t
                                                                                                                           =
                                                                                                                           t.clone();
                                                                                                                       move
                                                                                                                           |f|
                                                                                                                           &Func1::new({
                                                                                                                                           let f
                                                                                                                                               =
                                                                                                                                               f.clone();
                                                                                                                                           move
                                                                                                                                               |g|
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bitraverse_(),
                                                                                                                                                                                                                                                                                                                            &&&dictBifoldable),
                                                                                                                                                                                                                                                                                         &&&dictApplicative),
                                                                                                                                                                                                                                                      &&&f),
                                                                                                                                                                                                                   g),
                                                                                                                                                                                &&&t)
                                                                                                                                       })
                                                                                                                   })
                                                                                               })
                                                                           })))
    }
    pub fn Data_Bifoldable_bisequence_() -> &dyn Any {
        static Data_Bifoldable_bisequence_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bisequence_.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictBifoldable|
                                                                    &Func1::new({
                                                                                    let dictBifoldable
                                                                                        =
                                                                                        dictBifoldable.clone();
                                                                                    move
                                                                                        |dictApplicative|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bitraverse_(),
                                                                                                                                                                                                                                  &&&dictBifoldable),
                                                                                                                                                                                               dictApplicative),
                                                                                                                                                            &&&PureScript_Data_Bifoldable::Data_Bifoldable_identity()),
                                                                                                                         &&&PureScript_Data_Bifoldable::Data_Bifoldable_identity1())
                                                                                })))
    }
    pub fn Data_Bifoldable_bifoldl() -> &dyn Any {
        static Data_Bifoldable_bifoldl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldl.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("bifoldl"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bifoldable_bifoldableTuple() -> &dyn Any {
        static Data_Bifoldable_bifoldableTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldableTuple.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_Bifoldableusd_Dict(),
                                                                                         &&&add(string("bifoldMap"),
                                                                                                &&Func1::new(move
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
                                                                                                                                                                                     let matchValue_2:
                                                                                                                                                                                             LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                            &&&Semigroup0),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                            &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                               })),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                         &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                            }))
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                     })
                                                                                                                                 })
                                                                                                                 }),
                                                                                                add(string("bifoldr"),
                                                                                                    &&Func1::new(move
                                                                                                                     |f_1|
                                                                                                                     &Func1::new({
                                                                                                                                     let f_1
                                                                                                                                         =
                                                                                                                                         f_1.clone();
                                                                                                                                     move
                                                                                                                                         |g_1|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let g_1
                                                                                                                                                             =
                                                                                                                                                             g_1.clone();
                                                                                                                                                         move
                                                                                                                                                             |z|
                                                                                                                                                             &Func1::new({
                                                                                                                                                                             let z
                                                                                                                                                                                 =
                                                                                                                                                                                 z.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |v_1|
                                                                                                                                                                                 {
                                                                                                                                                                                     let matchValue_4 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                                     let matchValue_5 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&g_1);
                                                                                                                                                                                     let matchValue_6 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                                                     let matchValue_7:
                                                                                                                                                                                             LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                                         &&&match matchValue_7.as_ref()
                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_5,
                                                                                                                                                                                                                                                                                            &&&match matchValue_7.as_ref()
                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                         &&&matchValue_6))
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                     })
                                                                                                                                 })),
                                                                                                    add(string("bifoldl"),
                                                                                                        &&Func1::new(move
                                                                                                                         |f_2|
                                                                                                                         &Func1::new({
                                                                                                                                         let f_2
                                                                                                                                             =
                                                                                                                                             f_2.clone();
                                                                                                                                         move
                                                                                                                                             |g_2|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let g_2
                                                                                                                                                                 =
                                                                                                                                                                 g_2.clone();
                                                                                                                                                             move
                                                                                                                                                                 |z_1|
                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                 let z_1
                                                                                                                                                                                     =
                                                                                                                                                                                     z_1.clone();
                                                                                                                                                                                 move
                                                                                                                                                                                     |v_2|
                                                                                                                                                                                     {
                                                                                                                                                                                         let matchValue_9 =
                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                                         let matchValue_10 =
                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&g_2);
                                                                                                                                                                                         let matchValue_11 =
                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                                                         let matchValue_12:
                                                                                                                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                             Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_10,
                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_9,
                                                                                                                                                                                                                                                                                                                                   &&&matchValue_11),
                                                                                                                                                                                                                                                                                                &&&match matchValue_12.as_ref()
                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                   })),
                                                                                                                                                                                                                          &&&match matchValue_12.as_ref()
                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                             })
                                                                                                                                                                                     }
                                                                                                                                                                             })
                                                                                                                                                         })
                                                                                                                                     })),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))))
    }
    pub fn Data_Bifoldable_bifoldableJoker() -> &dyn Any {
        static Data_Bifoldable_bifoldableJoker: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldableJoker.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictFoldable|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_Bifoldableusd_Dict(),
                                                                                                         &&&add(string("bifoldr"),
                                                                                                                &&Func1::new({
                                                                                                                                 let dictFoldable
                                                                                                                                     =
                                                                                                                                     dictFoldable.clone();
                                                                                                                                 move
                                                                                                                                     |v|
                                                                                                                                     &Func1::new({
                                                                                                                                                     let v
                                                                                                                                                         =
                                                                                                                                                         v.clone();
                                                                                                                                                     move
                                                                                                                                                         |r|
                                                                                                                                                         &Func1::new({
                                                                                                                                                                         let r
                                                                                                                                                                             =
                                                                                                                                                                             r.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |u|
                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                             let u
                                                                                                                                                                                                 =
                                                                                                                                                                                                 u.clone();
                                                                                                                                                                                             move
                                                                                                                                                                                                 |v1|
                                                                                                                                                                                                 {
                                                                                                                                                                                                     let matchValue =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                     let matchValue_1 =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&r);
                                                                                                                                                                                                     let matchValue_2 =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&u);
                                                                                                                                                                                                     let matchValue_3 =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                               &&&dictFoldable),
                                                                                                                                                                                                                                                                                                            &&&matchValue_1),
                                                                                                                                                                                                                                                                         &&&matchValue_2),
                                                                                                                                                                                                                                      &&&matchValue_3)
                                                                                                                                                                                                 }
                                                                                                                                                                                         })
                                                                                                                                                                     })
                                                                                                                                                 })
                                                                                                                             }),
                                                                                                                add(string("bifoldl"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictFoldable
                                                                                                                                         =
                                                                                                                                         dictFoldable.clone();
                                                                                                                                     move
                                                                                                                                         |v_1|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let v_1
                                                                                                                                                             =
                                                                                                                                                             v_1.clone();
                                                                                                                                                         move
                                                                                                                                                             |r_1|
                                                                                                                                                             &Func1::new({
                                                                                                                                                                             let r_1
                                                                                                                                                                                 =
                                                                                                                                                                                 r_1.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |u_1|
                                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                                 let u_1
                                                                                                                                                                                                     =
                                                                                                                                                                                                     u_1.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     |v1_1|
                                                                                                                                                                                                     {
                                                                                                                                                                                                         let matchValue_5 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                                         let matchValue_6 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&r_1);
                                                                                                                                                                                                         let matchValue_7 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&u_1);
                                                                                                                                                                                                         let matchValue_8 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                                                   &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                &&&matchValue_6),
                                                                                                                                                                                                                                                                             &&&matchValue_7),
                                                                                                                                                                                                                                          &&&matchValue_8)
                                                                                                                                                                                                     }
                                                                                                                                                                                             })
                                                                                                                                                                         })
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    add(string("bifoldMap"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let dictFoldable
                                                                                                                                             =
                                                                                                                                             dictFoldable.clone();
                                                                                                                                         move
                                                                                                                                             |dictMonoid|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let dictMonoid
                                                                                                                                                                 =
                                                                                                                                                                 dictMonoid.clone();
                                                                                                                                                             move
                                                                                                                                                                 |v_2|
                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                 let v_2
                                                                                                                                                                                     =
                                                                                                                                                                                     v_2.clone();
                                                                                                                                                                                 move
                                                                                                                                                                                     |r_2|
                                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                                     let r_2
                                                                                                                                                                                                         =
                                                                                                                                                                                                         r_2.clone();
                                                                                                                                                                                                     move
                                                                                                                                                                                                         |v1_2|
                                                                                                                                                                                                         {
                                                                                                                                                                                                             let matchValue_10 =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                                                             let matchValue_11 =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&r_2);
                                                                                                                                                                                                             let matchValue_12 =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(v1_2);
                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                                       &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                    &&&dictMonoid),
                                                                                                                                                                                                                                                                                 &&&matchValue_11),
                                                                                                                                                                                                                                              &&&matchValue_12)
                                                                                                                                                                                                         }
                                                                                                                                                                                                 })
                                                                                                                                                                             })
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))))))
    }
    pub fn Data_Bifoldable_bifoldableEither() -> &dyn Any {
        static Data_Bifoldable_bifoldableEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldableEither.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_Bifoldableusd_Dict(),
                                                                                          &&&add(string("bifoldr"),
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
                                                                                                                                                          &Func1::new({
                                                                                                                                                                          let v2
                                                                                                                                                                              =
                                                                                                                                                                              v2.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |v3|
                                                                                                                                                                              {
                                                                                                                                                                                  let matchValue =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                  let matchValue_1 =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v1);
                                                                                                                                                                                  let matchValue_2 =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v2);
                                                                                                                                                                                  let matchValue_3:
                                                                                                                                                                                          LrcPtr<Data_Either_Either> =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(v3);
                                                                                                                                                                                  match matchValue_3.as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                      Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_3_1_0)
                                                                                                                                                                                      =>
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                          &&matchValue_3_1_0),
                                                                                                                                                                                                                       &&&matchValue_2),
                                                                                                                                                                                      Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_3_0_0)
                                                                                                                                                                                      =>
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                          &&matchValue_3_0_0),
                                                                                                                                                                                                                       &&&matchValue_2),
                                                                                                                                                                                  }
                                                                                                                                                                              }
                                                                                                                                                                      })
                                                                                                                                                  })
                                                                                                                              })),
                                                                                                 add(string("bifoldl"),
                                                                                                     &&Func1::new(move
                                                                                                                      |v_1|
                                                                                                                      &Func1::new({
                                                                                                                                      let v_1
                                                                                                                                          =
                                                                                                                                          v_1.clone();
                                                                                                                                      move
                                                                                                                                          |v1_1|
                                                                                                                                          &Func1::new({
                                                                                                                                                          let v1_1
                                                                                                                                                              =
                                                                                                                                                              v1_1.clone();
                                                                                                                                                          move
                                                                                                                                                              |v2_1|
                                                                                                                                                              &Func1::new({
                                                                                                                                                                              let v2_1
                                                                                                                                                                                  =
                                                                                                                                                                                  v2_1.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |v3_1|
                                                                                                                                                                                  {
                                                                                                                                                                                      let matchValue_5 =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                      let matchValue_6 =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v1_1);
                                                                                                                                                                                      let matchValue_7 =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v2_1);
                                                                                                                                                                                      let matchValue_8:
                                                                                                                                                                                              LrcPtr<Data_Either_Either> =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(v3_1);
                                                                                                                                                                                      match matchValue_8.as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                          Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_8_1_0)
                                                                                                                                                                                          =>
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_6,
                                                                                                                                                                                                                                                              &&&matchValue_7),
                                                                                                                                                                                                                           &&matchValue_8_1_0),
                                                                                                                                                                                          Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_8_0_0)
                                                                                                                                                                                          =>
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_5,
                                                                                                                                                                                                                                                              &&&matchValue_7),
                                                                                                                                                                                                                           &&matchValue_8_0_0),
                                                                                                                                                                                      }
                                                                                                                                                                                  }
                                                                                                                                                                          })
                                                                                                                                                      })
                                                                                                                                  })),
                                                                                                     add(string("bifoldMap"),
                                                                                                         &&Func1::new(move
                                                                                                                          |dictMonoid|
                                                                                                                          &Func1::new(move
                                                                                                                                          |v_2|
                                                                                                                                          &Func1::new({
                                                                                                                                                          let v_2
                                                                                                                                                              =
                                                                                                                                                              v_2.clone();
                                                                                                                                                          move
                                                                                                                                                              |v1_2|
                                                                                                                                                              &Func1::new({
                                                                                                                                                                              let v1_2
                                                                                                                                                                                  =
                                                                                                                                                                                  v1_2.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |v2_2|
                                                                                                                                                                                  {
                                                                                                                                                                                      let matchValue_10 =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                                      let matchValue_11 =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v1_2);
                                                                                                                                                                                      let matchValue_12:
                                                                                                                                                                                              LrcPtr<Data_Either_Either> =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(v2_2);
                                                                                                                                                                                      match matchValue_12.as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                          Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_12_1_0)
                                                                                                                                                                                          =>
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&matchValue_11,
                                                                                                                                                                                                                           &&matchValue_12_1_0),
                                                                                                                                                                                          Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_12_0_0)
                                                                                                                                                                                          =>
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&matchValue_10,
                                                                                                                                                                                                                           &&matchValue_12_0_0),
                                                                                                                                                                                      }
                                                                                                                                                                                  }
                                                                                                                                                                          })
                                                                                                                                                      }))),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))))
    }
    pub fn Data_Bifoldable_bifoldableConst() -> &dyn Any {
        static Data_Bifoldable_bifoldableConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldableConst.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_Bifoldableusd_Dict(),
                                                                                         &&&add(string("bifoldr"),
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
                                                                                                                                                         |z|
                                                                                                                                                         &Func1::new({
                                                                                                                                                                         let z
                                                                                                                                                                             =
                                                                                                                                                                             z.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v1|
                                                                                                                                                                             {
                                                                                                                                                                                 let matchValue =
                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                 let matchValue_1 =
                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                 let matchValue_2 =
                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::unbox(v1)),
                                                                                                                                                                                                                  &&&matchValue_2)
                                                                                                                                                                             }
                                                                                                                                                                     })
                                                                                                                                                 })
                                                                                                                             })),
                                                                                                add(string("bifoldl"),
                                                                                                    &&Func1::new(move
                                                                                                                     |f_1|
                                                                                                                     &Func1::new({
                                                                                                                                     let f_1
                                                                                                                                         =
                                                                                                                                         f_1.clone();
                                                                                                                                     move
                                                                                                                                         |v_1|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let v_1
                                                                                                                                                             =
                                                                                                                                                             v_1.clone();
                                                                                                                                                         move
                                                                                                                                                             |z_1|
                                                                                                                                                             &Func1::new({
                                                                                                                                                                             let z_1
                                                                                                                                                                                 =
                                                                                                                                                                                 z_1.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |v1_1|
                                                                                                                                                                                 {
                                                                                                                                                                                     let matchValue_5 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                                     let matchValue_6 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                     let matchValue_7 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                                                     let matchValue_8 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_5,
                                                                                                                                                                                                                                                         &&&matchValue_7),
                                                                                                                                                                                                                      &&&matchValue_8)
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                     })
                                                                                                                                 })),
                                                                                                    add(string("bifoldMap"),
                                                                                                        &&Func1::new(move
                                                                                                                         |dictMonoid|
                                                                                                                         &Func1::new(move
                                                                                                                                         |f_2|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let f_2
                                                                                                                                                             =
                                                                                                                                                             f_2.clone();
                                                                                                                                                         move
                                                                                                                                                             |v_2|
                                                                                                                                                             &Func1::new({
                                                                                                                                                                             let v_2
                                                                                                                                                                                 =
                                                                                                                                                                                 v_2.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |v1_2|
                                                                                                                                                                                 {
                                                                                                                                                                                     let matchValue_10 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                                     let matchValue_11 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&matchValue_10,
                                                                                                                                                                                                                      &&&Sharpurs_Prelude::unbox(v1_2))
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                     }))),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))))
    }
    pub fn Data_Bifoldable_bifoldableClown() -> &dyn Any {
        static Data_Bifoldable_bifoldableClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldableClown.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictFoldable|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_Bifoldableusd_Dict(),
                                                                                                         &&&add(string("bifoldr"),
                                                                                                                &&Func1::new({
                                                                                                                                 let dictFoldable
                                                                                                                                     =
                                                                                                                                     dictFoldable.clone();
                                                                                                                                 move
                                                                                                                                     |l|
                                                                                                                                     &Func1::new({
                                                                                                                                                     let l
                                                                                                                                                         =
                                                                                                                                                         l.clone();
                                                                                                                                                     move
                                                                                                                                                         |v|
                                                                                                                                                         &Func1::new({
                                                                                                                                                                         let v
                                                                                                                                                                             =
                                                                                                                                                                             v.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |u|
                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                             let u
                                                                                                                                                                                                 =
                                                                                                                                                                                                 u.clone();
                                                                                                                                                                                             move
                                                                                                                                                                                                 |v1|
                                                                                                                                                                                                 {
                                                                                                                                                                                                     let matchValue =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&l);
                                                                                                                                                                                                     let matchValue_1 =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                     let matchValue_2 =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&u);
                                                                                                                                                                                                     let matchValue_3 =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                               &&&dictFoldable),
                                                                                                                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                                                                                                                         &&&matchValue_2),
                                                                                                                                                                                                                                      &&&matchValue_3)
                                                                                                                                                                                                 }
                                                                                                                                                                                         })
                                                                                                                                                                     })
                                                                                                                                                 })
                                                                                                                             }),
                                                                                                                add(string("bifoldl"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictFoldable
                                                                                                                                         =
                                                                                                                                         dictFoldable.clone();
                                                                                                                                     move
                                                                                                                                         |l_1|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let l_1
                                                                                                                                                             =
                                                                                                                                                             l_1.clone();
                                                                                                                                                         move
                                                                                                                                                             |v_1|
                                                                                                                                                             &Func1::new({
                                                                                                                                                                             let v_1
                                                                                                                                                                                 =
                                                                                                                                                                                 v_1.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |u_1|
                                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                                 let u_1
                                                                                                                                                                                                     =
                                                                                                                                                                                                     u_1.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     |v1_1|
                                                                                                                                                                                                     {
                                                                                                                                                                                                         let matchValue_5 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&l_1);
                                                                                                                                                                                                         let matchValue_6 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                                         let matchValue_7 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&u_1);
                                                                                                                                                                                                         let matchValue_8 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                                                   &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                &&&matchValue_5),
                                                                                                                                                                                                                                                                             &&&matchValue_7),
                                                                                                                                                                                                                                          &&&matchValue_8)
                                                                                                                                                                                                     }
                                                                                                                                                                                             })
                                                                                                                                                                         })
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    add(string("bifoldMap"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let dictFoldable
                                                                                                                                             =
                                                                                                                                             dictFoldable.clone();
                                                                                                                                         move
                                                                                                                                             |dictMonoid|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let dictMonoid
                                                                                                                                                                 =
                                                                                                                                                                 dictMonoid.clone();
                                                                                                                                                             move
                                                                                                                                                                 |l_2|
                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                 let l_2
                                                                                                                                                                                     =
                                                                                                                                                                                     l_2.clone();
                                                                                                                                                                                 move
                                                                                                                                                                                     |v_2|
                                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                                     let v_2
                                                                                                                                                                                                         =
                                                                                                                                                                                                         v_2.clone();
                                                                                                                                                                                                     move
                                                                                                                                                                                                         |v1_2|
                                                                                                                                                                                                         {
                                                                                                                                                                                                             let matchValue_10 =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&l_2);
                                                                                                                                                                                                             let matchValue_11 =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                                                             let matchValue_12 =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(v1_2);
                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                                       &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                    &&&dictMonoid),
                                                                                                                                                                                                                                                                                 &&&matchValue_10),
                                                                                                                                                                                                                                              &&&matchValue_12)
                                                                                                                                                                                                         }
                                                                                                                                                                                                 })
                                                                                                                                                                             })
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))))))
    }
    pub fn Data_Bifoldable_bifoldMapDefaultR() -> &dyn Any {
        static Data_Bifoldable_bifoldMapDefaultR: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldMapDefaultR.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictBifoldable|
                                                                          &Func1::new({
                                                                                          let dictBifoldable
                                                                                              =
                                                                                              dictBifoldable.clone();
                                                                                          move
                                                                                              |dictMonoid|
                                                                                              {
                                                                                                  let append =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                  let mempty =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                       dictMonoid);
                                                                                                  &Func1::new({
                                                                                                                  let append
                                                                                                                      =
                                                                                                                      append.clone();
                                                                                                                  let mempty
                                                                                                                      =
                                                                                                                      mempty.clone();
                                                                                                                  move
                                                                                                                      |f|
                                                                                                                      &Func1::new({
                                                                                                                                      let f
                                                                                                                                          =
                                                                                                                                          f.clone();
                                                                                                                                      move
                                                                                                                                          |g|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldr(),
                                                                                                                                                                                                                                                                                    &&&dictBifoldable),
                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                       &&&append),
                                                                                                                                                                                                                                                                                    &&&f)),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                    &&&append),
                                                                                                                                                                                                                                                 g)),
                                                                                                                                                                           &&&mempty)
                                                                                                                                  })
                                                                                                              })
                                                                                              }
                                                                                      })))
    }
    pub fn Data_Bifoldable_bifoldMapDefaultL() -> &dyn Any {
        static Data_Bifoldable_bifoldMapDefaultL: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldMapDefaultL.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictBifoldable|
                                                                          &Func1::new({
                                                                                          let dictBifoldable
                                                                                              =
                                                                                              dictBifoldable.clone();
                                                                                          move
                                                                                              |dictMonoid|
                                                                                              {
                                                                                                  let Semigroup0 =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                              Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                  let mempty =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                       dictMonoid);
                                                                                                  &Func1::new({
                                                                                                                  let Semigroup0
                                                                                                                      =
                                                                                                                      Semigroup0.clone();
                                                                                                                  let mempty
                                                                                                                      =
                                                                                                                      mempty.clone();
                                                                                                                  move
                                                                                                                      |f|
                                                                                                                      &Func1::new({
                                                                                                                                      let f
                                                                                                                                          =
                                                                                                                                          f.clone();
                                                                                                                                      move
                                                                                                                                          |g|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldl(),
                                                                                                                                                                                                                                                                                    &&&dictBifoldable),
                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                   |m|
                                                                                                                                                                                                                                                                   &Func1::new({
                                                                                                                                                                                                                                                                                   let m
                                                                                                                                                                                                                                                                                       =
                                                                                                                                                                                                                                                                                       m.clone();
                                                                                                                                                                                                                                                                                   move
                                                                                                                                                                                                                                                                                       |a|
                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                              &&&Semigroup0),
                                                                                                                                                                                                                                                                                                                                                           &&&m),
                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                           a))
                                                                                                                                                                                                                                                                               }))),
                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                let g
                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                    g.clone();
                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                    |m_1|
                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                    let m_1
                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                        m_1.clone();
                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                        |b|
                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                               &&&Semigroup0),
                                                                                                                                                                                                                                                                                                                            &&&m_1),
                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                                                                                                            b))
                                                                                                                                                                                                                                                })
                                                                                                                                                                                                                            })),
                                                                                                                                                                           &&&mempty)
                                                                                                                                  })
                                                                                                              })
                                                                                              }
                                                                                      })))
    }
    pub fn Data_Bifoldable_bifoldMap() -> &dyn Any {
        static Data_Bifoldable_bifoldMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldMap.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("bifoldMap"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bifoldable_bifoldableFlip() -> &dyn Any {
        static Data_Bifoldable_bifoldableFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldableFlip.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictBifoldable|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_Bifoldableusd_Dict(),
                                                                                                        &&&add(string("bifoldr"),
                                                                                                               &&Func1::new({
                                                                                                                                let dictBifoldable
                                                                                                                                    =
                                                                                                                                    dictBifoldable.clone();
                                                                                                                                move
                                                                                                                                    |r|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let r
                                                                                                                                                        =
                                                                                                                                                        r.clone();
                                                                                                                                                    move
                                                                                                                                                        |l|
                                                                                                                                                        &Func1::new({
                                                                                                                                                                        let l
                                                                                                                                                                            =
                                                                                                                                                                            l.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |u|
                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                            let u
                                                                                                                                                                                                =
                                                                                                                                                                                                u.clone();
                                                                                                                                                                                            move
                                                                                                                                                                                                |v|
                                                                                                                                                                                                {
                                                                                                                                                                                                    let matchValue =
                                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&r);
                                                                                                                                                                                                    let matchValue_1 =
                                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&l);
                                                                                                                                                                                                    let matchValue_2 =
                                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&u);
                                                                                                                                                                                                    let matchValue_3 =
                                                                                                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldr(),
                                                                                                                                                                                                                                                                                                                                                                                 &&&dictBifoldable),
                                                                                                                                                                                                                                                                                                                                              &&&matchValue_1),
                                                                                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                                                                                        &&&matchValue_2),
                                                                                                                                                                                                                                     &&&matchValue_3)
                                                                                                                                                                                                }
                                                                                                                                                                                        })
                                                                                                                                                                    })
                                                                                                                                                })
                                                                                                                            }),
                                                                                                               add(string("bifoldl"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let dictBifoldable
                                                                                                                                        =
                                                                                                                                        dictBifoldable.clone();
                                                                                                                                    move
                                                                                                                                        |r_1|
                                                                                                                                        &Func1::new({
                                                                                                                                                        let r_1
                                                                                                                                                            =
                                                                                                                                                            r_1.clone();
                                                                                                                                                        move
                                                                                                                                                            |l_1|
                                                                                                                                                            &Func1::new({
                                                                                                                                                                            let l_1
                                                                                                                                                                                =
                                                                                                                                                                                l_1.clone();
                                                                                                                                                                            move
                                                                                                                                                                                |u_1|
                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                let u_1
                                                                                                                                                                                                    =
                                                                                                                                                                                                    u_1.clone();
                                                                                                                                                                                                move
                                                                                                                                                                                                    |v_1|
                                                                                                                                                                                                    {
                                                                                                                                                                                                        let matchValue_5 =
                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&r_1);
                                                                                                                                                                                                        let matchValue_6 =
                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&l_1);
                                                                                                                                                                                                        let matchValue_7 =
                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&u_1);
                                                                                                                                                                                                        let matchValue_8 =
                                                                                                                                                                                                            Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldl(),
                                                                                                                                                                                                                                                                                                                                                                                     &&&dictBifoldable),
                                                                                                                                                                                                                                                                                                                                                  &&&matchValue_6),
                                                                                                                                                                                                                                                                                                               &&&matchValue_5),
                                                                                                                                                                                                                                                                            &&&matchValue_7),
                                                                                                                                                                                                                                         &&&matchValue_8)
                                                                                                                                                                                                    }
                                                                                                                                                                                            })
                                                                                                                                                                        })
                                                                                                                                                    })
                                                                                                                                }),
                                                                                                                   add(string("bifoldMap"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let dictBifoldable
                                                                                                                                            =
                                                                                                                                            dictBifoldable.clone();
                                                                                                                                        move
                                                                                                                                            |dictMonoid|
                                                                                                                                            &Func1::new({
                                                                                                                                                            let dictMonoid
                                                                                                                                                                =
                                                                                                                                                                dictMonoid.clone();
                                                                                                                                                            move
                                                                                                                                                                |r_2|
                                                                                                                                                                &Func1::new({
                                                                                                                                                                                let r_2
                                                                                                                                                                                    =
                                                                                                                                                                                    r_2.clone();
                                                                                                                                                                                move
                                                                                                                                                                                    |l_2|
                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                    let l_2
                                                                                                                                                                                                        =
                                                                                                                                                                                                        l_2.clone();
                                                                                                                                                                                                    move
                                                                                                                                                                                                        |v_2|
                                                                                                                                                                                                        {
                                                                                                                                                                                                            let matchValue_10 =
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&r_2);
                                                                                                                                                                                                            let matchValue_11 =
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&l_2);
                                                                                                                                                                                                            let matchValue_12 =
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMap(),
                                                                                                                                                                                                                                                                                                                                                                                         &&&dictBifoldable),
                                                                                                                                                                                                                                                                                                                                                      &&&dictMonoid),
                                                                                                                                                                                                                                                                                                                   &&&matchValue_11),
                                                                                                                                                                                                                                                                                &&&matchValue_10),
                                                                                                                                                                                                                                             &&&matchValue_12)
                                                                                                                                                                                                        }
                                                                                                                                                                                                })
                                                                                                                                                                            })
                                                                                                                                                        })
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>()))))))
    }
    pub fn Data_Bifoldable_bifoldlDefault() -> &dyn Any {
        static Data_Bifoldable_bifoldlDefault: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldlDefault.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictBifoldable|
                                                                       &Func1::new({
                                                                                       let dictBifoldable
                                                                                           =
                                                                                           dictBifoldable.clone();
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
                                                                                                                                   |z|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let z
                                                                                                                                                       =
                                                                                                                                                       z.clone();
                                                                                                                                                   move
                                                                                                                                                       |p|
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMap(),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&dictBifoldable),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Bifoldable::Data_Bifoldable_monoidDual()),
                                                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual()),
                                                                                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo()),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&f)))),
                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual()),
                                                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo()),
                                                                                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&g)))),
                                                                                                                                                                                                                                                                                                 p))),
                                                                                                                                                                                        &&&z)
                                                                                                                                               })
                                                                                                                           })
                                                                                                       })
                                                                                   })))
    }
    pub fn Data_Bifoldable_bifoldrDefault() -> &dyn Any {
        static Data_Bifoldable_bifoldrDefault: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldrDefault.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictBifoldable|
                                                                       &Func1::new({
                                                                                       let dictBifoldable
                                                                                           =
                                                                                           dictBifoldable.clone();
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
                                                                                                                                   |z|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let z
                                                                                                                                                       =
                                                                                                                                                       z.clone();
                                                                                                                                                   move
                                                                                                                                                       |p|
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMap(),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&dictBifoldable),
                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Bifoldable::Data_Bifoldable_monoidEndo()),
                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo()),
                                                                                                                                                                                                                                                                                                                                                                       &&&f)),
                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo()),
                                                                                                                                                                                                                                                                                                                                    &&&g)),
                                                                                                                                                                                                                                                              p)),
                                                                                                                                                                                        &&&z)
                                                                                                                                               })
                                                                                                                           })
                                                                                                       })
                                                                                   })))
    }
    pub fn Data_Bifoldable_bifoldableProduct2_004054() -> &dyn Any {
        &Func1::new(move |dictBifoldable|
                        Func1::new({
                                       let dictBifoldable =
                                           dictBifoldable.clone();
                                       move |dictBifoldable1|
                                           PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableProduct2_tco(&dictBifoldable,
                                                                                                              dictBifoldable1)
                                   }))
    }
    pub fn Data_Bifoldable_bifoldableProduct2_004054_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Bifoldable_bifoldableProduct2_004054_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldableProduct2_004054_002d1.get_or_init(||
                                                                        Lazy(Data_Bifoldable_bifoldableProduct2_004054.clone()))
    }
    pub fn Data_Bifoldable_bifoldableProduct2_tco(dictBifoldable: &dyn Any,
                                                  dictBifoldable1: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_Bifoldableusd_Dict(),
                                         &&&add(string("bifoldr"),
                                                &&Func1::new({
                                                                 let dictBifoldable
                                                                     =
                                                                     dictBifoldable.clone();
                                                                 let dictBifoldable1
                                                                     =
                                                                     dictBifoldable1.clone();
                                                                 move |l|
                                                                     &Func1::new({
                                                                                     let l
                                                                                         =
                                                                                         l.clone();
                                                                                     move
                                                                                         |r|
                                                                                         &Func1::new({
                                                                                                         let r
                                                                                                             =
                                                                                                             r.clone();
                                                                                                         move
                                                                                                             |u|
                                                                                                             &Func1::new({
                                                                                                                             let u
                                                                                                                                 =
                                                                                                                                 u.clone();
                                                                                                                             move
                                                                                                                                 |m|
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldrDefault(),
                                                                                                                                                                                                                                                                                                              &&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableProduct2_tco(&&dictBifoldable,
                                                                                                                                                                                                                                                                                                                                                                                   &&dictBifoldable1)),
                                                                                                                                                                                                                                                                           &&&l),
                                                                                                                                                                                                                                        &&&r),
                                                                                                                                                                                                     &&&u),
                                                                                                                                                                  m)
                                                                                                                         })
                                                                                                     })
                                                                                 })
                                                             }),
                                                add(string("bifoldl"),
                                                    &&Func1::new({
                                                                     let dictBifoldable
                                                                         =
                                                                         dictBifoldable.clone();
                                                                     let dictBifoldable1
                                                                         =
                                                                         dictBifoldable1.clone();
                                                                     move
                                                                         |l_1|
                                                                         &Func1::new({
                                                                                         let l_1
                                                                                             =
                                                                                             l_1.clone();
                                                                                         move
                                                                                             |r_1|
                                                                                             &Func1::new({
                                                                                                             let r_1
                                                                                                                 =
                                                                                                                 r_1.clone();
                                                                                                             move
                                                                                                                 |u_1|
                                                                                                                 &Func1::new({
                                                                                                                                 let u_1
                                                                                                                                     =
                                                                                                                                     u_1.clone();
                                                                                                                                 move
                                                                                                                                     |m_1|
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldlDefault(),
                                                                                                                                                                                                                                                                                                                  &&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableProduct2_tco(&&dictBifoldable,
                                                                                                                                                                                                                                                                                                                                                                                       &&dictBifoldable1)),
                                                                                                                                                                                                                                                                               &&&l_1),
                                                                                                                                                                                                                                            &&&r_1),
                                                                                                                                                                                                         &&&u_1),
                                                                                                                                                                      m_1)
                                                                                                                             })
                                                                                                         })
                                                                                     })
                                                                 }),
                                                    add(string("bifoldMap"),
                                                        &&Func1::new({
                                                                         let dictBifoldable
                                                                             =
                                                                             dictBifoldable.clone();
                                                                         let dictBifoldable1
                                                                             =
                                                                             dictBifoldable1.clone();
                                                                         move
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
                                                                                                     |l_2|
                                                                                                     &Func1::new({
                                                                                                                     let l_2
                                                                                                                         =
                                                                                                                         l_2.clone();
                                                                                                                     move
                                                                                                                         |r_2|
                                                                                                                         &Func1::new({
                                                                                                                                         let r_2
                                                                                                                                             =
                                                                                                                                             r_2.clone();
                                                                                                                                         move
                                                                                                                                             |v|
                                                                                                                                             {
                                                                                                                                                 let matchValue =
                                                                                                                                                     Sharpurs_Prelude::unbox(&&l_2);
                                                                                                                                                 let matchValue_1 =
                                                                                                                                                     Sharpurs_Prelude::unbox(&&r_2);
                                                                                                                                                 let matchValue_2:
                                                                                                                                                         LrcPtr<Data_Functor_Product2_Product2> =
                                                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                                                 let r1 =
                                                                                                                                                     matchValue_1;
                                                                                                                                                 let l1 =
                                                                                                                                                     matchValue;
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                        &&&Semigroup0),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMap(),
                                                                                                                                                                                                                                                                                                                                                                                                    &&&dictBifoldable),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictMonoid),
                                                                                                                                                                                                                                                                                                                              &&&l1),
                                                                                                                                                                                                                                                                                           &&&r1),
                                                                                                                                                                                                                                                        &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                               Data_Functor_Product2_Product2::Data_Functor_Product2_Product2usd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                           })),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMap(),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictBifoldable1),
                                                                                                                                                                                                                                                                                                                              &&&dictMonoid),
                                                                                                                                                                                                                                                                                           &&&l1),
                                                                                                                                                                                                                                                        &&&r1),
                                                                                                                                                                                                                     &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                            {
                                                                                                                                                                                                                            Data_Functor_Product2_Product2::Data_Functor_Product2_Product2usd_Ctor(_,
                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                        }))
                                                                                                                                             }
                                                                                                                                     })
                                                                                                                 })
                                                                                             })
                                                                             }
                                                                     }),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Bifoldable_bifoldableProduct2() -> &dyn Any {
        static Data_Bifoldable_bifoldableProduct2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifoldableProduct2.get_or_init(||
                                                           Data_Bifoldable_bifoldableProduct2_004054_002d1.Value)
    }
    pub fn Data_Bifoldable_bifold() -> &dyn Any {
        static Data_Bifoldable_bifold: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_bifold.get_or_init(||
                                               &Func1::new(move
                                                               |dictBifoldable|
                                                               &Func1::new({
                                                                               let dictBifoldable
                                                                                   =
                                                                                   dictBifoldable.clone();
                                                                               move
                                                                                   |dictMonoid|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMap(),
                                                                                                                                                                                                                             &&&dictBifoldable),
                                                                                                                                                                                          dictMonoid),
                                                                                                                                                       &&&PureScript_Data_Bifoldable::Data_Bifoldable_identity2()),
                                                                                                                    &&&PureScript_Data_Bifoldable::Data_Bifoldable_identity2())
                                                                           })))
    }
    pub fn Data_Bifoldable_biany() -> &dyn Any {
        static Data_Bifoldable_biany: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_biany.get_or_init(||
                                              &Func1::new(move
                                                              |dictBifoldable|
                                                              &Func1::new({
                                                                              let dictBifoldable
                                                                                  =
                                                                                  dictBifoldable.clone();
                                                                              move
                                                                                  |dictBooleanAlgebra|
                                                                                  {
                                                                                      let monoidDisj =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_monoidDisj(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("HeytingAlgebra0"),
                                                                                                                                                                     Sharpurs_Prelude::unbox(dictBooleanAlgebra)),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                      &Func1::new({
                                                                                                      let monoidDisj
                                                                                                          =
                                                                                                          monoidDisj.clone();
                                                                                                      move
                                                                                                          |p|
                                                                                                          &Func1::new({
                                                                                                                          let p
                                                                                                                              =
                                                                                                                              p.clone();
                                                                                                                          move
                                                                                                                              |q|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                  &&&PureScript_Data_Bifoldable::Data_Bifoldable_unwrap()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMap(),
                                                                                                                                                                                                                                                                                                           &&&dictBifoldable),
                                                                                                                                                                                                                                                                        &&&monoidDisj),
                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_Disj()),
                                                                                                                                                                                                                                                                        &&&p)),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                        &&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_Disj()),
                                                                                                                                                                                                                                     q)))
                                                                                                                      })
                                                                                                  })
                                                                                  }
                                                                          })))
    }
    pub fn Data_Bifoldable_biall() -> &dyn Any {
        static Data_Bifoldable_biall: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifoldable_biall.get_or_init(||
                                              &Func1::new(move
                                                              |dictBifoldable|
                                                              &Func1::new({
                                                                              let dictBifoldable
                                                                                  =
                                                                                  dictBifoldable.clone();
                                                                              move
                                                                                  |dictBooleanAlgebra|
                                                                                  {
                                                                                      let monoidConj =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_monoidConj(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("HeytingAlgebra0"),
                                                                                                                                                                     Sharpurs_Prelude::unbox(dictBooleanAlgebra)),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                      &Func1::new({
                                                                                                      let monoidConj
                                                                                                          =
                                                                                                          monoidConj.clone();
                                                                                                      move
                                                                                                          |p|
                                                                                                          &Func1::new({
                                                                                                                          let p
                                                                                                                              =
                                                                                                                              p.clone();
                                                                                                                          move
                                                                                                                              |q|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                  &&&PureScript_Data_Bifoldable::Data_Bifoldable_unwrap()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMap(),
                                                                                                                                                                                                                                                                                                           &&&dictBifoldable),
                                                                                                                                                                                                                                                                        &&&monoidConj),
                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj()),
                                                                                                                                                                                                                                                                        &&&p)),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                        &&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj()),
                                                                                                                                                                                                                                     q)))
                                                                                                                      })
                                                                                                  })
                                                                                  }
                                                                          })))
    }
}
