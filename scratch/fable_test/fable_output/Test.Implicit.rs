pub mod PureScript_Test_Implicit {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Test_Implicit_foo() -> &dyn Any {
        static Test_Implicit_foo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Implicit_foo.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Test_Implicit_bar() -> &dyn Any {
        static Test_Implicit_bar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Implicit_bar.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Implicit::Test_Implicit_foo(),
                                                                           &&&42_i32))
    }
}
