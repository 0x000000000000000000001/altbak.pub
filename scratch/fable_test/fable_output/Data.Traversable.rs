pub mod PureScript_Data_Traversable {
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
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_a8445950::PureScript_Data_Const;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_419ece9e::PureScript_Data_Foldable;
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
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_3b72fe33::PureScript_Data_Monoid_Additive;
    use crate::module_89cc47bd::PureScript_Data_Monoid_Conj;
    use crate::module_6a9d0e21::PureScript_Data_Monoid_Disj;
    use crate::module_5023b7e9::PureScript_Data_Monoid_Dual;
    use crate::module_7d83a5e5::PureScript_Data_Monoid_Multiplicative;
    use crate::module_833c8c54::PureScript_Data_Traversable_Accum_Internal;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub mod Data_Traversable_FFI {
        use super::*;
        use fable_library_rust::Array_::append;
        use fable_library_rust::Native_::Func2;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::Native_::fix2;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_array;
        use fable_library_rust::NativeArray_::new_init;
        pub fn traverseArrayImpl(apply: &dyn Any, mapFn: &dyn Any,
                                 pureFn: &dyn Any, f: &dyn Any,
                                 arrayVal: &dyn Any) -> &dyn Any {
            let arr = arrayVal.clone();
            let go =
                Func2::new({
                               let apply = apply.clone();
                               let arr = arr.clone();
                               let f = f.clone();
                               let mapFn = mapFn.clone();
                               let pureFn = pureFn.clone();
                               move |bot: i32, top: i32|
                                   fix2(&(move |go, bot: i32, top: i32|
                                              {
                                                  let diff: i32 = top - bot;
                                                  if diff == 0_i32 {
                                                      Sharpurs_Prelude::sharpurs_apply(&pureFn,
                                                                                       &&new_init(&defaultOf(),
                                                                                                  0_i32))
                                                  } else {
                                                      if diff == 1_i32 {
                                                          Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&mapFn,
                                                                                                                             &&Func1::new(move
                                                                                                                                              |a|
                                                                                                                                              &new_array(&[a.clone()]))),
                                                                                           &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                             &arr[bot].clone()))
                                                      } else {
                                                          if diff == 2_i32 {
                                                              Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&apply,
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&mapFn,
                                                                                                                                                                                                     &&Func1::new(move
                                                                                                                                                                                                                      |a_1|
                                                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                                                      let a_1
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          a_1.clone();
                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                          |b|
                                                                                                                                                                                                                                          &new_array(&[a_1,
                                                                                                                                                                                                                                                       b.clone()])
                                                                                                                                                                                                                                  }))),
                                                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                                                                                     &arr[bot].clone()))),
                                                                                               &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                 &arr[bot
                                                                                                                                          +
                                                                                                                                          1_i32].clone()))
                                                          } else {
                                                              if diff == 3_i32
                                                                 {
                                                                  Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&apply,
                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&apply,
                                                                                                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&mapFn,
                                                                                                                                                                                                                                                                             &&Func1::new(move
                                                                                                                                                                                                                                                                                              |a_2|
                                                                                                                                                                                                                                                                                              &Func1::new({
                                                                                                                                                                                                                                                                                                              let a_2
                                                                                                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                                                                                                  a_2.clone();
                                                                                                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                                                                                                  |b_1|
                                                                                                                                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                                                                                                                                  let b_1
                                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                                      b_1.clone();
                                                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                                                      |c|
                                                                                                                                                                                                                                                                                                                                      &new_array(&[a_2,
                                                                                                                                                                                                                                                                                                                                                   b_1,
                                                                                                                                                                                                                                                                                                                                                   c.clone()])
                                                                                                                                                                                                                                                                                                                              })
                                                                                                                                                                                                                                                                                                          }))),
                                                                                                                                                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                                                                                                                                                             &arr[bot].clone()))),
                                                                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                                                                                         &arr[bot
                                                                                                                                                                                                                  +
                                                                                                                                                                                                                  1_i32].clone()))),
                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                     &arr[bot
                                                                                                                                              +
                                                                                                                                              2_i32].clone()))
                                                              } else {
                                                                  let pivot:
                                                                          i32 =
                                                                      bot +
                                                                          diff
                                                                              /
                                                                              4_i32
                                                                              *
                                                                              2_i32;
                                                                  Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&apply,
                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&mapFn,
                                                                                                                                                                                                         &&Func1::new(move
                                                                                                                                                                                                                          |xsVal|
                                                                                                                                                                                                                          &Func1::new({
                                                                                                                                                                                                                                          let xsVal
                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                              xsVal.clone();
                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                              |ysVal|
                                                                                                                                                                                                                                              &append(xsVal,
                                                                                                                                                                                                                                                      ysVal.clone())
                                                                                                                                                                                                                                      }))),
                                                                                                                                                                       &go(bot,
                                                                                                                                                                           pivot))),
                                                                                                   &go(pivot,
                                                                                                       top))
                                                              }
                                                          }
                                                      }
                                                  }
                                              }), bot, top)
                           });
            go(0_i32, count(arr.clone()))
        }
    }
    pub fn Data_Traversable_traverseArrayImpl() -> &dyn Any {
        static Data_Traversable_traverseArrayImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traverseArrayImpl.get_or_init(||
                                                           &Func1::new(move
                                                                           |apply|
                                                                           Func1::new({
                                                                                          let apply
                                                                                              =
                                                                                              apply.clone();
                                                                                          move
                                                                                              |mapFn|
                                                                                              Func1::new({
                                                                                                             let mapFn
                                                                                                                 =
                                                                                                                 mapFn.clone();
                                                                                                             move
                                                                                                                 |pureFn|
                                                                                                                 Func1::new({
                                                                                                                                let pureFn
                                                                                                                                    =
                                                                                                                                    pureFn.clone();
                                                                                                                                move
                                                                                                                                    |f|
                                                                                                                                    Func1::new({
                                                                                                                                                   let f
                                                                                                                                                       =
                                                                                                                                                       f.clone();
                                                                                                                                                   move
                                                                                                                                                       |arrayVal|
                                                                                                                                                       PureScript_Data_Traversable::Data_Traversable_FFI::traverseArrayImpl(&apply,
                                                                                                                                                                                                                            &mapFn,
                                                                                                                                                                                                                            &pureFn,
                                                                                                                                                                                                                            &f,
                                                                                                                                                                                                                            arrayVal)
                                                                                                                                               })
                                                                                                                            })
                                                                                                         })
                                                                                      })))
    }
    pub fn Data_Traversable_identity() -> &dyn Any {
        static Data_Traversable_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_identity.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                   &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Traversable_Traversableusd_Dict() -> &dyn Any {
        static Data_Traversable_Traversableusd_Dict: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Traversable_Traversableusd_Dict.get_or_init(||
                                                             &Func1::new(move
                                                                             |x|
                                                                             x.clone()))
    }
    pub fn Data_Traversable_traverse() -> &dyn Any {
        static Data_Traversable_traverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traverse.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("traverse"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Traversable_traversableTuple() -> &dyn Any {
        static Data_Traversable_traversableTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableTuple.get_or_init(||
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
                                                                                                      add(string("Functor0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                                                                          add(string("Foldable1"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused_1|
                                                                                                                               &PureScript_Data_Foldable::Data_Foldable_foldableTuple()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableMultiplicative() -> &dyn Any {
        static Data_Traversable_traversableMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableMultiplicative.get_or_init(||
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
                                                                                                                                                    move
                                                                                                                                                        |v_1|
                                                                                                                                                        {
                                                                                                                                                            let x_1 =
                                                                                                                                                                Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                   &&&Functor0_1),
                                                                                                                                                                                                                                &&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_Multiplicative()),
                                                                                                                                                                                             &&&x_1)
                                                                                                                                                        }
                                                                                                                                                })
                                                                                                                                }),
                                                                                                               add(string("Functor0"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |usd__unused|
                                                                                                                                    &PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_functorMultiplicative()),
                                                                                                                   add(string("Foldable1"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |usd__unused_1|
                                                                                                                                        &PureScript_Data_Foldable::Data_Foldable_foldableMultiplicative()),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableMaybe() -> &dyn Any {
        static Data_Traversable_traversableMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableMaybe.get_or_init(||
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
                                                                                                                                                                           LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                   match matchValue_1.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                       =>
                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                              &&&Functor0),
                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                             |usd__arg1|
                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                           &&matchValue_1_1_0)),
                                                                                                                                                                       _
                                                                                                                                                                       =>
                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                           &&&dictApplicative),
                                                                                                                                                                                                        &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                                   }
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
                                                                                                                                                   let matchValue_3:
                                                                                                                                                           LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                       Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                   match matchValue_3.as_ref()
                                                                                                                                                       {
                                                                                                                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_3_1_0)
                                                                                                                                                       =>
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                              &&&Functor0_1),
                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                             |usd__arg1_1|
                                                                                                                                                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                        &&matchValue_3_1_0),
                                                                                                                                                       _
                                                                                                                                                       =>
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                           &&&dictApplicative_1),
                                                                                                                                                                                        &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                   }
                                                                                                                                               }
                                                                                                                                       })
                                                                                                                       }),
                                                                                                      add(string("Functor0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                          add(string("Foldable1"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused_1|
                                                                                                                               &PureScript_Data_Foldable::Data_Foldable_foldableMaybe()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableIdentity() -> &dyn Any {
        static Data_Traversable_traversableIdentity: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Traversable_traversableIdentity.get_or_init(||
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
                                                                                                         add(string("Functor0"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused|
                                                                                                                              &PureScript_Data_Identity::Data_Identity_functorIdentity()),
                                                                                                             add(string("Foldable1"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |usd__unused_1|
                                                                                                                                  &PureScript_Data_Foldable::Data_Foldable_foldableIdentity()),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableEither() -> &dyn Any {
        static Data_Traversable_traversableEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableEither.get_or_init(||
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
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
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
                                                                                                                                                    let matchValue_3:
                                                                                                                                                            LrcPtr<Data_Either_Either> =
                                                                                                                                                        Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                    match matchValue_3.as_ref()
                                                                                                                                                        {
                                                                                                                                                        Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_3_1_0)
                                                                                                                                                        =>
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                               &&&Functor0_1),
                                                                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                                                                              |usd__arg1_1|
                                                                                                                                                                                                                                              &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                         &&matchValue_3_1_0),
                                                                                                                                                        Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_3_0_0)
                                                                                                                                                        =>
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                            &&&dictApplicative_1),
                                                                                                                                                                                         &&&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_3_0_0))),
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                        })
                                                                                                                        }),
                                                                                                       add(string("Functor0"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused|
                                                                                                                            &PureScript_Data_Either::Data_Either_functorEither()),
                                                                                                           add(string("Foldable1"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused_1|
                                                                                                                                &PureScript_Data_Foldable::Data_Foldable_foldableEither()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableDual() -> &dyn Any {
        static Data_Traversable_traversableDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableDual.get_or_init(||
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
                                                                                                                                          move
                                                                                                                                              |v_1|
                                                                                                                                              {
                                                                                                                                                  let x_1 =
                                                                                                                                                      Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                         &&&Functor0_1),
                                                                                                                                                                                                                      &&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual()),
                                                                                                                                                                                   &&&x_1)
                                                                                                                                              }
                                                                                                                                      })
                                                                                                                      }),
                                                                                                     add(string("Functor0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_Monoid_Dual::Data_Monoid_Dual_functorDual()),
                                                                                                         add(string("Foldable1"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused_1|
                                                                                                                              &PureScript_Data_Foldable::Data_Foldable_foldableDual()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableDisj() -> &dyn Any {
        static Data_Traversable_traversableDisj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableDisj.get_or_init(||
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
                                                                                                                                                                                                                                      &&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_Disj()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                      &&&matchValue_1))
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
                                                                                                                                          move
                                                                                                                                              |v_1|
                                                                                                                                              {
                                                                                                                                                  let x_1 =
                                                                                                                                                      Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                         &&&Functor0_1),
                                                                                                                                                                                                                      &&&PureScript_Data_Monoid_Disj::Data_Monoid_Disj_Disj()),
                                                                                                                                                                                   &&&x_1)
                                                                                                                                              }
                                                                                                                                      })
                                                                                                                      }),
                                                                                                     add(string("Functor0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_Monoid_Disj::Data_Monoid_Disj_functorDisj()),
                                                                                                         add(string("Foldable1"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused_1|
                                                                                                                              &PureScript_Data_Foldable::Data_Foldable_foldableDisj()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableConst() -> &dyn Any {
        static Data_Traversable_traversableConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableConst.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                                                                           &&&add(string("traverse"),
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
                                                                                                  add(string("sequence"),
                                                                                                      &&Func1::new(move
                                                                                                                       |dictApplicative_1|
                                                                                                                       &Func1::new({
                                                                                                                                       let dictApplicative_1
                                                                                                                                           =
                                                                                                                                           dictApplicative_1.clone();
                                                                                                                                       move
                                                                                                                                           |v_1|
                                                                                                                                           {
                                                                                                                                               let x_1 =
                                                                                                                                                   Sharpurs_Prelude::unbox(v_1);
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                   &&&dictApplicative_1),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_Const(),
                                                                                                                                                                                                                   &&&x_1))
                                                                                                                                           }
                                                                                                                                   })),
                                                                                                      add(string("Functor0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Const::Data_Const_functorConst()),
                                                                                                          add(string("Foldable1"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused_1|
                                                                                                                               &PureScript_Data_Foldable::Data_Foldable_foldableConst()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableConj() -> &dyn Any {
        static Data_Traversable_traversableConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableConj.get_or_init(||
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
                                                                                                                                                                                                                                      &&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                      &&&matchValue_1))
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
                                                                                                                                          move
                                                                                                                                              |v_1|
                                                                                                                                              {
                                                                                                                                                  let x_1 =
                                                                                                                                                      Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                         &&&Functor0_1),
                                                                                                                                                                                                                      &&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj()),
                                                                                                                                                                                   &&&x_1)
                                                                                                                                              }
                                                                                                                                      })
                                                                                                                      }),
                                                                                                     add(string("Functor0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_Monoid_Conj::Data_Monoid_Conj_functorConj()),
                                                                                                         add(string("Foldable1"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused_1|
                                                                                                                              &PureScript_Data_Foldable::Data_Foldable_foldableConj()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableCompose_004078() -> &dyn Any {
        &Func1::new(move |dictTraversable|
                        PureScript_Data_Traversable::Data_Traversable_traversableCompose_tco(dictTraversable))
    }
    pub fn Data_Traversable_traversableCompose_004078_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Traversable_traversableCompose_004078_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Traversable_traversableCompose_004078_002d1.get_or_init(||
                                                                         Lazy(Data_Traversable_traversableCompose_004078.clone()))
    }
    pub fn Data_Traversable_traversableCompose_tco(dictTraversable: &dyn Any)
     -> &dyn Any {
        let functorCompose =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_functorCompose(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                       Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        let foldableCompose =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldableCompose(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                       Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        &Func1::new({
                        let dictTraversable = dictTraversable.clone();
                        let foldableCompose = foldableCompose.clone();
                        let functorCompose = functorCompose.clone();
                        move |dictTraversable1|
                            {
                                let functorCompose1 =
                                    Sharpurs_Prelude::sharpurs_apply(&&&functorCompose,
                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                               Sharpurs_Prelude::unbox(dictTraversable1)),
                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                let foldableCompose1 =
                                    Sharpurs_Prelude::sharpurs_apply(&&&foldableCompose,
                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                                               Sharpurs_Prelude::unbox(dictTraversable1)),
                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                                                 &&&add(string("traverse"),
                                                                        &&Func1::new({
                                                                                         let dictTraversable1
                                                                                             =
                                                                                             dictTraversable1.clone();
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
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                       &&&Functor0),
                                                                                                                                                                                                                                                    &&&PureScript_Data_Functor_Compose::Data_Functor_Compose_Compose())),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                          &&&dictTraversable),
                                                                                                                                                                                                                                                                                       &&&dictApplicative),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                             &&&dictTraversable1),
                                                                                                                                                                                                                                                                                                                          &&&dictApplicative),
                                                                                                                                                                                                                                                                                       &&&matchValue)),
                                                                                                                                                                                                                 &&&matchValue_1))
                                                                                                                                         }
                                                                                                                                 })
                                                                                                             })
                                                                                             }
                                                                                     }),
                                                                        add(string("sequence"),
                                                                            &&Func1::new({
                                                                                             let dictTraversable1
                                                                                                 =
                                                                                                 dictTraversable1.clone();
                                                                                             move
                                                                                                 |dictApplicative_1|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&PureScript_Data_Traversable::Data_Traversable_traversableCompose_tco(&&dictTraversable),
                                                                                                                                                                                                                                           &&&dictTraversable1)),
                                                                                                                                                                     dictApplicative_1),
                                                                                                                                  &&&PureScript_Data_Traversable::Data_Traversable_identity())
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
                                                                                add(string("Foldable1"),
                                                                                    &&Func1::new({
                                                                                                     let foldableCompose1
                                                                                                         =
                                                                                                         foldableCompose1.clone();
                                                                                                     move
                                                                                                         |usd__unused_1|
                                                                                                         &foldableCompose1
                                                                                                 }),
                                                                                    empty::<string,
                                                                                            &dyn Any>())))))
                            }
                    })
    }
    pub fn Data_Traversable_traversableCompose() -> &dyn Any {
        static Data_Traversable_traversableCompose: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Traversable_traversableCompose.get_or_init(||
                                                            Data_Traversable_traversableCompose_004078_002d1.Value)
    }
    pub fn Data_Traversable_traversableAdditive() -> &dyn Any {
        static Data_Traversable_traversableAdditive: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Traversable_traversableAdditive.get_or_init(||
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
                                                                                                                                                                                                                                          &&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive()),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                          &&&matchValue_1))
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
                                                                                                                                              move
                                                                                                                                                  |v_1|
                                                                                                                                                  {
                                                                                                                                                      let x_1 =
                                                                                                                                                          Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                             &&&Functor0_1),
                                                                                                                                                                                                                          &&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive()),
                                                                                                                                                                                       &&&x_1)
                                                                                                                                                  }
                                                                                                                                          })
                                                                                                                          }),
                                                                                                         add(string("Functor0"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused|
                                                                                                                              &PureScript_Data_Monoid_Additive::Data_Monoid_Additive_functorAdditive()),
                                                                                                             add(string("Foldable1"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |usd__unused_1|
                                                                                                                                  &PureScript_Data_Foldable::Data_Foldable_foldableAdditive()),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>()))))))
    }
    pub fn Data_Traversable_sequenceDefault() -> &dyn Any {
        static Data_Traversable_sequenceDefault: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_sequenceDefault.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictTraversable|
                                                                         &Func1::new({
                                                                                         let dictTraversable
                                                                                             =
                                                                                             dictTraversable.clone();
                                                                                         move
                                                                                             |dictApplicative|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                    &&&dictTraversable),
                                                                                                                                                                 dictApplicative),
                                                                                                                              &&&PureScript_Data_Traversable::Data_Traversable_identity())
                                                                                     })))
    }
    pub fn Data_Traversable_traversableArray_004085() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                         &&&add(string("traverse"),
                                                &&Func1::new(move
                                                                 |dictApplicative|
                                                                 {
                                                                     let Apply0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                 Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverseArrayImpl(),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                               &&&Apply0)),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&Apply0)),
                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                         dictApplicative))
                                                                 }),
                                                add(string("sequence"),
                                                    &&Func1::new({
                                                                     let Data_Traversable_traversableArray_004085_002d1
                                                                         =
                                                                         Data_Traversable_traversableArray_004085_002d1.clone();
                                                                     move
                                                                         |dictApplicative_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequenceDefault(),
                                                                                                                                             &&&Data_Traversable_traversableArray_004085_002d1.Value),
                                                                                                          dictApplicative_1)
                                                                 }),
                                                    add(string("Functor0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Functor::Data_Functor_functorArray()),
                                                        add(string("Foldable1"),
                                                            &&Func1::new(move
                                                                             |usd__unused_1|
                                                                             &PureScript_Data_Foldable::Data_Foldable_foldableArray()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_Traversable_traversableArray_004085_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Traversable_traversableArray_004085_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Traversable_traversableArray_004085_002d1.get_or_init(||
                                                                       Lazy(Data_Traversable_traversableArray_004085.clone()))
    }
    pub fn Data_Traversable_traversableArray() -> &dyn Any {
        static Data_Traversable_traversableArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableArray.get_or_init(||
                                                          Data_Traversable_traversableArray_004085_002d1.Value)
    }
    pub fn Data_Traversable_sequence() -> &dyn Any {
        static Data_Traversable_sequence: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_sequence.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("sequence"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Traversable_traversableApp() -> &dyn Any {
        static Data_Traversable_traversableApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableApp.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictTraversable|
                                                                        {
                                                                            let functorApp =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_functorApp(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            let foldableApp =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldableApp(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                                                                                             &&&add(string("traverse"),
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
                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                                      &&&dictTraversable),
                                                                                                                                                                                                                                                                                                                                   &&&dictApplicative),
                                                                                                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                                                                                                             &&&matchValue_1))
                                                                                                                                                                                     }
                                                                                                                                                                             })
                                                                                                                                                         })
                                                                                                                                         }
                                                                                                                                 }),
                                                                                                                    add(string("sequence"),
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
                                                                                                                                                                         let x_1 =
                                                                                                                                                                             Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                &&&Functor0_1),
                                                                                                                                                                                                                                             &&&PureScript_Data_Functor_App::Data_Functor_App_App()),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                                                                                   &&&dictTraversable),
                                                                                                                                                                                                                                                                                &&&dictApplicative_1),
                                                                                                                                                                                                                                             &&&x_1))
                                                                                                                                                                     }
                                                                                                                                                             })
                                                                                                                                             }
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
                                                                                                                            add(string("Foldable1"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let foldableApp
                                                                                                                                                     =
                                                                                                                                                     foldableApp.clone();
                                                                                                                                                 move
                                                                                                                                                     |usd__unused_1|
                                                                                                                                                     &foldableApp
                                                                                                                                             }),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>())))))
                                                                        }))
    }
    pub fn Data_Traversable_traversableCoproduct() -> &dyn Any {
        static Data_Traversable_traversableCoproduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableCoproduct.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictTraversable|
                                                                              {
                                                                                  let sequence1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                       dictTraversable);
                                                                                  let functorCoproduct =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_functorCoproduct(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  let foldableCoproduct =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldableCoproduct(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  &Func1::new({
                                                                                                  let dictTraversable
                                                                                                      =
                                                                                                      dictTraversable.clone();
                                                                                                  let foldableCoproduct
                                                                                                      =
                                                                                                      foldableCoproduct.clone();
                                                                                                  let functorCoproduct
                                                                                                      =
                                                                                                      functorCoproduct.clone();
                                                                                                  let sequence1
                                                                                                      =
                                                                                                      sequence1.clone();
                                                                                                  move
                                                                                                      |dictTraversable1|
                                                                                                      {
                                                                                                          let sequence2 =
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                               dictTraversable1);
                                                                                                          let functorCoproduct1 =
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&functorCoproduct,
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictTraversable1)),
                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                          let foldableCoproduct1 =
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&foldableCoproduct,
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictTraversable1)),
                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                                                                                                                           &&&add(string("traverse"),
                                                                                                                                                  &&Func1::new({
                                                                                                                                                                   let dictTraversable1
                                                                                                                                                                       =
                                                                                                                                                                       dictTraversable1.clone();
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
                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                                                                               &&&dictTraversable),
                                                                                                                                                                                                                                                                                                                                                                            &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                         f))),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                            &&&Functor0),
                                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                                                                                                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                              |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))))),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                                            &&&dictTraversable1),
                                                                                                                                                                                                                                                                                                                                         &&&dictApplicative),
                                                                                                                                                                                                                                                                                                      f)))
                                                                                                                                                                                       })
                                                                                                                                                                       }
                                                                                                                                                               }),
                                                                                                                                                  add(string("sequence"),
                                                                                                                                                      &&Func1::new({
                                                                                                                                                                       let sequence2
                                                                                                                                                                           =
                                                                                                                                                                           sequence2.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |dictApplicative_1|
                                                                                                                                                                           {
                                                                                                                                                                               let Functor0_1 =
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictApplicative_1)),
                                                                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                               &&&Functor0_1),
                                                                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                                                                                                                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                 |usd__arg1_2|
                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_2.clone())))))),
                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&sequence1,
                                                                                                                                                                                                                                                                                                                         dictApplicative_1))),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                            &&&Functor0_1),
                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                                                                                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                              |usd__arg1_3|
                                                                                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_3.clone())))))),
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&sequence2,
                                                                                                                                                                                                                                                                                      dictApplicative_1)))
                                                                                                                                                                           }
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
                                                                                                                                                          add(string("Foldable1"),
                                                                                                                                                              &&Func1::new({
                                                                                                                                                                               let foldableCoproduct1
                                                                                                                                                                                   =
                                                                                                                                                                                   foldableCoproduct1.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |usd__unused_1|
                                                                                                                                                                                   &foldableCoproduct1
                                                                                                                                                                           }),
                                                                                                                                                              empty::<string,
                                                                                                                                                                      &dyn Any>())))))
                                                                                                      }
                                                                                              })
                                                                              }))
    }
    pub fn Data_Traversable_traversableFirst() -> &dyn Any {
        static Data_Traversable_traversableFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableFirst.get_or_init(||
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
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                                                                                                       &&&PureScript_Data_Maybe_First::Data_Maybe_First_First()),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Traversable::Data_Traversable_traversableMaybe()),
                                                                                                                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                                                       &&&matchValue_1))
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
                                                                                                                                                   let x_1 =
                                                                                                                                                       Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                          &&&Functor0_1),
                                                                                                                                                                                                                       &&&PureScript_Data_Maybe_First::Data_Maybe_First_First()),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Traversable::Data_Traversable_traversableMaybe()),
                                                                                                                                                                                                                                                          &&&dictApplicative_1),
                                                                                                                                                                                                                       &&&x_1))
                                                                                                                                               }
                                                                                                                                       })
                                                                                                                       }),
                                                                                                      add(string("Functor0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Maybe_First::Data_Maybe_First_functorFirst()),
                                                                                                          add(string("Foldable1"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused_1|
                                                                                                                               &PureScript_Data_Foldable::Data_Foldable_foldableFirst()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableLast() -> &dyn Any {
        static Data_Traversable_traversableLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traversableLast.get_or_init(||
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
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                         &&&Functor0),
                                                                                                                                                                                                                                      &&&PureScript_Data_Maybe_Last::Data_Maybe_Last_Last()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Traversable::Data_Traversable_traversableMaybe()),
                                                                                                                                                                                                                                                                                                            &&&dictApplicative),
                                                                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                                                                      &&&matchValue_1))
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
                                                                                                                                                  let x_1 =
                                                                                                                                                      Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                         &&&Functor0_1),
                                                                                                                                                                                                                      &&&PureScript_Data_Maybe_Last::Data_Maybe_Last_Last()),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Traversable::Data_Traversable_traversableMaybe()),
                                                                                                                                                                                                                                                         &&&dictApplicative_1),
                                                                                                                                                                                                                      &&&x_1))
                                                                                                                                              }
                                                                                                                                      })
                                                                                                                      }),
                                                                                                     add(string("Functor0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_Maybe_Last::Data_Maybe_Last_functorLast()),
                                                                                                         add(string("Foldable1"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused_1|
                                                                                                                              &PureScript_Data_Foldable::Data_Foldable_foldableLast()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))))
    }
    pub fn Data_Traversable_traversableProduct() -> &dyn Any {
        static Data_Traversable_traversableProduct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Traversable_traversableProduct.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictTraversable|
                                                                            {
                                                                                let functorProduct =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Product::Data_Functor_Product_functorProduct(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                               Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                let foldableProduct =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldableProduct(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                                                                                               Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                &Func1::new({
                                                                                                let dictTraversable
                                                                                                    =
                                                                                                    dictTraversable.clone();
                                                                                                let foldableProduct
                                                                                                    =
                                                                                                    foldableProduct.clone();
                                                                                                let functorProduct
                                                                                                    =
                                                                                                    functorProduct.clone();
                                                                                                move
                                                                                                    |dictTraversable1|
                                                                                                    {
                                                                                                        let functorProduct1 =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&functorProduct,
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                       Sharpurs_Prelude::unbox(dictTraversable1)),
                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                        let foldableProduct1 =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&foldableProduct,
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                                                                                                                       Sharpurs_Prelude::unbox(dictTraversable1)),
                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                                                                                                                         &&&add(string("traverse"),
                                                                                                                                                &&Func1::new({
                                                                                                                                                                 let dictTraversable1
                                                                                                                                                                     =
                                                                                                                                                                     dictTraversable1.clone();
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
                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&dictTraversable),
                                                                                                                                                                                                                                                                                                                                                                                                  &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                                                               &&&f1),
                                                                                                                                                                                                                                                                                                                            &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                               })),
                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                                                                  &&&dictTraversable1),
                                                                                                                                                                                                                                                                                                                                                               &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                            &&&f1),
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
                                                                                                                                                add(string("sequence"),
                                                                                                                                                    &&Func1::new({
                                                                                                                                                                     let dictTraversable1
                                                                                                                                                                         =
                                                                                                                                                                         dictTraversable1.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |dictApplicative_1|
                                                                                                                                                                         {
                                                                                                                                                                             let Apply0_1 =
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictApplicative_1)),
                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                             let Apply0_1
                                                                                                                                                                                                 =
                                                                                                                                                                                                 Apply0_1.clone();
                                                                                                                                                                                             let dictApplicative_1
                                                                                                                                                                                                 =
                                                                                                                                                                                                 dictApplicative_1.clone();
                                                                                                                                                                                             move
                                                                                                                                                                                                 |v_1|
                                                                                                                                                                                                 {
                                                                                                                                                                                                     let matchValue_3:
                                                                                                                                                                                                             LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                                                                                                                                                                               &&&Apply0_1),
                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Functor_Product::Data_Functor_Product_product()),
                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                                                                                                                                                  &&&dictTraversable),
                                                                                                                                                                                                                                                                                                                                               &&&dictApplicative_1),
                                                                                                                                                                                                                                                                                                            &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                               })),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                                                                                                               &&&dictTraversable1),
                                                                                                                                                                                                                                                                                                            &&&dictApplicative_1),
                                                                                                                                                                                                                                                                         &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                            }))
                                                                                                                                                                                                 }
                                                                                                                                                                                         })
                                                                                                                                                                         }
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
                                                                                                                                                        add(string("Foldable1"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let foldableProduct1
                                                                                                                                                                                 =
                                                                                                                                                                                 foldableProduct1.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |usd__unused_1|
                                                                                                                                                                                 &foldableProduct1
                                                                                                                                                                         }),
                                                                                                                                                            empty::<string,
                                                                                                                                                                    &dyn Any>())))))
                                                                                                    }
                                                                                            })
                                                                            }))
    }
    pub fn Data_Traversable_traverseDefault() -> &dyn Any {
        static Data_Traversable_traverseDefault: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_traverseDefault.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictTraversable|
                                                                         {
                                                                             let Functor0 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                         Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined());
                                                                             &Func1::new({
                                                                                             let Functor0
                                                                                                 =
                                                                                                 Functor0.clone();
                                                                                             let dictTraversable
                                                                                                 =
                                                                                                 dictTraversable.clone();
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
                                                                                                                                         |ta|
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                &&&dictTraversable),
                                                                                                                                                                                                             &&&dictApplicative),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                   &&&Functor0),
                                                                                                                                                                                                                                                &&&f),
                                                                                                                                                                                                             ta))
                                                                                                                                 })
                                                                                                             })
                                                                                         })
                                                                         }))
    }
    pub fn Data_Traversable_mapAccumR() -> &dyn Any {
        static Data_Traversable_mapAccumR: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_mapAccumR.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictTraversable|
                                                                   &Func1::new({
                                                                                   let dictTraversable
                                                                                       =
                                                                                       dictTraversable.clone();
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
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                               &&&dictTraversable),
                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_applicativeStateR()),
                                                                                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                                                                                           |a|
                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateR(),
                                                                                                                                                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                              let a
                                                                                                                                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                                                                                                                                  a.clone();
                                                                                                                                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                                                                                                                                  |s|
                                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                      s),
                                                                                                                                                                                                                                                                                                                                                                                   &&&a)
                                                                                                                                                                                                                                                                                                                                          })))),
                                                                                                                                                                                                                                      xs)),
                                                                                                                                                                &&&s0)
                                                                                                                       })
                                                                                                   })
                                                                               })))
    }
    pub fn Data_Traversable_scanr() -> &dyn Any {
        static Data_Traversable_scanr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_scanr.get_or_init(||
                                               &Func1::new(move
                                                               |dictTraversable|
                                                               &Func1::new({
                                                                               let dictTraversable
                                                                                   =
                                                                                   dictTraversable.clone();
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
                                                                                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_mapAccumR(),
                                                                                                                                                                                                                                                                                                   &&&dictTraversable),
                                                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                                                  |b|
                                                                                                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                                                                                                  let b
                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                      b.clone();
                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                      |a|
                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                          let b_prime =
                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                  a),
                                                                                                                                                                                                                                                                                                                                               &&&b);
                                                                                                                                                                                                                                                                                                          &add(string("accum"),
                                                                                                                                                                                                                                                                                                               &&b_prime,
                                                                                                                                                                                                                                                                                                               add(string("value"),
                                                                                                                                                                                                                                                                                                                   &&b_prime,
                                                                                                                                                                                                                                                                                                                   empty::<string,
                                                                                                                                                                                                                                                                                                                           &dyn Any>()))
                                                                                                                                                                                                                                                                                                      }
                                                                                                                                                                                                                                                                                              }))),
                                                                                                                                                                                                                             &&&b0),
                                                                                                                                                                                          xs)))
                                                                                                                   })
                                                                                               })
                                                                           })))
    }
    pub fn Data_Traversable_mapAccumL() -> &dyn Any {
        static Data_Traversable_mapAccumL: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_mapAccumL.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictTraversable|
                                                                   &Func1::new({
                                                                                   let dictTraversable
                                                                                       =
                                                                                       dictTraversable.clone();
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
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                               &&&dictTraversable),
                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_applicativeStateL()),
                                                                                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                                                                                           |a|
                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateL(),
                                                                                                                                                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                              let a
                                                                                                                                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                                                                                                                                  a.clone();
                                                                                                                                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                                                                                                                                  |s|
                                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                      s),
                                                                                                                                                                                                                                                                                                                                                                                   &&&a)
                                                                                                                                                                                                                                                                                                                                          })))),
                                                                                                                                                                                                                                      xs)),
                                                                                                                                                                &&&s0)
                                                                                                                       })
                                                                                                   })
                                                                               })))
    }
    pub fn Data_Traversable_scanl() -> &dyn Any {
        static Data_Traversable_scanl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_scanl.get_or_init(||
                                               &Func1::new(move
                                                               |dictTraversable|
                                                               &Func1::new({
                                                                               let dictTraversable
                                                                                   =
                                                                                   dictTraversable.clone();
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
                                                                                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_mapAccumL(),
                                                                                                                                                                                                                                                                                                   &&&dictTraversable),
                                                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                                                  |b|
                                                                                                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                                                                                                  let b
                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                      b.clone();
                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                      |a|
                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                          let b_prime =
                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                  &&&b),
                                                                                                                                                                                                                                                                                                                                               a);
                                                                                                                                                                                                                                                                                                          &add(string("accum"),
                                                                                                                                                                                                                                                                                                               &&b_prime,
                                                                                                                                                                                                                                                                                                               add(string("value"),
                                                                                                                                                                                                                                                                                                                   &&b_prime,
                                                                                                                                                                                                                                                                                                                   empty::<string,
                                                                                                                                                                                                                                                                                                                           &dyn Any>()))
                                                                                                                                                                                                                                                                                                      }
                                                                                                                                                                                                                                                                                              }))),
                                                                                                                                                                                                                             &&&b0),
                                                                                                                                                                                          xs)))
                                                                                                                   })
                                                                                               })
                                                                           })))
    }
    pub fn Data_Traversable_for() -> &dyn Any {
        static Data_Traversable_for: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_for.get_or_init(||
                                             &Func1::new(move
                                                             |dictApplicative|
                                                             &Func1::new({
                                                                             let dictApplicative
                                                                                 =
                                                                                 dictApplicative.clone();
                                                                             move
                                                                                 |dictTraversable|
                                                                                 &Func1::new({
                                                                                                 let dictTraversable
                                                                                                     =
                                                                                                     dictTraversable.clone();
                                                                                                 move
                                                                                                     |x|
                                                                                                     &Func1::new({
                                                                                                                     let x
                                                                                                                         =
                                                                                                                         x.clone();
                                                                                                                     move
                                                                                                                         |f|
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                   &&&dictTraversable),
                                                                                                                                                                                                                                &&&dictApplicative),
                                                                                                                                                                                             f),
                                                                                                                                                          &&&x)
                                                                                                                 })
                                                                                             })
                                                                         })))
    }
}
