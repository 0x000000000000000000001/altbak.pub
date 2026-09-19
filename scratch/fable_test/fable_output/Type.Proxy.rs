pub mod PureScript_Type_Proxy {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
    pub enum Type_Proxy_Proxy { Type_Proxy_Proxyusd_Ctor, }
    impl core::fmt::Display for PureScript_Type_Proxy::Type_Proxy_Proxy {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Type_Proxy_Proxy() -> &dyn Any {
        static Type_Proxy_Proxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Type_Proxy_Proxy.get_or_init(||
                                         &LrcPtr::new(PureScript_Type_Proxy::Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))
    }
}
