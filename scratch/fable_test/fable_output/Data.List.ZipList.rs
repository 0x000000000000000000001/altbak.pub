pub mod PureScript_Data_List_ZipList {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_44df2f32::PureScript_Data_List_Lazy_Types;
    use crate::module_529acc77::PureScript_Data_List_Lazy;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_List_ZipList_ZipList() -> &dyn Any {
        static Data_List_ZipList_ZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_ZipList.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  x.clone()))
    }
    pub fn Data_List_ZipList_traversableZipList() -> &dyn Any {
        static Data_List_ZipList_traversableZipList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_ZipList_traversableZipList.get_or_init(||
                                                             &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_traversableList())
    }
    pub fn Data_List_ZipList_showZipList() -> &dyn Any {
        static Data_List_ZipList_showZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_showZipList.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictShow|
                                                                      {
                                                                          let showList =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_showList(),
                                                                                                               dictShow);
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                           &&&add(string("show"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let showList
                                                                                                                                       =
                                                                                                                                       showList.clone();
                                                                                                                                   move
                                                                                                                                       |v|
                                                                                                                                       {
                                                                                                                                           let xs =
                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                               &&&string("(ZipList ")),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                        &&&showList),
                                                                                                                                                                                                                                                                                     &&&xs)),
                                                                                                                                                                                                               &&&string(")")))
                                                                                                                                       }
                                                                                                                               }),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))
                                                                      }))
    }
    pub fn Data_List_ZipList_semigroupZipList() -> &dyn Any {
        static Data_List_ZipList_semigroupZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_semigroupZipList.get_or_init(||
                                                           &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList())
    }
    pub fn Data_List_ZipList_ordZipList() -> &dyn Any {
        static Data_List_ZipList_ordZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_ordZipList.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictOrd|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_ordList(),
                                                                                                      dictOrd)))
    }
    pub fn Data_List_ZipList_newtypeZipList() -> &dyn Any {
        static Data_List_ZipList_newtypeZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_newtypeZipList.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                          &&&add(string("Coercible0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &Sharpurs_Prelude::Prim_undefined()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>())))
    }
    pub fn Data_List_ZipList_monoidZipList() -> &dyn Any {
        static Data_List_ZipList_monoidZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_monoidZipList.get_or_init(||
                                                        &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_monoidList())
    }
    pub fn Data_List_ZipList_functorZipList() -> &dyn Any {
        static Data_List_ZipList_functorZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_functorZipList.get_or_init(||
                                                         &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorList())
    }
    pub fn Data_List_ZipList_foldableZipList() -> &dyn Any {
        static Data_List_ZipList_foldableZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_foldableZipList.get_or_init(||
                                                          &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList())
    }
    pub fn Data_List_ZipList_eqZipList() -> &dyn Any {
        static Data_List_ZipList_eqZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_eqZipList.get_or_init(||
                                                    &Func1::new(move |dictEq|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eqList(),
                                                                                                     dictEq)))
    }
    pub fn Data_List_ZipList_applyZipList() -> &dyn Any {
        static Data_List_ZipList_applyZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_applyZipList.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                        &&&add(string("apply"),
                                                                                               &&Func1::new(move
                                                                                                                |v|
                                                                                                                &Func1::new({
                                                                                                                                let v
                                                                                                                                    =
                                                                                                                                    v.clone();
                                                                                                                                move
                                                                                                                                    |v1|
                                                                                                                                    {
                                                                                                                                        let matchValue =
                                                                                                                                            Sharpurs_Prelude::unbox(&&v);
                                                                                                                                        let matchValue_1 =
                                                                                                                                            Sharpurs_Prelude::unbox(v1);
                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_ZipList::Data_List_ZipList_ZipList(),
                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_zipWith(),
                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Function::Data_Function_apply()),
                                                                                                                                                                                                                                               &&&matchValue),
                                                                                                                                                                                                            &&&matchValue_1))
                                                                                                                                    }
                                                                                                                            })),
                                                                                               add(string("Functor0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_List_ZipList::Data_List_ZipList_functorZipList()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Data_List_ZipList_zipListIsNotBind() -> &dyn Any {
        static Data_List_ZipList_zipListIsNotBind: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_zipListIsNotBind.get_or_init(||
                                                           &Func1::new(move
                                                                           |usd__unused|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                                            &&&add(string("bind"),
                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafeCrashWith(),
                                                                                                                                                     &&&string("bind: unreachable")),
                                                                                                                   add(string("Apply0"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |usd__unused_1|
                                                                                                                                        &PureScript_Data_List_ZipList::Data_List_ZipList_applyZipList()),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>())))))
    }
    pub fn Data_List_ZipList_applicativeZipList() -> &dyn Any {
        static Data_List_ZipList_applicativeZipList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_ZipList_applicativeZipList.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                              &&&add(string("pure"),
                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                          &&&PureScript_Data_List_ZipList::Data_List_ZipList_ZipList()),
                                                                                                                                       &&&PureScript_Data_List_Lazy::Data_List_Lazy_repeat()),
                                                                                                     add(string("Apply0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_List_ZipList::Data_List_ZipList_applyZipList()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
    pub fn Data_List_ZipList_altZipList() -> &dyn Any {
        static Data_List_ZipList_altZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_altZipList.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                      &&&add(string("alt"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              &Func1::new({
                                                                                                                              let v
                                                                                                                                  =
                                                                                                                                  v.clone();
                                                                                                                              move
                                                                                                                                  |v1|
                                                                                                                                  {
                                                                                                                                      let matchValue =
                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                      let matchValue_1 =
                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                      let xs =
                                                                                                                                          matchValue;
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                          &&&PureScript_Data_List_ZipList::Data_List_ZipList_ZipList()),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                                                                                                                                                                             &&&xs),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_drop(),
                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_length(),
                                                                                                                                                                                                                                                                                                                   &&&xs)),
                                                                                                                                                                                                                                             &&&matchValue_1)))
                                                                                                                                  }
                                                                                                                          })),
                                                                                             add(string("Functor0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_List_ZipList::Data_List_ZipList_functorZipList()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_List_ZipList_plusZipList() -> &dyn Any {
        static Data_List_ZipList_plusZipList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ZipList_plusZipList.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                       &&&add(string("empty"),
                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                &&&PureScript_Data_List_ZipList::Data_List_ZipList_monoidZipList()),
                                                                                              add(string("Alt0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_List_ZipList::Data_List_ZipList_altZipList()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>()))))
    }
    pub fn Data_List_ZipList_alternativeZipList() -> &dyn Any {
        static Data_List_ZipList_alternativeZipList: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_ZipList_alternativeZipList.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                              &&&add(string("Applicative0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Data_List_ZipList::Data_List_ZipList_applicativeZipList()),
                                                                                                     add(string("Plus1"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused_1|
                                                                                                                          &PureScript_Data_List_ZipList::Data_List_ZipList_plusZipList()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
}
