pub mod PureScript_Test_RowToList {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub fn Test_RowToList_RecordKeysusd_Dict() -> &dyn Any {
        static Test_RowToList_RecordKeysusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToList_RecordKeysusd_Dict.get_or_init(||
                                                          &Func1::new(move |x|
                                                                          x.clone()))
    }
    pub fn Test_RowToList_keysNil() -> &dyn Any {
        static Test_RowToList_keysNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToList_keysNil.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_RecordKeysusd_Dict(),
                                                                                &&&add(string("keysImpl"),
                                                                                       &&Func1::new(move
                                                                                                        |v|
                                                                                                        &0_i32),
                                                                                       empty::<string,
                                                                                               &dyn Any>())))
    }
    pub fn Test_RowToList_keysImpl() -> &dyn Any {
        static Test_RowToList_keysImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToList_keysImpl.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("keysImpl"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Test_RowToList_keysCons() -> &dyn Any {
        static Test_RowToList_keysCons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToList_keysCons.get_or_init(||
                                                &Func1::new(move
                                                                |dictRecordKeys|
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_RecordKeysusd_Dict(),
                                                                                                 &&&add(string("keysImpl"),
                                                                                                        &&Func1::new({
                                                                                                                         let dictRecordKeys
                                                                                                                             =
                                                                                                                             dictRecordKeys.clone();
                                                                                                                         move
                                                                                                                             |v|
                                                                                                                             &(Sharpurs_Prelude::unbox(&&&1_i32)
                                                                                                                                   +
                                                                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_keysImpl(),
                                                                                                                                                                                                                                 &&&dictRecordKeys),
                                                                                                                                                                                              &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))))
                                                                                                                     }),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))))
    }
    pub fn Test_RowToList_keysCons1() -> &dyn Any {
        static Test_RowToList_keysCons1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToList_keysCons1.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_keysCons(),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_keysCons(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_keysCons(),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_keysCons(),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_keysCons(),
                                                                                                                                                                                                                              &&&PureScript_Test_RowToList::Test_RowToList_keysNil()))))))
    }
    pub fn Test_RowToList_keys() -> &dyn Any {
        static Test_RowToList_keys: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToList_keys.get_or_init(||
                                            &Func1::new(move |usd__unused|
                                                            &Func1::new(move
                                                                            |dictRecordKeys|
                                                                            &Func1::new({
                                                                                            let dictRecordKeys
                                                                                                =
                                                                                                dictRecordKeys.clone();
                                                                                            move
                                                                                                |v|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_keysImpl(),
                                                                                                                                                                    &&&dictRecordKeys),
                                                                                                                                 &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))
                                                                                        }))))
    }
    pub fn Test_RowToList_describe() -> &dyn Any {
        static Test_RowToList_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToList_describe.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                 &&&string("RowToList (Keys Count):")))
    }
    pub fn Test_RowToList_act() -> &dyn Any {
        static Test_RowToList_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToList_act.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                  &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                  &&&10000_i32)),
                                                                            &&&Func1::new(move
                                                                                              |usd__unused|
                                                                                              {
                                                                                                  let rec_var =
                                                                                                      &add(string("a"),
                                                                                                           &&1_i32,
                                                                                                           add(string("b"),
                                                                                                               &&string("two"),
                                                                                                               add(string("c"),
                                                                                                                   &&true,
                                                                                                                   add(string("d"),
                                                                                                                       &&4.0_f64,
                                                                                                                       add(string("e"),
                                                                                                                           &&string("five"),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>())))));
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                      &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                         &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToList::Test_RowToList_keys(),
                                                                                                                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                            &&&PureScript_Test_RowToList::Test_RowToList_keysCons1()),
                                                                                                                                                                                                         &&&rec_var)))
                                                                                              })))
    }
}
