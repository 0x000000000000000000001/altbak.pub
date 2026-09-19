pub mod PureScript_Data_Bitraversable {
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
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_637e1ff5::PureScript_Data_Bifoldable;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_a8445950::PureScript_Data_Const;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_58c69d5::PureScript_Data_Functor_Clown;
    use crate::module_f8895b9f::PureScript_Data_Functor_Flip;
    use crate::module_8377b1b5::PureScript_Data_Functor_Joker;
    use crate::module_3079b2b5::PureScript_Data_Functor_Product2;
    use crate::module_3079b2b5::PureScript_Data_Functor_Product2::Data_Functor_Product2_Product2;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Bitraversable_identity() -> &dyn Any {
        static Data_Bitraversable_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_identity.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                     &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Bitraversable_identity1() -> &dyn Any {
        static Data_Bitraversable_identity1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_identity1.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                      &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Bitraversable_Bitraversableusd_Dict() -> &dyn Any {
        static Data_Bitraversable_Bitraversableusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_Bitraversableusd_Dict.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |x|
                                                                                 x.clone()))
    }
    pub fn Data_Bitraversable_bitraverse() -> &dyn Any {
        static Data_Bitraversable_bitraverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_bitraverse.get_or_init(||
                                                      &Func1::new(move |dict|
                                                                      find(string("bitraverse"),
                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bitraversable_lfor() -> &dyn Any {
        static Data_Bitraversable_lfor: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_lfor.get_or_init(||
                                                &Func1::new(move
                                                                |dictBitraversable|
                                                                &Func1::new({
                                                                                let dictBitraversable
                                                                                    =
                                                                                    dictBitraversable.clone();
                                                                                move
                                                                                    |dictApplicative|
                                                                                    {
                                                                                        let pure_var =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                             dictApplicative);
                                                                                        &Func1::new({
                                                                                                        let dictApplicative
                                                                                                            =
                                                                                                            dictApplicative.clone();
                                                                                                        let pure_var
                                                                                                            =
                                                                                                            pure_var.clone();
                                                                                                        move
                                                                                                            |t|
                                                                                                            &Func1::new({
                                                                                                                            let t
                                                                                                                                =
                                                                                                                                t.clone();
                                                                                                                            move
                                                                                                                                |f|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                                                                                                                             &&&dictBitraversable),
                                                                                                                                                                                                                                                                          &&&dictApplicative),
                                                                                                                                                                                                                                       f),
                                                                                                                                                                                                    &&&pure_var),
                                                                                                                                                                 &&&t)
                                                                                                                        })
                                                                                                    })
                                                                                    }
                                                                            })))
    }
    pub fn Data_Bitraversable_ltraverse() -> &dyn Any {
        static Data_Bitraversable_ltraverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_ltraverse.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBitraversable|
                                                                     &Func1::new({
                                                                                     let dictBitraversable
                                                                                         =
                                                                                         dictBitraversable.clone();
                                                                                     move
                                                                                         |dictApplicative|
                                                                                         {
                                                                                             let pure_var =
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                  dictApplicative);
                                                                                             &Func1::new({
                                                                                                             let dictApplicative
                                                                                                                 =
                                                                                                                 dictApplicative.clone();
                                                                                                             let pure_var
                                                                                                                 =
                                                                                                                 pure_var.clone();
                                                                                                             move
                                                                                                                 |f|
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                                                                           &&&dictBitraversable),
                                                                                                                                                                                                                        &&&dictApplicative),
                                                                                                                                                                                     f),
                                                                                                                                                  &&&pure_var)
                                                                                                         })
                                                                                         }
                                                                                 })))
    }
    pub fn Data_Bitraversable_rfor() -> &dyn Any {
        static Data_Bitraversable_rfor: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_rfor.get_or_init(||
                                                &Func1::new(move
                                                                |dictBitraversable|
                                                                &Func1::new({
                                                                                let dictBitraversable
                                                                                    =
                                                                                    dictBitraversable.clone();
                                                                                move
                                                                                    |dictApplicative|
                                                                                    {
                                                                                        let pure_var =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                             dictApplicative);
                                                                                        &Func1::new({
                                                                                                        let dictApplicative
                                                                                                            =
                                                                                                            dictApplicative.clone();
                                                                                                        let pure_var
                                                                                                            =
                                                                                                            pure_var.clone();
                                                                                                        move
                                                                                                            |t|
                                                                                                            &Func1::new({
                                                                                                                            let t
                                                                                                                                =
                                                                                                                                t.clone();
                                                                                                                            move
                                                                                                                                |f|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                                                                                                                             &&&dictBitraversable),
                                                                                                                                                                                                                                                                          &&&dictApplicative),
                                                                                                                                                                                                                                       &&&pure_var),
                                                                                                                                                                                                    f),
                                                                                                                                                                 &&&t)
                                                                                                                        })
                                                                                                    })
                                                                                    }
                                                                            })))
    }
    pub fn Data_Bitraversable_rtraverse() -> &dyn Any {
        static Data_Bitraversable_rtraverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_rtraverse.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBitraversable|
                                                                     &Func1::new({
                                                                                     let dictBitraversable
                                                                                         =
                                                                                         dictBitraversable.clone();
                                                                                     move
                                                                                         |dictApplicative|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                &&&dictBitraversable),
                                                                                                                                                             dictApplicative),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                             dictApplicative))
                                                                                 })))
    }
    pub fn Data_Bitraversable_bitraversableTuple() -> &dyn Any {
        static Data_Bitraversable_bitraversableTuple:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_bitraversableTuple.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_Bitraversableusd_Dict(),
                                                                                               &&&add(string("bitraverse"),
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
                                                                                                                                           move
                                                                                                                                               |f|
                                                                                                                                               &Func1::new({
                                                                                                                                                               let f
                                                                                                                                                                   =
                                                                                                                                                                   f.clone();
                                                                                                                                                               move
                                                                                                                                                                   |g|
                                                                                                                                                                   &Func1::new({
                                                                                                                                                                                   let g
                                                                                                                                                                                       =
                                                                                                                                                                                       g.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |v|
                                                                                                                                                                                       {
                                                                                                                                                                                           let matchValue =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                           let matchValue_1 =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&g);
                                                                                                                                                                                           let matchValue_2:
                                                                                                                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                  &&&Apply0),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                        &&&Functor0),
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
                                                                                                                                                                                                                                                                                                                                                                  }))),
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                     &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                        }))),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                               &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                  }))
                                                                                                                                                                                       }
                                                                                                                                                                               })
                                                                                                                                                           })
                                                                                                                                       })
                                                                                                                       }),
                                                                                                      add(string("bisequence"),
                                                                                                          &&Func1::new(move
                                                                                                                           |dictApplicative_1|
                                                                                                                           {
                                                                                                                               let Apply0_1 =
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                           Sharpurs_Prelude::unbox(dictApplicative_1)),
                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                               let Functor0_1 =
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictApplicative_1)),
                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                               &Func1::new({
                                                                                                                                               let Apply0_1
                                                                                                                                                   =
                                                                                                                                                   Apply0_1.clone();
                                                                                                                                               let Functor0_1
                                                                                                                                                   =
                                                                                                                                                   Functor0_1.clone();
                                                                                                                                               move
                                                                                                                                                   |v_1|
                                                                                                                                                   {
                                                                                                                                                       let matchValue_4:
                                                                                                                                                               LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                           Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                              &&&Apply0_1),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                    &&&Functor0_1),
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
                                                                                                                                                                                                                                                                                                                              }))),
                                                                                                                                                                                                                                                              &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                        &&&match matchValue_4.as_ref()
                                                                                                                                                                                               {
                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                          x)
                                                                                                                                                                                               =>
                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                           })
                                                                                                                                                   }
                                                                                                                                           })
                                                                                                                           }),
                                                                                                          add(string("Bifunctor0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorTuple()),
                                                                                                              add(string("Bifoldable1"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |usd__unused_1|
                                                                                                                                   &PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableTuple()),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))))
    }
    pub fn Data_Bitraversable_bitraversableJoker() -> &dyn Any {
        static Data_Bitraversable_bitraversableJoker:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_bitraversableJoker.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictTraversable|
                                                                              {
                                                                                  let bifunctorJoker =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Joker::Data_Functor_Joker_bifunctorJoker(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  let bifoldableJoker =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableJoker(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_Bitraversableusd_Dict(),
                                                                                                                   &&&add(string("bitraverse"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictTraversable
                                                                                                                                               =
                                                                                                                                               dictTraversable.clone();
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
                                                                                                                                                                       |v|
                                                                                                                                                                       &Func1::new({
                                                                                                                                                                                       let v
                                                                                                                                                                                           =
                                                                                                                                                                                           v.clone();
                                                                                                                                                                                       move
                                                                                                                                                                                           |r|
                                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                                           let r
                                                                                                                                                                                                               =
                                                                                                                                                                                                               r.clone();
                                                                                                                                                                                                           move
                                                                                                                                                                                                               |v1|
                                                                                                                                                                                                               {
                                                                                                                                                                                                                   let matchValue =
                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                                   let matchValue_1 =
                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&r);
                                                                                                                                                                                                                   let matchValue_2 =
                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker()),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&dictTraversable),
                                                                                                                                                                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                          &&&matchValue_1),
                                                                                                                                                                                                                                                                                       &&&matchValue_2))
                                                                                                                                                                                                               }
                                                                                                                                                                                                       })
                                                                                                                                                                                   })
                                                                                                                                                               })
                                                                                                                                               }
                                                                                                                                       }),
                                                                                                                          add(string("bisequence"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let dictTraversable
                                                                                                                                                   =
                                                                                                                                                   dictTraversable.clone();
                                                                                                                                               move
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
                                                                                                                                                                               let f_1 =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                      &&&Functor0_1),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Functor_Joker::Data_Functor_Joker_Joker()),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                                                                                         &&&dictTraversable),
                                                                                                                                                                                                                                                                                      &&&dictApplicative_1),
                                                                                                                                                                                                                                                   &&&f_1))
                                                                                                                                                                           }
                                                                                                                                                                   })
                                                                                                                                                   }
                                                                                                                                           }),
                                                                                                                              add(string("Bifunctor0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let bifunctorJoker
                                                                                                                                                       =
                                                                                                                                                       bifunctorJoker.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &bifunctorJoker
                                                                                                                                               }),
                                                                                                                                  add(string("Bifoldable1"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let bifoldableJoker
                                                                                                                                                           =
                                                                                                                                                           bifoldableJoker.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused_1|
                                                                                                                                                           &bifoldableJoker
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))))
                                                                              }))
    }
    pub fn Data_Bitraversable_bitraversableEither() -> &dyn Any {
        static Data_Bitraversable_bitraversableEither:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_bitraversableEither.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_Bitraversableusd_Dict(),
                                                                                                &&&add(string("bitraverse"),
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
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                       &&&Functor0),
                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                      |usd__arg1_1|
                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                    &&matchValue_2_1_0)),
                                                                                                                                                                                                Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_2_0_0)
                                                                                                                                                                                                =>
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                       &&&Functor0),
                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                      |usd__arg1|
                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                    &&matchValue_2_0_0)),
                                                                                                                                                                                            }
                                                                                                                                                                                        }
                                                                                                                                                                                })
                                                                                                                                                            })
                                                                                                                                        })
                                                                                                                        }),
                                                                                                       add(string("bisequence"),
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
                                                                                                                                                    |v_1|
                                                                                                                                                    {
                                                                                                                                                        let matchValue_4:
                                                                                                                                                                LrcPtr<Data_Either_Either> =
                                                                                                                                                            Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                        match matchValue_4.as_ref()
                                                                                                                                                            {
                                                                                                                                                            Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_4_1_0)
                                                                                                                                                            =>
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                   &&&Functor0_1),
                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                  |usd__arg1_3|
                                                                                                                                                                                                                                                  &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_3.clone())))),
                                                                                                                                                                                             &&matchValue_4_1_0),
                                                                                                                                                            Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_4_0_0)
                                                                                                                                                            =>
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                   &&&Functor0_1),
                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                  |usd__arg1_2|
                                                                                                                                                                                                                                                  &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_2.clone())))),
                                                                                                                                                                                             &&matchValue_4_0_0),
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                            })
                                                                                                                            }),
                                                                                                           add(string("Bifunctor0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorEither()),
                                                                                                               add(string("Bifoldable1"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |usd__unused_1|
                                                                                                                                    &PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableEither()),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))))
    }
    pub fn Data_Bitraversable_bitraversableConst() -> &dyn Any {
        static Data_Bitraversable_bitraversableConst:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_bitraversableConst.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_Bitraversableusd_Dict(),
                                                                                               &&&add(string("bitraverse"),
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
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                  &&&Functor0),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Const::Data_Const_Const()),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                               &&&matchValue_2))
                                                                                                                                                                                       }
                                                                                                                                                                               })
                                                                                                                                                           })
                                                                                                                                       })
                                                                                                                       }),
                                                                                                      add(string("bisequence"),
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
                                                                                                                                                   |v_1|
                                                                                                                                                   {
                                                                                                                                                       let a_1 =
                                                                                                                                                           Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                              &&&Functor0_1),
                                                                                                                                                                                                                           &&&PureScript_Data_Const::Data_Const_Const()),
                                                                                                                                                                                        &&&a_1)
                                                                                                                                                   }
                                                                                                                                           })
                                                                                                                           }),
                                                                                                          add(string("Bifunctor0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorConst()),
                                                                                                              add(string("Bifoldable1"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |usd__unused_1|
                                                                                                                                   &PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableConst()),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))))
    }
    pub fn Data_Bitraversable_bitraversableClown() -> &dyn Any {
        static Data_Bitraversable_bitraversableClown:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_bitraversableClown.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictTraversable|
                                                                              {
                                                                                  let bifunctorClown =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Clown::Data_Functor_Clown_bifunctorClown(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  let bifoldableClown =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableClown(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_Bitraversableusd_Dict(),
                                                                                                                   &&&add(string("bitraverse"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictTraversable
                                                                                                                                               =
                                                                                                                                               dictTraversable.clone();
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
                                                                                                                                                                       |l|
                                                                                                                                                                       &Func1::new({
                                                                                                                                                                                       let l
                                                                                                                                                                                           =
                                                                                                                                                                                           l.clone();
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
                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&l);
                                                                                                                                                                                                                   let matchValue_1 =
                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                                   let matchValue_2 =
                                                                                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Functor_Clown::Data_Functor_Clown_Clown()),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&dictTraversable),
                                                                                                                                                                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                                                                                                       &&&matchValue_2))
                                                                                                                                                                                                               }
                                                                                                                                                                                                       })
                                                                                                                                                                                   })
                                                                                                                                                               })
                                                                                                                                               }
                                                                                                                                       }),
                                                                                                                          add(string("bisequence"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let dictTraversable
                                                                                                                                                   =
                                                                                                                                                   dictTraversable.clone();
                                                                                                                                               move
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
                                                                                                                                                                               let f_1 =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                      &&&Functor0_1),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Functor_Clown::Data_Functor_Clown_Clown()),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                                                                                         &&&dictTraversable),
                                                                                                                                                                                                                                                                                      &&&dictApplicative_1),
                                                                                                                                                                                                                                                   &&&f_1))
                                                                                                                                                                           }
                                                                                                                                                                   })
                                                                                                                                                   }
                                                                                                                                           }),
                                                                                                                              add(string("Bifunctor0"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let bifunctorClown
                                                                                                                                                       =
                                                                                                                                                       bifunctorClown.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &bifunctorClown
                                                                                                                                               }),
                                                                                                                                  add(string("Bifoldable1"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let bifoldableClown
                                                                                                                                                           =
                                                                                                                                                           bifoldableClown.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused_1|
                                                                                                                                                           &bifoldableClown
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))))
                                                                              }))
    }
    pub fn Data_Bitraversable_bisequenceDefault() -> &dyn Any {
        static Data_Bitraversable_bisequenceDefault: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Bitraversable_bisequenceDefault.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictBitraversable|
                                                                             &Func1::new({
                                                                                             let dictBitraversable
                                                                                                 =
                                                                                                 dictBitraversable.clone();
                                                                                             move
                                                                                                 |dictApplicative|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                                                           &&&dictBitraversable),
                                                                                                                                                                                                        dictApplicative),
                                                                                                                                                                     &&&PureScript_Data_Bitraversable::Data_Bitraversable_identity()),
                                                                                                                                  &&&PureScript_Data_Bitraversable::Data_Bitraversable_identity1())
                                                                                         })))
    }
    pub fn Data_Bitraversable_bisequence() -> &dyn Any {
        static Data_Bitraversable_bisequence: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_bisequence.get_or_init(||
                                                      &Func1::new(move |dict|
                                                                      find(string("bisequence"),
                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bitraversable_bitraversableFlip() -> &dyn Any {
        static Data_Bitraversable_bitraversableFlip: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Bitraversable_bitraversableFlip.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictBitraversable|
                                                                             {
                                                                                 let bifunctorFlip =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_bifunctorFlip(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictBitraversable)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 let bifoldableFlip =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableFlip(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifoldable1"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictBitraversable)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_Bitraversableusd_Dict(),
                                                                                                                  &&&add(string("bitraverse"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let dictBitraversable
                                                                                                                                              =
                                                                                                                                              dictBitraversable.clone();
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
                                                                                                                                                                      |r|
                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                      let r
                                                                                                                                                                                          =
                                                                                                                                                                                          r.clone();
                                                                                                                                                                                      move
                                                                                                                                                                                          |l|
                                                                                                                                                                                          &Func1::new({
                                                                                                                                                                                                          let l
                                                                                                                                                                                                              =
                                                                                                                                                                                                              l.clone();
                                                                                                                                                                                                          move
                                                                                                                                                                                                              |v|
                                                                                                                                                                                                              {
                                                                                                                                                                                                                  let matchValue =
                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&r);
                                                                                                                                                                                                                  let matchValue_1 =
                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&l);
                                                                                                                                                                                                                  let matchValue_2 =
                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                         &&&Functor0),
                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Functor_Flip::Data_Functor_Flip_Flip()),
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&dictBitraversable),
                                                                                                                                                                                                                                                                                                                                                                                               &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                                            &&&matchValue_1),
                                                                                                                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                                                                                                                      &&&matchValue_2))
                                                                                                                                                                                                              }
                                                                                                                                                                                                      })
                                                                                                                                                                                  })
                                                                                                                                                              })
                                                                                                                                              }
                                                                                                                                      }),
                                                                                                                         add(string("bisequence"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let dictBitraversable
                                                                                                                                                  =
                                                                                                                                                  dictBitraversable.clone();
                                                                                                                                              move
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
                                                                                                                                                                              let p_1 =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                     &&&Functor0_1),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Functor_Flip::Data_Functor_Flip_Flip()),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bisequence(),
                                                                                                                                                                                                                                                                                                                        &&&dictBitraversable),
                                                                                                                                                                                                                                                                                     &&&dictApplicative_1),
                                                                                                                                                                                                                                                  &&&p_1))
                                                                                                                                                                          }
                                                                                                                                                                  })
                                                                                                                                                  }
                                                                                                                                          }),
                                                                                                                             add(string("Bifunctor0"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let bifunctorFlip
                                                                                                                                                      =
                                                                                                                                                      bifunctorFlip.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused|
                                                                                                                                                      &bifunctorFlip
                                                                                                                                              }),
                                                                                                                                 add(string("Bifoldable1"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let bifoldableFlip
                                                                                                                                                          =
                                                                                                                                                          bifoldableFlip.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused_1|
                                                                                                                                                          &bifoldableFlip
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))))
                                                                             }))
    }
    pub fn Data_Bitraversable_bitraversableProduct2() -> &dyn Any {
        static Data_Bitraversable_bitraversableProduct2:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_bitraversableProduct2.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictBitraversable|
                                                                                 {
                                                                                     let bifunctorProduct2 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product2::Data_Functor_Product2_bifunctorProduct2(),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                                                                    Sharpurs_Prelude::unbox(dictBitraversable)),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                     let bifoldableProduct2 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldableProduct2(),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifoldable1"),
                                                                                                                                                                    Sharpurs_Prelude::unbox(dictBitraversable)),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                     &Func1::new({
                                                                                                     let bifoldableProduct2
                                                                                                         =
                                                                                                         bifoldableProduct2.clone();
                                                                                                     let bifunctorProduct2
                                                                                                         =
                                                                                                         bifunctorProduct2.clone();
                                                                                                     let dictBitraversable
                                                                                                         =
                                                                                                         dictBitraversable.clone();
                                                                                                     move
                                                                                                         |dictBitraversable1|
                                                                                                         {
                                                                                                             let bifunctorProduct21 =
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&bifunctorProduct2,
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictBitraversable1)),
                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                             let bifoldableProduct21 =
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&bifoldableProduct2,
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifoldable1"),
                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictBitraversable1)),
                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_Bitraversableusd_Dict(),
                                                                                                                                              &&&add(string("bitraverse"),
                                                                                                                                                     &&Func1::new({
                                                                                                                                                                      let dictBitraversable1
                                                                                                                                                                          =
                                                                                                                                                                          dictBitraversable1.clone();
                                                                                                                                                                      move
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
                                                                                                                                                                                                  |l|
                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                  let l
                                                                                                                                                                                                                      =
                                                                                                                                                                                                                      l.clone();
                                                                                                                                                                                                                  move
                                                                                                                                                                                                                      |r|
                                                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                                                      let r
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          r.clone();
                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                          |v|
                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                              let matchValue =
                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&l);
                                                                                                                                                                                                                                              let matchValue_1 =
                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&r);
                                                                                                                                                                                                                                              let matchValue_2:
                                                                                                                                                                                                                                                      LrcPtr<Data_Functor_Product2_Product2> =
                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                                                              let r1 =
                                                                                                                                                                                                                                                  matchValue_1;
                                                                                                                                                                                                                                              let l1 =
                                                                                                                                                                                                                                                  matchValue;
                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                     &&&Apply0),
                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                          |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                                                          Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                         let usd__arg1
                                                                                                                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                                                                                                                             usd__arg1.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                                                                                                                             |usd__arg2|
                                                                                                                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Functor_Product2_Product2::Data_Functor_Product2_Product2usd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                                                                                                                                     }))),
                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&dictBitraversable),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&l1),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&r1),
                                                                                                                                                                                                                                                                                                                                                                                        &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                                               Data_Functor_Product2_Product2::Data_Functor_Product2_Product2usd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                           }))),
                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&dictBitraversable1),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                                                                        &&&l1),
                                                                                                                                                                                                                                                                                                                                                     &&&r1),
                                                                                                                                                                                                                                                                                                                  &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                         Data_Functor_Product2_Product2::Data_Functor_Product2_Product2usd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                                                                          }
                                                                                                                                                                                                                                  })
                                                                                                                                                                                                              })
                                                                                                                                                                                          })
                                                                                                                                                                          }
                                                                                                                                                                  }),
                                                                                                                                                     add(string("bisequence"),
                                                                                                                                                         &&Func1::new({
                                                                                                                                                                          let dictBitraversable1
                                                                                                                                                                              =
                                                                                                                                                                              dictBitraversable1.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |dictApplicative_1|
                                                                                                                                                                              {
                                                                                                                                                                                  let Apply0_1 =
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(dictApplicative_1)),
                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                                                  let Functor0_1 =
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictApplicative_1)),
                                                                                                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                  let Apply0_1
                                                                                                                                                                                                      =
                                                                                                                                                                                                      Apply0_1.clone();
                                                                                                                                                                                                  let Functor0_1
                                                                                                                                                                                                      =
                                                                                                                                                                                                      Functor0_1.clone();
                                                                                                                                                                                                  let dictApplicative_1
                                                                                                                                                                                                      =
                                                                                                                                                                                                      dictApplicative_1.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |v_1|
                                                                                                                                                                                                      {
                                                                                                                                                                                                          let matchValue_4:
                                                                                                                                                                                                                  LrcPtr<Data_Functor_Product2_Product2> =
                                                                                                                                                                                                              Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                 &&&Apply0_1),
                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                       &&&Functor0_1),
                                                                                                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                      |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                      Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                     let usd__arg1_1
                                                                                                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                                                                                                         usd__arg1_1.clone();
                                                                                                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                                                                                                         |usd__arg2_1|
                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Functor_Product2_Product2::Data_Functor_Product2_Product2usd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                             usd__arg2_1.clone()))
                                                                                                                                                                                                                                                                                                                                                                                 }))),
                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bisequence(),
                                                                                                                                                                                                                                                                                                                                                                                                                          &&&dictBitraversable),
                                                                                                                                                                                                                                                                                                                                                                                       &&&dictApplicative_1),
                                                                                                                                                                                                                                                                                                                                                    &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                                           Data_Functor_Product2_Product2::Data_Functor_Product2_Product2usd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                                       }))),
                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bisequence(),
                                                                                                                                                                                                                                                                                                                                                    &&&dictBitraversable1),
                                                                                                                                                                                                                                                                                                                 &&&dictApplicative_1),
                                                                                                                                                                                                                                                                              &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                     Data_Functor_Product2_Product2::Data_Functor_Product2_Product2usd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                 }))
                                                                                                                                                                                                      }
                                                                                                                                                                                              })
                                                                                                                                                                              }
                                                                                                                                                                      }),
                                                                                                                                                         add(string("Bifunctor0"),
                                                                                                                                                             &&Func1::new({
                                                                                                                                                                              let bifunctorProduct21
                                                                                                                                                                                  =
                                                                                                                                                                                  bifunctorProduct21.clone();
                                                                                                                                                                              move
                                                                                                                                                                                  |usd__unused|
                                                                                                                                                                                  &bifunctorProduct21
                                                                                                                                                                          }),
                                                                                                                                                             add(string("Bifoldable1"),
                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                  let bifoldableProduct21
                                                                                                                                                                                      =
                                                                                                                                                                                      bifoldableProduct21.clone();
                                                                                                                                                                                  move
                                                                                                                                                                                      |usd__unused_1|
                                                                                                                                                                                      &bifoldableProduct21
                                                                                                                                                                              }),
                                                                                                                                                                 empty::<string,
                                                                                                                                                                         &dyn Any>())))))
                                                                                                         }
                                                                                                 })
                                                                                 }))
    }
    pub fn Data_Bitraversable_bitraverseDefault() -> &dyn Any {
        static Data_Bitraversable_bitraverseDefault: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Bitraversable_bitraverseDefault.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictBitraversable|
                                                                             {
                                                                                 let Bifunctor0 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                             Sharpurs_Prelude::unbox(dictBitraversable)),
                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                 &Func1::new({
                                                                                                 let Bifunctor0
                                                                                                     =
                                                                                                     Bifunctor0.clone();
                                                                                                 let dictBitraversable
                                                                                                     =
                                                                                                     dictBitraversable.clone();
                                                                                                 move
                                                                                                     |dictApplicative|
                                                                                                     &Func1::new({
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
                                                                                                                                             |g|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let g
                                                                                                                                                                 =
                                                                                                                                                                 g.clone();
                                                                                                                                                             move
                                                                                                                                                                 |t|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bisequence(),
                                                                                                                                                                                                                                                                        &&&dictBitraversable),
                                                                                                                                                                                                                                     &&&dictApplicative),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                              &&&Bifunctor0),
                                                                                                                                                                                                                                                                                                           &&&f),
                                                                                                                                                                                                                                                                        &&&g),
                                                                                                                                                                                                                                     t))
                                                                                                                                                         })
                                                                                                                                     })
                                                                                                                 })
                                                                                             })
                                                                             }))
    }
    pub fn Data_Bitraversable_bifor() -> &dyn Any {
        static Data_Bitraversable_bifor: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bitraversable_bifor.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictBitraversable|
                                                                 &Func1::new({
                                                                                 let dictBitraversable
                                                                                     =
                                                                                     dictBitraversable.clone();
                                                                                 move
                                                                                     |dictApplicative|
                                                                                     &Func1::new({
                                                                                                     let dictApplicative
                                                                                                         =
                                                                                                         dictApplicative.clone();
                                                                                                     move
                                                                                                         |t|
                                                                                                         &Func1::new({
                                                                                                                         let t
                                                                                                                             =
                                                                                                                             t.clone();
                                                                                                                         move
                                                                                                                             |f|
                                                                                                                             &Func1::new({
                                                                                                                                             let f
                                                                                                                                                 =
                                                                                                                                                 f.clone();
                                                                                                                                             move
                                                                                                                                                 |g|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                                                                                                                                              &&&dictBitraversable),
                                                                                                                                                                                                                                                                                           &&&dictApplicative),
                                                                                                                                                                                                                                                        &&&f),
                                                                                                                                                                                                                     g),
                                                                                                                                                                                  &&&t)
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
}
