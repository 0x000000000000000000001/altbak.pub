pub mod PureScript_Test_PrimesFFICheatcode {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Test_PrimesFFICheatcode_FFI {
        use super::*;
        pub fn isPrime(n: i32) -> bool {
            if n < 2_i32 {
                false
            } else {
                let p: MutCell<bool> = MutCell::new(true);
                let i: MutCell<i32> = MutCell::new(2_i32);
                while if p.get() { i.get() * i.get() <= n } else { false } {
                    if n % i.get() == 0_i32 { p.set(false); }
                    i.set(i.get() + 1_i32)
                }
                p.get()
            }
        }
        pub fn runPrimesFFICheatcode(n: &dyn Any) -> &dyn Any {
            let n_: i32 = Sharpurs_Prelude::unbox(n);
            let sum: MutCell<i32> = MutCell::new(0_i32);
            for i in 2_i32..=n_ {
                if PureScript_Test_PrimesFFICheatcode::Test_PrimesFFICheatcode_FFI::isPrime(i)
                   {
                    sum.set(sum.get() + i);
                };
            }
            &sum.get()
        }
    }
    pub fn Test_PrimesFFICheatcode_runPrimesFFICheatcode() -> &dyn Any {
        static Test_PrimesFFICheatcode_runPrimesFFICheatcode:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_PrimesFFICheatcode_runPrimesFFICheatcode.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |arg0|
                                                                                      &PureScript_Test_PrimesFFICheatcode::Test_PrimesFFICheatcode_FFI::runPrimesFFICheatcode(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_PrimesFFICheatcode_describe() -> &dyn Any {
        static Test_PrimesFFICheatcode_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_PrimesFFICheatcode_describe.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                          &&&string("Prime Sieve FFICheatcode (sum primes up to 500):")))
    }
    pub fn Test_PrimesFFICheatcode_act() -> &dyn Any {
        static Test_PrimesFFICheatcode_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_PrimesFFICheatcode_act.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                           &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                           &&&500_i32)),
                                                                                     &&&Func1::new(move
                                                                                                       |dummy|
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                           &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                              &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_PrimesFFICheatcode::Test_PrimesFFICheatcode_runPrimesFFICheatcode(),
                                                                                                                                                                                                              dummy))))))
    }
}
