pub mod PureScript_Effect_Exception {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Effect_Exception_FFI {
        use super::*;
        use fable_library_rust::Exception_::try_catch;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::String_::isEmpty;
        use fable_library_rust::String_::toString;
        use fable_library_rust::System::Exception;
        pub fn showErrorImpl(errVal: &dyn Any) -> &dyn Any {
            &toString(errVal.clone())
        }
        pub fn error(msgVal: &dyn Any) -> &dyn Any {
            &Exception::_ctor__Z721C83C5(msgVal.clone())
        }
        pub fn errorWithCause(msgVal: &dyn Any, causeVal: &dyn Any)
         -> &dyn Any {
            &Exception::_ctor__68CE3CA2(msgVal.clone(), causeVal.clone())
        }
        pub fn errorWithName(nameVal: &dyn Any, msgVal: &dyn Any)
         -> &dyn Any {
            let ex = Exception::_ctor__Z721C83C5(msgVal.clone());
            defaultOf::<&dyn Any>().Add541DA560(&&string("Name"), nameVal);
            &ex
        }
        pub fn message(eVal: &dyn Any) -> &dyn Any { &eVal.get_Message() }
        pub fn name(eVal: &dyn Any) -> &dyn Any {
            if defaultOf::<&dyn Any>().Contains4E60E31B(&&string("Name")) {
                &defaultOf::<&dyn Any>().get_Item4E60E31B(&&string("Name"))
            } else { &string("Error") }
        }
        pub fn stackImpl(just: &dyn Any, nothing: &dyn Any, eVal: &dyn Any)
         -> &dyn Any {
            if !isEmpty(defaultOf::<&dyn Any>()) {
                just(&defaultOf::<&dyn Any>())
            } else { nothing.clone() }
        }
        pub fn throwException(eVal: &dyn Any, dummy: &dyn Any) -> &dyn Any {
            panic!("{}", eVal.get_Message(),);
            defaultOf()
        }
        pub fn catchException(cVal: &dyn Any, tVal: &dyn Any, dummy: &dyn Any)
         -> &dyn Any {
            let c = cVal.clone();
            try_catch(|| tVal(defaultOf()),
                      |matchValue: LrcPtr<Exception>|
                          if let Some(matchValue) =
                                 (matchValue as
                                      &dyn Any).downcast_ref::<LrcPtr<Exception>>()
                             {
                              c(&matchValue)(defaultOf())
                          } else {
                              c(&Exception::_ctor__Z721C83C5(toString(matchValue.clone())))(defaultOf())
                          })
        }
    }
    pub fn Effect_Exception_catchException() -> &dyn Any {
        static Effect_Exception_catchException: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_catchException.get_or_init(||
                                                        &Func1::new(move
                                                                        |cVal|
                                                                        Func1::new({
                                                                                       let cVal
                                                                                           =
                                                                                           cVal.clone();
                                                                                       move
                                                                                           |tVal|
                                                                                           Func1::new({
                                                                                                          let tVal
                                                                                                              =
                                                                                                              tVal.clone();
                                                                                                          move
                                                                                                              |dummy|
                                                                                                              PureScript_Effect_Exception::Effect_Exception_FFI::catchException(&cVal,
                                                                                                                                                                                &tVal,
                                                                                                                                                                                dummy)
                                                                                                      })
                                                                                   })))
    }
    pub fn Effect_Exception_error() -> &dyn Any {
        static Effect_Exception_error: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_error.get_or_init(||
                                               &Func1::new(move |msgVal|
                                                               PureScript_Effect_Exception::Effect_Exception_FFI::error(msgVal)))
    }
    pub fn Effect_Exception_errorWithCause() -> &dyn Any {
        static Effect_Exception_errorWithCause: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_errorWithCause.get_or_init(||
                                                        &Func1::new(move
                                                                        |msgVal|
                                                                        Func1::new({
                                                                                       let msgVal
                                                                                           =
                                                                                           msgVal.clone();
                                                                                       move
                                                                                           |causeVal|
                                                                                           PureScript_Effect_Exception::Effect_Exception_FFI::errorWithCause(&msgVal,
                                                                                                                                                             causeVal)
                                                                                   })))
    }
    pub fn Effect_Exception_errorWithName() -> &dyn Any {
        static Effect_Exception_errorWithName: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_errorWithName.get_or_init(||
                                                       &Func1::new(move
                                                                       |nameVal|
                                                                       Func1::new({
                                                                                      let nameVal
                                                                                          =
                                                                                          nameVal.clone();
                                                                                      move
                                                                                          |msgVal|
                                                                                          PureScript_Effect_Exception::Effect_Exception_FFI::errorWithName(&nameVal,
                                                                                                                                                           msgVal)
                                                                                  })))
    }
    pub fn Effect_Exception_message() -> &dyn Any {
        static Effect_Exception_message: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_message.get_or_init(||
                                                 &Func1::new(move |eVal|
                                                                 PureScript_Effect_Exception::Effect_Exception_FFI::message(eVal)))
    }
    pub fn Effect_Exception_name() -> &dyn Any {
        static Effect_Exception_name: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_name.get_or_init(||
                                              &Func1::new(move |eVal|
                                                              PureScript_Effect_Exception::Effect_Exception_FFI::name(eVal)))
    }
    pub fn Effect_Exception_showErrorImpl() -> &dyn Any {
        static Effect_Exception_showErrorImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_showErrorImpl.get_or_init(||
                                                       &Func1::new(move
                                                                       |errVal|
                                                                       PureScript_Effect_Exception::Effect_Exception_FFI::showErrorImpl(errVal)))
    }
    pub fn Effect_Exception_stackImpl() -> &dyn Any {
        static Effect_Exception_stackImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_stackImpl.get_or_init(||
                                                   &Func1::new(move |just|
                                                                   Func1::new({
                                                                                  let just
                                                                                      =
                                                                                      just.clone();
                                                                                  move
                                                                                      |nothing|
                                                                                      Func1::new({
                                                                                                     let nothing
                                                                                                         =
                                                                                                         nothing.clone();
                                                                                                     move
                                                                                                         |eVal|
                                                                                                         PureScript_Effect_Exception::Effect_Exception_FFI::stackImpl(&just,
                                                                                                                                                                      &nothing,
                                                                                                                                                                      eVal)
                                                                                                 })
                                                                              })))
    }
    pub fn Effect_Exception_throwException() -> &dyn Any {
        static Effect_Exception_throwException: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_throwException.get_or_init(||
                                                        &Func1::new(move
                                                                        |eVal|
                                                                        Func1::new({
                                                                                       let eVal
                                                                                           =
                                                                                           eVal.clone();
                                                                                       move
                                                                                           |dummy|
                                                                                           PureScript_Effect_Exception::Effect_Exception_FFI::throwException(&eVal,
                                                                                                                                                             dummy)
                                                                                   })))
    }
    pub fn Effect_Exception_pure() -> &dyn Any {
        static Effect_Exception_pure: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_pure.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                               &&&PureScript_Effect::Effect_applicativeEffect()))
    }
    pub fn Effect_Exception_try() -> &dyn Any {
        static Effect_Exception_try: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_try.get_or_init(||
                                             &Func1::new(move |action|
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Exception::Effect_Exception_catchException(),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                       &&&PureScript_Effect_Exception::Effect_Exception_pure()),
                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                      |usd__arg1|
                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone()))))),
                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                       &&&PureScript_Effect::Effect_functorEffect()),
                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                      |usd__arg1_1|
                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                 action))))
    }
    pub fn Effect_Exception_throw() -> &dyn Any {
        static Effect_Exception_throw: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_throw.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                   &&&PureScript_Effect_Exception::Effect_Exception_throwException()),
                                                                                &&&PureScript_Effect_Exception::Effect_Exception_error()))
    }
    pub fn Effect_Exception_stack() -> &dyn Any {
        static Effect_Exception_stack: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_stack.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Exception::Effect_Exception_stackImpl(),
                                                                                                                   &&&Func1::new(move
                                                                                                                                     |usd__arg1|
                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Effect_Exception_showError() -> &dyn Any {
        static Effect_Exception_showError: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_showError.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                    &&&add(string("show"),
                                                                                           &&PureScript_Effect_Exception::Effect_Exception_showErrorImpl(),
                                                                                           empty::<string,
                                                                                                   &dyn Any>())))
    }
}
