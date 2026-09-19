pub mod PureScript_Data_Decide {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_18b8108c::PureScript_Data_Comparison;
    use crate::module_d8620aa6::PureScript_Data_Divide;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_59dcd94b::PureScript_Data_Equivalence;
    use crate::module_851c93ca::PureScript_Data_Op;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_d23c04ec::PureScript_Data_Predicate;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Decide_identity() -> &dyn Any {
        static Data_Decide_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decide_identity.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                              &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Decide_Decideusd_Dict() -> &dyn Any {
        static Data_Decide_Decideusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decide_Decideusd_Dict.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Decide_choosePredicate() -> &dyn Any {
        static Data_Decide_choosePredicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decide_choosePredicate.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decide::Data_Decide_Decideusd_Dict(),
                                                                                     &&&add(string("choose"),
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
                                                                                                                                                     |v1|
                                                                                                                                                     {
                                                                                                                                                         let matchValue =
                                                                                                                                                             Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                         let matchValue_1 =
                                                                                                                                                             Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                         let matchValue_2 =
                                                                                                                                                             Sharpurs_Prelude::unbox(v1);
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Predicate::Data_Predicate_Predicate(),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                                                                                                                                                                                                      &&&matchValue_1),
                                                                                                                                                                                                                                                                                                   &&&matchValue_2)),
                                                                                                                                                                                                                             &&&matchValue))
                                                                                                                                                     }
                                                                                                                                             })
                                                                                                                         })),
                                                                                            add(string("Divide0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Data_Divide::Data_Divide_dividePredicate()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Decide_chooseOp() -> &dyn Any {
        static Data_Decide_chooseOp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decide_chooseOp.get_or_init(||
                                             &Func1::new(move |dictSemigroup|
                                                             {
                                                                 let divideOp =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Divide::Data_Divide_divideOp(),
                                                                                                      dictSemigroup);
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decide::Data_Decide_Decideusd_Dict(),
                                                                                                  &&&add(string("choose"),
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
                                                                                                                                                                  |v1|
                                                                                                                                                                  {
                                                                                                                                                                      let matchValue =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                      let matchValue_1 =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                      let matchValue_2 =
                                                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Op::Data_Op_Op(),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                                                                                                                                                                                                                   &&&matchValue_1),
                                                                                                                                                                                                                                                                                                                &&&matchValue_2)),
                                                                                                                                                                                                                                          &&&matchValue))
                                                                                                                                                                  }
                                                                                                                                                          })
                                                                                                                                      })),
                                                                                                         add(string("Divide0"),
                                                                                                             &&Func1::new({
                                                                                                                              let divideOp
                                                                                                                                  =
                                                                                                                                  divideOp.clone();
                                                                                                                              move
                                                                                                                                  |usd__unused|
                                                                                                                                  &divideOp
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
                                                             }))
    }
    pub fn Data_Decide_chooseEquivalence() -> &dyn Any {
        static Data_Decide_chooseEquivalence: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decide_chooseEquivalence.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decide::Data_Decide_Decideusd_Dict(),
                                                                                       &&&add(string("choose"),
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
                                                                                                                                                       |v1|
                                                                                                                                                       {
                                                                                                                                                           let matchValue =
                                                                                                                                                               Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                           let matchValue_1 =
                                                                                                                                                               Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                           let matchValue_2 =
                                                                                                                                                               Sharpurs_Prelude::unbox(v1);
                                                                                                                                                           let f1 =
                                                                                                                                                               matchValue;
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Equivalence::Data_Equivalence_Equivalence(),
                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                              let f1
                                                                                                                                                                                                                  =
                                                                                                                                                                                                                  f1.clone();
                                                                                                                                                                                                              let matchValue_1
                                                                                                                                                                                                                  =
                                                                                                                                                                                                                  matchValue_1.clone();
                                                                                                                                                                                                              let matchValue_2
                                                                                                                                                                                                                  =
                                                                                                                                                                                                                  matchValue_2.clone();
                                                                                                                                                                                                              move
                                                                                                                                                                                                                  |a|
                                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                                  let a
                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                      a.clone();
                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                      |b|
                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                          let matchValue_4:
                                                                                                                                                                                                                                                  LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                                         &&&a));
                                                                                                                                                                                                                                          match matchValue_4.as_ref()
                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                              Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_4_1_0)
                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                  let matchValue_6:
                                                                                                                                                                                                                                                          LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                                                 b));
                                                                                                                                                                                                                                                  match matchValue_6.as_ref()
                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                      Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_6_1_0)
                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_2,
                                                                                                                                                                                                                                                                                                                          &&matchValue_4_1_0),
                                                                                                                                                                                                                                                                                       &&matchValue_6_1_0),
                                                                                                                                                                                                                                                      _
                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                      &false,
                                                                                                                                                                                                                                                  }
                                                                                                                                                                                                                                              }
                                                                                                                                                                                                                                              Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_4_0_0)
                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                  let matchValue_5:
                                                                                                                                                                                                                                                          LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                                                 b));
                                                                                                                                                                                                                                                  match matchValue_5.as_ref()
                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                      Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_5_1_0)
                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                      &false,
                                                                                                                                                                                                                                                      Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_5_0_0)
                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                          &&matchValue_4_0_0),
                                                                                                                                                                                                                                                                                       &&matchValue_5_0_0),
                                                                                                                                                                                                                                                  }
                                                                                                                                                                                                                                              }
                                                                                                                                                                                                                                          }
                                                                                                                                                                                                                                      }
                                                                                                                                                                                                                              })
                                                                                                                                                                                                          }))
                                                                                                                                                       }
                                                                                                                                               })
                                                                                                                           })),
                                                                                              add(string("Divide0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_Divide::Data_Divide_divideEquivalence()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Decide_chooseComparison() -> &dyn Any {
        static Data_Decide_chooseComparison: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decide_chooseComparison.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decide::Data_Decide_Decideusd_Dict(),
                                                                                      &&&add(string("choose"),
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
                                                                                                                                                      |v1|
                                                                                                                                                      {
                                                                                                                                                          let matchValue =
                                                                                                                                                              Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                          let matchValue_1 =
                                                                                                                                                              Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                          let matchValue_2 =
                                                                                                                                                              Sharpurs_Prelude::unbox(v1);
                                                                                                                                                          let f1 =
                                                                                                                                                              matchValue;
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Comparison::Data_Comparison_Comparison(),
                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                             let f1
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 f1.clone();
                                                                                                                                                                                                             let matchValue_1
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 matchValue_1.clone();
                                                                                                                                                                                                             let matchValue_2
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 matchValue_2.clone();
                                                                                                                                                                                                             move
                                                                                                                                                                                                                 |a|
                                                                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                                                                 let a
                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                     a.clone();
                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                     |b|
                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                         let matchValue_4:
                                                                                                                                                                                                                                                 LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                                        &&&a));
                                                                                                                                                                                                                                         match matchValue_4.as_ref()
                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                             Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_4_1_0)
                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                 let matchValue_6:
                                                                                                                                                                                                                                                         LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                                                b));
                                                                                                                                                                                                                                                 match matchValue_6.as_ref()
                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                     Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_6_1_0)
                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_2,
                                                                                                                                                                                                                                                                                                                         &&matchValue_4_1_0),
                                                                                                                                                                                                                                                                                      &&matchValue_6_1_0),
                                                                                                                                                                                                                                                     _
                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor),
                                                                                                                                                                                                                                                 }
                                                                                                                                                                                                                                             }
                                                                                                                                                                                                                                             Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_4_0_0)
                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                 let matchValue_5:
                                                                                                                                                                                                                                                         LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                                                b));
                                                                                                                                                                                                                                                 match matchValue_5.as_ref()
                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                     Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_5_1_0)
                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor),
                                                                                                                                                                                                                                                     Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_5_0_0)
                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                         &&matchValue_4_0_0),
                                                                                                                                                                                                                                                                                      &&matchValue_5_0_0),
                                                                                                                                                                                                                                                 }
                                                                                                                                                                                                                                             }
                                                                                                                                                                                                                                         }
                                                                                                                                                                                                                                     }
                                                                                                                                                                                                                             })
                                                                                                                                                                                                         }))
                                                                                                                                                      }
                                                                                                                                              })
                                                                                                                          })),
                                                                                             add(string("Divide0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_Divide::Data_Divide_divideComparison()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_Decide_choose() -> &dyn Any {
        static Data_Decide_choose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decide_choose.get_or_init(||
                                           &Func1::new(move |dict|
                                                           find(string("choose"),
                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Decide_chosen() -> &dyn Any {
        static Data_Decide_chosen: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Decide_chosen.get_or_init(||
                                           &Func1::new(move |dictDecide|
                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Decide::Data_Decide_choose(),
                                                                                                                               dictDecide),
                                                                                            &&&PureScript_Data_Decide::Data_Decide_identity())))
    }
}
