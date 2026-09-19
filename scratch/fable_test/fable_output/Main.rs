pub mod PureScript_Main {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    use crate::module_33161eca::PureScript_App;
    pub fn Main_main() -> &dyn Any {
        static Main_main: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Main_main.get_or_init(|| &PureScript_App::App_main())
    }
}
