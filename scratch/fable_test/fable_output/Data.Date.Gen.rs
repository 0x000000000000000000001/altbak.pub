pub mod PureScript_Data_Date_Gen {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_851afbc9::PureScript_Control_Monad_Gen_Class;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_bdfbe862::PureScript_Data_Date_Component_Gen;
    use crate::module_b6aac8e0::PureScript_Data_Date_Component;
    use crate::module_d96ec2c1::PureScript_Data_Date;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_2eb6dec6::PureScript_Data_Int;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_c4b34810::PureScript_Data_Time_Duration;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Date_Gen_bottom() -> &dyn Any {
        static Data_Date_Gen_bottom: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Gen_bottom.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                              &&&PureScript_Data_Date_Component::Data_Date_Component_boundedMonth()))
    }
    pub fn Data_Date_Gen_bottom1() -> &dyn Any {
        static Data_Date_Gen_bottom1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Gen_bottom1.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                               &&&PureScript_Data_Date_Component::Data_Date_Component_boundedDay()))
    }
    pub fn Data_Date_Gen_genDate() -> &dyn Any {
        static Data_Date_Gen_genDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Gen_genDate.get_or_init(||
                                              &Func1::new(move |dictMonadGen|
                                                              {
                                                                  let Monad0 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                              Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                  let Bind1 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                              Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                  let Functor0 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                              Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                  let pure_var =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                 Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                         &&&Bind1),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date_Component_Gen::Data_Date_Component_Gen_genYear(),
                                                                                                                                                                         dictMonadGen)),
                                                                                                   &&&Func1::new({
                                                                                                                     let Bind1
                                                                                                                         =
                                                                                                                         Bind1.clone();
                                                                                                                     let Functor0
                                                                                                                         =
                                                                                                                         Functor0.clone();
                                                                                                                     let dictMonadGen
                                                                                                                         =
                                                                                                                         dictMonadGen.clone();
                                                                                                                     let pure_var
                                                                                                                         =
                                                                                                                         pure_var.clone();
                                                                                                                     move
                                                                                                                         |year|
                                                                                                                         {
                                                                                                                             let maxDays =
                                                                                                                                 {
                                                                                                                                     let matchValue =
                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_isLeapYear(),
                                                                                                                                                                                                   year));
                                                                                                                                     match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                      &matchValue)
                                                                                                                                         {
                                                                                                                                         0_i32
                                                                                                                                         =>
                                                                                                                                         &365_i32,
                                                                                                                                         _
                                                                                                                                         =>
                                                                                                                                         &364_i32,
                                                                                                                                     }
                                                                                                                                 };
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                    &&&Bind1),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Time_Duration::Data_Time_Duration_Days()),
                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Int::Data_Int_toNumber())),
                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Gen_Class::Control_Monad_Gen_Class_chooseInt(),
                                                                                                                                                                                                                                                                                                                                             &&&dictMonadGen),
                                                                                                                                                                                                                                                                                                          &&&0_i32),
                                                                                                                                                                                                                                                                       &&&maxDays))),
                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                let year
                                                                                                                                                                                    =
                                                                                                                                                                                    year.clone();
                                                                                                                                                                                move
                                                                                                                                                                                    |days|
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                        &&&pure_var),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                           &&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial()),
                                                                                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                                                                                          let days
                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                              days.clone();
                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                              |usd__unused|
                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_exactDate(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&year),
                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Date_Gen::Data_Date_Gen_bottom()),
                                                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Date_Gen::Data_Date_Gen_bottom1())),
                                                                                                                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                    |janFirst|
                                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_adjust(),
                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&days),
                                                                                                                                                                                                                                                                                                                                                                                                     janFirst))))
                                                                                                                                                                                                                                                                      })))
                                                                                                                                                                            }))
                                                                                                                         }
                                                                                                                 }))
                                                              }))
    }
}
