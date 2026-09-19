pub mod PureScript_Data_FoldableWithIndex {
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
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_97eca9ab::PureScript_Data_Functor_Coproduct;
    use crate::module_f8b1f47e::PureScript_Data_FunctorWithIndex;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_89cc47bd::PureScript_Data_Monoid_Conj;
    use crate::module_6a9d0e21::PureScript_Data_Monoid_Disj;
    use crate::module_5023b7e9::PureScript_Data_Monoid_Dual;
    use crate::module_f1079b5::PureScript_Data_Monoid_Endo;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_FoldableWithIndex_foldr() -> &dyn Any {
        static Data_FoldableWithIndex_foldr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldr.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                      &&&PureScript_Data_Foldable::Data_Foldable_foldableMultiplicative()))
    }
    pub fn Data_FoldableWithIndex_foldl() -> &dyn Any {
        static Data_FoldableWithIndex_foldl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldl.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                      &&&PureScript_Data_Foldable::Data_Foldable_foldableMultiplicative()))
    }
    pub fn Data_FoldableWithIndex_foldMap() -> &dyn Any {
        static Data_FoldableWithIndex_foldMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMap.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                        &&&PureScript_Data_Foldable::Data_Foldable_foldableMultiplicative()))
    }
    pub fn Data_FoldableWithIndex_foldr1() -> &dyn Any {
        static Data_FoldableWithIndex_foldr1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldr1.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableMaybe()))
    }
    pub fn Data_FoldableWithIndex_foldl1() -> &dyn Any {
        static Data_FoldableWithIndex_foldl1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldl1.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableMaybe()))
    }
    pub fn Data_FoldableWithIndex_foldMap1() -> &dyn Any {
        static Data_FoldableWithIndex_foldMap1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMap1.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_foldableMaybe()))
    }
    pub fn Data_FoldableWithIndex_foldr2() -> &dyn Any {
        static Data_FoldableWithIndex_foldr2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldr2.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableLast()))
    }
    pub fn Data_FoldableWithIndex_foldl2() -> &dyn Any {
        static Data_FoldableWithIndex_foldl2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldl2.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableLast()))
    }
    pub fn Data_FoldableWithIndex_foldMap2() -> &dyn Any {
        static Data_FoldableWithIndex_foldMap2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMap2.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_foldableLast()))
    }
    pub fn Data_FoldableWithIndex_foldr3() -> &dyn Any {
        static Data_FoldableWithIndex_foldr3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldr3.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableFirst()))
    }
    pub fn Data_FoldableWithIndex_foldl3() -> &dyn Any {
        static Data_FoldableWithIndex_foldl3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldl3.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableFirst()))
    }
    pub fn Data_FoldableWithIndex_foldMap3() -> &dyn Any {
        static Data_FoldableWithIndex_foldMap3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMap3.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_foldableFirst()))
    }
    pub fn Data_FoldableWithIndex_foldr4() -> &dyn Any {
        static Data_FoldableWithIndex_foldr4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldr4.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableDual()))
    }
    pub fn Data_FoldableWithIndex_foldl4() -> &dyn Any {
        static Data_FoldableWithIndex_foldl4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldl4.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableDual()))
    }
    pub fn Data_FoldableWithIndex_foldMap4() -> &dyn Any {
        static Data_FoldableWithIndex_foldMap4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMap4.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_foldableDual()))
    }
    pub fn Data_FoldableWithIndex_foldr5() -> &dyn Any {
        static Data_FoldableWithIndex_foldr5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldr5.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableDisj()))
    }
    pub fn Data_FoldableWithIndex_foldl5() -> &dyn Any {
        static Data_FoldableWithIndex_foldl5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldl5.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableDisj()))
    }
    pub fn Data_FoldableWithIndex_foldMap5() -> &dyn Any {
        static Data_FoldableWithIndex_foldMap5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMap5.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_foldableDisj()))
    }
    pub fn Data_FoldableWithIndex_foldr6() -> &dyn Any {
        static Data_FoldableWithIndex_foldr6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldr6.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableConj()))
    }
    pub fn Data_FoldableWithIndex_foldl6() -> &dyn Any {
        static Data_FoldableWithIndex_foldl6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldl6.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableConj()))
    }
    pub fn Data_FoldableWithIndex_foldMap6() -> &dyn Any {
        static Data_FoldableWithIndex_foldMap6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMap6.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_foldableConj()))
    }
    pub fn Data_FoldableWithIndex_foldr7() -> &dyn Any {
        static Data_FoldableWithIndex_foldr7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldr7.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableAdditive()))
    }
    pub fn Data_FoldableWithIndex_foldl7() -> &dyn Any {
        static Data_FoldableWithIndex_foldl7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldl7.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                       &&&PureScript_Data_Foldable::Data_Foldable_foldableAdditive()))
    }
    pub fn Data_FoldableWithIndex_foldMap7() -> &dyn Any {
        static Data_FoldableWithIndex_foldMap7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMap7.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                         &&&PureScript_Data_Foldable::Data_Foldable_foldableAdditive()))
    }
    pub fn Data_FoldableWithIndex_monoidDual() -> &dyn Any {
        static Data_FoldableWithIndex_monoidDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_monoidDual.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_monoidDual(),
                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_monoidEndo(),
                                                                                                                              &&&PureScript_Control_Category::Control_Category_categoryFn())))
    }
    pub fn Data_FoldableWithIndex_monoidEndo() -> &dyn Any {
        static Data_FoldableWithIndex_monoidEndo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_monoidEndo.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_monoidEndo(),
                                                                                           &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_FoldableWithIndex_monoidEndo1() -> &dyn Any {
        static Data_FoldableWithIndex_monoidEndo1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_monoidEndo1.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_monoidEndo(),
                                                                                            &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_FoldableWithIndex_unwrap() -> &dyn Any {
        static Data_FoldableWithIndex_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_unwrap.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                       &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_FoldableWithIndex_FoldableWithIndexusd_Dict() -> &dyn Any {
        static Data_FoldableWithIndex_FoldableWithIndexusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_FoldableWithIndexusd_Dict.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |x|
                                                                                         x.clone()))
    }
    pub fn Data_FoldableWithIndex_foldrWithIndex() -> &dyn Any {
        static Data_FoldableWithIndex_foldrWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldrWithIndex.get_or_init(||
                                                              &Func1::new(move
                                                                              |dict|
                                                                              find(string("foldrWithIndex"),
                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_FoldableWithIndex_traverseWithIndex_() -> &dyn Any {
        static Data_FoldableWithIndex_traverseWithIndex_:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_traverseWithIndex_.get_or_init(||
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
                                                                                                          |dictFoldableWithIndex|
                                                                                                          &Func1::new({
                                                                                                                          let dictFoldableWithIndex
                                                                                                                              =
                                                                                                                              dictFoldableWithIndex.clone();
                                                                                                                          move
                                                                                                                              |f|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                     &&&dictFoldableWithIndex),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |i|
                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                            &&&applySecond),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                            i))
                                                                                                                                                                                                                })),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                     &&&dictApplicative),
                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                      })
                                                                                                  })
                                                                                  }))
    }
    pub fn Data_FoldableWithIndex_forWithIndex_() -> &dyn Any {
        static Data_FoldableWithIndex_forWithIndex_: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_FoldableWithIndex_forWithIndex_.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictApplicative|
                                                                             {
                                                                                 let traverseWithIndex_1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_traverseWithIndex_(),
                                                                                                                      dictApplicative);
                                                                                 &Func1::new({
                                                                                                 let traverseWithIndex_1
                                                                                                     =
                                                                                                     traverseWithIndex_1.clone();
                                                                                                 move
                                                                                                     |dictFoldableWithIndex|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&traverseWithIndex_1,
                                                                                                                                                                         dictFoldableWithIndex))
                                                                                             })
                                                                             }))
    }
    pub fn Data_FoldableWithIndex_foldrDefault() -> &dyn Any {
        static Data_FoldableWithIndex_foldrDefault: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_FoldableWithIndex_foldrDefault.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFoldableWithIndex|
                                                                            &Func1::new({
                                                                                            let dictFoldableWithIndex
                                                                                                =
                                                                                                dictFoldableWithIndex.clone();
                                                                                            move
                                                                                                |f|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                    &&&dictFoldableWithIndex),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                    f))
                                                                                        })))
    }
    pub fn Data_FoldableWithIndex_foldlWithIndex() -> &dyn Any {
        static Data_FoldableWithIndex_foldlWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldlWithIndex.get_or_init(||
                                                              &Func1::new(move
                                                                              |dict|
                                                                              find(string("foldlWithIndex"),
                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_FoldableWithIndex_foldlDefault() -> &dyn Any {
        static Data_FoldableWithIndex_foldlDefault: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_FoldableWithIndex_foldlDefault.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFoldableWithIndex|
                                                                            &Func1::new({
                                                                                            let dictFoldableWithIndex
                                                                                                =
                                                                                                dictFoldableWithIndex.clone();
                                                                                            move
                                                                                                |f|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                    &&&dictFoldableWithIndex),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                    f))
                                                                                        })))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexTuple() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexTuple:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexTuple.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                       &&&add(string("foldrWithIndex"),
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
                                                                                                                                                                           let matchValue_2:
                                                                                                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                                               &&&match matchValue_2.as_ref()
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
                                                                                                              add(string("foldlWithIndex"),
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
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()),
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
                                                                                                                  add(string("foldMapWithIndex"),
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
                                                                                                                                                                           {
                                                                                                                                                                               let matchValue_8 =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                               let matchValue_9:
                                                                                                                                                                                       LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_8,
                                                                                                                                                                                                                                                   &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                       {
                                                                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                   })
                                                                                                                                                                           }
                                                                                                                                                                   }))),
                                                                                                                      add(string("Foldable0"),
                                                                                                                          &&Func1::new(move
                                                                                                                                           |usd__unused|
                                                                                                                                           &PureScript_Data_Foldable::Data_Foldable_foldableTuple()),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexMultiplicative()
     -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexMultiplicative.get_or_init(||
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                                &&&add(string("foldrWithIndex"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |f|
                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                            &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldr()),
                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                            &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                       add(string("foldlWithIndex"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |f_1|
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldl()),
                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f_1,
                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                           add(string("foldMapWithIndex"),
                                                                                                                               &&Func1::new(move
                                                                                                                                                |dictMonoid|
                                                                                                                                                {
                                                                                                                                                    let foldMap8 =
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMap(),
                                                                                                                                                                                         dictMonoid);
                                                                                                                                                    &Func1::new({
                                                                                                                                                                    let foldMap8
                                                                                                                                                                        =
                                                                                                                                                                        foldMap8.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |f_2|
                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                            &&&foldMap8),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(f_2,
                                                                                                                                                                                                                                            &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                                })
                                                                                                                                                }),
                                                                                                                               add(string("Foldable0"),
                                                                                                                                   &&Func1::new(move
                                                                                                                                                    |usd__unused|
                                                                                                                                                    &PureScript_Data_Foldable::Data_Foldable_foldableMultiplicative()),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexMaybe() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexMaybe:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexMaybe.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                       &&&add(string("foldrWithIndex"),
                                                                                                              &&Func1::new(move
                                                                                                                               |f|
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                   &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldr1()),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                   &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                              add(string("foldlWithIndex"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |f_1|
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                       &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldl1()),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(f_1,
                                                                                                                                                                                                       &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                  add(string("foldMapWithIndex"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |dictMonoid|
                                                                                                                                       {
                                                                                                                                           let foldMap8 =
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMap1(),
                                                                                                                                                                                dictMonoid);
                                                                                                                                           &Func1::new({
                                                                                                                                                           let foldMap8
                                                                                                                                                               =
                                                                                                                                                               foldMap8.clone();
                                                                                                                                                           move
                                                                                                                                                               |f_2|
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                   &&&foldMap8),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(f_2,
                                                                                                                                                                                                                                   &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                       })
                                                                                                                                       }),
                                                                                                                      add(string("Foldable0"),
                                                                                                                          &&Func1::new(move
                                                                                                                                           |usd__unused|
                                                                                                                                           &PureScript_Data_Foldable::Data_Foldable_foldableMaybe()),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexLast() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexLast:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexLast.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                      &&&add(string("foldrWithIndex"),
                                                                                                             &&Func1::new(move
                                                                                                                              |f|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                  &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldr2()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                             add(string("foldlWithIndex"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |f_1|
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                      &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldl2()),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(f_1,
                                                                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                 add(string("foldMapWithIndex"),
                                                                                                                     &&Func1::new(move
                                                                                                                                      |dictMonoid|
                                                                                                                                      {
                                                                                                                                          let foldMap8 =
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMap2(),
                                                                                                                                                                               dictMonoid);
                                                                                                                                          &Func1::new({
                                                                                                                                                          let foldMap8
                                                                                                                                                              =
                                                                                                                                                              foldMap8.clone();
                                                                                                                                                          move
                                                                                                                                                              |f_2|
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                  &&&foldMap8),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(f_2,
                                                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                      })
                                                                                                                                      }),
                                                                                                                     add(string("Foldable0"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused|
                                                                                                                                          &PureScript_Data_Foldable::Data_Foldable_foldableLast()),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexIdentity() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexIdentity:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexIdentity.get_or_init(||
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                          &&&add(string("foldrWithIndex"),
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
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                                                  &&&matchValue_2),
                                                                                                                                                                                                               &&&matchValue_1)
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
                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                                                      &&&matchValue_5),
                                                                                                                                                                                                                   &&&matchValue_6)
                                                                                                                                                                              }
                                                                                                                                                                      })
                                                                                                                                                  })),
                                                                                                                     add(string("foldMapWithIndex"),
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
                                                                                                                                                                              {
                                                                                                                                                                                  let matchValue_8 =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                                  let matchValue_9 =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_8,
                                                                                                                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                   &&&matchValue_9)
                                                                                                                                                                              }
                                                                                                                                                                      }))),
                                                                                                                         add(string("Foldable0"),
                                                                                                                             &&Func1::new(move
                                                                                                                                              |usd__unused|
                                                                                                                                              &PureScript_Data_Foldable::Data_Foldable_foldableIdentity()),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexFirst() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexFirst:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexFirst.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                       &&&add(string("foldrWithIndex"),
                                                                                                              &&Func1::new(move
                                                                                                                               |f|
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                   &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldr3()),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                   &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                              add(string("foldlWithIndex"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |f_1|
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                       &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldl3()),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(f_1,
                                                                                                                                                                                                       &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                  add(string("foldMapWithIndex"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |dictMonoid|
                                                                                                                                       {
                                                                                                                                           let foldMap8 =
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMap3(),
                                                                                                                                                                                dictMonoid);
                                                                                                                                           &Func1::new({
                                                                                                                                                           let foldMap8
                                                                                                                                                               =
                                                                                                                                                               foldMap8.clone();
                                                                                                                                                           move
                                                                                                                                                               |f_2|
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                   &&&foldMap8),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(f_2,
                                                                                                                                                                                                                                   &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                       })
                                                                                                                                       }),
                                                                                                                      add(string("Foldable0"),
                                                                                                                          &&Func1::new(move
                                                                                                                                           |usd__unused|
                                                                                                                                           &PureScript_Data_Foldable::Data_Foldable_foldableFirst()),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexEither() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexEither:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexEither.get_or_init(||
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                        &&&add(string("foldrWithIndex"),
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
                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                                                    &&matchValue_2_1_0),
                                                                                                                                                                                                                 &&&matchValue_1),
                                                                                                                                                                                _
                                                                                                                                                                                =>
                                                                                                                                                                                &matchValue_1,
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                })
                                                                                                                                            })),
                                                                                                               add(string("foldlWithIndex"),
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
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_4,
                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                                                        &&&matchValue_5),
                                                                                                                                                                                                                     &&matchValue_6_1_0),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    &matchValue_5,
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                    })
                                                                                                                                                })),
                                                                                                                   add(string("foldMapWithIndex"),
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
                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_8,
                                                                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                             &&matchValue_9_1_0),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            &mempty,
                                                                                                                                                                                        }
                                                                                                                                                                                    }
                                                                                                                                                                            })
                                                                                                                                                        })
                                                                                                                                        }),
                                                                                                                       add(string("Foldable0"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |usd__unused|
                                                                                                                                            &PureScript_Data_Foldable::Data_Foldable_foldableEither()),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexDual() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexDual:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexDual.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                      &&&add(string("foldrWithIndex"),
                                                                                                             &&Func1::new(move
                                                                                                                              |f|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                  &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldr4()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                             add(string("foldlWithIndex"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |f_1|
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                      &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldl4()),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(f_1,
                                                                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                 add(string("foldMapWithIndex"),
                                                                                                                     &&Func1::new(move
                                                                                                                                      |dictMonoid|
                                                                                                                                      {
                                                                                                                                          let foldMap8 =
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMap4(),
                                                                                                                                                                               dictMonoid);
                                                                                                                                          &Func1::new({
                                                                                                                                                          let foldMap8
                                                                                                                                                              =
                                                                                                                                                              foldMap8.clone();
                                                                                                                                                          move
                                                                                                                                                              |f_2|
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                  &&&foldMap8),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(f_2,
                                                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                      })
                                                                                                                                      }),
                                                                                                                     add(string("Foldable0"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused|
                                                                                                                                          &PureScript_Data_Foldable::Data_Foldable_foldableDual()),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexDisj() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexDisj:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexDisj.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                      &&&add(string("foldrWithIndex"),
                                                                                                             &&Func1::new(move
                                                                                                                              |f|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                  &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldr5()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                             add(string("foldlWithIndex"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |f_1|
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                      &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldl5()),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(f_1,
                                                                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                 add(string("foldMapWithIndex"),
                                                                                                                     &&Func1::new(move
                                                                                                                                      |dictMonoid|
                                                                                                                                      {
                                                                                                                                          let foldMap8 =
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMap5(),
                                                                                                                                                                               dictMonoid);
                                                                                                                                          &Func1::new({
                                                                                                                                                          let foldMap8
                                                                                                                                                              =
                                                                                                                                                              foldMap8.clone();
                                                                                                                                                          move
                                                                                                                                                              |f_2|
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                  &&&foldMap8),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(f_2,
                                                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                      })
                                                                                                                                      }),
                                                                                                                     add(string("Foldable0"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused|
                                                                                                                                          &PureScript_Data_Foldable::Data_Foldable_foldableDisj()),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexConst() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexConst:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexConst.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                       &&&add(string("foldrWithIndex"),
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
                                                                                                              add(string("foldlWithIndex"),
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
                                                                                                                  add(string("foldMapWithIndex"),
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
                                                                                                                      add(string("Foldable0"),
                                                                                                                          &&Func1::new(move
                                                                                                                                           |usd__unused|
                                                                                                                                           &PureScript_Data_Foldable::Data_Foldable_foldableConst()),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexConj() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexConj:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexConj.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                      &&&add(string("foldrWithIndex"),
                                                                                                             &&Func1::new(move
                                                                                                                              |f|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                  &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldr6()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                             add(string("foldlWithIndex"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |f_1|
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                      &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldl6()),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(f_1,
                                                                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                 add(string("foldMapWithIndex"),
                                                                                                                     &&Func1::new(move
                                                                                                                                      |dictMonoid|
                                                                                                                                      {
                                                                                                                                          let foldMap8 =
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMap6(),
                                                                                                                                                                               dictMonoid);
                                                                                                                                          &Func1::new({
                                                                                                                                                          let foldMap8
                                                                                                                                                              =
                                                                                                                                                              foldMap8.clone();
                                                                                                                                                          move
                                                                                                                                                              |f_2|
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                  &&&foldMap8),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(f_2,
                                                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                      })
                                                                                                                                      }),
                                                                                                                     add(string("Foldable0"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused|
                                                                                                                                          &PureScript_Data_Foldable::Data_Foldable_foldableConj()),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexAdditive() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexAdditive:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexAdditive.get_or_init(||
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                          &&&add(string("foldrWithIndex"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |f|
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                      &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldr7()),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                 add(string("foldlWithIndex"),
                                                                                                                     &&Func1::new(move
                                                                                                                                      |f_1|
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                          &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldl7()),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(f_1,
                                                                                                                                                                                                          &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                     add(string("foldMapWithIndex"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |dictMonoid|
                                                                                                                                          {
                                                                                                                                              let foldMap8 =
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMap7(),
                                                                                                                                                                                   dictMonoid);
                                                                                                                                              &Func1::new({
                                                                                                                                                              let foldMap8
                                                                                                                                                                  =
                                                                                                                                                                  foldMap8.clone();
                                                                                                                                                              move
                                                                                                                                                                  |f_2|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                      &&&foldMap8),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(f_2,
                                                                                                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                          })
                                                                                                                                          }),
                                                                                                                         add(string("Foldable0"),
                                                                                                                             &&Func1::new(move
                                                                                                                                              |usd__unused|
                                                                                                                                              &PureScript_Data_Foldable::Data_Foldable_foldableAdditive()),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>()))))))
    }
    pub fn Data_FoldableWithIndex_foldWithIndexM() -> &dyn Any {
        static Data_FoldableWithIndex_foldWithIndexM:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldWithIndexM.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictFoldableWithIndex|
                                                                              &Func1::new({
                                                                                              let dictFoldableWithIndex
                                                                                                  =
                                                                                                  dictFoldableWithIndex.clone();
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
                                                                                                                                              |a0|
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                                     &&&dictFoldableWithIndex),
                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                    |i|
                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                    let i
                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                        i.clone();
                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                        |ma|
                                                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                                                        let ma
                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                            ma.clone();
                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                            |b|
                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                   &&&Bind1),
                                                                                                                                                                                                                                                                                                                                                &&&ma),
                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                      &&&i)),
                                                                                                                                                                                                                                                                                                                                                b))
                                                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                                                }))),
                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                     &&&Applicative0),
                                                                                                                                                                                                                  a0))
                                                                                                                                      })
                                                                                                                  })
                                                                                                  }
                                                                                          })))
    }
    pub fn Data_FoldableWithIndex_foldMapWithIndexDefaultR() -> &dyn Any {
        static Data_FoldableWithIndex_foldMapWithIndexDefaultR:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMapWithIndexDefaultR.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictFoldableWithIndex|
                                                                                        &Func1::new({
                                                                                                        let dictFoldableWithIndex
                                                                                                            =
                                                                                                            dictFoldableWithIndex.clone();
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
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                           &&&dictFoldableWithIndex),
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
                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                             &&&Semigroup0),
                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                &&&i),
                                                                                                                                                                                                                                                                                                                                                                             &&&x)),
                                                                                                                                                                                                                                                                                                       acc)
                                                                                                                                                                                                                                                              })
                                                                                                                                                                                                                                          })
                                                                                                                                                                                                                      })),
                                                                                                                                                                     &&&mempty)
                                                                                                                            })
                                                                                                            }
                                                                                                    })))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexArray_0040105()
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                         &&&add(string("foldrWithIndex"),
                                                &&Func1::new(move |f|
                                                                 &Func1::new({
                                                                                 let f
                                                                                     =
                                                                                     f.clone();
                                                                                 move
                                                                                     |z|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()),
                                                                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                                                                 |v|
                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                     let matchValue:
                                                                                                                                                                                                                                                             LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                                                     &Func1::new(move
                                                                                                                                                                                                                                                                     |y|
                                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                                                      y))
                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                            z)),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                            &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexArray()),
                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                           |usd__arg1|
                                                                                                                                                                           Func1::new({
                                                                                                                                                                                          let usd__arg1
                                                                                                                                                                                              =
                                                                                                                                                                                              usd__arg1.clone();
                                                                                                                                                                                          move
                                                                                                                                                                                              |usd__arg2|
                                                                                                                                                                                              &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                      usd__arg2.clone()))
                                                                                                                                                                                      }))))
                                                                             })),
                                                add(string("foldlWithIndex"),
                                                    &&Func1::new(move |f_1|
                                                                     &Func1::new({
                                                                                     let f_1
                                                                                         =
                                                                                         f_1.clone();
                                                                                     move
                                                                                         |z_1|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                      &&&PureScript_Data_Foldable::Data_Foldable_foldableArray()),
                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                     |y_1|
                                                                                                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                                                                                                     let y_1
                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                         y_1.clone();
                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                         |v_1|
                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                             let matchValue_1:
                                                                                                                                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f_1,
                                                                                                                                                                                                                                                                                                                                                                                    &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                                                                                                                                                 &&&y_1),
                                                                                                                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                 })
                                                                                                                                                                                                                                                                         }
                                                                                                                                                                                                                                                                 }))),
                                                                                                                                                                                                z_1)),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexArray()),
                                                                                                                                                             &&&Func1::new(move
                                                                                                                                                                               |usd__arg1_1|
                                                                                                                                                                               Func1::new({
                                                                                                                                                                                              let usd__arg1_1
                                                                                                                                                                                                  =
                                                                                                                                                                                                  usd__arg1_1.clone();
                                                                                                                                                                                              move
                                                                                                                                                                                                  |usd__arg2_1|
                                                                                                                                                                                                  &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                          usd__arg2_1.clone()))
                                                                                                                                                                                          }))))
                                                                                 })),
                                                    add(string("foldMapWithIndex"),
                                                        &&Func1::new({
                                                                         let Data_FoldableWithIndex_foldableWithIndexArray_0040105_002d1
                                                                             =
                                                                             Data_FoldableWithIndex_foldableWithIndexArray_0040105_002d1.clone();
                                                                         move
                                                                             |dictMonoid|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndexDefaultR(),
                                                                                                                                                 &&&Data_FoldableWithIndex_foldableWithIndexArray_0040105_002d1.Value),
                                                                                                              dictMonoid)
                                                                     }),
                                                        add(string("Foldable0"),
                                                            &&Func1::new(move
                                                                             |usd__unused|
                                                                             &PureScript_Data_Foldable::Data_Foldable_foldableArray()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexArray_0040105_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_FoldableWithIndex_foldableWithIndexArray_0040105_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexArray_0040105_002d1.get_or_init(||
                                                                                    Lazy(Data_FoldableWithIndex_foldableWithIndexArray_0040105.clone()))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexArray() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexArray.get_or_init(||
                                                                      Data_FoldableWithIndex_foldableWithIndexArray_0040105_002d1.Value)
    }
    pub fn Data_FoldableWithIndex_foldMapWithIndexDefaultL() -> &dyn Any {
        static Data_FoldableWithIndex_foldMapWithIndexDefaultL:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMapWithIndexDefaultL.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictFoldableWithIndex|
                                                                                        &Func1::new({
                                                                                                        let dictFoldableWithIndex
                                                                                                            =
                                                                                                            dictFoldableWithIndex.clone();
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
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                           &&&dictFoldableWithIndex),
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
                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                             &&&i),
                                                                                                                                                                                                                                                                                                                                          x))
                                                                                                                                                                                                                                                              })
                                                                                                                                                                                                                                          })
                                                                                                                                                                                                                      })),
                                                                                                                                                                     &&&mempty)
                                                                                                                            })
                                                                                                            }
                                                                                                    })))
    }
    pub fn Data_FoldableWithIndex_foldMapWithIndex() -> &dyn Any {
        static Data_FoldableWithIndex_foldMapWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMapWithIndex.get_or_init(||
                                                                &Func1::new(move
                                                                                |dict|
                                                                                find(string("foldMapWithIndex"),
                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexApp() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexApp:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexApp.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictFoldableWithIndex|
                                                                                    {
                                                                                        let foldableApp =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldableApp(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(dictFoldableWithIndex)),
                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                                         &&&add(string("foldrWithIndex"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let dictFoldableWithIndex
                                                                                                                                                     =
                                                                                                                                                     dictFoldableWithIndex.clone();
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
                                                                                                                                                                                                 let matchValue_2 =
                                                                                                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                                                                                                           &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                                                     &&&matchValue_1),
                                                                                                                                                                                                                                  &&&matchValue_2)
                                                                                                                                                                                             }
                                                                                                                                                                                     })
                                                                                                                                                                 })
                                                                                                                                             }),
                                                                                                                                add(string("foldlWithIndex"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let dictFoldableWithIndex
                                                                                                                                                         =
                                                                                                                                                         dictFoldableWithIndex.clone();
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
                                                                                                                                                                                                     let matchValue_6 =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                                                                                                                               &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                            &&&matchValue_4),
                                                                                                                                                                                                                                                                         &&&matchValue_5),
                                                                                                                                                                                                                                      &&&matchValue_6)
                                                                                                                                                                                                 }
                                                                                                                                                                                         })
                                                                                                                                                                     })
                                                                                                                                                 }),
                                                                                                                                    add(string("foldMapWithIndex"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let dictFoldableWithIndex
                                                                                                                                                             =
                                                                                                                                                             dictFoldableWithIndex.clone();
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
                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                   &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                &&&dictMonoid),
                                                                                                                                                                                                                                                                             &&&matchValue_8),
                                                                                                                                                                                                                                          &&&matchValue_9)
                                                                                                                                                                                                     }
                                                                                                                                                                                             })
                                                                                                                                                                         })
                                                                                                                                                     }),
                                                                                                                                        add(string("Foldable0"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let foldableApp
                                                                                                                                                                 =
                                                                                                                                                                 foldableApp.clone();
                                                                                                                                                             move
                                                                                                                                                                 |usd__unused|
                                                                                                                                                                 &foldableApp
                                                                                                                                                         }),
                                                                                                                                            empty::<string,
                                                                                                                                                    &dyn Any>())))))
                                                                                    }))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexCompose() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexCompose:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexCompose.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictFoldableWithIndex|
                                                                                        {
                                                                                            let foldableCompose =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldableCompose(),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                           Sharpurs_Prelude::unbox(dictFoldableWithIndex)),
                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                            &Func1::new({
                                                                                                            let dictFoldableWithIndex
                                                                                                                =
                                                                                                                dictFoldableWithIndex.clone();
                                                                                                            let foldableCompose
                                                                                                                =
                                                                                                                foldableCompose.clone();
                                                                                                            move
                                                                                                                |dictFoldableWithIndex1|
                                                                                                                {
                                                                                                                    let foldlWithIndex1 =
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                         dictFoldableWithIndex1);
                                                                                                                    let foldMapWithIndex1 =
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                         dictFoldableWithIndex1);
                                                                                                                    let foldableCompose1 =
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&foldableCompose,
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictFoldableWithIndex1)),
                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                                                                     &&&add(string("foldrWithIndex"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let dictFoldableWithIndex1
                                                                                                                                                                                 =
                                                                                                                                                                                 dictFoldableWithIndex1.clone();
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
                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                       &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                      |a|
                                                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&dictFoldableWithIndex1),
                                                                                                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_curry(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                                                                                                                                                                                                                                                                                                             a))))),
                                                                                                                                                                                                                                                                                                 &&&matchValue_1),
                                                                                                                                                                                                                                                              &&&matchValue_2)
                                                                                                                                                                                                                         }
                                                                                                                                                                                                                 })
                                                                                                                                                                                             })
                                                                                                                                                                         }),
                                                                                                                                                            add(string("foldlWithIndex"),
                                                                                                                                                                &&Func1::new({
                                                                                                                                                                                 let foldlWithIndex1
                                                                                                                                                                                     =
                                                                                                                                                                                     foldlWithIndex1.clone();
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
                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                           &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                              &&&foldlWithIndex1),
                                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_curry(),
                                                                                                                                                                                                                                                                                                                                                                                                              &&&matchValue_4))),
                                                                                                                                                                                                                                                                                                     &&&matchValue_5),
                                                                                                                                                                                                                                                                  &&&matchValue_6)
                                                                                                                                                                                                                             }
                                                                                                                                                                                                                     })
                                                                                                                                                                                                 })
                                                                                                                                                                             }),
                                                                                                                                                                add(string("foldMapWithIndex"),
                                                                                                                                                                    &&Func1::new({
                                                                                                                                                                                     let foldMapWithIndex1
                                                                                                                                                                                         =
                                                                                                                                                                                         foldMapWithIndex1.clone();
                                                                                                                                                                                     move
                                                                                                                                                                                         |dictMonoid|
                                                                                                                                                                                         {
                                                                                                                                                                                             let foldMapWithIndex2 =
                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&foldMapWithIndex1,
                                                                                                                                                                                                                                  dictMonoid);
                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                             let dictMonoid
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 dictMonoid.clone();
                                                                                                                                                                                                             let foldMapWithIndex2
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 foldMapWithIndex2.clone();
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
                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                   &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                                &&&dictMonoid),
                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                   &&&foldMapWithIndex2),
                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_curry(),
                                                                                                                                                                                                                                                                                                                                                                                   &&&matchValue_8))),
                                                                                                                                                                                                                                                                          &&&matchValue_9)
                                                                                                                                                                                                                                     }
                                                                                                                                                                                                                             })
                                                                                                                                                                                                         })
                                                                                                                                                                                         }
                                                                                                                                                                                 }),
                                                                                                                                                                    add(string("Foldable0"),
                                                                                                                                                                        &&Func1::new({
                                                                                                                                                                                         let foldableCompose1
                                                                                                                                                                                             =
                                                                                                                                                                                             foldableCompose1.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |usd__unused|
                                                                                                                                                                                             &foldableCompose1
                                                                                                                                                                                     }),
                                                                                                                                                                        empty::<string,
                                                                                                                                                                                &dyn Any>())))))
                                                                                                                }
                                                                                                        })
                                                                                        }))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexCoproduct() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexCoproduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexCoproduct.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |dictFoldableWithIndex|
                                                                                          {
                                                                                              let foldableCoproduct =
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldableCoproduct(),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                             Sharpurs_Prelude::unbox(dictFoldableWithIndex)),
                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                              &Func1::new({
                                                                                                              let dictFoldableWithIndex
                                                                                                                  =
                                                                                                                  dictFoldableWithIndex.clone();
                                                                                                              let foldableCoproduct
                                                                                                                  =
                                                                                                                  foldableCoproduct.clone();
                                                                                                              move
                                                                                                                  |dictFoldableWithIndex1|
                                                                                                                  {
                                                                                                                      let foldableCoproduct1 =
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&foldableCoproduct,
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictFoldableWithIndex1)),
                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                                                                       &&&add(string("foldrWithIndex"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let dictFoldableWithIndex1
                                                                                                                                                                                   =
                                                                                                                                                                                   dictFoldableWithIndex1.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |f|
                                                                                                                                                                                   &Func1::new({
                                                                                                                                                                                                   let f
                                                                                                                                                                                                       =
                                                                                                                                                                                                       f.clone();
                                                                                                                                                                                                   move
                                                                                                                                                                                                       |z|
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&f),
                                                                                                                                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                      |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone()))))),
                                                                                                                                                                                                                                                                                                              z)),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                                                                                                                 &&&dictFoldableWithIndex1),
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                    &&&f),
                                                                                                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                   |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                   &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone()))))),
                                                                                                                                                                                                                                                                           z))
                                                                                                                                                                                               })
                                                                                                                                                                           }),
                                                                                                                                                              add(string("foldlWithIndex"),
                                                                                                                                                                  &&Func1::new({
                                                                                                                                                                                   let dictFoldableWithIndex1
                                                                                                                                                                                       =
                                                                                                                                                                                       dictFoldableWithIndex1.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |f_1|
                                                                                                                                                                                       &Func1::new({
                                                                                                                                                                                                       let f_1
                                                                                                                                                                                                           =
                                                                                                                                                                                                           f_1.clone();
                                                                                                                                                                                                       move
                                                                                                                                                                                                           |z_1|
                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                        &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&f_1),
                                                                                                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                          |usd__arg1_2|
                                                                                                                                                                                                                                                                                                                                                                                                          &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_2.clone()))))),
                                                                                                                                                                                                                                                                                                                  z_1)),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                                                                                                                                     &&&dictFoldableWithIndex1),
                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                        &&&f_1),
                                                                                                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                       |usd__arg1_3|
                                                                                                                                                                                                                                                                                                                                                                       &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_3.clone()))))),
                                                                                                                                                                                                                                                                               z_1))
                                                                                                                                                                                                   })
                                                                                                                                                                               }),
                                                                                                                                                                  add(string("foldMapWithIndex"),
                                                                                                                                                                      &&Func1::new({
                                                                                                                                                                                       let dictFoldableWithIndex1
                                                                                                                                                                                           =
                                                                                                                                                                                           dictFoldableWithIndex1.clone();
                                                                                                                                                                                       move
                                                                                                                                                                                           |dictMonoid|
                                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                                           let dictMonoid
                                                                                                                                                                                                               =
                                                                                                                                                                                                               dictMonoid.clone();
                                                                                                                                                                                                           move
                                                                                                                                                                                                               |f_2|
                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                            &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                                         &&&dictMonoid),
                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                            f_2),
                                                                                                                                                                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                           |usd__arg1_4|
                                                                                                                                                                                                                                                                                                                                                                           &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_4.clone())))))),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                         &&&dictFoldableWithIndex1),
                                                                                                                                                                                                                                                                                                                      &&&dictMonoid),
                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                         f_2),
                                                                                                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                        |usd__arg1_5|
                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_5.clone()))))))
                                                                                                                                                                                                       })
                                                                                                                                                                                   }),
                                                                                                                                                                      add(string("Foldable0"),
                                                                                                                                                                          &&Func1::new({
                                                                                                                                                                                           let foldableCoproduct1
                                                                                                                                                                                               =
                                                                                                                                                                                               foldableCoproduct1.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |usd__unused|
                                                                                                                                                                                               &foldableCoproduct1
                                                                                                                                                                                       }),
                                                                                                                                                                          empty::<string,
                                                                                                                                                                                  &dyn Any>())))))
                                                                                                                  }
                                                                                                          })
                                                                                          }))
    }
    pub fn Data_FoldableWithIndex_foldableWithIndexProduct() -> &dyn Any {
        static Data_FoldableWithIndex_foldableWithIndexProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldableWithIndexProduct.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictFoldableWithIndex|
                                                                                        {
                                                                                            let foldableProduct =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldableProduct(),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                           Sharpurs_Prelude::unbox(dictFoldableWithIndex)),
                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                            &Func1::new({
                                                                                                            let dictFoldableWithIndex
                                                                                                                =
                                                                                                                dictFoldableWithIndex.clone();
                                                                                                            let foldableProduct
                                                                                                                =
                                                                                                                foldableProduct.clone();
                                                                                                            move
                                                                                                                |dictFoldableWithIndex1|
                                                                                                                {
                                                                                                                    let foldableProduct1 =
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&foldableProduct,
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictFoldableWithIndex1)),
                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                                                                     &&&add(string("foldrWithIndex"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let dictFoldableWithIndex1
                                                                                                                                                                                 =
                                                                                                                                                                                 dictFoldableWithIndex1.clone();
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
                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                       &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&f1),
                                                                                                                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                         |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone()))))),
                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&dictFoldableWithIndex1),
                                                                                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                                               |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                                                                                                               &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone()))))),
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
                                                                                                                                                            add(string("foldlWithIndex"),
                                                                                                                                                                &&Func1::new({
                                                                                                                                                                                 let dictFoldableWithIndex1
                                                                                                                                                                                     =
                                                                                                                                                                                     dictFoldableWithIndex1.clone();
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
                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                           &&&dictFoldableWithIndex1),
                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                              &&&f1_1),
                                                                                                                                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                             |usd__arg1_2|
                                                                                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_2.clone()))))),
                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&f1_1),
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |usd__arg1_3|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_3.clone()))))),
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
                                                                                                                                                                add(string("foldMapWithIndex"),
                                                                                                                                                                    &&Func1::new({
                                                                                                                                                                                     let dictFoldableWithIndex1
                                                                                                                                                                                         =
                                                                                                                                                                                         dictFoldableWithIndex1.clone();
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
                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                                                                                                      &&&dictMonoid),
                                                                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&f1_2),
                                                                                                                                                                                                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                        |usd__arg1_4|
                                                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_4.clone()))))),
                                                                                                                                                                                                                                                                                                                                                &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                                                   })),
                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                      &&&dictFoldableWithIndex1),
                                                                                                                                                                                                                                                                                                                                                                                   &&&dictMonoid),
                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                      &&&f1_2),
                                                                                                                                                                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                     |usd__arg1_5|
                                                                                                                                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_5.clone()))))),
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
                                                                                                                                                                    add(string("Foldable0"),
                                                                                                                                                                        &&Func1::new({
                                                                                                                                                                                         let foldableProduct1
                                                                                                                                                                                             =
                                                                                                                                                                                             foldableProduct1.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |usd__unused|
                                                                                                                                                                                             &foldableProduct1
                                                                                                                                                                                     }),
                                                                                                                                                                        empty::<string,
                                                                                                                                                                                &dyn Any>())))))
                                                                                                                }
                                                                                                        })
                                                                                        }))
    }
    pub fn Data_FoldableWithIndex_foldlWithIndexDefault() -> &dyn Any {
        static Data_FoldableWithIndex_foldlWithIndexDefault:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldlWithIndexDefault.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictFoldableWithIndex|
                                                                                     &Func1::new({
                                                                                                     let dictFoldableWithIndex
                                                                                                         =
                                                                                                         dictFoldableWithIndex.clone();
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
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                    &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_monoidDual()),
                                                                                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                |i|
                                                                                                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual()),
                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo()),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&c,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          i)))))),
                                                                                                                                                                                                                                                                                           xs))),
                                                                                                                                                                                  &&&u)
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })))
    }
    pub fn Data_FoldableWithIndex_foldrWithIndexDefault() -> &dyn Any {
        static Data_FoldableWithIndex_foldrWithIndexDefault:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldrWithIndexDefault.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictFoldableWithIndex|
                                                                                     &Func1::new({
                                                                                                     let dictFoldableWithIndex
                                                                                                         =
                                                                                                         dictFoldableWithIndex.clone();
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
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_monoidEndo()),
                                                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                                                             |i|
                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo()),
                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&c,
                                                                                                                                                                                                                                                                                                                                                                                 i)))),
                                                                                                                                                                                                                                                        xs)),
                                                                                                                                                                                  &&&u)
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })))
    }
    pub fn Data_FoldableWithIndex_surroundMapWithIndex() -> &dyn Any {
        static Data_FoldableWithIndex_surroundMapWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_surroundMapWithIndex.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictFoldableWithIndex|
                                                                                    &Func1::new({
                                                                                                    let dictFoldableWithIndex
                                                                                                        =
                                                                                                        dictFoldableWithIndex.clone();
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
                                                                                                                                                                                            |i|
                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                            let i
                                                                                                                                                                                                                =
                                                                                                                                                                                                                i.clone();
                                                                                                                                                                                                            move
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
                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&t,
                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&i),
                                                                                                                                                                                                                                                                                                                                                                                                                 &&&a)),
                                                                                                                                                                                                                                                                                                                                           m))
                                                                                                                                                                                                                                                               }))
                                                                                                                                                                                                        }));
                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                        &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_monoidEndo1()),
                                                                                                                                                                                                                                                                                                                  &&&joined),
                                                                                                                                                                                                                                                                               f)),
                                                                                                                                                                                                         &&&d)
                                                                                                                                                                    }
                                                                                                                                                            })
                                                                                                                                        })
                                                                                                                    })
                                                                                                })))
    }
    pub fn Data_FoldableWithIndex_foldMapDefault() -> &dyn Any {
        static Data_FoldableWithIndex_foldMapDefault:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_foldMapDefault.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictFoldableWithIndex|
                                                                              &Func1::new({
                                                                                              let dictFoldableWithIndex
                                                                                                  =
                                                                                                  dictFoldableWithIndex.clone();
                                                                                              move
                                                                                                  |dictMonoid|
                                                                                                  &Func1::new({
                                                                                                                  let dictMonoid
                                                                                                                      =
                                                                                                                      dictMonoid.clone();
                                                                                                                  move
                                                                                                                      |f|
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                             &&&dictFoldableWithIndex),
                                                                                                                                                                                          &&&dictMonoid),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                          f))
                                                                                                              })
                                                                                          })))
    }
    pub fn Data_FoldableWithIndex_findWithIndex() -> &dyn Any {
        static Data_FoldableWithIndex_findWithIndex: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_FoldableWithIndex_findWithIndex.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictFoldableWithIndex|
                                                                             &Func1::new({
                                                                                             let dictFoldableWithIndex
                                                                                                 =
                                                                                                 dictFoldableWithIndex.clone();
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
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let v1
                                                                                                                                                                     =
                                                                                                                                                                     v1.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |v2|
                                                                                                                                                                     {
                                                                                                                                                                         let matchValue =
                                                                                                                                                                             Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                         let matchValue_1:
                                                                                                                                                                                 LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                             Sharpurs_Prelude::unbox(&&v1);
                                                                                                                                                                         let matchValue_2 =
                                                                                                                                                                             Sharpurs_Prelude::unbox(v2);
                                                                                                                                                                         if let Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                                                =
                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                             if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                                                                                                          &&&matchValue_2))
                                                                                                                                                                                {
                                                                                                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&add(string("index"),
                                                                                                                                                                                                                                             &&matchValue,
                                                                                                                                                                                                                                             add(string("value"),
                                                                                                                                                                                                                                                 &&matchValue_2,
                                                                                                                                                                                                                                                 empty::<string,
                                                                                                                                                                                                                                                         &dyn Any>()))))
                                                                                                                                                                             } else {
                                                                                                                                                                                 &matchValue_1
                                                                                                                                                                             }
                                                                                                                                                                         } else {
                                                                                                                                                                             &matchValue_1
                                                                                                                                                                         }
                                                                                                                                                                     }
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     });
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                            &&&dictFoldableWithIndex),
                                                                                                                                                                         &&&go),
                                                                                                                                      &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))
                                                                                                 }
                                                                                         })))
    }
    pub fn Data_FoldableWithIndex_findMapWithIndex() -> &dyn Any {
        static Data_FoldableWithIndex_findMapWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FoldableWithIndex_findMapWithIndex.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictFoldableWithIndex|
                                                                                &Func1::new({
                                                                                                let dictFoldableWithIndex
                                                                                                    =
                                                                                                    dictFoldableWithIndex.clone();
                                                                                                move
                                                                                                    |f|
                                                                                                    {
                                                                                                        let go =
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
                                                                                                                                                    &Func1::new({
                                                                                                                                                                    let v1
                                                                                                                                                                        =
                                                                                                                                                                        v1.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |v2|
                                                                                                                                                                        {
                                                                                                                                                                            let matchValue =
                                                                                                                                                                                Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                            let matchValue_1:
                                                                                                                                                                                    LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                Sharpurs_Prelude::unbox(&&v1);
                                                                                                                                                                            let matchValue_2 =
                                                                                                                                                                                Sharpurs_Prelude::unbox(v2);
                                                                                                                                                                            if let Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                                                   =
                                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                                 &&&matchValue_2)
                                                                                                                                                                            } else {
                                                                                                                                                                                &matchValue_1
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                })
                                                                                                                                            })
                                                                                                                        });
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                               &&&dictFoldableWithIndex),
                                                                                                                                                                            &&&go),
                                                                                                                                         &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))
                                                                                                    }
                                                                                            })))
    }
    pub fn Data_FoldableWithIndex_anyWithIndex() -> &dyn Any {
        static Data_FoldableWithIndex_anyWithIndex: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_FoldableWithIndex_anyWithIndex.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFoldableWithIndex|
                                                                            &Func1::new({
                                                                                            let dictFoldableWithIndex
                                                                                                =
                                                                                                dictFoldableWithIndex.clone();
                                                                                            move
                                                                                                |dictHeytingAlgebra|
                                                                                                {
                                                                                                    let monoidDisj =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_monoidDisj(),
                                                                                                                                         dictHeytingAlgebra);
                                                                                                    &Func1::new({
                                                                                                                    let monoidDisj
                                                                                                                        =
                                                                                                                        monoidDisj.clone();
                                                                                                                    move
                                                                                                                        |t|
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                            &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_unwrap()),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                  &&&dictFoldableWithIndex),
                                                                                                                                                                                                                               &&&monoidDisj),
                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                              let t
                                                                                                                                                                                                                  =
                                                                                                                                                                                                                  t.clone();
                                                                                                                                                                                                              move
                                                                                                                                                                                                                  |i|
                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_Disj()),
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&t,
                                                                                                                                                                                                                                                                                      i))
                                                                                                                                                                                                          })))
                                                                                                                })
                                                                                                }
                                                                                        })))
    }
    pub fn Data_FoldableWithIndex_allWithIndex() -> &dyn Any {
        static Data_FoldableWithIndex_allWithIndex: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_FoldableWithIndex_allWithIndex.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFoldableWithIndex|
                                                                            &Func1::new({
                                                                                            let dictFoldableWithIndex
                                                                                                =
                                                                                                dictFoldableWithIndex.clone();
                                                                                            move
                                                                                                |dictHeytingAlgebra|
                                                                                                {
                                                                                                    let monoidConj =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_monoidConj(),
                                                                                                                                         dictHeytingAlgebra);
                                                                                                    &Func1::new({
                                                                                                                    let monoidConj
                                                                                                                        =
                                                                                                                        monoidConj.clone();
                                                                                                                    move
                                                                                                                        |t|
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                            &&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_unwrap()),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                  &&&dictFoldableWithIndex),
                                                                                                                                                                                                                               &&&monoidConj),
                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                              let t
                                                                                                                                                                                                                  =
                                                                                                                                                                                                                  t.clone();
                                                                                                                                                                                                              move
                                                                                                                                                                                                                  |i|
                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj()),
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&t,
                                                                                                                                                                                                                                                                                      i))
                                                                                                                                                                                                          })))
                                                                                                                })
                                                                                                }
                                                                                        })))
    }
}
