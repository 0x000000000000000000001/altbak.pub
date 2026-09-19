pub mod PureScript_Data_Semigroup_Foldable {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5023b7e9::PureScript_Data_Monoid_Dual;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_a3a23b56::PureScript_Data_Ord_Max;
    use crate::module_a424d288::PureScript_Data_Ord_Min;
    use crate::module_5f769efb::PureScript_Data_Ordering;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    #[derive(Clone, Debug,)]
    pub enum Data_Semigroup_Foldable_FoldRight1 {
        Data_Semigroup_Foldable_FoldRight1usd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1 {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Semigroup_Foldable_identity() -> &dyn Any {
        static Data_Semigroup_Foldable_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_identity.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                          &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Semigroup_Foldable_identity1() -> &dyn Any {
        static Data_Semigroup_Foldable_identity1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_identity1.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                           &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Semigroup_Foldable_JoinWith() -> &dyn Any {
        static Data_Semigroup_Foldable_JoinWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_JoinWith.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Data_Semigroup_Foldable_Foldable1usd_Dict() -> &dyn Any {
        static Data_Semigroup_Foldable_Foldable1usd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_Foldable1usd_Dict.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |x|
                                                                                  x.clone()))
    }
    pub fn Data_Semigroup_Foldable_FoldRight1() -> &dyn Any {
        static Data_Semigroup_Foldable_FoldRight1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_FoldRight1.get_or_init(||
                                                           &Func1::new(move
                                                                           |usd__arg1|
                                                                           Func1::new({
                                                                                          let usd__arg1
                                                                                              =
                                                                                              usd__arg1.clone();
                                                                                          move
                                                                                              |usd__arg2|
                                                                                              &LrcPtr::new(PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1::Data_Semigroup_Foldable_FoldRight1usd_Ctor(usd__arg1,
                                                                                                                                                                                                                              usd__arg2.clone()))
                                                                                      })))
    }
    pub fn Data_Semigroup_Foldable_Act() -> &dyn Any {
        static Data_Semigroup_Foldable_Act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_Act.get_or_init(||
                                                    &Func1::new(move |x|
                                                                    x.clone()))
    }
    pub fn Data_Semigroup_Foldable_semigroupJoinWith() -> &dyn Any {
        static Data_Semigroup_Foldable_semigroupJoinWith:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_semigroupJoinWith.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictSemigroup|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                   &&&add(string("append"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictSemigroup
                                                                                                                                               =
                                                                                                                                               dictSemigroup.clone();
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
                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_JoinWith()),
                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                          let matchValue_1
                                                                                                                                                                                                                              =
                                                                                                                                                                                                                              matchValue_1.clone();
                                                                                                                                                                                                                          move
                                                                                                                                                                                                                              |j|
                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                     &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                     j)),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                        &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                                                     j),
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                     j)))
                                                                                                                                                                                                                      }))
                                                                                                                                                                   }
                                                                                                                                                           })
                                                                                                                                       }),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Semigroup_Foldable_semigroupAct() -> &dyn Any {
        static Data_Semigroup_Foldable_semigroupAct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_Foldable_semigroupAct.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictApply|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                              &&&add(string("append"),
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
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_Act(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_applySecond(),
                                                                                                                                                                                                                                                                                                            &&&dictApply),
                                                                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                                                                      &&&matchValue_1))
                                                                                                                                                              }
                                                                                                                                                      })
                                                                                                                                  }),
                                                                                                                     empty::<string,
                                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Semigroup_Foldable_runFoldRight1() -> &dyn Any {
        static Data_Semigroup_Foldable_runFoldRight1:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_runFoldRight1.get_or_init(||
                                                              &Func1::new(move
                                                                              |v|
                                                                              {
                                                                                  let matchValue:
                                                                                          LrcPtr<PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1> =
                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&match matchValue.as_ref()
                                                                                                                          {
                                                                                                                          PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1::Data_Semigroup_Foldable_FoldRight1usd_Ctor(x,
                                                                                                                                                                                                                                             _)
                                                                                                                          =>
                                                                                                                          x.clone(),
                                                                                                                      },
                                                                                                                   &&&match matchValue.as_ref()
                                                                                                                          {
                                                                                                                          PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1::Data_Semigroup_Foldable_FoldRight1usd_Ctor(_,
                                                                                                                                                                                                                                             x)
                                                                                                                          =>
                                                                                                                          x.clone(),
                                                                                                                      })
                                                                              }))
    }
    pub fn Data_Semigroup_Foldable_mkFoldRight1() -> &dyn Any {
        static Data_Semigroup_Foldable_mkFoldRight1: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_Foldable_mkFoldRight1.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                |usd__arg1|
                                                                                                                Func1::new({
                                                                                                                               let usd__arg1
                                                                                                                                   =
                                                                                                                                   usd__arg1.clone();
                                                                                                                               move
                                                                                                                                   |usd__arg2|
                                                                                                                                   &LrcPtr::new(PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1::Data_Semigroup_Foldable_FoldRight1usd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                   usd__arg2.clone()))
                                                                                                                           })),
                                                                                              &&&PureScript_Data_Function::Data_Function_const()))
    }
    pub fn Data_Semigroup_Foldable_joinee() -> &dyn Any {
        static Data_Semigroup_Foldable_joinee: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_joinee.get_or_init(||
                                                       &Func1::new(move |v|
                                                                       &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_Semigroup_Foldable_getAct() -> &dyn Any {
        static Data_Semigroup_Foldable_getAct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_getAct.get_or_init(||
                                                       &Func1::new(move |v|
                                                                       &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_Semigroup_Foldable_foldr1() -> &dyn Any {
        static Data_Semigroup_Foldable_foldr1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldr1.get_or_init(||
                                                       &Func1::new(move |dict|
                                                                       find(string("foldr1"),
                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semigroup_Foldable_foldl1() -> &dyn Any {
        static Data_Semigroup_Foldable_foldl1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldl1.get_or_init(||
                                                       &Func1::new(move |dict|
                                                                       find(string("foldl1"),
                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semigroup_Foldable_maximumBy() -> &dyn Any {
        static Data_Semigroup_Foldable_maximumBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_maximumBy.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictFoldable1|
                                                                          &Func1::new({
                                                                                          let dictFoldable1
                                                                                              =
                                                                                              dictFoldable1.clone();
                                                                                          move
                                                                                              |cmp|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldl1(),
                                                                                                                                                                  &&&dictFoldable1),
                                                                                                                               &&&Func1::new({
                                                                                                                                                 let cmp
                                                                                                                                                     =
                                                                                                                                                     cmp.clone();
                                                                                                                                                 move
                                                                                                                                                     |x|
                                                                                                                                                     &Func1::new({
                                                                                                                                                                     let x
                                                                                                                                                                         =
                                                                                                                                                                         x.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |y|
                                                                                                                                                                         {
                                                                                                                                                                             let matchValue =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                                                                                    &&&x),
                                                                                                                                                                                                                                                                                                                 y)),
                                                                                                                                                                                                                                           &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)));
                                                                                                                                                                             match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                              &matchValue)
                                                                                                                                                                                 {
                                                                                                                                                                                 0_i32
                                                                                                                                                                                 =>
                                                                                                                                                                                 &x,
                                                                                                                                                                                 _
                                                                                                                                                                                 =>
                                                                                                                                                                                 y.clone(),
                                                                                                                                                                             }
                                                                                                                                                                         }
                                                                                                                                                                 })
                                                                                                                                             }))
                                                                                      })))
    }
    pub fn Data_Semigroup_Foldable_minimumBy() -> &dyn Any {
        static Data_Semigroup_Foldable_minimumBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_minimumBy.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictFoldable1|
                                                                          &Func1::new({
                                                                                          let dictFoldable1
                                                                                              =
                                                                                              dictFoldable1.clone();
                                                                                          move
                                                                                              |cmp|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldl1(),
                                                                                                                                                                  &&&dictFoldable1),
                                                                                                                               &&&Func1::new({
                                                                                                                                                 let cmp
                                                                                                                                                     =
                                                                                                                                                     cmp.clone();
                                                                                                                                                 move
                                                                                                                                                     |x|
                                                                                                                                                     &Func1::new({
                                                                                                                                                                     let x
                                                                                                                                                                         =
                                                                                                                                                                         x.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |y|
                                                                                                                                                                         {
                                                                                                                                                                             let matchValue =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                                                                                    &&&x),
                                                                                                                                                                                                                                                                                                                 y)),
                                                                                                                                                                                                                                           &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)));
                                                                                                                                                                             match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                              &matchValue)
                                                                                                                                                                                 {
                                                                                                                                                                                 0_i32
                                                                                                                                                                                 =>
                                                                                                                                                                                 &x,
                                                                                                                                                                                 _
                                                                                                                                                                                 =>
                                                                                                                                                                                 y.clone(),
                                                                                                                                                                             }
                                                                                                                                                                         }
                                                                                                                                                                 })
                                                                                                                                             }))
                                                                                      })))
    }
    pub fn Data_Semigroup_Foldable_foldableTuple() -> &dyn Any {
        static Data_Semigroup_Foldable_foldableTuple:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldableTuple.get_or_init(||
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
                                                                                                                                                           |v|
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f),
                                                                                                                                                                                            &&&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                                                   {
                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                               })
                                                                                                                                                   }))),
                                                                                                      add(string("foldr1"),
                                                                                                          &&Func1::new(move
                                                                                                                           |v_1|
                                                                                                                           &Func1::new({
                                                                                                                                           let v_1
                                                                                                                                               =
                                                                                                                                               v_1.clone();
                                                                                                                                           move
                                                                                                                                               |v1|
                                                                                                                                               {
                                                                                                                                                   let matchValue_3 =
                                                                                                                                                       Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                   &match Sharpurs_Prelude::unbox(v1).as_ref()
                                                                                                                                                        {
                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                   x)
                                                                                                                                                        =>
                                                                                                                                                        x.clone(),
                                                                                                                                                    }
                                                                                                                                               }
                                                                                                                                       })),
                                                                                                          add(string("foldl1"),
                                                                                                              &&Func1::new(move
                                                                                                                               |v_2|
                                                                                                                               &Func1::new({
                                                                                                                                               let v_2
                                                                                                                                                   =
                                                                                                                                                   v_2.clone();
                                                                                                                                               move
                                                                                                                                                   |v1_1|
                                                                                                                                                   {
                                                                                                                                                       let matchValue_6 =
                                                                                                                                                           Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                       &match Sharpurs_Prelude::unbox(v1_1).as_ref()
                                                                                                                                                            {
                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                       x)
                                                                                                                                                            =>
                                                                                                                                                            x.clone(),
                                                                                                                                                        }
                                                                                                                                                   }
                                                                                                                                           })),
                                                                                                              add(string("Foldable0"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |usd__unused|
                                                                                                                                   &PureScript_Data_Foldable::Data_Foldable_foldableTuple()),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))))
    }
    pub fn Data_Semigroup_Foldable_foldableMultiplicative() -> &dyn Any {
        static Data_Semigroup_Foldable_foldableMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldableMultiplicative.get_or_init(||
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_Foldable1usd_Dict(),
                                                                                                        &&&add(string("foldr1"),
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
                                                                                                                                                        &Sharpurs_Prelude::unbox(v1)
                                                                                                                                                    }
                                                                                                                                            })),
                                                                                                               add(string("foldl1"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |v_1|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let v_1
                                                                                                                                                        =
                                                                                                                                                        v_1.clone();
                                                                                                                                                    move
                                                                                                                                                        |v1_1|
                                                                                                                                                        {
                                                                                                                                                            let matchValue_3 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                            &Sharpurs_Prelude::unbox(v1_1)
                                                                                                                                                        }
                                                                                                                                                })),
                                                                                                                   add(string("foldMap1"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |dictSemigroup|
                                                                                                                                        &Func1::new(move
                                                                                                                                                        |f|
                                                                                                                                                        &Func1::new({
                                                                                                                                                                        let f
                                                                                                                                                                            =
                                                                                                                                                                            f.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |v_2|
                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f),
                                                                                                                                                                                                             &&&Sharpurs_Prelude::unbox(v_2))
                                                                                                                                                                    }))),
                                                                                                                       add(string("Foldable0"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |usd__unused|
                                                                                                                                            &PureScript_Data_Foldable::Data_Foldable_foldableMultiplicative()),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))))))
    }
    pub fn Data_Semigroup_Foldable_foldableIdentity() -> &dyn Any {
        static Data_Semigroup_Foldable_foldableIdentity:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldableIdentity.get_or_init(||
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
                                                                                                                                                              |v|
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f),
                                                                                                                                                                                               &&&Sharpurs_Prelude::unbox(v))
                                                                                                                                                      }))),
                                                                                                         add(string("foldl1"),
                                                                                                             &&Func1::new(move
                                                                                                                              |v_1|
                                                                                                                              &Func1::new({
                                                                                                                                              let v_1
                                                                                                                                                  =
                                                                                                                                                  v_1.clone();
                                                                                                                                              move
                                                                                                                                                  |v1|
                                                                                                                                                  {
                                                                                                                                                      let matchValue_3 =
                                                                                                                                                          Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                      &Sharpurs_Prelude::unbox(v1)
                                                                                                                                                  }
                                                                                                                                          })),
                                                                                                             add(string("foldr1"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |v_2|
                                                                                                                                  &Func1::new({
                                                                                                                                                  let v_2
                                                                                                                                                      =
                                                                                                                                                      v_2.clone();
                                                                                                                                                  move
                                                                                                                                                      |v1_1|
                                                                                                                                                      {
                                                                                                                                                          let matchValue_6 =
                                                                                                                                                              Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                          &Sharpurs_Prelude::unbox(v1_1)
                                                                                                                                                      }
                                                                                                                                              })),
                                                                                                                 add(string("Foldable0"),
                                                                                                                     &&Func1::new(move
                                                                                                                                      |usd__unused|
                                                                                                                                      &PureScript_Data_Foldable::Data_Foldable_foldableIdentity()),
                                                                                                                     empty::<string,
                                                                                                                             &dyn Any>()))))))
    }
    pub fn Data_Semigroup_Foldable_foldableDual() -> &dyn Any {
        static Data_Semigroup_Foldable_foldableDual: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldableDual.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_Foldable1usd_Dict(),
                                                                                              &&&add(string("foldr1"),
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
                                                                                                                                              &Sharpurs_Prelude::unbox(v1)
                                                                                                                                          }
                                                                                                                                  })),
                                                                                                     add(string("foldl1"),
                                                                                                         &&Func1::new(move
                                                                                                                          |v_1|
                                                                                                                          &Func1::new({
                                                                                                                                          let v_1
                                                                                                                                              =
                                                                                                                                              v_1.clone();
                                                                                                                                          move
                                                                                                                                              |v1_1|
                                                                                                                                              {
                                                                                                                                                  let matchValue_3 =
                                                                                                                                                      Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                  &Sharpurs_Prelude::unbox(v1_1)
                                                                                                                                              }
                                                                                                                                      })),
                                                                                                         add(string("foldMap1"),
                                                                                                             &&Func1::new(move
                                                                                                                              |dictSemigroup|
                                                                                                                              &Func1::new(move
                                                                                                                                              |f|
                                                                                                                                              &Func1::new({
                                                                                                                                                              let f
                                                                                                                                                                  =
                                                                                                                                                                  f.clone();
                                                                                                                                                              move
                                                                                                                                                                  |v_2|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f),
                                                                                                                                                                                                   &&&Sharpurs_Prelude::unbox(v_2))
                                                                                                                                                          }))),
                                                                                                             add(string("Foldable0"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |usd__unused|
                                                                                                                                  &PureScript_Data_Foldable::Data_Foldable_foldableDual()),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>()))))))
    }
    pub fn Data_Semigroup_Foldable_foldRight1Semigroup() -> &dyn Any {
        static Data_Semigroup_Foldable_foldRight1Semigroup:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldRight1Semigroup.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                     &&&add(string("append"),
                                                                                                            &&Func1::new(move
                                                                                                                             |v|
                                                                                                                             &Func1::new({
                                                                                                                                             let v
                                                                                                                                                 =
                                                                                                                                                 v.clone();
                                                                                                                                             move
                                                                                                                                                 |v1|
                                                                                                                                                 {
                                                                                                                                                     let matchValue:
                                                                                                                                                             LrcPtr<PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1> =
                                                                                                                                                         Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                     let matchValue_1:
                                                                                                                                                             LrcPtr<PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1> =
                                                                                                                                                         Sharpurs_Prelude::unbox(v1);
                                                                                                                                                     &LrcPtr::new(PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1::Data_Semigroup_Foldable_FoldRight1usd_Ctor(&Func1::new({
                                                                                                                                                                                                                                                                                                     let matchValue_1
                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                         matchValue_1.clone();
                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                         |a|
                                                                                                                                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                                                                                                                                         let a
                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                             a.clone();
                                                                                                                                                                                                                                                                                                                         let matchValue_1
                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                             matchValue_1.clone();
                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                             |f|
                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                                                        PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1::Data_Semigroup_Foldable_FoldRight1usd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                    },
                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1::Data_Semigroup_Foldable_FoldRight1usd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          }),
                                                                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1::Data_Semigroup_Foldable_FoldRight1usd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             },
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                       f))),
                                                                                                                                                                                                                                                                                                                                                              f)
                                                                                                                                                                                                                                                                                                                     })
                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                                                     &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                          PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_FoldRight1::Data_Semigroup_Foldable_FoldRight1usd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                             x)
                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                      }))
                                                                                                                                                 }
                                                                                                                                         })),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>())))
    }
    pub fn Data_Semigroup_Foldable_semigroupDual() -> &dyn Any {
        static Data_Semigroup_Foldable_semigroupDual:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_semigroupDual.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_semigroupDual(),
                                                                                               &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldRight1Semigroup()))
    }
    pub fn Data_Semigroup_Foldable_foldMap1DefaultR() -> &dyn Any {
        static Data_Semigroup_Foldable_foldMap1DefaultR:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldMap1DefaultR.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictFoldable1|
                                                                                 &Func1::new({
                                                                                                 let dictFoldable1
                                                                                                     =
                                                                                                     dictFoldable1.clone();
                                                                                                 move
                                                                                                     |dictFunctor|
                                                                                                     &Func1::new({
                                                                                                                     let dictFunctor
                                                                                                                         =
                                                                                                                         dictFunctor.clone();
                                                                                                                     move
                                                                                                                         |dictSemigroup|
                                                                                                                         {
                                                                                                                             let append =
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                  dictSemigroup);
                                                                                                                             &Func1::new({
                                                                                                                                             let append
                                                                                                                                                 =
                                                                                                                                                 append.clone();
                                                                                                                                             move
                                                                                                                                                 |f|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                           &&&dictFunctor),
                                                                                                                                                                                                                                                        f)),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldr1(),
                                                                                                                                                                                                                                                        &&&dictFoldable1),
                                                                                                                                                                                                                     &&&append))
                                                                                                                                         })
                                                                                                                         }
                                                                                                                 })
                                                                                             })))
    }
    pub fn Data_Semigroup_Foldable_foldMap1DefaultL() -> &dyn Any {
        static Data_Semigroup_Foldable_foldMap1DefaultL:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldMap1DefaultL.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictFoldable1|
                                                                                 &Func1::new({
                                                                                                 let dictFoldable1
                                                                                                     =
                                                                                                     dictFoldable1.clone();
                                                                                                 move
                                                                                                     |dictFunctor|
                                                                                                     &Func1::new({
                                                                                                                     let dictFunctor
                                                                                                                         =
                                                                                                                         dictFunctor.clone();
                                                                                                                     move
                                                                                                                         |dictSemigroup|
                                                                                                                         {
                                                                                                                             let append =
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                  dictSemigroup);
                                                                                                                             &Func1::new({
                                                                                                                                             let append
                                                                                                                                                 =
                                                                                                                                                 append.clone();
                                                                                                                                             move
                                                                                                                                                 |f|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                           &&&dictFunctor),
                                                                                                                                                                                                                                                        f)),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldl1(),
                                                                                                                                                                                                                                                        &&&dictFoldable1),
                                                                                                                                                                                                                     &&&append))
                                                                                                                                         })
                                                                                                                         }
                                                                                                                 })
                                                                                             })))
    }
    pub fn Data_Semigroup_Foldable_foldMap1() -> &dyn Any {
        static Data_Semigroup_Foldable_foldMap1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldMap1.get_or_init(||
                                                         &Func1::new(move
                                                                         |dict|
                                                                         find(string("foldMap1"),
                                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semigroup_Foldable_foldl1Default() -> &dyn Any {
        static Data_Semigroup_Foldable_foldl1Default:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldl1Default.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictFoldable1|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_runFoldRight1()),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_alaF(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual()),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                                                                                                                                                                                    dictFoldable1),
                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_semigroupDual())),
                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_mkFoldRight1())))),
                                                                                                               &&&PureScript_Data_Function::Data_Function_flip())))
    }
    pub fn Data_Semigroup_Foldable_foldr1Default() -> &dyn Any {
        static Data_Semigroup_Foldable_foldr1Default:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_foldr1Default.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictFoldable1|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                     &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_runFoldRight1()),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                                                                           dictFoldable1),
                                                                                                                                                                                                                        &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldRight1Semigroup()),
                                                                                                                                                                                     &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_mkFoldRight1())))))
    }
    pub fn Data_Semigroup_Foldable_intercalateMap() -> &dyn Any {
        static Data_Semigroup_Foldable_intercalateMap:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_intercalateMap.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictFoldable1|
                                                                               &Func1::new({
                                                                                               let dictFoldable1
                                                                                                   =
                                                                                                   dictFoldable1.clone();
                                                                                               move
                                                                                                   |dictSemigroup|
                                                                                                   {
                                                                                                       let semigroupJoinWith1 =
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_semigroupJoinWith(),
                                                                                                                                            dictSemigroup);
                                                                                                       &Func1::new({
                                                                                                                       let semigroupJoinWith1
                                                                                                                           =
                                                                                                                           semigroupJoinWith1.clone();
                                                                                                                       move
                                                                                                                           |j|
                                                                                                                           &Func1::new({
                                                                                                                                           let j
                                                                                                                                               =
                                                                                                                                               j.clone();
                                                                                                                                           move
                                                                                                                                               |f|
                                                                                                                                               &Func1::new({
                                                                                                                                                               let f
                                                                                                                                                                   =
                                                                                                                                                                   f.clone();
                                                                                                                                                               move
                                                                                                                                                                   |foldable|
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_joinee(),
                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                                                                                                                                                                                                   &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                                                                &&&semigroupJoinWith1),
                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_JoinWith()),
                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Function::Data_Function_const()),
                                                                                                                                                                                                                                                                                                                                                                                   &&&f))),
                                                                                                                                                                                                                                                                          foldable)),
                                                                                                                                                                                                    &&&j)
                                                                                                                                                           })
                                                                                                                                       })
                                                                                                                   })
                                                                                                   }
                                                                                           })))
    }
    pub fn Data_Semigroup_Foldable_intercalate() -> &dyn Any {
        static Data_Semigroup_Foldable_intercalate: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_Foldable_intercalate.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFoldable1|
                                                                            {
                                                                                let intercalateMap1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_intercalateMap(),
                                                                                                                     dictFoldable1);
                                                                                &Func1::new({
                                                                                                let intercalateMap1
                                                                                                    =
                                                                                                    intercalateMap1.clone();
                                                                                                move
                                                                                                    |dictSemigroup|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&intercalateMap1,
                                                                                                                                                                                                           dictSemigroup)),
                                                                                                                                     &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_identity())
                                                                                            })
                                                                            }))
    }
    pub fn Data_Semigroup_Foldable_maximum() -> &dyn Any {
        static Data_Semigroup_Foldable_maximum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_maximum.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictOrd|
                                                                        {
                                                                            let semigroupMax =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Max::Data_Ord_Max_semigroupMax(),
                                                                                                                 dictOrd);
                                                                            &Func1::new({
                                                                                            let semigroupMax
                                                                                                =
                                                                                                semigroupMax.clone();
                                                                                            move
                                                                                                |dictFoldable1|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_ala(),
                                                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                    &&&PureScript_Data_Ord_Max::Data_Ord_Max_Max()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                       dictFoldable1),
                                                                                                                                                                    &&&semigroupMax))
                                                                                        })
                                                                        }))
    }
    pub fn Data_Semigroup_Foldable_minimum() -> &dyn Any {
        static Data_Semigroup_Foldable_minimum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_minimum.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictOrd|
                                                                        {
                                                                            let semigroupMin =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Min::Data_Ord_Min_semigroupMin(),
                                                                                                                 dictOrd);
                                                                            &Func1::new({
                                                                                            let semigroupMin
                                                                                                =
                                                                                                semigroupMin.clone();
                                                                                            move
                                                                                                |dictFoldable1|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_ala(),
                                                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                    &&&PureScript_Data_Ord_Min::Data_Ord_Min_Min()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                       dictFoldable1),
                                                                                                                                                                    &&&semigroupMin))
                                                                                        })
                                                                        }))
    }
    pub fn Data_Semigroup_Foldable_traverse1_() -> &dyn Any {
        static Data_Semigroup_Foldable_traverse1_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_traverse1_.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictFoldable1|
                                                                           &Func1::new({
                                                                                           let dictFoldable1
                                                                                               =
                                                                                               dictFoldable1.clone();
                                                                                           move
                                                                                               |dictApply|
                                                                                               {
                                                                                                   let Functor0 =
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                               Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                   let semigroupAct1 =
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_semigroupAct(),
                                                                                                                                        dictApply);
                                                                                                   &Func1::new({
                                                                                                                   let Functor0
                                                                                                                       =
                                                                                                                       Functor0.clone();
                                                                                                                   let semigroupAct1
                                                                                                                       =
                                                                                                                       semigroupAct1.clone();
                                                                                                                   move
                                                                                                                       |f|
                                                                                                                       &Func1::new({
                                                                                                                                       let f
                                                                                                                                           =
                                                                                                                                           f.clone();
                                                                                                                                       move
                                                                                                                                           |t|
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_voidRight(),
                                                                                                                                                                                                                                                  &&&Functor0),
                                                                                                                                                                                                               &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_getAct(),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                                                                                                                                                                           &&&dictFoldable1),
                                                                                                                                                                                                                                                                                                                        &&&semigroupAct1),
                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_Act()),
                                                                                                                                                                                                                                                                                                                        &&&f)),
                                                                                                                                                                                                                                                  t)))
                                                                                                                                   })
                                                                                                               })
                                                                                               }
                                                                                       })))
    }
    pub fn Data_Semigroup_Foldable_for1_() -> &dyn Any {
        static Data_Semigroup_Foldable_for1_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_for1_.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictFoldable1|
                                                                      {
                                                                          let traverse1_1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_traverse1_(),
                                                                                                               dictFoldable1);
                                                                          &Func1::new({
                                                                                          let traverse1_1
                                                                                              =
                                                                                              traverse1_1.clone();
                                                                                          move
                                                                                              |dictApply|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&traverse1_1,
                                                                                                                                                                  dictApply))
                                                                                      })
                                                                      }))
    }
    pub fn Data_Semigroup_Foldable_sequence1_() -> &dyn Any {
        static Data_Semigroup_Foldable_sequence1_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_sequence1_.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictFoldable1|
                                                                           &Func1::new({
                                                                                           let dictFoldable1
                                                                                               =
                                                                                               dictFoldable1.clone();
                                                                                           move
                                                                                               |dictApply|
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_traverse1_(),
                                                                                                                                                                                                      &&&dictFoldable1),
                                                                                                                                                                   dictApply),
                                                                                                                                &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_identity1())
                                                                                       })))
    }
    pub fn Data_Semigroup_Foldable_fold1() -> &dyn Any {
        static Data_Semigroup_Foldable_fold1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Foldable_fold1.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictFoldable1|
                                                                      &Func1::new({
                                                                                      let dictFoldable1
                                                                                          =
                                                                                          dictFoldable1.clone();
                                                                                      move
                                                                                          |dictSemigroup|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                 &&&dictFoldable1),
                                                                                                                                                              dictSemigroup),
                                                                                                                           &&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_identity())
                                                                                  })))
    }
}
