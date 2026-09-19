pub mod PureScript_Effect_Ref {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Effect_Ref_FFI {
        use super::*;
        use fable_library_rust::Map_::find;
        use fable_library_rust::Monitor_::lock;
        use fable_library_rust::Native_::Func0;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::Native_::refCell;
        pub fn _new<a: Clone + 'static, b: Clone + 'static>(s: a, _arg: b)
         -> &dyn Any {
            &refCell(s)
        }
        pub fn newWithSelf<a: Clone + 'static, b: Clone +
                           'static>(f: a, _arg: b) -> &dyn Any {
            let r = refCell(&defaultOf());
            let s =
                Sharpurs_Prelude::unbox(&{
                                             let arg = &r;
                                             (Sharpurs_Prelude::unbox(&&f))(arg)
                                         });
            r.set(s);
            &r
        }
        pub fn read<a: Clone + 'static, b: Clone + 'static>(r: a, _arg: b)
         -> &dyn Any {
            let rRef = Sharpurs_Prelude::unbox(&&r);
            lock(rRef.clone(),
                 Func0::new({ let rRef = rRef.clone(); move || rRef.get() }))
        }
        pub fn write<a: Clone + 'static, b: Clone +
                     'static>(s: &dyn Any, r: a, _arg: b) -> &dyn Any {
            let rRef = Sharpurs_Prelude::unbox(&&r);
            lock(rRef.clone(),
                 Func0::new({
                                let rRef = rRef.clone();
                                let s = s.clone();
                                move || rRef.set(s)
                            }));
            &defaultOf()
        }
        pub fn modifyImpl<a: Clone + 'static, b: Clone + 'static, c: Clone +
                          'static>(f: a, r: b, _arg: c) -> &dyn Any {
            let rRef = Sharpurs_Prelude::unbox(&&r);
            let f_ =
                {
                    let clo = Sharpurs_Prelude::unbox(&&f);
                    Func1::new({
                                   let clo = clo.clone();
                                   move |arg| clo(arg.clone())
                               })
                };
            lock(rRef.clone(),
                 Func0::new({
                                let f_ = f_.clone();
                                let rRef = rRef.clone();
                                move ||
                                    {
                                        let result = f_(rRef.get());
                                        rRef.set(find(string("state"),
                                                      result.clone()));
                                        find(string("value"), result)
                                    }
                            }))
        }
    }
    pub fn Effect_Ref__new() -> &dyn Any {
        static Effect_Ref__new: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Ref__new.get_or_init(||
                                        &Func1::new(move |arg0|
                                                        &{
                                                             let s =
                                                                 Sharpurs_Prelude::unbox(arg0);
                                                             Func1::new({
                                                                            let s
                                                                                =
                                                                                s.clone();
                                                                            move
                                                                                |arg10_0040|
                                                                                PureScript_Effect_Ref::Effect_Ref_FFI::_new(s,
                                                                                                                            arg10_0040.clone())
                                                                        })
                                                         }))
    }
    pub fn Effect_Ref_modifyImpl() -> &dyn Any {
        static Effect_Ref_modifyImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Ref_modifyImpl.get_or_init(||
                                              &Func1::new(move |arg0|
                                                              &Func1::new({
                                                                              let arg0
                                                                                  =
                                                                                  arg0.clone();
                                                                              move
                                                                                  |arg1|
                                                                                  &{
                                                                                       let f =
                                                                                           Sharpurs_Prelude::unbox(&arg0);
                                                                                       let r =
                                                                                           Sharpurs_Prelude::unbox(arg1);
                                                                                       Func1::new({
                                                                                                      let f
                                                                                                          =
                                                                                                          f.clone();
                                                                                                      let r
                                                                                                          =
                                                                                                          r.clone();
                                                                                                      move
                                                                                                          |arg20_0040|
                                                                                                          PureScript_Effect_Ref::Effect_Ref_FFI::modifyImpl(f,
                                                                                                                                                            r,
                                                                                                                                                            arg20_0040.clone())
                                                                                                  })
                                                                                   }
                                                                          })))
    }
    pub fn Effect_Ref_newWithSelf() -> &dyn Any {
        static Effect_Ref_newWithSelf: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Ref_newWithSelf.get_or_init(||
                                               &Func1::new(move |arg0|
                                                               &{
                                                                    let f =
                                                                        Sharpurs_Prelude::unbox(arg0);
                                                                    Func1::new({
                                                                                   let f
                                                                                       =
                                                                                       f.clone();
                                                                                   move
                                                                                       |arg10_0040|
                                                                                       PureScript_Effect_Ref::Effect_Ref_FFI::newWithSelf(f,
                                                                                                                                          arg10_0040.clone())
                                                                               })
                                                                }))
    }
    pub fn Effect_Ref_read() -> &dyn Any {
        static Effect_Ref_read: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Ref_read.get_or_init(||
                                        &Func1::new(move |arg0|
                                                        &{
                                                             let r =
                                                                 Sharpurs_Prelude::unbox(arg0);
                                                             Func1::new({
                                                                            let r
                                                                                =
                                                                                r.clone();
                                                                            move
                                                                                |arg10_0040|
                                                                                PureScript_Effect_Ref::Effect_Ref_FFI::read(r,
                                                                                                                            arg10_0040.clone())
                                                                        })
                                                         }))
    }
    pub fn Effect_Ref_write() -> &dyn Any {
        static Effect_Ref_write: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Ref_write.get_or_init(||
                                         &Func1::new(move |arg0|
                                                         &Func1::new({
                                                                         let arg0
                                                                             =
                                                                             arg0.clone();
                                                                         move
                                                                             |arg1|
                                                                             &{
                                                                                  let s =
                                                                                      Sharpurs_Prelude::unbox(&arg0);
                                                                                  let r =
                                                                                      Sharpurs_Prelude::unbox(arg1);
                                                                                  Func1::new({
                                                                                                 let r
                                                                                                     =
                                                                                                     r.clone();
                                                                                                 let s
                                                                                                     =
                                                                                                     s.clone();
                                                                                                 move
                                                                                                     |arg20_0040|
                                                                                                     PureScript_Effect_Ref::Effect_Ref_FFI::write(&s,
                                                                                                                                                  r,
                                                                                                                                                  arg20_0040.clone())
                                                                                             })
                                                                              }
                                                                     })))
    }
    pub fn Effect_Ref_void() -> &dyn Any {
        static Effect_Ref_void: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Ref_void.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                         &&&PureScript_Effect::Effect_functorEffect()))
    }
    pub fn Effect_Ref_new() -> &dyn Any {
        static Effect_Ref_new: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Effect_Ref_new.get_or_init(||
                                       &PureScript_Effect_Ref::Effect_Ref__new())
    }
    pub fn Effect_Ref_modify_prime() -> &dyn Any {
        static Effect_Ref_modify_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Ref_modify_prime.get_or_init(||
                                                &PureScript_Effect_Ref::Effect_Ref_modifyImpl())
    }
    pub fn Effect_Ref_modify() -> &dyn Any {
        static Effect_Ref_modify: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Ref_modify.get_or_init(||
                                          &Func1::new(move |f|
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Ref::Effect_Ref_modify_prime(),
                                                                                           &&&Func1::new({
                                                                                                             let f
                                                                                                                 =
                                                                                                                 f.clone();
                                                                                                             move
                                                                                                                 |s|
                                                                                                                 {
                                                                                                                     let s_prime =
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                          s);
                                                                                                                     &add(string("state"),
                                                                                                                          &&s_prime,
                                                                                                                          add(string("value"),
                                                                                                                              &&s_prime,
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>()))
                                                                                                                 }
                                                                                                         }))))
    }
    pub fn Effect_Ref_modify_() -> &dyn Any {
        static Effect_Ref_modify_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Ref_modify_.get_or_init(||
                                           &Func1::new(move |f|
                                                           &Func1::new({
                                                                           let f
                                                                               =
                                                                               f.clone();
                                                                           move
                                                                               |s|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                   &&&PureScript_Effect_Ref::Effect_Ref_void()),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Ref::Effect_Ref_modify(),
                                                                                                                                                                                      &&&f),
                                                                                                                                                   s))
                                                                       })))
    }
}
