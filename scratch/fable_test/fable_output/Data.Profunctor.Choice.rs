pub mod PureScript_Data_Profunctor_Choice {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Profunctor_Choice_identity() -> &dyn Any {
        static Data_Profunctor_Choice_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Choice_identity.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                         &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Profunctor_Choice_Choiceusd_Dict() -> &dyn Any {
        static Data_Profunctor_Choice_Choiceusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Choice_Choiceusd_Dict.get_or_init(||
                                                              &Func1::new(move
                                                                              |x|
                                                                              x.clone()))
    }
    pub fn Data_Profunctor_Choice_right() -> &dyn Any {
        static Data_Profunctor_Choice_right: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Choice_right.get_or_init(||
                                                     &Func1::new(move |dict|
                                                                     find(string("right"),
                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Profunctor_Choice_left() -> &dyn Any {
        static Data_Profunctor_Choice_left: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Choice_left.get_or_init(||
                                                    &Func1::new(move |dict|
                                                                    find(string("left"),
                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Profunctor_Choice_splitChoice() -> &dyn Any {
        static Data_Profunctor_Choice_splitChoice: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Choice_splitChoice.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictSemigroupoid|
                                                                           &Func1::new({
                                                                                           let dictSemigroupoid
                                                                                               =
                                                                                               dictSemigroupoid.clone();
                                                                                           move
                                                                                               |dictChoice|
                                                                                               &Func1::new({
                                                                                                               let dictChoice
                                                                                                                   =
                                                                                                                   dictChoice.clone();
                                                                                                               move
                                                                                                                   |l|
                                                                                                                   &Func1::new({
                                                                                                                                   let l
                                                                                                                                       =
                                                                                                                                       l.clone();
                                                                                                                                   move
                                                                                                                                       |r|
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                              &&&dictSemigroupoid),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Choice::Data_Profunctor_Choice_left(),
                                                                                                                                                                                                                                                                                 &&&dictChoice),
                                                                                                                                                                                                                                              &&&l)),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Choice::Data_Profunctor_Choice_right(),
                                                                                                                                                                                                                                              &&&dictChoice),
                                                                                                                                                                                                           r))
                                                                                                                               })
                                                                                                           })
                                                                                       })))
    }
    pub fn Data_Profunctor_Choice_fanin() -> &dyn Any {
        static Data_Profunctor_Choice_fanin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Choice_fanin.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictSemigroupoid|
                                                                     &Func1::new({
                                                                                     let dictSemigroupoid
                                                                                         =
                                                                                         dictSemigroupoid.clone();
                                                                                     move
                                                                                         |dictChoice|
                                                                                         {
                                                                                             let Profunctor0 =
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&find(string("Profunctor0"),
                                                                                                                                         Sharpurs_Prelude::unbox(dictChoice)),
                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined());
                                                                                             &Func1::new({
                                                                                                             let Profunctor0
                                                                                                                 =
                                                                                                                 Profunctor0.clone();
                                                                                                             let dictChoice
                                                                                                                 =
                                                                                                                 dictChoice.clone();
                                                                                                             move
                                                                                                                 |l|
                                                                                                                 &Func1::new({
                                                                                                                                 let l
                                                                                                                                     =
                                                                                                                                     l.clone();
                                                                                                                                 move
                                                                                                                                     |r|
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_rmap(),
                                                                                                                                                                                                                                            &&&Profunctor0),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                                                                                                                                               &&&PureScript_Data_Profunctor_Choice::Data_Profunctor_Choice_identity()),
                                                                                                                                                                                                                                            &&&PureScript_Data_Profunctor_Choice::Data_Profunctor_Choice_identity())),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Choice::Data_Profunctor_Choice_splitChoice(),
                                                                                                                                                                                                                                                                                                                  &&&dictSemigroupoid),
                                                                                                                                                                                                                                                                               &&&dictChoice),
                                                                                                                                                                                                                                            &&&l),
                                                                                                                                                                                                         r))
                                                                                                                             })
                                                                                                         })
                                                                                         }
                                                                                 })))
    }
    pub fn Data_Profunctor_Choice_choiceFn() -> &dyn Any {
        static Data_Profunctor_Choice_choiceFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Choice_choiceFn.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Choice::Data_Profunctor_Choice_Choiceusd_Dict(),
                                                                                         &&&add(string("left"),
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
                                                                                                                                         let matchValue_1:
                                                                                                                                                 LrcPtr<Data_Either_Either> =
                                                                                                                                             Sharpurs_Prelude::unbox(v1);
                                                                                                                                         match matchValue_1.as_ref()
                                                                                                                                             {
                                                                                                                                             Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)
                                                                                                                                             =>
                                                                                                                                             &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)),
                                                                                                                                             Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                                                                             =>
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                   |usd__arg1|
                                                                                                                                                                                                                                   &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                 &&matchValue_1_0_0)),
                                                                                                                                         }
                                                                                                                                     }
                                                                                                                             })),
                                                                                                add(string("right"),
                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                      &&&PureScript_Data_Either::Data_Either_functorEither()),
                                                                                                    add(string("Profunctor0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_Profunctor::Data_Profunctor_profunctorFn()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))))
    }
}
