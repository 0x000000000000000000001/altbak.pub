pub mod PureScript_Data_FunctorWithIndex {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_a8445950::PureScript_Data_Const;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_3ae611ad::PureScript_Data_Functor_App;
    use crate::module_ea451784::PureScript_Data_Functor_Compose;
    use crate::module_97eca9ab::PureScript_Data_Functor_Coproduct;
    use crate::module_4f0b17c7::PureScript_Data_Functor_Product;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_bde357b3::PureScript_Data_Maybe_First;
    use crate::module_a305e0e3::PureScript_Data_Maybe_Last;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_3b72fe33::PureScript_Data_Monoid_Additive;
    use crate::module_89cc47bd::PureScript_Data_Monoid_Conj;
    use crate::module_6a9d0e21::PureScript_Data_Monoid_Disj;
    use crate::module_5023b7e9::PureScript_Data_Monoid_Dual;
    use crate::module_7d83a5e5::PureScript_Data_Monoid_Multiplicative;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Data_FunctorWithIndex_FFI {
        use super::*;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_init;
        pub fn mapWithIndexArray(f: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            let res = new_init(&defaultOf(), count(arr.clone()));
            for i in 0_i32..=count(arr.clone()) - 1_i32 {
                let f1 = Sharpurs_Prelude::sharpurs_apply(f, &&i);
                res.get_mut()[i as usize] =
                    Sharpurs_Prelude::sharpurs_apply(&f1, &arr[i].clone())
            }
            &res
        }
    }
    pub fn Data_FunctorWithIndex_mapWithIndexArray() -> &dyn Any {
        static Data_FunctorWithIndex_mapWithIndexArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_mapWithIndexArray.get_or_init(||
                                                                &Func1::new(move
                                                                                |f|
                                                                                Func1::new({
                                                                                               let f
                                                                                                   =
                                                                                                   f.clone();
                                                                                               move
                                                                                                   |xs|
                                                                                                   PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FFI::mapWithIndexArray(&f,
                                                                                                                                                                                  xs)
                                                                                           })))
    }
    pub fn Data_FunctorWithIndex_map() -> &dyn Any {
        static Data_FunctorWithIndex_map: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                   &&&PureScript_Data_Tuple::Data_Tuple_functorTuple()))
    }
    pub fn Data_FunctorWithIndex_map1() -> &dyn Any {
        static Data_FunctorWithIndex_map1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map1.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_functorMultiplicative()))
    }
    pub fn Data_FunctorWithIndex_map2() -> &dyn Any {
        static Data_FunctorWithIndex_map2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map2.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()))
    }
    pub fn Data_FunctorWithIndex_map3() -> &dyn Any {
        static Data_FunctorWithIndex_map3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map3.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&&PureScript_Data_Maybe_Last::Data_Maybe_Last_functorLast()))
    }
    pub fn Data_FunctorWithIndex_map4() -> &dyn Any {
        static Data_FunctorWithIndex_map4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map4.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&&PureScript_Data_Maybe_First::Data_Maybe_First_functorFirst()))
    }
    pub fn Data_FunctorWithIndex_map5() -> &dyn Any {
        static Data_FunctorWithIndex_map5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map5.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&&PureScript_Data_Either::Data_Either_functorEither()))
    }
    pub fn Data_FunctorWithIndex_map6() -> &dyn Any {
        static Data_FunctorWithIndex_map6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map6.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_functorDual()))
    }
    pub fn Data_FunctorWithIndex_map7() -> &dyn Any {
        static Data_FunctorWithIndex_map7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map7.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_functorDisj()))
    }
    pub fn Data_FunctorWithIndex_map8() -> &dyn Any {
        static Data_FunctorWithIndex_map8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map8.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_functorConj()))
    }
    pub fn Data_FunctorWithIndex_map9() -> &dyn Any {
        static Data_FunctorWithIndex_map9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_map9.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                    &&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_functorAdditive()))
    }
    pub fn Data_FunctorWithIndex_FunctorWithIndexusd_Dict() -> &dyn Any {
        static Data_FunctorWithIndex_FunctorWithIndexusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_FunctorWithIndexusd_Dict.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |x|
                                                                                       x.clone()))
    }
    pub fn Data_FunctorWithIndex_mapWithIndex() -> &dyn Any {
        static Data_FunctorWithIndex_mapWithIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_mapWithIndex.get_or_init(||
                                                           &Func1::new(move
                                                                           |dict|
                                                                           find(string("mapWithIndex"),
                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_FunctorWithIndex_mapDefault() -> &dyn Any {
        static Data_FunctorWithIndex_mapDefault: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_mapDefault.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictFunctorWithIndex|
                                                                         &Func1::new({
                                                                                         let dictFunctorWithIndex
                                                                                             =
                                                                                             dictFunctorWithIndex.clone();
                                                                                         move
                                                                                             |f|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                 &&&dictFunctorWithIndex),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                 f))
                                                                                     })))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexTuple() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexTuple:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexTuple.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                     &&&add(string("mapWithIndex"),
                                                                                                            &&Func1::new(move
                                                                                                                             |f|
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                 &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map()),
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                 &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                            add(string("Functor0"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |usd__unused|
                                                                                                                                 &PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexProduct() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexProduct.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictFunctorWithIndex|
                                                                                      {
                                                                                          let functorProduct =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_functorProduct(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictFunctorWithIndex)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          &Func1::new({
                                                                                                          let dictFunctorWithIndex
                                                                                                              =
                                                                                                              dictFunctorWithIndex.clone();
                                                                                                          let functorProduct
                                                                                                              =
                                                                                                              functorProduct.clone();
                                                                                                          move
                                                                                                              |dictFunctorWithIndex1|
                                                                                                              {
                                                                                                                  let functorProduct1 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&functorProduct,
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictFunctorWithIndex1)),
                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                                                                   &&&add(string("mapWithIndex"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let dictFunctorWithIndex1
                                                                                                                                                                               =
                                                                                                                                                                               dictFunctorWithIndex1.clone();
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
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_Product(),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorTuple()),
                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&dictFunctorWithIndex),
                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                         |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))))),
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&dictFunctorWithIndex1),
                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                      |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))))),
                                                                                                                                                                                                                                                                           &&&matchValue_1))
                                                                                                                                                                                                   }
                                                                                                                                                                                           })
                                                                                                                                                                       }),
                                                                                                                                                          add(string("Functor0"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let functorProduct1
                                                                                                                                                                                   =
                                                                                                                                                                                   functorProduct1.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |usd__unused|
                                                                                                                                                                                   &functorProduct1
                                                                                                                                                                           }),
                                                                                                                                                              empty::<string,
                                                                                                                                                                      &dyn Any>())))
                                                                                                              }
                                                                                                      })
                                                                                      }))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexMultiplicative()
     -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexMultiplicative.get_or_init(||
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                              &&&add(string("mapWithIndex"),
                                                                                                                     &&Func1::new(move
                                                                                                                                      |f|
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                          &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map1()),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                          &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                     add(string("Functor0"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused|
                                                                                                                                          &PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_functorMultiplicative()),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexMaybe() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexMaybe:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexMaybe.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                     &&&add(string("mapWithIndex"),
                                                                                                            &&Func1::new(move
                                                                                                                             |f|
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                 &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map2()),
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                 &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                            add(string("Functor0"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |usd__unused|
                                                                                                                                 &PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexLast() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexLast:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexLast.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                    &&&add(string("mapWithIndex"),
                                                                                                           &&Func1::new(move
                                                                                                                            |f|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map3()),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                           add(string("Functor0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Data_Maybe_Last::Data_Maybe_Last_functorLast()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexIdentity() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexIdentity:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexIdentity.get_or_init(||
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
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Identity::Data_Identity_Identity(),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                               &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                            &&&matchValue_1))
                                                                                                                                                    }
                                                                                                                                            })),
                                                                                                               add(string("Functor0"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |usd__unused|
                                                                                                                                    &PureScript_Data_Identity::Data_Identity_functorIdentity()),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexFirst() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexFirst:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexFirst.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                     &&&add(string("mapWithIndex"),
                                                                                                            &&Func1::new(move
                                                                                                                             |f|
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                 &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map4()),
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                 &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                            add(string("Functor0"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |usd__unused|
                                                                                                                                 &PureScript_Data_Maybe_First::Data_Maybe_First_functorFirst()),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexEither() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexEither:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexEither.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                      &&&add(string("mapWithIndex"),
                                                                                                             &&Func1::new(move
                                                                                                                              |f|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                  &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map5()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                             add(string("Functor0"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |usd__unused|
                                                                                                                                  &PureScript_Data_Either::Data_Either_functorEither()),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexDual() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexDual:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexDual.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                    &&&add(string("mapWithIndex"),
                                                                                                           &&Func1::new(move
                                                                                                                            |f|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map6()),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                           add(string("Functor0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Data_Monoid_Dual::Data_Monoid_Dual_functorDual()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexDisj() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexDisj:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexDisj.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                    &&&add(string("mapWithIndex"),
                                                                                                           &&Func1::new(move
                                                                                                                            |f|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map7()),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                           add(string("Functor0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Data_Monoid_Disj::Data_Monoid_Disj_functorDisj()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexCoproduct() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexCoproduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexCoproduct.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictFunctorWithIndex|
                                                                                        {
                                                                                            let functorCoproduct =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_functorCoproduct(),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                           Sharpurs_Prelude::unbox(dictFunctorWithIndex)),
                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                            &Func1::new({
                                                                                                            let dictFunctorWithIndex
                                                                                                                =
                                                                                                                dictFunctorWithIndex.clone();
                                                                                                            let functorCoproduct
                                                                                                                =
                                                                                                                functorCoproduct.clone();
                                                                                                            move
                                                                                                                |dictFunctorWithIndex1|
                                                                                                                {
                                                                                                                    let functorCoproduct1 =
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&functorCoproduct,
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictFunctorWithIndex1)),
                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                                                                     &&&add(string("mapWithIndex"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let dictFunctorWithIndex1
                                                                                                                                                                                 =
                                                                                                                                                                                 dictFunctorWithIndex1.clone();
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
                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct(),
                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorEither()),
                                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                         &&&dictFunctorWithIndex),
                                                                                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                           |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                                                                                           &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))))),
                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                      &&&dictFunctorWithIndex1),
                                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                         &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                        |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))))),
                                                                                                                                                                                                                                                                             &&&matchValue_1))
                                                                                                                                                                                                     }
                                                                                                                                                                                             })
                                                                                                                                                                         }),
                                                                                                                                                            add(string("Functor0"),
                                                                                                                                                                &&Func1::new({
                                                                                                                                                                                 let functorCoproduct1
                                                                                                                                                                                     =
                                                                                                                                                                                     functorCoproduct1.clone();
                                                                                                                                                                                 move
                                                                                                                                                                                     |usd__unused|
                                                                                                                                                                                     &functorCoproduct1
                                                                                                                                                                             }),
                                                                                                                                                                empty::<string,
                                                                                                                                                                        &dyn Any>())))
                                                                                                                }
                                                                                                        })
                                                                                        }))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexConst() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexConst:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexConst.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                     &&&add(string("mapWithIndex"),
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
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_Const(),
                                                                                                                                                                                      &&&matchValue_1)
                                                                                                                                                 }
                                                                                                                                         })),
                                                                                                            add(string("Functor0"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |usd__unused|
                                                                                                                                 &PureScript_Data_Const::Data_Const_functorConst()),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexConj() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexConj:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexConj.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                    &&&add(string("mapWithIndex"),
                                                                                                           &&Func1::new(move
                                                                                                                            |f|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map8()),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                           add(string("Functor0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Data_Monoid_Conj::Data_Monoid_Conj_functorConj()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexCompose() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexCompose:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexCompose.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictFunctorWithIndex|
                                                                                      {
                                                                                          let functorCompose =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_functorCompose(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictFunctorWithIndex)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          &Func1::new({
                                                                                                          let dictFunctorWithIndex
                                                                                                              =
                                                                                                              dictFunctorWithIndex.clone();
                                                                                                          let functorCompose
                                                                                                              =
                                                                                                              functorCompose.clone();
                                                                                                          move
                                                                                                              |dictFunctorWithIndex1|
                                                                                                              {
                                                                                                                  let mapWithIndex1 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                       dictFunctorWithIndex1);
                                                                                                                  let functorCompose1 =
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&functorCompose,
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictFunctorWithIndex1)),
                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                                                                   &&&add(string("mapWithIndex"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let mapWithIndex1
                                                                                                                                                                               =
                                                                                                                                                                               mapWithIndex1.clone();
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
                                                                                                                                                                                                                                                                           &&&PureScript_Data_Functor_Compose::Data_Functor_Compose_Compose()),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                 &&&dictFunctorWithIndex),
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                    &&&mapWithIndex1),
                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_curry(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&matchValue))),
                                                                                                                                                                                                                                                                           &&&matchValue_1))
                                                                                                                                                                                                   }
                                                                                                                                                                                           })
                                                                                                                                                                       }),
                                                                                                                                                          add(string("Functor0"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let functorCompose1
                                                                                                                                                                                   =
                                                                                                                                                                                   functorCompose1.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |usd__unused|
                                                                                                                                                                                   &functorCompose1
                                                                                                                                                                           }),
                                                                                                                                                              empty::<string,
                                                                                                                                                                      &dyn Any>())))
                                                                                                              }
                                                                                                      })
                                                                                      }))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexArray() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexArray.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                     &&&add(string("mapWithIndex"),
                                                                                                            &&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndexArray(),
                                                                                                            add(string("Functor0"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |usd__unused|
                                                                                                                                 &PureScript_Data_Functor::Data_Functor_functorArray()),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexApp() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexApp:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexApp.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictFunctorWithIndex|
                                                                                  {
                                                                                      let functorApp =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_functorApp(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                     Sharpurs_Prelude::unbox(dictFunctorWithIndex)),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                                       &&&add(string("mapWithIndex"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let dictFunctorWithIndex
                                                                                                                                                   =
                                                                                                                                                   dictFunctorWithIndex.clone();
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
                                                                                                                                                                                                                                               &&&PureScript_Data_Functor_App::Data_Functor_App_App()),
                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                                                                                                                                     &&&dictFunctorWithIndex),
                                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                                               &&&matchValue_1))
                                                                                                                                                                       }
                                                                                                                                                               })
                                                                                                                                           }),
                                                                                                                              add(string("Functor0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let functorApp
                                                                                                                                                       =
                                                                                                                                                       functorApp.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &functorApp
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
    pub fn Data_FunctorWithIndex_functorWithIndexAdditive() -> &dyn Any {
        static Data_FunctorWithIndex_functorWithIndexAdditive:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_FunctorWithIndex_functorWithIndexAdditive.get_or_init(||
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                        &&&add(string("mapWithIndex"),
                                                                                                               &&Func1::new(move
                                                                                                                                |f|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                    &&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_map9()),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                    &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                               add(string("Functor0"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |usd__unused|
                                                                                                                                    &PureScript_Data_Monoid_Additive::Data_Monoid_Additive_functorAdditive()),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
}
