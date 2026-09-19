pub mod Sharpurs_Prelude {
    use super::*;
    use fable_library_rust::Array_::find;
    use fable_library_rust::Exception_::try_catch;
    use fable_library_rust::Map_::Map;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::tryFind;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::defaultOf;
    use fable_library_rust::NativeArray_::count;
    use fable_library_rust::Reflection_::name;
    use fable_library_rust::String_::string;
    use fable_library_rust::System::Reflection::TargetInvocationException;
    use fable_library_rust::System::Exception;
    use fable_library_rust::System::Reflection::MethodInfo;
    pub fn _007cLitBool_007c__007c(expected: bool, value: &dyn Any)
     -> Option<()> {
        if if let Some(value) = (value as &dyn Any).downcast_ref::<bool>() {
               value.clone() == expected
           } else { false } {
            Some(())
        } else { None::<()> }
    }
    pub fn _007cLitInt_007c__007c(expected: i32, value: &dyn Any)
     -> Option<()> {
        if if let Some(value) = (value as &dyn Any).downcast_ref::<i32>() {
               value.clone() == expected
           } else { false } {
            Some(())
        } else { None::<()> }
    }
    pub fn _007cLitNumber_007c__007c(expected: f64, value: &dyn Any)
     -> Option<()> {
        if if let Some(value) = (value as &dyn Any).downcast_ref::<f64>() {
               value.clone() == expected
           } else { false } {
            Some(())
        } else { None::<()> }
    }
    pub fn _007cLitString_007c__007c(expected: string, value: &dyn Any)
     -> Option<()> {
        if if let Some(value) = (value as &dyn Any).downcast_ref::<string>() {
               value.clone() == expected
           } else { false } {
            Some(())
        } else { None::<()> }
    }
    pub fn _007cLitChar_007c__007c(expected: char, value: &dyn Any)
     -> Option<()> {
        if if let Some(value) = (value as &dyn Any).downcast_ref::<char>() {
               value.clone() == expected
           } else { false } {
            Some(())
        } else { None::<()> }
    }
    pub fn _007cHasProp_007c__007c(key: string, value: &dyn Any)
     -> Option<&dyn Any> {
        if let Some(value) =
               (value as &dyn Any).downcast_ref::<Map<string, &dyn Any>>() {
            tryFind(key, value.clone())
        } else { None::<&dyn Any> }
    }
    pub mod SharpursRuntime {
        use super::*;
        use fable_library_rust::Native_::compare;
        pub fn eventLoopWg() -> &dyn Any {
            static eventLoopWg: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            eventLoopWg.get_or_init(|| defaultOf::<&dyn Any>())
        }
        pub fn loopHasTasks() -> LrcPtr<MutCell<bool>> {
            static loopHasTasks: MutCell<Option<LrcPtr<MutCell<bool>>>> =
                MutCell::new(None);
            loopHasTasks.get_or_init(|| LrcPtr::new(MutCell::new(false)))
        }
        pub fn EventLoopAdd(n: i32) -> &dyn Any {
            Sharpurs_Prelude::SharpursRuntime::loopHasTasks().set(true);
            defaultOf::<&dyn Any>()
        }
        pub fn EventLoopDone() {
            if compare(defaultOf::<&dyn Any>(), 0_i32) > 0_i32 { (); };
        }
        pub fn EventLoopWait() -> &dyn Any {
            if compare(defaultOf::<&dyn Any>(), 0_i32) > 0_i32 { (); }
            if Sharpurs_Prelude::SharpursRuntime::loopHasTasks().get() {
                defaultOf::<&dyn Any>();
            }
        }
    }
    pub fn objMap() -> Map<string, &dyn Any> {
        static objMap: MutCell<Option<Map<string, &dyn Any>>> =
            MutCell::new(None);
        objMap.get_or_init(|| empty::<string, &dyn Any>())
    }
    pub fn unbox<a: Clone + 'static>(x: &dyn Any) -> a {
        a::from(x.clone()).clone()
    }
    pub fn _007cUnbox_007c<a: Clone + 'static>(x: &dyn Any) -> a {
        Sharpurs_Prelude::unbox(x)
    }
    pub fn undefined() -> &dyn Any {
        static undefined: MutCell<Option<&dyn Any>> = MutCell::new(None);
        undefined.get_or_init(|| defaultOf())
    }
    pub fn Prim_undefined() -> &dyn Any {
        static Prim_undefined: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Prim_undefined.get_or_init(|| Sharpurs_Prelude::undefined())
    }
    pub fn intMod<a: Clone + 'static, b: Clone + 'static>(a: a, b: b) -> i32 {
        Sharpurs_Prelude::unbox(&&a) % Sharpurs_Prelude::unbox(&&b)
    }
    pub fn semiringInt() -> i32 {
        static semiringInt: MutCell<Option<i32>> = MutCell::new(None);
        semiringInt.get_or_init(|| 0_i32)
    }
    pub fn sharpurs_int_mod(left: i32, right: i32) -> i32 {
        if if right == 0_i32 { true } else { right == -1_i32 } {
            0_i32
        } else {
            let remainder: i32 = left % right;
            if remainder < 0_i32 {
                if right > 0_i32 {
                    remainder + right
                } else { remainder - right }
            } else { remainder }
        }
    }
    pub fn sharpurs_apply(func: &dyn Any, arg: &dyn Any) -> &dyn Any {
        if func.clone() == defaultOf::<&dyn Any>() {
            panic!("{}", string("sharpurs_apply: func is null!"),);
        }
        if let Some(func) =
               (func as &dyn Any).downcast_ref::<Func1<&dyn Any, &dyn Any>>()
           {
            try_catch(|| func(arg.clone()),
                      |ex: LrcPtr<Exception>|
                          panic!("{}",
                                 TargetInvocationException::_ctor__229D3F39(ex.clone()).get_Message(),))
        } else {
            let method: LrcPtr<dyn MethodInfo> =
                find(Func1::new(move |m: LrcPtr<dyn MethodInfo>|
                                    if name(m.clone()) == string("Invoke") {
                                        count(m[string("parameters")]) ==
                                            1_i32
                                    } else { false }),
                     defaultOf::<&dyn Any>());
            defaultOf::<&dyn Any>()
        }
    }
}
