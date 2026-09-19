pub mod PureScript_Data_TraversableWithIndex {
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
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_a8445950::PureScript_Data_Const;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_9201da02::PureScript_Data_FoldableWithIndex;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_3ae611ad::PureScript_Data_Functor_App;
    use crate::module_ea451784::PureScript_Data_Functor_Compose;
    use crate::module_97eca9ab::PureScript_Data_Functor_Coproduct;
    use crate::module_4f0b17c7::PureScript_Data_Functor_Product;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f8b1f47e::PureScript_Data_FunctorWithIndex;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_833c8c54::PureScript_Data_Traversable_Accum_Internal;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_TraversableWithIndex_traverse() -> &dyn Any {
        static Data_TraversableWithIndex_traverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traverse.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                            &&&PureScript_Data_Traversable::Data_Traversable_traversableMultiplicative()))
    }
    pub fn Data_TraversableWithIndex_traverse1() -> &dyn Any {
        static Data_TraversableWithIndex_traverse1: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_TraversableWithIndex_traverse1.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                             &&&PureScript_Data_Traversable::Data_Traversable_traversableMaybe()))
    }
    pub fn Data_TraversableWithIndex_traverse2() -> &dyn Any {
        static Data_TraversableWithIndex_traverse2: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_TraversableWithIndex_traverse2.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                             &&&PureScript_Data_Traversable::Data_Traversable_traversableLast()))
    }
    pub fn Data_TraversableWithIndex_traverse3() -> &dyn Any {
        static Data_TraversableWithIndex_traverse3: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_TraversableWithIndex_traverse3.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                             &&&PureScript_Data_Traversable::Data_Traversable_traversableFirst()))
    }
    pub fn Data_TraversableWithIndex_traverse4() -> &dyn Any {
        static Data_TraversableWithIndex_traverse4: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_TraversableWithIndex_traverse4.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                             &&&PureScript_Data_Traversable::Data_Traversable_traversableDual()))
    }
    pub fn Data_TraversableWithIndex_traverse5() -> &dyn Any {
        static Data_TraversableWithIndex_traverse5: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_TraversableWithIndex_traverse5.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                             &&&PureScript_Data_Traversable::Data_Traversable_traversableDisj()))
    }
    pub fn Data_TraversableWithIndex_traverse6() -> &dyn Any {
        static Data_TraversableWithIndex_traverse6: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_TraversableWithIndex_traverse6.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                             &&&PureScript_Data_Traversable::Data_Traversable_traversableConj()))
    }
    pub fn Data_TraversableWithIndex_traverse7() -> &dyn Any {
        static Data_TraversableWithIndex_traverse7: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_TraversableWithIndex_traverse7.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                             &&&PureScript_Data_Traversable::Data_Traversable_traversableAdditive()))
    }
    pub fn Data_TraversableWithIndex_TraversableWithIndexusd_Dict()
     -> &dyn Any {
        static Data_TraversableWithIndex_TraversableWithIndexusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_TraversableWithIndexusd_Dict.get_or_init(||
                                                                               &Func1::new(move
                                                                                               |x|
                                                                                               x.clone()))
    }
    pub fn Data_TraversableWithIndex_traverseWithIndexDefault() -> &dyn Any {
        static Data_TraversableWithIndex_traverseWithIndexDefault:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traverseWithIndexDefault.get_or_init(||
                                                                           &Func1::new(move
                                                                                           |dictTraversableWithIndex|
                                                                                           {
                                                                                               let sequence =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable2"),
                                                                                                                                                                              Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                               let FunctorWithIndex0 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("FunctorWithIndex0"),
                                                                                                                                           Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                               &Func1::new({
                                                                                                               let FunctorWithIndex0
                                                                                                                   =
                                                                                                                   FunctorWithIndex0.clone();
                                                                                                               let sequence
                                                                                                                   =
                                                                                                                   sequence.clone();
                                                                                                               move
                                                                                                                   |dictApplicative|
                                                                                                                   {
                                                                                                                       let sequence1 =
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&sequence,
                                                                                                                                                            dictApplicative);
                                                                                                                       &Func1::new({
                                                                                                                                       let sequence1
                                                                                                                                           =
                                                                                                                                           sequence1.clone();
                                                                                                                                       move
                                                                                                                                           |f|
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                               &&&sequence1),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                                                                  &&&FunctorWithIndex0),
                                                                                                                                                                                                               f))
                                                                                                                                   })
                                                                                                                   }
                                                                                                           })
                                                                                           }))
    }
    pub fn Data_TraversableWithIndex_traverseWithIndex() -> &dyn Any {
        static Data_TraversableWithIndex_traverseWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traverseWithIndex.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dict|
                                                                                    find(string("traverseWithIndex"),
                                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_TraversableWithIndex_traverseDefault() -> &dyn Any {
        static Data_TraversableWithIndex_traverseDefault:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traverseDefault.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictTraversableWithIndex|
                                                                                  &Func1::new({
                                                                                                  let dictTraversableWithIndex
                                                                                                      =
                                                                                                      dictTraversableWithIndex.clone();
                                                                                                  move
                                                                                                      |dictApplicative|
                                                                                                      &Func1::new({
                                                                                                                      let dictApplicative
                                                                                                                          =
                                                                                                                          dictApplicative.clone();
                                                                                                                      move
                                                                                                                          |f|
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                 &&&dictTraversableWithIndex),
                                                                                                                                                                                              &&&dictApplicative),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                              f))
                                                                                                                  })
                                                                                              })))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexTuple() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexTuple:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexTuple.get_or_init(||
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
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Unit::Data_Unit_unit()),
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
                                                                                                                    add(string("FunctorWithIndex0"),
                                                                                                                        &&Func1::new(move
                                                                                                                                         |usd__unused|
                                                                                                                                         &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexTuple()),
                                                                                                                        add(string("FoldableWithIndex1"),
                                                                                                                            &&Func1::new(move
                                                                                                                                             |usd__unused_1|
                                                                                                                                             &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexTuple()),
                                                                                                                            add(string("Traversable2"),
                                                                                                                                &&Func1::new(move
                                                                                                                                                 |usd__unused_2|
                                                                                                                                                 &PureScript_Data_Traversable::Data_Traversable_traversableTuple()),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexProduct()
     -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexProduct.get_or_init(||
                                                                              &Func1::new(move
                                                                                              |dictTraversableWithIndex|
                                                                                              {
                                                                                                  let functorWithIndexProduct =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexProduct(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FunctorWithIndex0"),
                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                  let foldableWithIndexProduct =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexProduct(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FoldableWithIndex1"),
                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                  let traversableProduct =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traversableProduct(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable2"),
                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                  &Func1::new({
                                                                                                                  let dictTraversableWithIndex
                                                                                                                      =
                                                                                                                      dictTraversableWithIndex.clone();
                                                                                                                  let foldableWithIndexProduct
                                                                                                                      =
                                                                                                                      foldableWithIndexProduct.clone();
                                                                                                                  let functorWithIndexProduct
                                                                                                                      =
                                                                                                                      functorWithIndexProduct.clone();
                                                                                                                  let traversableProduct
                                                                                                                      =
                                                                                                                      traversableProduct.clone();
                                                                                                                  move
                                                                                                                      |dictTraversableWithIndex1|
                                                                                                                      {
                                                                                                                          let functorWithIndexProduct1 =
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&functorWithIndexProduct,
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FunctorWithIndex0"),
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictTraversableWithIndex1)),
                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                          let foldableWithIndexProduct1 =
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&foldableWithIndexProduct,
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FoldableWithIndex1"),
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictTraversableWithIndex1)),
                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                          let traversableProduct1 =
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&traversableProduct,
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable2"),
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictTraversableWithIndex1)),
                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                                                                           &&&add(string("traverseWithIndex"),
                                                                                                                                                                  &&Func1::new({
                                                                                                                                                                                   let dictTraversableWithIndex1
                                                                                                                                                                                       =
                                                                                                                                                                                       dictTraversableWithIndex1.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |dictApplicative|
                                                                                                                                                                                       {
                                                                                                                                                                                           let Apply0 =
                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                                           let Apply0
                                                                                                                                                                                                               =
                                                                                                                                                                                                               Apply0.clone();
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
                                                                                                                                                                                                                                       let matchValue_1:
                                                                                                                                                                                                                                               LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                                       let f1 =
                                                                                                                                                                                                                                           matchValue;
                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                                                                                                                                                                                                                 &&&Apply0),
                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Functor_Product::Data_Functor_Product_product()),
                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&dictTraversableWithIndex),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                      |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone()))))),
                                                                                                                                                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&dictTraversableWithIndex1),
                                                                                                                                                                                                                                                                                                                                                                                 &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                   |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                                                   &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone()))))),
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
                                                                                                                                                                                       }
                                                                                                                                                                               }),
                                                                                                                                                                  add(string("FunctorWithIndex0"),
                                                                                                                                                                      &&Func1::new({
                                                                                                                                                                                       let functorWithIndexProduct1
                                                                                                                                                                                           =
                                                                                                                                                                                           functorWithIndexProduct1.clone();
                                                                                                                                                                                       move
                                                                                                                                                                                           |usd__unused|
                                                                                                                                                                                           &functorWithIndexProduct1
                                                                                                                                                                                   }),
                                                                                                                                                                      add(string("FoldableWithIndex1"),
                                                                                                                                                                          &&Func1::new({
                                                                                                                                                                                           let foldableWithIndexProduct1
                                                                                                                                                                                               =
                                                                                                                                                                                               foldableWithIndexProduct1.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |usd__unused_1|
                                                                                                                                                                                               &foldableWithIndexProduct1
                                                                                                                                                                                       }),
                                                                                                                                                                          add(string("Traversable2"),
                                                                                                                                                                              &&Func1::new({
                                                                                                                                                                                               let traversableProduct1
                                                                                                                                                                                                   =
                                                                                                                                                                                                   traversableProduct1.clone();
                                                                                                                                                                                               move
                                                                                                                                                                                                   |usd__unused_2|
                                                                                                                                                                                                   &traversableProduct1
                                                                                                                                                                                           }),
                                                                                                                                                                              empty::<string,
                                                                                                                                                                                      &dyn Any>())))))
                                                                                                                      }
                                                                                                              })
                                                                                              }))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexMultiplicative()
     -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexMultiplicative.get_or_init(||
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                                      &&&add(string("traverseWithIndex"),
                                                                                                                             &&Func1::new(move
                                                                                                                                              |dictApplicative|
                                                                                                                                              {
                                                                                                                                                  let traverse8 =
                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverse(),
                                                                                                                                                                                       dictApplicative);
                                                                                                                                                  &Func1::new({
                                                                                                                                                                  let traverse8
                                                                                                                                                                      =
                                                                                                                                                                      traverse8.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |f|
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                          &&&traverse8),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                          &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                              })
                                                                                                                                              }),
                                                                                                                             add(string("FunctorWithIndex0"),
                                                                                                                                 &&Func1::new(move
                                                                                                                                                  |usd__unused|
                                                                                                                                                  &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexMultiplicative()),
                                                                                                                                 add(string("FoldableWithIndex1"),
                                                                                                                                     &&Func1::new(move
                                                                                                                                                      |usd__unused_1|
                                                                                                                                                      &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexMultiplicative()),
                                                                                                                                     add(string("Traversable2"),
                                                                                                                                         &&Func1::new(move
                                                                                                                                                          |usd__unused_2|
                                                                                                                                                          &PureScript_Data_Traversable::Data_Traversable_traversableMultiplicative()),
                                                                                                                                         empty::<string,
                                                                                                                                                 &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexMaybe() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexMaybe:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexMaybe.get_or_init(||
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                             &&&add(string("traverseWithIndex"),
                                                                                                                    &&Func1::new(move
                                                                                                                                     |dictApplicative|
                                                                                                                                     {
                                                                                                                                         let traverse8 =
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverse1(),
                                                                                                                                                                              dictApplicative);
                                                                                                                                         &Func1::new({
                                                                                                                                                         let traverse8
                                                                                                                                                             =
                                                                                                                                                             traverse8.clone();
                                                                                                                                                         move
                                                                                                                                                             |f|
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                 &&&traverse8),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                 &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                     })
                                                                                                                                     }),
                                                                                                                    add(string("FunctorWithIndex0"),
                                                                                                                        &&Func1::new(move
                                                                                                                                         |usd__unused|
                                                                                                                                         &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexMaybe()),
                                                                                                                        add(string("FoldableWithIndex1"),
                                                                                                                            &&Func1::new(move
                                                                                                                                             |usd__unused_1|
                                                                                                                                             &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexMaybe()),
                                                                                                                            add(string("Traversable2"),
                                                                                                                                &&Func1::new(move
                                                                                                                                                 |usd__unused_2|
                                                                                                                                                 &PureScript_Data_Traversable::Data_Traversable_traversableMaybe()),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexLast() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexLast:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexLast.get_or_init(||
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                            &&&add(string("traverseWithIndex"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |dictApplicative|
                                                                                                                                    {
                                                                                                                                        let traverse8 =
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverse2(),
                                                                                                                                                                             dictApplicative);
                                                                                                                                        &Func1::new({
                                                                                                                                                        let traverse8
                                                                                                                                                            =
                                                                                                                                                            traverse8.clone();
                                                                                                                                                        move
                                                                                                                                                            |f|
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                &&&traverse8),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                    })
                                                                                                                                    }),
                                                                                                                   add(string("FunctorWithIndex0"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |usd__unused|
                                                                                                                                        &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexLast()),
                                                                                                                       add(string("FoldableWithIndex1"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |usd__unused_1|
                                                                                                                                            &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexLast()),
                                                                                                                           add(string("Traversable2"),
                                                                                                                               &&Func1::new(move
                                                                                                                                                |usd__unused_2|
                                                                                                                                                &PureScript_Data_Traversable::Data_Traversable_traversableLast()),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexIdentity()
     -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexIdentity:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexIdentity.get_or_init(||
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
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                                                            &&&matchValue_1))
                                                                                                                                                                                    }
                                                                                                                                                                            })
                                                                                                                                                        })
                                                                                                                                        }),
                                                                                                                       add(string("FunctorWithIndex0"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |usd__unused|
                                                                                                                                            &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexIdentity()),
                                                                                                                           add(string("FoldableWithIndex1"),
                                                                                                                               &&Func1::new(move
                                                                                                                                                |usd__unused_1|
                                                                                                                                                &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexIdentity()),
                                                                                                                               add(string("Traversable2"),
                                                                                                                                   &&Func1::new(move
                                                                                                                                                    |usd__unused_2|
                                                                                                                                                    &PureScript_Data_Traversable::Data_Traversable_traversableIdentity()),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexFirst() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexFirst:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexFirst.get_or_init(||
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                             &&&add(string("traverseWithIndex"),
                                                                                                                    &&Func1::new(move
                                                                                                                                     |dictApplicative|
                                                                                                                                     {
                                                                                                                                         let traverse8 =
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverse3(),
                                                                                                                                                                              dictApplicative);
                                                                                                                                         &Func1::new({
                                                                                                                                                         let traverse8
                                                                                                                                                             =
                                                                                                                                                             traverse8.clone();
                                                                                                                                                         move
                                                                                                                                                             |f|
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                 &&&traverse8),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                 &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                     })
                                                                                                                                     }),
                                                                                                                    add(string("FunctorWithIndex0"),
                                                                                                                        &&Func1::new(move
                                                                                                                                         |usd__unused|
                                                                                                                                         &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexFirst()),
                                                                                                                        add(string("FoldableWithIndex1"),
                                                                                                                            &&Func1::new(move
                                                                                                                                             |usd__unused_1|
                                                                                                                                             &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexFirst()),
                                                                                                                            add(string("Traversable2"),
                                                                                                                                &&Func1::new(move
                                                                                                                                                 |usd__unused_2|
                                                                                                                                                 &PureScript_Data_Traversable::Data_Traversable_traversableFirst()),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexEither()
     -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexEither:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexEither.get_or_init(||
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
                                                                                                                                                                                      let matchValue_1:
                                                                                                                                                                                              LrcPtr<Data_Either_Either> =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                      match matchValue_1.as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                          Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                          =>
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                 &&&Functor0),
                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                |usd__arg1|
                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                                                              &&matchValue_1_1_0)),
                                                                                                                                                                                          Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                                                                                                                          =>
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                              &&&dictApplicative),
                                                                                                                                                                                                                           &&&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0))),
                                                                                                                                                                                      }
                                                                                                                                                                                  }
                                                                                                                                                                          })
                                                                                                                                                      })
                                                                                                                                      }),
                                                                                                                     add(string("FunctorWithIndex0"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused|
                                                                                                                                          &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexEither()),
                                                                                                                         add(string("FoldableWithIndex1"),
                                                                                                                             &&Func1::new(move
                                                                                                                                              |usd__unused_1|
                                                                                                                                              &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexEither()),
                                                                                                                             add(string("Traversable2"),
                                                                                                                                 &&Func1::new(move
                                                                                                                                                  |usd__unused_2|
                                                                                                                                                  &PureScript_Data_Traversable::Data_Traversable_traversableEither()),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexDual() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexDual:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexDual.get_or_init(||
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                            &&&add(string("traverseWithIndex"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |dictApplicative|
                                                                                                                                    {
                                                                                                                                        let traverse8 =
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverse4(),
                                                                                                                                                                             dictApplicative);
                                                                                                                                        &Func1::new({
                                                                                                                                                        let traverse8
                                                                                                                                                            =
                                                                                                                                                            traverse8.clone();
                                                                                                                                                        move
                                                                                                                                                            |f|
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                &&&traverse8),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                    })
                                                                                                                                    }),
                                                                                                                   add(string("FunctorWithIndex0"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |usd__unused|
                                                                                                                                        &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexDual()),
                                                                                                                       add(string("FoldableWithIndex1"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |usd__unused_1|
                                                                                                                                            &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexDual()),
                                                                                                                           add(string("Traversable2"),
                                                                                                                               &&Func1::new(move
                                                                                                                                                |usd__unused_2|
                                                                                                                                                &PureScript_Data_Traversable::Data_Traversable_traversableDual()),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexDisj() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexDisj:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexDisj.get_or_init(||
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                            &&&add(string("traverseWithIndex"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |dictApplicative|
                                                                                                                                    {
                                                                                                                                        let traverse8 =
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverse5(),
                                                                                                                                                                             dictApplicative);
                                                                                                                                        &Func1::new({
                                                                                                                                                        let traverse8
                                                                                                                                                            =
                                                                                                                                                            traverse8.clone();
                                                                                                                                                        move
                                                                                                                                                            |f|
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                &&&traverse8),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                    })
                                                                                                                                    }),
                                                                                                                   add(string("FunctorWithIndex0"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |usd__unused|
                                                                                                                                        &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexDisj()),
                                                                                                                       add(string("FoldableWithIndex1"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |usd__unused_1|
                                                                                                                                            &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexDisj()),
                                                                                                                           add(string("Traversable2"),
                                                                                                                               &&Func1::new(move
                                                                                                                                                |usd__unused_2|
                                                                                                                                                &PureScript_Data_Traversable::Data_Traversable_traversableDisj()),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexCoproduct()
     -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexCoproduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexCoproduct.get_or_init(||
                                                                                &Func1::new(move
                                                                                                |dictTraversableWithIndex|
                                                                                                {
                                                                                                    let functorWithIndexCoproduct =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexCoproduct(),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FunctorWithIndex0"),
                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                    let foldableWithIndexCoproduct =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexCoproduct(),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FoldableWithIndex1"),
                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                    let traversableCoproduct =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traversableCoproduct(),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable2"),
                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                    &Func1::new({
                                                                                                                    let dictTraversableWithIndex
                                                                                                                        =
                                                                                                                        dictTraversableWithIndex.clone();
                                                                                                                    let foldableWithIndexCoproduct
                                                                                                                        =
                                                                                                                        foldableWithIndexCoproduct.clone();
                                                                                                                    let functorWithIndexCoproduct
                                                                                                                        =
                                                                                                                        functorWithIndexCoproduct.clone();
                                                                                                                    let traversableCoproduct
                                                                                                                        =
                                                                                                                        traversableCoproduct.clone();
                                                                                                                    move
                                                                                                                        |dictTraversableWithIndex1|
                                                                                                                        {
                                                                                                                            let functorWithIndexCoproduct1 =
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&functorWithIndexCoproduct,
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FunctorWithIndex0"),
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictTraversableWithIndex1)),
                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                            let foldableWithIndexCoproduct1 =
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&foldableWithIndexCoproduct,
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FoldableWithIndex1"),
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictTraversableWithIndex1)),
                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                            let traversableCoproduct1 =
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&traversableCoproduct,
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable2"),
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictTraversableWithIndex1)),
                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                                                                             &&&add(string("traverseWithIndex"),
                                                                                                                                                                    &&Func1::new({
                                                                                                                                                                                     let dictTraversableWithIndex1
                                                                                                                                                                                         =
                                                                                                                                                                                         dictTraversableWithIndex1.clone();
                                                                                                                                                                                     move
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
                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                                   |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                                                                                                   &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))))),
                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&dictTraversableWithIndex),
                                                                                                                                                                                                                                                                                                                                                                                              &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                 f),
                                                                                                                                                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_1.clone()))))))),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                              &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                |usd__arg1_2|
                                                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_2.clone())))))),
                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                              &&&dictTraversableWithIndex1),
                                                                                                                                                                                                                                                                                                                                                           &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                              f),
                                                                                                                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                             |usd__arg1_3|
                                                                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_3.clone())))))))
                                                                                                                                                                                                         })
                                                                                                                                                                                         }
                                                                                                                                                                                 }),
                                                                                                                                                                    add(string("FunctorWithIndex0"),
                                                                                                                                                                        &&Func1::new({
                                                                                                                                                                                         let functorWithIndexCoproduct1
                                                                                                                                                                                             =
                                                                                                                                                                                             functorWithIndexCoproduct1.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |usd__unused|
                                                                                                                                                                                             &functorWithIndexCoproduct1
                                                                                                                                                                                     }),
                                                                                                                                                                        add(string("FoldableWithIndex1"),
                                                                                                                                                                            &&Func1::new({
                                                                                                                                                                                             let foldableWithIndexCoproduct1
                                                                                                                                                                                                 =
                                                                                                                                                                                                 foldableWithIndexCoproduct1.clone();
                                                                                                                                                                                             move
                                                                                                                                                                                                 |usd__unused_1|
                                                                                                                                                                                                 &foldableWithIndexCoproduct1
                                                                                                                                                                                         }),
                                                                                                                                                                            add(string("Traversable2"),
                                                                                                                                                                                &&Func1::new({
                                                                                                                                                                                                 let traversableCoproduct1
                                                                                                                                                                                                     =
                                                                                                                                                                                                     traversableCoproduct1.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     |usd__unused_2|
                                                                                                                                                                                                     &traversableCoproduct1
                                                                                                                                                                                             }),
                                                                                                                                                                                empty::<string,
                                                                                                                                                                                        &dyn Any>())))))
                                                                                                                        }
                                                                                                                })
                                                                                                }))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexConst() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexConst:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexConst.get_or_init(||
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                             &&&add(string("traverseWithIndex"),
                                                                                                                    &&Func1::new(move
                                                                                                                                     |dictApplicative|
                                                                                                                                     &Func1::new({
                                                                                                                                                     let dictApplicative
                                                                                                                                                         =
                                                                                                                                                         dictApplicative.clone();
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
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                     &&&dictApplicative),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_Const(),
                                                                                                                                                                                                                                                     &&&matchValue_1))
                                                                                                                                                                             }
                                                                                                                                                                     })
                                                                                                                                                 })),
                                                                                                                    add(string("FunctorWithIndex0"),
                                                                                                                        &&Func1::new(move
                                                                                                                                         |usd__unused|
                                                                                                                                         &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexConst()),
                                                                                                                        add(string("FoldableWithIndex1"),
                                                                                                                            &&Func1::new(move
                                                                                                                                             |usd__unused_1|
                                                                                                                                             &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexConst()),
                                                                                                                            add(string("Traversable2"),
                                                                                                                                &&Func1::new(move
                                                                                                                                                 |usd__unused_2|
                                                                                                                                                 &PureScript_Data_Traversable::Data_Traversable_traversableConst()),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexConj() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexConj:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexConj.get_or_init(||
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                            &&&add(string("traverseWithIndex"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |dictApplicative|
                                                                                                                                    {
                                                                                                                                        let traverse8 =
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverse6(),
                                                                                                                                                                             dictApplicative);
                                                                                                                                        &Func1::new({
                                                                                                                                                        let traverse8
                                                                                                                                                            =
                                                                                                                                                            traverse8.clone();
                                                                                                                                                        move
                                                                                                                                                            |f|
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                &&&traverse8),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                    })
                                                                                                                                    }),
                                                                                                                   add(string("FunctorWithIndex0"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |usd__unused|
                                                                                                                                        &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexConj()),
                                                                                                                       add(string("FoldableWithIndex1"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |usd__unused_1|
                                                                                                                                            &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexConj()),
                                                                                                                           add(string("Traversable2"),
                                                                                                                               &&Func1::new(move
                                                                                                                                                |usd__unused_2|
                                                                                                                                                &PureScript_Data_Traversable::Data_Traversable_traversableConj()),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexCompose()
     -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexCompose:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexCompose.get_or_init(||
                                                                              &Func1::new(move
                                                                                              |dictTraversableWithIndex|
                                                                                              {
                                                                                                  let functorWithIndexCompose =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexCompose(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FunctorWithIndex0"),
                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                  let foldableWithIndexCompose =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexCompose(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FoldableWithIndex1"),
                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                  let traversableCompose =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traversableCompose(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable2"),
                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                  &Func1::new({
                                                                                                                  let dictTraversableWithIndex
                                                                                                                      =
                                                                                                                      dictTraversableWithIndex.clone();
                                                                                                                  let foldableWithIndexCompose
                                                                                                                      =
                                                                                                                      foldableWithIndexCompose.clone();
                                                                                                                  let functorWithIndexCompose
                                                                                                                      =
                                                                                                                      functorWithIndexCompose.clone();
                                                                                                                  let traversableCompose
                                                                                                                      =
                                                                                                                      traversableCompose.clone();
                                                                                                                  move
                                                                                                                      |dictTraversableWithIndex1|
                                                                                                                      {
                                                                                                                          let traverseWithIndex1 =
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                               dictTraversableWithIndex1);
                                                                                                                          let functorWithIndexCompose1 =
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&functorWithIndexCompose,
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FunctorWithIndex0"),
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictTraversableWithIndex1)),
                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                          let foldableWithIndexCompose1 =
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&foldableWithIndexCompose,
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FoldableWithIndex1"),
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictTraversableWithIndex1)),
                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                          let traversableCompose1 =
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&traversableCompose,
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable2"),
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictTraversableWithIndex1)),
                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                                                                           &&&add(string("traverseWithIndex"),
                                                                                                                                                                  &&Func1::new({
                                                                                                                                                                                   let traverseWithIndex1
                                                                                                                                                                                       =
                                                                                                                                                                                       traverseWithIndex1.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |dictApplicative|
                                                                                                                                                                                       {
                                                                                                                                                                                           let Functor0 =
                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                                                           let traverseWithIndex2 =
                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&traverseWithIndex1,
                                                                                                                                                                                                                                dictApplicative);
                                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                                           let Functor0
                                                                                                                                                                                                               =
                                                                                                                                                                                                               Functor0.clone();
                                                                                                                                                                                                           let dictApplicative
                                                                                                                                                                                                               =
                                                                                                                                                                                                               dictApplicative.clone();
                                                                                                                                                                                                           let traverseWithIndex2
                                                                                                                                                                                                               =
                                                                                                                                                                                                               traverseWithIndex2.clone();
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
                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Functor_Compose::Data_Functor_Compose_Compose())),
                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&dictTraversableWithIndex),
                                                                                                                                                                                                                                                                                                                                                                                 &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&traverseWithIndex2),
                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_curry(),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&matchValue))),
                                                                                                                                                                                                                                                                                                           &&&matchValue_1))
                                                                                                                                                                                                                                   }
                                                                                                                                                                                                                           })
                                                                                                                                                                                                       })
                                                                                                                                                                                       }
                                                                                                                                                                               }),
                                                                                                                                                                  add(string("FunctorWithIndex0"),
                                                                                                                                                                      &&Func1::new({
                                                                                                                                                                                       let functorWithIndexCompose1
                                                                                                                                                                                           =
                                                                                                                                                                                           functorWithIndexCompose1.clone();
                                                                                                                                                                                       move
                                                                                                                                                                                           |usd__unused|
                                                                                                                                                                                           &functorWithIndexCompose1
                                                                                                                                                                                   }),
                                                                                                                                                                      add(string("FoldableWithIndex1"),
                                                                                                                                                                          &&Func1::new({
                                                                                                                                                                                           let foldableWithIndexCompose1
                                                                                                                                                                                               =
                                                                                                                                                                                               foldableWithIndexCompose1.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |usd__unused_1|
                                                                                                                                                                                               &foldableWithIndexCompose1
                                                                                                                                                                                       }),
                                                                                                                                                                          add(string("Traversable2"),
                                                                                                                                                                              &&Func1::new({
                                                                                                                                                                                               let traversableCompose1
                                                                                                                                                                                                   =
                                                                                                                                                                                                   traversableCompose1.clone();
                                                                                                                                                                                               move
                                                                                                                                                                                                   |usd__unused_2|
                                                                                                                                                                                                   &traversableCompose1
                                                                                                                                                                                           }),
                                                                                                                                                                              empty::<string,
                                                                                                                                                                                      &dyn Any>())))))
                                                                                                                      }
                                                                                                              })
                                                                                              }))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexArray_004059()
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                         &&&add(string("traverseWithIndex"),
                                                &&Func1::new({
                                                                 let Data_TraversableWithIndex_traversableWithIndexArray_004059_002d1
                                                                     =
                                                                     Data_TraversableWithIndex_traversableWithIndexArray_004059_002d1.clone();
                                                                 move
                                                                     |dictApplicative|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndexDefault(),
                                                                                                                                         &&&Data_TraversableWithIndex_traversableWithIndexArray_004059_002d1.Value),
                                                                                                      dictApplicative)
                                                             }),
                                                add(string("FunctorWithIndex0"),
                                                    &&Func1::new(move
                                                                     |usd__unused|
                                                                     &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexArray()),
                                                    add(string("FoldableWithIndex1"),
                                                        &&Func1::new(move
                                                                         |usd__unused_1|
                                                                         &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexArray()),
                                                        add(string("Traversable2"),
                                                            &&Func1::new(move
                                                                             |usd__unused_2|
                                                                             &PureScript_Data_Traversable::Data_Traversable_traversableArray()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexArray_004059_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static
         Data_TraversableWithIndex_traversableWithIndexArray_004059_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexArray_004059_002d1.get_or_init(||
                                                                                         Lazy(Data_TraversableWithIndex_traversableWithIndexArray_004059.clone()))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexArray() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexArray.get_or_init(||
                                                                            Data_TraversableWithIndex_traversableWithIndexArray_004059_002d1.Value)
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexApp() -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexApp:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexApp.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |dictTraversableWithIndex|
                                                                                          {
                                                                                              let functorWithIndexApp =
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexApp(),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FunctorWithIndex0"),
                                                                                                                                                                             Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                              let foldableWithIndexApp =
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexApp(),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FoldableWithIndex1"),
                                                                                                                                                                             Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                              let traversableApp =
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traversableApp(),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable2"),
                                                                                                                                                                             Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                                               &&&add(string("traverseWithIndex"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let dictTraversableWithIndex
                                                                                                                                                           =
                                                                                                                                                           dictTraversableWithIndex.clone();
                                                                                                                                                       move
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
                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                  &&&Functor0),
                                                                                                                                                                                                                                                                               &&&PureScript_Data_Functor_App::Data_Functor_App_App()),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                        &&&dictTraversableWithIndex),
                                                                                                                                                                                                                                                                                                                                                     &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                                                                               &&&matchValue_1))
                                                                                                                                                                                                       }
                                                                                                                                                                                               })
                                                                                                                                                                           })
                                                                                                                                                           }
                                                                                                                                                   }),
                                                                                                                                      add(string("FunctorWithIndex0"),
                                                                                                                                          &&Func1::new({
                                                                                                                                                           let functorWithIndexApp
                                                                                                                                                               =
                                                                                                                                                               functorWithIndexApp.clone();
                                                                                                                                                           move
                                                                                                                                                               |usd__unused|
                                                                                                                                                               &functorWithIndexApp
                                                                                                                                                       }),
                                                                                                                                          add(string("FoldableWithIndex1"),
                                                                                                                                              &&Func1::new({
                                                                                                                                                               let foldableWithIndexApp
                                                                                                                                                                   =
                                                                                                                                                                   foldableWithIndexApp.clone();
                                                                                                                                                               move
                                                                                                                                                                   |usd__unused_1|
                                                                                                                                                                   &foldableWithIndexApp
                                                                                                                                                           }),
                                                                                                                                              add(string("Traversable2"),
                                                                                                                                                  &&Func1::new({
                                                                                                                                                                   let traversableApp
                                                                                                                                                                       =
                                                                                                                                                                       traversableApp.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |usd__unused_2|
                                                                                                                                                                       &traversableApp
                                                                                                                                                               }),
                                                                                                                                                  empty::<string,
                                                                                                                                                          &dyn Any>())))))
                                                                                          }))
    }
    pub fn Data_TraversableWithIndex_traversableWithIndexAdditive()
     -> &dyn Any {
        static Data_TraversableWithIndex_traversableWithIndexAdditive:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_traversableWithIndexAdditive.get_or_init(||
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                                &&&add(string("traverseWithIndex"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |dictApplicative|
                                                                                                                                        {
                                                                                                                                            let traverse8 =
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverse7(),
                                                                                                                                                                                 dictApplicative);
                                                                                                                                            &Func1::new({
                                                                                                                                                            let traverse8
                                                                                                                                                                =
                                                                                                                                                                traverse8.clone();
                                                                                                                                                            move
                                                                                                                                                                |f|
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                    &&&traverse8),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                                    &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                        })
                                                                                                                                        }),
                                                                                                                       add(string("FunctorWithIndex0"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |usd__unused|
                                                                                                                                            &PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_functorWithIndexAdditive()),
                                                                                                                           add(string("FoldableWithIndex1"),
                                                                                                                               &&Func1::new(move
                                                                                                                                                |usd__unused_1|
                                                                                                                                                &PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldableWithIndexAdditive()),
                                                                                                                               add(string("Traversable2"),
                                                                                                                                   &&Func1::new(move
                                                                                                                                                    |usd__unused_2|
                                                                                                                                                    &PureScript_Data_Traversable::Data_Traversable_traversableAdditive()),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>()))))))
    }
    pub fn Data_TraversableWithIndex_mapAccumRWithIndex() -> &dyn Any {
        static Data_TraversableWithIndex_mapAccumRWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_mapAccumRWithIndex.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictTraversableWithIndex|
                                                                                     &Func1::new({
                                                                                                     let dictTraversableWithIndex
                                                                                                         =
                                                                                                         dictTraversableWithIndex.clone();
                                                                                                     move
                                                                                                         |f|
                                                                                                         &Func1::new({
                                                                                                                         let f
                                                                                                                             =
                                                                                                                             f.clone();
                                                                                                                         move
                                                                                                                             |s0|
                                                                                                                             &Func1::new({
                                                                                                                                             let s0
                                                                                                                                                 =
                                                                                                                                                 s0.clone();
                                                                                                                                             move
                                                                                                                                                 |xs|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_stateR(),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictTraversableWithIndex),
                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_applicativeStateR()),
                                                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                                                             |i|
                                                                                                                                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                                                                                                                                             let i
                                                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                                                 i.clone();
                                                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                                                 |a|
                                                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateR(),
                                                                                                                                                                                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                    let a
                                                                                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                                                                                        a.clone();
                                                                                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                                                                                        |s|
                                                                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&i),
                                                                                                                                                                                                                                                                                                                                                                                                                                                            s),
                                                                                                                                                                                                                                                                                                                                                                                                                         &&&a)
                                                                                                                                                                                                                                                                                                                                                                                }))
                                                                                                                                                                                                                                                                                                                         }))),
                                                                                                                                                                                                                                                        xs)),
                                                                                                                                                                                  &&&s0)
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })))
    }
    pub fn Data_TraversableWithIndex_scanrWithIndex() -> &dyn Any {
        static Data_TraversableWithIndex_scanrWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_scanrWithIndex.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictTraversableWithIndex|
                                                                                 &Func1::new({
                                                                                                 let dictTraversableWithIndex
                                                                                                     =
                                                                                                     dictTraversableWithIndex.clone();
                                                                                                 move
                                                                                                     |f|
                                                                                                     &Func1::new({
                                                                                                                     let f
                                                                                                                         =
                                                                                                                         f.clone();
                                                                                                                     move
                                                                                                                         |b0|
                                                                                                                         &Func1::new({
                                                                                                                                         let b0
                                                                                                                                             =
                                                                                                                                             b0.clone();
                                                                                                                                         move
                                                                                                                                             |xs|
                                                                                                                                             find(string("value"),
                                                                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_mapAccumRWithIndex(),
                                                                                                                                                                                                                                                                                                                     &&&dictTraversableWithIndex),
                                                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                                                    |i|
                                                                                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                                                                                    let i
                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                        i.clone();
                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                        |b|
                                                                                                                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                                                                                                                        let b
                                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                                            b.clone();
                                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                                            |a|
                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                                let b_prime =
                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&i),
                                                                                                                                                                                                                                                                                                                                                                                                                        a),
                                                                                                                                                                                                                                                                                                                                                                                     &&&b);
                                                                                                                                                                                                                                                                                                                                                &add(string("accum"),
                                                                                                                                                                                                                                                                                                                                                     &&b_prime,
                                                                                                                                                                                                                                                                                                                                                     add(string("value"),
                                                                                                                                                                                                                                                                                                                                                         &&b_prime,
                                                                                                                                                                                                                                                                                                                                                         empty::<string,
                                                                                                                                                                                                                                                                                                                                                                 &dyn Any>()))
                                                                                                                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                                                                                                                }))),
                                                                                                                                                                                                                                               &&&b0),
                                                                                                                                                                                                            xs)))
                                                                                                                                     })
                                                                                                                 })
                                                                                             })))
    }
    pub fn Data_TraversableWithIndex_mapAccumLWithIndex() -> &dyn Any {
        static Data_TraversableWithIndex_mapAccumLWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_mapAccumLWithIndex.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictTraversableWithIndex|
                                                                                     &Func1::new({
                                                                                                     let dictTraversableWithIndex
                                                                                                         =
                                                                                                         dictTraversableWithIndex.clone();
                                                                                                     move
                                                                                                         |f|
                                                                                                         &Func1::new({
                                                                                                                         let f
                                                                                                                             =
                                                                                                                             f.clone();
                                                                                                                         move
                                                                                                                             |s0|
                                                                                                                             &Func1::new({
                                                                                                                                             let s0
                                                                                                                                                 =
                                                                                                                                                 s0.clone();
                                                                                                                                             move
                                                                                                                                                 |xs|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_stateL(),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictTraversableWithIndex),
                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_applicativeStateL()),
                                                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                                                             |i|
                                                                                                                                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                                                                                                                                             let i
                                                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                                                 i.clone();
                                                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                                                 |a|
                                                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateL(),
                                                                                                                                                                                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                    let a
                                                                                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                                                                                        a.clone();
                                                                                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                                                                                        |s|
                                                                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&i),
                                                                                                                                                                                                                                                                                                                                                                                                                                                            s),
                                                                                                                                                                                                                                                                                                                                                                                                                         &&&a)
                                                                                                                                                                                                                                                                                                                                                                                }))
                                                                                                                                                                                                                                                                                                                         }))),
                                                                                                                                                                                                                                                        xs)),
                                                                                                                                                                                  &&&s0)
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })))
    }
    pub fn Data_TraversableWithIndex_scanlWithIndex() -> &dyn Any {
        static Data_TraversableWithIndex_scanlWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_scanlWithIndex.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictTraversableWithIndex|
                                                                                 &Func1::new({
                                                                                                 let dictTraversableWithIndex
                                                                                                     =
                                                                                                     dictTraversableWithIndex.clone();
                                                                                                 move
                                                                                                     |f|
                                                                                                     &Func1::new({
                                                                                                                     let f
                                                                                                                         =
                                                                                                                         f.clone();
                                                                                                                     move
                                                                                                                         |b0|
                                                                                                                         &Func1::new({
                                                                                                                                         let b0
                                                                                                                                             =
                                                                                                                                             b0.clone();
                                                                                                                                         move
                                                                                                                                             |xs|
                                                                                                                                             find(string("value"),
                                                                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_mapAccumLWithIndex(),
                                                                                                                                                                                                                                                                                                                     &&&dictTraversableWithIndex),
                                                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                                                    |i|
                                                                                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                                                                                    let i
                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                        i.clone();
                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                        |b|
                                                                                                                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                                                                                                                        let b
                                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                                            b.clone();
                                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                                            |a|
                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                                let b_prime =
                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&i),
                                                                                                                                                                                                                                                                                                                                                                                                                        &&&b),
                                                                                                                                                                                                                                                                                                                                                                                     a);
                                                                                                                                                                                                                                                                                                                                                &add(string("accum"),
                                                                                                                                                                                                                                                                                                                                                     &&b_prime,
                                                                                                                                                                                                                                                                                                                                                     add(string("value"),
                                                                                                                                                                                                                                                                                                                                                         &&b_prime,
                                                                                                                                                                                                                                                                                                                                                         empty::<string,
                                                                                                                                                                                                                                                                                                                                                                 &dyn Any>()))
                                                                                                                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                                                                                                                }))),
                                                                                                                                                                                                                                               &&&b0),
                                                                                                                                                                                                            xs)))
                                                                                                                                     })
                                                                                                                 })
                                                                                             })))
    }
    pub fn Data_TraversableWithIndex_forWithIndex() -> &dyn Any {
        static Data_TraversableWithIndex_forWithIndex:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_TraversableWithIndex_forWithIndex.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictApplicative|
                                                                               &Func1::new({
                                                                                               let dictApplicative
                                                                                                   =
                                                                                                   dictApplicative.clone();
                                                                                               move
                                                                                                   |dictTraversableWithIndex|
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                          dictTraversableWithIndex),
                                                                                                                                                                       &&&dictApplicative))
                                                                                           })))
    }
}
