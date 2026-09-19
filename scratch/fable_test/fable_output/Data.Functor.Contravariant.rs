pub mod PureScript_Data_Functor_Contravariant {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_a8445950::PureScript_Data_Const;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_38c1e8e1::PureScript_Data_Void;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Functor_Contravariant_Contravariantusd_Dict() -> &dyn Any {
        static Data_Functor_Contravariant_Contravariantusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Contravariant_Contravariantusd_Dict.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |x|
                                                                                         x.clone()))
    }
    pub fn Data_Functor_Contravariant_contravariantConst() -> &dyn Any {
        static Data_Functor_Contravariant_contravariantConst:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Contravariant_contravariantConst.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Contravariant::Data_Functor_Contravariant_Contravariantusd_Dict(),
                                                                                                       &&&add(string("cmap"),
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
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_Const(),
                                                                                                                                                                                        &&&matchValue_1)
                                                                                                                                                   }
                                                                                                                                           })),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))
    }
    pub fn Data_Functor_Contravariant_cmap() -> &dyn Any {
        static Data_Functor_Contravariant_cmap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Contravariant_cmap.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("cmap"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Functor_Contravariant_cmapFlipped() -> &dyn Any {
        static Data_Functor_Contravariant_cmapFlipped:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Contravariant_cmapFlipped.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictContravariant|
                                                                               &Func1::new({
                                                                                               let dictContravariant
                                                                                                   =
                                                                                                   dictContravariant.clone();
                                                                                               move
                                                                                                   |x|
                                                                                                   &Func1::new({
                                                                                                                   let x
                                                                                                                       =
                                                                                                                       x.clone();
                                                                                                                   move
                                                                                                                       |f|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Contravariant::Data_Functor_Contravariant_cmap(),
                                                                                                                                                                                                                              &&&dictContravariant),
                                                                                                                                                                                           f),
                                                                                                                                                        &&&x)
                                                                                                               })
                                                                                           })))
    }
    pub fn Data_Functor_Contravariant_coerce() -> &dyn Any {
        static Data_Functor_Contravariant_coerce: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Contravariant_coerce.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictContravariant|
                                                                          &Func1::new({
                                                                                          let dictContravariant
                                                                                              =
                                                                                              dictContravariant.clone();
                                                                                          move
                                                                                              |dictFunctor|
                                                                                              &Func1::new({
                                                                                                              let dictFunctor
                                                                                                                  =
                                                                                                                  dictFunctor.clone();
                                                                                                              move
                                                                                                                  |a|
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                         &&&dictFunctor),
                                                                                                                                                                                      &&&PureScript_Data_Void::Data_Void_absurd()),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Contravariant::Data_Functor_Contravariant_cmap(),
                                                                                                                                                                                                                                                            &&&dictContravariant),
                                                                                                                                                                                                                         &&&PureScript_Data_Void::Data_Void_absurd()),
                                                                                                                                                                                      a))
                                                                                                          })
                                                                                      })))
    }
    pub fn Data_Functor_Contravariant_imapC() -> &dyn Any {
        static Data_Functor_Contravariant_imapC: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Contravariant_imapC.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictContravariant|
                                                                         &Func1::new({
                                                                                         let dictContravariant
                                                                                             =
                                                                                             dictContravariant.clone();
                                                                                         move
                                                                                             |v|
                                                                                             &Func1::new(move
                                                                                                             |f|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Contravariant::Data_Functor_Contravariant_cmap(),
                                                                                                                                                                                 &&&dictContravariant),
                                                                                                                                              f))
                                                                                     })))
    }
}
