pub mod PureScript_Data_Function {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::fix1;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_Function_on() -> &dyn Any {
        static Data_Function_on: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_on.get_or_init(||
                                         &Func1::new(move |f|
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
                                                                                                 |x|
                                                                                                 &Func1::new({
                                                                                                                 let x
                                                                                                                     =
                                                                                                                     x.clone();
                                                                                                                 move
                                                                                                                     |y|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                            &&&x)),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                         y))
                                                                                                             })
                                                                                         })
                                                                     })))
    }
    pub fn Data_Function_flip() -> &dyn Any {
        static Data_Function_flip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_flip.get_or_init(||
                                           &Func1::new(move |f|
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
                                                                                                   |a|
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                       a),
                                                                                                                                    &&&b)
                                                                                           })
                                                                       })))
    }
    pub fn Data_Function_const() -> &dyn Any {
        static Data_Function_const: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_const.get_or_init(||
                                            &Func1::new(move |a|
                                                            &Func1::new({
                                                                            let a
                                                                                =
                                                                                a.clone();
                                                                            move
                                                                                |v|
                                                                                &a
                                                                        })))
    }
    pub fn Data_Function_applyN() -> &dyn Any {
        static Data_Function_applyN: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_applyN.get_or_init(||
                                             &Func1::new(move |f|
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
                                                                                                            |n|
                                                                                                            Func1::new({
                                                                                                                           let go_tco
                                                                                                                               =
                                                                                                                               go_tco.clone();
                                                                                                                           let n
                                                                                                                               =
                                                                                                                               n.clone();
                                                                                                                           move
                                                                                                                               |acc|
                                                                                                                               go_tco(n)(acc.clone())
                                                                                                                       })
                                                                                                    })
                                                                                });
                                                                 let go_1 =
                                                                     Lazy(go_2);
                                                                 let go_tco =
                                                                     Func1::new({
                                                                                    let f
                                                                                        =
                                                                                        f.clone();
                                                                                    move
                                                                                        |n_1|
                                                                                        fix1(&(move
                                                                                                   |go_tco,
                                                                                                    n_1|
                                                                                                   Func1::new({
                                                                                                                  let go_tco
                                                                                                                      =
                                                                                                                      go_tco.clone();
                                                                                                                  let n_1
                                                                                                                      =
                                                                                                                      n_1.clone();
                                                                                                                  move
                                                                                                                      |acc_1|
                                                                                                                      {
                                                                                                                          let matchValue =
                                                                                                                              Sharpurs_Prelude::unbox(&&n_1);
                                                                                                                          let matchValue_1 =
                                                                                                                              Sharpurs_Prelude::unbox(acc_1);
                                                                                                                          if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                             &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                       &&&0_i32))
                                                                                                                             {
                                                                                                                              &matchValue_1
                                                                                                                          } else {
                                                                                                                              if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                 {
                                                                                                                                  go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                                          &&&1_i32))(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                      &&&matchValue_1))
                                                                                                                              } else {
                                                                                                                                  panic!("{}",
                                                                                                                                         LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Function.fs"),
                                  Data1: 15_i32,
                                  Data2: 251_i32,}).get_Message(),)
                                                                                                                              }
                                                                                                                          }
                                                                                                                      }
                                                                                                              })),
                                                                                             n_1.clone())
                                                                                });
                                                                 let go =
                                                                     go_1.Value;
                                                                 &go
                                                             }))
    }
    pub fn Data_Function_applyFlipped() -> &dyn Any {
        static Data_Function_applyFlipped: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_applyFlipped.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   &Func1::new({
                                                                                   let x
                                                                                       =
                                                                                       x.clone();
                                                                                   move
                                                                                       |f|
                                                                                       Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                        &&&x)
                                                                               })))
    }
    pub fn Data_Function_apply() -> &dyn Any {
        static Data_Function_apply: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_apply.get_or_init(||
                                            &Func1::new(move |f|
                                                            &Func1::new({
                                                                            let f
                                                                                =
                                                                                f.clone();
                                                                            move
                                                                                |x|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                 x)
                                                                        })))
    }
}
