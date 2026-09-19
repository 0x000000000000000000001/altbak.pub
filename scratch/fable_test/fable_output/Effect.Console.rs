pub mod PureScript_Effect_Console {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Effect_Console_FFI {
        use super::*;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::String_::string;
        pub fn log(s: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let s = s.clone();
                            move |_arg|
                                {
                                    {
                                        let arg: string =
                                            Sharpurs_Prelude::unbox(&s);
                                        println!("{}", arg)
                                    }
                                    &defaultOf()
                                }
                        })
        }
        pub fn time<a: Clone + 'static>(_arg: a) -> &dyn Any { &defaultOf() }
        pub fn timeLog<a: Clone + 'static>(_arg: a) -> &dyn Any {
            &defaultOf()
        }
        pub fn timeEnd<a: Clone + 'static>(_arg: a) -> &dyn Any {
            &defaultOf()
        }
        pub fn info(s: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let s = s.clone();
                            move |_arg|
                                {
                                    {
                                        let arg: string =
                                            Sharpurs_Prelude::unbox(&s);
                                        println!("{}", arg)
                                    }
                                    &defaultOf()
                                }
                        })
        }
        pub fn groupEnd<a: Clone + 'static>(_arg: a) -> &dyn Any {
            &defaultOf()
        }
        pub fn groupCollapsed<a: Clone + 'static>(_arg: a) -> &dyn Any {
            &defaultOf()
        }
        pub fn group<a: Clone + 'static>(_arg: a) -> &dyn Any { &defaultOf() }
        pub fn error(s: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let s = s.clone();
                            move |_arg|
                                {
                                    {
                                        let arg: string =
                                            Sharpurs_Prelude::unbox(&s);
                                        eprintln!("{}", arg)
                                    }
                                    &defaultOf()
                                }
                        })
        }
        pub fn debug(s: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let s = s.clone();
                            move |_arg|
                                {
                                    {
                                        let arg: string =
                                            Sharpurs_Prelude::unbox(&s);
                                        println!("{}", arg)
                                    }
                                    &defaultOf()
                                }
                        })
        }
        pub fn clear<a: Clone + 'static>(_arg: a) -> &dyn Any { &defaultOf() }
        pub fn warn(s: &dyn Any) -> &dyn Any {
            &Func1::new({
                            let s = s.clone();
                            move |_arg|
                                {
                                    {
                                        let arg: string =
                                            Sharpurs_Prelude::unbox(&s);
                                        eprintln!("{}", arg)
                                    }
                                    &defaultOf()
                                }
                        })
        }
    }
    pub fn Effect_Console_clear() -> &dyn Any {
        static Effect_Console_clear: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_clear.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &PureScript_Effect_Console::Effect_Console_FFI::clear(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_debug() -> &dyn Any {
        static Effect_Console_debug: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_debug.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &PureScript_Effect_Console::Effect_Console_FFI::debug(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_error() -> &dyn Any {
        static Effect_Console_error: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_error.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &PureScript_Effect_Console::Effect_Console_FFI::error(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_group() -> &dyn Any {
        static Effect_Console_group: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_group.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &PureScript_Effect_Console::Effect_Console_FFI::group(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_groupCollapsed() -> &dyn Any {
        static Effect_Console_groupCollapsed: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_groupCollapsed.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Effect_Console::Effect_Console_FFI::groupCollapsed(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_groupEnd() -> &dyn Any {
        static Effect_Console_groupEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_groupEnd.get_or_init(||
                                                &Func1::new(move |arg0|
                                                                &PureScript_Effect_Console::Effect_Console_FFI::groupEnd(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_info() -> &dyn Any {
        static Effect_Console_info: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_info.get_or_init(||
                                            &Func1::new(move |arg0|
                                                            &PureScript_Effect_Console::Effect_Console_FFI::info(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_log() -> &dyn Any {
        static Effect_Console_log: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_log.get_or_init(||
                                           &Func1::new(move |arg0|
                                                           &PureScript_Effect_Console::Effect_Console_FFI::log(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_time() -> &dyn Any {
        static Effect_Console_time: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_time.get_or_init(||
                                            &Func1::new(move |arg0|
                                                            &PureScript_Effect_Console::Effect_Console_FFI::time(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_timeEnd() -> &dyn Any {
        static Effect_Console_timeEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_timeEnd.get_or_init(||
                                               &Func1::new(move |arg0|
                                                               &PureScript_Effect_Console::Effect_Console_FFI::timeEnd(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_timeLog() -> &dyn Any {
        static Effect_Console_timeLog: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_timeLog.get_or_init(||
                                               &Func1::new(move |arg0|
                                                               &PureScript_Effect_Console::Effect_Console_FFI::timeLog(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_warn() -> &dyn Any {
        static Effect_Console_warn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_warn.get_or_init(||
                                            &Func1::new(move |arg0|
                                                            &PureScript_Effect_Console::Effect_Console_FFI::warn(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Console_warnShow() -> &dyn Any {
        static Effect_Console_warnShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_warnShow.get_or_init(||
                                                &Func1::new(move |dictShow|
                                                                &Func1::new({
                                                                                let dictShow
                                                                                    =
                                                                                    dictShow.clone();
                                                                                move
                                                                                    |a|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_warn(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                           &&&dictShow),
                                                                                                                                                        a))
                                                                            })))
    }
    pub fn Effect_Console_logShow() -> &dyn Any {
        static Effect_Console_logShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_logShow.get_or_init(||
                                               &Func1::new(move |dictShow|
                                                               &Func1::new({
                                                                               let dictShow
                                                                                   =
                                                                                   dictShow.clone();
                                                                               move
                                                                                   |a|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                          &&&dictShow),
                                                                                                                                                       a))
                                                                           })))
    }
    pub fn Effect_Console_infoShow() -> &dyn Any {
        static Effect_Console_infoShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_infoShow.get_or_init(||
                                                &Func1::new(move |dictShow|
                                                                &Func1::new({
                                                                                let dictShow
                                                                                    =
                                                                                    dictShow.clone();
                                                                                move
                                                                                    |a|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_info(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                           &&&dictShow),
                                                                                                                                                        a))
                                                                            })))
    }
    pub fn Effect_Console_grouped() -> &dyn Any {
        static Effect_Console_grouped: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_grouped.get_or_init(||
                                               &Func1::new(move |name|
                                                               &Func1::new({
                                                                               let name
                                                                                   =
                                                                                   name.clone();
                                                                               move
                                                                                   |inner|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                             &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                          &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_group(),
                                                                                                                                                                                          &&&name)),
                                                                                                                    &&&Func1::new({
                                                                                                                                      let inner
                                                                                                                                          =
                                                                                                                                          inner.clone();
                                                                                                                                      move
                                                                                                                                          |usd__unused|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                 &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                                              &&&inner),
                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                             |result|
                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                    &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                                                                                                                                                                 &&&PureScript_Effect_Console::Effect_Console_groupEnd()),
                                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                                let result
                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                    result.clone();
                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                    |usd__unused_1|
                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                        &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                                                                                                                                                                     &&&result)
                                                                                                                                                                                                                                            }))))
                                                                                                                                  }))
                                                                           })))
    }
    pub fn Effect_Console_errorShow() -> &dyn Any {
        static Effect_Console_errorShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_errorShow.get_or_init(||
                                                 &Func1::new(move |dictShow|
                                                                 &Func1::new({
                                                                                 let dictShow
                                                                                     =
                                                                                     dictShow.clone();
                                                                                 move
                                                                                     |a|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_error(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                            &&&dictShow),
                                                                                                                                                         a))
                                                                             })))
    }
    pub fn Effect_Console_debugShow() -> &dyn Any {
        static Effect_Console_debugShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Console_debugShow.get_or_init(||
                                                 &Func1::new(move |dictShow|
                                                                 &Func1::new({
                                                                                 let dictShow
                                                                                     =
                                                                                     dictShow.clone();
                                                                                 move
                                                                                     |a|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_debug(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                            &&&dictShow),
                                                                                                                                                         a))
                                                                             })))
    }
}
