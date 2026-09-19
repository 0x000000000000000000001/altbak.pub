pub mod PureScript_Test_Main {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Test_Main_main() -> &dyn Any {
        static Test_Main_main: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Test_Main_main.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                           &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                        &&&PureScript_Data_Unit::Data_Unit_unit()))
    }
}
