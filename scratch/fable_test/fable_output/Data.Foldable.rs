pub mod PureScript_Data_Foldable {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty as empty_1;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::fix1;
    use fable_library_rust::Option_::getValue;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_97eca9ab::PureScript_Data_Functor_Coproduct;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_bde357b3::PureScript_Data_Maybe_First;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_89cc47bd::PureScript_Data_Monoid_Conj;
    use crate::module_6a9d0e21::PureScript_Data_Monoid_Disj;
    use crate::module_f1079b5::PureScript_Data_Monoid_Endo;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    use fable_library_rust::System::Lazy_1;
    pub mod Data_Foldable_FFI {
        use super::*;
        use fable_library_rust::NativeArray_::count;
        pub fn foldrArray(f: &dyn Any, init: &dyn Any, xs: &dyn Any)
         -> &dyn Any {
            let arr = xs.clone();
            let acc = init.clone();
            for i in (0_i32..=count(arr.clone()) - 1_i32).rev() {
                let step1 =
                    Sharpurs_Prelude::sharpurs_apply(f, &arr[i].clone());
                acc.set(Sharpurs_Prelude::sharpurs_apply(&step1, &acc))
            }
            acc
        }
        pub fn foldlArray(f: &dyn Any, init: &dyn Any, xs: &dyn Any)
         -> &dyn Any {
            let arr = xs.clone();
            let acc = init.clone();
            for i in 0_i32..=count(arr.clone()) - 1_i32 {
                let step1 = Sharpurs_Prelude::sharpurs_apply(f, &acc);
                acc.set(Sharpurs_Prelude::sharpurs_apply(&step1,
                                                         &arr[i].clone()))
            }
            acc
        }
    }
    pub fn Data_Foldable_foldlArray() -> &dyn Any {
        static Data_Foldable_foldlArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldlArray.get_or_init(||
                                                 &Func1::new(move |f|
                                                                 Func1::new({
                                                                                let f
                                                                                    =
                                                                                    f.clone();
                                                                                move
                                                                                    |init|
                                                                                    Func1::new({
                                                                                                   let init
                                                                                                       =
                                                                                                       init.clone();
                                                                                                   move
                                                                                                       |xs|
                                                                                                       PureScript_Data_Foldable::Data_Foldable_FFI::foldlArray(&f,
                                                                                                                                                               &init,
                                                                                                                                                               xs)
                                                                                               })
                                                                            })))
    }
    pub fn Data_Foldable_foldrArray() -> &dyn Any {
        static Data_Foldable_foldrArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldrArray.get_or_init(||
                                                 &Func1::new(move |f|
                                                                 Func1::new({
                                                                                let f
                                                                                    =
                                                                                    f.clone();
                                                                                move
                                                                                    |init|
                                                                                    Func1::new({
                                                                                                   let init
                                                                                                       =
                                                                                                       init.clone();
                                                                                                   move
                                                                                                       |xs|
                                                                                                       PureScript_Data_Foldable::Data_Foldable_FFI::foldrArray(&f,
                                                                                                                                                               &init,
                                                                                                                                                               xs)
                                                                                               })
                                                                            })))
    }
    #[derive(Clone, Debug,)]
    pub enum Data_Foldable_FreeMonoidTree {
        Data_Foldable_Emptyusd_Ctor,
        Data_Foldable_Nodeusd_Ctor(&dyn Any),
        Data_Foldable_Appendusd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Foldable_identity() -> &dyn Any {
        static Data_Foldable_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_identity.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Foldable_unwrap() -> &dyn Any {
        static Data_Foldable_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_unwrap.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                              &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Foldable_monoidEndo() -> &dyn Any {
        static Data_Foldable_monoidEndo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_monoidEndo.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_monoidEndo(),
                                                                                  &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Foldable_identity1() -> &dyn Any {
        static Data_Foldable_identity1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_identity1.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                 &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Foldable_not() -> &dyn Any {
        static Data_Foldable_not: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_not.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                           &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()))
    }
    pub fn Data_Foldable_identity2() -> &dyn Any {
        static Data_Foldable_identity2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_identity2.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                 &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Foldable_Empty() -> &dyn Any {
        static Data_Foldable_Empty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_Empty.get_or_init(||
                                            &LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor))
    }
    pub fn Data_Foldable_Node() -> &dyn Any {
        static Data_Foldable_Node: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_Node.get_or_init(||
                                           &Func1::new(move |usd__arg1|
                                                           &LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Nodeusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Foldable_Append() -> &dyn Any {
        static Data_Foldable_Append: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_Append.get_or_init(||
                                             &Func1::new(move |usd__arg1|
                                                             Func1::new({
                                                                            let usd__arg1
                                                                                =
                                                                                usd__arg1.clone();
                                                                            move
                                                                                |usd__arg2|
                                                                                &LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Appendusd_Ctor(usd__arg1,
                                                                                                                                                                                  usd__arg2.clone()))
                                                                        })))
    }
    pub fn Data_Foldable_Foldableusd_Dict() -> &dyn Any {
        static Data_Foldable_Foldableusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_Foldableusd_Dict.get_or_init(||
                                                       &Func1::new(move |x|
                                                                       x.clone()))
    }
    pub fn Data_Foldable_semigroupFreeMonoidTree() -> &dyn Any {
        static Data_Foldable_semigroupFreeMonoidTree:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_semigroupFreeMonoidTree.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                               &&&add(string("append"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__arg1|
                                                                                                                       Func1::new({
                                                                                                                                      let usd__arg1
                                                                                                                                          =
                                                                                                                                          usd__arg1.clone();
                                                                                                                                      move
                                                                                                                                          |usd__arg2|
                                                                                                                                          &LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Appendusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                            usd__arg2.clone()))
                                                                                                                                  })),
                                                                                                      empty_1::<string,
                                                                                                                &dyn Any>())))
    }
    pub fn Data_Foldable_monoidFreeMonoidTree() -> &dyn Any {
        static Data_Foldable_monoidFreeMonoidTree: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_monoidFreeMonoidTree.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                            &&&add(string("mempty"),
                                                                                                   &&LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor),
                                                                                                   add(string("Semigroup0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_Foldable::Data_Foldable_semigroupFreeMonoidTree()),
                                                                                                       empty_1::<string,
                                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Foldable_foldr() -> &dyn Any {
        static Data_Foldable_foldr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldr.get_or_init(||
                                            &Func1::new(move |dict|
                                                            find(string("foldr"),
                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Foldable_indexr() -> &dyn Any {
        static Data_Foldable_indexr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_indexr.get_or_init(||
                                             &Func1::new(move |dictFoldable|
                                                             &Func1::new({
                                                                             let dictFoldable
                                                                                 =
                                                                                 dictFoldable.clone();
                                                                             move
                                                                                 |idx|
                                                                                 {
                                                                                     let go =
                                                                                         &Func1::new({
                                                                                                         let idx
                                                                                                             =
                                                                                                             idx.clone();
                                                                                                         move
                                                                                                             |a|
                                                                                                             &Func1::new({
                                                                                                                             let a
                                                                                                                                 =
                                                                                                                                 a.clone();
                                                                                                                             move
                                                                                                                                 |cursor|
                                                                                                                                 if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(_)
                                                                                                                                        =
                                                                                                                                        Sharpurs_Prelude::unbox(&find(string("elem"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(cursor))).as_ref()
                                                                                                                                    {
                                                                                                                                     cursor.clone()
                                                                                                                                 } else {
                                                                                                                                     let matchValue_1 =
                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                      &&find(string("pos"),
                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(cursor))),
                                                                                                                                                                                                   &&&idx));
                                                                                                                                     match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                      &matchValue_1)
                                                                                                                                         {
                                                                                                                                         0_i32
                                                                                                                                         =>
                                                                                                                                         &add(string("elem"),
                                                                                                                                              &&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&a)),
                                                                                                                                              add(string("pos"),
                                                                                                                                                  &find(string("pos"),
                                                                                                                                                        Sharpurs_Prelude::unbox(cursor)),
                                                                                                                                                  empty_1::<string,
                                                                                                                                                            &dyn Any>())),
                                                                                                                                         _
                                                                                                                                         =>
                                                                                                                                         &add(string("pos"),
                                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                   &&find(string("pos"),
                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(cursor))),
                                                                                                                                                                                &&&1_i32),
                                                                                                                                              add(string("elem"),
                                                                                                                                                  &find(string("elem"),
                                                                                                                                                        Sharpurs_Prelude::unbox(cursor)),
                                                                                                                                                  empty_1::<string,
                                                                                                                                                            &dyn Any>())),
                                                                                                                                     }
                                                                                                                                 }
                                                                                                                         })
                                                                                                     });
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                           |v|
                                                                                                                                                                           find(string("elem"),
                                                                                                                                                                                Sharpurs_Prelude::unbox(v)))),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                               &&&dictFoldable),
                                                                                                                                                                                            &&&go),
                                                                                                                                                         &&&add(string("elem"),
                                                                                                                                                                &&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                add(string("pos"),
                                                                                                                                                                    &&0_i32,
                                                                                                                                                                    empty_1::<string,
                                                                                                                                                                              &dyn Any>()))))
                                                                                 }
                                                                         })))
    }
    pub fn Data_Foldable_null() -> &dyn Any {
        static Data_Foldable_null: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_null.get_or_init(||
                                           &Func1::new(move |dictFoldable|
                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                  dictFoldable),
                                                                                                                               &&&Func1::new(move
                                                                                                                                                 |v|
                                                                                                                                                 &Func1::new(move
                                                                                                                                                                 |v1|
                                                                                                                                                                 &false))),
                                                                                            &&&true)))
    }
    pub fn Data_Foldable_oneOf() -> &dyn Any {
        static Data_Foldable_oneOf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_oneOf.get_or_init(||
                                            &Func1::new(move |dictFoldable|
                                                            &Func1::new({
                                                                            let dictFoldable
                                                                                =
                                                                                dictFoldable.clone();
                                                                            move
                                                                                |dictPlus|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                       &&&dictFoldable),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictPlus)),
                                                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                                                                    dictPlus))
                                                                        })))
    }
    pub fn Data_Foldable_oneOfMap() -> &dyn Any {
        static Data_Foldable_oneOfMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_oneOfMap.get_or_init(||
                                               &Func1::new(move |dictFoldable|
                                                               &Func1::new({
                                                                               let dictFoldable
                                                                                   =
                                                                                   dictFoldable.clone();
                                                                               move
                                                                                   |dictPlus|
                                                                                   {
                                                                                       let alt =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(dictPlus)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       let empty =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                                            dictPlus);
                                                                                       &Func1::new({
                                                                                                       let alt
                                                                                                           =
                                                                                                           alt.clone();
                                                                                                       let empty
                                                                                                           =
                                                                                                           empty.clone();
                                                                                                       move
                                                                                                           |f|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                  &&&dictFoldable),
                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                     &&&alt),
                                                                                                                                                                                                                  f)),
                                                                                                                                            &&&empty)
                                                                                                   })
                                                                                   }
                                                                           })))
    }
    pub fn Data_Foldable_traverse_() -> &dyn Any {
        static Data_Foldable_traverse_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_traverse_.get_or_init(||
                                                &Func1::new(move
                                                                |dictApplicative|
                                                                {
                                                                    let applySecond =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_applySecond(),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                   Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                    &Func1::new({
                                                                                    let applySecond
                                                                                        =
                                                                                        applySecond.clone();
                                                                                    let dictApplicative
                                                                                        =
                                                                                        dictApplicative.clone();
                                                                                    move
                                                                                        |dictFoldable|
                                                                                        &Func1::new({
                                                                                                        let dictFoldable
                                                                                                            =
                                                                                                            dictFoldable.clone();
                                                                                                        move
                                                                                                            |f|
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                   &&&dictFoldable),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                      &&&applySecond),
                                                                                                                                                                                                                   f)),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                   &&&dictApplicative),
                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                    })
                                                                                })
                                                                }))
    }
    pub fn Data_Foldable_for_() -> &dyn Any {
        static Data_Foldable_for_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_for_.get_or_init(||
                                           &Func1::new(move |dictApplicative|
                                                           {
                                                               let traverse_1 =
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_traverse_(),
                                                                                                    dictApplicative);
                                                               &Func1::new({
                                                                               let traverse_1
                                                                                   =
                                                                                   traverse_1.clone();
                                                                               move
                                                                                   |dictFoldable|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&traverse_1,
                                                                                                                                                       dictFoldable))
                                                                           })
                                                           }))
    }
    pub fn Data_Foldable_sequence_() -> &dyn Any {
        static Data_Foldable_sequence_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_sequence_.get_or_init(||
                                                &Func1::new(move
                                                                |dictApplicative|
                                                                &Func1::new({
                                                                                let dictApplicative
                                                                                    =
                                                                                    dictApplicative.clone();
                                                                                move
                                                                                    |dictFoldable|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_traverse_(),
                                                                                                                                                                                           &&&dictApplicative),
                                                                                                                                                        dictFoldable),
                                                                                                                     &&&PureScript_Data_Foldable::Data_Foldable_identity())
                                                                            })))
    }
    pub fn Data_Foldable_foldl() -> &dyn Any {
        static Data_Foldable_foldl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldl.get_or_init(||
                                            &Func1::new(move |dict|
                                                            find(string("foldl"),
                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Foldable_indexl() -> &dyn Any {
        static Data_Foldable_indexl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_indexl.get_or_init(||
                                             &Func1::new(move |dictFoldable|
                                                             &Func1::new({
                                                                             let dictFoldable
                                                                                 =
                                                                                 dictFoldable.clone();
                                                                             move
                                                                                 |idx|
                                                                                 {
                                                                                     let go =
                                                                                         &Func1::new({
                                                                                                         let idx
                                                                                                             =
                                                                                                             idx.clone();
                                                                                                         move
                                                                                                             |cursor|
                                                                                                             &Func1::new({
                                                                                                                             let cursor
                                                                                                                                 =
                                                                                                                                 cursor.clone();
                                                                                                                             move
                                                                                                                                 |a|
                                                                                                                                 if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(_)
                                                                                                                                        =
                                                                                                                                        Sharpurs_Prelude::unbox(&find(string("elem"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&cursor))).as_ref()
                                                                                                                                    {
                                                                                                                                     &cursor
                                                                                                                                 } else {
                                                                                                                                     let matchValue_1 =
                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                      &&find(string("pos"),
                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&cursor))),
                                                                                                                                                                                                   &&&idx));
                                                                                                                                     match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                      &matchValue_1)
                                                                                                                                         {
                                                                                                                                         0_i32
                                                                                                                                         =>
                                                                                                                                         &add(string("elem"),
                                                                                                                                              &&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(a.clone())),
                                                                                                                                              add(string("pos"),
                                                                                                                                                  &find(string("pos"),
                                                                                                                                                        Sharpurs_Prelude::unbox(&&cursor)),
                                                                                                                                                  empty_1::<string,
                                                                                                                                                            &dyn Any>())),
                                                                                                                                         _
                                                                                                                                         =>
                                                                                                                                         &add(string("pos"),
                                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                   &&find(string("pos"),
                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&cursor))),
                                                                                                                                                                                &&&1_i32),
                                                                                                                                              add(string("elem"),
                                                                                                                                                  &find(string("elem"),
                                                                                                                                                        Sharpurs_Prelude::unbox(&&cursor)),
                                                                                                                                                  empty_1::<string,
                                                                                                                                                            &dyn Any>())),
                                                                                                                                     }
                                                                                                                                 }
                                                                                                                         })
                                                                                                     });
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                           |v|
                                                                                                                                                                           find(string("elem"),
                                                                                                                                                                                Sharpurs_Prelude::unbox(v)))),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                               &&&dictFoldable),
                                                                                                                                                                                            &&&go),
                                                                                                                                                         &&&add(string("elem"),
                                                                                                                                                                &&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                add(string("pos"),
                                                                                                                                                                    &&0_i32,
                                                                                                                                                                    empty_1::<string,
                                                                                                                                                                              &dyn Any>()))))
                                                                                 }
                                                                         })))
    }
    pub fn Data_Foldable_intercalate() -> &dyn Any {
        static Data_Foldable_intercalate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_intercalate.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictFoldable|
                                                                  &Func1::new({
                                                                                  let dictFoldable
                                                                                      =
                                                                                      dictFoldable.clone();
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
                                                                                                              |sep|
                                                                                                              &Func1::new({
                                                                                                                              let sep
                                                                                                                                  =
                                                                                                                                  sep.clone();
                                                                                                                              move
                                                                                                                                  |xs|
                                                                                                                                  {
                                                                                                                                      let go =
                                                                                                                                          &Func1::new(move
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
                                                                                                                                                                                  {
                                                                                                                                                                                      let activePatternResult =
                                                                                                                                                                                          Sharpurs_Prelude::_007cHasProp_007c__007c(string("init"),
                                                                                                                                                                                                                                    &matchValue);
                                                                                                                                                                                      if activePatternResult.is_some()
                                                                                                                                                                                         {
                                                                                                                                                                                          if Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                       &getValue(activePatternResult)).is_some()
                                                                                                                                                                                             {
                                                                                                                                                                                              &add(string("init"),
                                                                                                                                                                                                   &&false,
                                                                                                                                                                                                   add(string("acc"),
                                                                                                                                                                                                       &&matchValue_1,
                                                                                                                                                                                                       empty_1::<string,
                                                                                                                                                                                                                 &dyn Any>()))
                                                                                                                                                                                          } else {
                                                                                                                                                                                              let activePatternResult_2 =
                                                                                                                                                                                                  Sharpurs_Prelude::_007cHasProp_007c__007c(string("acc"),
                                                                                                                                                                                                                                            &matchValue);
                                                                                                                                                                                              if activePatternResult_2.is_some()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  let acc =
                                                                                                                                                                                                      getValue(activePatternResult_2);
                                                                                                                                                                                                  &add(string("init"),
                                                                                                                                                                                                       &&false,
                                                                                                                                                                                                       add(string("acc"),
                                                                                                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                   &&&Semigroup0),
                                                                                                                                                                                                                                                                                &&&acc),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                      &&&Semigroup0),
                                                                                                                                                                                                                                                                                                                   &&&sep),
                                                                                                                                                                                                                                                                                &&&matchValue_1)),
                                                                                                                                                                                                           empty_1::<string,
                                                                                                                                                                                                                     &dyn Any>()))
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  panic!("{}",
                                                                                                                                                                                                         LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Foldable.fs"),
                                  Data1: 80_i32,
                                  Data2: 441_i32,}).get_Message(),)
                                                                                                                                                                                              }
                                                                                                                                                                                          }
                                                                                                                                                                                      } else {
                                                                                                                                                                                          let activePatternResult_2 =
                                                                                                                                                                                              Sharpurs_Prelude::_007cHasProp_007c__007c(string("acc"),
                                                                                                                                                                                                                                        &matchValue);
                                                                                                                                                                                          if activePatternResult_2.is_some()
                                                                                                                                                                                             {
                                                                                                                                                                                              let acc =
                                                                                                                                                                                                  getValue(activePatternResult_2);
                                                                                                                                                                                              &add(string("init"),
                                                                                                                                                                                                   &&false,
                                                                                                                                                                                                   add(string("acc"),
                                                                                                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                               &&&Semigroup0),
                                                                                                                                                                                                                                                                            &&&acc),
                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                  &&&Semigroup0),
                                                                                                                                                                                                                                                                                                               &&&sep),
                                                                                                                                                                                                                                                                            &&&matchValue_1)),
                                                                                                                                                                                                       empty_1::<string,
                                                                                                                                                                                                                 &dyn Any>()))
                                                                                                                                                                                          } else {
                                                                                                                                                                                              panic!("{}",
                                                                                                                                                                                                     LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Foldable.fs"),
                                  Data1: 80_i32,
                                  Data2: 441_i32,}).get_Message(),)
                                                                                                                                                                                          }
                                                                                                                                                                                      }
                                                                                                                                                                                  }
                                                                                                                                                                              }
                                                                                                                                                                      }));
                                                                                                                                      find(string("acc"),
                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                              &&&dictFoldable),
                                                                                                                                                                                                                                                                           &&&go),
                                                                                                                                                                                                                                        &&&add(string("init"),
                                                                                                                                                                                                                                               &&true,
                                                                                                                                                                                                                                               add(string("acc"),
                                                                                                                                                                                                                                                   &&mempty,
                                                                                                                                                                                                                                                   empty_1::<string,
                                                                                                                                                                                                                                                             &dyn Any>()))),
                                                                                                                                                                                                     xs)))
                                                                                                                                  }
                                                                                                                          })
                                                                                                      })
                                                                                      }
                                                                              })))
    }
    pub fn Data_Foldable_length() -> &dyn Any {
        static Data_Foldable_length: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_length.get_or_init(||
                                             &Func1::new(move |dictFoldable|
                                                             &Func1::new({
                                                                             let dictFoldable
                                                                                 =
                                                                                 dictFoldable.clone();
                                                                             move
                                                                                 |dictSemiring|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                        &&&dictFoldable),
                                                                                                                                                     &&&Func1::new({
                                                                                                                                                                       let dictSemiring
                                                                                                                                                                           =
                                                                                                                                                                           dictSemiring.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |c|
                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                           let c
                                                                                                                                                                                               =
                                                                                                                                                                                               c.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |v|
                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                      &&&dictSemiring),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                                                                                                                                                      &&&dictSemiring)),
                                                                                                                                                                                                                                &&&c)
                                                                                                                                                                                       })
                                                                                                                                                                   })),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                     dictSemiring))
                                                                         })))
    }
    pub fn Data_Foldable_maximumBy() -> &dyn Any {
        static Data_Foldable_maximumBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_maximumBy.get_or_init(||
                                                &Func1::new(move
                                                                |dictFoldable|
                                                                &Func1::new({
                                                                                let dictFoldable
                                                                                    =
                                                                                    dictFoldable.clone();
                                                                                move
                                                                                    |cmp|
                                                                                    {
                                                                                        let max_prime =
                                                                                            &Func1::new({
                                                                                                            let cmp
                                                                                                                =
                                                                                                                cmp.clone();
                                                                                                            move
                                                                                                                |v|
                                                                                                                &Func1::new({
                                                                                                                                let v
                                                                                                                                    =
                                                                                                                                    v.clone();
                                                                                                                                move
                                                                                                                                    |v1|
                                                                                                                                    {
                                                                                                                                        let matchValue:
                                                                                                                                                LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                            Sharpurs_Prelude::unbox(&&v);
                                                                                                                                        let matchValue_1 =
                                                                                                                                            Sharpurs_Prelude::unbox(v1);
                                                                                                                                        match matchValue.as_ref()
                                                                                                                                            {
                                                                                                                                            Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                            =>
                                                                                                                                            {
                                                                                                                                                let y =
                                                                                                                                                    matchValue_1;
                                                                                                                                                let x_1 =
                                                                                                                                                    matchValue_1_0.clone();
                                                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor({
                                                                                                                                                                                                           let matchValue_3 =
                                                                                                                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                                                                                                                  &&&x_1),
                                                                                                                                                                                                                                                                                                                                               &&&y)),
                                                                                                                                                                                                                                                                         &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)));
                                                                                                                                                                                                           match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                            &matchValue_3)
                                                                                                                                                                                                               {
                                                                                                                                                                                                               0_i32
                                                                                                                                                                                                               =>
                                                                                                                                                                                                               &x_1,
                                                                                                                                                                                                               _
                                                                                                                                                                                                               =>
                                                                                                                                                                                                               &y,
                                                                                                                                                                                                           }
                                                                                                                                                                                                       }))
                                                                                                                                            }
                                                                                                                                            _
                                                                                                                                            =>
                                                                                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&matchValue_1)),
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                            })
                                                                                                        });
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                               &&&dictFoldable),
                                                                                                                                                            &&&max_prime),
                                                                                                                         &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Foldable_maximum() -> &dyn Any {
        static Data_Foldable_maximum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_maximum.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              {
                                                                  let compare =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                       dictOrd);
                                                                  &Func1::new({
                                                                                  let compare
                                                                                      =
                                                                                      compare.clone();
                                                                                  move
                                                                                      |dictFoldable|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_maximumBy(),
                                                                                                                                                          dictFoldable),
                                                                                                                       &&&compare)
                                                                              })
                                                              }))
    }
    pub fn Data_Foldable_minimumBy() -> &dyn Any {
        static Data_Foldable_minimumBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_minimumBy.get_or_init(||
                                                &Func1::new(move
                                                                |dictFoldable|
                                                                &Func1::new({
                                                                                let dictFoldable
                                                                                    =
                                                                                    dictFoldable.clone();
                                                                                move
                                                                                    |cmp|
                                                                                    {
                                                                                        let min_prime =
                                                                                            &Func1::new({
                                                                                                            let cmp
                                                                                                                =
                                                                                                                cmp.clone();
                                                                                                            move
                                                                                                                |v|
                                                                                                                &Func1::new({
                                                                                                                                let v
                                                                                                                                    =
                                                                                                                                    v.clone();
                                                                                                                                move
                                                                                                                                    |v1|
                                                                                                                                    {
                                                                                                                                        let matchValue:
                                                                                                                                                LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                            Sharpurs_Prelude::unbox(&&v);
                                                                                                                                        let matchValue_1 =
                                                                                                                                            Sharpurs_Prelude::unbox(v1);
                                                                                                                                        match matchValue.as_ref()
                                                                                                                                            {
                                                                                                                                            Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                            =>
                                                                                                                                            {
                                                                                                                                                let y =
                                                                                                                                                    matchValue_1;
                                                                                                                                                let x_1 =
                                                                                                                                                    matchValue_1_0.clone();
                                                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor({
                                                                                                                                                                                                           let matchValue_3 =
                                                                                                                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                                                                                                                  &&&x_1),
                                                                                                                                                                                                                                                                                                                                               &&&y)),
                                                                                                                                                                                                                                                                         &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)));
                                                                                                                                                                                                           match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                            &matchValue_3)
                                                                                                                                                                                                               {
                                                                                                                                                                                                               0_i32
                                                                                                                                                                                                               =>
                                                                                                                                                                                                               &x_1,
                                                                                                                                                                                                               _
                                                                                                                                                                                                               =>
                                                                                                                                                                                                               &y,
                                                                                                                                                                                                           }
                                                                                                                                                                                                       }))
                                                                                                                                            }
                                                                                                                                            _
                                                                                                                                            =>
                                                                                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&matchValue_1)),
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                            })
                                                                                                        });
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                               &&&dictFoldable),
                                                                                                                                                            &&&min_prime),
                                                                                                                         &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Foldable_minimum() -> &dyn Any {
        static Data_Foldable_minimum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_minimum.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              {
                                                                  let compare =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                       dictOrd);
                                                                  &Func1::new({
                                                                                  let compare
                                                                                      =
                                                                                      compare.clone();
                                                                                  move
                                                                                      |dictFoldable|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_minimumBy(),
                                                                                                                                                          dictFoldable),
                                                                                                                       &&&compare)
                                                                              })
                                                              }))
    }
    pub fn Data_Foldable_product() -> &dyn Any {
        static Data_Foldable_product: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_product.get_or_init(||
                                              &Func1::new(move |dictFoldable|
                                                              &Func1::new({
                                                                              let dictFoldable
                                                                                  =
                                                                                  dictFoldable.clone();
                                                                              move
                                                                                  |dictSemiring|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                         &&&dictFoldable),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                         dictSemiring)),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                      dictSemiring))
                                                                          })))
    }
    pub fn Data_Foldable_sum() -> &dyn Any {
        static Data_Foldable_sum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_sum.get_or_init(||
                                          &Func1::new(move |dictFoldable|
                                                          &Func1::new({
                                                                          let dictFoldable
                                                                              =
                                                                              dictFoldable.clone();
                                                                          move
                                                                              |dictSemiring|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                     &&&dictFoldable),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                     dictSemiring)),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                  dictSemiring))
                                                                      })))
    }
    pub fn Data_Foldable_foldableTuple() -> &dyn Any {
        static Data_Foldable_foldableTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableTuple.get_or_init(||
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
                                                                                                                                                     |v|
                                                                                                                                                     {
                                                                                                                                                         let matchValue =
                                                                                                                                                             Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                         let matchValue_1 =
                                                                                                                                                             Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                             &&&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                }),
                                                                                                                                                                                          &&&matchValue_1)
                                                                                                                                                     }
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
                                                                                                                                                         |v_1|
                                                                                                                                                         {
                                                                                                                                                             let matchValue_4 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                             let matchValue_5 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                             let matchValue_6:
                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                 Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                 &&&matchValue_5),
                                                                                                                                                                                              &&&match matchValue_6.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                     =>
                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                 })
                                                                                                                                                         }
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
                                                                                                                                                         |v_2|
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f_2),
                                                                                                                                                                                          &&&match Sharpurs_Prelude::unbox(v_2).as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                 =>
                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                             })
                                                                                                                                                 }))),
                                                                                                    empty_1::<string,
                                                                                                              &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableMultiplicative() -> &dyn Any {
        static Data_Foldable_foldableMultiplicative: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Foldable_foldableMultiplicative.get_or_init(||
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
                                                                                                                                                              |v|
                                                                                                                                                              {
                                                                                                                                                                  let matchValue =
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                  let matchValue_1 =
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                      &&&Sharpurs_Prelude::unbox(v)),
                                                                                                                                                                                                   &&&matchValue_1)
                                                                                                                                                              }
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
                                                                                                                                                                  |v_1|
                                                                                                                                                                  {
                                                                                                                                                                      let matchValue_4 =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                      let matchValue_5 =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                                      let matchValue_6 =
                                                                                                                                                                          Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                          &&&matchValue_5),
                                                                                                                                                                                                       &&&matchValue_6)
                                                                                                                                                                  }
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
                                                                                                                                                                  |v_2|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f_2),
                                                                                                                                                                                                   &&&Sharpurs_Prelude::unbox(v_2))
                                                                                                                                                          }))),
                                                                                                             empty_1::<string,
                                                                                                                       &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableMaybe() -> &dyn Any {
        static Data_Foldable_foldableMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableMaybe.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                     &&&add(string("foldr"),
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
                                                                                                                                                                 LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                             Sharpurs_Prelude::unbox(v2);
                                                                                                                                                         match matchValue_2.as_ref()
                                                                                                                                                             {
                                                                                                                                                             Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_2_1_0)
                                                                                                                                                             =>
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                 &&matchValue_2_1_0),
                                                                                                                                                                                              &&&matchValue_1),
                                                                                                                                                             _
                                                                                                                                                             =>
                                                                                                                                                             &matchValue_1,
                                                                                                                                                         }
                                                                                                                                                     }
                                                                                                                                             })
                                                                                                                         })),
                                                                                            add(string("foldl"),
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
                                                                                                                                                         {
                                                                                                                                                             let matchValue_4 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                             let matchValue_5 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&v1_1);
                                                                                                                                                             let matchValue_6:
                                                                                                                                                                     LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                 Sharpurs_Prelude::unbox(v2_1);
                                                                                                                                                             match matchValue_6.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                 Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_6_1_0)
                                                                                                                                                                 =>
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                     &&&matchValue_5),
                                                                                                                                                                                                  &&matchValue_6_1_0),
                                                                                                                                                                 _
                                                                                                                                                                 =>
                                                                                                                                                                 &matchValue_5,
                                                                                                                                                             }
                                                                                                                                                         }
                                                                                                                                                 })
                                                                                                                             })),
                                                                                                add(string("foldMap"),
                                                                                                    &&Func1::new(move
                                                                                                                     |dictMonoid|
                                                                                                                     {
                                                                                                                         let mempty =
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                              dictMonoid);
                                                                                                                         &Func1::new({
                                                                                                                                         let mempty
                                                                                                                                             =
                                                                                                                                             mempty.clone();
                                                                                                                                         move
                                                                                                                                             |v_2|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let v_2
                                                                                                                                                                 =
                                                                                                                                                                 v_2.clone();
                                                                                                                                                             move
                                                                                                                                                                 |v1_2|
                                                                                                                                                                 {
                                                                                                                                                                     let matchValue_8 =
                                                                                                                                                                         Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                     let matchValue_9:
                                                                                                                                                                             LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                         Sharpurs_Prelude::unbox(v1_2);
                                                                                                                                                                     match matchValue_9.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                         Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_9_1_0)
                                                                                                                                                                         =>
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&matchValue_8,
                                                                                                                                                                                                          &&matchValue_9_1_0),
                                                                                                                                                                         _
                                                                                                                                                                         =>
                                                                                                                                                                         &mempty,
                                                                                                                                                                     }
                                                                                                                                                                 }
                                                                                                                                                         })
                                                                                                                                     })
                                                                                                                     }),
                                                                                                    empty_1::<string,
                                                                                                              &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableIdentity() -> &dyn Any {
        static Data_Foldable_foldableIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableIdentity.get_or_init(||
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
                                                                                                                                                        |v|
                                                                                                                                                        {
                                                                                                                                                            let matchValue =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                            let matchValue_1 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                &&&Sharpurs_Prelude::unbox(v)),
                                                                                                                                                                                             &&&matchValue_1)
                                                                                                                                                        }
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
                                                                                                                                                            |v_1|
                                                                                                                                                            {
                                                                                                                                                                let matchValue_4 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                let matchValue_5 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                                let matchValue_6 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                    &&&matchValue_5),
                                                                                                                                                                                                 &&&matchValue_6)
                                                                                                                                                            }
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
                                                                                                                                                            |v_2|
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f_2),
                                                                                                                                                                                             &&&Sharpurs_Prelude::unbox(v_2))
                                                                                                                                                    }))),
                                                                                                       empty_1::<string,
                                                                                                                 &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableEither() -> &dyn Any {
        static Data_Foldable_foldableEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableEither.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                      &&&add(string("foldr"),
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
                                                                                                                                                                  LrcPtr<Data_Either_Either> =
                                                                                                                                                              Sharpurs_Prelude::unbox(v2);
                                                                                                                                                          match matchValue_2.as_ref()
                                                                                                                                                              {
                                                                                                                                                              Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_2_1_0)
                                                                                                                                                              =>
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                  &&matchValue_2_1_0),
                                                                                                                                                                                               &&&matchValue_1),
                                                                                                                                                              _
                                                                                                                                                              =>
                                                                                                                                                              &matchValue_1,
                                                                                                                                                          }
                                                                                                                                                      }
                                                                                                                                              })
                                                                                                                          })),
                                                                                             add(string("foldl"),
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
                                                                                                                                                          {
                                                                                                                                                              let matchValue_4 =
                                                                                                                                                                  Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                              let matchValue_5 =
                                                                                                                                                                  Sharpurs_Prelude::unbox(&&v1_1);
                                                                                                                                                              let matchValue_6:
                                                                                                                                                                      LrcPtr<Data_Either_Either> =
                                                                                                                                                                  Sharpurs_Prelude::unbox(v2_1);
                                                                                                                                                              match matchValue_6.as_ref()
                                                                                                                                                                  {
                                                                                                                                                                  Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_6_1_0)
                                                                                                                                                                  =>
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                      &&&matchValue_5),
                                                                                                                                                                                                   &&matchValue_6_1_0),
                                                                                                                                                                  _
                                                                                                                                                                  =>
                                                                                                                                                                  &matchValue_5,
                                                                                                                                                              }
                                                                                                                                                          }
                                                                                                                                                  })
                                                                                                                              })),
                                                                                                 add(string("foldMap"),
                                                                                                     &&Func1::new(move
                                                                                                                      |dictMonoid|
                                                                                                                      {
                                                                                                                          let mempty =
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                               dictMonoid);
                                                                                                                          &Func1::new({
                                                                                                                                          let mempty
                                                                                                                                              =
                                                                                                                                              mempty.clone();
                                                                                                                                          move
                                                                                                                                              |v_2|
                                                                                                                                              &Func1::new({
                                                                                                                                                              let v_2
                                                                                                                                                                  =
                                                                                                                                                                  v_2.clone();
                                                                                                                                                              move
                                                                                                                                                                  |v1_2|
                                                                                                                                                                  {
                                                                                                                                                                      let matchValue_8 =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                      let matchValue_9:
                                                                                                                                                                              LrcPtr<Data_Either_Either> =
                                                                                                                                                                          Sharpurs_Prelude::unbox(v1_2);
                                                                                                                                                                      match matchValue_9.as_ref()
                                                                                                                                                                          {
                                                                                                                                                                          Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_9_1_0)
                                                                                                                                                                          =>
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&matchValue_8,
                                                                                                                                                                                                           &&matchValue_9_1_0),
                                                                                                                                                                          _
                                                                                                                                                                          =>
                                                                                                                                                                          &mempty,
                                                                                                                                                                      }
                                                                                                                                                                  }
                                                                                                                                                          })
                                                                                                                                      })
                                                                                                                      }),
                                                                                                     empty_1::<string,
                                                                                                               &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableDual() -> &dyn Any {
        static Data_Foldable_foldableDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableDual.get_or_init(||
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
                                                                                                                                                    |v|
                                                                                                                                                    {
                                                                                                                                                        let matchValue =
                                                                                                                                                            Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                        let matchValue_1 =
                                                                                                                                                            Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                            &&&Sharpurs_Prelude::unbox(v)),
                                                                                                                                                                                         &&&matchValue_1)
                                                                                                                                                    }
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
                                                                                                                                                        |v_1|
                                                                                                                                                        {
                                                                                                                                                            let matchValue_4 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                            let matchValue_5 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                            let matchValue_6 =
                                                                                                                                                                Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                &&&matchValue_5),
                                                                                                                                                                                             &&&matchValue_6)
                                                                                                                                                        }
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
                                                                                                                                                        |v_2|
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f_2),
                                                                                                                                                                                         &&&Sharpurs_Prelude::unbox(v_2))
                                                                                                                                                }))),
                                                                                                   empty_1::<string,
                                                                                                             &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableDisj() -> &dyn Any {
        static Data_Foldable_foldableDisj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableDisj.get_or_init(||
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
                                                                                                                                                    |v|
                                                                                                                                                    {
                                                                                                                                                        let matchValue =
                                                                                                                                                            Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                        let matchValue_1 =
                                                                                                                                                            Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                            &&&Sharpurs_Prelude::unbox(v)),
                                                                                                                                                                                         &&&matchValue_1)
                                                                                                                                                    }
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
                                                                                                                                                        |v_1|
                                                                                                                                                        {
                                                                                                                                                            let matchValue_4 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                            let matchValue_5 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                            let matchValue_6 =
                                                                                                                                                                Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                &&&matchValue_5),
                                                                                                                                                                                             &&&matchValue_6)
                                                                                                                                                        }
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
                                                                                                                                                        |v_2|
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f_2),
                                                                                                                                                                                         &&&Sharpurs_Prelude::unbox(v_2))
                                                                                                                                                }))),
                                                                                                   empty_1::<string,
                                                                                                             &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableConst() -> &dyn Any {
        static Data_Foldable_foldableConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableConst.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                     &&&add(string("foldr"),
                                                                                            &&Func1::new(move
                                                                                                             |v|
                                                                                                             &Func1::new(move
                                                                                                                             |z|
                                                                                                                             &Func1::new({
                                                                                                                                             let z
                                                                                                                                                 =
                                                                                                                                                 z.clone();
                                                                                                                                             move
                                                                                                                                                 |v1|
                                                                                                                                                 &z
                                                                                                                                         }))),
                                                                                            add(string("foldl"),
                                                                                                &&Func1::new(move
                                                                                                                 |v_1|
                                                                                                                 &Func1::new(move
                                                                                                                                 |z_1|
                                                                                                                                 &Func1::new({
                                                                                                                                                 let z_1
                                                                                                                                                     =
                                                                                                                                                     z_1.clone();
                                                                                                                                                 move
                                                                                                                                                     |v1_1|
                                                                                                                                                     &z_1
                                                                                                                                             }))),
                                                                                                add(string("foldMap"),
                                                                                                    &&Func1::new(move
                                                                                                                     |dictMonoid|
                                                                                                                     {
                                                                                                                         let mempty =
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                              dictMonoid);
                                                                                                                         &Func1::new({
                                                                                                                                         let mempty
                                                                                                                                             =
                                                                                                                                             mempty.clone();
                                                                                                                                         move
                                                                                                                                             |v_2|
                                                                                                                                             &Func1::new(move
                                                                                                                                                             |v1_2|
                                                                                                                                                             &mempty)
                                                                                                                                     })
                                                                                                                     }),
                                                                                                    empty_1::<string,
                                                                                                              &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableConj() -> &dyn Any {
        static Data_Foldable_foldableConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableConj.get_or_init(||
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
                                                                                                                                                    |v|
                                                                                                                                                    {
                                                                                                                                                        let matchValue =
                                                                                                                                                            Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                        let matchValue_1 =
                                                                                                                                                            Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                            &&&Sharpurs_Prelude::unbox(v)),
                                                                                                                                                                                         &&&matchValue_1)
                                                                                                                                                    }
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
                                                                                                                                                        |v_1|
                                                                                                                                                        {
                                                                                                                                                            let matchValue_4 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                            let matchValue_5 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                            let matchValue_6 =
                                                                                                                                                                Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                &&&matchValue_5),
                                                                                                                                                                                             &&&matchValue_6)
                                                                                                                                                        }
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
                                                                                                                                                        |v_2|
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f_2),
                                                                                                                                                                                         &&&Sharpurs_Prelude::unbox(v_2))
                                                                                                                                                }))),
                                                                                                   empty_1::<string,
                                                                                                             &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableAdditive() -> &dyn Any {
        static Data_Foldable_foldableAdditive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableAdditive.get_or_init(||
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
                                                                                                                                                        |v|
                                                                                                                                                        {
                                                                                                                                                            let matchValue =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                            let matchValue_1 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                &&&Sharpurs_Prelude::unbox(v)),
                                                                                                                                                                                             &&&matchValue_1)
                                                                                                                                                        }
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
                                                                                                                                                            |v_1|
                                                                                                                                                            {
                                                                                                                                                                let matchValue_4 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                let matchValue_5 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                                let matchValue_6 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                    &&&matchValue_5),
                                                                                                                                                                                                 &&&matchValue_6)
                                                                                                                                                            }
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
                                                                                                                                                            |v_2|
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f_2),
                                                                                                                                                                                             &&&Sharpurs_Prelude::unbox(v_2))
                                                                                                                                                    }))),
                                                                                                       empty_1::<string,
                                                                                                                 &dyn Any>())))))
    }
    pub fn Data_Foldable_foldMapDefaultR() -> &dyn Any {
        static Data_Foldable_foldMapDefaultR: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldMapDefaultR.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictFoldable|
                                                                      &Func1::new({
                                                                                      let dictFoldable
                                                                                          =
                                                                                          dictFoldable.clone();
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
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                         &&&dictFoldable),
                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                        let f
                                                                                                                                                                                                            =
                                                                                                                                                                                                            f.clone();
                                                                                                                                                                                                        move
                                                                                                                                                                                                            |x|
                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                            let x
                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                x.clone();
                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                |acc|
                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                       &&&Semigroup0),
                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                       &&&x)),
                                                                                                                                                                                                                                                                 acc)
                                                                                                                                                                                                                        })
                                                                                                                                                                                                    })),
                                                                                                                                                   &&&mempty)
                                                                                                          })
                                                                                          }
                                                                                  })))
    }
    pub fn Data_Foldable_foldableArray_0040118() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                         &&&add(string("foldr"),
                                                &&PureScript_Data_Foldable::Data_Foldable_foldrArray(),
                                                add(string("foldl"),
                                                    &&PureScript_Data_Foldable::Data_Foldable_foldlArray(),
                                                    add(string("foldMap"),
                                                        &&Func1::new({
                                                                         let Data_Foldable_foldableArray_0040118_002d1
                                                                             =
                                                                             Data_Foldable_foldableArray_0040118_002d1.clone();
                                                                         move
                                                                             |dictMonoid|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMapDefaultR(),
                                                                                                                                                 &&&Data_Foldable_foldableArray_0040118_002d1.Value),
                                                                                                              dictMonoid)
                                                                     }),
                                                        empty_1::<string,
                                                                  &dyn Any>()))))
    }
    pub fn Data_Foldable_foldableArray_0040118_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Foldable_foldableArray_0040118_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Foldable_foldableArray_0040118_002d1.get_or_init(||
                                                                  Lazy(Data_Foldable_foldableArray_0040118.clone()))
    }
    pub fn Data_Foldable_foldableArray() -> &dyn Any {
        static Data_Foldable_foldableArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableArray.get_or_init(||
                                                    Data_Foldable_foldableArray_0040118_002d1.Value)
    }
    pub fn Data_Foldable_foldableFreeMonoidTree_0040121() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                         &&&add(string("foldl"),
                                                &&Func1::new(move |r#fn|
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
                                                                                                                |acc|
                                                                                                                Func1::new({
                                                                                                                               let acc
                                                                                                                                   =
                                                                                                                                   acc.clone();
                                                                                                                               let go_tco
                                                                                                                                   =
                                                                                                                                   go_tco.clone();
                                                                                                                               move
                                                                                                                                   |lhs|
                                                                                                                                   Func1::new({
                                                                                                                                                  let go_tco
                                                                                                                                                      =
                                                                                                                                                      go_tco.clone();
                                                                                                                                                  let lhs
                                                                                                                                                      =
                                                                                                                                                      lhs.clone();
                                                                                                                                                  move
                                                                                                                                                      |rhs|
                                                                                                                                                      go_tco(acc)(lhs)(rhs.clone())
                                                                                                                                              })
                                                                                                                           })
                                                                                                        })
                                                                                    });
                                                                     let go_1 =
                                                                         Lazy(go_2);
                                                                     let go_tco =
                                                                         Func1::new({
                                                                                        let r#fn
                                                                                            =
                                                                                            r#fn.clone();
                                                                                        move
                                                                                            |acc_1|
                                                                                            fix1(&(move
                                                                                                       |go_tco,
                                                                                                        acc_1|
                                                                                                       Func1::new({
                                                                                                                      let acc_1
                                                                                                                          =
                                                                                                                          acc_1.clone();
                                                                                                                      let go_tco
                                                                                                                          =
                                                                                                                          go_tco.clone();
                                                                                                                      move
                                                                                                                          |lhs_1|
                                                                                                                          Func1::new({
                                                                                                                                         let go_tco
                                                                                                                                             =
                                                                                                                                             go_tco.clone();
                                                                                                                                         let lhs_1
                                                                                                                                             =
                                                                                                                                             lhs_1.clone();
                                                                                                                                         move
                                                                                                                                             |rhs_1|
                                                                                                                                             {
                                                                                                                                                 let matchValue:
                                                                                                                                                         LrcPtr<PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree> =
                                                                                                                                                     Sharpurs_Prelude::unbox(&&lhs_1);
                                                                                                                                                 match matchValue.as_ref()
                                                                                                                                                     {
                                                                                                                                                     PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Appendusd_Ctor(matchValue_2_0,
                                                                                                                                                                                                                                          matchValue_2_1)
                                                                                                                                                     =>
                                                                                                                                                     {
                                                                                                                                                         let ys =
                                                                                                                                                             matchValue_2_1.clone();
                                                                                                                                                         let xs =
                                                                                                                                                             matchValue_2_0.clone();
                                                                                                                                                         if let PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&ys).as_ref()
                                                                                                                                                            {
                                                                                                                                                             go_tco(&acc_1)(&xs)(rhs_1.clone())
                                                                                                                                                         } else {
                                                                                                                                                             if let PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor
                                                                                                                                                                    =
                                                                                                                                                                    Sharpurs_Prelude::unbox(rhs_1).as_ref()
                                                                                                                                                                {
                                                                                                                                                                 go_tco(&acc_1)(&xs)(&ys)
                                                                                                                                                             } else {
                                                                                                                                                                 go_tco(&acc_1)(&xs)(&LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Appendusd_Ctor(&ys,
                                                                                                                                                                                                                                                                                       rhs_1.clone())))
                                                                                                                                                             }
                                                                                                                                                         }
                                                                                                                                                     }
                                                                                                                                                     PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor
                                                                                                                                                     =>
                                                                                                                                                     if let PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor
                                                                                                                                                            =
                                                                                                                                                            Sharpurs_Prelude::unbox(rhs_1).as_ref()
                                                                                                                                                        {
                                                                                                                                                         &acc_1
                                                                                                                                                     } else {
                                                                                                                                                         go_tco(&acc_1)(rhs_1.clone())(&LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor))
                                                                                                                                                     },
                                                                                                                                                     PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Nodeusd_Ctor(matchValue_1_0)
                                                                                                                                                     =>
                                                                                                                                                     go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&r#fn,
                                                                                                                                                                                                                                &&&acc_1),
                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Nodeusd_Ctor(x)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                }))(rhs_1.clone())(&LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor)),
                                                                                                                                                 }
                                                                                                                                             }
                                                                                                                                     })
                                                                                                                  })),
                                                                                                 acc_1.clone())
                                                                                    });
                                                                     let go =
                                                                         go_1.Value;
                                                                     &Func1::new({
                                                                                     let go_tco
                                                                                         =
                                                                                         go_tco.clone();
                                                                                     move
                                                                                         |a_1|
                                                                                         &Func1::new({
                                                                                                         let a_1
                                                                                                             =
                                                                                                             a_1.clone();
                                                                                                         let go_tco
                                                                                                             =
                                                                                                             go_tco.clone();
                                                                                                         move
                                                                                                             |b|
                                                                                                             go_tco(&a_1)(b.clone())(&LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor))
                                                                                                     })
                                                                                 })
                                                                 }),
                                                add(string("foldr"),
                                                    &&Func1::new(move |fn_1|
                                                                     {
                                                                         let go_5 =
                                                                             Func0::new({
                                                                                            let go_tco_1
                                                                                                =
                                                                                                go_tco_1.clone();
                                                                                            move
                                                                                                ||
                                                                                                &Func1::new({
                                                                                                                let go_tco_1
                                                                                                                    =
                                                                                                                    go_tco_1.clone();
                                                                                                                move
                                                                                                                    |acc_2|
                                                                                                                    Func1::new({
                                                                                                                                   let acc_2
                                                                                                                                       =
                                                                                                                                       acc_2.clone();
                                                                                                                                   let go_tco_1
                                                                                                                                       =
                                                                                                                                       go_tco_1.clone();
                                                                                                                                   move
                                                                                                                                       |lhs_2|
                                                                                                                                       Func1::new({
                                                                                                                                                      let go_tco_1
                                                                                                                                                          =
                                                                                                                                                          go_tco_1.clone();
                                                                                                                                                      let lhs_2
                                                                                                                                                          =
                                                                                                                                                          lhs_2.clone();
                                                                                                                                                      move
                                                                                                                                                          |rhs_2|
                                                                                                                                                          go_tco_1(acc_2)(lhs_2)(rhs_2.clone())
                                                                                                                                                  })
                                                                                                                               })
                                                                                                            })
                                                                                        });
                                                                         let go_4 =
                                                                             Lazy(go_5);
                                                                         let go_tco_1 =
                                                                             Func1::new({
                                                                                            let fn_1
                                                                                                =
                                                                                                fn_1.clone();
                                                                                            move
                                                                                                |acc_3|
                                                                                                fix1(&(move
                                                                                                           |go_tco_1,
                                                                                                            acc_3|
                                                                                                           Func1::new({
                                                                                                                          let acc_3
                                                                                                                              =
                                                                                                                              acc_3.clone();
                                                                                                                          let go_tco_1
                                                                                                                              =
                                                                                                                              go_tco_1.clone();
                                                                                                                          move
                                                                                                                              |lhs_3|
                                                                                                                              Func1::new({
                                                                                                                                             let go_tco_1
                                                                                                                                                 =
                                                                                                                                                 go_tco_1.clone();
                                                                                                                                             let lhs_3
                                                                                                                                                 =
                                                                                                                                                 lhs_3.clone();
                                                                                                                                             move
                                                                                                                                                 |rhs_3|
                                                                                                                                                 {
                                                                                                                                                     let matchValue_4:
                                                                                                                                                             LrcPtr<PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree> =
                                                                                                                                                         Sharpurs_Prelude::unbox(rhs_3);
                                                                                                                                                     match matchValue_4.as_ref()
                                                                                                                                                         {
                                                                                                                                                         PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Appendusd_Ctor(matchValue_4_2_0,
                                                                                                                                                                                                                                              matchValue_4_2_1)
                                                                                                                                                         =>
                                                                                                                                                         {
                                                                                                                                                             let ys_1 =
                                                                                                                                                                 matchValue_4_2_1.clone();
                                                                                                                                                             let xs_1 =
                                                                                                                                                                 matchValue_4_2_0.clone();
                                                                                                                                                             if let PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor
                                                                                                                                                                    =
                                                                                                                                                                    Sharpurs_Prelude::unbox(&&xs_1).as_ref()
                                                                                                                                                                {
                                                                                                                                                                 go_tco_1(&acc_3)(&lhs_3)(&ys_1)
                                                                                                                                                             } else {
                                                                                                                                                                 if let PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor
                                                                                                                                                                        =
                                                                                                                                                                        Sharpurs_Prelude::unbox(&&lhs_3).as_ref()
                                                                                                                                                                    {
                                                                                                                                                                     go_tco_1(&acc_3)(&xs_1)(&ys_1)
                                                                                                                                                                 } else {
                                                                                                                                                                     go_tco_1(&acc_3)(&LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Appendusd_Ctor(&lhs_3,
                                                                                                                                                                                                                                                                                        &xs_1)))(&ys_1)
                                                                                                                                                                 }
                                                                                                                                                             }
                                                                                                                                                         }
                                                                                                                                                         PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor
                                                                                                                                                         =>
                                                                                                                                                         if let PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&lhs_3).as_ref()
                                                                                                                                                            {
                                                                                                                                                             &acc_3
                                                                                                                                                         } else {
                                                                                                                                                             go_tco_1(&acc_3)(&LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor))(&lhs_3)
                                                                                                                                                         },
                                                                                                                                                         PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Nodeusd_Ctor(matchValue_4_1_0)
                                                                                                                                                         =>
                                                                                                                                                         go_tco_1(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&fn_1,
                                                                                                                                                                                                                                      &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                             PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Nodeusd_Ctor(x)
                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                   &&&acc_3))(&LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor))(&lhs_3),
                                                                                                                                                     }
                                                                                                                                                 }
                                                                                                                                         })
                                                                                                                      })),
                                                                                                     acc_3.clone())
                                                                                        });
                                                                         let go_3 =
                                                                             go_4.Value;
                                                                         &Func1::new({
                                                                                         let go_tco_1
                                                                                             =
                                                                                             go_tco_1.clone();
                                                                                         move
                                                                                             |a_3|
                                                                                             &Func1::new({
                                                                                                             let a_3
                                                                                                                 =
                                                                                                                 a_3.clone();
                                                                                                             let go_tco_1
                                                                                                                 =
                                                                                                                 go_tco_1.clone();
                                                                                                             move
                                                                                                                 |b_1|
                                                                                                                 go_tco_1(&a_3)(&LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Emptyusd_Ctor))(b_1.clone())
                                                                                                         })
                                                                                     })
                                                                     }),
                                                    add(string("foldMap"),
                                                        &&Func1::new({
                                                                         let Data_Foldable_foldableFreeMonoidTree_0040121_002d1
                                                                             =
                                                                             Data_Foldable_foldableFreeMonoidTree_0040121_002d1.clone();
                                                                         move
                                                                             |dictMonoid|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMapDefaultR(),
                                                                                                                                                 &&&Data_Foldable_foldableFreeMonoidTree_0040121_002d1.Value),
                                                                                                              dictMonoid)
                                                                     }),
                                                        empty_1::<string,
                                                                  &dyn Any>()))))
    }
    pub fn Data_Foldable_foldableFreeMonoidTree_0040121_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Foldable_foldableFreeMonoidTree_0040121_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Foldable_foldableFreeMonoidTree_0040121_002d1.get_or_init(||
                                                                           Lazy(Data_Foldable_foldableFreeMonoidTree_0040121.clone()))
    }
    pub fn Data_Foldable_foldableFreeMonoidTree() -> &dyn Any {
        static Data_Foldable_foldableFreeMonoidTree: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Foldable_foldableFreeMonoidTree.get_or_init(||
                                                             Data_Foldable_foldableFreeMonoidTree_0040121_002d1.Value)
    }
    pub fn Data_Foldable_foldMapDefaultL() -> &dyn Any {
        static Data_Foldable_foldMapDefaultL: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldMapDefaultL.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictFoldable|
                                                                      &Func1::new({
                                                                                      let dictFoldable
                                                                                          =
                                                                                          dictFoldable.clone();
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
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                         &&&dictFoldable),
                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                        let f
                                                                                                                                                                                                            =
                                                                                                                                                                                                            f.clone();
                                                                                                                                                                                                        move
                                                                                                                                                                                                            |acc|
                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                            let acc
                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                acc.clone();
                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                |x|
                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                       &&&Semigroup0),
                                                                                                                                                                                                                                                                                                    &&&acc),
                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                    x))
                                                                                                                                                                                                                        })
                                                                                                                                                                                                    })),
                                                                                                                                                   &&&mempty)
                                                                                                          })
                                                                                          }
                                                                                  })))
    }
    pub fn Data_Foldable_foldMap() -> &dyn Any {
        static Data_Foldable_foldMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldMap.get_or_init(||
                                              &Func1::new(move |dict|
                                                              find(string("foldMap"),
                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Foldable_foldableApp() -> &dyn Any {
        static Data_Foldable_foldableApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableApp.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictFoldable|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                                   &&&add(string("foldr"),
                                                                                                          &&Func1::new({
                                                                                                                           let dictFoldable
                                                                                                                               =
                                                                                                                               dictFoldable.clone();
                                                                                                                           move
                                                                                                                               |f|
                                                                                                                               &Func1::new({
                                                                                                                                               let f
                                                                                                                                                   =
                                                                                                                                                   f.clone();
                                                                                                                                               move
                                                                                                                                                   |i|
                                                                                                                                                   &Func1::new({
                                                                                                                                                                   let i
                                                                                                                                                                       =
                                                                                                                                                                       i.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |v|
                                                                                                                                                                       {
                                                                                                                                                                           let matchValue =
                                                                                                                                                                               Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                           let matchValue_1 =
                                                                                                                                                                               Sharpurs_Prelude::unbox(&&i);
                                                                                                                                                                           let matchValue_2 =
                                                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                     &&&dictFoldable),
                                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                                               &&&matchValue_1),
                                                                                                                                                                                                            &&&matchValue_2)
                                                                                                                                                                       }
                                                                                                                                                               })
                                                                                                                                           })
                                                                                                                       }),
                                                                                                          add(string("foldl"),
                                                                                                              &&Func1::new({
                                                                                                                               let dictFoldable
                                                                                                                                   =
                                                                                                                                   dictFoldable.clone();
                                                                                                                               move
                                                                                                                                   |f_1|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let f_1
                                                                                                                                                       =
                                                                                                                                                       f_1.clone();
                                                                                                                                                   move
                                                                                                                                                       |i_1|
                                                                                                                                                       &Func1::new({
                                                                                                                                                                       let i_1
                                                                                                                                                                           =
                                                                                                                                                                           i_1.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |v_1|
                                                                                                                                                                           {
                                                                                                                                                                               let matchValue_4 =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                               let matchValue_5 =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&i_1);
                                                                                                                                                                               let matchValue_6 =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                         &&&dictFoldable),
                                                                                                                                                                                                                                                                                      &&&matchValue_4),
                                                                                                                                                                                                                                                   &&&matchValue_5),
                                                                                                                                                                                                                &&&matchValue_6)
                                                                                                                                                                           }
                                                                                                                                                                   })
                                                                                                                                               })
                                                                                                                           }),
                                                                                                              add(string("foldMap"),
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
                                                                                                                                                           |f_2|
                                                                                                                                                           &Func1::new({
                                                                                                                                                                           let f_2
                                                                                                                                                                               =
                                                                                                                                                                               f_2.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |v_2|
                                                                                                                                                                               {
                                                                                                                                                                                   let matchValue_8 =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                                   let matchValue_9 =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                             &&&dictFoldable),
                                                                                                                                                                                                                                                                                          &&&dictMonoid),
                                                                                                                                                                                                                                                       &&&matchValue_8),
                                                                                                                                                                                                                    &&&matchValue_9)
                                                                                                                                                                               }
                                                                                                                                                                       })
                                                                                                                                                   })
                                                                                                                               }),
                                                                                                                  empty_1::<string,
                                                                                                                            &dyn Any>()))))))
    }
    pub fn Data_Foldable_foldableCompose() -> &dyn Any {
        static Data_Foldable_foldableCompose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableCompose.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictFoldable|
                                                                      &Func1::new({
                                                                                      let dictFoldable
                                                                                          =
                                                                                          dictFoldable.clone();
                                                                                      move
                                                                                          |dictFoldable1|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                                                           &&&add(string("foldr"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let dictFoldable1
                                                                                                                                                       =
                                                                                                                                                       dictFoldable1.clone();
                                                                                                                                                   move
                                                                                                                                                       |f|
                                                                                                                                                       &Func1::new({
                                                                                                                                                                       let f
                                                                                                                                                                           =
                                                                                                                                                                           f.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |i|
                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                           let i
                                                                                                                                                                                               =
                                                                                                                                                                                               i.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |v|
                                                                                                                                                                                               {
                                                                                                                                                                                                   let matchValue =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                                   let matchValue_1 =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&i);
                                                                                                                                                                                                   let matchValue_2 =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                             &&&dictFoldable),
                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                                                                                                   &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                                                                                                &&&matchValue))),
                                                                                                                                                                                                                                                                       &&&matchValue_1),
                                                                                                                                                                                                                                    &&&matchValue_2)
                                                                                                                                                                                               }
                                                                                                                                                                                       })
                                                                                                                                                                   })
                                                                                                                                               }),
                                                                                                                                  add(string("foldl"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let dictFoldable1
                                                                                                                                                           =
                                                                                                                                                           dictFoldable1.clone();
                                                                                                                                                       move
                                                                                                                                                           |f_1|
                                                                                                                                                           &Func1::new({
                                                                                                                                                                           let f_1
                                                                                                                                                                               =
                                                                                                                                                                               f_1.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |i_1|
                                                                                                                                                                               &Func1::new({
                                                                                                                                                                                               let i_1
                                                                                                                                                                                                   =
                                                                                                                                                                                                   i_1.clone();
                                                                                                                                                                                               move
                                                                                                                                                                                                   |v_1|
                                                                                                                                                                                                   {
                                                                                                                                                                                                       let matchValue_4 =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                                                       let matchValue_5 =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&i_1);
                                                                                                                                                                                                       let matchValue_6 =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                                                 &&&dictFoldable),
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                                                                 &&&matchValue_4)),
                                                                                                                                                                                                                                                                           &&&matchValue_5),
                                                                                                                                                                                                                                        &&&matchValue_6)
                                                                                                                                                                                                   }
                                                                                                                                                                                           })
                                                                                                                                                                       })
                                                                                                                                                   }),
                                                                                                                                      add(string("foldMap"),
                                                                                                                                          &&Func1::new({
                                                                                                                                                           let dictFoldable1
                                                                                                                                                               =
                                                                                                                                                               dictFoldable1.clone();
                                                                                                                                                           move
                                                                                                                                                               |dictMonoid|
                                                                                                                                                               &Func1::new({
                                                                                                                                                                               let dictMonoid
                                                                                                                                                                                   =
                                                                                                                                                                                   dictMonoid.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |f_2|
                                                                                                                                                                                   &Func1::new({
                                                                                                                                                                                                   let f_2
                                                                                                                                                                                                       =
                                                                                                                                                                                                       f_2.clone();
                                                                                                                                                                                                   move
                                                                                                                                                                                                       |v_2|
                                                                                                                                                                                                       {
                                                                                                                                                                                                           let matchValue_8 =
                                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                                                           let matchValue_9 =
                                                                                                                                                                                                               Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                                     &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                  &&&dictMonoid),
                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                                                                        &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                                                                     &&&dictMonoid),
                                                                                                                                                                                                                                                                                                                  &&&matchValue_8)),
                                                                                                                                                                                                                                            &&&matchValue_9)
                                                                                                                                                                                                       }
                                                                                                                                                                                               })
                                                                                                                                                                           })
                                                                                                                                                       }),
                                                                                                                                          empty_1::<string,
                                                                                                                                                    &dyn Any>()))))
                                                                                  })))
    }
    pub fn Data_Foldable_foldableCoproduct() -> &dyn Any {
        static Data_Foldable_foldableCoproduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableCoproduct.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictFoldable|
                                                                        &Func1::new({
                                                                                        let dictFoldable
                                                                                            =
                                                                                            dictFoldable.clone();
                                                                                        move
                                                                                            |dictFoldable1|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                                                             &&&add(string("foldr"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let dictFoldable1
                                                                                                                                                         =
                                                                                                                                                         dictFoldable1.clone();
                                                                                                                                                     move
                                                                                                                                                         |f|
                                                                                                                                                         &Func1::new({
                                                                                                                                                                         let f
                                                                                                                                                                             =
                                                                                                                                                                             f.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |z|
                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                                          &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                       &&&f),
                                                                                                                                                                                                                                                                                    z)),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                       &&&dictFoldable1),
                                                                                                                                                                                                                                                                                    &&&f),
                                                                                                                                                                                                                                                 z))
                                                                                                                                                                     })
                                                                                                                                                 }),
                                                                                                                                    add(string("foldl"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let dictFoldable1
                                                                                                                                                             =
                                                                                                                                                             dictFoldable1.clone();
                                                                                                                                                         move
                                                                                                                                                             |f_1|
                                                                                                                                                             &Func1::new({
                                                                                                                                                                             let f_1
                                                                                                                                                                                 =
                                                                                                                                                                                 f_1.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |z_1|
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                                                              &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                           &&&f_1),
                                                                                                                                                                                                                                                                                        z_1)),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                           &&&dictFoldable1),
                                                                                                                                                                                                                                                                                        &&&f_1),
                                                                                                                                                                                                                                                     z_1))
                                                                                                                                                                         })
                                                                                                                                                     }),
                                                                                                                                        add(string("foldMap"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let dictFoldable1
                                                                                                                                                                 =
                                                                                                                                                                 dictFoldable1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |dictMonoid|
                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                 let dictMonoid
                                                                                                                                                                                     =
                                                                                                                                                                                     dictMonoid.clone();
                                                                                                                                                                                 move
                                                                                                                                                                                     |f_2|
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                                                  &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                               &&&dictMonoid),
                                                                                                                                                                                                                                                                                            f_2)),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                               &&&dictFoldable1),
                                                                                                                                                                                                                                                                                            &&&dictMonoid),
                                                                                                                                                                                                                                                         f_2))
                                                                                                                                                                             })
                                                                                                                                                         }),
                                                                                                                                            empty_1::<string,
                                                                                                                                                      &dyn Any>()))))
                                                                                    })))
    }
    pub fn Data_Foldable_foldableFirst() -> &dyn Any {
        static Data_Foldable_foldableFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableFirst.get_or_init(||
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
                                                                                                                                                     |v|
                                                                                                                                                     {
                                                                                                                                                         let matchValue =
                                                                                                                                                             Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                         let matchValue_1 =
                                                                                                                                                             Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                         let matchValue_2 =
                                                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Foldable::Data_Foldable_foldableMaybe()),
                                                                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                                                                             &&&matchValue_1),
                                                                                                                                                                                          &&&matchValue_2)
                                                                                                                                                     }
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
                                                                                                                                                         |v_1|
                                                                                                                                                         {
                                                                                                                                                             let matchValue_4 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                             let matchValue_5 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                             let matchValue_6 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableMaybe()),
                                                                                                                                                                                                                                                                    &&&matchValue_4),
                                                                                                                                                                                                                                 &&&matchValue_5),
                                                                                                                                                                                              &&&matchValue_6)
                                                                                                                                                         }
                                                                                                                                                 })
                                                                                                                             })),
                                                                                                add(string("foldMap"),
                                                                                                    &&Func1::new(move
                                                                                                                     |dictMonoid|
                                                                                                                     &Func1::new({
                                                                                                                                     let dictMonoid
                                                                                                                                         =
                                                                                                                                         dictMonoid.clone();
                                                                                                                                     move
                                                                                                                                         |f_2|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let f_2
                                                                                                                                                             =
                                                                                                                                                             f_2.clone();
                                                                                                                                                         move
                                                                                                                                                             |v_2|
                                                                                                                                                             {
                                                                                                                                                                 let matchValue_8 =
                                                                                                                                                                     Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                 let matchValue_9 =
                                                                                                                                                                     Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Foldable::Data_Foldable_foldableMaybe()),
                                                                                                                                                                                                                                                                        &&&dictMonoid),
                                                                                                                                                                                                                                     &&&matchValue_8),
                                                                                                                                                                                                  &&&matchValue_9)
                                                                                                                                                             }
                                                                                                                                                     })
                                                                                                                                 })),
                                                                                                    empty_1::<string,
                                                                                                              &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableLast() -> &dyn Any {
        static Data_Foldable_foldableLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableLast.get_or_init(||
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
                                                                                                                                                    |v|
                                                                                                                                                    {
                                                                                                                                                        let matchValue =
                                                                                                                                                            Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                        let matchValue_1 =
                                                                                                                                                            Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                        let matchValue_2 =
                                                                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Foldable::Data_Foldable_foldableMaybe()),
                                                                                                                                                                                                                                                               &&&matchValue),
                                                                                                                                                                                                                            &&&matchValue_1),
                                                                                                                                                                                         &&&matchValue_2)
                                                                                                                                                    }
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
                                                                                                                                                        |v_1|
                                                                                                                                                        {
                                                                                                                                                            let matchValue_4 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                            let matchValue_5 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                            let matchValue_6 =
                                                                                                                                                                Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Foldable::Data_Foldable_foldableMaybe()),
                                                                                                                                                                                                                                                                   &&&matchValue_4),
                                                                                                                                                                                                                                &&&matchValue_5),
                                                                                                                                                                                             &&&matchValue_6)
                                                                                                                                                        }
                                                                                                                                                })
                                                                                                                            })),
                                                                                               add(string("foldMap"),
                                                                                                   &&Func1::new(move
                                                                                                                    |dictMonoid|
                                                                                                                    &Func1::new({
                                                                                                                                    let dictMonoid
                                                                                                                                        =
                                                                                                                                        dictMonoid.clone();
                                                                                                                                    move
                                                                                                                                        |f_2|
                                                                                                                                        &Func1::new({
                                                                                                                                                        let f_2
                                                                                                                                                            =
                                                                                                                                                            f_2.clone();
                                                                                                                                                        move
                                                                                                                                                            |v_2|
                                                                                                                                                            {
                                                                                                                                                                let matchValue_8 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                let matchValue_9 =
                                                                                                                                                                    Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Foldable::Data_Foldable_foldableMaybe()),
                                                                                                                                                                                                                                                                       &&&dictMonoid),
                                                                                                                                                                                                                                    &&&matchValue_8),
                                                                                                                                                                                                 &&&matchValue_9)
                                                                                                                                                            }
                                                                                                                                                    })
                                                                                                                                })),
                                                                                                   empty_1::<string,
                                                                                                             &dyn Any>())))))
    }
    pub fn Data_Foldable_foldableProduct() -> &dyn Any {
        static Data_Foldable_foldableProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldableProduct.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictFoldable|
                                                                      &Func1::new({
                                                                                      let dictFoldable
                                                                                          =
                                                                                          dictFoldable.clone();
                                                                                      move
                                                                                          |dictFoldable1|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                                                           &&&add(string("foldr"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let dictFoldable1
                                                                                                                                                       =
                                                                                                                                                       dictFoldable1.clone();
                                                                                                                                                   move
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
                                                                                                                                                                                               |v|
                                                                                                                                                                                               {
                                                                                                                                                                                                   let matchValue =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                                   let matchValue_1 =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&z);
                                                                                                                                                                                                   let matchValue_2:
                                                                                                                                                                                                           LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                   let f1 =
                                                                                                                                                                                                       matchValue;
                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                             &&&dictFoldable),
                                                                                                                                                                                                                                                                                                          &&&f1),
                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                                                                                                   &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                                                                                                &&&f1),
                                                                                                                                                                                                                                                                                                                                             &&&matchValue_1),
                                                                                                                                                                                                                                                                                                          &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                             })),
                                                                                                                                                                                                                                    &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                       })
                                                                                                                                                                                               }
                                                                                                                                                                                       })
                                                                                                                                                                   })
                                                                                                                                               }),
                                                                                                                                  add(string("foldl"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let dictFoldable1
                                                                                                                                                           =
                                                                                                                                                           dictFoldable1.clone();
                                                                                                                                                       move
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
                                                                                                                                                                                                   |v_1|
                                                                                                                                                                                                   {
                                                                                                                                                                                                       let matchValue_4 =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                                                       let matchValue_5 =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&z_1);
                                                                                                                                                                                                       let matchValue_6:
                                                                                                                                                                                                               LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                       let f1_1 =
                                                                                                                                                                                                           matchValue_4;
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                                                 &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                              &&&f1_1),
                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                                                                                    &&&f1_1),
                                                                                                                                                                                                                                                                                                                                                 &&&matchValue_5),
                                                                                                                                                                                                                                                                                                              &&&match matchValue_6.as_ref()
                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                                                                        &&&match matchValue_6.as_ref()
                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                           })
                                                                                                                                                                                                   }
                                                                                                                                                                                           })
                                                                                                                                                                       })
                                                                                                                                                   }),
                                                                                                                                      add(string("foldMap"),
                                                                                                                                          &&Func1::new({
                                                                                                                                                           let dictFoldable1
                                                                                                                                                               =
                                                                                                                                                               dictFoldable1.clone();
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
                                                                                                                                                                                       |f_2|
                                                                                                                                                                                       &Func1::new({
                                                                                                                                                                                                       let f_2
                                                                                                                                                                                                           =
                                                                                                                                                                                                           f_2.clone();
                                                                                                                                                                                                       move
                                                                                                                                                                                                           |v_2|
                                                                                                                                                                                                           {
                                                                                                                                                                                                               let matchValue_8 =
                                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                                                               let matchValue_9:
                                                                                                                                                                                                                       LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                   Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                                               let f1_2 =
                                                                                                                                                                                                                   matchValue_8;
                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                      &&&Semigroup0),
                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                                                                                                               &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                                                                                            &&&dictMonoid),
                                                                                                                                                                                                                                                                                                                                                         &&&f1_2),
                                                                                                                                                                                                                                                                                                                      &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                         })),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                                                                            &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                                                                         &&&dictMonoid),
                                                                                                                                                                                                                                                                                                                      &&&f1_2),
                                                                                                                                                                                                                                                                                   &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                      }))
                                                                                                                                                                                                           }
                                                                                                                                                                                                   })
                                                                                                                                                                               })
                                                                                                                                                               }
                                                                                                                                                       }),
                                                                                                                                          empty_1::<string,
                                                                                                                                                    &dyn Any>()))))
                                                                                  })))
    }
    pub fn Data_Foldable_foldlDefault() -> &dyn Any {
        static Data_Foldable_foldlDefault: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldlDefault.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictFoldable|
                                                                   &Func1::new({
                                                                                   let dictFoldable
                                                                                       =
                                                                                       dictFoldable.clone();
                                                                                   move
                                                                                       |c|
                                                                                       &Func1::new({
                                                                                                       let c
                                                                                                           =
                                                                                                           c.clone();
                                                                                                       move
                                                                                                           |u|
                                                                                                           &Func1::new({
                                                                                                                           let u
                                                                                                                               =
                                                                                                                               u.clone();
                                                                                                                           move
                                                                                                                               |xs|
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Foldable::Data_Foldable_foldableFreeMonoidTree()),
                                                                                                                                                                                                                                                                         &&&c),
                                                                                                                                                                                                                                      &&&u)),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                            &&&dictFoldable),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_monoidFreeMonoidTree()),
                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                        |usd__arg1|
                                                                                                                                                                                                                                                        &LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Nodeusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                   xs))
                                                                                                                       })
                                                                                                   })
                                                                               })))
    }
    pub fn Data_Foldable_foldrDefault() -> &dyn Any {
        static Data_Foldable_foldrDefault: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldrDefault.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictFoldable|
                                                                   &Func1::new({
                                                                                   let dictFoldable
                                                                                       =
                                                                                       dictFoldable.clone();
                                                                                   move
                                                                                       |c|
                                                                                       &Func1::new({
                                                                                                       let c
                                                                                                           =
                                                                                                           c.clone();
                                                                                                       move
                                                                                                           |u|
                                                                                                           &Func1::new({
                                                                                                                           let u
                                                                                                                               =
                                                                                                                               u.clone();
                                                                                                                           move
                                                                                                                               |xs|
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Foldable::Data_Foldable_foldableFreeMonoidTree()),
                                                                                                                                                                                                                                                                         &&&c),
                                                                                                                                                                                                                                      &&&u)),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                            &&&dictFoldable),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_monoidFreeMonoidTree()),
                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                        |usd__arg1|
                                                                                                                                                                                                                                                        &LrcPtr::new(PureScript_Data_Foldable::Data_Foldable_FreeMonoidTree::Data_Foldable_Nodeusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                   xs))
                                                                                                                       })
                                                                                                   })
                                                                               })))
    }
    pub fn Data_Foldable_lookup() -> &dyn Any {
        static Data_Foldable_lookup: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_lookup.get_or_init(||
                                             &Func1::new(move |dictFoldable|
                                                             &Func1::new({
                                                                             let dictFoldable
                                                                                 =
                                                                                 dictFoldable.clone();
                                                                             move
                                                                                 |dictEq|
                                                                                 &Func1::new({
                                                                                                 let dictEq
                                                                                                     =
                                                                                                     dictEq.clone();
                                                                                                 move
                                                                                                     |a|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_unwrap()),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                               &&&dictFoldable),
                                                                                                                                                                                                            &&&PureScript_Data_Maybe_First::Data_Maybe_First_monoidFirst()),
                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                           let a
                                                                                                                                                                                               =
                                                                                                                                                                                               a.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |v|
                                                                                                                                                                                               {
                                                                                                                                                                                                   let matchValue:
                                                                                                                                                                                                           LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe_First::Data_Maybe_First_First(),
                                                                                                                                                                                                                                    &&{
                                                                                                                                                                                                                                          let matchValue_1 =
                                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                                                              &&&dictEq),
                                                                                                                                                                                                                                                                                                                                           &&&a),
                                                                                                                                                                                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                           }));
                                                                                                                                                                                                                                          match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                                           &matchValue_1)
                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                              0_i32
                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                      })),
                                                                                                                                                                                                                                              _
                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                                                                                          }
                                                                                                                                                                                                                                      })
                                                                                                                                                                                               }
                                                                                                                                                                                       })))
                                                                                             })
                                                                         })))
    }
    pub fn Data_Foldable_surroundMap() -> &dyn Any {
        static Data_Foldable_surroundMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_surroundMap.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictFoldable|
                                                                  &Func1::new({
                                                                                  let dictFoldable
                                                                                      =
                                                                                      dictFoldable.clone();
                                                                                  move
                                                                                      |dictSemigroup|
                                                                                      &Func1::new({
                                                                                                      let dictSemigroup
                                                                                                          =
                                                                                                          dictSemigroup.clone();
                                                                                                      move
                                                                                                          |d|
                                                                                                          &Func1::new({
                                                                                                                          let d
                                                                                                                              =
                                                                                                                              d.clone();
                                                                                                                          move
                                                                                                                              |t|
                                                                                                                              &Func1::new({
                                                                                                                                              let t
                                                                                                                                                  =
                                                                                                                                                  t.clone();
                                                                                                                                              move
                                                                                                                                                  |f|
                                                                                                                                                  {
                                                                                                                                                      let joined =
                                                                                                                                                          &Func1::new(move
                                                                                                                                                                          |a|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo(),
                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                             let a
                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                 a.clone();
                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                 |m|
                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                        &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                     &&&d),
                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                           &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&t,
                                                                                                                                                                                                                                                                                                                                                                           &&&a)),
                                                                                                                                                                                                                                                                                                     m))
                                                                                                                                                                                                                         })));
                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                                                      &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Foldable::Data_Foldable_monoidEndo()),
                                                                                                                                                                                                                                                                                                &&&joined),
                                                                                                                                                                                                                                                             f)),
                                                                                                                                                                                       &&&d)
                                                                                                                                                  }
                                                                                                                                          })
                                                                                                                      })
                                                                                                  })
                                                                              })))
    }
    pub fn Data_Foldable_surround() -> &dyn Any {
        static Data_Foldable_surround: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_surround.get_or_init(||
                                               &Func1::new(move |dictFoldable|
                                                               &Func1::new({
                                                                               let dictFoldable
                                                                                   =
                                                                                   dictFoldable.clone();
                                                                               move
                                                                                   |dictSemigroup|
                                                                                   &Func1::new({
                                                                                                   let dictSemigroup
                                                                                                       =
                                                                                                       dictSemigroup.clone();
                                                                                                   move
                                                                                                       |d|
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_surroundMap(),
                                                                                                                                                                                                                                                 &&&dictFoldable),
                                                                                                                                                                                                              &&&dictSemigroup),
                                                                                                                                                                           d),
                                                                                                                                        &&&PureScript_Data_Foldable::Data_Foldable_identity1())
                                                                                               })
                                                                           })))
    }
    pub fn Data_Foldable_foldM() -> &dyn Any {
        static Data_Foldable_foldM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_foldM.get_or_init(||
                                            &Func1::new(move |dictFoldable|
                                                            &Func1::new({
                                                                            let dictFoldable
                                                                                =
                                                                                dictFoldable.clone();
                                                                            move
                                                                                |dictMonad|
                                                                                {
                                                                                    let Bind1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    let Applicative0 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    &Func1::new({
                                                                                                    let Applicative0
                                                                                                        =
                                                                                                        Applicative0.clone();
                                                                                                    let Bind1
                                                                                                        =
                                                                                                        Bind1.clone();
                                                                                                    move
                                                                                                        |f|
                                                                                                        &Func1::new({
                                                                                                                        let f
                                                                                                                            =
                                                                                                                            f.clone();
                                                                                                                        move
                                                                                                                            |b0|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                   &&&dictFoldable),
                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                  |b|
                                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                                  let b
                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                      b.clone();
                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                      |a|
                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                             &&&Bind1),
                                                                                                                                                                                                                                                                                                          &&&b),
                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                             &&&f),
                                                                                                                                                                                                                                                                                                          a))
                                                                                                                                                                                                                              }))),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                   &&&Applicative0),
                                                                                                                                                                                                b0))
                                                                                                                    })
                                                                                                })
                                                                                }
                                                                        })))
    }
    pub fn Data_Foldable_fold() -> &dyn Any {
        static Data_Foldable_fold: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_fold.get_or_init(||
                                           &Func1::new(move |dictFoldable|
                                                           &Func1::new({
                                                                           let dictFoldable
                                                                               =
                                                                               dictFoldable.clone();
                                                                           move
                                                                               |dictMonoid|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                      &&&dictFoldable),
                                                                                                                                                   dictMonoid),
                                                                                                                &&&PureScript_Data_Foldable::Data_Foldable_identity1())
                                                                       })))
    }
    pub fn Data_Foldable_findMap() -> &dyn Any {
        static Data_Foldable_findMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_findMap.get_or_init(||
                                              &Func1::new(move |dictFoldable|
                                                              &Func1::new({
                                                                              let dictFoldable
                                                                                  =
                                                                                  dictFoldable.clone();
                                                                              move
                                                                                  |p|
                                                                                  {
                                                                                      let go =
                                                                                          &Func1::new({
                                                                                                          let p
                                                                                                              =
                                                                                                              p.clone();
                                                                                                          move
                                                                                                              |v|
                                                                                                              &Func1::new({
                                                                                                                              let v
                                                                                                                                  =
                                                                                                                                  v.clone();
                                                                                                                              move
                                                                                                                                  |v1|
                                                                                                                                  {
                                                                                                                                      let matchValue:
                                                                                                                                              LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                      if let Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                             =
                                                                                                                                             matchValue.as_ref()
                                                                                                                                         {
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                           &&&Sharpurs_Prelude::unbox(v1))
                                                                                                                                      } else {
                                                                                                                                          &matchValue
                                                                                                                                      }
                                                                                                                                  }
                                                                                                                          })
                                                                                                      });
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                             &&&dictFoldable),
                                                                                                                                                          &&&go),
                                                                                                                       &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))
                                                                                  }
                                                                          })))
    }
    pub fn Data_Foldable_find() -> &dyn Any {
        static Data_Foldable_find: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_find.get_or_init(||
                                           &Func1::new(move |dictFoldable|
                                                           &Func1::new({
                                                                           let dictFoldable
                                                                               =
                                                                               dictFoldable.clone();
                                                                           move
                                                                               |p|
                                                                               {
                                                                                   let go =
                                                                                       &Func1::new({
                                                                                                       let p
                                                                                                           =
                                                                                                           p.clone();
                                                                                                       move
                                                                                                           |v|
                                                                                                           &Func1::new({
                                                                                                                           let v
                                                                                                                               =
                                                                                                                               v.clone();
                                                                                                                           move
                                                                                                                               |v1|
                                                                                                                               {
                                                                                                                                   let matchValue:
                                                                                                                                           LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                                                                                   let matchValue_1 =
                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                   if let Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                          =
                                                                                                                                          matchValue.as_ref()
                                                                                                                                      {
                                                                                                                                       if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                                    &&&matchValue_1))
                                                                                                                                          {
                                                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&matchValue_1))
                                                                                                                                       } else {
                                                                                                                                           &matchValue
                                                                                                                                       }
                                                                                                                                   } else {
                                                                                                                                       &matchValue
                                                                                                                                   }
                                                                                                                               }
                                                                                                                       })
                                                                                                   });
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                          &&&dictFoldable),
                                                                                                                                                       &&&go),
                                                                                                                    &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))
                                                                               }
                                                                       })))
    }
    pub fn Data_Foldable_any() -> &dyn Any {
        static Data_Foldable_any: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_any.get_or_init(||
                                          &Func1::new(move |dictFoldable|
                                                          {
                                                              let foldMap1 =
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                   dictFoldable);
                                                              &Func1::new({
                                                                              let foldMap1
                                                                                  =
                                                                                  foldMap1.clone();
                                                                              move
                                                                                  |dictHeytingAlgebra|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_alaF(),
                                                                                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                      &&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_Disj()),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&foldMap1,
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_monoidDisj(),
                                                                                                                                                                                         dictHeytingAlgebra)))
                                                                          })
                                                          }))
    }
    pub fn Data_Foldable_elem() -> &dyn Any {
        static Data_Foldable_elem: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_elem.get_or_init(||
                                           &Func1::new(move |dictFoldable|
                                                           {
                                                               let any1 =
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_any(),
                                                                                                                                       dictFoldable),
                                                                                                    &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean());
                                                               &Func1::new({
                                                                               let any1
                                                                                   =
                                                                                   any1.clone();
                                                                               move
                                                                                   |dictEq|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                       &&&any1),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                       dictEq))
                                                                           })
                                                           }))
    }
    pub fn Data_Foldable_notElem() -> &dyn Any {
        static Data_Foldable_notElem: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_notElem.get_or_init(||
                                              &Func1::new(move |dictFoldable|
                                                              &Func1::new({
                                                                              let dictFoldable
                                                                                  =
                                                                                  dictFoldable.clone();
                                                                              move
                                                                                  |dictEq|
                                                                                  &Func1::new({
                                                                                                  let dictEq
                                                                                                      =
                                                                                                      dictEq.clone();
                                                                                                  move
                                                                                                      |x|
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                          &&&PureScript_Data_Foldable::Data_Foldable_not()),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_elem(),
                                                                                                                                                                                                                                                &&&dictFoldable),
                                                                                                                                                                                                             &&&dictEq),
                                                                                                                                                                          x))
                                                                                              })
                                                                          })))
    }
    pub fn Data_Foldable_or() -> &dyn Any {
        static Data_Foldable_or: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_or.get_or_init(||
                                         &Func1::new(move |dictFoldable|
                                                         &Func1::new({
                                                                         let dictFoldable
                                                                             =
                                                                             dictFoldable.clone();
                                                                         move
                                                                             |dictHeytingAlgebra|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_any(),
                                                                                                                                                                                    &&&dictFoldable),
                                                                                                                                                 dictHeytingAlgebra),
                                                                                                              &&&PureScript_Data_Foldable::Data_Foldable_identity2())
                                                                     })))
    }
    pub fn Data_Foldable_all() -> &dyn Any {
        static Data_Foldable_all: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_all.get_or_init(||
                                          &Func1::new(move |dictFoldable|
                                                          {
                                                              let foldMap1 =
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                   dictFoldable);
                                                              &Func1::new({
                                                                              let foldMap1
                                                                                  =
                                                                                  foldMap1.clone();
                                                                              move
                                                                                  |dictHeytingAlgebra|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_alaF(),
                                                                                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                      &&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj()),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&foldMap1,
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_monoidConj(),
                                                                                                                                                                                         dictHeytingAlgebra)))
                                                                          })
                                                          }))
    }
    pub fn Data_Foldable_and() -> &dyn Any {
        static Data_Foldable_and: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Foldable_and.get_or_init(||
                                          &Func1::new(move |dictFoldable|
                                                          &Func1::new({
                                                                          let dictFoldable
                                                                              =
                                                                              dictFoldable.clone();
                                                                          move
                                                                              |dictHeytingAlgebra|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_all(),
                                                                                                                                                                                     &&&dictFoldable),
                                                                                                                                                  dictHeytingAlgebra),
                                                                                                               &&&PureScript_Data_Foldable::Data_Foldable_identity2())
                                                                      })))
    }
}
