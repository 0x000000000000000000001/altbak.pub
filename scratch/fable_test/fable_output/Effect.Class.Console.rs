pub mod PureScript_Effect_Class_Console {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_21d6b3bc::PureScript_Effect_Class;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Effect_Class_Console_warnShow() -> &dyn Any {
        static Effect_Class_Console_warnShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_warnShow.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonadEffect|
                                                                      {
                                                                          let liftEffect =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                               dictMonadEffect);
                                                                          &Func1::new({
                                                                                          let liftEffect
                                                                                              =
                                                                                              liftEffect.clone();
                                                                                          move
                                                                                              |dictShow|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                  &&&liftEffect),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_warnShow(),
                                                                                                                                                                  dictShow))
                                                                                      })
                                                                      }))
    }
    pub fn Effect_Class_Console_warn() -> &dyn Any {
        static Effect_Class_Console_warn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_warn.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictMonadEffect|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                         dictMonadEffect)),
                                                                                                   &&&PureScript_Effect_Console::Effect_Console_warn())))
    }
    pub fn Effect_Class_Console_timeLog() -> &dyn Any {
        static Effect_Class_Console_timeLog: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_timeLog.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictMonadEffect|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                            dictMonadEffect)),
                                                                                                      &&&PureScript_Effect_Console::Effect_Console_timeLog())))
    }
    pub fn Effect_Class_Console_timeEnd() -> &dyn Any {
        static Effect_Class_Console_timeEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_timeEnd.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictMonadEffect|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                            dictMonadEffect)),
                                                                                                      &&&PureScript_Effect_Console::Effect_Console_timeEnd())))
    }
    pub fn Effect_Class_Console_time() -> &dyn Any {
        static Effect_Class_Console_time: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_time.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictMonadEffect|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                         dictMonadEffect)),
                                                                                                   &&&PureScript_Effect_Console::Effect_Console_time())))
    }
    pub fn Effect_Class_Console_logShow() -> &dyn Any {
        static Effect_Class_Console_logShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_logShow.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictMonadEffect|
                                                                     {
                                                                         let liftEffect =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                              dictMonadEffect);
                                                                         &Func1::new({
                                                                                         let liftEffect
                                                                                             =
                                                                                             liftEffect.clone();
                                                                                         move
                                                                                             |dictShow|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                 &&&liftEffect),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_logShow(),
                                                                                                                                                                 dictShow))
                                                                                     })
                                                                     }))
    }
    pub fn Effect_Class_Console_log() -> &dyn Any {
        static Effect_Class_Console_log: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_log.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictMonadEffect|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                        dictMonadEffect)),
                                                                                                  &&&PureScript_Effect_Console::Effect_Console_log())))
    }
    pub fn Effect_Class_Console_infoShow() -> &dyn Any {
        static Effect_Class_Console_infoShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_infoShow.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonadEffect|
                                                                      {
                                                                          let liftEffect =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                               dictMonadEffect);
                                                                          &Func1::new({
                                                                                          let liftEffect
                                                                                              =
                                                                                              liftEffect.clone();
                                                                                          move
                                                                                              |dictShow|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                  &&&liftEffect),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_infoShow(),
                                                                                                                                                                  dictShow))
                                                                                      })
                                                                      }))
    }
    pub fn Effect_Class_Console_info() -> &dyn Any {
        static Effect_Class_Console_info: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_info.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictMonadEffect|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                         dictMonadEffect)),
                                                                                                   &&&PureScript_Effect_Console::Effect_Console_info())))
    }
    pub fn Effect_Class_Console_groupEnd() -> &dyn Any {
        static Effect_Class_Console_groupEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_groupEnd.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonadEffect|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                          dictMonadEffect),
                                                                                                       &&&PureScript_Effect_Console::Effect_Console_groupEnd())))
    }
    pub fn Effect_Class_Console_groupCollapsed() -> &dyn Any {
        static Effect_Class_Console_groupCollapsed: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Class_Console_groupCollapsed.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictMonadEffect|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                                   dictMonadEffect)),
                                                                                                             &&&PureScript_Effect_Console::Effect_Console_groupCollapsed())))
    }
    pub fn Effect_Class_Console_group() -> &dyn Any {
        static Effect_Class_Console_group: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_group.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonadEffect|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                          dictMonadEffect)),
                                                                                                    &&&PureScript_Effect_Console::Effect_Console_group())))
    }
    pub fn Effect_Class_Console_grouped() -> &dyn Any {
        static Effect_Class_Console_grouped: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_grouped.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictMonadEffect|
                                                                     {
                                                                         let Monad0 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictMonadEffect)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                         let Bind1 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                     Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                         let groupEnd1 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class_Console::Effect_Class_Console_groupEnd(),
                                                                                                              dictMonadEffect);
                                                                         let Applicative0 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                     Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                         &Func1::new({
                                                                                         let Applicative0
                                                                                             =
                                                                                             Applicative0.clone();
                                                                                         let Bind1
                                                                                             =
                                                                                             Bind1.clone();
                                                                                         let dictMonadEffect
                                                                                             =
                                                                                             dictMonadEffect.clone();
                                                                                         let groupEnd1
                                                                                             =
                                                                                             groupEnd1.clone();
                                                                                         move
                                                                                             |name|
                                                                                             &Func1::new({
                                                                                                             let name
                                                                                                                 =
                                                                                                                 name.clone();
                                                                                                             move
                                                                                                                 |inner|
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                           &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                        &&&Bind1),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class_Console::Effect_Class_Console_group(),
                                                                                                                                                                                                                                                           &&&dictMonadEffect),
                                                                                                                                                                                                                        &&&name)),
                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                    let inner
                                                                                                                                                                        =
                                                                                                                                                                        inner.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |usd__unused|
                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                               &&&Bind1),
                                                                                                                                                                                                                                            &&&inner),
                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                           |result|
                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                                                  &&&Bind1),
                                                                                                                                                                                                                                                                                               &&&groupEnd1),
                                                                                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                                                                                              let result
                                                                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                                                                  result.clone();
                                                                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                                                                  |usd__unused_1|
                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                                                      &&&Applicative0),
                                                                                                                                                                                                                                                                                                                   &&&result)
                                                                                                                                                                                                                                                                          }))))
                                                                                                                                                                }))
                                                                                                         })
                                                                                     })
                                                                     }))
    }
    pub fn Effect_Class_Console_errorShow() -> &dyn Any {
        static Effect_Class_Console_errorShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_errorShow.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadEffect|
                                                                       {
                                                                           let liftEffect =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                dictMonadEffect);
                                                                           &Func1::new({
                                                                                           let liftEffect
                                                                                               =
                                                                                               liftEffect.clone();
                                                                                           move
                                                                                               |dictShow|
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                   &&&liftEffect),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_errorShow(),
                                                                                                                                                                   dictShow))
                                                                                       })
                                                                       }))
    }
    pub fn Effect_Class_Console_error() -> &dyn Any {
        static Effect_Class_Console_error: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_error.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonadEffect|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                          dictMonadEffect)),
                                                                                                    &&&PureScript_Effect_Console::Effect_Console_error())))
    }
    pub fn Effect_Class_Console_debugShow() -> &dyn Any {
        static Effect_Class_Console_debugShow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_debugShow.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadEffect|
                                                                       {
                                                                           let liftEffect =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                dictMonadEffect);
                                                                           &Func1::new({
                                                                                           let liftEffect
                                                                                               =
                                                                                               liftEffect.clone();
                                                                                           move
                                                                                               |dictShow|
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                   &&&liftEffect),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_debugShow(),
                                                                                                                                                                   dictShow))
                                                                                       })
                                                                       }))
    }
    pub fn Effect_Class_Console_debug() -> &dyn Any {
        static Effect_Class_Console_debug: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_debug.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonadEffect|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                                                          dictMonadEffect)),
                                                                                                    &&&PureScript_Effect_Console::Effect_Console_debug())))
    }
    pub fn Effect_Class_Console_clear() -> &dyn Any {
        static Effect_Class_Console_clear: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_Console_clear.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonadEffect|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_liftEffect(),
                                                                                                                                       dictMonadEffect),
                                                                                                    &&&PureScript_Effect_Console::Effect_Console_clear())))
    }
}
