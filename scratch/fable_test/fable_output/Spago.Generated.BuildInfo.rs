pub mod PureScript_Spago_Generated_BuildInfo {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    pub fn Spago_Generated_BuildInfo_spagoVersion() -> &dyn Any {
        static Spago_Generated_BuildInfo_spagoVersion:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Spago_Generated_BuildInfo_spagoVersion.get_or_init(||
                                                               &string("1.0.4"))
    }
    pub fn Spago_Generated_BuildInfo_pursVersion() -> &dyn Any {
        static Spago_Generated_BuildInfo_pursVersion:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Spago_Generated_BuildInfo_pursVersion.get_or_init(||
                                                              &string("0.15.16"))
    }
    pub fn Spago_Generated_BuildInfo_packages() -> &dyn Any {
        static Spago_Generated_BuildInfo_packages: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Spago_Generated_BuildInfo_packages.get_or_init(||
                                                           &add(string("ps-cs-test"),
                                                                &&string("0.0.0"),
                                                                empty::<string,
                                                                        &dyn Any>()))
    }
}
