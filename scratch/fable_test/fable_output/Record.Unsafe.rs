pub mod PureScript_Record_Unsafe {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    pub mod Record_Unsafe_FFI {
        use super::*;
        use fable_library_rust::Map_::add;
        use fable_library_rust::Map_::containsKey;
        use fable_library_rust::Map_::find;
        use fable_library_rust::Map_::remove;
        use fable_library_rust::Native_::Func1;
        use crate::module_aa21d1e7::Sharpurs_Prelude;
        pub fn unsafeHas() -> &dyn Any {
            static unsafeHas: MutCell<Option<&dyn Any>> = MutCell::new(None);
            unsafeHas.get_or_init(||
                                      &Func1::new(move |k|
                                                      &Func1::new({
                                                                      let k =
                                                                          k.clone();
                                                                      move
                                                                          |map|
                                                                          &containsKey(Sharpurs_Prelude::unbox(&k),
                                                                                       Sharpurs_Prelude::unbox(map))
                                                                  })))
        }
        pub fn unsafeGet() -> &dyn Any {
            static unsafeGet: MutCell<Option<&dyn Any>> = MutCell::new(None);
            unsafeGet.get_or_init(||
                                      &Func1::new(move |k|
                                                      &Func1::new({
                                                                      let k =
                                                                          k.clone();
                                                                      move
                                                                          |map|
                                                                          find(Sharpurs_Prelude::unbox(&k),
                                                                               Sharpurs_Prelude::unbox(map))
                                                                  })))
        }
        pub fn unsafeSet() -> &dyn Any {
            static unsafeSet: MutCell<Option<&dyn Any>> = MutCell::new(None);
            unsafeSet.get_or_init(||
                                      &Func1::new(move |k|
                                                      &Func1::new({
                                                                      let k =
                                                                          k.clone();
                                                                      move |v|
                                                                          &Func1::new({
                                                                                          let v
                                                                                              =
                                                                                              v.clone();
                                                                                          move
                                                                                              |map|
                                                                                              &add(Sharpurs_Prelude::unbox(&k),
                                                                                                   v,
                                                                                                   Sharpurs_Prelude::unbox(map))
                                                                                      })
                                                                  })))
        }
        pub fn unsafeDelete() -> &dyn Any {
            static unsafeDelete: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            unsafeDelete.get_or_init(||
                                         &Func1::new(move |k|
                                                         &Func1::new({
                                                                         let k
                                                                             =
                                                                             k.clone();
                                                                         move
                                                                             |map|
                                                                             &remove(Sharpurs_Prelude::unbox(&k),
                                                                                     Sharpurs_Prelude::unbox(map))
                                                                     })))
        }
    }
    pub fn Record_Unsafe_unsafeDelete() -> &dyn Any {
        static Record_Unsafe_unsafeDelete: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Record_Unsafe_unsafeDelete.get_or_init(||
                                                   &PureScript_Record_Unsafe::Record_Unsafe_FFI::unsafeDelete())
    }
    pub fn Record_Unsafe_unsafeGet() -> &dyn Any {
        static Record_Unsafe_unsafeGet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Record_Unsafe_unsafeGet.get_or_init(||
                                                &PureScript_Record_Unsafe::Record_Unsafe_FFI::unsafeGet())
    }
    pub fn Record_Unsafe_unsafeHas() -> &dyn Any {
        static Record_Unsafe_unsafeHas: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Record_Unsafe_unsafeHas.get_or_init(||
                                                &PureScript_Record_Unsafe::Record_Unsafe_FFI::unsafeHas())
    }
    pub fn Record_Unsafe_unsafeSet() -> &dyn Any {
        static Record_Unsafe_unsafeSet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Record_Unsafe_unsafeSet.get_or_init(||
                                                &PureScript_Record_Unsafe::Record_Unsafe_FFI::unsafeSet())
    }
}
