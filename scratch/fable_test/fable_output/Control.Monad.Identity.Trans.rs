pub mod PureScript_Control_Monad_Identity_Trans {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_f5fe307f::PureScript_Control_Monad_Trans_Class;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Identity_Trans_IdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_IdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_IdentityT.get_or_init(||
                                                               &Func1::new(move
                                                                               |x|
                                                                               x.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadSTIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadSTIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadSTIdentityT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictMonadST|
                                                                                      dictMonadST.clone()))
    }
    pub fn Control_Monad_Identity_Trans_traversableIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_traversableIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_traversableIdentityT.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |dictTraversable|
                                                                                          dictTraversable.clone()))
    }
    pub fn Control_Monad_Identity_Trans_runIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_runIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_runIdentityT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |v|
                                                                                  &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Control_Monad_Identity_Trans_plusIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_plusIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_plusIdentityT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictPlus|
                                                                                   dictPlus.clone()))
    }
    pub fn Control_Monad_Identity_Trans_newtypeIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_newtypeIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_newtypeIdentityT.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                       &&&add(string("Coercible0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &Sharpurs_Prelude::Prim_undefined()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))
    }
    pub fn Control_Monad_Identity_Trans_monadWriterIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadWriterIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadWriterIdentityT.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |dictMonadWriter|
                                                                                          dictMonadWriter.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadTransIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadTransIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadTransIdentityT.get_or_init(||
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Trans_Class::Control_Monad_Trans_Class_MonadTransusd_Dict(),
                                                                                                          &&&add(string("lift"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |dictMonad|
                                                                                                                                  &PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_IdentityT()),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>())))
    }
    pub fn Control_Monad_Identity_Trans_monadThrowIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadThrowIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadThrowIdentityT.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |dictMonadThrow|
                                                                                         dictMonadThrow.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadTellIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadTellIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadTellIdentityT.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictMonadTell|
                                                                                        dictMonadTell.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadStateIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadStateIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadStateIdentityT.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |dictMonadState|
                                                                                         dictMonadState.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadRecIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadRecIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadRecIdentityT.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictMonadRec|
                                                                                       dictMonadRec.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadReaderIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadReaderIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadReaderIdentityT.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |dictMonadReader|
                                                                                          dictMonadReader.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadPlusIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadPlusIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadPlusIdentityT.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictMonadPlus|
                                                                                        dictMonadPlus.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadIdentityT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictMonad|
                                                                                    dictMonad.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadErrorIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadErrorIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadErrorIdentityT.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |dictMonadError|
                                                                                         dictMonadError.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadEffectIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadEffectIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadEffectIdentityT.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |dictMonadEffect|
                                                                                          dictMonadEffect.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadContIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadContIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadContIdentityT.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictMonadCont|
                                                                                        dictMonadCont.clone()))
    }
    pub fn Control_Monad_Identity_Trans_monadAskIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_monadAskIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_monadAskIdentityT.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictMonadAsk|
                                                                                       dictMonadAsk.clone()))
    }
    pub fn Control_Monad_Identity_Trans_mapIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_mapIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_mapIdentityT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |f|
                                                                                  &Func1::new({
                                                                                                  let f
                                                                                                      =
                                                                                                      f.clone();
                                                                                                  move
                                                                                                      |v|
                                                                                                      {
                                                                                                          let matchValue =
                                                                                                              Sharpurs_Prelude::unbox(&&f);
                                                                                                          let matchValue_1 =
                                                                                                              Sharpurs_Prelude::unbox(v);
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_IdentityT(),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                              &&&matchValue_1))
                                                                                                      }
                                                                                              })))
    }
    pub fn Control_Monad_Identity_Trans_functorIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_functorIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_functorIdentityT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictFunctor|
                                                                                      dictFunctor.clone()))
    }
    pub fn Control_Monad_Identity_Trans_foldableIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_foldableIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_foldableIdentityT.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictFoldable|
                                                                                       dictFoldable.clone()))
    }
    pub fn Control_Monad_Identity_Trans_extendIdentityI() -> &dyn Any {
        static Control_Monad_Identity_Trans_extendIdentityI:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_extendIdentityI.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictExtend|
                                                                                     {
                                                                                         let functorIdentityT1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_functorIdentityT(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                        Sharpurs_Prelude::unbox(dictExtend)),
                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                                                          &&&add(string("extend"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let dictExtend
                                                                                                                                                      =
                                                                                                                                                      dictExtend.clone();
                                                                                                                                                  move
                                                                                                                                                      |f|
                                                                                                                                                      &Func1::new({
                                                                                                                                                                      let f
                                                                                                                                                                          =
                                                                                                                                                                          f.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |v|
                                                                                                                                                                          {
                                                                                                                                                                              let matchValue =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                              let matchValue_1 =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_IdentityT(),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                                                                                                                        &&&dictExtend),
                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_IdentityT())),
                                                                                                                                                                                                                                                  &&&matchValue_1))
                                                                                                                                                                          }
                                                                                                                                                                  })
                                                                                                                                              }),
                                                                                                                                 add(string("Functor0"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let functorIdentityT1
                                                                                                                                                          =
                                                                                                                                                          functorIdentityT1.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused|
                                                                                                                                                          &functorIdentityT1
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
                                                                                     }))
    }
    pub fn Control_Monad_Identity_Trans_eqIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_eqIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_eqIdentityT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictEq1|
                                                                                 &Func1::new({
                                                                                                 let dictEq1
                                                                                                     =
                                                                                                     dictEq1.clone();
                                                                                                 move
                                                                                                     |dictEq|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                                                      &&&add(string("eq"),
                                                                                                                                             &&Func1::new({
                                                                                                                                                              let dictEq
                                                                                                                                                                  =
                                                                                                                                                                  dictEq.clone();
                                                                                                                                                              move
                                                                                                                                                                  |x|
                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                  let x
                                                                                                                                                                                      =
                                                                                                                                                                                      x.clone();
                                                                                                                                                                                  move
                                                                                                                                                                                      |y|
                                                                                                                                                                                      {
                                                                                                                                                                                          let matchValue =
                                                                                                                                                                                              Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                                          let matchValue_1 =
                                                                                                                                                                                              Sharpurs_Prelude::unbox(y);
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                                                                                                                                                                                                                    &&&dictEq1),
                                                                                                                                                                                                                                                                                                 &&&dictEq),
                                                                                                                                                                                                                                                              &&&matchValue),
                                                                                                                                                                                                                           &&&matchValue_1)
                                                                                                                                                                                      }
                                                                                                                                                                              })
                                                                                                                                                          }),
                                                                                                                                             empty::<string,
                                                                                                                                                     &dyn Any>()))
                                                                                             })))
    }
    pub fn Control_Monad_Identity_Trans_ordIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_ordIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_ordIdentityT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictOrd1|
                                                                                  {
                                                                                      let eqIdentityT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_eqIdentityT(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                                     Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                      &Func1::new({
                                                                                                      let dictOrd1
                                                                                                          =
                                                                                                          dictOrd1.clone();
                                                                                                      let eqIdentityT1
                                                                                                          =
                                                                                                          eqIdentityT1.clone();
                                                                                                      move
                                                                                                          |dictOrd|
                                                                                                          {
                                                                                                              let eqIdentityT2 =
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&eqIdentityT1,
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                                             Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                                               &&&add(string("compare"),
                                                                                                                                                      &&Func1::new({
                                                                                                                                                                       let dictOrd
                                                                                                                                                                           =
                                                                                                                                                                           dictOrd.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |x|
                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                           let x
                                                                                                                                                                                               =
                                                                                                                                                                                               x.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |y|
                                                                                                                                                                                               {
                                                                                                                                                                                                   let matchValue =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                                                   let matchValue_1 =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(y);
                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                                                                                                                                                                                                                             &&&dictOrd1),
                                                                                                                                                                                                                                                                                                          &&&dictOrd),
                                                                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                                                                    &&&matchValue_1)
                                                                                                                                                                                               }
                                                                                                                                                                                       })
                                                                                                                                                                   }),
                                                                                                                                                      add(string("Eq0"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let eqIdentityT2
                                                                                                                                                                               =
                                                                                                                                                                               eqIdentityT2.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |usd__unused|
                                                                                                                                                                               &eqIdentityT2
                                                                                                                                                                       }),
                                                                                                                                                          empty::<string,
                                                                                                                                                                  &dyn Any>())))
                                                                                                          }
                                                                                                  })
                                                                                  }))
    }
    pub fn Control_Monad_Identity_Trans_eq1IdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_eq1IdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_eq1IdentityT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictEq1|
                                                                                  {
                                                                                      let eqIdentityT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_eqIdentityT(),
                                                                                                                           dictEq1);
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                                                       &&&add(string("eq1"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let eqIdentityT1
                                                                                                                                                   =
                                                                                                                                                   eqIdentityT1.clone();
                                                                                                                                               move
                                                                                                                                                   |dictEq|
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&eqIdentityT1,
                                                                                                                                                                                                                       dictEq))
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>()))
                                                                                  }))
    }
    pub fn Control_Monad_Identity_Trans_ord1IdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_ord1IdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_ord1IdentityT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictOrd1|
                                                                                   {
                                                                                       let ordIdentityT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_ordIdentityT(),
                                                                                                                            dictOrd1);
                                                                                       let eq1IdentityT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_eq1IdentityT(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                                                        &&&add(string("compare1"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let ordIdentityT1
                                                                                                                                                    =
                                                                                                                                                    ordIdentityT1.clone();
                                                                                                                                                move
                                                                                                                                                    |dictOrd|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&ordIdentityT1,
                                                                                                                                                                                                                        dictOrd))
                                                                                                                                            }),
                                                                                                                               add(string("Eq10"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let eq1IdentityT1
                                                                                                                                                        =
                                                                                                                                                        eq1IdentityT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &eq1IdentityT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Monad_Identity_Trans_comonadIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_comonadIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_comonadIdentityT.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictComonad|
                                                                                      {
                                                                                          let extendIdentityI1 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_extendIdentityI(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictComonad)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_Comonadusd_Dict(),
                                                                                                                           &&&add(string("extract"),
                                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                                          dictComonad)),
                                                                                                                                                                    &&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_runIdentityT()),
                                                                                                                                  add(string("Extend0"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let extendIdentityI1
                                                                                                                                                           =
                                                                                                                                                           extendIdentityI1.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused|
                                                                                                                                                           &extendIdentityI1
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))
                                                                                      }))
    }
    pub fn Control_Monad_Identity_Trans_bindIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_bindIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_bindIdentityT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictBind|
                                                                                   dictBind.clone()))
    }
    pub fn Control_Monad_Identity_Trans_applyIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_applyIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_applyIdentityT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictApply|
                                                                                    dictApply.clone()))
    }
    pub fn Control_Monad_Identity_Trans_applicativeIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_applicativeIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_applicativeIdentityT.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |dictApplicative|
                                                                                          dictApplicative.clone()))
    }
    pub fn Control_Monad_Identity_Trans_alternativeIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_alternativeIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_alternativeIdentityT.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |dictAlternative|
                                                                                          dictAlternative.clone()))
    }
    pub fn Control_Monad_Identity_Trans_altIdentityT() -> &dyn Any {
        static Control_Monad_Identity_Trans_altIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Identity_Trans_altIdentityT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictAlt|
                                                                                  dictAlt.clone()))
    }
}
