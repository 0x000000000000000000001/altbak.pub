pub mod PureScript_Data_Functor {
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
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub mod Data_Functor_FFI {
        use super::*;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_init;
        pub fn arrayMap() -> &dyn Any {
            static arrayMap: MutCell<Option<&dyn Any>> = MutCell::new(None);
            arrayMap.get_or_init(||
                                     &Func1::new(move |f|
                                                     &Func1::new({
                                                                     let f =
                                                                         f.clone();
                                                                     move
                                                                         |arr|
                                                                         {
                                                                             let a =
                                                                                 Sharpurs_Prelude::unbox(arr);
                                                                             let l:
                                                                                     i32 =
                                                                                 count(a.clone());
                                                                             let result =
                                                                                 new_init(&defaultOf(),
                                                                                          l);
                                                                             for i
                                                                                 in
                                                                                 0_i32..=l
                                                                                             -
                                                                                             1_i32
                                                                                 {
                                                                                 result.get_mut()[i
                                                                                                      as
                                                                                                      usize]
                                                                                     =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                      &a[i].clone());
                                                                             }
                                                                             &result
                                                                         }
                                                                 })))
        }
    }
    pub fn Data_Functor_arrayMap() -> &dyn Any {
        static Data_Functor_arrayMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_arrayMap.get_or_init(||
                                              &PureScript_Data_Functor::Data_Functor_FFI::arrayMap())
    }
    pub fn Data_Functor_Functorusd_Dict() -> &dyn Any {
        static Data_Functor_Functorusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Functorusd_Dict.get_or_init(||
                                                     &Func1::new(move |x|
                                                                     x.clone()))
    }
    pub fn Data_Functor_map() -> &dyn Any {
        static Data_Functor_map: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_map.get_or_init(||
                                         &Func1::new(move |dict|
                                                         find(string("map"),
                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Functor_mapFlipped() -> &dyn Any {
        static Data_Functor_mapFlipped: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_mapFlipped.get_or_init(||
                                                &Func1::new(move |dictFunctor|
                                                                &Func1::new({
                                                                                let dictFunctor
                                                                                    =
                                                                                    dictFunctor.clone();
                                                                                move
                                                                                    |fa|
                                                                                    &Func1::new({
                                                                                                    let fa
                                                                                                        =
                                                                                                        fa.clone();
                                                                                                    move
                                                                                                        |f|
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                               &&&dictFunctor),
                                                                                                                                                                            f),
                                                                                                                                         &&&fa)
                                                                                                })
                                                                            })))
    }
    pub fn Data_Functor_void() -> &dyn Any {
        static Data_Functor_void: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_void.get_or_init(||
                                          &Func1::new(move |dictFunctor|
                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                              dictFunctor),
                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                              &&&PureScript_Data_Unit::Data_Unit_unit()))))
    }
    pub fn Data_Functor_voidLeft() -> &dyn Any {
        static Data_Functor_voidLeft: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_voidLeft.get_or_init(||
                                              &Func1::new(move |dictFunctor|
                                                              &Func1::new({
                                                                              let dictFunctor
                                                                                  =
                                                                                  dictFunctor.clone();
                                                                              move
                                                                                  |f|
                                                                                  &Func1::new({
                                                                                                  let f
                                                                                                      =
                                                                                                      f.clone();
                                                                                                  move
                                                                                                      |x|
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                             &&&dictFunctor),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                             x)),
                                                                                                                                       &&&f)
                                                                                              })
                                                                          })))
    }
    pub fn Data_Functor_voidRight() -> &dyn Any {
        static Data_Functor_voidRight: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_voidRight.get_or_init(||
                                               &Func1::new(move |dictFunctor|
                                                               &Func1::new({
                                                                               let dictFunctor
                                                                                   =
                                                                                   dictFunctor.clone();
                                                                               move
                                                                                   |x|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                       &&&dictFunctor),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                       x))
                                                                           })))
    }
    pub fn Data_Functor_functorProxy() -> &dyn Any {
        static Data_Functor_functorProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_functorProxy.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                   &&&add(string("map"),
                                                                                          &&Func1::new(move
                                                                                                           |v|
                                                                                                           &Func1::new(move
                                                                                                                           |v1|
                                                                                                                           &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Data_Functor_functorFn() -> &dyn Any {
        static Data_Functor_functorFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_functorFn.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                &&&add(string("map"),
                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                       empty::<string,
                                                                                               &dyn Any>())))
    }
    pub fn Data_Functor_functorArray() -> &dyn Any {
        static Data_Functor_functorArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_functorArray.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                   &&&add(string("map"),
                                                                                          &&PureScript_Data_Functor::Data_Functor_arrayMap(),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Data_Functor_flap() -> &dyn Any {
        static Data_Functor_flap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_flap.get_or_init(||
                                          &Func1::new(move |dictFunctor|
                                                          &Func1::new({
                                                                          let dictFunctor
                                                                              =
                                                                              dictFunctor.clone();
                                                                          move
                                                                              |ff|
                                                                              &Func1::new({
                                                                                              let ff
                                                                                                  =
                                                                                                  ff.clone();
                                                                                              move
                                                                                                  |x|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                         &&&dictFunctor),
                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                        let x
                                                                                                                                                                                            =
                                                                                                                                                                                            x.clone();
                                                                                                                                                                                        move
                                                                                                                                                                                            |f|
                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                                             &&&x)
                                                                                                                                                                                    })),
                                                                                                                                   &&&ff)
                                                                                          })
                                                                      })))
    }
}
