pub mod PureScript_Data_Lazy {
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
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_209e0d2c::PureScript_Control_Lazy;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_13840e4f::PureScript_Data_BooleanAlgebra;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_cb047b05::PureScript_Data_CommutativeRing;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_9201da02::PureScript_Data_FoldableWithIndex;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f8b1f47e::PureScript_Data_FunctorWithIndex;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_2563115d::PureScript_Data_Semigroup_Foldable;
    use crate::module_abab3d09::PureScript_Data_Semigroup_Traversable;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_829cacf6::PureScript_Data_TraversableWithIndex;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Data_Lazy_FFI {
        use super::*;
        use fable_library_rust::Native_::Func0;
        use fable_library_rust::Util_::Lazy;
        pub fn defer(thunk: &dyn Any) -> &dyn Any {
            &Lazy(Func0::new({
                                 let thunk = thunk.clone();
                                 move ||
                                     Sharpurs_Prelude::sharpurs_apply(&thunk,
                                                                      &Sharpurs_Prelude::undefined())
                             }))
        }
        pub fn force(l: &dyn Any) -> &dyn Any {
            let lazyObj = Sharpurs_Prelude::unbox(l);
            lazyObj.Value
        }
    }
    pub fn Data_Lazy_defer() -> &dyn Any {
        static Data_Lazy_defer: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_defer.get_or_init(||
                                        &Func1::new(move |thunk|
                                                        PureScript_Data_Lazy::Data_Lazy_FFI::defer(thunk)))
    }
    pub fn Data_Lazy_force() -> &dyn Any {
        static Data_Lazy_force: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_force.get_or_init(||
                                        &Func1::new(move |l|
                                                        PureScript_Data_Lazy::Data_Lazy_FFI::force(l)))
    }
    pub fn Data_Lazy_showLazy() -> &dyn Any {
        static Data_Lazy_showLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_showLazy.get_or_init(||
                                           &Func1::new(move |dictShow|
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                            &&&add(string("show"),
                                                                                                   &&Func1::new({
                                                                                                                    let dictShow
                                                                                                                        =
                                                                                                                        dictShow.clone();
                                                                                                                    move
                                                                                                                        |x|
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                               &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                            &&&string("(defer \\_ -> ")),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                     &&&dictShow),
                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                     x))),
                                                                                                                                                                                            &&&string(")")))
                                                                                                                }),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Lazy_semiringLazy() -> &dyn Any {
        static Data_Lazy_semiringLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_semiringLazy.get_or_init(||
                                               &Func1::new(move |dictSemiring|
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                                &&&add(string("add"),
                                                                                                       &&Func1::new({
                                                                                                                        let dictSemiring
                                                                                                                            =
                                                                                                                            dictSemiring.clone();
                                                                                                                        move
                                                                                                                            |a|
                                                                                                                            &Func1::new({
                                                                                                                                            let a
                                                                                                                                                =
                                                                                                                                                a.clone();
                                                                                                                                            move
                                                                                                                                                |b|
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                                   let b
                                                                                                                                                                                                       =
                                                                                                                                                                                                       b.clone();
                                                                                                                                                                                                   move
                                                                                                                                                                                                       |v|
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                              &&&dictSemiring),
                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                              &&&a)),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                           &&&b))
                                                                                                                                                                                               }))
                                                                                                                                        })
                                                                                                                    }),
                                                                                                       add(string("zero"),
                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                             &&&Func1::new({
                                                                                                                                                               let dictSemiring
                                                                                                                                                                   =
                                                                                                                                                                   dictSemiring.clone();
                                                                                                                                                               move
                                                                                                                                                                   |v_1|
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                                                                    &&&dictSemiring)
                                                                                                                                                           })),
                                                                                                           add(string("mul"),
                                                                                                               &&Func1::new({
                                                                                                                                let dictSemiring
                                                                                                                                    =
                                                                                                                                    dictSemiring.clone();
                                                                                                                                move
                                                                                                                                    |a_1|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let a_1
                                                                                                                                                        =
                                                                                                                                                        a_1.clone();
                                                                                                                                                    move
                                                                                                                                                        |b_1|
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                                           let b_1
                                                                                                                                                                                                               =
                                                                                                                                                                                                               b_1.clone();
                                                                                                                                                                                                           move
                                                                                                                                                                                                               |v_2|
                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                                                                                      &&&dictSemiring),
                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                                      &&&a_1)),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                   &&&b_1))
                                                                                                                                                                                                       }))
                                                                                                                                                })
                                                                                                                            }),
                                                                                                               add(string("one"),
                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                     &&&Func1::new({
                                                                                                                                                                       let dictSemiring
                                                                                                                                                                           =
                                                                                                                                                                           dictSemiring.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |v_3|
                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                                                            &&&dictSemiring)
                                                                                                                                                                   })),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>())))))))
    }
    pub fn Data_Lazy_semigroupLazy() -> &dyn Any {
        static Data_Lazy_semigroupLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_semigroupLazy.get_or_init(||
                                                &Func1::new(move
                                                                |dictSemigroup|
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                 &&&add(string("append"),
                                                                                                        &&Func1::new({
                                                                                                                         let dictSemigroup
                                                                                                                             =
                                                                                                                             dictSemigroup.clone();
                                                                                                                         move
                                                                                                                             |a|
                                                                                                                             &Func1::new({
                                                                                                                                             let a
                                                                                                                                                 =
                                                                                                                                                 a.clone();
                                                                                                                                             move
                                                                                                                                                 |b|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                    let b
                                                                                                                                                                                                        =
                                                                                                                                                                                                        b.clone();
                                                                                                                                                                                                    move
                                                                                                                                                                                                        |v|
                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                               &&&dictSemigroup),
                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                               &&&a)),
                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                            &&&b))
                                                                                                                                                                                                }))
                                                                                                                                         })
                                                                                                                     }),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))))
    }
    pub fn Data_Lazy_ringLazy() -> &dyn Any {
        static Data_Lazy_ringLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_ringLazy.get_or_init(||
                                           &Func1::new(move |dictRing|
                                                           {
                                                               let semiringLazy1 =
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_semiringLazy(),
                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                                              Sharpurs_Prelude::unbox(dictRing)),
                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_Ringusd_Dict(),
                                                                                                &&&add(string("sub"),
                                                                                                       &&Func1::new({
                                                                                                                        let dictRing
                                                                                                                            =
                                                                                                                            dictRing.clone();
                                                                                                                        move
                                                                                                                            |a|
                                                                                                                            &Func1::new({
                                                                                                                                            let a
                                                                                                                                                =
                                                                                                                                                a.clone();
                                                                                                                                            move
                                                                                                                                                |b|
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                                   let b
                                                                                                                                                                                                       =
                                                                                                                                                                                                       b.clone();
                                                                                                                                                                                                   move
                                                                                                                                                                                                       |v|
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                              &&&dictRing),
                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                              &&&a)),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                           &&&b))
                                                                                                                                                                                               }))
                                                                                                                                        })
                                                                                                                    }),
                                                                                                       add(string("Semiring0"),
                                                                                                           &&Func1::new({
                                                                                                                            let semiringLazy1
                                                                                                                                =
                                                                                                                                semiringLazy1.clone();
                                                                                                                            move
                                                                                                                                |usd__unused|
                                                                                                                                &semiringLazy1
                                                                                                                        }),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))
                                                           }))
    }
    pub fn Data_Lazy_monoidLazy() -> &dyn Any {
        static Data_Lazy_monoidLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_monoidLazy.get_or_init(||
                                             &Func1::new(move |dictMonoid|
                                                             {
                                                                 let semigroupLazy1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_semigroupLazy(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                  &&&add(string("mempty"),
                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                           &&&Func1::new({
                                                                                                                                                             let dictMonoid
                                                                                                                                                                 =
                                                                                                                                                                 dictMonoid.clone();
                                                                                                                                                             move
                                                                                                                                                                 |v|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                  &&&dictMonoid)
                                                                                                                                                         })),
                                                                                                         add(string("Semigroup0"),
                                                                                                             &&Func1::new({
                                                                                                                              let semigroupLazy1
                                                                                                                                  =
                                                                                                                                  semigroupLazy1.clone();
                                                                                                                              move
                                                                                                                                  |usd__unused|
                                                                                                                                  &semigroupLazy1
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
                                                             }))
    }
    pub fn Data_Lazy_lazyLazy() -> &dyn Any {
        static Data_Lazy_lazyLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_lazyLazy.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_Lazyusd_Dict(),
                                                                            &&&add(string("defer"),
                                                                                   &&Func1::new(move
                                                                                                    |f|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                     &&&Func1::new({
                                                                                                                                                       let f
                                                                                                                                                           =
                                                                                                                                                           f.clone();
                                                                                                                                                       move
                                                                                                                                                           |v|
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                               &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                   }))),
                                                                                   empty::<string,
                                                                                           &dyn Any>())))
    }
    pub fn Data_Lazy_functorLazy() -> &dyn Any {
        static Data_Lazy_functorLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_functorLazy.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                               &&&add(string("map"),
                                                                                      &&Func1::new(move
                                                                                                       |f|
                                                                                                       &Func1::new({
                                                                                                                       let f
                                                                                                                           =
                                                                                                                           f.clone();
                                                                                                                       move
                                                                                                                           |l|
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                              let l
                                                                                                                                                                                  =
                                                                                                                                                                                  l.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |v|
                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                      &&&l))
                                                                                                                                                                          }))
                                                                                                                   })),
                                                                                      empty::<string,
                                                                                              &dyn Any>())))
    }
    pub fn Data_Lazy_map() -> &dyn Any {
        static Data_Lazy_map: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Lazy_map.get_or_init(||
                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                       &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()))
    }
    pub fn Data_Lazy_functorWithIndexLazy() -> &dyn Any {
        static Data_Lazy_functorWithIndexLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_functorWithIndexLazy.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                        &&&add(string("mapWithIndex"),
                                                                                               &&Func1::new(move
                                                                                                                |f|
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                    &&&PureScript_Data_Lazy::Data_Lazy_map()),
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                    &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                               add(string("Functor0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Lazy_invariantLazy() -> &dyn Any {
        static Data_Lazy_invariantLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_invariantLazy.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                 &&&add(string("imap"),
                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imapF(),
                                                                                                                          &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                        empty::<string,
                                                                                                &dyn Any>())))
    }
    pub fn Data_Lazy_foldableLazy() -> &dyn Any {
        static Data_Lazy_foldableLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_foldableLazy.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                &&&add(string("foldr"),
                                                                                       &&Func1::new(move
                                                                                                        |f|
                                                                                                        &Func1::new({
                                                                                                                        let f
                                                                                                                            =
                                                                                                                            f.clone();
                                                                                                                        move
                                                                                                                            |z|
                                                                                                                            &Func1::new({
                                                                                                                                            let z
                                                                                                                                                =
                                                                                                                                                z.clone();
                                                                                                                                            move
                                                                                                                                                |l|
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                       l)),
                                                                                                                                                                                 &&&z)
                                                                                                                                        })
                                                                                                                    })),
                                                                                       add(string("foldl"),
                                                                                           &&Func1::new(move
                                                                                                            |f_1|
                                                                                                            &Func1::new({
                                                                                                                            let f_1
                                                                                                                                =
                                                                                                                                f_1.clone();
                                                                                                                            move
                                                                                                                                |z_1|
                                                                                                                                &Func1::new({
                                                                                                                                                let z_1
                                                                                                                                                    =
                                                                                                                                                    z_1.clone();
                                                                                                                                                move
                                                                                                                                                    |l_1|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f_1,
                                                                                                                                                                                                                        &&&z_1),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                        l_1))
                                                                                                                                            })
                                                                                                                        })),
                                                                                           add(string("foldMap"),
                                                                                               &&Func1::new(move
                                                                                                                |dictMonoid|
                                                                                                                &Func1::new(move
                                                                                                                                |f_2|
                                                                                                                                &Func1::new({
                                                                                                                                                let f_2
                                                                                                                                                    =
                                                                                                                                                    f_2.clone();
                                                                                                                                                move
                                                                                                                                                    |l_2|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&f_2,
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                        l_2))
                                                                                                                                            }))),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))))
    }
    pub fn Data_Lazy_foldr() -> &dyn Any {
        static Data_Lazy_foldr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_foldr.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                         &&&PureScript_Data_Lazy::Data_Lazy_foldableLazy()))
    }
    pub fn Data_Lazy_foldl() -> &dyn Any {
        static Data_Lazy_foldl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_foldl.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                         &&&PureScript_Data_Lazy::Data_Lazy_foldableLazy()))
    }
    pub fn Data_Lazy_foldMap() -> &dyn Any {
        static Data_Lazy_foldMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_foldMap.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                           &&&PureScript_Data_Lazy::Data_Lazy_foldableLazy()))
    }
    pub fn Data_Lazy_foldableWithIndexLazy() -> &dyn Any {
        static Data_Lazy_foldableWithIndexLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_foldableWithIndexLazy.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                         &&&add(string("foldrWithIndex"),
                                                                                                &&Func1::new(move
                                                                                                                 |f|
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                     &&&PureScript_Data_Lazy::Data_Lazy_foldr()),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                     &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                add(string("foldlWithIndex"),
                                                                                                    &&Func1::new(move
                                                                                                                     |f_1|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                         &&&PureScript_Data_Lazy::Data_Lazy_foldl()),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(f_1,
                                                                                                                                                                                         &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                    add(string("foldMapWithIndex"),
                                                                                                        &&Func1::new(move
                                                                                                                         |dictMonoid|
                                                                                                                         {
                                                                                                                             let foldMap1 =
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_foldMap(),
                                                                                                                                                                  dictMonoid);
                                                                                                                             &Func1::new({
                                                                                                                                             let foldMap1
                                                                                                                                                 =
                                                                                                                                                 foldMap1.clone();
                                                                                                                                             move
                                                                                                                                                 |f_2|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                     &&&foldMap1),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(f_2,
                                                                                                                                                                                                                     &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                         })
                                                                                                                         }),
                                                                                                        add(string("Foldable0"),
                                                                                                            &&Func1::new(move
                                                                                                                             |usd__unused|
                                                                                                                             &PureScript_Data_Lazy::Data_Lazy_foldableLazy()),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>()))))))
    }
    pub fn Data_Lazy_traversableLazy() -> &dyn Any {
        static Data_Lazy_traversableLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_traversableLazy.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                                                                   &&&add(string("traverse"),
                                                                                          &&Func1::new(move
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
                                                                                                                               move
                                                                                                                                   |f|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let f
                                                                                                                                                       =
                                                                                                                                                       f.clone();
                                                                                                                                                   move
                                                                                                                                                       |l|
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                              &&&Functor0),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Lazy::Data_Lazy_defer()),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Function::Data_Function_const())),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                              l)))
                                                                                                                                               })
                                                                                                                           })
                                                                                                           }),
                                                                                          add(string("sequence"),
                                                                                              &&Func1::new(move
                                                                                                               |dictApplicative_1|
                                                                                                               {
                                                                                                                   let Functor0_1 =
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                Sharpurs_Prelude::unbox(dictApplicative_1)),
                                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                   &Func1::new({
                                                                                                                                   let Functor0_1
                                                                                                                                       =
                                                                                                                                       Functor0_1.clone();
                                                                                                                                   move
                                                                                                                                       |l_1|
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                              &&&Functor0_1),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Lazy::Data_Lazy_defer()),
                                                                                                                                                                                                                                              &&&PureScript_Data_Function::Data_Function_const())),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                           l_1))
                                                                                                                               })
                                                                                                               }),
                                                                                              add(string("Functor0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                  add(string("Foldable1"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused_1|
                                                                                                                       &PureScript_Data_Lazy::Data_Lazy_foldableLazy()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>()))))))
    }
    pub fn Data_Lazy_traverse() -> &dyn Any {
        static Data_Lazy_traverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_traverse.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                            &&&PureScript_Data_Lazy::Data_Lazy_traversableLazy()))
    }
    pub fn Data_Lazy_traversableWithIndexLazy() -> &dyn Any {
        static Data_Lazy_traversableWithIndexLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_traversableWithIndexLazy.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                            &&&add(string("traverseWithIndex"),
                                                                                                   &&Func1::new(move
                                                                                                                    |dictApplicative|
                                                                                                                    {
                                                                                                                        let traverse1 =
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_traverse(),
                                                                                                                                                             dictApplicative);
                                                                                                                        &Func1::new({
                                                                                                                                        let traverse1
                                                                                                                                            =
                                                                                                                                            traverse1.clone();
                                                                                                                                        move
                                                                                                                                            |f|
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                &&&traverse1),
                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                    })
                                                                                                                    }),
                                                                                                   add(string("FunctorWithIndex0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_Lazy::Data_Lazy_functorWithIndexLazy()),
                                                                                                       add(string("FoldableWithIndex1"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused_1|
                                                                                                                            &PureScript_Data_Lazy::Data_Lazy_foldableWithIndexLazy()),
                                                                                                           add(string("Traversable2"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused_2|
                                                                                                                                &PureScript_Data_Lazy::Data_Lazy_traversableLazy()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))))
    }
    pub fn Data_Lazy_foldable1Lazy() -> &dyn Any {
        static Data_Lazy_foldable1Lazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_foldable1Lazy.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_Foldable1usd_Dict(),
                                                                                 &&&add(string("foldMap1"),
                                                                                        &&Func1::new(move
                                                                                                         |dictSemigroup|
                                                                                                         &Func1::new(move
                                                                                                                         |f|
                                                                                                                         &Func1::new({
                                                                                                                                         let f
                                                                                                                                             =
                                                                                                                                             f.clone();
                                                                                                                                         move
                                                                                                                                             |l|
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                 l))
                                                                                                                                     }))),
                                                                                        add(string("foldr1"),
                                                                                            &&Func1::new(move
                                                                                                             |v|
                                                                                                             &Func1::new(move
                                                                                                                             |l_1|
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                              l_1))),
                                                                                            add(string("foldl1"),
                                                                                                &&Func1::new(move
                                                                                                                 |v_1|
                                                                                                                 &Func1::new(move
                                                                                                                                 |l_2|
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                  l_2))),
                                                                                                add(string("Foldable0"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &PureScript_Data_Lazy::Data_Lazy_foldableLazy()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))))
    }
    pub fn Data_Lazy_traversable1Lazy() -> &dyn Any {
        static Data_Lazy_traversable1Lazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_traversable1Lazy.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_Traversable1usd_Dict(),
                                                                                    &&&add(string("traverse1"),
                                                                                           &&Func1::new(move
                                                                                                            |dictApply|
                                                                                                            {
                                                                                                                let Functor0 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                &Func1::new({
                                                                                                                                let Functor0
                                                                                                                                    =
                                                                                                                                    Functor0.clone();
                                                                                                                                move
                                                                                                                                    |f|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let f
                                                                                                                                                        =
                                                                                                                                                        f.clone();
                                                                                                                                                    move
                                                                                                                                                        |l|
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                               &&&Functor0),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_defer()),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Function::Data_Function_const())),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                               l)))
                                                                                                                                                })
                                                                                                                            })
                                                                                                            }),
                                                                                           add(string("sequence1"),
                                                                                               &&Func1::new(move
                                                                                                                |dictApply_1|
                                                                                                                {
                                                                                                                    let Functor0_1 =
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictApply_1)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                    &Func1::new({
                                                                                                                                    let Functor0_1
                                                                                                                                        =
                                                                                                                                        Functor0_1.clone();
                                                                                                                                    move
                                                                                                                                        |l_1|
                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                               &&&Functor0_1),
                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_defer()),
                                                                                                                                                                                                                                               &&&PureScript_Data_Function::Data_Function_const())),
                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                            l_1))
                                                                                                                                })
                                                                                                                }),
                                                                                               add(string("Foldable10"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_Lazy::Data_Lazy_foldable1Lazy()),
                                                                                                   add(string("Traversable1"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused_1|
                                                                                                                        &PureScript_Data_Lazy::Data_Lazy_traversableLazy()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))))
    }
    pub fn Data_Lazy_extendLazy() -> &dyn Any {
        static Data_Lazy_extendLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_extendLazy.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                              &&&add(string("extend"),
                                                                                     &&Func1::new(move
                                                                                                      |f|
                                                                                                      &Func1::new({
                                                                                                                      let f
                                                                                                                          =
                                                                                                                          f.clone();
                                                                                                                      move
                                                                                                                          |x|
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                             let x
                                                                                                                                                                                 =
                                                                                                                                                                                 x.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |v|
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                  &&&x)
                                                                                                                                                                         }))
                                                                                                                  })),
                                                                                     add(string("Functor0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                         empty::<string,
                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Lazy_eqLazy() -> &dyn Any {
        static Data_Lazy_eqLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_eqLazy.get_or_init(||
                                         &Func1::new(move |dictEq|
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                          &&&add(string("eq"),
                                                                                                 &&Func1::new({
                                                                                                                  let dictEq
                                                                                                                      =
                                                                                                                      dictEq.clone();
                                                                                                                  move
                                                                                                                      |x|
                                                                                                                      &Func1::new({
                                                                                                                                      let x
                                                                                                                                          =
                                                                                                                                          x.clone();
                                                                                                                                      move
                                                                                                                                          |y|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                 &&&dictEq),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                 &&&x)),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                              y))
                                                                                                                                  })
                                                                                                              }),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_Lazy_ordLazy() -> &dyn Any {
        static Data_Lazy_ordLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_ordLazy.get_or_init(||
                                          &Func1::new(move |dictOrd|
                                                          {
                                                              let eqLazy1 =
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_eqLazy(),
                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                             Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                               &&&add(string("compare"),
                                                                                                      &&Func1::new({
                                                                                                                       let dictOrd
                                                                                                                           =
                                                                                                                           dictOrd.clone();
                                                                                                                       move
                                                                                                                           |x|
                                                                                                                           &Func1::new({
                                                                                                                                           let x
                                                                                                                                               =
                                                                                                                                               x.clone();
                                                                                                                                           move
                                                                                                                                               |y|
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                      &&&dictOrd),
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                      &&&x)),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                   y))
                                                                                                                                       })
                                                                                                                   }),
                                                                                                      add(string("Eq0"),
                                                                                                          &&Func1::new({
                                                                                                                           let eqLazy1
                                                                                                                               =
                                                                                                                               eqLazy1.clone();
                                                                                                                           move
                                                                                                                               |usd__unused|
                                                                                                                               &eqLazy1
                                                                                                                       }),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
                                                          }))
    }
    pub fn Data_Lazy_eq1Lazy() -> &dyn Any {
        static Data_Lazy_eq1Lazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_eq1Lazy.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                           &&&add(string("eq1"),
                                                                                  &&Func1::new(move
                                                                                                   |dictEq|
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_eqLazy(),
                                                                                                                                                                       dictEq))),
                                                                                  empty::<string,
                                                                                          &dyn Any>())))
    }
    pub fn Data_Lazy_ord1Lazy() -> &dyn Any {
        static Data_Lazy_ord1Lazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_ord1Lazy.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                            &&&add(string("compare1"),
                                                                                   &&Func1::new(move
                                                                                                    |dictOrd|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_ordLazy(),
                                                                                                                                                                        dictOrd))),
                                                                                   add(string("Eq10"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &PureScript_Data_Lazy::Data_Lazy_eq1Lazy()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
    pub fn Data_Lazy_comonadLazy() -> &dyn Any {
        static Data_Lazy_comonadLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_comonadLazy.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_Comonadusd_Dict(),
                                                                               &&&add(string("extract"),
                                                                                      &&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                      add(string("Extend0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &PureScript_Data_Lazy::Data_Lazy_extendLazy()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Lazy_commutativeRingLazy() -> &dyn Any {
        static Data_Lazy_commutativeRingLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_commutativeRingLazy.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictCommutativeRing|
                                                                      {
                                                                          let ringLazy1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_ringLazy(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ring0"),
                                                                                                                                                         Sharpurs_Prelude::unbox(dictCommutativeRing)),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingusd_Dict(),
                                                                                                           &&&add(string("Ring0"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let ringLazy1
                                                                                                                                       =
                                                                                                                                       ringLazy1.clone();
                                                                                                                                   move
                                                                                                                                       |usd__unused|
                                                                                                                                       &ringLazy1
                                                                                                                               }),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))
                                                                      }))
    }
    pub fn Data_Lazy_euclideanRingLazy() -> &dyn Any {
        static Data_Lazy_euclideanRingLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_euclideanRingLazy.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictEuclideanRing|
                                                                    {
                                                                        let commutativeRingLazy1 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_commutativeRingLazy(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("CommutativeRing0"),
                                                                                                                                                       Sharpurs_Prelude::unbox(dictEuclideanRing)),
                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_EuclideanRingusd_Dict(),
                                                                                                         &&&add(string("degree"),
                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_degree(),
                                                                                                                                                                                                                        dictEuclideanRing)),
                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_force()),
                                                                                                                add(string("div"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictEuclideanRing
                                                                                                                                         =
                                                                                                                                         dictEuclideanRing.clone();
                                                                                                                                     move
                                                                                                                                         |a|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let a
                                                                                                                                                             =
                                                                                                                                                             a.clone();
                                                                                                                                                         move
                                                                                                                                                             |b|
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                let b
                                                                                                                                                                                                                    =
                                                                                                                                                                                                                    b.clone();
                                                                                                                                                                                                                move
                                                                                                                                                                                                                    |v|
                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                                                                                                                           &&&dictEuclideanRing),
                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                                           &&&a)),
                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                        &&&b))
                                                                                                                                                                                                            }))
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    add(string("mod"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let dictEuclideanRing
                                                                                                                                             =
                                                                                                                                             dictEuclideanRing.clone();
                                                                                                                                         move
                                                                                                                                             |a_1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let a_1
                                                                                                                                                                 =
                                                                                                                                                                 a_1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |b_1|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let b_1
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        b_1.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |v_1|
                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                                                                                                                                                               &&&dictEuclideanRing),
                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                                               &&&a_1)),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                            &&&b_1))
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        add(string("CommutativeRing0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let commutativeRingLazy1
                                                                                                                                                 =
                                                                                                                                                 commutativeRingLazy1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &commutativeRingLazy1
                                                                                                                                         }),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>())))))
                                                                    }))
    }
    pub fn Data_Lazy_boundedLazy() -> &dyn Any {
        static Data_Lazy_boundedLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_boundedLazy.get_or_init(||
                                              &Func1::new(move |dictBounded|
                                                              {
                                                                  let ordLazy1 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_ordLazy(),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                 Sharpurs_Prelude::unbox(dictBounded)),
                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                                   &&&add(string("top"),
                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                            &&&Func1::new({
                                                                                                                                                              let dictBounded
                                                                                                                                                                  =
                                                                                                                                                                  dictBounded.clone();
                                                                                                                                                              move
                                                                                                                                                                  |v|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                   &&&dictBounded)
                                                                                                                                                          })),
                                                                                                          add(string("bottom"),
                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                &&&Func1::new({
                                                                                                                                                                  let dictBounded
                                                                                                                                                                      =
                                                                                                                                                                      dictBounded.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |v_1|
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                       &&&dictBounded)
                                                                                                                                                              })),
                                                                                                              add(string("Ord0"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let ordLazy1
                                                                                                                                       =
                                                                                                                                       ordLazy1.clone();
                                                                                                                                   move
                                                                                                                                       |usd__unused|
                                                                                                                                       &ordLazy1
                                                                                                                               }),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))
                                                              }))
    }
    pub fn Data_Lazy_applyLazy() -> &dyn Any {
        static Data_Lazy_applyLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_applyLazy.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                             &&&add(string("apply"),
                                                                                    &&Func1::new(move
                                                                                                     |f|
                                                                                                     &Func1::new({
                                                                                                                     let f
                                                                                                                         =
                                                                                                                         f.clone();
                                                                                                                     move
                                                                                                                         |x|
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                            let x
                                                                                                                                                                                =
                                                                                                                                                                                x.clone();
                                                                                                                                                                            move
                                                                                                                                                                                |v|
                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                    &&&f),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                    &&&x))
                                                                                                                                                                        }))
                                                                                                                 })),
                                                                                    add(string("Functor0"),
                                                                                        &&Func1::new(move
                                                                                                         |usd__unused|
                                                                                                         &PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                        empty::<string,
                                                                                                &dyn Any>()))))
    }
    pub fn Data_Lazy_bindLazy() -> &dyn Any {
        static Data_Lazy_bindLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_bindLazy.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                            &&&add(string("bind"),
                                                                                   &&Func1::new(move
                                                                                                    |l|
                                                                                                    &Func1::new({
                                                                                                                    let l
                                                                                                                        =
                                                                                                                        l.clone();
                                                                                                                    move
                                                                                                                        |f|
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                           let f
                                                                                                                                                                               =
                                                                                                                                                                               f.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |v|
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Lazy::Data_Lazy_force()),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                      &&&l)))
                                                                                                                                                                       }))
                                                                                                                })),
                                                                                   add(string("Apply0"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &PureScript_Data_Lazy::Data_Lazy_applyLazy()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
    pub fn Data_Lazy_heytingAlgebraLazy() -> &dyn Any {
        static Data_Lazy_heytingAlgebraLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_heytingAlgebraLazy.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictHeytingAlgebra|
                                                                     {
                                                                         let implies =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_implies(),
                                                                                                              dictHeytingAlgebra);
                                                                         let conj =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                              dictHeytingAlgebra);
                                                                         let disj =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                              dictHeytingAlgebra);
                                                                         let not_var =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                              dictHeytingAlgebra);
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_HeytingAlgebrausd_Dict(),
                                                                                                          &&&add(string("ff"),
                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                     let dictHeytingAlgebra
                                                                                                                                                                         =
                                                                                                                                                                         dictHeytingAlgebra.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |v|
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ff(),
                                                                                                                                                                                                          &&&dictHeytingAlgebra)
                                                                                                                                                                 })),
                                                                                                                 add(string("tt"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let dictHeytingAlgebra
                                                                                                                                                                             =
                                                                                                                                                                             dictHeytingAlgebra.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v_1|
                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_tt(),
                                                                                                                                                                                                              &&&dictHeytingAlgebra)
                                                                                                                                                                     })),
                                                                                                                     add(string("implies"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let implies
                                                                                                                                              =
                                                                                                                                              implies.clone();
                                                                                                                                          move
                                                                                                                                              |a|
                                                                                                                                              &Func1::new({
                                                                                                                                                              let a
                                                                                                                                                                  =
                                                                                                                                                                  a.clone();
                                                                                                                                                              move
                                                                                                                                                                  |b|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Lazy::Data_Lazy_applyLazy()),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                                            &&&implies),
                                                                                                                                                                                                                                                                         &&&a)),
                                                                                                                                                                                                   b)
                                                                                                                                                          })
                                                                                                                                      }),
                                                                                                                         add(string("conj"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let conj
                                                                                                                                                  =
                                                                                                                                                  conj.clone();
                                                                                                                                              move
                                                                                                                                                  |a_1|
                                                                                                                                                  &Func1::new({
                                                                                                                                                                  let a_1
                                                                                                                                                                      =
                                                                                                                                                                      a_1.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |b_1|
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_Lazy::Data_Lazy_applyLazy()),
                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                                                &&&conj),
                                                                                                                                                                                                                                                                             &&&a_1)),
                                                                                                                                                                                                       b_1)
                                                                                                                                                              })
                                                                                                                                          }),
                                                                                                                             add(string("disj"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let disj
                                                                                                                                                      =
                                                                                                                                                      disj.clone();
                                                                                                                                                  move
                                                                                                                                                      |a_2|
                                                                                                                                                      &Func1::new({
                                                                                                                                                                      let a_2
                                                                                                                                                                          =
                                                                                                                                                                          a_2.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |b_2|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Lazy::Data_Lazy_applyLazy()),
                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                                                                                    &&&disj),
                                                                                                                                                                                                                                                                                 &&&a_2)),
                                                                                                                                                                                                           b_2)
                                                                                                                                                                  })
                                                                                                                                              }),
                                                                                                                                 add(string("not"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let not_var
                                                                                                                                                          =
                                                                                                                                                          not_var.clone();
                                                                                                                                                      move
                                                                                                                                                          |a_3|
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                 &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                              &&&not_var),
                                                                                                                                                                                           a_3)
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))))))
                                                                     }))
    }
    pub fn Data_Lazy_booleanAlgebraLazy() -> &dyn Any {
        static Data_Lazy_booleanAlgebraLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_booleanAlgebraLazy.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBooleanAlgebra|
                                                                     {
                                                                         let heytingAlgebraLazy1 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_heytingAlgebraLazy(),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("HeytingAlgebra0"),
                                                                                                                                                        Sharpurs_Prelude::unbox(dictBooleanAlgebra)),
                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_BooleanAlgebrausd_Dict(),
                                                                                                          &&&add(string("HeytingAlgebra0"),
                                                                                                                 &&Func1::new({
                                                                                                                                  let heytingAlgebraLazy1
                                                                                                                                      =
                                                                                                                                      heytingAlgebraLazy1.clone();
                                                                                                                                  move
                                                                                                                                      |usd__unused|
                                                                                                                                      &heytingAlgebraLazy1
                                                                                                                              }),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>()))
                                                                     }))
    }
    pub fn Data_Lazy_applicativeLazy() -> &dyn Any {
        static Data_Lazy_applicativeLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_applicativeLazy.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                   &&&add(string("pure"),
                                                                                          &&Func1::new(move
                                                                                                           |a|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                            &&&Func1::new({
                                                                                                                                                              let a
                                                                                                                                                                  =
                                                                                                                                                                  a.clone();
                                                                                                                                                              move
                                                                                                                                                                  |v|
                                                                                                                                                                  &a
                                                                                                                                                          }))),
                                                                                          add(string("Apply0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Data_Lazy::Data_Lazy_applyLazy()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Lazy_monadLazy() -> &dyn Any {
        static Data_Lazy_monadLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Lazy_monadLazy.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                             &&&add(string("Applicative0"),
                                                                                    &&Func1::new(move
                                                                                                     |usd__unused|
                                                                                                     &PureScript_Data_Lazy::Data_Lazy_applicativeLazy()),
                                                                                    add(string("Bind1"),
                                                                                        &&Func1::new(move
                                                                                                         |usd__unused_1|
                                                                                                         &PureScript_Data_Lazy::Data_Lazy_bindLazy()),
                                                                                        empty::<string,
                                                                                                &dyn Any>()))))
    }
}
