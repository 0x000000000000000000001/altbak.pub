pub mod PureScript_Data_Semigroup_Traversable {
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
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_5023b7e9::PureScript_Data_Monoid_Dual;
    use crate::module_7d83a5e5::PureScript_Data_Monoid_Multiplicative;
    use crate::module_2563115d::PureScript_Data_Semigroup_Foldable;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_Semigroup_Traversable_identity() -> &dyn Any {
        static Data_Semigroup_Traversable_identity: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_Traversable_identity.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                             &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Semigroup_Traversable_Traversable1usd_Dict() -> &dyn Any {
        static Data_Semigroup_Traversable_Traversable1usd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Traversable_Traversable1usd_Dict.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |x|
                                                                                        x.clone()))
    }
    pub fn Data_Semigroup_Traversable_traverse1() -> &dyn Any {
        static Data_Semigroup_Traversable_traverse1: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_Traversable_traverse1.get_or_init(||
                                                             &Func1::new(move
                                                                             |dict|
                                                                             find(string("traverse1"),
                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semigroup_Traversable_traversableTuple() -> &dyn Any {
        static Data_Semigroup_Traversable_traversableTuple:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Traversable_traversableTuple.get_or_init(||
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
                                                                                                                                                                         |v|
                                                                                                                                                                         {
                                                                                                                                                                             let matchValue =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                             let matchValue_1:
                                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(v);
                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                    &&&Functor0),
                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                                                                      |usd__arg1|
                                                                                                                                                                                                                                                                                                      Func1::new({
                                                                                                                                                                                                                                                                                                                     let usd__arg1
                                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                                         usd__arg1.clone();
                                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                                         |usd__arg2|
                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                 usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                                                                                                                    &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                       })),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                    }))
                                                                                                                                                                         }
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
                                                                                                                                                         |v_1|
                                                                                                                                                         {
                                                                                                                                                             let matchValue_3:
                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                 Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                    &&&Functor0_1),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                                                      |usd__arg1_1|
                                                                                                                                                                                                                                                                                      Func1::new({
                                                                                                                                                                                                                                                                                                     let usd__arg1_1
                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                         usd__arg1_1.clone();
                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                         |usd__arg2_1|
                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                                                                                                                                 usd__arg2_1.clone()))
                                                                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                                                                                                    &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                       })),
                                                                                                                                                                                              &&&match matchValue_3.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                     =>
                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                 })
                                                                                                                                                         }
                                                                                                                                                 })
                                                                                                                                 }),
                                                                                                                add(string("Foldable10"),
                                                                                                                    &&Func1::new(move
                                                                                                                                     |usd__unused|
                                                                                                                                     &PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldableTuple()),
                                                                                                                    add(string("Traversable1"),
                                                                                                                        &&Func1::new(move
                                                                                                                                         |usd__unused_1|
                                                                                                                                         &PureScript_Data_Traversable::Data_Traversable_traversableTuple()),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))))))
    }
    pub fn Data_Semigroup_Traversable_traversableIdentity() -> &dyn Any {
        static Data_Semigroup_Traversable_traversableIdentity:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Traversable_traversableIdentity.get_or_init(||
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
                                                                                                                                                                            |v|
                                                                                                                                                                            {
                                                                                                                                                                                let matchValue =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                let matchValue_1 =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                       &&&Functor0),
                                                                                                                                                                                                                                                    &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                    &&&matchValue_1))
                                                                                                                                                                            }
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
                                                                                                                                                            |v_1|
                                                                                                                                                            {
                                                                                                                                                                let x_1 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                       &&&Functor0_1),
                                                                                                                                                                                                                                    &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                                                                                 &&&x_1)
                                                                                                                                                            }
                                                                                                                                                    })
                                                                                                                                    }),
                                                                                                                   add(string("Foldable10"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |usd__unused|
                                                                                                                                        &PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldableIdentity()),
                                                                                                                       add(string("Traversable1"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |usd__unused_1|
                                                                                                                                            &PureScript_Data_Traversable::Data_Traversable_traversableIdentity()),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))))))
    }
    pub fn Data_Semigroup_Traversable_sequence1Default() -> &dyn Any {
        static Data_Semigroup_Traversable_sequence1Default:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Traversable_sequence1Default.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictTraversable1|
                                                                                    &Func1::new({
                                                                                                    let dictTraversable1
                                                                                                        =
                                                                                                        dictTraversable1.clone();
                                                                                                    move
                                                                                                        |dictApply|
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_traverse1(),
                                                                                                                                                                                                               &&&dictTraversable1),
                                                                                                                                                                            dictApply),
                                                                                                                                         &&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_identity())
                                                                                                })))
    }
    pub fn Data_Semigroup_Traversable_traversableDual_004019() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_Traversable1usd_Dict(),
                                         &&&add(string("traverse1"),
                                                &&Func1::new(move |dictApply|
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
                                                                                                             |v|
                                                                                                             {
                                                                                                                 let matchValue =
                                                                                                                     Sharpurs_Prelude::unbox(&&f);
                                                                                                                 let matchValue_1 =
                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                        &&&Functor0),
                                                                                                                                                                                     &&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual()),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                     &&&matchValue_1))
                                                                                                             }
                                                                                                     })
                                                                                 })
                                                                 }),
                                                add(string("sequence1"),
                                                    &&Func1::new({
                                                                     let Data_Semigroup_Traversable_traversableDual_004019_002d1
                                                                         =
                                                                         Data_Semigroup_Traversable_traversableDual_004019_002d1.clone();
                                                                     move
                                                                         |dictApply_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_sequence1Default(),
                                                                                                                                             &&&Data_Semigroup_Traversable_traversableDual_004019_002d1.Value),
                                                                                                          dictApply_1)
                                                                 }),
                                                    add(string("Foldable10"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldableDual()),
                                                        add(string("Traversable1"),
                                                            &&Func1::new(move
                                                                             |usd__unused_1|
                                                                             &PureScript_Data_Traversable::Data_Traversable_traversableDual()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_Semigroup_Traversable_traversableDual_004019_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Semigroup_Traversable_traversableDual_004019_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Semigroup_Traversable_traversableDual_004019_002d1.get_or_init(||
                                                                                Lazy(Data_Semigroup_Traversable_traversableDual_004019.clone()))
    }
    pub fn Data_Semigroup_Traversable_traversableDual() -> &dyn Any {
        static Data_Semigroup_Traversable_traversableDual:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Traversable_traversableDual.get_or_init(||
                                                                   Data_Semigroup_Traversable_traversableDual_004019_002d1.Value)
    }
    pub fn Data_Semigroup_Traversable_traversableMultiplicative_004022()
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_Traversable1usd_Dict(),
                                         &&&add(string("traverse1"),
                                                &&Func1::new(move |dictApply|
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
                                                                                                             |v|
                                                                                                             {
                                                                                                                 let matchValue =
                                                                                                                     Sharpurs_Prelude::unbox(&&f);
                                                                                                                 let matchValue_1 =
                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                        &&&Functor0),
                                                                                                                                                                                     &&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_Multiplicative()),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                     &&&matchValue_1))
                                                                                                             }
                                                                                                     })
                                                                                 })
                                                                 }),
                                                add(string("sequence1"),
                                                    &&Func1::new({
                                                                     let Data_Semigroup_Traversable_traversableMultiplicative_004022_002d1
                                                                         =
                                                                         Data_Semigroup_Traversable_traversableMultiplicative_004022_002d1.clone();
                                                                     move
                                                                         |dictApply_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_sequence1Default(),
                                                                                                                                             &&&Data_Semigroup_Traversable_traversableMultiplicative_004022_002d1.Value),
                                                                                                          dictApply_1)
                                                                 }),
                                                    add(string("Foldable10"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldableMultiplicative()),
                                                        add(string("Traversable1"),
                                                            &&Func1::new(move
                                                                             |usd__unused_1|
                                                                             &PureScript_Data_Traversable::Data_Traversable_traversableMultiplicative()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_Semigroup_Traversable_traversableMultiplicative_004022_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static
         Data_Semigroup_Traversable_traversableMultiplicative_004022_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Semigroup_Traversable_traversableMultiplicative_004022_002d1.get_or_init(||
                                                                                          Lazy(Data_Semigroup_Traversable_traversableMultiplicative_004022.clone()))
    }
    pub fn Data_Semigroup_Traversable_traversableMultiplicative()
     -> &dyn Any {
        static Data_Semigroup_Traversable_traversableMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Traversable_traversableMultiplicative.get_or_init(||
                                                                             Data_Semigroup_Traversable_traversableMultiplicative_004022_002d1.Value)
    }
    pub fn Data_Semigroup_Traversable_sequence1() -> &dyn Any {
        static Data_Semigroup_Traversable_sequence1: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_Traversable_sequence1.get_or_init(||
                                                             &Func1::new(move
                                                                             |dict|
                                                                             find(string("sequence1"),
                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semigroup_Traversable_traverse1Default() -> &dyn Any {
        static Data_Semigroup_Traversable_traverse1Default:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Traversable_traverse1Default.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictTraversable1|
                                                                                    {
                                                                                        let Functor0 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable1"),
                                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictTraversable1)),
                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        &Func1::new({
                                                                                                        let Functor0
                                                                                                            =
                                                                                                            Functor0.clone();
                                                                                                        let dictTraversable1
                                                                                                            =
                                                                                                            dictTraversable1.clone();
                                                                                                        move
                                                                                                            |dictApply|
                                                                                                            &Func1::new({
                                                                                                                            let dictApply
                                                                                                                                =
                                                                                                                                dictApply.clone();
                                                                                                                            move
                                                                                                                                |f|
                                                                                                                                &Func1::new({
                                                                                                                                                let f
                                                                                                                                                    =
                                                                                                                                                    f.clone();
                                                                                                                                                move
                                                                                                                                                    |ta|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_sequence1(),
                                                                                                                                                                                                                                                           &&&dictTraversable1),
                                                                                                                                                                                                                        &&&dictApply),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                              &&&Functor0),
                                                                                                                                                                                                                                                           &&&f),
                                                                                                                                                                                                                        ta))
                                                                                                                                            })
                                                                                                                        })
                                                                                                    })
                                                                                    }))
    }
}
