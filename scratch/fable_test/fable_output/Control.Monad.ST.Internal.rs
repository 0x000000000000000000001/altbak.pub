pub mod PureScript_Control_Monad_ST_Internal {
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
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub mod Control_Monad_ST_Internal_FFI {
        use super::*;
        use fable_library_rust::Native_::Func0;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::Native_::refCell;
        use fable_library_rust::NativeArray_::count;
        pub fn map_(f: &dyn Any, a: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let a = a.clone();
                            let f = f.clone();
                            move || Sharpurs_Prelude::sharpurs_apply(&f, &a())
                        })
        }
        pub fn bind_(a: &dyn Any, f: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let a = a.clone();
                            let f = f.clone();
                            move ||
                                (Sharpurs_Prelude::sharpurs_apply(&f, &a()))()
                        })
        }
        pub fn pure_(a: &dyn Any) -> &dyn Any {
            &Func0::new({ let a = a.clone(); move || a })
        }
        pub fn run(f: &dyn Any) -> &dyn Any { f() }
        pub fn r#while(cond: &dyn Any, a: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let a = a.clone();
                            let cond = cond.clone();
                            move ||
                                {
                                    while Sharpurs_Prelude::unbox(&cond()) {
                                        let value = a();
                                        ()
                                    }
                                    defaultOf()
                                }
                        })
        }
        pub fn new(val_: &dyn Any) -> &dyn Any {
            &Func0::new({ let val_ = val_.clone(); move || &refCell(val_) })
        }
        pub fn read(r: &dyn Any) -> &dyn Any {
            &Func0::new({ let r = r.clone(); move || r.get() })
        }
        pub fn modifyImpl(f: &dyn Any, r: &dyn Any) -> &dyn Any {
            let r_ = r.clone();
            &Func0::new({
                            let f = f.clone();
                            let r_ = r_.clone();
                            move ||
                                {
                                    let res =
                                        Sharpurs_Prelude::sharpurs_apply(&f,
                                                                         &r_.get());
                                    r_.set(find(string("state"),
                                                res.clone()));
                                    find(string("value"), res)
                                }
                        })
        }
        pub fn write(a: &dyn Any, r: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let a = a.clone();
                            let r = r.clone();
                            move || { r = a; a }
                        })
        }
        pub fn r#for(lo: &dyn Any, hi: &dyn Any, f: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let f = f.clone();
                            let hi = hi.clone();
                            let lo = lo.clone();
                            move ||
                                {
                                    for i in lo..=hi - 1_i32 {
                                        let value =
                                            (Sharpurs_Prelude::sharpurs_apply(&f,
                                                                              &&i))();
                                        ()
                                    }
                                    defaultOf()
                                }
                        })
        }
        pub fn foreach(xs: &dyn Any, f: &dyn Any) -> &dyn Any {
            let arr = xs.clone();
            &Func0::new({
                            let arr = arr.clone();
                            let f = f.clone();
                            move ||
                                {
                                    for idx in
                                        0_i32..=count(arr.clone()) - 1_i32 {
                                        let value =
                                            (Sharpurs_Prelude::sharpurs_apply(&f,
                                                                              &arr[idx].clone()))();
                                        ()
                                    }
                                    defaultOf()
                                }
                        })
        }
    }
    pub fn Control_Monad_ST_Internal_bind_() -> &dyn Any {
        static Control_Monad_ST_Internal_bind_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_bind_.get_or_init(||
                                                        &Func1::new(move |a|
                                                                        Func1::new({
                                                                                       let a
                                                                                           =
                                                                                           a.clone();
                                                                                       move
                                                                                           |f|
                                                                                           PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::bind_(&a,
                                                                                                                                                                      f)
                                                                                   })))
    }
    pub fn Control_Monad_ST_Internal_for() -> &dyn Any {
        static Control_Monad_ST_Internal_for: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_for.get_or_init(||
                                                      &Func1::new(move |lo|
                                                                      Func1::new({
                                                                                     let lo
                                                                                         =
                                                                                         lo.clone();
                                                                                     move
                                                                                         |hi|
                                                                                         Func1::new({
                                                                                                        let hi
                                                                                                            =
                                                                                                            hi.clone();
                                                                                                        move
                                                                                                            |f|
                                                                                                            PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::r#for(&lo,
                                                                                                                                                                                       &hi,
                                                                                                                                                                                       f)
                                                                                                    })
                                                                                 })))
    }
    pub fn Control_Monad_ST_Internal_foreach() -> &dyn Any {
        static Control_Monad_ST_Internal_foreach: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_foreach.get_or_init(||
                                                          &Func1::new(move
                                                                          |xs|
                                                                          Func1::new({
                                                                                         let xs
                                                                                             =
                                                                                             xs.clone();
                                                                                         move
                                                                                             |f|
                                                                                             PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::foreach(&xs,
                                                                                                                                                                          f)
                                                                                     })))
    }
    pub fn Control_Monad_ST_Internal_map_() -> &dyn Any {
        static Control_Monad_ST_Internal_map_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_map_.get_or_init(||
                                                       &Func1::new(move |f|
                                                                       Func1::new({
                                                                                      let f
                                                                                          =
                                                                                          f.clone();
                                                                                      move
                                                                                          |a|
                                                                                          PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::map_(&f,
                                                                                                                                                                    a)
                                                                                  })))
    }
    pub fn Control_Monad_ST_Internal_modifyImpl() -> &dyn Any {
        static Control_Monad_ST_Internal_modifyImpl: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_ST_Internal_modifyImpl.get_or_init(||
                                                             &Func1::new(move
                                                                             |f|
                                                                             Func1::new({
                                                                                            let f
                                                                                                =
                                                                                                f.clone();
                                                                                            move
                                                                                                |r|
                                                                                                PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::modifyImpl(&f,
                                                                                                                                                                                r)
                                                                                        })))
    }
    pub fn Control_Monad_ST_Internal_new() -> &dyn Any {
        static Control_Monad_ST_Internal_new: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_new.get_or_init(||
                                                      &Func1::new(move |val_|
                                                                      PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::new(val_)))
    }
    pub fn Control_Monad_ST_Internal_pure_() -> &dyn Any {
        static Control_Monad_ST_Internal_pure_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_pure_.get_or_init(||
                                                        &Func1::new(move |a|
                                                                        PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::pure_(a)))
    }
    pub fn Control_Monad_ST_Internal_read() -> &dyn Any {
        static Control_Monad_ST_Internal_read: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_read.get_or_init(||
                                                       &Func1::new(move |r|
                                                                       PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::read(r)))
    }
    pub fn Control_Monad_ST_Internal_run() -> &dyn Any {
        static Control_Monad_ST_Internal_run: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_run.get_or_init(||
                                                      &Func1::new(move |f|
                                                                      PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::run(f)))
    }
    pub fn Control_Monad_ST_Internal_while() -> &dyn Any {
        static Control_Monad_ST_Internal_while: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_while.get_or_init(||
                                                        &Func1::new(move
                                                                        |cond|
                                                                        Func1::new({
                                                                                       let cond
                                                                                           =
                                                                                           cond.clone();
                                                                                       move
                                                                                           |a|
                                                                                           PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::r#while(&cond,
                                                                                                                                                                        a)
                                                                                   })))
    }
    pub fn Control_Monad_ST_Internal_write() -> &dyn Any {
        static Control_Monad_ST_Internal_write: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_write.get_or_init(||
                                                        &Func1::new(move |a|
                                                                        Func1::new({
                                                                                       let a
                                                                                           =
                                                                                           a.clone();
                                                                                       move
                                                                                           |r|
                                                                                           PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_FFI::write(&a,
                                                                                                                                                                      r)
                                                                                   })))
    }
    pub fn Control_Monad_ST_Internal_modify_prime() -> &dyn Any {
        static Control_Monad_ST_Internal_modify_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_modify_prime.get_or_init(||
                                                               &PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_modifyImpl())
    }
    pub fn Control_Monad_ST_Internal_modify() -> &dyn Any {
        static Control_Monad_ST_Internal_modify: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_modify.get_or_init(||
                                                         &Func1::new(move |f|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_modify_prime(),
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
    pub fn Control_Monad_ST_Internal_functorST() -> &dyn Any {
        static Control_Monad_ST_Internal_functorST: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_ST_Internal_functorST.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                             &&&add(string("map"),
                                                                                                    &&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_map_(),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Control_Monad_ST_Internal_monadST_0040101() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                         &&&add(string("Applicative0"),
                                                &&Func1::new({
                                                                 let Control_Monad_ST_Internal_applicativeST_0040104_002d1
                                                                     =
                                                                     Control_Monad_ST_Internal_applicativeST_0040104_002d1.clone();
                                                                 move
                                                                     |usd__unused|
                                                                     &Control_Monad_ST_Internal_applicativeST_0040104_002d1.Value
                                                             }),
                                                add(string("Bind1"),
                                                    &&Func1::new({
                                                                     let Control_Monad_ST_Internal_bindST_0040102_002d1
                                                                         =
                                                                         Control_Monad_ST_Internal_bindST_0040102_002d1.clone();
                                                                     move
                                                                         |usd__unused_1|
                                                                         &Control_Monad_ST_Internal_bindST_0040102_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_ST_Internal_monadST_0040101_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_ST_Internal_monadST_0040101_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_monadST_0040101_002d1.get_or_init(||
                                                                        Lazy(Control_Monad_ST_Internal_monadST_0040101.clone()))
    }
    pub fn Control_Monad_ST_Internal_bindST_0040102() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                         &&&add(string("bind"),
                                                &&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bind_(),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let Control_Monad_ST_Internal_applyST_0040103_002d1
                                                                         =
                                                                         Control_Monad_ST_Internal_applyST_0040103_002d1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &Control_Monad_ST_Internal_applyST_0040103_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_ST_Internal_bindST_0040102_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_ST_Internal_bindST_0040102_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_bindST_0040102_002d1.get_or_init(||
                                                                       Lazy(Control_Monad_ST_Internal_bindST_0040102.clone()))
    }
    pub fn Control_Monad_ST_Internal_applyST_0040103() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                         &&&add(string("apply"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_ap(),
                                                                                  &&&Control_Monad_ST_Internal_monadST_0040101_002d1.Value),
                                                add(string("Functor0"),
                                                    &&Func1::new(move
                                                                     |usd__unused|
                                                                     &PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_ST_Internal_applyST_0040103_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_ST_Internal_applyST_0040103_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_applyST_0040103_002d1.get_or_init(||
                                                                        Lazy(Control_Monad_ST_Internal_applyST_0040103.clone()))
    }
    pub fn Control_Monad_ST_Internal_applicativeST_0040104() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                         &&&add(string("pure"),
                                                &&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_pure_(),
                                                add(string("Apply0"),
                                                    &&Func1::new({
                                                                     let Control_Monad_ST_Internal_applyST_0040103_002d1
                                                                         =
                                                                         Control_Monad_ST_Internal_applyST_0040103_002d1.clone();
                                                                     move
                                                                         |usd__unused|
                                                                         &Control_Monad_ST_Internal_applyST_0040103_002d1.Value
                                                                 }),
                                                    empty::<string,
                                                            &dyn Any>())))
    }
    pub fn Control_Monad_ST_Internal_applicativeST_0040104_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Control_Monad_ST_Internal_applicativeST_0040104_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_applicativeST_0040104_002d1.get_or_init(||
                                                                              Lazy(Control_Monad_ST_Internal_applicativeST_0040104.clone()))
    }
    pub fn Control_Monad_ST_Internal_monadST() -> &dyn Any {
        static Control_Monad_ST_Internal_monadST: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_monadST.get_or_init(||
                                                          Control_Monad_ST_Internal_monadST_0040101_002d1.Value)
    }
    pub fn Control_Monad_ST_Internal_bindST() -> &dyn Any {
        static Control_Monad_ST_Internal_bindST: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_bindST.get_or_init(||
                                                         Control_Monad_ST_Internal_bindST_0040102_002d1.Value)
    }
    pub fn Control_Monad_ST_Internal_applyST() -> &dyn Any {
        static Control_Monad_ST_Internal_applyST: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_applyST.get_or_init(||
                                                          Control_Monad_ST_Internal_applyST_0040103_002d1.Value)
    }
    pub fn Control_Monad_ST_Internal_applicativeST() -> &dyn Any {
        static Control_Monad_ST_Internal_applicativeST:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_applicativeST.get_or_init(||
                                                                Control_Monad_ST_Internal_applicativeST_0040104_002d1.Value)
    }
    pub fn Control_Monad_ST_Internal_semigroupST() -> &dyn Any {
        static Control_Monad_ST_Internal_semigroupST:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_semigroupST.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictSemigroup|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                               &&&add(string("append"),
                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                           &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applyST()),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                           dictSemigroup)),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))))
    }
    pub fn Control_Monad_ST_Internal_monadRecST() -> &dyn Any {
        static Control_Monad_ST_Internal_monadRecST: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_ST_Internal_monadRecST.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_MonadRecusd_Dict(),
                                                                                              &&&add(string("tailRecM"),
                                                                                                     &&Func1::new(move
                                                                                                                      |f|
                                                                                                                      &Func1::new({
                                                                                                                                      let f
                                                                                                                                          =
                                                                                                                                          f.clone();
                                                                                                                                      move
                                                                                                                                          |a|
                                                                                                                                          {
                                                                                                                                              let isLooping =
                                                                                                                                                  &Func1::new(move
                                                                                                                                                                  |v|
                                                                                                                                                                  if let Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(_)
                                                                                                                                                                         =
                                                                                                                                                                         Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      &true
                                                                                                                                                                  } else {
                                                                                                                                                                      &false
                                                                                                                                                                  });
                                                                                                                                              let fromDone =
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                     |usd__unused|
                                                                                                                                                                                                     &Func1::new(move
                                                                                                                                                                                                                     |v_1|
                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&Func1::new({
                                                                                                                                                                                                                                                                        let v_1
                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                            v_1.clone();
                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                            |usd__unused_1|
                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                let matchValue_1:
                                                                                                                                                                                                                                                                                        LrcPtr<Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                                                                                                                if let Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                                                                                                                       =
                                                                                                                                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                    &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                         Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(x)
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                                                     }
                                                                                                                                                                                                                                                                                } else {
                                                                                                                                                                                                                                                                                    panic!("{}",
                                                                                                                                                                                                                                                                                           string("Match failure"),)
                                                                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))));
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bindFlipped(),
                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_new()),
                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                        a))),
                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                 let fromDone
                                                                                                                                                                                                     =
                                                                                                                                                                                                     fromDone.clone();
                                                                                                                                                                                                 let isLooping
                                                                                                                                                                                                     =
                                                                                                                                                                                                     isLooping.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     |r|
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_while(),
                                                                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&&isLooping),
                                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_read(),
                                                                                                                                                                                                                                                                                                                                                                                                                     r))),
                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_read(),
                                                                                                                                                                                                                                                                                                                                                                                                                     r)),
                                                                                                                                                                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                                 let r
                                                                                                                                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                                                                                                                                     r.clone();
                                                                                                                                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                                                                                                                                     |v_2|
                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                         let matchValue_2:
                                                                                                                                                                                                                                                                                                                                                                                 LrcPtr<Control_Monad_Rec_Class_Step> =
                                                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                                                                                                                                                                                                         match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                             Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(matchValue_2_1_0)
                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                                                                                                                                                                                                                             Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(matchValue_2_0_0)
                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&matchValue_2_0_0)),
                                                                                                                                                                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                |e|
                                                                                                                                                                                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_write(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&r)))),
                                                                                                                                                                                                                                                                                                                                                                         }
                                                                                                                                                                                                                                                                                                                                                                     }
                                                                                                                                                                                                                                                                                                                                                             })))),
                                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                                        let r
                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                            r.clone();
                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                            |usd__unused_2|
                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                                                                                                                                                                                &&&fromDone),
                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_read(),
                                                                                                                                                                                                                                                                                                                                &&&r))
                                                                                                                                                                                                                                                    }))
                                                                                                                                                                                             }))
                                                                                                                                          }
                                                                                                                                  })),
                                                                                                     add(string("Monad0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused_3|
                                                                                                                          &PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_monadST()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
    pub fn Control_Monad_ST_Internal_monoidST() -> &dyn Any {
        static Control_Monad_ST_Internal_monoidST: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Internal_monoidST.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictMonoid|
                                                                           {
                                                                               let semigroupST1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_semigroupST(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                &&&add(string("mempty"),
                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                            &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                            dictMonoid)),
                                                                                                                       add(string("Semigroup0"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let semigroupST1
                                                                                                                                                =
                                                                                                                                                semigroupST1.clone();
                                                                                                                                            move
                                                                                                                                                |usd__unused|
                                                                                                                                                &semigroupST1
                                                                                                                                        }),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>())))
                                                                           }))
    }
}
