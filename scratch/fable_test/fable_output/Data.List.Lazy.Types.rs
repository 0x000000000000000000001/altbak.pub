pub mod PureScript_Data_List_Lazy_Types {
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
    use fable_library_rust::Option_::getValue;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_209e0d2c::PureScript_Control_Lazy;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_53a2e11::PureScript_Control_MonadPlus;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_9201da02::PureScript_Data_FoldableWithIndex;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f8b1f47e::PureScript_Data_FunctorWithIndex;
    use crate::module_720e12db::PureScript_Data_Lazy;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty::Data_NonEmpty_NonEmpty;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_829cacf6::PureScript_Data_TraversableWithIndex;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_b1754f14::PureScript_Data_Unfoldable1;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    use fable_library_rust::System::Lazy_1;
    #[derive(Clone, Debug,)]
    pub enum Data_List_Lazy_Types_Step {
        Data_List_Lazy_Types_Nilusd_Ctor,
        Data_List_Lazy_Types_Consusd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_List_Lazy_Types_identity() -> &dyn Any {
        static Data_List_Lazy_Types_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_identity.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                       &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_List_Lazy_Types_unwrap() -> &dyn Any {
        static Data_List_Lazy_Types_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_unwrap.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_List_Lazy_Types_List() -> &dyn Any {
        static Data_List_Lazy_Types_List: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_List.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  x.clone()))
    }
    pub fn Data_List_Lazy_Types_Nil() -> &dyn Any {
        static Data_List_Lazy_Types_Nil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_Nil.get_or_init(||
                                                 &LrcPtr::new(PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor))
    }
    pub fn Data_List_Lazy_Types_Cons() -> &dyn Any {
        static Data_List_Lazy_Types_Cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_Cons.get_or_init(||
                                                  &Func1::new(move |usd__arg1|
                                                                  Func1::new({
                                                                                 let usd__arg1
                                                                                     =
                                                                                     usd__arg1.clone();
                                                                                 move
                                                                                     |usd__arg2|
                                                                                     &LrcPtr::new(PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(usd__arg1,
                                                                                                                                                                                                usd__arg2.clone()))
                                                                             })))
    }
    pub fn Data_List_Lazy_Types_NonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_NonEmptyList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_NonEmptyList.get_or_init(||
                                                          &Func1::new(move |x|
                                                                          x.clone()))
    }
    pub fn Data_List_Lazy_Types_nil() -> &dyn Any {
        static Data_List_Lazy_Types_nil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_nil.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                     &&&Func1::new(move
                                                                                                                                       |v|
                                                                                                                                       &LrcPtr::new(PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)))))
    }
    pub fn Data_List_Lazy_Types_newtypeNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_newtypeNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_newtypeNonEmptyList.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                  &&&add(string("Coercible0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &Sharpurs_Prelude::Prim_undefined()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_newtypeList() -> &dyn Any {
        static Data_List_Lazy_Types_newtypeList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_newtypeList.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                          &&&add(string("Coercible0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &Sharpurs_Prelude::Prim_undefined()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_step() -> &dyn Any {
        static Data_List_Lazy_Types_step: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_step.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                      &&&PureScript_Data_Lazy::Data_Lazy_force()),
                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined())))
    }
    pub fn Data_List_Lazy_Types_semigroupList_004031() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                         &&&add(string("append"),
                                                &&Func1::new({
                                                                 let Data_List_Lazy_Types_semigroupList_004031_002d1
                                                                     =
                                                                     Data_List_Lazy_Types_semigroupList_004031_002d1.clone();
                                                                 move |xs|
                                                                     &Func1::new({
                                                                                     let Data_List_Lazy_Types_semigroupList_004031_002d1
                                                                                         =
                                                                                         Data_List_Lazy_Types_semigroupList_004031_002d1.clone();
                                                                                     let xs
                                                                                         =
                                                                                         xs.clone();
                                                                                     move
                                                                                         |ys|
                                                                                         {
                                                                                             let go =
                                                                                                 &Func1::new({
                                                                                                                 let Data_List_Lazy_Types_semigroupList_004031_002d1
                                                                                                                     =
                                                                                                                     Data_List_Lazy_Types_semigroupList_004031_002d1.clone();
                                                                                                                 let ys
                                                                                                                     =
                                                                                                                     ys.clone();
                                                                                                                 move
                                                                                                                     |v|
                                                                                                                     {
                                                                                                                         let matchValue:
                                                                                                                                 LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                         match matchValue.as_ref()
                                                                                                                             {
                                                                                                                             PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                           matchValue_1_1)
                                                                                                                             =>
                                                                                                                             &LrcPtr::new(PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                               &&&Data_List_Lazy_Types_semigroupList_004031_002d1.Value),
                                                                                                                                                                                                                                                                                                            &&matchValue_1_1),
                                                                                                                                                                                                                                                                         &&&ys))),
                                                                                                                             _
                                                                                                                             =>
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                              &&&ys),
                                                                                                                         }
                                                                                                                     }
                                                                                                             });
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                       &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                    &&&go),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                    &&&xs)))
                                                                                         }
                                                                                 })
                                                             }),
                                                empty::<string, &dyn Any>()))
    }
    pub fn Data_List_Lazy_Types_semigroupList_004031_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_Types_semigroupList_004031_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_Types_semigroupList_004031_002d1.get_or_init(||
                                                                        Lazy(Data_List_Lazy_Types_semigroupList_004031.clone()))
    }
    pub fn Data_List_Lazy_Types_semigroupList() -> &dyn Any {
        static Data_List_Lazy_Types_semigroupList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_semigroupList.get_or_init(||
                                                           Data_List_Lazy_Types_semigroupList_004031_002d1.Value)
    }
    pub fn Data_List_Lazy_Types_monoidList() -> &dyn Any {
        static Data_List_Lazy_Types_monoidList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_monoidList.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                         &&&add(string("mempty"),
                                                                                                &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                add(string("Semigroup0"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_lazyList() -> &dyn Any {
        static Data_List_Lazy_Types_lazyList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_lazyList.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_Lazyusd_Dict(),
                                                                                       &&&add(string("defer"),
                                                                                              &&Func1::new(move
                                                                                                               |f|
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                   &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                         &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step()),
                                                                                                                                                                                                                      f)))),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_functorList_004038() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                         &&&add(string("map"),
                                                &&Func1::new({
                                                                 let Data_List_Lazy_Types_functorList_004038_002d1
                                                                     =
                                                                     Data_List_Lazy_Types_functorList_004038_002d1.clone();
                                                                 move |f|
                                                                     &Func1::new({
                                                                                     let Data_List_Lazy_Types_functorList_004038_002d1
                                                                                         =
                                                                                         Data_List_Lazy_Types_functorList_004038_002d1.clone();
                                                                                     let f
                                                                                         =
                                                                                         f.clone();
                                                                                     move
                                                                                         |xs|
                                                                                         {
                                                                                             let go =
                                                                                                 &Func1::new({
                                                                                                                 let Data_List_Lazy_Types_functorList_004038_002d1
                                                                                                                     =
                                                                                                                     Data_List_Lazy_Types_functorList_004038_002d1.clone();
                                                                                                                 move
                                                                                                                     |v|
                                                                                                                     {
                                                                                                                         let matchValue:
                                                                                                                                 LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                         match matchValue.as_ref()
                                                                                                                             {
                                                                                                                             PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                           matchValue_1_1)
                                                                                                                             =>
                                                                                                                             &LrcPtr::new(PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                         &&matchValue_1_0),
                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                               &&&Data_List_Lazy_Types_functorList_004038_002d1.Value),
                                                                                                                                                                                                                                                                                                            &&&f),
                                                                                                                                                                                                                                                                         &&matchValue_1_1))),
                                                                                                                             _
                                                                                                                             =>
                                                                                                                             &LrcPtr::new(PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor),
                                                                                                                         }
                                                                                                                     }
                                                                                                             });
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                       &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                    &&&go),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                    xs)))
                                                                                         }
                                                                                 })
                                                             }),
                                                empty::<string, &dyn Any>()))
    }
    pub fn Data_List_Lazy_Types_functorList_004038_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_Types_functorList_004038_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_Types_functorList_004038_002d1.get_or_init(||
                                                                      Lazy(Data_List_Lazy_Types_functorList_004038.clone()))
    }
    pub fn Data_List_Lazy_Types_functorList() -> &dyn Any {
        static Data_List_Lazy_Types_functorList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_functorList.get_or_init(||
                                                         Data_List_Lazy_Types_functorList_004038_002d1.Value)
    }
    pub fn Data_List_Lazy_Types_functorNonEmpty() -> &dyn Any {
        static Data_List_Lazy_Types_functorNonEmpty: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_functorNonEmpty.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_functorNonEmpty(),
                                                                                              &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorList()))
    }
    pub fn Data_List_Lazy_Types_functorNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_functorNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_functorNonEmptyList.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                  &&&add(string("map"),
                                                                                                         &&Func1::new(move
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
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorNonEmpty()),
                                                                                                                                                                                                                                                                                            &&&matchValue)),
                                                                                                                                                                                                                      &&&matchValue_1))
                                                                                                                                              }
                                                                                                                                      })),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_eq1List() -> &dyn Any {
        static Data_List_Lazy_Types_eq1List: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_eq1List.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                      &&&add(string("eq1"),
                                                                                             &&Func1::new(move
                                                                                                              |dictEq|
                                                                                                              &Func1::new({
                                                                                                                              let dictEq
                                                                                                                                  =
                                                                                                                                  dictEq.clone();
                                                                                                                              move
                                                                                                                                  |xs|
                                                                                                                                  &Func1::new({
                                                                                                                                                  let xs
                                                                                                                                                      =
                                                                                                                                                      xs.clone();
                                                                                                                                                  move
                                                                                                                                                      |ys|
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
                                                                                                                                                                                                     Func1::new({
                                                                                                                                                                                                                    let go_tco
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        go_tco.clone();
                                                                                                                                                                                                                    let v
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        v.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |v1|
                                                                                                                                                                                                                        go_tco(v)(v1.clone())
                                                                                                                                                                                                                })
                                                                                                                                                                                             })
                                                                                                                                                                         });
                                                                                                                                                          let go_1 =
                                                                                                                                                              Lazy(go_2);
                                                                                                                                                          fn go_tco(v_1:
                                                                                                                                                                        _)
                                                                                                                                                           ->
                                                                                                                                                               Func1<&dyn Any,
                                                                                                                                                                     &dyn Any> {
                                                                                                                                                              Func1::new({
                                                                                                                                                                             let go_tco
                                                                                                                                                                                 =
                                                                                                                                                                                 go_tco.clone();
                                                                                                                                                                             let v_1
                                                                                                                                                                                 =
                                                                                                                                                                                 v_1.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |v1_1|
                                                                                                                                                                                 {
                                                                                                                                                                                     let matchValue:
                                                                                                                                                                                             LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                     let matchValue_1:
                                                                                                                                                                                             LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                     if let PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                                                                          matchValue_1_1)
                                                                                                                                                                                            =
                                                                                                                                                                                            matchValue.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                         if let PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                                                                                              matchValue_1_1_1)
                                                                                                                                                                                                =
                                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                             if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                &&&dictEq),
                                                                                                                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                    PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                                                                                                                }),
                                                                                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                 PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                 _
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                                                                                                             }))
                                                                                                                                                                                                {
                                                                                                                                                                                                 go_tco(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                                            }))(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                        PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                                                                    }))
                                                                                                                                                                                             } else {
                                                                                                                                                                                                 &false
                                                                                                                                                                                             }
                                                                                                                                                                                         } else {
                                                                                                                                                                                             &false
                                                                                                                                                                                         }
                                                                                                                                                                                     } else {
                                                                                                                                                                                         if let PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor
                                                                                                                                                                                                =
                                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                             &true
                                                                                                                                                                                         } else {
                                                                                                                                                                                             &false
                                                                                                                                                                                         }
                                                                                                                                                                                     }
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                          }
                                                                                                                                                          let go =
                                                                                                                                                              go_1.Value;
                                                                                                                                                          go_tco(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                  &&&xs))(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                                           ys))
                                                                                                                                                      }
                                                                                                                                              })
                                                                                                                          })),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_eqNonEmpty() -> &dyn Any {
        static Data_List_Lazy_Types_eqNonEmpty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_eqNonEmpty.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_eqNonEmpty(),
                                                                                         &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eq1List()))
    }
    pub fn Data_List_Lazy_Types_eq1() -> &dyn Any {
        static Data_List_Lazy_Types_eq1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_eq1.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                  &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eq1List()))
    }
    pub fn Data_List_Lazy_Types_eq1NonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_eq1NonEmptyList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_eq1NonEmptyList.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                              &&&add(string("eq1"),
                                                                                                     &&Func1::new(move
                                                                                                                      |dictEq|
                                                                                                                      {
                                                                                                                          let eqNonEmpty1 =
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eqNonEmpty(),
                                                                                                                                                               dictEq);
                                                                                                                          &Func1::new({
                                                                                                                                          let eqNonEmpty1
                                                                                                                                              =
                                                                                                                                              eqNonEmpty1.clone();
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
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Lazy::Data_Lazy_eq1Lazy()),
                                                                                                                                                                                                                                                                             &&&eqNonEmpty1),
                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                       &&&matchValue_1)
                                                                                                                                                                  }
                                                                                                                                                          })
                                                                                                                                      })
                                                                                                                      }),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_eqList() -> &dyn Any {
        static Data_List_Lazy_Types_eqList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_eqList.get_or_init(||
                                                    &Func1::new(move |dictEq|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                     &&&add(string("eq"),
                                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eq1(),
                                                                                                                                              dictEq),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_eqNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_eqNonEmptyList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_eqNonEmptyList.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictEq|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_eqLazy(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eqNonEmpty(),
                                                                                                                                                dictEq))))
    }
    pub fn Data_List_Lazy_Types_ord1List() -> &dyn Any {
        static Data_List_Lazy_Types_ord1List: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_ord1List.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                       &&&add(string("compare1"),
                                                                                              &&Func1::new(move
                                                                                                               |dictOrd|
                                                                                                               &Func1::new({
                                                                                                                               let dictOrd
                                                                                                                                   =
                                                                                                                                   dictOrd.clone();
                                                                                                                               move
                                                                                                                                   |xs|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let xs
                                                                                                                                                       =
                                                                                                                                                       xs.clone();
                                                                                                                                                   move
                                                                                                                                                       |ys|
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
                                                                                                                                                                                                      Func1::new({
                                                                                                                                                                                                                     let go_tco
                                                                                                                                                                                                                         =
                                                                                                                                                                                                                         go_tco.clone();
                                                                                                                                                                                                                     let v
                                                                                                                                                                                                                         =
                                                                                                                                                                                                                         v.clone();
                                                                                                                                                                                                                     move
                                                                                                                                                                                                                         |v1|
                                                                                                                                                                                                                         go_tco(v)(v1.clone())
                                                                                                                                                                                                                 })
                                                                                                                                                                                              })
                                                                                                                                                                          });
                                                                                                                                                           let go_1 =
                                                                                                                                                               Lazy(go_2);
                                                                                                                                                           fn go_tco(v_1:
                                                                                                                                                                         _)
                                                                                                                                                            ->
                                                                                                                                                                Func1<&dyn Any,
                                                                                                                                                                      &dyn Any> {
                                                                                                                                                               Func1::new({
                                                                                                                                                                              let go_tco
                                                                                                                                                                                  =
                                                                                                                                                                                  go_tco.clone();
                                                                                                                                                                              let v_1
                                                                                                                                                                                  =
                                                                                                                                                                                  v_1.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |v1_1|
                                                                                                                                                                                  {
                                                                                                                                                                                      let matchValue:
                                                                                                                                                                                              LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                      let matchValue_1:
                                                                                                                                                                                              LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                      if let PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                                                                           matchValue_1_1)
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          if let PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                                                                                               matchValue_1_1_1)
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              let matchValue_3:
                                                                                                                                                                                                      LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                                                   &&&dictOrd),
                                                                                                                                                                                                                                                                                                &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                       PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                       _
                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                       unreachable!(),
                                                                                                                                                                                                                                                                                                   }),
                                                                                                                                                                                                                                                             &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                    PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                                                                                }));
                                                                                                                                                                                              if let Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_3.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  go_tco(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                 PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                 _
                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                                                                                             }))(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                         PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  &matchValue_3
                                                                                                                                                                                              }
                                                                                                                                                                                          } else {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                          }
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                          }
                                                                                                                                                                                      }
                                                                                                                                                                                  }
                                                                                                                                                                          })
                                                                                                                                                           }
                                                                                                                                                           let go =
                                                                                                                                                               go_1.Value;
                                                                                                                                                           go_tco(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                   &&&xs))(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                                            ys))
                                                                                                                                                       }
                                                                                                                                               })
                                                                                                                           })),
                                                                                              add(string("Eq10"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eq1List()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_ordNonEmpty() -> &dyn Any {
        static Data_List_Lazy_Types_ordNonEmpty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_ordNonEmpty.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_ordNonEmpty(),
                                                                                          &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_ord1List()))
    }
    pub fn Data_List_Lazy_Types_compare1() -> &dyn Any {
        static Data_List_Lazy_Types_compare1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_compare1.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_ord1List()))
    }
    pub fn Data_List_Lazy_Types_ord1NonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_ord1NonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_ord1NonEmptyList.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                               &&&add(string("compare1"),
                                                                                                      &&Func1::new(move
                                                                                                                       |dictOrd|
                                                                                                                       {
                                                                                                                           let ordNonEmpty1 =
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_ordNonEmpty(),
                                                                                                                                                                dictOrd);
                                                                                                                           &Func1::new({
                                                                                                                                           let ordNonEmpty1
                                                                                                                                               =
                                                                                                                                               ordNonEmpty1.clone();
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
                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Lazy::Data_Lazy_ord1Lazy()),
                                                                                                                                                                                                                                                                              &&&ordNonEmpty1),
                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                        &&&matchValue_1)
                                                                                                                                                                   }
                                                                                                                                                           })
                                                                                                                                       })
                                                                                                                       }),
                                                                                                      add(string("Eq10"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eq1NonEmptyList()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_ordList() -> &dyn Any {
        static Data_List_Lazy_Types_ordList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_ordList.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictOrd|
                                                                     {
                                                                         let eqList1 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eqList(),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                        Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                          &&&add(string("compare"),
                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_compare1(),
                                                                                                                                                   dictOrd),
                                                                                                                 add(string("Eq0"),
                                                                                                                     &&Func1::new({
                                                                                                                                      let eqList1
                                                                                                                                          =
                                                                                                                                          eqList1.clone();
                                                                                                                                      move
                                                                                                                                          |usd__unused|
                                                                                                                                          &eqList1
                                                                                                                                  }),
                                                                                                                     empty::<string,
                                                                                                                             &dyn Any>())))
                                                                     }))
    }
    pub fn Data_List_Lazy_Types_ordNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_ordNonEmptyList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_ordNonEmptyList.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictOrd|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_ordLazy(),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_ordNonEmpty(),
                                                                                                                                                 dictOrd))))
    }
    pub fn Data_List_Lazy_Types_cons() -> &dyn Any {
        static Data_List_Lazy_Types_cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_cons.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  &Func1::new({
                                                                                  let x
                                                                                      =
                                                                                      x.clone();
                                                                                  move
                                                                                      |xs|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                          &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                            let xs
                                                                                                                                                                                =
                                                                                                                                                                                xs.clone();
                                                                                                                                                                            move
                                                                                                                                                                                |v|
                                                                                                                                                                                &LrcPtr::new(PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&x,
                                                                                                                                                                                                                                                                                           &xs))
                                                                                                                                                                        })))
                                                                              })))
    }
    pub fn Data_List_Lazy_Types_foldableList_004083() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                         &&&add(string("foldr"),
                                                &&Func1::new({
                                                                 let Data_List_Lazy_Types_foldableList_004083_002d1
                                                                     =
                                                                     Data_List_Lazy_Types_foldableList_004083_002d1.clone();
                                                                 move |op|
                                                                     &Func1::new({
                                                                                     let Data_List_Lazy_Types_foldableList_004083_002d1
                                                                                         =
                                                                                         Data_List_Lazy_Types_foldableList_004083_002d1.clone();
                                                                                     let op
                                                                                         =
                                                                                         op.clone();
                                                                                     move
                                                                                         |z|
                                                                                         &Func1::new({
                                                                                                         let Data_List_Lazy_Types_foldableList_004083_002d1
                                                                                                             =
                                                                                                             Data_List_Lazy_Types_foldableList_004083_002d1.clone();
                                                                                                         let z
                                                                                                             =
                                                                                                             z.clone();
                                                                                                         move
                                                                                                             |xs|
                                                                                                             {
                                                                                                                 let rev =
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                            &&&Data_List_Lazy_Types_foldableList_004083_002d1.Value),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons())),
                                                                                                                                                      &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil());
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                           &&&Data_List_Lazy_Types_foldableList_004083_002d1.Value),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                           &&&op)),
                                                                                                                                                                                     &&&z),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&rev,
                                                                                                                                                                                     xs))
                                                                                                             }
                                                                                                     })
                                                                                 })
                                                             }),
                                                add(string("foldl"),
                                                    &&Func1::new(move |op_1|
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
                                                                                                                    |b|
                                                                                                                    Func1::new({
                                                                                                                                   let b
                                                                                                                                       =
                                                                                                                                       b.clone();
                                                                                                                                   let go_tco
                                                                                                                                       =
                                                                                                                                       go_tco.clone();
                                                                                                                                   move
                                                                                                                                       |xs_1|
                                                                                                                                       go_tco(b)(xs_1.clone())
                                                                                                                               })
                                                                                                            })
                                                                                        });
                                                                         let go_1 =
                                                                             Lazy(go_2);
                                                                         let go_tco =
                                                                             Func1::new({
                                                                                            let op_1
                                                                                                =
                                                                                                op_1.clone();
                                                                                            move
                                                                                                |b_1|
                                                                                                fix1(&(move
                                                                                                           |go_tco,
                                                                                                            b_1|
                                                                                                           Func1::new({
                                                                                                                          let b_1
                                                                                                                              =
                                                                                                                              b_1.clone();
                                                                                                                          let go_tco
                                                                                                                              =
                                                                                                                              go_tco.clone();
                                                                                                                          move
                                                                                                                              |xs_2|
                                                                                                                              {
                                                                                                                                  let matchValue:
                                                                                                                                          LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                 xs_2));
                                                                                                                                  match matchValue.as_ref()
                                                                                                                                      {
                                                                                                                                      PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                    matchValue_1_1)
                                                                                                                                      =>
                                                                                                                                      go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&op_1,
                                                                                                                                                                                                                 &&&b_1),
                                                                                                                                                                              &&matchValue_1_0))(matchValue_1_1),
                                                                                                                                      _
                                                                                                                                      =>
                                                                                                                                      &b_1,
                                                                                                                                  }
                                                                                                                              }
                                                                                                                      })),
                                                                                                     b_1.clone())
                                                                                        });
                                                                         let go =
                                                                             go_1.Value;
                                                                         &go
                                                                     }),
                                                    add(string("foldMap"),
                                                        &&Func1::new({
                                                                         let Data_List_Lazy_Types_foldableList_004083_002d1
                                                                             =
                                                                             Data_List_Lazy_Types_foldableList_004083_002d1.clone();
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
                                                                                                 let Data_List_Lazy_Types_foldableList_004083_002d1
                                                                                                     =
                                                                                                     Data_List_Lazy_Types_foldableList_004083_002d1.clone();
                                                                                                 let Semigroup0
                                                                                                     =
                                                                                                     Semigroup0.clone();
                                                                                                 let mempty
                                                                                                     =
                                                                                                     mempty.clone();
                                                                                                 move
                                                                                                     |f|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                            &&&Data_List_Lazy_Types_foldableList_004083_002d1.Value),
                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                           let f
                                                                                                                                                                                               =
                                                                                                                                                                                               f.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |b_2|
                                                                                                                                                                                               &Func1::new({
                                                                                                                                                                                                               let b_2
                                                                                                                                                                                                                   =
                                                                                                                                                                                                                   b_2.clone();
                                                                                                                                                                                                               move
                                                                                                                                                                                                                   |a|
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                          &&&Semigroup0),
                                                                                                                                                                                                                                                                                       &&&b_2),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                       a))
                                                                                                                                                                                                           })
                                                                                                                                                                                       })),
                                                                                                                                      &&&mempty)
                                                                                             })
                                                                             }
                                                                     }),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_foldableList_004083_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_Types_foldableList_004083_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_Types_foldableList_004083_002d1.get_or_init(||
                                                                       Lazy(Data_List_Lazy_Types_foldableList_004083.clone()))
    }
    pub fn Data_List_Lazy_Types_foldableList() -> &dyn Any {
        static Data_List_Lazy_Types_foldableList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_foldableList.get_or_init(||
                                                          Data_List_Lazy_Types_foldableList_004083_002d1.Value)
    }
    pub fn Data_List_Lazy_Types_foldableNonEmpty() -> &dyn Any {
        static Data_List_Lazy_Types_foldableNonEmpty:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_foldableNonEmpty.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_foldableNonEmpty(),
                                                                                               &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()))
    }
    pub fn Data_List_Lazy_Types_extendList() -> &dyn Any {
        static Data_List_Lazy_Types_extendList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_extendList.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                         &&&add(string("extend"),
                                                                                                &&Func1::new(move
                                                                                                                 |f|
                                                                                                                 &Func1::new({
                                                                                                                                 let f
                                                                                                                                     =
                                                                                                                                     f.clone();
                                                                                                                                 move
                                                                                                                                     |l|
                                                                                                                                     {
                                                                                                                                         let go =
                                                                                                                                             &Func1::new(move
                                                                                                                                                             |a|
                                                                                                                                                             &Func1::new({
                                                                                                                                                                             let a
                                                                                                                                                                                 =
                                                                                                                                                                                 a.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |v|
                                                                                                                                                                                 {
                                                                                                                                                                                     let matchValue =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&a);
                                                                                                                                                                                     let matchValue_1 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                     {
                                                                                                                                                                                         let activePatternResult =
                                                                                                                                                                                             Sharpurs_Prelude::_007cHasProp_007c__007c(string("val"),
                                                                                                                                                                                                                                       &matchValue_1);
                                                                                                                                                                                         if activePatternResult.is_some()
                                                                                                                                                                                            {
                                                                                                                                                                                             let activePatternResult_1 =
                                                                                                                                                                                                 Sharpurs_Prelude::_007cHasProp_007c__007c(string("acc"),
                                                                                                                                                                                                                                           &matchValue_1);
                                                                                                                                                                                             if activePatternResult_1.is_some()
                                                                                                                                                                                                {
                                                                                                                                                                                                 let acc =
                                                                                                                                                                                                     getValue(activePatternResult_1);
                                                                                                                                                                                                 let val_var =
                                                                                                                                                                                                     getValue(activePatternResult);
                                                                                                                                                                                                 let acc_prime =
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                                                                      &&&acc);
                                                                                                                                                                                                 &add(string("val"),
                                                                                                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                              &&&acc_prime)),
                                                                                                                                                                                                                                        &&&val_var),
                                                                                                                                                                                                      add(string("acc"),
                                                                                                                                                                                                          &&acc_prime,
                                                                                                                                                                                                          empty::<string,
                                                                                                                                                                                                                  &dyn Any>()))
                                                                                                                                                                                             } else {
                                                                                                                                                                                                 panic!("{}",
                                                                                                                                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.Types.fs"),
                                  Data1: 94_i32,
                                  Data2: 240_i32,}).get_Message(),)
                                                                                                                                                                                             }
                                                                                                                                                                                         } else {
                                                                                                                                                                                             panic!("{}",
                                                                                                                                                                                                    LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.Types.fs"),
                                  Data1: 94_i32,
                                  Data2: 240_i32,}).get_Message(),)
                                                                                                                                                                                         }
                                                                                                                                                                                     }
                                                                                                                                                                                 }
                                                                                                                                                                         }));
                                                                                                                                         let matchValue_3:
                                                                                                                                                 LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                        l));
                                                                                                                                         match matchValue_3.as_ref()
                                                                                                                                             {
                                                                                                                                             PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_3_1_0,
                                                                                                                                                                                                                                           matchValue_3_1_1)
                                                                                                                                             =>
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                    l)),
                                                                                                                                                                              &&find(string("val"),
                                                                                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                                                                                                                                                     &&&go),
                                                                                                                                                                                                                                                                                  &&&add(string("val"),
                                                                                                                                                                                                                                                                                         &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                                                                                                                                                                                                         add(string("acc"),
                                                                                                                                                                                                                                                                                             &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                                                                                                                                                                                                             empty::<string,
                                                                                                                                                                                                                                                                                                     &dyn Any>()))),
                                                                                                                                                                                                                                               &&matchValue_3_1_1)))),
                                                                                                                                             _
                                                                                                                                             =>
                                                                                                                                             &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                                                         }
                                                                                                                                     }
                                                                                                                             })),
                                                                                                add(string("Functor0"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorList()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_extendNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_extendNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_extendNonEmptyList.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                                 &&&add(string("extend"),
                                                                                                        &&Func1::new(move
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
                                                                                                                                                 let go =
                                                                                                                                                     &Func1::new({
                                                                                                                                                                     let f1
                                                                                                                                                                         =
                                                                                                                                                                         f1.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |a|
                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                         let a
                                                                                                                                                                                             =
                                                                                                                                                                                             a.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |v1|
                                                                                                                                                                                             {
                                                                                                                                                                                                 let matchValue_3 =
                                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&a);
                                                                                                                                                                                                 let matchValue_4 =
                                                                                                                                                                                                     Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                 {
                                                                                                                                                                                                     let activePatternResult =
                                                                                                                                                                                                         Sharpurs_Prelude::_007cHasProp_007c__007c(string("val"),
                                                                                                                                                                                                                                                   &matchValue_4);
                                                                                                                                                                                                     if activePatternResult.is_some()
                                                                                                                                                                                                        {
                                                                                                                                                                                                         let activePatternResult_1 =
                                                                                                                                                                                                             Sharpurs_Prelude::_007cHasProp_007c__007c(string("acc"),
                                                                                                                                                                                                                                                       &matchValue_4);
                                                                                                                                                                                                         if activePatternResult_1.is_some()
                                                                                                                                                                                                            {
                                                                                                                                                                                                             let acc =
                                                                                                                                                                                                                 getValue(activePatternResult_1);
                                                                                                                                                                                                             let val_var =
                                                                                                                                                                                                                 getValue(activePatternResult);
                                                                                                                                                                                                             &add(string("val"),
                                                                                                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                  let acc
                                                                                                                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                                                                                                                      acc.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                  let matchValue_3
                                                                                                                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                                                                                                                      matchValue_3.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                                                                                                                                      |v2|
                                                                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&matchValue_3,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &acc))
                                                                                                                                                                                                                                                                                                                                                                                                              }))))),
                                                                                                                                                                                                                                                    &&&val_var),
                                                                                                                                                                                                                  add(string("acc"),
                                                                                                                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                           &&&matchValue_3),
                                                                                                                                                                                                                                                        &&&acc),
                                                                                                                                                                                                                      empty::<string,
                                                                                                                                                                                                                              &dyn Any>()))
                                                                                                                                                                                                         } else {
                                                                                                                                                                                                             panic!("{}",
                                                                                                                                                                                                                    LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.Types.fs"),
                                  Data1: 96_i32,
                                  Data2: 325_i32,}).get_Message(),)
                                                                                                                                                                                                         }
                                                                                                                                                                                                     } else {
                                                                                                                                                                                                         panic!("{}",
                                                                                                                                                                                                                LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.Types.fs"),
                                  Data1: 96_i32,
                                  Data2: 325_i32,}).get_Message(),)
                                                                                                                                                                                                     }
                                                                                                                                                                                                 }
                                                                                                                                                                                             }
                                                                                                                                                                                     })
                                                                                                                                                                 });
                                                                                                                                                 let matchValue_6:
                                                                                                                                                         LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                &&&matchValue_1));
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList()),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                     &&&Func1::new({
                                                                                                                                                                                                                                       let f1
                                                                                                                                                                                                                                           =
                                                                                                                                                                                                                                           f1.clone();
                                                                                                                                                                                                                                       let go
                                                                                                                                                                                                                                           =
                                                                                                                                                                                                                                           go.clone();
                                                                                                                                                                                                                                       let matchValue_1
                                                                                                                                                                                                                                           =
                                                                                                                                                                                                                                           matchValue_1.clone();
                                                                                                                                                                                                                                       let matchValue_6
                                                                                                                                                                                                                                           =
                                                                                                                                                                                                                                           matchValue_6.clone();
                                                                                                                                                                                                                                       move
                                                                                                                                                                                                                                           |v2_1|
                                                                                                                                                                                                                                           &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                                                                                &&&matchValue_1),
                                                                                                                                                                                                                                                                                                               find(string("val"),
                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&go),
                                                                                                                                                                                                                                                                                                                                                                                                                 &&&add(string("val"),
                                                                                                                                                                                                                                                                                                                                                                                                                        &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                                                                                                                                                                                                                                                                                                                                        add(string("acc"),
                                                                                                                                                                                                                                                                                                                                                                                                                            &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                                                                                                                                                                                                                                                                                                                                            empty::<string,
                                                                                                                                                                                                                                                                                                                                                                                                                                    &dyn Any>()))),
                                                                                                                                                                                                                                                                                                                                                                              &&&match matchValue_6.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                     Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                 })))))
                                                                                                                                                                                                                                   })))
                                                                                                                                             }
                                                                                                                                     })),
                                                                                                        add(string("Functor0"),
                                                                                                            &&Func1::new(move
                                                                                                                             |usd__unused|
                                                                                                                             &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorNonEmptyList()),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_foldableNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_foldableNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_foldableNonEmptyList.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                                   &&&add(string("foldr"),
                                                                                                          &&Func1::new(move
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
                                                                                                                                                                   |v|
                                                                                                                                                                   {
                                                                                                                                                                       let matchValue =
                                                                                                                                                                           Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                       let matchValue_1 =
                                                                                                                                                                           Sharpurs_Prelude::unbox(&&b);
                                                                                                                                                                       let matchValue_2 =
                                                                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableNonEmpty()),
                                                                                                                                                                                                                                                                              &&&matchValue),
                                                                                                                                                                                                                                           &&&matchValue_1),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                           &&&matchValue_2))
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
                                                                                                                                                   |b_1|
                                                                                                                                                   &Func1::new({
                                                                                                                                                                   let b_1
                                                                                                                                                                       =
                                                                                                                                                                       b_1.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |v_1|
                                                                                                                                                                       {
                                                                                                                                                                           let matchValue_4 =
                                                                                                                                                                               Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                           let matchValue_5 =
                                                                                                                                                                               Sharpurs_Prelude::unbox(&&b_1);
                                                                                                                                                                           let matchValue_6 =
                                                                                                                                                                               Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableNonEmpty()),
                                                                                                                                                                                                                                                                                  &&&matchValue_4),
                                                                                                                                                                                                                                               &&&matchValue_5),
                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                               &&&matchValue_6))
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
                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableNonEmpty()),
                                                                                                                                                                                                                                                                                      &&&dictMonoid),
                                                                                                                                                                                                                                                   &&&matchValue_8),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                   &&&matchValue_9))
                                                                                                                                                                           }
                                                                                                                                                                   })
                                                                                                                                               })),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>())))))
    }
    pub fn Data_List_Lazy_Types_showList() -> &dyn Any {
        static Data_List_Lazy_Types_showList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_showList.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictShow|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                       &&&add(string("show"),
                                                                                                              &&Func1::new({
                                                                                                                               let dictShow
                                                                                                                                   =
                                                                                                                                   dictShow.clone();
                                                                                                                               move
                                                                                                                                   |xs|
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                       &&&string("(fromFoldable [")),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                          &&{
                                                                                                                                                                                                                                                let matchValue:
                                                                                                                                                                                                                                                        LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                                                                                                               xs));
                                                                                                                                                                                                                                                match matchValue.as_ref()
                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                    PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                                                                                                                                  matchValue_1_1)
                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                              &&&dictShow),
                                                                                                                                                                                                                                                                                                                                                           &&matchValue_1_0)),
                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                |shown|
                                                                                                                                                                                                                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                let shown
                                                                                                                                                                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                                                                                                                                                                    shown.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                                                                                                                                                                    |x_prime|
                                                                                                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&shown),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&string(",")),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&dictShow),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           x_prime)))
                                                                                                                                                                                                                                                                                                                                                                                                                            }))),
                                                                                                                                                                                                                                                                                                                                                           &&&string("")),
                                                                                                                                                                                                                                                                                                                        &&matchValue_1_1)),
                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                    &string(""),
                                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                       &&&string("])")))
                                                                                                                           }),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_showNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_showNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_showNonEmptyList.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictShow|
                                                                              {
                                                                                  let showLazy =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_showLazy(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_showNonEmpty(),
                                                                                                                                                                                             dictShow),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_showList(),
                                                                                                                                                                                             dictShow)));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                   &&&add(string("show"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let showLazy
                                                                                                                                               =
                                                                                                                                               showLazy.clone();
                                                                                                                                           move
                                                                                                                                               |v|
                                                                                                                                               {
                                                                                                                                                   let nel =
                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                       &&&string("(NonEmptyList ")),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                &&&showLazy),
                                                                                                                                                                                                                                                                                             &&&nel)),
                                                                                                                                                                                                                       &&&string(")")))
                                                                                                                                               }
                                                                                                                                       }),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))
                                                                              }))
    }
    pub fn Data_List_Lazy_Types_showStep() -> &dyn Any {
        static Data_List_Lazy_Types_showStep: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_showStep.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictShow|
                                                                      {
                                                                          let showList1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_showList(),
                                                                                                               dictShow);
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                           &&&add(string("show"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let dictShow
                                                                                                                                       =
                                                                                                                                       dictShow.clone();
                                                                                                                                   let showList1
                                                                                                                                       =
                                                                                                                                       showList1.clone();
                                                                                                                                   move
                                                                                                                                       |v|
                                                                                                                                       {
                                                                                                                                           let matchValue:
                                                                                                                                                   LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                           match matchValue.as_ref()
                                                                                                                                               {
                                                                                                                                               PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                             matchValue_1_1)
                                                                                                                                               =>
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                   &&&string("(")),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                            &&&dictShow),
                                                                                                                                                                                                                                                                                         &&matchValue_1_0)),
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                         &&&string(" : ")),
                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                  &&&showList1),
                                                                                                                                                                                                                                                                                                                                                               &&matchValue_1_1)),
                                                                                                                                                                                                                                                                                         &&&string(")"))))),
                                                                                                                                               _
                                                                                                                                               =>
                                                                                                                                               &string("Nil"),
                                                                                                                                           }
                                                                                                                                       }
                                                                                                                               }),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))
                                                                      }))
    }
    pub fn Data_List_Lazy_Types_foldableWithIndexList_0040106() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                         &&&add(string("foldrWithIndex"),
                                                &&Func1::new(move |f|
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
                                                                                                         |xs|
                                                                                                         {
                                                                                                             let matchValue_1:
                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                      &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                                     |v1|
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                         let matchValue:
                                                                                                                                                                                                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                                                                         &Func1::new(move
                                                                                                                                                                                                                                                                                         |a|
                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                                                                                                                                                  &&&1_i32),
                                                                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                                                                                                                                                     a),
                                                                                                                                                                                                                                                                                                                                                                                  &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                     }))))
                                                                                                                                                                                                                                                                     })),
                                                                                                                                                                                                               &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&0_i32,
                                                                                                                                                                                                                                                                         &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()))),
                                                                                                                                                                            xs));
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                 &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                          &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                         |v1_1|
                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                             let matchValue_2:
                                                                                                                                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                                                                                                             let i_1 =
                                                                                                                                                                                                                                                                                 match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                 };
                                                                                                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                                                                                                             let i_1
                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                 i_1.clone();
                                                                                                                                                                                                                                                                                             let matchValue_2
                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                 matchValue_2.clone();
                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                 |a_1|
                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                             &&&i_1),
                                                                                                                                                                                                                                                                                                                                                                                          &&&1_i32),
                                                                                                                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&i_1),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&1_i32)),
                                                                                                                                                                                                                                                                                                                                                                                                                             a_1),
                                                                                                                                                                                                                                                                                                                                                                                          &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                             })))
                                                                                                                                                                                                                                                                                         })
                                                                                                                                                                                                                                                                         })),
                                                                                                                                                                                                                    &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               },
                                                                                                                                                                                                                                                                              &b))),
                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                   x)
                                                                                                                                                                                        =>
                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                    }))
                                                                                                         }
                                                                                                 })
                                                                             })),
                                                add(string("foldlWithIndex"),
                                                    &&Func1::new(move |f_1|
                                                                     &Func1::new({
                                                                                     let f_1
                                                                                         =
                                                                                         f_1.clone();
                                                                                     move
                                                                                         |acc_1|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                             &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                   &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                  |v_1|
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                      let matchValue_3:
                                                                                                                                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                                      let i_2 =
                                                                                                                                                                                                                          match matchValue_3.as_ref()
                                                                                                                                                                                                                              {
                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                          };
                                                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                                                      let i_2
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          i_2.clone();
                                                                                                                                                                                                                                      let matchValue_3
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          matchValue_3.clone();
                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                          |a_2|
                                                                                                                                                                                                                                          &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                                      &&&i_2),
                                                                                                                                                                                                                                                                                                                                   &&&1_i32),
                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f_1,
                                                                                                                                                                                                                                                                                                                                                                                                         &&&i_2),
                                                                                                                                                                                                                                                                                                                                                                      &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                                                                                                   a_2)))
                                                                                                                                                                                                                                  })
                                                                                                                                                                                                                  })),
                                                                                                                                                             &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&0_i32,
                                                                                                                                                                                                                       acc_1.clone()))))
                                                                                 })),
                                                    add(string("foldMapWithIndex"),
                                                        &&Func1::new({
                                                                         let Data_List_Lazy_Types_foldableWithIndexList_0040106_002d1
                                                                             =
                                                                             Data_List_Lazy_Types_foldableWithIndexList_0040106_002d1.clone();
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
                                                                                                 let Data_List_Lazy_Types_foldableWithIndexList_0040106_002d1
                                                                                                     =
                                                                                                     Data_List_Lazy_Types_foldableWithIndexList_0040106_002d1.clone();
                                                                                                 let Semigroup0
                                                                                                     =
                                                                                                     Semigroup0.clone();
                                                                                                 let mempty
                                                                                                     =
                                                                                                     mempty.clone();
                                                                                                 move
                                                                                                     |f_2|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                            &&&Data_List_Lazy_Types_foldableWithIndexList_0040106_002d1.Value),
                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                           let f_2
                                                                                                                                                                                               =
                                                                                                                                                                                               f_2.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |i_3|
                                                                                                                                                                                               &Func1::new({
                                                                                                                                                                                                               let i_3
                                                                                                                                                                                                                   =
                                                                                                                                                                                                                   i_3.clone();
                                                                                                                                                                                                               move
                                                                                                                                                                                                                   |acc_2|
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                             &&&Semigroup0),
                                                                                                                                                                                                                                                                                                                          acc_2)),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&f_2,
                                                                                                                                                                                                                                                                                       &&&i_3))
                                                                                                                                                                                                           })
                                                                                                                                                                                       })),
                                                                                                                                      &&&mempty)
                                                                                             })
                                                                             }
                                                                     }),
                                                        add(string("Foldable0"),
                                                            &&Func1::new(move
                                                                             |usd__unused|
                                                                             &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_List_Lazy_Types_foldableWithIndexList_0040106_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_Types_foldableWithIndexList_0040106_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_Types_foldableWithIndexList_0040106_002d1.get_or_init(||
                                                                                 Lazy(Data_List_Lazy_Types_foldableWithIndexList_0040106.clone()))
    }
    pub fn Data_List_Lazy_Types_foldableWithIndexList() -> &dyn Any {
        static Data_List_Lazy_Types_foldableWithIndexList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_foldableWithIndexList.get_or_init(||
                                                                   Data_List_Lazy_Types_foldableWithIndexList_0040106_002d1.Value)
    }
    pub fn Data_List_Lazy_Types_foldableWithIndexNonEmpty() -> &dyn Any {
        static Data_List_Lazy_Types_foldableWithIndexNonEmpty:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_foldableWithIndexNonEmpty.get_or_init(||
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_foldableWithIndexNonEmpty(),
                                                                                                        &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableWithIndexList()))
    }
    pub fn Data_List_Lazy_Types_foldableWithIndexNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_foldableWithIndexNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_foldableWithIndexNonEmptyList.get_or_init(||
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                            &&&add(string("foldMapWithIndex"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |dictMonoid|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let dictMonoid
                                                                                                                                                        =
                                                                                                                                                        dictMonoid.clone();
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
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableWithIndexNonEmpty()),
                                                                                                                                                                                                                                                                                                                          &&&dictMonoid),
                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&0_i32),
                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                &&&1_i32))))),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                    &&&matchValue_1))
                                                                                                                                                                            }
                                                                                                                                                                    })
                                                                                                                                                })),
                                                                                                                   add(string("foldlWithIndex"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |f_1|
                                                                                                                                        &Func1::new({
                                                                                                                                                        let f_1
                                                                                                                                                            =
                                                                                                                                                            f_1.clone();
                                                                                                                                                        move
                                                                                                                                                            |b|
                                                                                                                                                            &Func1::new({
                                                                                                                                                                            let b
                                                                                                                                                                                =
                                                                                                                                                                                b.clone();
                                                                                                                                                                            move
                                                                                                                                                                                |v_1|
                                                                                                                                                                                {
                                                                                                                                                                                    let matchValue_3 =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                                    let matchValue_4 =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&b);
                                                                                                                                                                                    let matchValue_5 =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableWithIndexNonEmpty()),
                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                    &&&matchValue_3),
                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&0_i32),
                                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&1_i32)))),
                                                                                                                                                                                                                                                                                           &&&matchValue_4)),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                        &&&matchValue_5))
                                                                                                                                                                                }
                                                                                                                                                                        })
                                                                                                                                                    })),
                                                                                                                       add(string("foldrWithIndex"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |f_2|
                                                                                                                                            &Func1::new({
                                                                                                                                                            let f_2
                                                                                                                                                                =
                                                                                                                                                                f_2.clone();
                                                                                                                                                            move
                                                                                                                                                                |b_1|
                                                                                                                                                                &Func1::new({
                                                                                                                                                                                let b_1
                                                                                                                                                                                    =
                                                                                                                                                                                    b_1.clone();
                                                                                                                                                                                move
                                                                                                                                                                                    |v_2|
                                                                                                                                                                                    {
                                                                                                                                                                                        let matchValue_7 =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                                        let matchValue_8 =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&b_1);
                                                                                                                                                                                        let matchValue_9 =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableWithIndexNonEmpty()),
                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                        &&&matchValue_7),
                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&0_i32),
                                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&1_i32)))),
                                                                                                                                                                                                                                                                                               &&&matchValue_8)),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                            &&&matchValue_9))
                                                                                                                                                                                    }
                                                                                                                                                                            })
                                                                                                                                                        })),
                                                                                                                           add(string("Foldable0"),
                                                                                                                               &&Func1::new(move
                                                                                                                                                |usd__unused|
                                                                                                                                                &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableNonEmptyList()),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>()))))))
    }
    pub fn Data_List_Lazy_Types_functorWithIndexList() -> &dyn Any {
        static Data_List_Lazy_Types_functorWithIndexList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_functorWithIndexList.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                   &&&add(string("mapWithIndex"),
                                                                                                          &&Func1::new(move
                                                                                                                           |f|
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                  &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableWithIndexList()),
                                                                                                                                                                                               &&&Func1::new({
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
                                                                                                                                                                                                                                         |x|
                                                                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                                                                         let x
                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                             x.clone();
                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                             |acc|
                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                       &&&i),
                                                                                                                                                                                                                                                                                                                                                                    &&&x)),
                                                                                                                                                                                                                                                                                              acc)
                                                                                                                                                                                                                                                     })
                                                                                                                                                                                                                                 })
                                                                                                                                                                                                             })),
                                                                                                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())),
                                                                                                          add(string("Functor0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorList()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_functorWithIndex() -> &dyn Any {
        static Data_List_Lazy_Types_functorWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_functorWithIndex.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_functorWithIndex(),
                                                                                               &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorWithIndexList()))
    }
    pub fn Data_List_Lazy_Types_functorWithIndexNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_functorWithIndexNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_functorWithIndexNonEmptyList.get_or_init(||
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                           &&&add(string("mapWithIndex"),
                                                                                                                  &&Func1::new(move
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
                                                                                                                                                                                                                               &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList()),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                                                 let matchValue_1
                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                     matchValue_1.clone();
                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                     |v1|
                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorWithIndex()),
                                                                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&0_i32),
                                                                                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&1_i32))))),
                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                                                         &&&matchValue_1))
                                                                                                                                                                                                                                             })))
                                                                                                                                                       }
                                                                                                                                               })),
                                                                                                                  add(string("Functor0"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |usd__unused|
                                                                                                                                       &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorNonEmptyList()),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_toList() -> &dyn Any {
        static Data_List_Lazy_Types_toList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_toList.get_or_init(||
                                                    &Func1::new(move |v|
                                                                    {
                                                                        let nel =
                                                                            Sharpurs_Prelude::unbox(v);
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_lazyList()),
                                                                                                         &&&Func1::new({
                                                                                                                           let nel
                                                                                                                               =
                                                                                                                               nel.clone();
                                                                                                                           move
                                                                                                                               |v1|
                                                                                                                               {
                                                                                                                                   let matchValue:
                                                                                                                                           LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                  &&&nel));
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                          }),
                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                           Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                  x)
                                                                                                                                                                           =>
                                                                                                                                                                           x.clone(),
                                                                                                                                                                       })
                                                                                                                               }
                                                                                                                       }))
                                                                    }))
    }
    pub fn Data_List_Lazy_Types_semigroupNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_semigroupNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_semigroupNonEmptyList.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                    &&&add(string("append"),
                                                                                                           &&Func1::new(move
                                                                                                                            |v|
                                                                                                                            &Func1::new({
                                                                                                                                            let v
                                                                                                                                                =
                                                                                                                                                v.clone();
                                                                                                                                            move
                                                                                                                                                |as_prime|
                                                                                                                                                {
                                                                                                                                                    let matchValue =
                                                                                                                                                        Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                    let matchValue_1 =
                                                                                                                                                        Sharpurs_Prelude::unbox(as_prime);
                                                                                                                                                    let matchValue_3:
                                                                                                                                                            LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                                        Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                   &&&matchValue));
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                          let matchValue_1
                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                              matchValue_1.clone();
                                                                                                                                                                                                                                          let matchValue_3
                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                              matchValue_3.clone();
                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                              |v2|
                                                                                                                                                                                                                                              &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                       Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                   },
                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                                                                                                                                                                                                                                                                                                                      &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                                             Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_toList(),
                                                                                                                                                                                                                                                                                                                                                                                      &&&matchValue_1))))
                                                                                                                                                                                                                                      })))
                                                                                                                                                }
                                                                                                                                        })),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_traversableList_0040123() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                         &&&add(string("traverse"),
                                                &&Func1::new(move
                                                                 |dictApplicative|
                                                                 {
                                                                     let Apply0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                 Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     let Functor0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     &Func1::new({
                                                                                     let Apply0
                                                                                         =
                                                                                         Apply0.clone();
                                                                                     let Functor0
                                                                                         =
                                                                                         Functor0.clone();
                                                                                     let dictApplicative
                                                                                         =
                                                                                         dictApplicative.clone();
                                                                                     move
                                                                                         |f|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                             &&&Func1::new({
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
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                              &&&Apply0),
                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons()),
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                 &&&a))),
                                                                                                                                                                                                                                        b)
                                                                                                                                                                                               })
                                                                                                                                                                           })),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                &&&dictApplicative),
                                                                                                                                                             &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()))
                                                                                 })
                                                                 }),
                                                add(string("sequence"),
                                                    &&Func1::new({
                                                                     let Data_List_Lazy_Types_traversableList_0040123_002d1
                                                                         =
                                                                         Data_List_Lazy_Types_traversableList_0040123_002d1.clone();
                                                                     move
                                                                         |dictApplicative_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                &&&Data_List_Lazy_Types_traversableList_0040123_002d1.Value),
                                                                                                                                             dictApplicative_1),
                                                                                                          &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_identity())
                                                                 }),
                                                    add(string("Functor0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorList()),
                                                        add(string("Foldable1"),
                                                            &&Func1::new(move
                                                                             |usd__unused_1|
                                                                             &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_List_Lazy_Types_traversableList_0040123_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_Types_traversableList_0040123_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_Types_traversableList_0040123_002d1.get_or_init(||
                                                                           Lazy(Data_List_Lazy_Types_traversableList_0040123.clone()))
    }
    pub fn Data_List_Lazy_Types_traversableList() -> &dyn Any {
        static Data_List_Lazy_Types_traversableList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_traversableList.get_or_init(||
                                                             Data_List_Lazy_Types_traversableList_0040123_002d1.Value)
    }
    pub fn Data_List_Lazy_Types_traversableNonEmpty() -> &dyn Any {
        static Data_List_Lazy_Types_traversableNonEmpty:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_traversableNonEmpty.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_traversableNonEmpty(),
                                                                                                  &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_traversableList()))
    }
    pub fn Data_List_Lazy_Types_traversableNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_traversableNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_traversableNonEmptyList.get_or_init(||
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
                                                                                                                                                                          |v|
                                                                                                                                                                          {
                                                                                                                                                                              let matchValue =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                              let matchValue_1 =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                        &&&Functor0),
                                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                                       |xxs|
                                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList()),
                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                             let xxs
                                                                                                                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                                                                                                                 xxs.clone();
                                                                                                                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                                                                                                                 |v1|
                                                                                                                                                                                                                                                                                                                                                                                                 &xxs
                                                                                                                                                                                                                                                                                                                                                                                         })))))),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_traversableNonEmpty()),
                                                                                                                                                                                                                                                                                                                        &&&dictApplicative),
                                                                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                     &&&matchValue_1)))
                                                                                                                                                                          }
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
                                                                                                                                                      let dictApplicative_1
                                                                                                                                                          =
                                                                                                                                                          dictApplicative_1.clone();
                                                                                                                                                      move
                                                                                                                                                          |v_1|
                                                                                                                                                          {
                                                                                                                                                              let nel_1 =
                                                                                                                                                                  Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                        &&&Functor0_1),
                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                       |xxs_1|
                                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList()),
                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                             let xxs_1
                                                                                                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                                                                                                 xxs_1.clone();
                                                                                                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                                                                                                 |v1_1|
                                                                                                                                                                                                                                                                                                                                                                                 &xxs_1
                                                                                                                                                                                                                                                                                                                                                                         })))))),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_traversableNonEmpty()),
                                                                                                                                                                                                                                                                     &&&dictApplicative_1),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                     &&&nel_1)))
                                                                                                                                                          }
                                                                                                                                                  })
                                                                                                                                  }),
                                                                                                                 add(string("Functor0"),
                                                                                                                     &&Func1::new(move
                                                                                                                                      |usd__unused|
                                                                                                                                      &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorNonEmptyList()),
                                                                                                                     add(string("Foldable1"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused_1|
                                                                                                                                          &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableNonEmptyList()),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))))
    }
    pub fn Data_List_Lazy_Types_traversableWithIndexList() -> &dyn Any {
        static Data_List_Lazy_Types_traversableWithIndexList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_traversableWithIndexList.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                       &&&add(string("traverseWithIndex"),
                                                                                                              &&Func1::new(move
                                                                                                                               |dictApplicative|
                                                                                                                               {
                                                                                                                                   let Apply0 =
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                               Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                   let Functor0 =
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                   &Func1::new({
                                                                                                                                                   let Apply0
                                                                                                                                                       =
                                                                                                                                                       Apply0.clone();
                                                                                                                                                   let Functor0
                                                                                                                                                       =
                                                                                                                                                       Functor0.clone();
                                                                                                                                                   let dictApplicative
                                                                                                                                                       =
                                                                                                                                                       dictApplicative.clone();
                                                                                                                                                   move
                                                                                                                                                       |f|
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableWithIndexList()),
                                                                                                                                                                                                                           &&&Func1::new({
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
                                                                                                                                                                                                                                                                     |a|
                                                                                                                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                                                                                                                     let a
                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                         a.clone();
                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                         |b|
                                                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&Apply0),
                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons()),
                                                                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&i),
                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&a))),
                                                                                                                                                                                                                                                                                                                          b)
                                                                                                                                                                                                                                                                                 })
                                                                                                                                                                                                                                                             })
                                                                                                                                                                                                                                         })),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                              &&&dictApplicative),
                                                                                                                                                                                                                           &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()))
                                                                                                                                               })
                                                                                                                               }),
                                                                                                              add(string("FunctorWithIndex0"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |usd__unused|
                                                                                                                                   &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorWithIndexList()),
                                                                                                                  add(string("FoldableWithIndex1"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |usd__unused_1|
                                                                                                                                       &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableWithIndexList()),
                                                                                                                      add(string("Traversable2"),
                                                                                                                          &&Func1::new(move
                                                                                                                                           |usd__unused_2|
                                                                                                                                           &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_traversableList()),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))))
    }
    pub fn Data_List_Lazy_Types_traversableWithIndexNonEmpty() -> &dyn Any {
        static Data_List_Lazy_Types_traversableWithIndexNonEmpty:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_traversableWithIndexNonEmpty.get_or_init(||
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_traversableWithIndexNonEmpty(),
                                                                                                           &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_traversableWithIndexList()))
    }
    pub fn Data_List_Lazy_Types_traversableWithIndexNonEmptyList()
     -> &dyn Any {
        static Data_List_Lazy_Types_traversableWithIndexNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_traversableWithIndexNonEmptyList.get_or_init(||
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                               &&&add(string("traverseWithIndex"),
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
                                                                                                                                                                                   |v|
                                                                                                                                                                                   {
                                                                                                                                                                                       let matchValue =
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                       let matchValue_1 =
                                                                                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                 &&&Functor0),
                                                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                |xxs|
                                                                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList()),
                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                      let xxs
                                                                                                                                                                                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                                                                                                                                                                                          xxs.clone();
                                                                                                                                                                                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                                                                                                                                                                                          |v1|
                                                                                                                                                                                                                                                                                                                                                                                                          &xxs
                                                                                                                                                                                                                                                                                                                                                                                                  })))))),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_traversableWithIndexNonEmpty()),
                                                                                                                                                                                                                                                                                                                                                                    &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&0_i32),
                                                                                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&1_i32))))),
                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                                                              &&&matchValue_1)))
                                                                                                                                                                                   }
                                                                                                                                                                           })
                                                                                                                                                       })
                                                                                                                                       }),
                                                                                                                      add(string("FunctorWithIndex0"),
                                                                                                                          &&Func1::new(move
                                                                                                                                           |usd__unused|
                                                                                                                                           &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorWithIndexNonEmptyList()),
                                                                                                                          add(string("FoldableWithIndex1"),
                                                                                                                              &&Func1::new(move
                                                                                                                                               |usd__unused_1|
                                                                                                                                               &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableWithIndexNonEmptyList()),
                                                                                                                              add(string("Traversable2"),
                                                                                                                                  &&Func1::new(move
                                                                                                                                                   |usd__unused_2|
                                                                                                                                                   &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_traversableNonEmptyList()),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>()))))))
    }
    pub fn Data_List_Lazy_Types_unfoldable1List() -> &dyn Any {
        static Data_List_Lazy_Types_unfoldable1List: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_unfoldable1List.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_Unfoldable1usd_Dict(),
                                                                                              &&&add(string("unfoldr1"),
                                                                                                     &{
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
                                                                                                                                                     |f|
                                                                                                                                                     Func1::new({
                                                                                                                                                                    let f
                                                                                                                                                                        =
                                                                                                                                                                        f.clone();
                                                                                                                                                                    let go_tco
                                                                                                                                                                        =
                                                                                                                                                                        go_tco.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |b|
                                                                                                                                                                        go_tco(f)(b.clone())
                                                                                                                                                                })
                                                                                                                                             })
                                                                                                                         });
                                                                                                          let go_1 =
                                                                                                              Lazy(go_2);
                                                                                                          fn go_tco(f_1:
                                                                                                                        &dyn Any)
                                                                                                           ->
                                                                                                               Func1<&dyn Any,
                                                                                                                     &dyn Any> {
                                                                                                              Func1::new({
                                                                                                                             let f_1
                                                                                                                                 =
                                                                                                                                 f_1.clone();
                                                                                                                             let go_tco
                                                                                                                                 =
                                                                                                                                 go_tco.clone();
                                                                                                                             move
                                                                                                                                 |b_1|
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_lazyList()),
                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                    let b_1
                                                                                                                                                                                        =
                                                                                                                                                                                        b_1.clone();
                                                                                                                                                                                    let go_tco
                                                                                                                                                                                        =
                                                                                                                                                                                        go_tco.clone();
                                                                                                                                                                                    move
                                                                                                                                                                                        |v|
                                                                                                                                                                                        {
                                                                                                                                                                                            let matchValue:
                                                                                                                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f_1,
                                                                                                                                                                                                                                                           &&&b_1));
                                                                                                                                                                                            {
                                                                                                                                                                                                let activePatternResult:
                                                                                                                                                                                                        LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                    Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                       });
                                                                                                                                                                                                if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(activePatternResult_1_0)
                                                                                                                                                                                                       =
                                                                                                                                                                                                       activePatternResult.as_ref()
                                                                                                                                                                                                   {
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                     &&go_tco(&f_1)(&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                         Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                                } else {
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
                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                         &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())
                                                                                                                                                                                                    } else {
                                                                                                                                                                                                        panic!("{}",
                                                                                                                                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.Types.fs"),
                                  Data1: 138_i32,
                                  Data2: 457_i32,}).get_Message(),)
                                                                                                                                                                                                    }
                                                                                                                                                                                                }
                                                                                                                                                                                            }
                                                                                                                                                                                        }
                                                                                                                                                                                }))
                                                                                                                         })
                                                                                                          }
                                                                                                          let go =
                                                                                                              go_1.Value;
                                                                                                          &go
                                                                                                      },
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_unfoldableList() -> &dyn Any {
        static Data_List_Lazy_Types_unfoldableList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_unfoldableList.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_Unfoldableusd_Dict(),
                                                                                             &&&add(string("unfoldr"),
                                                                                                    &{
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
                                                                                                                                                    |f|
                                                                                                                                                    Func1::new({
                                                                                                                                                                   let f
                                                                                                                                                                       =
                                                                                                                                                                       f.clone();
                                                                                                                                                                   let go_tco
                                                                                                                                                                       =
                                                                                                                                                                       go_tco.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |b|
                                                                                                                                                                       go_tco(f)(b.clone())
                                                                                                                                                               })
                                                                                                                                            })
                                                                                                                        });
                                                                                                         let go_1 =
                                                                                                             Lazy(go_2);
                                                                                                         fn go_tco(f_1:
                                                                                                                       &dyn Any)
                                                                                                          ->
                                                                                                              Func1<&dyn Any,
                                                                                                                    &dyn Any> {
                                                                                                             Func1::new({
                                                                                                                            let f_1
                                                                                                                                =
                                                                                                                                f_1.clone();
                                                                                                                            let go_tco
                                                                                                                                =
                                                                                                                                go_tco.clone();
                                                                                                                            move
                                                                                                                                |b_1|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_lazyList()),
                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                   let b_1
                                                                                                                                                                                       =
                                                                                                                                                                                       b_1.clone();
                                                                                                                                                                                   let go_tco
                                                                                                                                                                                       =
                                                                                                                                                                                       go_tco.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |v|
                                                                                                                                                                                       {
                                                                                                                                                                                           let matchValue:
                                                                                                                                                                                                   LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f_1,
                                                                                                                                                                                                                                                          &&&b_1));
                                                                                                                                                                                           match matchValue.as_ref()
                                                                                                                                                                                               {
                                                                                                                                                                                               Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                               =>
                                                                                                                                                                                               {
                                                                                                                                                                                                   let activePatternResult:
                                                                                                                                                                                                           LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                       Sharpurs_Prelude::_007cUnbox_007c(matchValue_1_0);
                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                       &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                          }),
                                                                                                                                                                                                                                    &&go_tco(&f_1)(&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                    }))
                                                                                                                                                                                               }
                                                                                                                                                                                               _
                                                                                                                                                                                               =>
                                                                                                                                                                                               &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                                                                                                           }
                                                                                                                                                                                       }
                                                                                                                                                                               }))
                                                                                                                        })
                                                                                                         }
                                                                                                         let go =
                                                                                                             go_1.Value;
                                                                                                         &go
                                                                                                     },
                                                                                                    add(string("Unfoldable10"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_unfoldable1List()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_unfoldable1NonEmpty() -> &dyn Any {
        static Data_List_Lazy_Types_unfoldable1NonEmpty:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_unfoldable1NonEmpty.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_unfoldable1NonEmpty(),
                                                                                                  &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_unfoldableList()))
    }
    pub fn Data_List_Lazy_Types_unfoldable1NonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_unfoldable1NonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_unfoldable1NonEmptyList.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_Unfoldable1usd_Dict(),
                                                                                                      &&&add(string("unfoldr1"),
                                                                                                             &&Func1::new(move
                                                                                                                              |f|
                                                                                                                              &Func1::new({
                                                                                                                                              let f
                                                                                                                                                  =
                                                                                                                                                  f.clone();
                                                                                                                                              move
                                                                                                                                                  |b|
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                      &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList()),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                        let b
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            b.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |v|
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_unfoldable1NonEmpty()),
                                                                                                                                                                                                                                                                                                                &&&f),
                                                                                                                                                                                                                                                                             &&&b)
                                                                                                                                                                                                                                    })))
                                                                                                                                          })),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_comonadNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_comonadNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_comonadNonEmptyList.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_Comonadusd_Dict(),
                                                                                                  &&&add(string("extract"),
                                                                                                         &&Func1::new(move
                                                                                                                          |v|
                                                                                                                          {
                                                                                                                              let nel =
                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                  &&&PureScript_Data_NonEmpty::Data_NonEmpty_head()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                  &&&nel))
                                                                                                                          }),
                                                                                                         add(string("Extend0"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused|
                                                                                                                              &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_extendNonEmptyList()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_monadList_0040158() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                         &&&add(string("Applicative0"),
                                                &&Func1::new({
                                                                 let Data_List_Lazy_Types_applicativeList_0040161_002d1
                                                                     =
                                                                     Data_List_Lazy_Types_applicativeList_0040161_002d1.clone();
                                                                 move
                                                                     |usd__unused|
                                                                     &Data_List_Lazy_Types_applicativeList_0040161_002d1.Value
                                                             }),
                                                add(string("Bind1"),
                                                    &&Func1::new({
                                                                     let Data_List_Lazy_Types_bindList_0040159_002d1
                                                                         =
                                                                         Data_List_Lazy_Types_bindList_0040159_002d1.clone();
                                                                     move
                                                                         |usd__unused_1|
                                                                         &Data_List_Lazy_Types_bindList_0040159_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_monadList_0040158_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_Types_monadList_0040158_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_Types_monadList_0040158_002d1.get_or_init(||
                                                                     Lazy(Data_List_Lazy_Types_monadList_0040158.clone()))
    }
    pub fn Data_List_Lazy_Types_bindList_0040159() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                         &&&add(string("bind"),
                                                &&Func1::new({
                                                                 let Data_List_Lazy_Types_bindList_0040159_002d1
                                                                     =
                                                                     Data_List_Lazy_Types_bindList_0040159_002d1.clone();
                                                                 move |xs|
                                                                     &Func1::new({
                                                                                     let Data_List_Lazy_Types_bindList_0040159_002d1
                                                                                         =
                                                                                         Data_List_Lazy_Types_bindList_0040159_002d1.clone();
                                                                                     let xs
                                                                                         =
                                                                                         xs.clone();
                                                                                     move
                                                                                         |f|
                                                                                         {
                                                                                             let go =
                                                                                                 &Func1::new({
                                                                                                                 let Data_List_Lazy_Types_bindList_0040159_002d1
                                                                                                                     =
                                                                                                                     Data_List_Lazy_Types_bindList_0040159_002d1.clone();
                                                                                                                 let f
                                                                                                                     =
                                                                                                                     f.clone();
                                                                                                                 move
                                                                                                                     |v|
                                                                                                                     {
                                                                                                                         let matchValue:
                                                                                                                                 LrcPtr<PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step> =
                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                         match matchValue.as_ref()
                                                                                                                             {
                                                                                                                             PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                           matchValue_1_1)
                                                                                                                             =>
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                       &&matchValue_1_0)),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                          &&&Data_List_Lazy_Types_bindList_0040159_002d1.Value),
                                                                                                                                                                                                                                                                       &&matchValue_1_1),
                                                                                                                                                                                                                                    &&&f))),
                                                                                                                             _
                                                                                                                             =>
                                                                                                                             &LrcPtr::new(PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor),
                                                                                                                         }
                                                                                                                     }
                                                                                                             });
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                       &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                    &&&go),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                    &&&xs)))
                                                                                         }
                                                                                 })
                                                             }),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let Data_List_Lazy_Types_applyList_0040160_002d1
                                                                         =
                                                                         Data_List_Lazy_Types_applyList_0040160_002d1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &Data_List_Lazy_Types_applyList_0040160_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_bindList_0040159_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_Types_bindList_0040159_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_Types_bindList_0040159_002d1.get_or_init(||
                                                                    Lazy(Data_List_Lazy_Types_bindList_0040159.clone()))
    }
    pub fn Data_List_Lazy_Types_applyList_0040160() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                         &&&add(string("apply"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_ap(),
                                                                                  &&&Data_List_Lazy_Types_monadList_0040158_002d1.Value),
                                                add(string("Functor0"),
                                                    &&Func1::new(move
                                                                     |usd__unused|
                                                                     &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorList()),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_applyList_0040160_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_Types_applyList_0040160_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_Types_applyList_0040160_002d1.get_or_init(||
                                                                     Lazy(Data_List_Lazy_Types_applyList_0040160.clone()))
    }
    pub fn Data_List_Lazy_Types_applicativeList_0040161() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                         &&&add(string("pure"),
                                                &&Func1::new(move |a|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                     a),
                                                                                                  &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let Data_List_Lazy_Types_applyList_0040160_002d1
                                                                         =
                                                                         Data_List_Lazy_Types_applyList_0040160_002d1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &Data_List_Lazy_Types_applyList_0040160_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Data_List_Lazy_Types_applicativeList_0040161_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_Types_applicativeList_0040161_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_Types_applicativeList_0040161_002d1.get_or_init(||
                                                                           Lazy(Data_List_Lazy_Types_applicativeList_0040161.clone()))
    }
    pub fn Data_List_Lazy_Types_monadList() -> &dyn Any {
        static Data_List_Lazy_Types_monadList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_monadList.get_or_init(||
                                                       Data_List_Lazy_Types_monadList_0040158_002d1.Value)
    }
    pub fn Data_List_Lazy_Types_bindList() -> &dyn Any {
        static Data_List_Lazy_Types_bindList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_bindList.get_or_init(||
                                                      Data_List_Lazy_Types_bindList_0040159_002d1.Value)
    }
    pub fn Data_List_Lazy_Types_applyList() -> &dyn Any {
        static Data_List_Lazy_Types_applyList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_applyList.get_or_init(||
                                                       Data_List_Lazy_Types_applyList_0040160_002d1.Value)
    }
    pub fn Data_List_Lazy_Types_applicativeList() -> &dyn Any {
        static Data_List_Lazy_Types_applicativeList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_applicativeList.get_or_init(||
                                                             Data_List_Lazy_Types_applicativeList_0040161_002d1.Value)
    }
    pub fn Data_List_Lazy_Types_applyNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_applyNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_applyNonEmptyList.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                &&&add(string("apply"),
                                                                                                       &&Func1::new(move
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
                                                                                                                                                let v2 =
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                     &&&matchValue_1);
                                                                                                                                                let matchValue_3:
                                                                                                                                                        LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                               &&&matchValue));
                                                                                                                                                let matchValue_4:
                                                                                                                                                        LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                                    Sharpurs_Prelude::unbox(&&v2);
                                                                                                                                                let fs =
                                                                                                                                                    match matchValue_3.as_ref()
                                                                                                                                                        {
                                                                                                                                                        Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                               x)
                                                                                                                                                        =>
                                                                                                                                                        x.clone(),
                                                                                                                                                    };
                                                                                                                                                let f =
                                                                                                                                                    match matchValue_3.as_ref()
                                                                                                                                                        {
                                                                                                                                                        Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                               _)
                                                                                                                                                        =>
                                                                                                                                                        x.clone(),
                                                                                                                                                    };
                                                                                                                                                let a =
                                                                                                                                                    match matchValue_4.as_ref()
                                                                                                                                                        {
                                                                                                                                                        Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                               _)
                                                                                                                                                        =>
                                                                                                                                                        x.clone(),
                                                                                                                                                    };
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                                                      let a
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          a.clone();
                                                                                                                                                                                                                                      let f
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          f.clone();
                                                                                                                                                                                                                                      let fs
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          fs.clone();
                                                                                                                                                                                                                                      let matchValue_4
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          matchValue_4.clone();
                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                          |v4|
                                                                                                                                                                                                                                          &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                               &&&a),
                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_applyList()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&fs),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()))),
                                                                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_applyList()),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&fs)),
                                                                                                                                                                                                                                                                                                                                                                                  &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                                                         Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                     }))))
                                                                                                                                                                                                                                  })))
                                                                                                                                            }
                                                                                                                                    })),
                                                                                                       add(string("Functor0"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused|
                                                                                                                            &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorNonEmptyList()),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_bindNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_bindNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_bindNonEmptyList.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                               &&&add(string("bind"),
                                                                                                      &&Func1::new(move
                                                                                                                       |v|
                                                                                                                       &Func1::new({
                                                                                                                                       let v
                                                                                                                                           =
                                                                                                                                           v.clone();
                                                                                                                                       move
                                                                                                                                           |f|
                                                                                                                                           {
                                                                                                                                               let matchValue =
                                                                                                                                                   Sharpurs_Prelude::unbox(&&v);
                                                                                                                                               let f1 =
                                                                                                                                                   Sharpurs_Prelude::unbox(f);
                                                                                                                                               let matchValue_3:
                                                                                                                                                       LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                              &&&matchValue));
                                                                                                                                               let matchValue_4:
                                                                                                                                                       LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                 &&&PureScript_Data_Lazy::Data_Lazy_force()),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_unwrap()),
                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                    &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                           Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                       }))));
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                                                     let f1
                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                         f1.clone();
                                                                                                                                                                                                                                     let matchValue_3
                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                         matchValue_3.clone();
                                                                                                                                                                                                                                     let matchValue_4
                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                         matchValue_4.clone();
                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                         |v3|
                                                                                                                                                                                                                                         &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                  Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                              },
                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                                                                                                                                                                                                                                                                                                                 &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                                        Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_bindList()),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                                                                                                           Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_toList()),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&f1)))))
                                                                                                                                                                                                                                 })))
                                                                                                                                           }
                                                                                                                                   })),
                                                                                                      add(string("Apply0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_applyNonEmptyList()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_altNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_altNonEmptyList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_altNonEmptyList.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                              &&&add(string("alt"),
                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupNonEmptyList()),
                                                                                                     add(string("Functor0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorNonEmptyList()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_altList() -> &dyn Any {
        static Data_List_Lazy_Types_altList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_altList.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                      &&&add(string("alt"),
                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                               &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                             add(string("Functor0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorList()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_plusList() -> &dyn Any {
        static Data_List_Lazy_Types_plusList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_plusList.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                       &&&add(string("empty"),
                                                                                              &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                              add(string("Alt0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_altList()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_alternativeList() -> &dyn Any {
        static Data_List_Lazy_Types_alternativeList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_Types_alternativeList.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                              &&&add(string("Applicative0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_applicativeList()),
                                                                                                     add(string("Plus1"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused_1|
                                                                                                                          &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_plusList()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_monadPlusList() -> &dyn Any {
        static Data_List_Lazy_Types_monadPlusList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_monadPlusList.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_MonadPlus::Control_MonadPlus_MonadPlususd_Dict(),
                                                                                            &&&add(string("Monad0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_monadList()),
                                                                                                   add(string("Alternative1"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused_1|
                                                                                                                        &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_alternativeList()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_applicativeNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_applicativeNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_applicativeNonEmptyList.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                      &&&add(string("pure"),
                                                                                                             &&Func1::new(move
                                                                                                                              |a|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let a
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        a.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |v|
                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_singleton(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_plusList()),
                                                                                                                                                                                                                                                         &&&a)
                                                                                                                                                                                                                })))),
                                                                                                             add(string("Apply0"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |usd__unused|
                                                                                                                                  &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_applyNonEmptyList()),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>()))))
    }
    pub fn Data_List_Lazy_Types_monadNonEmptyList() -> &dyn Any {
        static Data_List_Lazy_Types_monadNonEmptyList:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Types_monadNonEmptyList.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                                &&&add(string("Applicative0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_applicativeNonEmptyList()),
                                                                                                       add(string("Bind1"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused_1|
                                                                                                                            &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_bindNonEmptyList()),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>()))))
    }
}
