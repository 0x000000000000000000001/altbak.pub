pub mod PureScript_Data_Void {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Util_::Lazy;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Void_Void() -> &dyn Any {
        static Data_Void_Void: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Void_Void.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Void_absurd() -> &dyn Any {
        static Data_Void_absurd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Void_absurd.get_or_init(||
                                         &Func1::new(move |a|
                                                         {
                                                             let spin_2 =
                                                                 Func0::new({
                                                                                let spin_tco
                                                                                    =
                                                                                    spin_tco.clone();
                                                                                move
                                                                                    ||
                                                                                    &Func1::new({
                                                                                                    let spin_tco
                                                                                                        =
                                                                                                        spin_tco.clone();
                                                                                                    move
                                                                                                        |v|
                                                                                                        spin_tco(v.clone())
                                                                                                })
                                                                            });
                                                             let spin_1 =
                                                                 Lazy(spin_2);
                                                             fn spin_tco(v_1:
                                                                             _)
                                                              -> &dyn Any {
                                                                 let v_1 =
                                                                     v_1.clone();
                                                                 '_spin_tco:
                                                                     loop  {
                                                                         break
                                                                             '_spin_tco
                                                                             ({
                                                                                  let v_1_temp =
                                                                                      &Sharpurs_Prelude::unbox(&&v_1);
                                                                                  v_1.set(v_1_temp);
                                                                                  continue
                                                                                      '_spin_tco

                                                                              })
                                                                             ;
                                                                     }
                                                             }
                                                             let spin =
                                                                 spin_1.Value;
                                                             spin_tco(a.clone())
                                                         }))
    }
}
