pub mod PureScript_Data_Functor_Product {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Functor_Product_unwrap() -> &dyn Any {
        static Data_Functor_Product_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_unwrap.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Functor_Product_unwrap1() -> &dyn Any {
        static Data_Functor_Product_unwrap1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_unwrap1.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Functor_Product_Product() -> &dyn Any {
        static Data_Functor_Product_Product: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_Product.get_or_init(||
                                                     &Func1::new(move |x|
                                                                     x.clone()))
    }
    pub fn Data_Functor_Product_showProduct() -> &dyn Any {
        static Data_Functor_Product_showProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_showProduct.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictShow|
                                                                         &Func1::new({
                                                                                         let dictShow
                                                                                             =
                                                                                             dictShow.clone();
                                                                                         move
                                                                                             |dictShow1|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                              &&&add(string("show"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let dictShow1
                                                                                                                                                          =
                                                                                                                                                          dictShow1.clone();
                                                                                                                                                      move
                                                                                                                                                          |v|
                                                                                                                                                          {
                                                                                                                                                              let matchValue:
                                                                                                                                                                      LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                     &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                  &&&string("(product ")),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                           &&&dictShow),
                                                                                                                                                                                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                           })),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                        &&&string(" ")),
                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                                 &&&dictShow1),
                                                                                                                                                                                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                                                                                                                                        &&&string(")")))))
                                                                                                                                                          }
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>()))
                                                                                     })))
    }
    pub fn Data_Functor_Product_product() -> &dyn Any {
        static Data_Functor_Product_product: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_product.get_or_init(||
                                                     &Func1::new(move |fa|
                                                                     &Func1::new({
                                                                                     let fa
                                                                                         =
                                                                                         fa.clone();
                                                                                     move
                                                                                         |ga|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_Product(),
                                                                                                                          &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&fa,
                                                                                                                                                                                    ga.clone())))
                                                                                 })))
    }
    pub fn Data_Functor_Product_newtypeProduct() -> &dyn Any {
        static Data_Functor_Product_newtypeProduct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Product_newtypeProduct.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                             &&&add(string("Coercible0"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &Sharpurs_Prelude::Prim_undefined()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Functor_Product_functorProduct() -> &dyn Any {
        static Data_Functor_Product_functorProduct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Product_functorProduct.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFunctor|
                                                                            &Func1::new({
                                                                                            let dictFunctor
                                                                                                =
                                                                                                dictFunctor.clone();
                                                                                            move
                                                                                                |dictFunctor1|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                                                 &&&add(string("map"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let dictFunctor1
                                                                                                                                                             =
                                                                                                                                                             dictFunctor1.clone();
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
                                                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&dictFunctor),
                                                                                                                                                                                                                                                                                                                                                                  &&&f1)),
                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                  &&&dictFunctor1),
                                                                                                                                                                                                                                                                                                                               &&&f1)),
                                                                                                                                                                                                                                                         &&&matchValue_1))
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                     }),
                                                                                                                                        empty::<string,
                                                                                                                                                &dyn Any>()))
                                                                                        })))
    }
    pub fn Data_Functor_Product_eq1Product() -> &dyn Any {
        static Data_Functor_Product_eq1Product: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_eq1Product.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictEq1|
                                                                        &Func1::new({
                                                                                        let dictEq1
                                                                                            =
                                                                                            dictEq1.clone();
                                                                                        move
                                                                                            |dictEq11|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                                                             &&&add(string("eq1"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let dictEq11
                                                                                                                                                         =
                                                                                                                                                         dictEq11.clone();
                                                                                                                                                     move
                                                                                                                                                         |dictEq|
                                                                                                                                                         &Func1::new({
                                                                                                                                                                         let dictEq
                                                                                                                                                                             =
                                                                                                                                                                             dictEq.clone();
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
                                                                                                                                                                                                             LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                     let matchValue_1:
                                                                                                                                                                                                             LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&&dictEq1),
                                                                                                                                                                                                                                                                                                                                                                                  &&&dictEq),
                                                                                                                                                                                                                                                                                                                                               &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                                                                                            &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                               })),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                                                                                                                                                                                                                                                                  &&&dictEq11),
                                                                                                                                                                                                                                                                                                                                               &&&dictEq),
                                                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                               }),
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
                                                                                                                                    empty::<string,
                                                                                                                                            &dyn Any>()))
                                                                                    })))
    }
    pub fn Data_Functor_Product_eqProduct() -> &dyn Any {
        static Data_Functor_Product_eqProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_eqProduct.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictEq1|
                                                                       {
                                                                           let eq1Product1 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_eq1Product(),
                                                                                                                dictEq1);
                                                                           &Func1::new({
                                                                                           let eq1Product1
                                                                                               =
                                                                                               eq1Product1.clone();
                                                                                           move
                                                                                               |dictEq11|
                                                                                               {
                                                                                                   let eq1 =
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&eq1Product1,
                                                                                                                                                                           dictEq11));
                                                                                                   &Func1::new({
                                                                                                                   let eq1
                                                                                                                       =
                                                                                                                       eq1.clone();
                                                                                                                   move
                                                                                                                       |dictEq|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                                                                        &&&add(string("eq"),
                                                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&&eq1,
                                                                                                                                                                                                 dictEq),
                                                                                                                                                               empty::<string,
                                                                                                                                                                       &dyn Any>()))
                                                                                                               })
                                                                                               }
                                                                                       })
                                                                       }))
    }
    pub fn Data_Functor_Product_ord1Product() -> &dyn Any {
        static Data_Functor_Product_ord1Product: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_ord1Product.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictOrd1|
                                                                         {
                                                                             let eq1Product1 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_eq1Product(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             &Func1::new({
                                                                                             let dictOrd1
                                                                                                 =
                                                                                                 dictOrd1.clone();
                                                                                             let eq1Product1
                                                                                                 =
                                                                                                 eq1Product1.clone();
                                                                                             move
                                                                                                 |dictOrd11|
                                                                                                 {
                                                                                                     let eq1Product2 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&eq1Product1,
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                                                    Sharpurs_Prelude::unbox(dictOrd11)),
                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                                                                      &&&add(string("compare1"),
                                                                                                                                             &&Func1::new({
                                                                                                                                                              let dictOrd11
                                                                                                                                                                  =
                                                                                                                                                                  dictOrd11.clone();
                                                                                                                                                              move
                                                                                                                                                                  |dictOrd|
                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                  let dictOrd
                                                                                                                                                                                      =
                                                                                                                                                                                      dictOrd.clone();
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
                                                                                                                                                                                                                      LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                              let matchValue_1:
                                                                                                                                                                                                                      LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                              let matchValue_3:
                                                                                                                                                                                                                      LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                                                                                                                                                                                                                                                                      &&&dictOrd1),
                                                                                                                                                                                                                                                                                                                                                   &&&dictOrd),
                                                                                                                                                                                                                                                                                                                &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                   }),
                                                                                                                                                                                                                                                                             &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                }));
                                                                                                                                                                                                              if let Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                                                                                                     =
                                                                                                                                                                                                                     matchValue_3.as_ref()
                                                                                                                                                                                                                 {
                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                                                                                                                                                                                                                                            &&&dictOrd11),
                                                                                                                                                                                                                                                                                                                         &&&dictOrd),
                                                                                                                                                                                                                                                                                      &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                   &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                      })
                                                                                                                                                                                                              } else {
                                                                                                                                                                                                                  &matchValue_3
                                                                                                                                                                                                              }
                                                                                                                                                                                                          }
                                                                                                                                                                                                  })
                                                                                                                                                                              })
                                                                                                                                                          }),
                                                                                                                                             add(string("Eq10"),
                                                                                                                                                 &&Func1::new({
                                                                                                                                                                  let eq1Product2
                                                                                                                                                                      =
                                                                                                                                                                      eq1Product2.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |usd__unused|
                                                                                                                                                                      &eq1Product2
                                                                                                                                                              }),
                                                                                                                                                 empty::<string,
                                                                                                                                                         &dyn Any>())))
                                                                                                 }
                                                                                         })
                                                                         }))
    }
    pub fn Data_Functor_Product_ordProduct() -> &dyn Any {
        static Data_Functor_Product_ordProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_ordProduct.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictOrd1|
                                                                        {
                                                                            let ord1Product1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_ord1Product(),
                                                                                                                 dictOrd1);
                                                                            let eqProduct1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_eqProduct(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            &Func1::new({
                                                                                            let eqProduct1
                                                                                                =
                                                                                                eqProduct1.clone();
                                                                                            let ord1Product1
                                                                                                =
                                                                                                ord1Product1.clone();
                                                                                            move
                                                                                                |dictOrd11|
                                                                                                {
                                                                                                    let compare1 =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&ord1Product1,
                                                                                                                                                                            dictOrd11));
                                                                                                    let eqProduct2 =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&eqProduct1,
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictOrd11)),
                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                    &Func1::new({
                                                                                                                    let compare1
                                                                                                                        =
                                                                                                                        compare1.clone();
                                                                                                                    let eqProduct2
                                                                                                                        =
                                                                                                                        eqProduct2.clone();
                                                                                                                    move
                                                                                                                        |dictOrd|
                                                                                                                        {
                                                                                                                            let eqProduct3 =
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&eqProduct2,
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                                                             &&&add(string("compare"),
                                                                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&compare1,
                                                                                                                                                                                                      dictOrd),
                                                                                                                                                                    add(string("Eq0"),
                                                                                                                                                                        &&Func1::new({
                                                                                                                                                                                         let eqProduct3
                                                                                                                                                                                             =
                                                                                                                                                                                             eqProduct3.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |usd__unused|
                                                                                                                                                                                             &eqProduct3
                                                                                                                                                                                     }),
                                                                                                                                                                        empty::<string,
                                                                                                                                                                                &dyn Any>())))
                                                                                                                        }
                                                                                                                })
                                                                                                }
                                                                                        })
                                                                        }))
    }
    pub fn Data_Functor_Product_bihoistProduct() -> &dyn Any {
        static Data_Functor_Product_bihoistProduct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Product_bihoistProduct.get_or_init(||
                                                            &Func1::new(move
                                                                            |natF|
                                                                            &Func1::new({
                                                                                            let natF
                                                                                                =
                                                                                                natF.clone();
                                                                                            move
                                                                                                |natG|
                                                                                                &Func1::new({
                                                                                                                let natG
                                                                                                                    =
                                                                                                                    natG.clone();
                                                                                                                move
                                                                                                                    |v|
                                                                                                                    {
                                                                                                                        let matchValue =
                                                                                                                            Sharpurs_Prelude::unbox(&&natF);
                                                                                                                        let matchValue_1 =
                                                                                                                            Sharpurs_Prelude::unbox(&&natG);
                                                                                                                        let matchValue_2 =
                                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_Product(),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorTuple()),
                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                               &&&matchValue_1),
                                                                                                                                                                                            &&&matchValue_2))
                                                                                                                    }
                                                                                                            })
                                                                                        })))
    }
    pub fn Data_Functor_Product_applyProduct() -> &dyn Any {
        static Data_Functor_Product_applyProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_applyProduct.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictApply|
                                                                          {
                                                                              let functorProduct1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_functorProduct(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                              &Func1::new({
                                                                                              let dictApply
                                                                                                  =
                                                                                                  dictApply.clone();
                                                                                              let functorProduct1
                                                                                                  =
                                                                                                  functorProduct1.clone();
                                                                                              move
                                                                                                  |dictApply1|
                                                                                                  {
                                                                                                      let functorProduct2 =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&functorProduct1,
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictApply1)),
                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                                                       &&&add(string("apply"),
                                                                                                                                              &&Func1::new({
                                                                                                                                                               let dictApply1
                                                                                                                                                                   =
                                                                                                                                                                   dictApply1.clone();
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
                                                                                                                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                           let matchValue_1:
                                                                                                                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_product(),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                        &&&dictApply),
                                                                                                                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                     })),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                     &&&dictApply1),
                                                                                                                                                                                                                                                                                                  &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                  }))
                                                                                                                                                                                       }
                                                                                                                                                                               })
                                                                                                                                                           }),
                                                                                                                                              add(string("Functor0"),
                                                                                                                                                  &&Func1::new({
                                                                                                                                                                   let functorProduct2
                                                                                                                                                                       =
                                                                                                                                                                       functorProduct2.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |usd__unused|
                                                                                                                                                                       &functorProduct2
                                                                                                                                                               }),
                                                                                                                                                  empty::<string,
                                                                                                                                                          &dyn Any>())))
                                                                                                  }
                                                                                          })
                                                                          }))
    }
    pub fn Data_Functor_Product_bindProduct() -> &dyn Any {
        static Data_Functor_Product_bindProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_bindProduct.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictBind|
                                                                         {
                                                                             let applyProduct1 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_applyProduct(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictBind)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             &Func1::new({
                                                                                             let applyProduct1
                                                                                                 =
                                                                                                 applyProduct1.clone();
                                                                                             let dictBind
                                                                                                 =
                                                                                                 dictBind.clone();
                                                                                             move
                                                                                                 |dictBind1|
                                                                                                 {
                                                                                                     let applyProduct2 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&applyProduct1,
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                    Sharpurs_Prelude::unbox(dictBind1)),
                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                                                                      &&&add(string("bind"),
                                                                                                                                             &&Func1::new({
                                                                                                                                                              let dictBind1
                                                                                                                                                                  =
                                                                                                                                                                  dictBind1.clone();
                                                                                                                                                              move
                                                                                                                                                                  |v|
                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                  let v
                                                                                                                                                                                      =
                                                                                                                                                                                      v.clone();
                                                                                                                                                                                  move
                                                                                                                                                                                      |f|
                                                                                                                                                                                      {
                                                                                                                                                                                          let matchValue:
                                                                                                                                                                                                  LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                              Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                          let f1 =
                                                                                                                                                                                              Sharpurs_Prelude::unbox(f);
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_product(),
                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                       &&&dictBind),
                                                                                                                                                                                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Tuple::Data_Tuple_fst()),
                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Functor_Product::Data_Functor_Product_unwrap()),
                                                                                                                                                                                                                                                                                                                                                                       &&&f1)))),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                    &&&dictBind1),
                                                                                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Functor_Product::Data_Functor_Product_unwrap1()),
                                                                                                                                                                                                                                                                                                                                    &&&f1))))
                                                                                                                                                                                      }
                                                                                                                                                                              })
                                                                                                                                                          }),
                                                                                                                                             add(string("Apply0"),
                                                                                                                                                 &&Func1::new({
                                                                                                                                                                  let applyProduct2
                                                                                                                                                                      =
                                                                                                                                                                      applyProduct2.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |usd__unused|
                                                                                                                                                                      &applyProduct2
                                                                                                                                                              }),
                                                                                                                                                 empty::<string,
                                                                                                                                                         &dyn Any>())))
                                                                                                 }
                                                                                         })
                                                                         }))
    }
    pub fn Data_Functor_Product_applicativeProduct() -> &dyn Any {
        static Data_Functor_Product_applicativeProduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_applicativeProduct.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictApplicative|
                                                                                {
                                                                                    let applyProduct1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_applyProduct(),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                   Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                    &Func1::new({
                                                                                                    let applyProduct1
                                                                                                        =
                                                                                                        applyProduct1.clone();
                                                                                                    let dictApplicative
                                                                                                        =
                                                                                                        dictApplicative.clone();
                                                                                                    move
                                                                                                        |dictApplicative1|
                                                                                                        {
                                                                                                            let applyProduct2 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&applyProduct1,
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictApplicative1)),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                                                             &&&add(string("pure"),
                                                                                                                                                    &&Func1::new({
                                                                                                                                                                     let dictApplicative1
                                                                                                                                                                         =
                                                                                                                                                                         dictApplicative1.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |a|
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_product(),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                   &&&dictApplicative),
                                                                                                                                                                                                                                                                                a)),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                &&&dictApplicative1),
                                                                                                                                                                                                                                             a))
                                                                                                                                                                 }),
                                                                                                                                                    add(string("Apply0"),
                                                                                                                                                        &&Func1::new({
                                                                                                                                                                         let applyProduct2
                                                                                                                                                                             =
                                                                                                                                                                             applyProduct2.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |usd__unused|
                                                                                                                                                                             &applyProduct2
                                                                                                                                                                     }),
                                                                                                                                                        empty::<string,
                                                                                                                                                                &dyn Any>())))
                                                                                                        }
                                                                                                })
                                                                                }))
    }
    pub fn Data_Functor_Product_monadProduct() -> &dyn Any {
        static Data_Functor_Product_monadProduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Product_monadProduct.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonad|
                                                                          {
                                                                              let applicativeProduct1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_applicativeProduct(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                              let bindProduct1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_bindProduct(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                             Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                              &Func1::new({
                                                                                              let applicativeProduct1
                                                                                                  =
                                                                                                  applicativeProduct1.clone();
                                                                                              let bindProduct1
                                                                                                  =
                                                                                                  bindProduct1.clone();
                                                                                              move
                                                                                                  |dictMonad1|
                                                                                                  {
                                                                                                      let applicativeProduct2 =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&applicativeProduct1,
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonad1)),
                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                      let bindProduct2 =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&bindProduct1,
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonad1)),
                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                                                                       &&&add(string("Applicative0"),
                                                                                                                                              &&Func1::new({
                                                                                                                                                               let applicativeProduct2
                                                                                                                                                                   =
                                                                                                                                                                   applicativeProduct2.clone();
                                                                                                                                                               move
                                                                                                                                                                   |usd__unused|
                                                                                                                                                                   &applicativeProduct2
                                                                                                                                                           }),
                                                                                                                                              add(string("Bind1"),
                                                                                                                                                  &&Func1::new({
                                                                                                                                                                   let bindProduct2
                                                                                                                                                                       =
                                                                                                                                                                       bindProduct2.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |usd__unused_1|
                                                                                                                                                                       &bindProduct2
                                                                                                                                                               }),
                                                                                                                                                  empty::<string,
                                                                                                                                                          &dyn Any>())))
                                                                                                  }
                                                                                          })
                                                                          }))
    }
}
